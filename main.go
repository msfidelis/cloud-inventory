package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ryanuber/columnize"

	"encoding/csv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

// A Resource to export on reports
type Resource struct {
	Name    string // resource name - this value come from tag:Name
	Arn     string // resource ARN
	Region  string // resource Region
	Service string // resource Service. This is parsed by aws/arn
}

func main() {
	tagName := flag.String("tag-name", "", "Tag to search")
	tagValue := flag.String("tag-value", "", "Tag to search")
	region := flag.String("region", "us-east-1", "Region to search inventory; default: us-east-1")
	resourceType := flag.String("resource", "", "Optional resource type; ex: ec2, s3, acm")
	outputFormat := flag.String("output", "default", "Output report type; ex: default, arn, csv")
	pipeGrep := flag.String("grep", "", "String pattern to search on Tag:Name and ARN resource")

	flag.Parse()

	resources := getResources(*tagName, *tagValue, *region, *resourceType, *pipeGrep)

	output := createOutput(resources, *outputFormat)

	fmt.Println(output)

	fmt.Printf("\nFound %v resources\n", len(resources))

}

func getResources(tagKey string, tagValue string, region string, resourceType string, stringPattern string) map[string]Resource {

	fmt.Printf("\nSearching for resources using Tag %v:%v\n\n", tagKey, tagValue)

	cloudResources := make(map[string]Resource)
	cloudResourcesFiltered := make(map[string]Resource)

	items := int64(100)
	sess, err := getAWSSession(region)

	if err != nil {
		fmt.Errorf("Unable to initialize AWS session: %v", err)
	}

	svr := resourcegroupstaggingapi.New(sess)

	// Pagination Loop
	var token *string

	for {

		tagFilters := &resourcegroupstaggingapi.TagFilter{}
		tagFilters.Key = aws.String(tagKey)
		tagFilters.Values = append(tagFilters.Values, aws.String(tagValue))

		getResourcesInput := &resourcegroupstaggingapi.GetResourcesInput{}
		getResourcesInput.TagFilters = append(getResourcesInput.TagFilters, tagFilters)

		getResourcesInput.ResourcesPerPage = &items
		getResourcesInput.PaginationToken = token

		if resourceType != "" {
			getResourcesInput.ResourceTypeFilters = []*string{
				aws.String(resourceType),
			}
		}

		resources, err := svr.GetResources(getResourcesInput)
		if err != nil {
			fmt.Errorf("Unable to initialize AWS session: %v", err)
		}

		for _, resource := range resources.ResourceTagMappingList {

			name := findNameTag(resource.Tags)
			arnLong := *resource.ResourceARN

			arnInfos, _ := arn.Parse(arnLong)

			cloudResources[arnLong] = Resource{
				Name:    name,
				Arn:     arnLong,
				Region:  region,
				Service: arnInfos.Service,
			}

		}

		token = resources.PaginationToken

		if token == nil || *token == "" {
			break
		}
	}

	// Return cloud resources if not string pattern is set
	if stringPattern == "" {
		return cloudResources
	}

	// Filter string pattern

	cloudResourcesFiltered = filterResources(cloudResources, stringPattern)

	return cloudResourcesFiltered

}

func findNameTag(tags []*resourcegroupstaggingapi.Tag) string {
	for _, tag := range tags {
		if strings.ToLower(*tag.Key) == "name" {
			return *tag.Value
		}
	}
	return "-"
}

func getAWSSession(region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true)
	return session.NewSession(awsConfig)
}

func createOutput(items map[string]Resource, format string) string {

	var output string

	switch strings.ToLower(format) {
	case "default":
		output = createDefaultOutput(items)
		break
	case "arn":
		output = createArnOutput(items)
		break
	case "csv":
		output = createCsvOutput(items)
		break
	default:
		output = createDefaultOutput(items)
	}

	return output
}

func createDefaultOutput(items map[string]Resource) string {

	output := []string{
		"Tag:Name | ARN | Region | Service",
	}

	for _, item := range items {
		output = append(output, fmt.Sprintf("%s | %s | %s | %s", item.Name, item.Arn, item.Region, item.Service))
	}

	result := columnize.SimpleFormat(output)

	return result
}

func createArnOutput(items map[string]Resource) string {

	output := []string{}

	for _, item := range items {
		output = append(output, item.Arn)
	}

	result := columnize.SimpleFormat(output)

	return result
}

func createCsvOutput(items map[string]Resource) string {

	var data = [][]string{}

	file, err := os.Create("results.csv")
	defer file.Close()

	if err != nil {
		fmt.Errorf("Cannot create file: %v", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range items {
		var output = []string{}
		output = append(output, item.Name)
		output = append(output, item.Arn)
		output = append(output, item.Region)
		output = append(output, item.Service)

		data = append(data, output)
	}

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			fmt.Errorf("Cannot write to file: %v", err)
		}
	}

	return "Output file: results.csv"
}

func filterResources(resources map[string]Resource, pattern string) map[string]Resource {

	var filtered = make(map[string]Resource)

	if pattern == "" {
		return resources
	}

	fmt.Printf("Grepping by %s pattern on resources\n\n", pattern)

	for arn, resource := range resources {

		matchArn, err := regexp.MatchString(pattern, resource.Arn)

		if err != nil {
			fmt.Errorf("Error to parsing regex on ARN %v", err)
		}

		matchArnName, err := regexp.MatchString(pattern, resource.Name)

		if err != nil {
			fmt.Errorf("Error to parsing regex on Tag:Name %v", err)
		}

		if matchArn == true || matchArnName == true {
			filtered[arn] = resource
		}
	}

	return filtered

}
