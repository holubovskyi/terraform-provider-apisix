package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TimeoutType struct {
	Connect types.Int64 `tfsdk:"connect"`
	Send    types.Int64 `tfsdk:"send"`
	Read    types.Int64 `tfsdk:"read"`
}

var TimeoutSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Sets the timeout (in seconds) for connecting to, and sending and receiving messages to and from the Upstream.",
	Optional:            true,
	Computed:            true,
	Attributes: map[string]schema.Attribute{
		"connect": schema.Int64Attribute{
			Required: true,
		},
		"send": schema.Int64Attribute{
			Required: true,
		},
		"read": schema.Int64Attribute{
			Required: true,
		},
	},
}
