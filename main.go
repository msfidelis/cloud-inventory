package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ryanuber/columnize"

	"encoding/csv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

type Resource struct {
	Name   string
	Arn    string
	Region string
}

func main() {
	tag_name := flag.String("tag-name", "", "Tag to search")
	tag_value := flag.String("tag-value", "", "Tag to search")
	region := flag.String("region", "us-east-1", "region to search inventory; default: us-east-1")
	resource_type := flag.String("resource", "", "Resource type; ex: ec2, s3, acm")
	output_format := flag.String("output", "default", "Resource type; ex: default, arn, csv")
	flag.Parse()

	resources := getResources(*tag_name, *tag_value, *region, *resource_type)

	output := createOutput(resources, *output_format)

	fmt.Println(output)
}

func getResources(tag_key string, tag_value string, region string, resource_type string) map[string]Resource {

	fmt.Printf("Searching for resources using Tag %v:%v\n", tag_key, tag_value)

	cloud_resources := make(map[string]Resource)
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
		tagFilters.Key = aws.String(tag_key)
		tagFilters.Values = append(tagFilters.Values, aws.String(tag_value))

		getResourcesInput := &resourcegroupstaggingapi.GetResourcesInput{}
		getResourcesInput.TagFilters = append(getResourcesInput.TagFilters, tagFilters)

		getResourcesInput.ResourcesPerPage = &items
		getResourcesInput.PaginationToken = token

		if resource_type != "" {
			getResourcesInput.ResourceTypeFilters = []*string{
				aws.String(resource_type),
			}
		}

		resources, err := svr.GetResources(getResourcesInput)
		if err != nil {
			fmt.Errorf("Unable to initialize AWS session: %v", err)
		}

		for _, resource := range resources.ResourceTagMappingList {

			name := findNameTag(resource.Tags)
			arn := *resource.ResourceARN

			cloud_resources[arn] = Resource{
				Name:   name,
				Arn:    arn,
				Region: region,
			}

		}

		token = resources.PaginationToken

		if token == nil || *token == "" {
			break
		}
	}

	return cloud_resources

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
		"Tag:Name | ARN | Region",
	}

	for _, item := range items {
		output = append(output, fmt.Sprintf("%s | %s | %s", item.Name, item.Arn, item.Region))
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
