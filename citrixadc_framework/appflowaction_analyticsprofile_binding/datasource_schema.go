package appflowaction_analyticsprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppflowactionAnalyticsprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"analyticsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Analytics profile to be bound to the appflow action",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
		},
	}
}
