package model

import (
	"context"

	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type UpstreamChecksType struct {
	Active  *UpstreamChecksActiveType  `tfsdk:"active"`
	Passive *UpstreamChecksPassiveType `tfsdk:"passive"`
}

var UpstreamChecksSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Configures the parameters for the health check.",
	Optional:            true,

	Attributes: map[string]schema.Attribute{
		"active":  UpstreamChecksActiveSchemaAttribute,
		"passive": UpstreamChecksPassiveSchemaAttribute,
	},
}

func UpstreamChecksFromTerraformToAPI(ctx context.Context, terraformDataModel *UpstreamChecksType) (apiDataModel api_client.UpstreamChecksType) {
	if terraformDataModel == nil {
		return
	}

	apiDataModel.Active = UpstreamChecksActiveFromTerraformToApi(ctx, terraformDataModel.Active)
	apiDataModel.Passive = UpstreamChecksPassiveFromTerraformToApi(ctx, terraformDataModel.Passive)

	return apiDataModel
}
