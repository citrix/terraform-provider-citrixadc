package dnsaddrec

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsaddrecDataSourceModel is the data-source-specific model, decoupled from
// DnsaddrecResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits.
type DnsaddrecDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Hostname  types.String `tfsdk:"hostname"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Ttl       types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsaddrec.json). Never settable; populated from GET.
	Vservername types.String `tfsdk:"vservername"`
	Authtype    types.String `tfsdk:"authtype"`
}

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

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"vservername": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server name.",
			},
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnsaddrecDataSourceSetAttrFromGet projects a NITRO dnsaddrec GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnsaddrecDataSourceSetAttrFromGet(ctx context.Context, data *DnsaddrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsaddrecDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Hostname = utils.MapGetString(g, "hostname")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Vservername = utils.MapGetString(g, "vservername")
	data.Authtype = utils.MapGetString(g, "authtype")

	// Backward-compatible SDK v2 ID format: "hostname,ipaddress".
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Hostname.ValueString(), data.Ipaddress.ValueString()))
}
