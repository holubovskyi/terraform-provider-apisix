package apisix

import (
	"context"
	"fmt"

	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

// sslCertificateResourceModel maps the resource schema data.
type sslCertificateResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Status      types.Int64  `tfsdk:"status"`
	Certificate types.String `tfsdk:"certificate"`
	PrivateKey  types.String `tfsdk:"private_key"`
	Snis        types.List   `tfsdk:"snis"`
	Type        types.String `tfsdk:"type"`
	Labels      types.Map    `tfsdk:"labels"`
}

// Metadata returns the resource type name.
func (r *sslCertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_certificate"
}

// Schema defines the schema for the resource.
func (r *sslCertificateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages SSL certificates.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the certificate.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate": schema.StringAttribute{
				Description: "HTTPS certificate.",
				Required:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "HTTPS private key.",
				Required:    true,
				Sensitive:   false,
			},
			"snis": schema.ListAttribute{
				MarkdownDescription: "A non-empty array of HTTPS SNI. Required if `type` is `server`",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Identifies the type of certificate, default `server`\n" +
					"`client` Indicates that the certificate is a client certificate, which is used when APISIX accesses the upstream;\n" +
					"`server` Indicates that the certificate is a server-side certificate, which is used by APISIX when verifying client requests.",
				Required: true,
				Validators: []validator.String{
					// Validate string value must be "server" or "client"
					stringvalidator.OneOfCaseInsensitive([]string{"server", "client"}...),
				},
			},
			"labels": schema.MapAttribute{
				Description: "Attributes of the resource specified as key-value pairs.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"status": schema.Int64Attribute{
				MarkdownDescription: "Enables the current SSL. Set to `1` (enabled) by default. `1` to enable, `0` to disable",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators: []validator.Int64{
					// Validate integer value must be 0 or 1
					int64validator.OneOf([]int64{0, 1}...),
				},
			},
		},
	}
}

// Validate that snis are specified when certificate type is `server`
func (r *sslCertificateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data sslCertificateResourceModel

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
	// Retrieve values from plan
	var plan sslCertificateResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var certificate api_client.SSLCertificate

	certificate.Status = uint(plan.Status.ValueInt64())
	certificate.Certificate = plan.Certificate.ValueString()
	certificate.PrivateKey = plan.PrivateKey.ValueString()
	certificate.Type = plan.Type.ValueString()

	resp.Diagnostics.Append(plan.Snis.ElementsAs(ctx, &certificate.SNIs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.Labels.ElementsAs(ctx, &certificate.Labels, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "SSLCertificate struct contains:", map[string]any{
		"snis":   certificate.SNIs,
		"labels": certificate.Labels,
	})

	// Create new certificate
	newCertificateRequest, err := r.client.CreateSslCertificate(certificate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SSL certificate",
			"Could not create SSL certificate, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(newCertificateRequest.ID)
	plan.Status = types.Int64Value(int64(newCertificateRequest.Status))
	plan.Certificate = types.StringValue(newCertificateRequest.Certificate)
	//plan.PrivateKey = types.StringValue(newCertificateRequest.PrivateKey)
	plan.Type = types.StringValue(newCertificateRequest.Type)

	plan.Snis, diags = types.ListValueFrom(ctx, types.StringType, newCertificateRequest.SNIs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Labels, diags = types.MapValueFrom(ctx, types.StringType, newCertificateRequest.Labels)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *sslCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state sslCertificateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed certificate from the APISIX
	certificate, err := r.client.GetSslCertificate(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading APISIX SSL Certificate",
			"Could not read APISIX SSL Certificate ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite with refreshed state
	newState := sslCertificateResourceModel{}

	newState.ID = types.StringValue(certificate.ID)
	newState.Status = types.Int64Value(int64(certificate.Status))
	newState.Certificate = types.StringValue(certificate.Certificate)
	newState.PrivateKey = types.StringValue(state.PrivateKey.ValueString())
	newState.Type = types.StringValue(certificate.Type)

	newState.Snis, diags = types.ListValueFrom(ctx, types.StringType, certificate.SNIs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.Labels, diags = types.MapValueFrom(ctx, types.StringType, certificate.Labels)
	resp.Diagnostics.Append(diags...)
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
	return
}

// Delete resource.
func (r *sslCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	return
}

// Import resource into state
func (r *sslCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	return
}
