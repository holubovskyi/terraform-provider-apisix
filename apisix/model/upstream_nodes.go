package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamNodeType struct {
	Host   types.String `tfsdk:"host"`
	Port   types.Number `tfsdk:"port"`
	Weight types.Number `tfsdk:"weight"`
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
