package appfwgrpccontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwgrpccontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"grpccontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as gRPC",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is gRPC content type a regular expression?",
			},
		},
	}
}
