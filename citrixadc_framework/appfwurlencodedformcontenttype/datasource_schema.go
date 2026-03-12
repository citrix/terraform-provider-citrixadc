package appfwurlencodedformcontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwurlencodedformcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is urlencoded form content type a regular expression?",
			},
			"urlencodedformcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as urlencoded form",
			},
		},
	}
}
