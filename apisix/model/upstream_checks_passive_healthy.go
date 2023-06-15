package model

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamChecksPassiveHealthyType struct {
	HTTPStatuses types.List  `tfsdk:"http_statuses"`
	Successes    types.Int64 `tfsdk:"successes"`
}

var UpstreamChecksPassiveHealthySchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Passive health check refers to judging whether the corresponding upstream node is healthy by judging the response status of the request forwarded from APISIX to the upstream node.",
	Optional:            true,
	Attributes: map[string]schema.Attribute{
		"http_statuses": schema.ListAttribute{
			MarkdownDescription: "Passive check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node.",
			ElementType:         types.Int64Type,
			Optional:            true,
			Validators: []validator.List{
				listvalidator.ValueInt64sAre(int64validator.Between(200, 599)),
			},
		},
		"successes": schema.Int64Attribute{
			MarkdownDescription: "Passive checks (healthy node) determine the number of times a node is healthy.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 254),
			},
			Default: int64default.StaticInt64(5),
		},
	},
}
