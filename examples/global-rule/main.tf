terraform {
  required_providers {
    apisix = {
      source = "hashicorp.com/edu/apisix"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}


provider "random" {}

provider "apisix" {
  endpoint = "http://127.0.0.1:9180"
  api_key  = "edd1c9f034335f136f87ad84b625c8f1"
}

resource "random_id" "rule_id" {
  byte_length = 4
}

resource "apisix_global_rule" "example" {
  id = random_id.rule_id.dec
  plugins = jsonencode(
    {
      prometheus = {
        prefer_name = true
      }
    }
  )
}