package apisix

import (
	"context"
	"fmt"

	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &routeDataSource{}
	_ datasource.DataSourceWithConfigure = &routeDataSource{}
)

// NewRouteDataSource is a helper function to simplify the provider implementation.
func NewRouteDataSource() datasource.DataSource {
	return &routeDataSource{}
}

// routeDataSource is the data source implementation
type routeDataSource struct {

	client *api_client.ApiClient
}

// Metadata returns the data source type name.
func (d *routeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

// Configure adds the provider configured client to the data source.
func (d *routeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api_client.ApiClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api_client.ApiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

// Schema defines the schema for the data source.
func (d *routeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Numeric identifier of the route.",
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Identifier for the Route.",
				Computed: true,
			},
			"desc": schema.StringAttribute{
				Description: "Description of usage scenarios.",
				Computed: true,
			},
			"uri": schema.StringAttribute{
				Description: "Matches the uri. For more advanced matching see Router.",
				Computed: true,
			},							
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *routeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}