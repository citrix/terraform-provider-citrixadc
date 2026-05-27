package appfwgrpcwebtextcontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwgrpcwebtextcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"grpcwebtextcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as gRPC-web-text",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is gRPC-web-text content type a regular expression?",
			},
		},
	}
}
