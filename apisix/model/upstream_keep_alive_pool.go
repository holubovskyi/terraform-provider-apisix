package model

import (
	"math/big"
	"terraform-provider-apisix/apisix/plan_modifier"
	"terraform-provider-apisix/apisix/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamKeepAlivePoolType struct {
	Size        types.Number `tfsdk:"size"`
	IdleTimeout types.Number `tfsdk:"idle_timeout"`
	Requests    types.Number `tfsdk:"requests"`
}

var UpstreamKeepAlivePoolSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Computed: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"size": {
			Required: true,
			Type:     types.NumberType,
		},
		"idle_timeout": {
			Required: true,
			Type:     types.NumberType,
		},
		"requests": {
			Required: true,
			Type:     types.NumberType,
		},
	}),
	PlanModifiers: []tfsdk.AttributePlanModifier{
		plan_modifier.DefaultObject(
			map[string]attr.Type{
				"size":         types.NumberType,
				"idle_timeout": types.NumberType,
				"requests":     types.NumberType,
			},
			map[string]attr.Value{
				"size":         types.Number{Value: big.NewFloat(320)},
				"idle_timeout": types.Number{Value: big.NewFloat(60)},
				"requests":     types.Number{Value: big.NewFloat(1000)},
			},
		),
	},
}

func UpstreamKeepAlivePoolMapToState(data map[string]interface{}) *UpstreamKeepAlivePoolType {
	v := data["keepalive_pool"]

	if v == nil {
		return nil
	}
	value := v.(map[string]interface{})
	output := UpstreamKeepAlivePoolType{}

	utils.MapValueToNumberTypeValue(value, "size", &output.Size)
	utils.MapValueToNumberTypeValue(value, "idle_timeout", &output.IdleTimeout)
	utils.MapValueToNumberTypeValue(value, "requests", &output.Requests)

	return &output
}

func UpstreamKeepAlivePoolStateToMap(state *UpstreamKeepAlivePoolType, dMap map[string]interface{}) {
	if state == nil {
		return
	}

	output := make(map[string]interface{})
	utils.NumberTypeValueToMap(state.Size, output, "size")
	utils.NumberTypeValueToMap(state.IdleTimeout, output, "idle_timeout")
	utils.NumberTypeValueToMap(state.Requests, output, "requests")

	dMap["keepalive_pool"] = output
}
