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
  labels = {
    "version" : "v1"
  }
  retry_timeout = 30
  timeout = {
    connect = 10
    send    = 5
    read    = 5
  }
  nodes = [
    {
      host   = "127.0.0.1"
      port   = 1980
      weight = 1
    },
    {
      host   = "127.0.0.1"
      port   = 1970
      weight = 1
    },
  ]
  checks = {
    active = {
      host      = "example.com"
      port      = 8888
      timeout   = 5
      http_path = "/status"
      healthy = {
        interval  = 2,
        successes = 1
      }
      unhealthy = {
        interval      = 1
        http_failures = 2
      }
    }
    passive = {
      healthy = {
        http_statuses = [200, 201]
      }
      unhealthy = {
        http_statuses = [500]
        http_failures = 3
        tcp_failures  = 3
      }
    }
  }
}