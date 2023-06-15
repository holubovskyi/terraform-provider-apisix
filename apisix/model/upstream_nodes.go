package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api_client "github.com/holubovskyi/apisix-client-go"
)

type UpstreamNodeType struct {
	Host   types.String `tfsdk:"host"`
	Port   types.Int64  `tfsdk:"port"`
	Weight types.Int64  `tfsdk:"weight"`
}

var UpstreamNodesSchemaAttribute = schema.ListNestedAttribute{
	MarkdownDescription: "Configures the parameters for the health check.",
	Optional:            true,

	NestedObject: schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required: true,
			},
			"port": schema.Int64Attribute{
				Required: true,
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1),
			},
		},
	},
}

func UpstreamNodesFromTerraformToAPI(terraformDataModel *[]UpstreamNodeType) (apiDataModel []api_client.UpstreamNodeType) {
	if terraformDataModel == nil {
		return
	}

	for i, v := range *terraformDataModel {
		apiDataModel[i].Host = v.Host.ValueString()
		apiDataModel[i].Port = uint(v.Port.ValueInt64())
		apiDataModel[i].Weight = uint(v.Weight.ValueInt64())
	}
	return apiDataModel
}
