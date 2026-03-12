package snmpcommunity

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpcommunityDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"communityname": schema.StringAttribute{
				Required:    true,
				Description: "The SNMP community string. Can consist of 1 to 31 characters that include uppercase and lowercase letters,numbers and special characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the string includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my string\" or 'my string').",
			},
			"permissions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The SNMP V1 or V2 query-type privilege that you want to associate with this SNMP community.",
			},
		},
	}
}
