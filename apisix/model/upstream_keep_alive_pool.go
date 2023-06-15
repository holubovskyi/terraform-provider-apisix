package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamKeepAlivePoolType struct {
	Size        types.Number `tfsdk:"size"`
	IdleTimeout types.Number `tfsdk:"idle_timeout"`
	Requests    types.Number `tfsdk:"requests"`
}

var UpstreamKeepAlivePoolSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Sets the `keepalive_pool`.",
	Optional:            true,
	Computed:            true,
	Attributes: map[string]schema.Attribute{
		"size": schema.Int64Attribute{
			Required: true,
		},
		"idle_timeout": schema.Int64Attribute{
			Required: true,
		},
		"requests": schema.Int64Attribute{
			Required: true,
		},
	},
}
