package model

import (
	"terraform-provider-apisix/apisix/plan_modifier"
	"terraform-provider-apisix/apisix/utils"
	"terraform-provider-apisix/apisix/validator"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamChecksPassiveHealthyType struct {
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	Successes    types.Number `tfsdk:"successes"`
}

var UpstreamChecksPassiveHealthySchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"http_statuses": {
			Type:     types.ListType{ElemType: types.NumberType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfNumbers(200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308),
			},
			Description: "Passive check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node",
		},

		"successes": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(254),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(5),
			},
			Description: "Passive checks (healthy node) determine the number of times a node is healthy",
		},
	}),
}

func UpstreamChecksPassiveHealthyMapToState(data map[string]interface{}) *UpstreamChecksPassiveHealthyType {
	v := data["healthy"]
	if v == nil {
		return nil
	}

	value := v.(map[string]interface{})
	output := UpstreamChecksPassiveHealthyType{}

	utils.MapValueToNumberTypeValue(value, "successes", &output.Successes)
	utils.MapValueToListTypeValue(value, "http_statuses", &output.HTTPStatuses)

	return &output
}

func UpstreamChecksPassiveHealthyStateToMap(state *UpstreamChecksPassiveHealthyType, dMap map[string]interface{}) {
	if state == nil {
		return
	}

	output := make(map[string]interface{})

	utils.ListTypeValueToMap(state.HTTPStatuses, output, "http_statuses")
	utils.NumberTypeValueToMap(state.Successes, output, "successes")

	dMap["healthy"] = output

}
