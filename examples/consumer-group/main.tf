terraform {
  required_providers {
    apisix = {
      source = "hashicorp.com/edu/apisix"
    }
  }
}


provider "apisix" {
  endpoint = "http://127.0.0.1:9180"
  api_key  = "edd1c9f034335f136f87ad84b625c8f1"
}


resource "apisix_consumer_group" "example" {
  id   = "007"
  desc = "Example of the consumer group resource usage"
  plugins = jsonencode(
    {
      prometheus = {
        prefer_name = false
      }
    }
  )
  labels = {
    version = "v1"
  }
}