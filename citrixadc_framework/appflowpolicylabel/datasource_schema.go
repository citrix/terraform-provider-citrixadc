package appflowpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppflowpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AppFlow policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policylabel\" or 'my appflow policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\n                    The following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow policylabel\" or 'my appflow policylabel')",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of traffic evaluated by the policies bound to the policy label.",
			},
		},
	}
}
