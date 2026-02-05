package nsservicefunction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsservicefunctionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ingressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN ID on which the traffic from service function reaches Citrix ADC.",
			},
			"servicefunctionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service function to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
		},
	}
}
