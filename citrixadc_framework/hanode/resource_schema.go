package hanode

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// HanodeResourceModel describes the resource data model.
type HanodeResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Deadinterval         types.Int64  `tfsdk:"deadinterval"`
	Failsafe             types.String `tfsdk:"failsafe"`
	Haprop               types.String `tfsdk:"haprop"`
	Hastatus             types.String `tfsdk:"hastatus"`
	Hasync               types.String `tfsdk:"hasync"`
	Hellointerval        types.Int64  `tfsdk:"hellointerval"`
	Hanodeid             types.Int64  `tfsdk:"hanode_id"`
	Inc                  types.String `tfsdk:"inc"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Maxflips             types.Int64  `tfsdk:"maxflips"`
	Maxfliptime          types.Int64  `tfsdk:"maxfliptime"`
	Rpcnodepassword      types.String `tfsdk:"rpcnodepassword"`
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
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.",
			},
			"failsafe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.",
			},
			"haprop": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Automatically propagate all commands from the primary to the secondary node, except the following:\n* All HA configuration related commands. For example, add ha node, set ha node, and bind ha node.\n* All Interface related commands. For example, set interface and unset interface.\n* All channels related commands. For example, add channel, set channel, and bind channel.\nThe propagated command is executed on the secondary node before it is executed on the primary. If command propagation fails, or if command execution fails on the secondary, the primary node executes the command and logs an error.  Command propagation uses port 3010.\nNote: After enabling propagation, run force synchronization on either node.",
			},
			"hastatus": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node.\n	 STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.",
			},
			"hasync": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.",
			},
			"hanode_id": schema.Int64Attribute{
				Required:    true,
				Description: "Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.",
			},
			"inc": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.",
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
			"rpcnodepassword": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to be used in authentication with the peer rpc node.",
			},
			"syncstatusstrictmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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

func hanodeGetThePayloadFromtheConfig(ctx context.Context, data *HanodeResourceModel) ha.Hanode {
	tflog.Debug(ctx, "In hanodeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	hanode := ha.Hanode{}
	if !data.Deadinterval.IsNull() {
		hanode.Deadinterval = utils.IntPtr(int(data.Deadinterval.ValueInt64()))
	}
	if !data.Failsafe.IsNull() {
		hanode.Failsafe = data.Failsafe.ValueString()
	}
	if !data.Haprop.IsNull() {
		hanode.Haprop = data.Haprop.ValueString()
	}
	if !data.Hastatus.IsNull() {
		hanode.Hastatus = data.Hastatus.ValueString()
	}
	if !data.Hasync.IsNull() {
		hanode.Hasync = data.Hasync.ValueString()
	}
	if !data.Hellointerval.IsNull() {
		hanode.Hellointerval = utils.IntPtr(int(data.Hellointerval.ValueInt64()))
	}
	if !data.Hanodeid.IsNull() {
		hanode.Id = utils.IntPtr(int(data.Hanodeid.ValueInt64()))
	}
	if !data.Inc.IsNull() {
		hanode.Inc = data.Inc.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		hanode.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxflips.IsNull() {
		hanode.Maxflips = utils.IntPtr(int(data.Maxflips.ValueInt64()))
	}
	if !data.Maxfliptime.IsNull() {
		hanode.Maxfliptime = utils.IntPtr(int(data.Maxfliptime.ValueInt64()))
	}
	if !data.Rpcnodepassword.IsNull() {
		hanode.Rpcnodepassword = data.Rpcnodepassword.ValueString()
	}
	if !data.Syncstatusstrictmode.IsNull() {
		hanode.Syncstatusstrictmode = data.Syncstatusstrictmode.ValueString()
	}
	if !data.Syncvlan.IsNull() {
		hanode.Syncvlan = utils.IntPtr(int(data.Syncvlan.ValueInt64()))
	}

	return hanode
}

func hanodeSetAttrFromGet(ctx context.Context, data *HanodeResourceModel, getResponseData map[string]interface{}) *HanodeResourceModel {
	tflog.Debug(ctx, "In hanodeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["deadinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deadinterval = types.Int64Value(intVal)
		}
	} else {
		data.Deadinterval = types.Int64Null()
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
	if val, ok := getResponseData["hastatus"]; ok && val != nil {
		data.Hastatus = types.StringValue(val.(string))
	} else {
		data.Hastatus = types.StringNull()
	}
	if val, ok := getResponseData["hasync"]; ok && val != nil {
		data.Hasync = types.StringValue(val.(string))
	} else {
		data.Hasync = types.StringNull()
	}
	if val, ok := getResponseData["hellointerval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hellointerval = types.Int64Value(intVal)
		}
	} else {
		data.Hellointerval = types.Int64Null()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hanodeid = types.Int64Value(intVal)
		}
	} else {
		data.Hanodeid = types.Int64Null()
	}
	if val, ok := getResponseData["inc"]; ok && val != nil {
		data.Inc = types.StringValue(val.(string))
	} else {
		data.Inc = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
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
	if val, ok := getResponseData["rpcnodepassword"]; ok && val != nil {
		data.Rpcnodepassword = types.StringValue(val.(string))
	} else {
		data.Rpcnodepassword = types.StringNull()
	}
	if val, ok := getResponseData["syncstatusstrictmode"]; ok && val != nil {
		data.Syncstatusstrictmode = types.StringValue(val.(string))
	} else {
		data.Syncstatusstrictmode = types.StringNull()
	}
	if val, ok := getResponseData["syncvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Syncvlan = types.Int64Value(intVal)
		}
	} else {
		data.Syncvlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Id.ValueString())

	return data
}
