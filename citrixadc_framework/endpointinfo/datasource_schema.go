package endpointinfo

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func EndpointinfoDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"endpointkind": schema.StringAttribute{
				Required:    true,
				Description: "Endpoint kind. Currently, IP endpoints are supported",
			},
			"endpointlabelsjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String representing labels in json form. Maximum length 16K",
			},
			"endpointmetadata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of qualifiers, in dotted notation, structured metadata for an endpoint. Each qualifier is more specific than the one that precedes it, as in cluster.namespace.service. For example: cluster.default.frontend. \nNote: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.",
			},
			"endpointname": schema.StringAttribute{
				Required:    true,
				Description: "Name of endpoint, depends on kind. For IP Endpoint - IP address.",
			},
		},
	}
}
