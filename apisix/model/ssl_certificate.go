package model

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// SSLCertificateResourceModel maps the resource schema data.
type SSLCertificateResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Status      types.Int64  `tfsdk:"status"`
	Certificate types.String `tfsdk:"certificate"`
	PrivateKey  types.String `tfsdk:"private_key"`
	Snis        types.List   `tfsdk:"snis"`
	Type        types.String `tfsdk:"type"`
	Labels      types.Map    `tfsdk:"labels"`
}

var SSLCertificateSchema = schema.Schema{
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
			Sensitive:   true,
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
			MarkdownDescription: "Attributes of the resource specified as key-value pairs. An individual pair cannot be deleted using APISIX API" +
				"In order to delete an individual pair, you can delete all labels and reapply the resource with the desired labels map",
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

func SSLCertificateFromTerraformToAPI(ctx context.Context, terraformDataModel *SSLCertificateResourceModel) (apiDataModel api_client.SSLCertificate, snisDiag diag.Diagnostics, labelsDiag diag.Diagnostics) {
	apiDataModel.Status = uint(terraformDataModel.Status.ValueInt64())
	apiDataModel.Certificate = terraformDataModel.Certificate.ValueString()
	apiDataModel.PrivateKey = terraformDataModel.PrivateKey.ValueString()
	apiDataModel.Type = terraformDataModel.Type.ValueString()

	snisDiag = terraformDataModel.Snis.ElementsAs(ctx, &apiDataModel.SNIs, false)
	labelsDiag = terraformDataModel.Labels.ElementsAs(ctx, &apiDataModel.Labels, false)

	return apiDataModel, snisDiag, labelsDiag
}

func SSLCertificateFromAPIToTerraform(ctx context.Context, apiDataModel *api_client.SSLCertificate) (terraformDataModel SSLCertificateResourceModel, snisDiag diag.Diagnostics, labelsDiag diag.Diagnostics) {
	terraformDataModel.ID = types.StringValue(apiDataModel.ID)
	terraformDataModel.Status = types.Int64Value(int64(apiDataModel.Status))
	terraformDataModel.Certificate = types.StringValue(apiDataModel.Certificate)
	// APISIX API returns the private key in base64 form
	//terraformDataModel.PrivateKey = types.StringValue(apiDataModel.PrivateKey)
	terraformDataModel.Type = types.StringValue(apiDataModel.Type)

	terraformDataModel.Snis, snisDiag = types.ListValueFrom(ctx, types.StringType, apiDataModel.SNIs)
	terraformDataModel.Labels, labelsDiag = types.MapValueFrom(ctx, types.StringType, apiDataModel.Labels)

	return terraformDataModel, snisDiag, labelsDiag
}
