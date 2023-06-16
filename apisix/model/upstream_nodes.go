package model

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/holubovskyi/apisix-client-go"
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

func UpstreamNodesFromTerraformToAPI(ctx context.Context, terraformDataModel *[]UpstreamNodeType) (apiDataModel *[]api_client.UpstreamNodeType) {
	if terraformDataModel == nil {
		tflog.Error(ctx, "Can't transform upstream nodes to api model")
		return
	}

	var result = []api_client.UpstreamNodeType{}

	for _, v := range *terraformDataModel {
		result = append(result, api_client.UpstreamNodeType{
			Host:   v.Host.ValueString(),
			Port:   uint(v.Port.ValueInt64()),
			Weight: uint(v.Weight.ValueInt64())})
	}
	return &result
}
