# Terraform Apisix Provider

Forked from the https://github.com/JorgeAraujo123/terraform-provider-apisix

### Local Provider Install
Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Create a new file called `.terraformrc` in your home directory (`~`), then add the `dev_overrides` block below. Change the `<PATH>` to the value returned from the `go env GOBIN` command.

If the `GOBIN` go environment variable is not set, use the default path, `/Users/<Username>/go/bin`.

```terraform
provider_installation {

  dev_overrides {
      "hashicorp.com/edu/hashicups-pf" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

```

### Docker Compose
The Docker compose configuration is from the [apisix-docker](https://github.com/apache/apisix-docker/blob/master/example/docker-compose.yml) repository