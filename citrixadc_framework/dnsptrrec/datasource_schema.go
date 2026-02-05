package dnsptrrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsptrrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to configure reverse mapping.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached PTR record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"reversedomain": schema.StringAttribute{
				Required:    true,
				Description: "Reversed domain name representation of the IPv4 or IPv6 address for which to create the PTR record. Use the \"in-addr.arpa.\" suffix for IPv4 addresses and the \"ip6.arpa.\" suffix for IPv6 addresses.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}
