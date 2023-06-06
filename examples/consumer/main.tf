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

resource "apisix_consumer" "example" {
  username = "example"
  desc     = "Example of the consumer"
  plugins = {
    basic_auth = {
      username = "example"
      password = "changeMe"
    }
  }
}