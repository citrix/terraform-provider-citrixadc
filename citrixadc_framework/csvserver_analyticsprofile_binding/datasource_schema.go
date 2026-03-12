package csvserver_analyticsprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CsvserverAnalyticsprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"analyticsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the analytics profile bound to the LB vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
		},
	}
}
