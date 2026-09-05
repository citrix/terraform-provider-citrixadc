package inat

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InatDataSourceModel is the data-source-specific model, decoupled from
// InatResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (flags). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type InatDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"` // Required lookup key
	Connfailover types.String `tfsdk:"connfailover"`
	Ftp          types.String `tfsdk:"ftp"`
	Mode         types.String `tfsdk:"mode"`
	Privateip    types.String `tfsdk:"privateip"`
	Proxyip      types.String `tfsdk:"proxyip"`
	Publicip     types.String `tfsdk:"publicip"`
	Tcpproxy     types.String `tfsdk:"tcpproxy"`
	Td           types.Int64  `tfsdk:"td"`
	Tftp         types.String `tfsdk:"tftp"`
	Useproxyport types.String `tfsdk:"useproxyport"`
	Usip         types.String `tfsdk:"usip"`
	Usnip        types.String `tfsdk:"usnip"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/inat.json). Never settable; populated from GET.
	Flags types.Int64 `tfsdk:"flags"`
}

func InatDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the INAT session",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the FTP protocol on the server for transferring files between the client and the server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Stateless translation.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Inbound NAT (INAT) entry. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"privateip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.",
			},
			"proxyip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"tftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To enable/disable TFTP (Default DISABLED).",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server.",
			},
			"usnip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags for different modes.",
			},
		},
	}
}

// inatDataSourceSetAttrFromGet projects a NITRO inat GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func inatDataSourceSetAttrFromGet(ctx context.Context, data *InatDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In inatDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Connfailover = utils.MapGetString(g, "connfailover")
	data.Ftp = utils.MapGetString(g, "ftp")
	data.Mode = utils.MapGetString(g, "mode")
	data.Privateip = utils.MapGetString(g, "privateip")
	data.Proxyip = utils.MapGetString(g, "proxyip")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Tcpproxy = utils.MapGetString(g, "tcpproxy")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Tftp = utils.MapGetString(g, "tftp")
	data.Useproxyport = utils.MapGetString(g, "useproxyport")
	data.Usip = utils.MapGetString(g, "usip")
	data.Usnip = utils.MapGetString(g, "usnip")

	// Read-only attributes.
	data.Flags = utils.MapGetInt64(g, "flags")
}
