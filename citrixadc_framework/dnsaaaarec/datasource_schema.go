package dnsaaaarec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsaaaarecDataSourceModel is the data-source-specific model, decoupled from
// DnsaaaarecResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type DnsaaaarecDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ecssubnet   types.String `tfsdk:"ecssubnet"`
	Hostname    types.String `tfsdk:"hostname"`
	Ipv6address types.String `tfsdk:"ipv6address"`
	Nodeid      types.Int64  `tfsdk:"nodeid"`
	Ttl         types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsaaaarec.json). Never settable; populated from GET.
	Vservername types.String `tfsdk:"vservername"`
	Authtype    types.String `tfsdk:"authtype"`
}

func DnsaaaarecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached records need to be removed.",
			},
			"hostname": schema.StringAttribute{
				Required:    true,
				Description: "Domain name.",
			},
			"ipv6address": schema.StringAttribute{
				Required:    true,
				Description: "One or more IPv6 addresses to assign to the domain name.",
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

// dnsaaaarecDataSourceSetAttrFromGet projects a NITRO dnsaaaarec GET response
// onto the data-source model. Attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func dnsaaaarecDataSourceSetAttrFromGet(ctx context.Context, data *DnsaaaarecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsaaaarecDataSourceSetAttrFromGet Function")

	if v, ok := g["hostname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Hostname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Ipv6address = utils.MapGetString(g, "ipv6address")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Vservername = utils.MapGetString(g, "vservername")
	data.Authtype = utils.MapGetString(g, "authtype")
}
