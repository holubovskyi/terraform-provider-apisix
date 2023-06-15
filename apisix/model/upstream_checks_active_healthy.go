package model

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamChecksActiveHealthyType struct {
	Interval     types.Int64 `tfsdk:"interval"`
	HTTPStatuses types.List  `tfsdk:"http_statuses"`
	Successes    types.Int64 `tfsdk:"successes"`
}

var UpstreamChecksActiveHealthySchemaAttribute = schema.SingleNestedAttribute{
	Optional: true,
	Attributes: map[string]schema.Attribute{
		"interval": schema.Int64Attribute{
			MarkdownDescription: "Active check (healthy node) check interval (unit: second)",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(1),
			Validators: []validator.Int64{
				int64validator.AtLeast(1),
			},
		},
		"http_statuses": schema.ListAttribute{
			MarkdownDescription: "Active check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node.",
			ElementType:         types.Int64Type,
			Optional:            true,
			Computed:            true,
			Validators: []validator.List{
				listvalidator.ValueInt64sAre(int64validator.Between(200, 599)),
			},
		},
		"successes": schema.Int64Attribute{
			MarkdownDescription: "Active check (healthy node) determine the number of times a node is healthy.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 254),
			},
			Default: int64default.StaticInt64(2),
		},
	},
}
