package model

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/holubovskyi/apisix-client-go"
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

func UpstreamChecksActiveFromTerraformToApi(ctx context.Context, terraformDataModel *UpstreamChecksActiveType) (apiDataModel *api_client.UpstreamChecksActiveType) {
	if terraformDataModel == nil {
		return
	}

	result := api_client.UpstreamChecksActiveType{
		Type:                   terraformDataModel.Type.ValueString(),
		Timeout:                uint(terraformDataModel.Timeout.ValueInt64()),
		Concurrency:            uint(terraformDataModel.Concurrency.ValueInt64()),
		HTTPPath:               terraformDataModel.HTTPPath.ValueString(),
		Host:                   terraformDataModel.Host.ValueString(),
		Port:                   uint(terraformDataModel.Port.ValueInt64()),
		HTTPSVerifyCertificate: terraformDataModel.HTTPSVerifyCertificate.ValueBool(),
		Healthy:                UpstreamChecksActiveHealthyFromTerraformToApi(ctx, terraformDataModel.Healthy),
		Unhealthy:              UpstreamChecksActiveUnhealthyFromTerraformToApi(ctx, terraformDataModel.Unhealthy),
	}

	_ = terraformDataModel.ReqHeaders.ElementsAs(ctx, &result.ReqHeaders, false)

	return &result
}

func UpstreamChecksActiveFromApiToTerraform(ctx context.Context, apiDataModel *api_client.UpstreamChecksActiveType) (terraformDataModel *UpstreamChecksActiveType) {
	if apiDataModel == nil {
		return
	}

	result := UpstreamChecksActiveType{
		Type:                   types.StringValue(apiDataModel.Type),
		Timeout:                types.Int64Value(int64(apiDataModel.Timeout)),
		Concurrency:            types.Int64Value(int64(apiDataModel.Concurrency)),
		HTTPPath:               types.StringValue(apiDataModel.HTTPPath),
		Host:                   types.StringValue(apiDataModel.Host),
		Port:                   types.Int64Value(int64(apiDataModel.Port)),
		HTTPSVerifyCertificate: types.BoolValue(apiDataModel.HTTPSVerifyCertificate),
		Healthy:                UpstreamChecksActiveHealthyFromApiToTerraform(ctx, apiDataModel.Healthy),
		Unhealthy:              UpstreamChecksActiveUnhealthyFromApiToTerraform(ctx, apiDataModel.Unhealthy),
	}
	result.ReqHeaders, _ = types.ListValueFrom(ctx, types.StringType, apiDataModel.ReqHeaders)

	return &result
}
