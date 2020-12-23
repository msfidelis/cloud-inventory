<h1 align="left">Welcome to Cloud Inventory Tags 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v0.0.4-blue.svg?cacheSeconds=2592000" />
  <a href="README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="LICENSE " target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/fidelissauro" target="_blank">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" />
  </a>
  <a href="/" target="_blank">
    <img alt="Build CI" src="https://github.com/msfidelis/cloud-inventory/workflows/cloud-inventory%20ci/badge.svg" />
  </a>  
  <a href="/" target="_blank">
    <img alt="Release" src="https://github.com/msfidelis/cloud-inventory/workflows/release%20packages/badge.svg" />
  </a>
</p>

> Simple tool to search tagged resources around all AWS Account

## Installation

### MacOS / OSX

```bash
wget https://github.com/msfidelis/cloud-inventory/releases/download/v0.0.4/cloud-inventory_0.0.4_darwin_amd64 -O /usr/local/bin/cloud-inventory
chmod +x /usr/local/bin/cloud-inventory
```

### Linux 

```bash
wget https://github.com/msfidelis/cloud-inventory/releases/download/v0.0.4/cloud-inventory_0.0.4_linux_amd64 -O /usr/local/bin/cloud-inventory
chmod +x /usr/local/bin/cloud-inventory
```

## Usage

```sh
cloud-inventory -h

  -output string
    	Output report type; ex: default, arn, csv (default "default")
  -region string
    	Region to search inventory; default: us-east-1 (default "us-east-1")
  -resource string
    	Optional resource type; ex: ec2, s3, acm
  -tag-name string
    	Tag to search
  -tag-value string
    	Tag to search
```

## Searching for a tag

```sh
cloud-inventory --tag-name Project --tag-value k8s-with-cri-o
Searching for resources using Tag Project:k8s-with-cri-o
Tag:Name                          ARN                                                                     Region
k8s-node-3                        arn:aws:ec2:us-east-1:181560427716:instance/i-08658a0eca5db90ae         us-east-1
k8s-master                        arn:aws:ec2:us-east-1:181560427716:instance/i-09471d9233d5eaf8b         us-east-1
k8s-with-cri-o-kubernetes-sg      arn:aws:ec2:us-east-1:181560427716:security-group/sg-0354149cb0868e23f  us-east-1
k8s-with-cri-o-public-us-east-1a  arn:aws:ec2:us-east-1:181560427716:subnet/subnet-07e7fab89a7f38e77      us-east-1
k8s-with-cri-o-vpc                arn:aws:ec2:us-east-1:181560427716:vpc/vpc-068939c361da4db98            us-east-1
k8s-node-2                        arn:aws:ec2:us-east-1:181560427716:instance/i-08064cc28841d1e06         us-east-1
k8s-node-1                        arn:aws:ec2:us-east-1:181560427716:instance/i-01640c6de2121a8f4         us-east-1
k8s-node-0                        arn:aws:ec2:us-east-1:181560427716:instance/i-0634e8bd2de7e2615         us-east-1
-                                 arn:aws:ec2:us-east-1:181560427716:key-pair/key-054c61aaddb4cbb9c       us-east-1
k8s-with-cri-o-public-us-east-1b  arn:aws:ec2:us-east-1:181560427716:subnet/subnet-04c9e8e42473dd2b9      us-east-1

Found 10 resources
```


## Searching for a tag on a specific AWS service

```sh
cloud-inventory --tag-name Project --tag-value k8s-with-cri-o --resource ec2
Searching for resources using Tag Project:k8s-with-cri-o
Tag:Name                          ARN                                                                     Region
k8s-node-1                        arn:aws:ec2:us-east-1:181560427716:instance/i-01640c6de2121a8f4         us-east-1
-                                 arn:aws:ec2:us-east-1:181560427716:key-pair/key-054c61aaddb4cbb9c       us-east-1
k8s-node-2                        arn:aws:ec2:us-east-1:181560427716:instance/i-08064cc28841d1e06         us-east-1
k8s-node-0                        arn:aws:ec2:us-east-1:181560427716:instance/i-0634e8bd2de7e2615         us-east-1
k8s-master                        arn:aws:ec2:us-east-1:181560427716:instance/i-09471d9233d5eaf8b         us-east-1
k8s-node-3                        arn:aws:ec2:us-east-1:181560427716:instance/i-08658a0eca5db90ae         us-east-1

Found 6 resources
```

## Customize output format 

```sh
cloud-inventory --tag-name Project --tag-value CarsAndBus --resource rds --output csv

Searching for resources using Tag Project:CarsAndBus

Found 6 resources

Output file: results.csv
```

## Docker usage

```sh
docker run -it fidelissauro/cloud-inventory:latest --tag-name Project --tag-value CarsAndBus
```


## Run tests

```sh
go test -v 
```

## Author

👤 **Matheus Fidelis**

* Website: https://raj.ninja
* Twitter: [@fidelissauro](https://twitter.com/fidelissauro)
* Github: [@msfidelis](https://github.com/msfidelis)
* LinkedIn: [@msfidelis](https://linkedin.com/in/msfidelis)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Matheus Fidelis](https://github.com/msfidelis).<br />
This project is [MIT](LICENSE ) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
