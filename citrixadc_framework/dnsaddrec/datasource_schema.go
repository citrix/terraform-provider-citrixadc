package dnsaddrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsaddrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached address records need to be removed.",
			},
			"hostname": schema.StringAttribute{
				Required:    true,
				Description: "Domain name.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "One or more IPv4 addresses to assign to the domain name.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}
