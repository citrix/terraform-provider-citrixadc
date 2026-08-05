package hanode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// HanodeResourceModel describes the resource data model.
type HanodeResourceModel struct {
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
	Hanodeid             types.Int64  `tfsdk:"hanode_id"`
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
}

func (r *HanodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hanode resource.",
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
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.",
			},
			"inc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.",
			},
			"ipaddress": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
		},
	}
}

// hanodeGetThePayloadFromthePlan builds the full add/create payload, mirroring the
// SDK v2 createHanodeFunc contract.
func hanodeGetThePayloadFromthePlan(ctx context.Context, data *HanodeResourceModel) ha.Hanode {
	tflog.Debug(ctx, "In hanodeGetThePayloadFromthePlan Function")

	hanode := ha.Hanode{}
	// id is *int and zero (self node) is a valid value; hanode_id is Required so always known.
	if !data.Hanodeid.IsNull() && !data.Hanodeid.IsUnknown() {
		hanode.Id = utils.IntPtr(int(data.Hanodeid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		hanode.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Inc.IsNull() && !data.Inc.IsUnknown() {
		hanode.Inc = data.Inc.ValueString()
	}
	if !data.Rpcnodepassword.IsNull() && !data.Rpcnodepassword.IsUnknown() {
		hanode.Rpcnodepassword = data.Rpcnodepassword.ValueString()
	}
	if !data.Hastatus.IsNull() && !data.Hastatus.IsUnknown() {
		hanode.Hastatus = data.Hastatus.ValueString()
	}
	if !data.Hasync.IsNull() && !data.Hasync.IsUnknown() {
		hanode.Hasync = data.Hasync.ValueString()
	}
	if !data.Haprop.IsNull() && !data.Haprop.IsUnknown() {
		hanode.Haprop = data.Haprop.ValueString()
	}
	if !data.Failsafe.IsNull() && !data.Failsafe.IsUnknown() {
		hanode.Failsafe = data.Failsafe.ValueString()
	}
	if !data.Syncstatusstrictmode.IsNull() && !data.Syncstatusstrictmode.IsUnknown() {
		hanode.Syncstatusstrictmode = data.Syncstatusstrictmode.ValueString()
	}
	if !data.Deadinterval.IsNull() && !data.Deadinterval.IsUnknown() {
		hanode.Deadinterval = utils.IntPtr(int(data.Deadinterval.ValueInt64()))
	}
	if !data.Hellointerval.IsNull() && !data.Hellointerval.IsUnknown() {
		hanode.Hellointerval = utils.IntPtr(int(data.Hellointerval.ValueInt64()))
	}
	if !data.Maxflips.IsNull() && !data.Maxflips.IsUnknown() {
		hanode.Maxflips = utils.IntPtr(int(data.Maxflips.ValueInt64()))
	}
	if !data.Maxfliptime.IsNull() && !data.Maxfliptime.IsUnknown() {
		hanode.Maxfliptime = utils.IntPtr(int(data.Maxfliptime.ValueInt64()))
	}
	if !data.Syncvlan.IsNull() && !data.Syncvlan.IsUnknown() {
		hanode.Syncvlan = utils.IntPtr(int(data.Syncvlan.ValueInt64()))
	}

	return hanode
}

// hanodeGetTheUpdatablePayloadFromThePlan builds the PUT (update) payload, restricted
// to NITRO-updatable fields. ipaddress and rpcnodepassword are ForceNew/RequiresReplace
// and are therefore never part of an in-place update.
func hanodeGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *HanodeResourceModel) ha.Hanode {
	tflog.Debug(ctx, "In hanodeGetTheUpdatablePayloadFromThePlan Function")

	hanode := ha.Hanode{}
	// id must accompany every PUT /config/hanode payload.
	if !data.Hanodeid.IsNull() && !data.Hanodeid.IsUnknown() {
		hanode.Id = utils.IntPtr(int(data.Hanodeid.ValueInt64()))
	}
	if !data.Deadinterval.IsNull() && !data.Deadinterval.IsUnknown() {
		hanode.Deadinterval = utils.IntPtr(int(data.Deadinterval.ValueInt64()))
	}
	if !data.Failsafe.IsNull() && !data.Failsafe.IsUnknown() {
		hanode.Failsafe = data.Failsafe.ValueString()
	}
	if !data.Haprop.IsNull() && !data.Haprop.IsUnknown() {
		hanode.Haprop = data.Haprop.ValueString()
	}
	if !data.Hastatus.IsNull() && !data.Hastatus.IsUnknown() {
		hanode.Hastatus = data.Hastatus.ValueString()
	}
	if !data.Hasync.IsNull() && !data.Hasync.IsUnknown() {
		hanode.Hasync = data.Hasync.ValueString()
	}
	if !data.Hellointerval.IsNull() && !data.Hellointerval.IsUnknown() {
		hanode.Hellointerval = utils.IntPtr(int(data.Hellointerval.ValueInt64()))
	}
	if !data.Inc.IsNull() && !data.Inc.IsUnknown() {
		hanode.Inc = data.Inc.ValueString()
	}
	if !data.Maxflips.IsNull() && !data.Maxflips.IsUnknown() {
		hanode.Maxflips = utils.IntPtr(int(data.Maxflips.ValueInt64()))
	}
	if !data.Maxfliptime.IsNull() && !data.Maxfliptime.IsUnknown() {
		hanode.Maxfliptime = utils.IntPtr(int(data.Maxfliptime.ValueInt64()))
	}
	if !data.Syncstatusstrictmode.IsNull() && !data.Syncstatusstrictmode.IsUnknown() {
		hanode.Syncstatusstrictmode = data.Syncstatusstrictmode.ValueString()
	}
	if !data.Syncvlan.IsNull() && !data.Syncvlan.IsUnknown() {
		hanode.Syncvlan = utils.IntPtr(int(data.Syncvlan.ValueInt64()))
	}

	return hanode
}

// hanodeSetAttrFromGet maps a NITRO GET response onto the resource state. It preserves
// user-configured values for attributes that NITRO does not (reliably) echo back
// (ipaddress, rpcnodepassword) and guards Optional+Computed attributes so a configured
// value is not clobbered when NITRO omits the attribute from the response.
func hanodeSetAttrFromGet(ctx context.Context, data *HanodeResourceModel, getResponseData map[string]interface{}) *HanodeResourceModel {
	tflog.Debug(ctx, "In hanodeSetAttrFromGet Function")

	// hanode_id (NITRO "id")
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hanodeid = types.Int64Value(intVal)
		}
	}

	// ---- Optional+Computed numeric attributes ----
	if val, ok := getResponseData["deadinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deadinterval = types.Int64Value(intVal)
		}
	} else if data.Deadinterval.IsUnknown() {
		data.Deadinterval = types.Int64Null()
	}
	if val, ok := getResponseData["hellointerval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hellointerval = types.Int64Value(intVal)
		}
	} else if data.Hellointerval.IsUnknown() {
		data.Hellointerval = types.Int64Null()
	}
	if val, ok := getResponseData["maxflips"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxflips = types.Int64Value(intVal)
		}
	} else if data.Maxflips.IsUnknown() {
		data.Maxflips = types.Int64Null()
	}
	if val, ok := getResponseData["maxfliptime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxfliptime = types.Int64Value(intVal)
		}
	} else if data.Maxfliptime.IsUnknown() {
		data.Maxfliptime = types.Int64Null()
	}
	if val, ok := getResponseData["syncvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Syncvlan = types.Int64Value(intVal)
		}
	} else if data.Syncvlan.IsUnknown() {
		data.Syncvlan = types.Int64Null()
	}

	// ---- Optional+Computed string attributes ----
	if val, ok := getResponseData["failsafe"]; ok && val != nil {
		data.Failsafe = types.StringValue(val.(string))
	} else if data.Failsafe.IsUnknown() {
		data.Failsafe = types.StringNull()
	}
	if val, ok := getResponseData["haprop"]; ok && val != nil {
		data.Haprop = types.StringValue(val.(string))
	} else if data.Haprop.IsUnknown() {
		data.Haprop = types.StringNull()
	}
	if val, ok := getResponseData["hasync"]; ok && val != nil {
		data.Hasync = types.StringValue(val.(string))
	} else if data.Hasync.IsUnknown() {
		data.Hasync = types.StringNull()
	}
	if val, ok := getResponseData["inc"]; ok && val != nil {
		data.Inc = types.StringValue(val.(string))
	} else if data.Inc.IsUnknown() {
		data.Inc = types.StringNull()
	}
	if val, ok := getResponseData["syncstatusstrictmode"]; ok && val != nil {
		data.Syncstatusstrictmode = types.StringValue(val.(string))
	} else if data.Syncstatusstrictmode.IsUnknown() {
		data.Syncstatusstrictmode = types.StringNull()
	}

	// hastatus: NITRO returns "UP" where the user configured "ENABLED"; normalize for
	// backward compatibility with the SDK v2 resource.
	if val, ok := getResponseData["hastatus"]; ok && val != nil {
		s := val.(string)
		if s == "UP" {
			s = "ENABLED"
		}
		data.Hastatus = types.StringValue(s)
	} else if data.Hastatus.IsUnknown() {
		data.Hastatus = types.StringNull()
	}

	// ipaddress: Optional (not Computed) and, per the SDK v2 readHanodeFunc, never read
	// back from NITRO ("neither propagated nor synchronized"). Preserve the configured
	// value exactly so state always equals the plan for this non-Computed attribute
	// (adopting a GET value here would cause "inconsistent result after apply" for a
	// self node whose config omits ipaddress).

	// rpcnodepassword: secret; NITRO does not reliably return it. Preserve the configured
	// value and only resolve an unknown (unconfigured) value to null.
	if data.Rpcnodepassword.IsUnknown() {
		data.Rpcnodepassword = types.StringNull()
	}

	// ---- Read only attributes (strings per the NITRO Hanode struct) ----
	if val, ok := getResponseData["completedfliptime"]; ok && val != nil {
		data.Completedfliptime = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Completedfliptime = types.StringNull()
	}
	if val, ok := getResponseData["curflips"]; ok && val != nil {
		data.Curflips = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Curflips = types.StringNull()
	}
	if val, ok := getResponseData["enaifaces"]; ok && val != nil {
		data.Enaifaces = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Enaifaces = types.StringNull()
	}
	if val, ok := getResponseData["masterstatetime"]; ok && val != nil {
		data.Masterstatetime = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Masterstatetime = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["routemonitor"]; ok && val != nil {
		data.Routemonitor = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Routemonitor = types.StringNull()
	}
	if val, ok := getResponseData["routemonitorstate"]; ok && val != nil {
		data.Routemonitorstate = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Routemonitorstate = types.StringNull()
	}
	if val, ok := getResponseData["ssl2"]; ok && val != nil {
		data.Ssl2 = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Ssl2 = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource (single unique attribute: the hanode id), matching the
	// SDK v2 d.SetId(strconv.Itoa(hanode_id)) contract.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Hanodeid.ValueInt64()))

	return data
}

// hanodeSetAttrFromGetForDatasource maps a NITRO GET response onto the model for the
// datasource, copying every attribute straight from the response (the datasource has no
// prior user config to preserve) and setting the datasource ID.
func hanodeSetAttrFromGetForDatasource(ctx context.Context, data *HanodeResourceModel, getResponseData map[string]interface{}) *HanodeResourceModel {
	tflog.Debug(ctx, "In hanodeSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hanodeid = types.Int64Value(intVal)
		}
	} else {
		data.Hanodeid = types.Int64Null()
	}
	if val, ok := getResponseData["deadinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deadinterval = types.Int64Value(intVal)
		}
	} else {
		data.Deadinterval = types.Int64Null()
	}
	if val, ok := getResponseData["hellointerval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hellointerval = types.Int64Value(intVal)
		}
	} else {
		data.Hellointerval = types.Int64Null()
	}
	if val, ok := getResponseData["maxflips"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxflips = types.Int64Value(intVal)
		}
	} else {
		data.Maxflips = types.Int64Null()
	}
	if val, ok := getResponseData["maxfliptime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxfliptime = types.Int64Value(intVal)
		}
	} else {
		data.Maxfliptime = types.Int64Null()
	}
	if val, ok := getResponseData["syncvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Syncvlan = types.Int64Value(intVal)
		}
	} else {
		data.Syncvlan = types.Int64Null()
	}
	if val, ok := getResponseData["failsafe"]; ok && val != nil {
		data.Failsafe = types.StringValue(val.(string))
	} else {
		data.Failsafe = types.StringNull()
	}
	if val, ok := getResponseData["haprop"]; ok && val != nil {
		data.Haprop = types.StringValue(val.(string))
	} else {
		data.Haprop = types.StringNull()
	}
	if val, ok := getResponseData["hasync"]; ok && val != nil {
		data.Hasync = types.StringValue(val.(string))
	} else {
		data.Hasync = types.StringNull()
	}
	if val, ok := getResponseData["inc"]; ok && val != nil {
		data.Inc = types.StringValue(val.(string))
	} else {
		data.Inc = types.StringNull()
	}
	if val, ok := getResponseData["syncstatusstrictmode"]; ok && val != nil {
		data.Syncstatusstrictmode = types.StringValue(val.(string))
	} else {
		data.Syncstatusstrictmode = types.StringNull()
	}
	if val, ok := getResponseData["hastatus"]; ok && val != nil {
		s := val.(string)
		if s == "UP" {
			s = "ENABLED"
		}
		data.Hastatus = types.StringValue(s)
	} else {
		data.Hastatus = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	// rpcnodepassword is a secret and is never returned by NITRO.
	data.Rpcnodepassword = types.StringNull()

	if val, ok := getResponseData["completedfliptime"]; ok && val != nil {
		data.Completedfliptime = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Completedfliptime = types.StringNull()
	}
	if val, ok := getResponseData["curflips"]; ok && val != nil {
		data.Curflips = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Curflips = types.StringNull()
	}
	if val, ok := getResponseData["enaifaces"]; ok && val != nil {
		data.Enaifaces = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Enaifaces = types.StringNull()
	}
	if val, ok := getResponseData["masterstatetime"]; ok && val != nil {
		data.Masterstatetime = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Masterstatetime = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["routemonitor"]; ok && val != nil {
		data.Routemonitor = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Routemonitor = types.StringNull()
	}
	if val, ok := getResponseData["routemonitorstate"]; ok && val != nil {
		data.Routemonitorstate = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Routemonitorstate = types.StringNull()
	}
	if val, ok := getResponseData["ssl2"]; ok && val != nil {
		data.Ssl2 = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.Ssl2 = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(fmt.Sprintf("%v", val))
	} else {
		data.State = types.StringNull()
	}

	data.Id = types.StringValue(fmt.Sprintf("%d", data.Hanodeid.ValueInt64()))

	return data
}
