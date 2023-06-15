package apisix

import (
	"context"
	"fmt"

	"github.com/holubovskyi/apisix-client-go"

	"terraform-provider-apisix/apisix/model"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                   = &sslCertificateResource{}
	_ resource.ResourceWithConfigure      = &sslCertificateResource{}
	_ resource.ResourceWithImportState    = &sslCertificateResource{}
	_ resource.ResourceWithValidateConfig = &sslCertificateResource{}
)

// NewSSLCertificateResource is a helper function to simplify the provider implementation.
func NewSSLCertificateResource() resource.Resource {
	return &sslCertificateResource{}
}

// sslCertificateResource is the resource implementation.
type sslCertificateResource struct {
	client *api_client.ApiClient
}

// Metadata returns the resource type name.
func (r *sslCertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_certificate"
}

// Schema defines the schema for the resource.
func (r *sslCertificateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = model.SSLCertificateSchema
}

// Validate that snis are specified when certificate type is `server`
func (r *sslCertificateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data model.SSLCertificateResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get Certificate type as a string
	certificateType, diags := data.Type.ToStringValue(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if certificateType == basetypes.NewStringValue("server") && (data.Snis.IsNull() || len(data.Snis.Elements()) == 0) {
		resp.Diagnostics.AddAttributeError(
			path.Root("snis"),
			"Missing Attribute Configuration",
			"Expected snis to be configured if the certificate type is server",
		)
	}
}

// Configure adds the provider configured client to the resource.
func (r *sslCertificateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *sslCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start of the SSL certificate resource creation")
	// Retrieve values from plan
	var plan model.SSLCertificateResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	newCertificateRequest, snisDiag, labelsDiag := model.SSLCertificateFromTerraformToAPI(ctx, &plan)

	resp.Diagnostics.Append(snisDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new certificate
	newCertificateResponse, err := r.client.CreateSslCertificate(newCertificateRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SSL certificate",
			"Could not create SSL certificate, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	newState, snisDiag, labelsDiag := model.SSLCertificateFromAPIToTerraform(ctx, newCertificateResponse)
	newState.PrivateKey = types.StringValue(plan.PrivateKey.ValueString())

	resp.Diagnostics.Append(snisDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *sslCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start of the resource read")
	// Get current state
	var state model.SSLCertificateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed certificate from the APISIX
	certificateStatusResponse, err := r.client.GetSslCertificate(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading APISIX SSL Certificate",
			"Could not read APISIX SSL Certificate ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite with refreshed state
	newState, snisDiag, labelsDiag := model.SSLCertificateFromAPIToTerraform(ctx, certificateStatusResponse)
	newState.PrivateKey = types.StringValue(state.PrivateKey.ValueString())

	resp.Diagnostics.Append(snisDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource.
func (r *sslCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start of the resource Update")
	// Retrieve values from plan
	var plan model.SSLCertificateResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateCertificateRequest, snisDiag, labelsDiag := model.SSLCertificateFromTerraformToAPI(ctx, &plan)

	resp.Diagnostics.Append(snisDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing certificate
	_, err := r.client.UpdateSslCertificate(plan.ID.ValueString(), updateCertificateRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating APISIX SSL Certificate",
			"Could not update SSL certificate, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated certificate
	updatedCertificate, err := r.client.GetSslCertificate(plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading APISIX SSL Certificate",
			"Could not read APISIX SSL Certificate ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	newState, snisDiag, labelsDiag := model.SSLCertificateFromAPIToTerraform(ctx, updatedCertificate)
	newState.PrivateKey = types.StringValue(plan.PrivateKey.ValueString())

	resp.Diagnostics.Append(snisDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(labelsDiag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete resource.
func (r *sslCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start of the resource deletion")
	// Get current state
	var state model.SSLCertificateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing certificate
	err := r.client.DeleteSslCertificate(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting APISIX SSL Certificate",
			"Could not delete certificate, unexpected error: "+err.Error(),
		)
		return
	}
}

// Import resource into state
func (r *sslCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
