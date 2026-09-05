package vlan

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VlanDataSourceModel is the data-source-specific model, decoupled from
// VlanResourceModel. A data source is a pure read surface, so it exposes the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only VLAN membership/metadata the resource deliberately omits
// (linklocalipv6addr, the bitmaps, member interfaces, etc.). Every non-key
// attribute is Computed.
type VlanDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Aliasname          types.String `tfsdk:"aliasname"`
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Vlanid             types.Int64  `tfsdk:"vlanid"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`
	Mtu                types.Int64  `tfsdk:"mtu"`
	Sharing            types.String `tfsdk:"sharing"`

	// Read-only (GET-only) attributes from zion73x_readonly. Never settable.
	Linklocalipv6addr types.String `tfsdk:"linklocalipv6addr"`
	Rnat              types.Bool   `tfsdk:"rnat"`
	Portbitmap        types.Int64  `tfsdk:"portbitmap"`
	Lsbitmap          types.Int64  `tfsdk:"lsbitmap"`
	Tagbitmap         types.Int64  `tfsdk:"tagbitmap"`
	Lstagbitmap       types.Int64  `tfsdk:"lstagbitmap"`
	Ifaces            types.String `tfsdk:"ifaces"`
	Tagifaces         types.String `tfsdk:"tagifaces"`
	Ifnum             types.String `tfsdk:"ifnum"`
	Tagged            types.Bool   `tfsdk:"tagged"`
	Vlantd            types.Int64  `tfsdk:"vlantd"`
	Sdxvlan           types.String `tfsdk:"sdxvlan"`
	Partitionname     types.String `tfsdk:"partitionname"`
	Vxlan             types.Int64  `tfsdk:"vxlan"`
}

func VlanDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A name for the VLAN. Must begin with a letter, a number, or the underscore symbol, and can consist of from 1 to 31 letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the VLAN. However, you cannot perform any VLAN operation by specifying this name instead of the VLAN ID.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing on this VLAN.",
			},
			"vlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer that uniquely identifies a VLAN.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on this VLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the maximum transmission unit (MTU), in bytes. The MTU is the largest packet size, excluding 14 bytes of ethernet header and 4 bytes of crc, that can be transmitted and received over this VLAN.",
			},
			"sharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If sharing is enabled, then this vlan can be shared across multiple partitions by binding it to all those partitions. If sharing is disabled, then this vlan can be bound to only one of the partitions.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"linklocalipv6addr": schema.StringAttribute{
				Computed:    true,
				Description: "The link-local IP address assigned to the VLAN.",
			},
			"rnat": schema.BoolAttribute{
				Computed:    true,
				Description: "Temporary flag used for internal purpose.",
			},
			"portbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Member interfaces of this vlan.",
			},
			"lsbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Member linksets of this vlan.",
			},
			"tagbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Tagged members of this vlan.",
			},
			"lstagbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Tagged linksets of this vlan.",
			},
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Names of all member interfaces of this vlan.",
			},
			"tagifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Names of all tagged member interfaces of this vlan.",
			},
			"ifnum": schema.StringAttribute{
				Computed:    true,
				Description: "The interface bound to the VLAN, specified in slot/port notation (for example, 1/3).",
			},
			"tagged": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the interface is an 802.1q tagged interface.",
			},
			"vlantd": schema.Int64Attribute{
				Computed:    true,
				Description: "Traffic domain associated with vlan.",
			},
			"sdxvlan": schema.StringAttribute{
				Computed:    true,
				Description: "SDX vlan. Possible values = YES, NO.",
			},
			"partitionname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the Partition to which this vlan is bound.",
			},
			"vxlan": schema.Int64Attribute{
				Computed:    true,
				Description: "The VXLAN that extends this vlan.",
			},
		},
	}
}

// vlanDataSourceSetAttrFromGet projects a NITRO vlan GET response onto the
// data-source model. The VLAN id is carried by the NITRO "id" key; vlanid mirrors
// it and id (tfsdk) is its string form. All other attributes are filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func vlanDataSourceSetAttrFromGet(ctx context.Context, data *VlanDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vlanDataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}
	data.Vlanid = utils.MapGetInt64(g, "id")

	data.Aliasname = utils.MapGetString(g, "aliasname")
	data.Dynamicrouting = utils.MapGetString(g, "dynamicrouting")
	data.Ipv6dynamicrouting = utils.MapGetString(g, "ipv6dynamicrouting")
	data.Mtu = utils.MapGetInt64(g, "mtu")
	data.Sharing = utils.MapGetString(g, "sharing")

	// Read-only attributes.
	data.Linklocalipv6addr = utils.MapGetString(g, "linklocalipv6addr")
	data.Rnat = utils.MapGetBool(g, "rnat")
	data.Portbitmap = utils.MapGetInt64(g, "portbitmap")
	data.Lsbitmap = utils.MapGetInt64(g, "lsbitmap")
	data.Tagbitmap = utils.MapGetInt64(g, "tagbitmap")
	data.Lstagbitmap = utils.MapGetInt64(g, "lstagbitmap")
	data.Ifaces = utils.MapGetString(g, "ifaces")
	data.Tagifaces = utils.MapGetString(g, "tagifaces")
	data.Ifnum = utils.MapGetString(g, "ifnum")
	data.Tagged = utils.MapGetBool(g, "tagged")
	data.Vlantd = utils.MapGetInt64(g, "vlantd")
	data.Sdxvlan = utils.MapGetString(g, "sdxvlan")
	data.Partitionname = utils.MapGetString(g, "partitionname")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")
}
