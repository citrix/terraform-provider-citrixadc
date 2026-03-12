package rsskeytype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RsskeytypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rsstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of RSS key, possible values are SYMMETRIC and ASYMMETRIC.",
			},
		},
	}
}
