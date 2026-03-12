package vpnglobal_staserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnglobalStaserverBindingDataSourceSchema() schema.Schema {
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
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},
		},
	}
}
