package model

import (
	"terraform-provider-apisix/apisix/plan_modifier"
	"terraform-provider-apisix/apisix/utils"
	"terraform-provider-apisix/apisix/validator"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PluginResponseRewriteType struct {
	Disable    types.Bool   `tfsdk:"disable"`
	StatusCode types.Number `tfsdk:"status_code"`
	Body       types.String `tfsdk:"body"`
	BodyBase64 types.Bool   `tfsdk:"body_base64"`
	Headers    types.Map    `tfsdk:"headers"`
	Vars       types.String `tfsdk:"vars"`
}

var PluginResponseRewriteSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"disable": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"status_code": {
			Optional: true,
			Type:     types.NumberType,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(200),
				validator.NumberLessOrEqualThan(598),
			},
		},
		"body": {
			Optional: true,
			Type:     types.StringType,
		},
		"body_base64": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},

		"headers": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},

		"vars": {
			Optional:    true,
			Type:        types.StringType,
			Description: "JSON string",
		},
	}),
}

func (s PluginResponseRewriteType) Name() string { return "response-rewrite" }

func (s PluginResponseRewriteType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginResponseRewriteType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToNumberTypeValue(jsonData, "status_code", &item.StatusCode)
	utils.MapValueToStringTypeValue(jsonData, "body", &item.Body)
	utils.MapValueToBoolTypeValue(jsonData, "body_base64", &item.BodyBase64)
	utils.MapValueToMapTypeValue(jsonData, "headers", &item.Headers)

	item.Vars = varsMapToState(jsonData)

	pluginsType.ResponseRewrite = &item
}

func (s PluginResponseRewriteType) StateToMap(m map[string]interface{}) {
	var pluginValue = make(map[string]interface{})

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.NumberTypeValueToMap(s.StatusCode, pluginValue, "status_code")
	utils.StringTypeValueToMap(s.Body, pluginValue, "body")
	utils.BoolTypeValueToMap(s.BodyBase64, pluginValue, "body_base64")
	utils.MapTypeValueToMap(s.Headers, pluginValue, "headers")

	varsStateToMap(s.Vars, pluginValue)

	m[s.Name()] = pluginValue
}
