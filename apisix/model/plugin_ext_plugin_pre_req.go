package model

import (
	"terraform-provider-apisix/apisix/plan_modifier"
	"terraform-provider-apisix/apisix/utils"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PluginExtPluginPreReqType struct {
	Disable types.Bool                      `tfsdk:"disable"`
	Config  []PluginExtPluginPreReqConfType `tfsdk:"conf"`
}

type PluginExtPluginPreReqConfType struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var PluginExtPluginPreReqSchemaAttribute = tfsdk.Attribute{
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
		"conf": {
			Required: true,
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"name": {
					Required: true,
					Type:     types.StringType,
				},
				"value": {
					Required: true,
					Type:     types.StringType,
				},
			}, tfsdk.ListNestedAttributesOptions{}),
		},
	}),
}

func (s PluginExtPluginPreReqType) Name() string { return "ext-plugin-pre-req" }

func (s PluginExtPluginPreReqType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginExtPluginPreReqType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)

	var subItems []PluginExtPluginPreReqConfType
	if v := jsonData["conf"]; v != nil {
		for _, vv := range v.([]interface{}) {
			subItem := PluginExtPluginPreReqConfType{}
			subV := vv.(map[string]interface{})
			utils.MapValueToStringTypeValue(subV, "name", &subItem.Name)
			utils.MapValueToStringTypeValue(subV, "value", &subItem.Value)
			subItems = append(subItems, subItem)
		}
	}

	item.Config = subItems
	pluginsType.ExtPluginPreReqType = &item
}

func (s PluginExtPluginPreReqType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")

	var subItems []map[string]interface{}
	for _, vv := range s.Config {
		subItem := make(map[string]interface{})
		utils.StringTypeValueToMap(vv.Name, subItem, "name")
		utils.StringTypeValueToMap(vv.Value, subItem, "value")
		subItems = append(subItems, subItem)
	}

	pluginValue["conf"] = subItems
	m[s.Name()] = pluginValue
}
