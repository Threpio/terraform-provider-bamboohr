terraform {
  required_providers {
    bamboohr = {
      source = "threpio/bamboohr"
      version = "0.1.0"
    }
  }
}

provider "bamboohr"{
  subdomain = "aiia"
  api_key = "9842ead9507b12ea112f977f7b1e4c59a0328315"
}

data "bamboohr_employee" "employee123" {
  id = 128
}

data "bamboohr_employees" "all_employees" {}

output "bamboohr_employees" {
  value = data.bamboohr_employees
}