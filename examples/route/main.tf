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
  type = "roundrobin"
  nodes = [
    {
      host   = "127.0.0.1"
      port   = "8080"
      weight = 1
    }
  ]
}

resource "apisix_route" "example" {
  name        = "testroute"
  desc        = "Example of the route configuration"
  uri         = "/test"
  hosts       = ["foo.com", "*.bar.com"]
  remote_addr = "10.0.0.0/8"
  upstream_id = apisix_upstream.example.id
  methods     = ["GET", "POST"]
  priority    = 2
  labels = {
    "version" : "0.1"
  }
  plugins = {
    ip_restriction = {
      blacklist = ["10.20.10.77"]
      message   = "Access denied"
    }
  }
}
