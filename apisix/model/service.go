package model

import (
	"context"

	"encoding/json"

	"github.com/holubovskyi/apisix-client-go"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SSLCertificateResourceModel maps the resource schema data.
type ServiceResourceModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"desc"`
	EnableWebsocket types.Bool   `tfsdk:"enable_websocket"`
	Hosts           types.List   `tfsdk:"hosts"`
	Labels          types.Map    `tfsdk:"labels"`
	Plugins         types.String `tfsdk:"plugins"`
	UpstreamId      types.String `tfsdk:"upstream_id"`
}

var ServiceSchema = schema.Schema{
	Description: "Manages services",
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "Identifier of the service.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Description: "Identifier for the service.",
			Optional:    true,
		},
		"desc": schema.StringAttribute{
			Description: "Description of usage scenarios.",
			Optional:    true,
		},
		"enable_websocket": schema.BoolAttribute{
			MarkdownDescription: "Enables a websocket. Set to `false` by default.",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"hosts": schema.ListAttribute{
			MarkdownDescription: "Matches with any one of the multiple `hosts` specified in the form of a non-empty list.",
			ElementType:         types.StringType,
			Optional:            true,
		},
		"labels": schema.MapAttribute{
			Description: "Attributes of the Service specified as key-value pairs.",
			ElementType: types.StringType,
			Optional:    true,
		},
		"plugins": schema.StringAttribute{
			Description: "Plugins that are executed during the request/response cycle.",
			Optional:    true,
		},
		"upstream_id": schema.StringAttribute{
			Description: "Id of the Upstream service.",
			Required:    true,
		},
	},
}

func ServiceFromTerraformToApi(ctx context.Context, terraformDataModel *ServiceResourceModel) (apiDataModel api_client.Service) {
	apiDataModel.Name = terraformDataModel.Name.ValueStringPointer()
	apiDataModel.Description = terraformDataModel.Description.ValueStringPointer()
	apiDataModel.EnableWebsocket = terraformDataModel.EnableWebsocket.ValueBoolPointer()
	apiDataModel.UpstreamId = terraformDataModel.UpstreamId.ValueStringPointer()

	_ = terraformDataModel.Hosts.ElementsAs(ctx, &apiDataModel.Hosts, true)
	_ = terraformDataModel.Labels.ElementsAs(ctx, &apiDataModel.Labels, true)

	if !terraformDataModel.Plugins.IsNull() {
		tflog.Info(ctx, "Plugins configuration is defined")

		data := []byte(terraformDataModel.Plugins.ValueString())
		err := json.Unmarshal(data, &apiDataModel.Plugins)
		if err != nil {
			tflog.Error(ctx, "Error", map[string]interface{ any }{
				"Error converting plugins to json": err,
			})
			panic(err)
		}

	}

	tflog.Info(ctx, "Result of ServiceFromTerraformToApi", map[string]interface{ any }{
		"Name":            apiDataModel.Name,
		"Description":     apiDataModel.Description,
		"EnableWebsocket": apiDataModel.EnableWebsocket,
		"UpstreamId":      apiDataModel.UpstreamId,
		"Hosts":           apiDataModel.Hosts,
		"Labels":          apiDataModel.Labels,
		"Plugins":         apiDataModel.Plugins,
	})

	return apiDataModel
}

func ServiceFromTerraformToApiUpdate(ctx context.Context, terraformDataModel *ServiceResourceModel) (apiDataModel api_client.ServiceUpdate) {
	apiDataModel.Name = terraformDataModel.Name.ValueStringPointer()
	apiDataModel.Description = terraformDataModel.Description.ValueStringPointer()
	apiDataModel.EnableWebsocket = terraformDataModel.EnableWebsocket.ValueBoolPointer()
	apiDataModel.UpstreamId = terraformDataModel.UpstreamId.ValueStringPointer()

	_ = terraformDataModel.Hosts.ElementsAs(ctx, &apiDataModel.Hosts, true)
	_ = terraformDataModel.Labels.ElementsAs(ctx, &apiDataModel.Labels, true)

	if !terraformDataModel.Plugins.IsNull() {
		tflog.Info(ctx, "Plugins configuration is defined")

		data := []byte(terraformDataModel.Plugins.ValueString())
		err := json.Unmarshal(data, &apiDataModel.Plugins)
		if err != nil {
			tflog.Error(ctx, "Error", map[string]interface{ any }{
				"Error converting plugins to json": err,
			})
			panic(err)
		}

	}

	tflog.Info(ctx, "ServiceFromTerraformToApiUpdate", map[string]interface{ any }{
		"Name":            apiDataModel.Name,
		"Description":     apiDataModel.Description,
		"EnableWebsocket": apiDataModel.EnableWebsocket,
		"UpstreamId":      apiDataModel.UpstreamId,
		"Hosts":           apiDataModel.Hosts,
		"Labels":          apiDataModel.Labels,
		"Plugins":         apiDataModel.Plugins,
	})

	return apiDataModel
}

func ServiceFromApiToTerraform(ctx context.Context, apiDataModel *api_client.Service) (terraformDataModel ServiceResourceModel) {
	terraformDataModel.ID = types.StringPointerValue(apiDataModel.ID)
	terraformDataModel.Name = types.StringPointerValue(apiDataModel.Name)
	terraformDataModel.Description = types.StringPointerValue(apiDataModel.Description)
	terraformDataModel.EnableWebsocket = types.BoolPointerValue(apiDataModel.EnableWebsocket)
	terraformDataModel.UpstreamId = types.StringPointerValue(apiDataModel.UpstreamId)

	terraformDataModel.Hosts, _ = types.ListValueFrom(ctx, types.StringType, apiDataModel.Hosts)
	terraformDataModel.Labels, _ = types.MapValueFrom(ctx, types.StringType, apiDataModel.Labels)

	if !(apiDataModel.Plugins == nil) {
		data, err := json.Marshal(apiDataModel.Plugins)
		if err != nil {
			tflog.Error(ctx, "Error converting plugins to terraform values")
			panic(err)
		}

		jsonStr := string(data)
		terraformDataModel.Plugins = types.StringValue(jsonStr)
		tflog.Info(ctx, "Result of ServiceFromApiToTerraform", map[string]interface{ any }{
			"jsonSrt": jsonStr,
		})
	} else {
		terraformDataModel.Plugins = types.StringNull()
	}

	tflog.Info(ctx, "Result of ServiceFromApiToTerraform", map[string]interface{ any }{
		"API":       apiDataModel.Plugins,
		"terraform": terraformDataModel.Plugins,
	})
	return terraformDataModel
}
