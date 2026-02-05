package dnsmxrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsmxrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to add the MX record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached MX record need to be removed.",
			},
			"mx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host name of the mail exchange server.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"pref": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}
