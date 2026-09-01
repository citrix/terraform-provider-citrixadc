package hanode

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HanodeDataSourceModel is the data-source-specific model, decoupled from
// HanodeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only NITRO attributes the resource deliberately omits
// (flags, interface lists, hasyncfailurereason, ...). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type HanodeDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Completedfliptime    types.String `tfsdk:"completedfliptime"`
	Curflips             types.String `tfsdk:"curflips"`
	Deadinterval         types.Int64  `tfsdk:"deadinterval"`
	Enaifaces            types.String `tfsdk:"enaifaces"`
	Failsafe             types.String `tfsdk:"failsafe"`
	Haprop               types.String `tfsdk:"haprop"`
	Hastatus             types.String `tfsdk:"hastatus"`
	Hasync               types.String `tfsdk:"hasync"`
	Hellointerval        types.Int64  `tfsdk:"hellointerval"`
	Hanodeid             types.Int64  `tfsdk:"hanode_id"` // Required lookup key
	Inc                  types.String `tfsdk:"inc"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Masterstatetime      types.String `tfsdk:"masterstatetime"`
	Maxflips             types.Int64  `tfsdk:"maxflips"`
	Maxfliptime          types.Int64  `tfsdk:"maxfliptime"`
	Netmask              types.String `tfsdk:"netmask"`
	Routemonitor         types.String `tfsdk:"routemonitor"`
	Routemonitorstate    types.String `tfsdk:"routemonitorstate"`
	Rpcnodepassword      types.String `tfsdk:"rpcnodepassword"`
	Ssl2                 types.String `tfsdk:"ssl2"`
	State                types.String `tfsdk:"state"`
	Syncstatusstrictmode types.String `tfsdk:"syncstatusstrictmode"`
	Syncvlan             types.Int64  `tfsdk:"syncvlan"`

	// Read-only (GET-only) NITRO attributes from the read-only set
	// (zion73x_readonly/hanode.json) not already surfaced above. Never settable;
	// populated from GET.
	Name                 types.String `tfsdk:"name"`
	Flags                types.Int64  `tfsdk:"flags"`
	Disifaces            types.String `tfsdk:"disifaces"`
	Hamonifaces          types.String `tfsdk:"hamonifaces"`
	Haheartbeatifaces    types.String `tfsdk:"haheartbeatifaces"`
	Pfifaces             types.String `tfsdk:"pfifaces"`
	Ifaces               types.String `tfsdk:"ifaces"`
	Hasyncfailurereason  types.String `tfsdk:"hasyncfailurereason"`
	Secureheartbeatstate types.String `tfsdk:"secureheartbeatstate"`
}

func HanodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"completedfliptime": schema.StringAttribute{
				Computed:    true,
				Description: "To inform user whether flip time is elapsed or not.",
			},
			"curflips": schema.StringAttribute{
				Computed:    true,
				Description: "Keeps track of number of flips that have happened till now in current interval.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.",
			},
			"enaifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Enabled interfaces.",
			},
			"failsafe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.",
			},
			"haprop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Automatically propagate all commands from the primary to the secondary node, except the following:\n* All HA configuration related commands. For example, add ha node, set ha node, and bind ha node.\n* All Interface related commands. For example, set interface and unset interface.\n* All channels related commands. For example, add channel, set channel, and bind channel.\nThe propagated command is executed on the secondary node before it is executed on the primary. If command propagation fails, or if command execution fails on the secondary, the primary node executes the command and logs an error.  Command propagation uses port 3010.\nNote: After enabling propagation, run force synchronization on either node.",
			},
			"hastatus": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node.\n	 STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.",
			},
			"hasync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.",
			},
			"hanode_id": schema.Int64Attribute{
				Required:    true,
				Description: "Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.",
			},
			"inc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.",
			},
			"masterstatetime": schema.StringAttribute{
				Computed:    true,
				Description: "Time elapsed in current master state.",
			},
			"maxflips": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Max number of flips allowed before becoming sticky primary",
			},
			"maxfliptime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval after which flipping of node states can again start",
			},
			"netmask": schema.StringAttribute{
				Computed:    true,
				Description: "The netmask.",
			},
			"routemonitor": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
			"routemonitorstate": schema.StringAttribute{
				Computed:    true,
				Description: "State for route monitor.",
			},
			"rpcnodepassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to be used in authentication with the peer rpc node.",
			},
			"ssl2": schema.StringAttribute{
				Computed:    true,
				Description: "SSL card status.",
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "HA master state.",
			},
			"syncstatusstrictmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "strict mode flag for sync status",
			},
			"syncvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.",
			},

			// Read-only (GET-only) NITRO attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Node Name.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "The flags for this entry.",
			},
			"disifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Disabled interfaces.",
			},
			"hamonifaces": schema.StringAttribute{
				Computed:    true,
				Description: "HAMON ON interfaces.",
			},
			"haheartbeatifaces": schema.StringAttribute{
				Computed:    true,
				Description: "HAHEARTBEAT OFF interfaces.",
			},
			"pfifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Interfaces causing Partial Failure.",
			},
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Interfaces on which non-multicast is not seen.",
			},
			"hasyncfailurereason": schema.StringAttribute{
				Computed:    true,
				Description: "Displays the reason for HA SYNC Failure.",
			},
			"secureheartbeatstate": schema.StringAttribute{
				Computed:    true,
				Description: "Displays the current state of HA secure heartbeats (ENABLED, DISABLED).",
			},
		},
	}
}

// hanodeDataSourceSetAttrFromGet projects a NITRO hanode GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) — no unknown->null resolution or plan preservation is required. The
// shared utils.MapGet* helpers implement that projection.
func hanodeDataSourceSetAttrFromGet(ctx context.Context, data *HanodeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In hanodeDataSourceSetAttrFromGet Function")

	// hanode_id is the NITRO "id" key; also drives the datasource ID.
	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		if iv, err := utils.ConvertToInt64(v); err == nil {
			data.Hanodeid = types.Int64Value(iv)
		}
	}

	// Read/write attributes as read-back outputs.
	data.Completedfliptime = utils.MapGetString(g, "completedfliptime")
	data.Curflips = utils.MapGetString(g, "curflips")
	data.Deadinterval = utils.MapGetInt64(g, "deadinterval")
	data.Enaifaces = utils.MapGetString(g, "enaifaces")
	data.Failsafe = utils.MapGetString(g, "failsafe")
	data.Haprop = utils.MapGetString(g, "haprop")
	// hastatus: NITRO returns "UP" where the user configured "ENABLED"; normalize
	// for consistency with the resource read-back.
	data.Hastatus = utils.MapGetString(g, "hastatus")
	if data.Hastatus.ValueString() == "UP" {
		data.Hastatus = types.StringValue("ENABLED")
	}
	data.Hasync = utils.MapGetString(g, "hasync")
	data.Hellointerval = utils.MapGetInt64(g, "hellointerval")
	data.Inc = utils.MapGetString(g, "inc")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Masterstatetime = utils.MapGetString(g, "masterstatetime")
	data.Maxflips = utils.MapGetInt64(g, "maxflips")
	data.Maxfliptime = utils.MapGetInt64(g, "maxfliptime")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Routemonitor = utils.MapGetString(g, "routemonitor")
	data.Routemonitorstate = utils.MapGetString(g, "routemonitorstate")
	data.Ssl2 = utils.MapGetString(g, "ssl2")
	data.State = utils.MapGetString(g, "state")
	data.Syncstatusstrictmode = utils.MapGetString(g, "syncstatusstrictmode")
	data.Syncvlan = utils.MapGetInt64(g, "syncvlan")

	// rpcnodepassword is a secret the GET never returns -> Null.
	data.Rpcnodepassword = types.StringNull()

	// Read-only NITRO attributes.
	data.Name = utils.MapGetString(g, "name")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Disifaces = utils.MapGetString(g, "disifaces")
	data.Hamonifaces = utils.MapGetString(g, "hamonifaces")
	data.Haheartbeatifaces = utils.MapGetString(g, "haheartbeatifaces")
	data.Pfifaces = utils.MapGetString(g, "pfifaces")
	data.Ifaces = utils.MapGetString(g, "ifaces")
	data.Hasyncfailurereason = utils.MapGetString(g, "hasyncfailurereason")
	data.Secureheartbeatstate = utils.MapGetString(g, "secureheartbeatstate")
}
