package dnsaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DnsactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns action.",
			},
			"actiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of DNS action that is being configured.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the transaction for which the action is chosen",
			},
			"ipaddress": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type.\n        In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.",
			},
			"preferredloclist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The location list in priority order used for the given action.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to live, in seconds.",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The view name that must be used for the given action.",
			},
		},
	}
}
