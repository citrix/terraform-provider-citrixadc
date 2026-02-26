package vpnglobal_intranetip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnglobalIntranetipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetip": schema.StringAttribute{
				Required:    true,
				Description: "The intranet ip address or range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}
