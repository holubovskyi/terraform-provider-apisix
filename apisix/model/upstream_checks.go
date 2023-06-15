package model

import (
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
