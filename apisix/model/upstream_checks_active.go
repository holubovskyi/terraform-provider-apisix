package model

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamChecksActiveType struct {
	Type                   types.String                       `tfsdk:"type"`
	Timeout                types.Int64                        `tfsdk:"timeout"`
	Concurrency            types.Int64                        `tfsdk:"concurrency"`
	HTTPPath               types.String                       `tfsdk:"http_path"`
	Host                   types.String                       `tfsdk:"host"`
	Port                   types.Int64                        `tfsdk:"port"`
	HTTPSVerifyCertificate types.Bool                         `tfsdk:"https_verify_certificate"`
	ReqHeaders             types.List                         `tfsdk:"req_headers"`
	Healthy                *UpstreamChecksActiveHealthyType   `tfsdk:"healthy"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `tfsdk:"unhealthy"`
}

var UpstreamChecksActiveSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Active health check mainly means that APISIX actively detects the survivability of upstream nodes through probes.",
	Optional:            true,
	Attributes: map[string]schema.Attribute{
		"type": schema.StringAttribute{
			MarkdownDescription: "The type of active check. Valid values are `http`, `https`, and `tcp`",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("http"),
			Validators: []validator.String{
				stringvalidator.OneOfCaseInsensitive([]string{"http", "https", "tcp"}...),
			},
		},
		"timeout": schema.Int64Attribute{
			MarkdownDescription: "The timeout period of the active check (seconds).",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(1),
		},
		"concurrency": schema.Int64Attribute{
			MarkdownDescription: "The number of targets to be checked at the same time during the active check.",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(10),
		},
		"http_path": schema.StringAttribute{
			MarkdownDescription: "The HTTP request path that is actively checked.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("/"),
		},
		"host": schema.StringAttribute{
			MarkdownDescription: "The hostname of the HTTP request actively checked.",
			Optional:            true,
			Computed:            true,
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "The host port of the HTTP request that is actively checked.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 65535),
			},
		},
		"https_verify_certificate": schema.BoolAttribute{
			MarkdownDescription: "Active check whether to check the SSL certificate of the remote host when HTTPS type checking is used.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"req_headers": schema.ListAttribute{
			MarkdownDescription: "Active check When using HTTP or HTTPS type checking, set additional request header information.",
			ElementType:         types.StringType,
			Optional:            true,
		},
		"healthy":   UpstreamChecksActiveHealthySchemaAttribute,
		"unhealthy": UpstreamChecksActiveUnhealthySchemaAttribute,
	},
}
