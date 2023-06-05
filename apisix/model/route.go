package model

import (
	"reflect"
	"terraform-provider-apisix/apisix/common"
	"terraform-provider-apisix/apisix/plan_modifier"
	"terraform-provider-apisix/apisix/utils"
	"terraform-provider-apisix/apisix/validator"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RouteType struct {
	ID              types.String  `tfsdk:"id"`
	Description     types.String  `tfsdk:"desc"`
	EnableWebsocket types.Bool    `tfsdk:"enable_websocket"`
	FilterFunc      types.String  `tfsdk:"filter_func"`
	Host            types.String  `tfsdk:"host"`
	Hosts           types.List    `tfsdk:"hosts"`
	IsEnabled       types.Bool    `tfsdk:"is_enabled"`
	Labels          types.Map     `tfsdk:"labels"`
	Methods         types.List    `tfsdk:"methods"`
	Name            types.String  `tfsdk:"name"`
	PluginConfigId  types.String  `tfsdk:"plugin_config_id"`
	Plugins         *PluginsType  `tfsdk:"plugins"`
	Priority        types.Number  `tfsdk:"priority"`
	RemoteAddr      types.String  `tfsdk:"remote_addr"`
	RemoteAddrs     types.List    `tfsdk:"remote_addrs"`
	Script          types.String  `tfsdk:"script"`
	ServiceId       types.String  `tfsdk:"service_id"`
	URI             types.String  `tfsdk:"uri"`
	URIS            types.List    `tfsdk:"uris"`
	Upstream        *UpstreamType `tfsdk:"upstream"`
	UpstreamId      types.String  `tfsdk:"upstream_id"`
	Vars            types.String  `tfsdk:"vars"`
}

var RouteSchema = tfsdk.Schema{
	Version: 8,
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"desc": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("Managed by Terraform"),
			},
		},
		"uri": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("uris"),
			},
		},
		"uris": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("uri"),
			},
		},
		"host": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("hosts"),
			},
		},
		"hosts": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("host"),
			},
		},
		"remote_addr": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("remote_addrs"),
			},
		},
		"remote_addrs": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("remote_addr"),
			},
		},
		"methods": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringOfStringInSlice(common.HttpMethods...),
			},
		},
		"priority": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(0),
			},
		},

		"is_enabled": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
		},
		"enable_websocket": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"service_id": {
			Type:     types.StringType,
			Optional: true,
		},
		"upstream_id": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("upstream"),
			},
		},
		//
		"labels": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},

		//"timeout": TimeoutSchemaAttribute,
		"script": {
			Optional: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("plugin_config_id"),
			},
		},
		"plugin_config_id": {
			Optional: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("script"),
			},
		},
		"filter_func": {
			Optional: true,
			Type:     types.StringType,
		},
		"plugins": {
			Optional:   true,
			Attributes: PluginsSchemaAttribute,
		},
		"upstream": UpstreamSchemaSAttribute,
		"vars": {
			Optional: true,
			Type:     types.StringType,
		},
	},
}

func RouteTypeMapToState(jsonMap map[string]interface{}) (*RouteType, error) {
	newState := RouteType{}

	utils.MapValueToStringTypeValue(jsonMap, "id", &newState.ID)
	utils.MapValueToStringTypeValue(jsonMap, "name", &newState.Name)
	utils.MapValueToStringTypeValue(jsonMap, "desc", &newState.Description)
	utils.MapValueToStringTypeValue(jsonMap, "uri", &newState.URI)
	utils.MapValueToListTypeValue(jsonMap, "uris", &newState.URIS)
	utils.MapValueToStringTypeValue(jsonMap, "host", &newState.Host)
	utils.MapValueToListTypeValue(jsonMap, "hosts", &newState.Hosts)
	utils.MapValueToStringTypeValue(jsonMap, "remote_addr", &newState.RemoteAddr)
	utils.MapValueToListTypeValue(jsonMap, "remote_addrs", &newState.RemoteAddrs)
	utils.MapValueToListTypeValue(jsonMap, "methods", &newState.Methods)
	utils.MapValueToNumberTypeValue(jsonMap, "priority", &newState.Priority)
	utils.MapValueToStringTypeValue(jsonMap, "filter_func", &newState.FilterFunc)
	utils.MapValueToStringTypeValue(jsonMap, "script", &newState.Script)
	utils.MapValueToStringTypeValue(jsonMap, "upstream_id", &newState.UpstreamId)
	utils.MapValueToStringTypeValue(jsonMap, "service_id", &newState.ServiceId)
	utils.MapValueToStringTypeValue(jsonMap, "plugin_config_id", &newState.PluginConfigId)
	utils.MapValueToBoolTypeValue(jsonMap, "enable_websocket", &newState.EnableWebsocket)
	newState.Vars = varsMapToState(jsonMap)

	if v := jsonMap["status"]; v != nil {
		if v.(float64) == 1 {
			newState.IsEnabled = types.Bool{Value: true}
		} else {
			newState.IsEnabled = types.Bool{Value: false}
		}
	} else {
		newState.IsEnabled = types.Bool{Null: true}
	}

	utils.MapValueToMapTypeValue(jsonMap, "labels", &newState.Labels)

	upstream, err := UpstreamTypeMapToState(jsonMap, false)
	if err != nil {
		return nil, err
	}
	newState.Upstream = upstream

	if v := jsonMap["plugins"]; v != nil {
		value := v.(map[string]interface{})
		pluginsType := PluginsType{}

		e := reflect.ValueOf(&pluginsType).Elem()
		for i := 0; i < e.NumField(); i++ {
			switch e.Field(i).Interface().(type) {
			case PluginCommonInterface:
				reflect.New(e.Type().Field(i).Type.Elem()).Interface().(PluginCommonInterface).MapToState(value, &pluginsType)
			default:

			}
		}

		//PluginCustomTypeMapToState(value, &pluginsType, plan, state)
		newState.Plugins = &pluginsType
	} else {
		newState.Plugins = nil
	}

	return &newState, nil
}

func RouteTypeStateToMap(plan RouteType) (map[string]interface{}, error) {

	output := make(map[string]interface{})

	utils.StringTypeValueToMap(plan.Name, output, "name")
	utils.StringTypeValueToMap(plan.Description, output, "desc")
	utils.StringTypeValueToMap(plan.URI, output, "uri")
	utils.ListTypeValueToMap(plan.URIS, output, "uris")
	utils.StringTypeValueToMap(plan.Host, output, "host")
	utils.ListTypeValueToMap(plan.Hosts, output, "hosts")
	utils.StringTypeValueToMap(plan.RemoteAddr, output, "remote_addr")
	utils.ListTypeValueToMap(plan.RemoteAddrs, output, "remote_addrs")
	utils.ListTypeValueToMap(plan.Methods, output, "methods")
	utils.NumberTypeValueToMap(plan.Priority, output, "priority")

	if !plan.IsEnabled.Null {
		if plan.IsEnabled.Value {
			output["status"] = 1
		} else {
			output["status"] = 0
		}
	}

	utils.BoolTypeValueToMap(plan.EnableWebsocket, output, "enable_websocket")
	utils.StringTypeValueToMap(plan.ServiceId, output, "service_id")
	utils.StringTypeValueToMap(plan.UpstreamId, output, "upstream_id")
	utils.MapTypeValueToMap(plan.Labels, output, "labels")
	utils.StringTypeValueToMap(plan.Script, output, "script")
	utils.StringTypeValueToMap(plan.PluginConfigId, output, "plugin_config_id")
	utils.StringTypeValueToMap(plan.FilterFunc, output, "filter_func")

	varsStateToMap(plan.Vars, output)

	plugins := make(map[string]interface{})
	if plan.Plugins != nil {
		planPlugins := plan.Plugins

		e := reflect.ValueOf(planPlugins).Elem()
		for i := 0; i < e.NumField(); i++ {

			if !e.Field(i).IsNil() {
				switch e.Field(i).Interface().(type) {
				case PluginCommonInterface:
					e.Field(i).Interface().(PluginCommonInterface).StateToMap(plugins)
				default:

				}

			}
			//else if isUpdate {
			//	switch e.Field(i).Interface().(type) {
			//	case PluginCommonInterface:
			//		plugins[reflect.New(e.Type().Field(i).Type.Elem()).Interface().(PluginCommonInterface).Name()] = nil
			//	default:
			//	}
			//}
		}

		//PluginCustomTypeStateToMap(plugins, plan, state, isUpdate)

		output["plugins"] = plugins
	}

	upstream, err := UpstreamTypeStateToMap(plan.Upstream)

	if err != nil {
		return nil, err
	}

	if upstream != nil {
		output["upstream"] = upstream
	}

	return output, nil
}
