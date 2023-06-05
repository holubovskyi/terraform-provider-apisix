# Terraform Apisix Provider

Forked from the https://github.com/JorgeAraujo123/terraform-provider-apisix

## Local Provider Install
Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Create a new file called `.terraformrc` in your home directory (`~`), then add the `dev_overrides` block below. Change the `<PATH>` to the value returned from the `go env GOBIN` command.

If the `GOBIN` go environment variable is not set, use the default path, `/home/<Username>/go/bin`.

```terraform
provider_installation {

  dev_overrides {
      "hashicorp.com/edu/apisix" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

```
### Locally install provider and verify with Terraform
Use the go install command from the example repository's root directory to compile the provider into a binary and install it in your `GOBIN` path.
```bash
$ go install .
```
Run Terraform plan in the `examples/provider-install-verification` dirrectory.
Running a Terraform plan will report the provider override, as well as an error about the missing provider configuration.

```bash
$ cd examples/provider-install-verification/
$ terraform plan
╷                                                                                                                                                         
│ Warning: Provider development overrides are in effect                                                                                                           
│                                                                                                                                                                 
│ The following provider development overrides are set in the CLI configuration:                                                                                  
│  - hashicorp.com/edu/hashicups-pf in /home/mholubovskyi/go/bin                                                                                                  
│  - hashicorp.com/edu/apisix in /home/mholubovskyi/go/bin                                                                                                        
│                                                                                                                                                                 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible with published        
│ releases.                                                                                                                                                       
╵                                                                                                                                                                 
╷                                                                                
│ Error: Missing APISIX API Endpoint                                             
│                                                                                                                                                                 
│   with provider["hashicorp.com/edu/apisix"],                                                                                                                    
│   on main.tf line 9, in provider "apisix":                                     
│    9: provider "apisix" {}                                                                                                                                      
│                                                                                                                                                                 
│ The provider cannot create the APISIX API client as there is a missing or empty value for the APISIX API endpoint. Set the endpoint value in the configuration,
│ or use the APISIX_ENDPOINT environment variable. If either is already set, ensure the value is not empty.
╵                                                                    
```

### Start Apache APISIX Locally
You can start local APISIX instance using the provided Docker Compose file. The file was adapted from the [apisix-docker repository](https://github.com/apache/apisix-docker/blob/master/example/docker-compose.yml).

In another terminal window, navigate to the `docker_compose` directory.
```bash
cd docker_compose
```
Run `docker-compose up` to spin up a local instance of HashiCups on port 9080.
```bash
docker-compose up
```
Leave this process running in your terminal window. In the original terminal window, verify that APISIX is running by sending a request.
```bash
$ curl "http://127.0.0.1:9180/apisix/admin/services/" \
-H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1'
```
The response indicates that apisix is running successfully:
```json
{"total":0,"list":[]}
```
The credentials for the test user are defined in the `docker_compose/docker-compose.yml` file

## Provider Configuration
The provider configuration method loads configuration data either from environment variables, or from the provider block in Terraform configuration. 

```terraform
provider "apisix" {
  endpoint = "http://127.0.0.1:9180"
  api_key  = "edd1c9f034335f136f87ad84b625c8f1"
}
```
Verify the environment variable behavior by setting the provider-defined HASHICUPS_HOST, HASHICUPS_USERNAME, and HASHICUPS_PASSWORD environment variables when executing a Terraform plan. Terraform will configure the HashiCups client via these environment variables.
```bash
$ APISIX_ENDPOINT=http://127.0.0.1:9180 \
APISIX_API_KEY=edd1c9f034335f136f87ad84b625c8f1 \
terraform plan
```