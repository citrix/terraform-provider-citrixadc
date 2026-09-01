package iptunnel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IptunnelDataSourceModel is the data-source-specific model, decoupled from
// IptunnelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (type, encapip, channel, tunneltype, ipsectunnelstatus, refcnt, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type IptunnelDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Destport         types.Int64  `tfsdk:"destport"`
	Grepayload       types.String `tfsdk:"grepayload"`
	Ipsecprofilename types.String `tfsdk:"ipsecprofilename"`
	Local            types.String `tfsdk:"local"`
	Name             types.String `tfsdk:"name"`
	Ownergroup       types.String `tfsdk:"ownergroup"`
	Protocol         types.String `tfsdk:"protocol"`
	Remote           types.String `tfsdk:"remote"`
	Remotesubnetmask types.String `tfsdk:"remotesubnetmask"`
	Tosinherit       types.String `tfsdk:"tosinherit"`
	Vlan             types.Int64  `tfsdk:"vlan"`
	Vlantagging      types.String `tfsdk:"vlantagging"`
	Vnid             types.Int64  `tfsdk:"vnid"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/iptunnel.json). Never settable; populated from GET.
	Sysname           types.String `tfsdk:"sysname"`
	Type              types.Int64  `tfsdk:"type"`
	Encapip           types.String `tfsdk:"encapip"`
	Channel           types.Int64  `tfsdk:"channel"`
	Tunneltype        types.List   `tfsdk:"tunneltype"`
	Ipsectunnelstatus types.String `tfsdk:"ipsectunnelstatus"`
	Refcnt            types.Int64  `tfsdk:"refcnt"`
}

func IptunnelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies UDP destination port for Geneve packets. Default port is 6081.",
			},
			"grepayload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The payload GRE will carry",
			},
			"ipsecprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of IPSec profile to be associated.",
			},
			"local": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for the iptunnel.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the protocol to be used on this tunnel.",
			},
			"remote": schema.StringAttribute{
				Required:    true,
				Description: "Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.",
			},
			"remotesubnetmask": schema.StringAttribute{
				Required:    true,
				Description: "Subnet mask of the remote IP address of the tunnel.",
			},
			"tosinherit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan for mulicast packets",
			},
			"vlantagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to select Vlan Tagging.",
			},
			"vnid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"sysname": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the ip tunnel.",
			},
			"type": schema.Int64Attribute{
				Computed:    true,
				Description: "The type of this tunnel.",
			},
			"encapip": schema.StringAttribute{
				Computed:    true,
				Description: "The effective local IP address of the tunnel. Used as the source of the encapsulated packets.",
			},
			"channel": schema.Int64Attribute{
				Computed:    true,
				Description: "The tunnel that is bound to a netbridge.",
			},
			"tunneltype": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a tunnel is User-Configured, Internal or DELETE-IN-PROGRESS.",
			},
			"ipsectunnelstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the ipsec on this tunnel is up or down.",
			},
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of PBRs to bound to this iptunnel.",
			},
		},
	}
}

// iptunnelDataSourceSetAttrFromGet projects a NITRO iptunnel GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func iptunnelDataSourceSetAttrFromGet(ctx context.Context, data *IptunnelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In iptunnelDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	// iptunnel is read via an array filter matched on remote/remotesubnetmask; the
	// returned rows do not echo "name", so derive the id from the config-supplied
	// name (already populated) instead of from the GET row.
	data.Id = types.StringValue(data.Name.ValueString())

	// Read/write attributes as read-back outputs.
	data.Destport = utils.MapGetInt64(g, "destport")
	data.Grepayload = utils.MapGetString(g, "grepayload")
	data.Ipsecprofilename = utils.MapGetString(g, "ipsecprofilename")
	data.Local = utils.MapGetString(g, "local")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Remote = utils.MapGetString(g, "remote")
	data.Remotesubnetmask = utils.MapGetString(g, "remotesubnetmask")
	data.Tosinherit = utils.MapGetString(g, "tosinherit")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vlantagging = utils.MapGetString(g, "vlantagging")
	data.Vnid = utils.MapGetInt64(g, "vnid")

	// Read-only metadata.
	data.Sysname = utils.MapGetString(g, "sysname")
	data.Type = utils.MapGetInt64(g, "type")
	data.Encapip = utils.MapGetString(g, "encapip")
	data.Channel = utils.MapGetInt64(g, "channel")
	data.Tunneltype = utils.MapGetStringList(g, "tunneltype")
	data.Ipsectunnelstatus = utils.MapGetString(g, "ipsectunnelstatus")
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
}
