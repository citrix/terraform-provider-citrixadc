package vpnvserver_appfwpolicy_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnvserverAppfwpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a specific vpnvserver_appfwpolicy_binding.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Terraform ID. Format: <name>,<policy>",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy, if any, bound to the VPN virtual server.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Computed:    true,
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers.",
			},
			"secondary": schema.BoolAttribute{
				Computed:    true,
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method.",
			},
		},
	}
}
