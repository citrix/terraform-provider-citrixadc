package dnsptrrec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsptrrecDataSourceModel is the data-source-specific model, decoupled from
// DnsptrrecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (authtype). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type DnsptrrecDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Domain        types.String `tfsdk:"domain"`
	Ecssubnet     types.String `tfsdk:"ecssubnet"`
	Nodeid        types.Int64  `tfsdk:"nodeid"`
	Reversedomain types.String `tfsdk:"reversedomain"` // Required lookup key
	Ttl           types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnsptrrec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

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

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnsptrrecDataSourceSetAttrFromGet projects a NITRO dnsptrrec GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnsptrrecDataSourceSetAttrFromGet(ctx context.Context, data *DnsptrrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsptrrecDataSourceSetAttrFromGet Function")

	if v, ok := g["reversedomain"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Reversedomain = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Domain = utils.MapGetString(g, "domain")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only metadata.
	data.Authtype = utils.MapGetString(g, "authtype")
}
