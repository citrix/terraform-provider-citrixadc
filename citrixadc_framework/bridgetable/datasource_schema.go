package bridgetable

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BridgetableDataSourceModel is the data-source-specific model, decoupled from
// BridgetableResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type BridgetableDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Bridgeage  types.Int64  `tfsdk:"bridgeage"`
	Devicevlan types.Int64  `tfsdk:"devicevlan"`
	Ifnum      types.String `tfsdk:"ifnum"`
	Mac        types.String `tfsdk:"mac"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Vlan       types.Int64  `tfsdk:"vlan"`
	Vni        types.Int64  `tfsdk:"vni"`
	Vtep       types.String `tfsdk:"vtep"`
	Vxlan      types.Int64  `tfsdk:"vxlan"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/bridgetable.json). Never settable; populated from GET.
	Flags        types.Int64  `tfsdk:"flags"`
	Type         types.String `tfsdk:"type"`
	Channel      types.Int64  `tfsdk:"channel"`
	Controlplane types.Bool   `tfsdk:"controlplane"`
}

func BridgetableDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgeage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.",
			},
			"devicevlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.",
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "INTERFACE  whose entries are to be removed.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The MAC address of the target.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN  whose entries are to be removed.",
			},
			"vni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint.  If omitted the value specified as vxlan will be used.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VXLAN to which this address is associated.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Display flags.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Whether static or dynamic. Possible values = STATIC, PERMANENT, DYNAMIC",
			},
			"channel": schema.Int64Attribute{
				Computed:    true,
				Description: "The Tunnel through which bridge entry is learned.",
			},
			"controlplane": schema.BoolAttribute{
				Computed:    true,
				Description: "This bridge table entry is populated by a control plane protocol.",
			},
		},
	}
}

// bridgetableDataSourceSetAttrFromGet projects a NITRO bridgetable GET response
// entry onto the data-source model. Attributes are filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func bridgetableDataSourceSetAttrFromGet(ctx context.Context, data *BridgetableDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In bridgetableDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Bridgeage = utils.MapGetInt64(g, "bridgeage")
	data.Devicevlan = utils.MapGetInt64(g, "devicevlan")
	data.Ifnum = utils.MapGetString(g, "ifnum")
	data.Mac = utils.MapGetString(g, "mac")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vni = utils.MapGetInt64(g, "vni")
	data.Vtep = utils.MapGetString(g, "vtep")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")

	// Read-only metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Type = utils.MapGetString(g, "type")
	data.Channel = utils.MapGetInt64(g, "channel")
	data.Controlplane = utils.MapGetBool(g, "controlplane")

	// Backward-compatible composite ID: "mac,vxlan,vtep" (matches SDK v2).
	data.Id = types.StringValue(fmt.Sprintf("%s,%d,%s", data.Mac.ValueString(), data.Vxlan.ValueInt64(), data.Vtep.ValueString()))
}
