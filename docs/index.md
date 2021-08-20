# BambooHR Provider

The BambooHR provider can be used to configure [BambooHR](https://www.bamboohr.com/) using their production API. Documentation regarding the [Data Sources](https://www.terraform.io/docs/language/data-sources/index.html) and [Resources](https://www.terraform.io/docs/language/resources/index.html) supported by the Azure Active Directory Provider can be found in the navigation to the left.

Interested in the provider's latest features, or want to make sure you're up to date? Check out the [changelog](https://github.com/threpio/terraform-provider-bamboohr/blob/main/CHANGELOG.md) for version information and release notes.

## Example Usage

```hcl
#Configure Terraform
terraform {
 required_providers {
  bamboohr = {
    source = "threpio/bamboohr"
    version = "~> 0.1.0"
  } 
 }
}

# Configure the Bamboo HR Provider
provider "bamboohr" {
  subdomain = "<yourSubdomain>"
  api_key = "<yourAPIKey>"
}

# Retrieve All Users Directory 
data "bamboohr_employees" "example" {}

# Retrieve A specific User by ID
data "bamboohr_employee" "example_user" {
  id = 128
}
```