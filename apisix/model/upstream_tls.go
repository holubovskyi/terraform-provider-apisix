package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UpstreamTLSType struct {
	ClientCert types.String `tfsdk:"client_cert"`
	ClientKey  types.String `tfsdk:"client_key"`
}

var UpstreamTLSSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Sets the client certificate while connecting to a TLS Upstream.",
	Optional:            true,
	Attributes: map[string]schema.Attribute{
		"client_cert": schema.Int64Attribute{
			MarkdownDescription: "Sets the client certificate while connecting to a TLS Upstream.",
			Required:            true,
		},
		"client_key": schema.Int64Attribute{
			MarkdownDescription: "Sets the client private key while connecting to a TLS Upstream.",
			Required:            true,
		},
	},
}
