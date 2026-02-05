package appfwmultipartformcontenttype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwmultipartformcontenttypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is multipart_form content type a regular expression?",
			},
			"multipartformcontenttypevalue": schema.StringAttribute{
				Required:    true,
				Description: "Content type to be classified as multipart form",
			},
		},
	}
}
