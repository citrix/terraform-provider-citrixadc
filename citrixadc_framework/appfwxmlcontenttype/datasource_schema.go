package appfwxmlcontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwxmlcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is field name a regular expression?",
			},
			"xmlcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as XML",
			},
		},
	}
}
