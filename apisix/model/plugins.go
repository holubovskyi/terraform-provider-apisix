package model

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func PluginsStringToJson(ctx context.Context, srt types.String) (jsonPointer *map[string]interface{}) {

	if srt.IsNull() {
		return nil
	}

	var result map[string]interface{}
	bytes := []byte(srt.ValueString())

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		tflog.Error(ctx, "Error", map[string]interface{ any }{
			"Error converting plugins to json": err,
			"Input string is":                  srt.ValueString(),
		})
		panic(err)
	}

	return &result
}

func PluginsFromJsonToString(ctx context.Context, jsonPointer *map[string]interface{}) (str types.String) {
	if jsonPointer == nil {
		return types.StringNull()
	}

	data, err := json.Marshal(jsonPointer)
	if err != nil {
		tflog.Error(ctx, "Error converting plugins to terraform values")
		panic(err)
	}

	jsonStr := string(data)

	return types.StringValue(jsonStr)

}
