package dnssoarec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnssoarecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.",
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to add the SOA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached SOA record need to be removed.",
			},
			"expire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.",
			},
			"minimum": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default time to live (TTL) for all records in the zone. Can be overridden for individual records.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"originserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the name server that responds authoritatively for the domain.",
			},
			"refresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.",
			},
			"retry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.",
			},
			"serial": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}
