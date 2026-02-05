package lbgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period, in minutes, for which backup persistence is in effect.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain attribute for the HTTP cookie.",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"mastervserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the load balancing virtual server group.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the load balancing virtual server group.",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of backup persistence for the group.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the group. Available settings function as follows:\n* SOURCEIP - Create persistence sessions based on the client IP.\n* COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client.\n* RULE - Create persistence sessions based on a user defined rule.\n* NONE - Disable persistence for the group.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which a persistence session is in effect.",
			},
			"usevserverpersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.",
			},
		},
	}
}
