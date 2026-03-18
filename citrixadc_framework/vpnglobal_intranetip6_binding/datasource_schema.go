package vpnglobal_intranetip6_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnglobalIntranetip6BindingDataSourceSchema() schema.Schema {
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
			"intranetip6": schema.StringAttribute{
				Required:    true,
				Description: "The intranet ip address or range.",
			},
			"numaddr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}
