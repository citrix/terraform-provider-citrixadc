package dnsnaptrrec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsnaptrrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Name of the domain for the NAPTR record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached NAPTR record need to be removed.",
			},
			"flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "flags for this NAPTR.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"preference": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"recordid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying\ndomain name and all other naptr record attributes as was supplied during the add command.",
			},
			"regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The regular expression, that specifies the substitution expression for this NAPTR",
			},
			"replacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The replacement domain name for this NAPTR.",
			},
			"services": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service Parameters applicable to this delegation path.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}
