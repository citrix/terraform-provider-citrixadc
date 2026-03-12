package spilloveraction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SpilloveractionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Spillover action. Currently only type SPILLOVER is supported",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the spillover action.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters. \nChoose a name that can be correlated with the function that the action performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}
