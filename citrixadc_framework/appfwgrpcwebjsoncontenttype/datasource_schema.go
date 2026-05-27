package appfwgrpcwebjsoncontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwgrpcwebjsoncontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"grpcwebjsoncontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as gRPC-web-json",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is gRPC-web-json content type a regular expression?",
			},
		},
	}
}
