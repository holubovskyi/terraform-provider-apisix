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

resource "apisix_upstream" "example" {
  name = "Example"
  desc = "Example of the upstream resource usage"
  type = "roundrobin"
  nodes = [
    {
      host   = "127.0.0.1"
      port   = 8080
      weight = 1
    },
    {
      host = "10.10.10.10"
      port = 80
      weight = 2
    }
  ]
}