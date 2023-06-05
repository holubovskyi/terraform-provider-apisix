package validator

import (
	"fmt"
	"terraform-provider-apisix/apisix/utils"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/context"
)

type CountOfElementsType struct {
	Min  int
	Max  int
	Type string
}

func (j CountOfElementsType) Description(ctx context.Context) string {
	return ""
}

func (j CountOfElementsType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j CountOfElementsType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {

	if utils.IsAttributeIsNull(request.AttributeConfig) {
		return
	}

	v := len(request.AttributeConfig.(types.List).Elems)

	switch j.Type {
	case "gt":
		if v > j.Min {
			return
		}
	case "gte":
		if v >= j.Min {
			return
		}
	case "lt":
		if v < j.Max {
			return
		}
	case "lte":
		if v <= j.Max {
			return
		}

	}

	response.Diagnostics.AddError(
		fmt.Sprintf("Wrong value total items in array, field: %v", request.AttributePath.String()),
		fmt.Sprintf("Values must be more than: %v, or less than: %v", j.Min, j.Max),
	)
}

func ElementsGreatThan(v int) CountOfElementsType {
	return CountOfElementsType{Min: v, Type: "gt"}
}

func ElementsGreatOrEqualThan(v int) CountOfElementsType {
	return CountOfElementsType{Min: v, Type: "gte"}
}

func ElementsLessThan(v int) CountOfElementsType {
	return CountOfElementsType{Max: v, Type: "lt"}
}

func ElementsLessOrEqualThan(v int) CountOfElementsType {
	return CountOfElementsType{Max: v, Type: "lte"}
}
