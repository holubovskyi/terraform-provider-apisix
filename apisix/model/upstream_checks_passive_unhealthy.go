package model

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamChecksPassiveUnhealthyType struct {
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	TCPFailures  types.Number `tfsdk:"tcp_failures"`
	Timeouts     types.Number `tfsdk:"timeouts"`
	HTTPFailures types.Number `tfsdk:"http_failures"`
}

var UpstreamChecksPassiveUnhealthySchemaAttribute = schema.SingleNestedAttribute{
	Optional: true,
	Attributes: map[string]schema.Attribute{
		"http_statuses": schema.ListAttribute{
			MarkdownDescription: "Passive check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node.",
			ElementType:         types.Int64Type,
			Optional:            true,
			Computed:            true,
			Validators: []validator.List{
				listvalidator.ValueInt64sAre(int64validator.Between(200, 599)),
			},
		},
		"http_failures": schema.Int64Attribute{
			MarkdownDescription: "Passive check (unhealthy node) The number of times that the node is not healthy during HTTP or HTTPS type checking.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 254),
			},
			Default: int64default.StaticInt64(5),
		},
		"tcp_failures": schema.Int64Attribute{
			MarkdownDescription: "Passive check (unhealthy node) When TCP type is checked, determine the number of times that the node is not healthy.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 254),
			},
			Default: int64default.StaticInt64(2),
		},
		"timeouts": schema.Int64Attribute{
			MarkdownDescription: "Passive checks (unhealthy node) determine the number of timeouts for unhealthy nodes.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 254),
			},
			Default: int64default.StaticInt64(7),
		},
	},
}
