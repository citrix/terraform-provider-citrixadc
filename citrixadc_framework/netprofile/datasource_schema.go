package netprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NetprofileDataSourceModel is the data-source-specific model, decoupled from
// NetprofileResourceModel. A data source is a pure read surface, so it exposes
// the read/write attributes (as Computed outputs) plus the read-only (GET-only)
// attributes the resource deliberately omits.
type NetprofileDataSourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Badipactionthreshold           types.Int64  `tfsdk:"badipactionthreshold"`
	Mbf                            types.String `tfsdk:"mbf"`
	Name                           types.String `tfsdk:"name"`
	Overridelsn                    types.String `tfsdk:"overridelsn"`
	Proxyprotocol                  types.String `tfsdk:"proxyprotocol"`
	Proxyprotocolaftertlshandshake types.String `tfsdk:"proxyprotocolaftertlshandshake"`
	Proxyprotocoltxversion         types.String `tfsdk:"proxyprotocoltxversion"`
	Srcip                          types.String `tfsdk:"srcip"`
	Srcippersistency               types.String `tfsdk:"srcippersistency"`
	Td                             types.Int64  `tfsdk:"td"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/netprofile.json). Never settable; populated from GET.
	Proxyprotocoltlvoptions types.List `tfsdk:"proxyprotocoltlvoptions"`
}

func NetprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"badipactionthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of protocol violation from an IP address before taking action. Default value: 0 Minimum value =  0 Maximum value =  100000",
			},
			"mbf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.",
			},
			"overridelsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "USNIP/USIP settings override LSN settings for configured\n              service/virtual server traffic..",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Action (Enabled/Disabled)",
			},
			"proxyprotocolaftertlshandshake": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake",
			},
			"proxyprotocoltxversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Version (V1/V2)",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or the name of an IP set.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"proxyprotocoltlvoptions": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Proxy protocol TLV options (for example cert-cn).",
			},
		},
	}
}

// netprofileDataSourceSetAttrFromGet projects a NITRO netprofile GET response
// onto the data-source model using the shared utils.MapGet* helpers.
func netprofileDataSourceSetAttrFromGet(ctx context.Context, data *NetprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In netprofileDataSourceSetAttrFromGet Function")

	data.Badipactionthreshold = utils.MapGetInt64(g, "badipactionthreshold")
	data.Mbf = utils.MapGetString(g, "mbf")
	data.Name = utils.MapGetString(g, "name")
	data.Overridelsn = utils.MapGetString(g, "overridelsn")
	data.Proxyprotocol = utils.MapGetString(g, "proxyprotocol")
	data.Proxyprotocolaftertlshandshake = utils.MapGetString(g, "proxyprotocolaftertlshandshake")
	data.Proxyprotocoltxversion = utils.MapGetString(g, "proxyprotocoltxversion")
	data.Srcip = utils.MapGetString(g, "srcip")
	data.Srcippersistency = utils.MapGetString(g, "srcippersistency")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}

	// Read-only (GET-only) attribute.
	data.Proxyprotocoltlvoptions = utils.MapGetStringList(g, "proxyprotocoltlvoptions")

	// Set ID from the single unique key (name).
	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = types.StringValue(data.Name.ValueString())
	}
}
