package dnscnamerec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnscnamerecDataSourceModel is the data-source-specific model, decoupled from
// DnscnamerecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type DnscnamerecDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Aliasname     types.String `tfsdk:"aliasname"` // Required lookup key
	Canonicalname types.String `tfsdk:"canonicalname"`
	Ecssubnet     types.String `tfsdk:"ecssubnet"`
	Nodeid        types.Int64  `tfsdk:"nodeid"`
	Ttl           types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnscnamerec.json). Never settable; populated from GET.
	Vservername types.String `tfsdk:"vservername"`
	Authtype    types.String `tfsdk:"authtype"`
}

func DnscnamerecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aliasname": schema.StringAttribute{
				Required:    true,
				Description: "Alias for the canonical domain name.",
			},
			"canonicalname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Canonical domain name.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached CNAME record need to be removed.",
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

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"vservername": schema.StringAttribute{
				Computed:    true,
				Description: "GSLB Virtual server name to which this domain is bound.",
			},
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Record type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnscnamerecDataSourceSetAttrFromGet projects a NITRO dnscnamerec GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func dnscnamerecDataSourceSetAttrFromGet(ctx context.Context, data *DnscnamerecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnscnamerecDataSourceSetAttrFromGet Function")

	if v, ok := g["aliasname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Aliasname = types.StringValue(utils.AnyToString(v))
	}

	data.Canonicalname = utils.MapGetString(g, "canonicalname")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Vservername = utils.MapGetString(g, "vservername")
	data.Authtype = utils.MapGetString(g, "authtype")
}
