package forwardingsession

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ForwardingsessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acl6name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as a forwarding session rule.",
			},
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured ACL whose action is ALLOW. The rule of the ACL is used as a forwarding session rule.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the forwarding session.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the forwarding session rule. Can begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rule\" or 'my rule').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask associated with the network.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An IPv4 network address or IPv6 prefix of a network from which the forwarded traffic originates or to which it is destined.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabling this option on forwarding session will not steer the packet to flow processor. Instead, packet will be routed.",
			},
			"sourceroutecache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache the source ip address and mac address of the DA servers.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
