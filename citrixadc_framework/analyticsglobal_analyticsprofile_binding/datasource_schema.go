package analyticsglobal_analyticsprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AnalyticsglobalAnalyticsprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"analyticsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the analytics profile bound.",
			},
		},
	}
}
