package model

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type UpstreamChecksPassiveType struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `tfsdk:"healthy"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `tfsdk:"unhealthy"`
}

var UpstreamChecksPassiveSchemaAttribute = schema.SingleNestedAttribute{
	MarkdownDescription: "Passive health check refers to judging whether the corresponding upstream node is healthy by judging the response status of the request forwarded from APISIX to the upstream node.",
	Optional:            true,
	Attributes: map[string]schema.Attribute{
		"healthy":   UpstreamChecksPassiveHealthySchemaAttribute,
		"unhealthy": UpstreamChecksPassiveUnhealthySchemaAttribute,
	},
}
