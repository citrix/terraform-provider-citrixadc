package dnssrvrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnssrvrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name, which, by convention, is prefixed by the symbolic name of the desired service and the symbolic name of the desired protocol, each with an underscore (_) prepended. For example, if an SRV-aware client wants to discover a SIP service that is provided over UDP, in the domain example.com, the client performs a lookup for _sip._udp.example.com.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached SRV record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the target host listens for client requests.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.",
			},
			"target": schema.StringAttribute{
				Required:    true,
				Description: "Target host for the specified service.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.",
			},
		},
	}
}
