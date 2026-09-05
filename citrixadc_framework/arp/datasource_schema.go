package arp

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ArpDataSourceModel is the data-source-specific model, decoupled from
// ArpResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only ARP metadata attributes that the resource
// deliberately omits (timeout, state, flags, type, channel, controlplane). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type ArpDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	All       types.Bool   `tfsdk:"all"`
	Ifnum     types.String `tfsdk:"ifnum"`
	Ipaddress types.String `tfsdk:"ipaddress"` // Required lookup key
	Mac       types.String `tfsdk:"mac"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
	Td        types.Int64  `tfsdk:"td"`
	Vlan      types.Int64  `tfsdk:"vlan"`
	Vtep      types.String `tfsdk:"vtep"`
	Vxlan     types.Int64  `tfsdk:"vxlan"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/arp.json). Never settable; populated from GET.
	Timeout      types.Int64  `tfsdk:"timeout"`
	State        types.Int64  `tfsdk:"state"`
	Flags        types.Int64  `tfsdk:"flags"`
	Type         types.String `tfsdk:"type"`
	Channel      types.Int64  `tfsdk:"channel"`
	Controlplane types.Bool   `tfsdk:"controlplane"`
}

func ArpDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all ARP entries from the ARP table of the Citrix ADC.",
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the network device is accessible. Specify the interface in (slot/port) notation. For example, 1/3.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the network device that you want to add to the ARP table.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address of the network device.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "The owner node for the Arp entry.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VLAN ID through which packets are to be sent after matching the ARP entry. This is a numeric value.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IP address of this ARP entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN on which the IP address of this ARP entry is reachable.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"timeout": schema.Int64Attribute{
				Computed:    true,
				Description: "The time, in seconds, after which the entry times out.",
			},
			"state": schema.Int64Attribute{
				Computed:    true,
				Description: "The state of the ARP entry.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "The flags for the entry.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates whether this ARP entry was added manually or dynamically (STATIC, PERMANENT, or DYNAMIC).",
			},
			"channel": schema.Int64Attribute{
				Computed:    true,
				Description: "The tunnel, channel, or physical interface through which the ARP entry is identified.",
			},
			"controlplane": schema.BoolAttribute{
				Computed:    true,
				Description: "This arp entry is populated by a control plane protocol.",
			},
		},
	}
}

// arpDataSourceSetAttrFromGet projects a NITRO arp GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func arpDataSourceSetAttrFromGet(ctx context.Context, data *ArpDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In arpDataSourceSetAttrFromGet Function")

	if v, ok := g["ipaddress"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Ipaddress = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ifnum = utils.MapGetString(g, "ifnum")
	data.Mac = utils.MapGetString(g, "mac")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ownernode = utils.MapGetInt64(g, "ownernode")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vtep = utils.MapGetString(g, "vtep")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")

	// all is an action-only input the GET never returns -> Null.
	data.All = types.BoolNull()

	// Read-only metadata.
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.State = utils.MapGetInt64(g, "state")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Type = utils.MapGetString(g, "type")
	data.Channel = utils.MapGetInt64(g, "channel")
	data.Controlplane = utils.MapGetBool(g, "controlplane")
}
