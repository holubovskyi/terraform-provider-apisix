package apisix

import (
	"context"
	"fmt"

	"github.com/holubovskyi/apisix-client-go"

	"terraform-provider-apisix/apisix/model"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
//	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                     = &upstreamResource{}
	_ resource.ResourceWithConfigure        = &upstreamResource{}
	_ resource.ResourceWithImportState      = &upstreamResource{}
	_ resource.ResourceWithConfigValidators = &upstreamResource{}
	// _ resource.ResourceWithValidateConfig = &upstreamResource{}
)

// NewUpstreamResource is a helper function to simplify the provider implementation.
func NewUpstreamResource() resource.Resource {
	return &upstreamResource{}
}

// upstreamResource is the resource implementation.
type upstreamResource struct {
	client *api_client.ApiClient
}

// Metadata returns the resource type name.
func (r *upstreamResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_upstream"
}

// Schema defines the schema for the resource.
func (r *upstreamResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = model.UpstreamSchema
}

// Validate Config
func (r *upstreamResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(
			path.MatchRoot("service_name"),
			path.MatchRoot("nodes"),
		),
		resourcevalidator.RequiredTogether(
			path.MatchRoot("service_name"),
			path.MatchRoot("discovery_type"),
		),
	}
}

// Configure adds the provider configured client to the resource.
func (r *upstreamResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// Create a new resource.
func (r *upstreamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start of the upstream resource creation")
	// Retrieve values from plan
	var plan model.UpstreamResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	newUpstreamRequest, labelsDiag := model.UpstreamFromTerraformToAPI(ctx, &plan)

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new upstream
	newUpstreamResponse, err := r.client.CreateUpstream(newUpstreamRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Upstream",
			"Could not create Upstream, unexpected error: "+err.Error(),
		)
		return
	}	
	tflog.Info(ctx, "Response after upstream creation", map[string]any{
		"Info": newUpstreamResponse.Type,
	})

}

// Read resource information.
func (r *upstreamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	return
}

// Update resource.
func (r *upstreamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	return
}

// Delete resource.
func (r *upstreamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	return
}

// Import resource into state
func (r *upstreamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	return
}
