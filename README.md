<h1 align="left">Welcome to Cloud Inventory Tags 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v0.0.0-blue.svg?cacheSeconds=2592000" />
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
wget https://github.com/msfidelis/cloud-inventory/releases/download/v0.0.3/cloud-inventory_0.0.3_darwin_amd64 -O /usr/local/bin/cloud-inventory
chmod +x /usr/local/bin/cloud-inventory
```

### Linux 

```bash
wget https://github.com/msfidelis/cloud-inventory/releases/download/v0.0.3/cloud-inventory_0.0.3_linux_amd64 -O /usr/local/bin/cloud-inventory
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
cloud-inventory --tag-name Project --tag-value CarsAndBus
```


## Searching for a tag on a specific AWS service

```sh
cloud-inventory --tag-name Project --tag-value CarsAndBus --resource rds
```

## Customize output format 

```sh
cloud-inventory --tag-name Project --tag-value CarsAndBus --resource rds --output csv

Searching for resources using Tag Project:CarsAndBus
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
