package appfwjsoncontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwjsoncontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is json content type a regular expression?",
			},
			"jsoncontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as JSON",
			},
		},
	}
}
