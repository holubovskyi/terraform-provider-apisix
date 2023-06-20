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

resource "apisix_service" "example" {
  name  = "Example"
  hosts = ["foo.com", "*.bar.com"]
  labels = {
    "version" = "v1"
  }
  enable_websocket = false
  upstream_id      = apisix_upstream.example.id
  plugins          = <<EOF
{
	"limit-count": {
		"count": 2,
		"time_window": 100,
		"rejected_code": 503,
		"key": "remote_addr"
	},
	"prometheus": {}
}
EOF  
}