package crvserver_analyticsprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CrvserverAnalyticsprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"analyticsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the analytics profile bound to the CR vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
		},
	}
}
