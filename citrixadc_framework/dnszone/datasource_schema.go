package dnszone

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnszoneDataSourceModel is the data-source-specific model, decoupled from
// DnszoneResourceModel. A data source is a pure read surface, so it can expose
// the FULL GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (flags).
type DnszoneDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Zonename types.String `tfsdk:"zonename"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Dnssecoffload types.String `tfsdk:"dnssecoffload"`
	Keyname       types.List   `tfsdk:"keyname"`
	Nsec          types.String `tfsdk:"nsec"`
	Proxymode     types.String `tfsdk:"proxymode"`
	Type          types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnszone.json). Never settable; populated from GET.
	Flags types.Int64 `tfsdk:"flags"`
}

func DnszoneDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dnssecoffload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dnssec offload for this zone.",
			},
			"keyname": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.",
			},
			"nsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable nsec generation for dnssec offload.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deploy the zone in proxy mode. Enable in the following scenarios:\n* The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.\n* The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers.\nIn either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.\nDisable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows:\n* ADNS - Display all the zones for which the Citrix ADC is authoritative.\n* PROXY - Display all the zones for which the Citrix ADC is functioning as a proxy server.\n* ALL - Display all the zones configured on the appliance.",
			},
			"zonename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the zone to create.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags controlling display.",
			},
		},
	}
}

// dnszoneDataSourceSetAttrFromGet projects a NITRO dnszone GET response onto the
// data-source model. Attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func dnszoneDataSourceSetAttrFromGet(ctx context.Context, data *DnszoneDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnszoneDataSourceSetAttrFromGet Function")

	if v, ok := g["zonename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Zonename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Dnssecoffload = utils.MapGetString(g, "dnssecoffload")
	data.Keyname = utils.MapGetStringList(g, "keyname")
	data.Nsec = utils.MapGetString(g, "nsec")
	data.Proxymode = utils.MapGetString(g, "proxymode")
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
}
