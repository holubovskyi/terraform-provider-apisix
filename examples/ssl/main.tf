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

resource "apisix_ssl_certificate" "example" {
  certificate = file("example.crt")
  private_key = file("example.key")
  type        = "server"
  labels = {
    "version" = "v1"
  }
}