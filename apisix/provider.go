package apisix

import (
	"context"
	"os"

	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &apisixProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &apisixProvider{
			version: version,
		}
	}
}

// apisixProvider is the provider implementation.
type apisixProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// apisixProviderModel maps provider schema data to a Go type.
type apisixProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ApiKey   types.String `tfsdk:"api_key"`
}

// Metadata returns the provider type name.
func (p *apisixProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "apisix"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *apisixProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Apache APISIX,",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Description: "URI for Apache APISIX API. May also be provided via APISIX_ENDPOINT environment variable.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Description: "API Key for Apache APISIX API. May also be provided via APISIX_API_KEY environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *apisixProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configring APISIX client")

	// Retrieve provider data from configuration
	var config apisixProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown APISIX API Endpoint",
			"The provider cannot create the APISIX API client as there is an unknown configuration value for the APISIX API endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the APISIX_ENDPOINT environment variable.",
		)
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown APISIX API Key ",
			"The provider cannot create the APISIX API client as there is an unknown configuration value for the APISIX API Key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the APISIX_API_KEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	endpoint := os.Getenv("APISIX_ENDPOINT")
	api_key := os.Getenv("APISIX_API_KEY")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.ApiKey.IsNull() {
		api_key = config.ApiKey.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing APISIX API Endpoint",
			"The provider cannot create the APISIX API client as there is a missing or empty value for the APISIX API endpoint. "+
				"Set the endpoint value in the configuration, or use the APISIX_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing APISIX API Key ",
			"The provider cannot create the APISIX API client as there is a missing or empty value for the APISIX API Key. "+
				"Set the API Key value in the configuration, or use the APISIX_API_KEY environment variable."+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "apisix_endpoint", endpoint)
	ctx = tflog.SetField(ctx, "apisix_apikey", api_key)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "apisix_apikey")

	tflog.Debug(ctx, "Creating APISIX client")

	// Create a new APISIX client using the configuration values
	client := api_client.GetCl(config.ApiKey.ValueString(), config.Endpoint.ValueString())
	// TODO: Client should return an error
	//client, err := api_client.GetCl(config.ApiKey.ValueString(), config.Endpoint.ValueString())
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Unable to Create HashiCups API Client",
	// 		"An unexpected error occurred when creating the HashiCups API client. "+
	// 			"If the error is not clear, please contact the provider developers.\n\n"+
	// 			"HashiCups Client Error: "+err.Error(),
	// 	)
	// 	return
	// }

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured APISIX client", map[string]any{"success": true})
}

// func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
// 	var config ProviderData
// 	diags := req.Config.Get(ctx, &config)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	if config.Endpoint.Unknown || config.Endpoint.Null || config.Endpoint.Value == "" {
// 		// Cannot connect to client with an unknown value
// 		resp.Diagnostics.AddWarning(
// 			"Unable to create client",
// 			"Cannot use unknown value as username",
// 		)
// 		return
// 	}

// 	p.client = api_client.GetCl(config.ApiKey.Value, config.Endpoint.Value)
// }

// DataSources defines the data sources implemented in the provider.
func (p *apisixProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRouteDataSource,
	}
}

// func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
// 	return map[string]tfsdk.DataSourceType{}, nil
// }

func (p *apisixProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
// 	return map[string]tfsdk.ResourceType{
// 		"apisix_consumer":                   ResourceConsumerType{},
// 		"apisix_global_rule":                ResourceGlobalRuleType{},
// 		"apisix_plugin_metadata_log_format": ResourcePluginMetadataLogFormatType{},
// 		"apisix_route":                      ResourceRouteType{},
// 		"apisix_service":                    ResourceServiceType{},
// 		"apisix_ssl_certificate":            ResourceSslCertificateType{},
// 		"apisix_stream_route":               ResourceStreamRouteType{},
// 		"apisix_upstream":                   ResourceUpstreamType{},
// 	}, nil
// }
