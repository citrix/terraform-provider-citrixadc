package rnat6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Rnat6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acl6name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as an RNAT6 rule.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RNAT6 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT6 rule.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address of the network on whose traffic you want the Citrix ADC to do RNAT processing.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
			"redirectport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to which the IPv6 packets are redirected. Applicable to TCP and UDP protocols.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
