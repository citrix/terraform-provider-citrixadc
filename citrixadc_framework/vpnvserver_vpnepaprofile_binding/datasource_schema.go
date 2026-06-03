package vpnvserver_vpnepaprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverVpnepaprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"epaprofile": schema.StringAttribute{
				Required:    true,
				Description: "Advanced EPA profile to bind",
			},
			"epaprofileoptional": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark the EPA profile optional for preauthentication EPA profile. User would be shown a logon page even if the EPA profile fails to evaluate.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}
