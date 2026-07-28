package dnscaarec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnscaarecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name of the CAA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached CAA record need to be removed.",
			},
			"flag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag associated with the CAA record.",
			},
			"recordid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique, internally generated record ID. View the details of the CAA record to obtain its record ID. Records can be removedby either specifying the domain name and record id OR by specifying domain name and all other CAA record attributes as was supplied during the add command.",
			},
			"tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - issue, issuwild and iodef.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"valuestring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value associated with the chosen property tag in the CAA resource record. Enclose the string in single or double quotation marks.",
			},
		},
	}
}
