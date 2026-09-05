package vpnintranetapplication

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnintranetapplicationDataSourceModel is the data-source-specific model,
// decoupled from VpnintranetapplicationResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnintranetapplicationDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Intranetapplication types.String `tfsdk:"intranetapplication"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Clientapplication types.List   `tfsdk:"clientapplication"`
	Destip            types.String `tfsdk:"destip"`
	Destport          types.String `tfsdk:"destport"`
	Hostname          types.String `tfsdk:"hostname"`
	Interception      types.String `tfsdk:"interception"`
	Iprange           types.String `tfsdk:"iprange"`
	Netmask           types.String `tfsdk:"netmask"`
	Protocol          types.String `tfsdk:"protocol"`
	Spoofiip          types.String `tfsdk:"spoofiip"`
	Srcip             types.String `tfsdk:"srcip"`
	Srcport           types.Int64  `tfsdk:"srcport"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnintranetapplication.json). Never settable; populated
	// from GET.
	Ipaddress types.String `tfsdk:"ipaddress"`
}

func VpnintranetapplicationDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientapplication": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Names of the client applications, such as PuTTY and Xshell.",
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.",
			},
			"destport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.",
			},
			"interception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.",
			},
			"intranetapplication": schema.StringAttribute{
				Required:    true,
				Description: "Name of the intranet application.",
			},
			"iprange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination subnet mask for the intranet application.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.",
			},
			"spoofiip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address that the intranet application will use to route the connection through the virtual adapter.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.",
			},
			"srcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address for the application. This address is the real application server IP address.",
			},
		},
	}
}

// vpnintranetapplicationDataSourceSetAttrFromGet projects a NITRO
// vpnintranetapplication GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func vpnintranetapplicationDataSourceSetAttrFromGet(ctx context.Context, data *VpnintranetapplicationDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnintranetapplicationDataSourceSetAttrFromGet Function")

	if v, ok := g["intranetapplication"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Intranetapplication = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Clientapplication = utils.MapGetStringList(g, "clientapplication")
	data.Destip = utils.MapGetString(g, "destip")
	data.Destport = utils.MapGetString(g, "destport")
	data.Hostname = utils.MapGetString(g, "hostname")
	data.Interception = utils.MapGetString(g, "interception")
	data.Iprange = utils.MapGetString(g, "iprange")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Spoofiip = utils.MapGetString(g, "spoofiip")
	data.Srcip = utils.MapGetString(g, "srcip")
	data.Srcport = utils.MapGetInt64(g, "srcport")

	// Read-only metadata.
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
}
