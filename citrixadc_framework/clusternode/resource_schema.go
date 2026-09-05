package clusternode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusternodeResourceModel describes the resource data model.
type ClusternodeResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Backplane            types.String `tfsdk:"backplane"`
	Clearnodegroupconfig types.String `tfsdk:"clearnodegroupconfig"`
	Delay                types.Int64  `tfsdk:"delay"`
	Force                types.Bool   `tfsdk:"force"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Nodegroup            types.String `tfsdk:"nodegroup"`
	Nodeid               types.Int64  `tfsdk:"nodeid"`
	Priority             types.Int64  `tfsdk:"priority"`
	State                types.String `tfsdk:"state"`
	Tunnelmode           types.String `tfsdk:"tunnelmode"`
}

func (r *ClusternodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternode resource.",
			},
			"backplane": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.",
			},
			"clearnodegroupconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to remove nodegroup config",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Applicable for Passive node and node becomes passive after this timeout (in minutes)",
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Node will be removed from cluster without prompting for user confirmation.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default node group in a Cluster system.",
			},
			"nodeid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(31),
				Description: "Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator.\nWhen the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator.\nNote: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
				Description: "Admin state of the cluster node. The available settings function as follows:\nACTIVE - The node serves traffic.\nSPARE - The node does not serve traffic unless an ACTIVE node goes down.\nPASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.",
			},
			"tunnelmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "To set the tunnel mode",
			},
		},
	}
}

// clusternodeGetThePayloadFromthePlan builds the full create (add) payload.
// force and clearnodegroupconfig are delete-time arguments (not part of the add/update
// body per the NITRO doc and the SDK v2 resource), so they are intentionally excluded.
func clusternodeGetThePayloadFromthePlan(ctx context.Context, data *ClusternodeResourceModel) cluster.Clusternode {
	tflog.Debug(ctx, "In clusternodeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	clusternode := cluster.Clusternode{}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		clusternode.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		clusternode.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		clusternode.State = data.State.ValueString()
	}
	if !data.Backplane.IsNull() && !data.Backplane.IsUnknown() {
		clusternode.Backplane = data.Backplane.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		clusternode.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Nodegroup.IsNull() && !data.Nodegroup.IsUnknown() {
		clusternode.Nodegroup = data.Nodegroup.ValueString()
	}
	if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
		clusternode.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Tunnelmode.IsNull() && !data.Tunnelmode.IsUnknown() {
		clusternode.Tunnelmode = data.Tunnelmode.ValueString()
	}

	return clusternode
}

// clusternodeGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload restricted
// to NITRO-updatable fields. It excludes the create-only params (ipaddress and nodegroup),
// which are present in the NITRO add payload but ABSENT from the update payload and must not
// be sent in the PUT. force/clearnodegroupconfig are delete-time args and are also excluded.
// This mirrors the reference gslbservicegroup update builder (name key + only updatable attrs).
func clusternodeGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *ClusternodeResourceModel) cluster.Clusternode {
	tflog.Debug(ctx, "In clusternodeGetTheUpdatablePayloadFromThePlan Function")

	clusternode := cluster.Clusternode{}
	// nodeid is the key and is mandatory in the update body.
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		clusternode.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		clusternode.State = data.State.ValueString()
	}
	if !data.Backplane.IsNull() && !data.Backplane.IsUnknown() {
		clusternode.Backplane = data.Backplane.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		clusternode.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
		clusternode.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Tunnelmode.IsNull() && !data.Tunnelmode.IsUnknown() {
		clusternode.Tunnelmode = data.Tunnelmode.ValueString()
	}

	return clusternode
}

func clusternodeSetAttrFromGet(ctx context.Context, data *ClusternodeResourceModel, getResponseData map[string]interface{}) *ClusternodeResourceModel {
	tflog.Debug(ctx, "In clusternodeSetAttrFromGet Function")

	// Convert API response to model.
	// Optional+Computed attributes use a guarded else-branch: a configured (known) value
	// is preserved when NITRO omits the attribute from GET (e.g. delete-time args force and
	// clearnodegroupconfig are never returned; delay=0 may be omitted). Nulling is only done
	// when the value is still unknown (Computed and not configured), which prevents the
	// "inconsistent result after apply" error.
	if val, ok := getResponseData["backplane"]; ok && val != nil {
		data.Backplane = types.StringValue(val.(string))
	} else if data.Backplane.IsUnknown() {
		data.Backplane = types.StringNull()
	}
	if val, ok := getResponseData["clearnodegroupconfig"]; ok && val != nil {
		data.Clearnodegroupconfig = types.StringValue(val.(string))
	} else if data.Clearnodegroupconfig.IsUnknown() {
		data.Clearnodegroupconfig = types.StringNull()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else if data.Delay.IsUnknown() {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["force"]; ok && val != nil {
		data.Force = types.BoolValue(val.(bool))
	} else if data.Force.IsUnknown() {
		data.Force = types.BoolNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["nodegroup"]; ok && val != nil {
		data.Nodegroup = types.StringValue(val.(string))
	} else if data.Nodegroup.IsUnknown() {
		data.Nodegroup = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tunnelmode"]; ok && val != nil {
		data.Tunnelmode = types.StringValue(val.(string))
	} else if data.Tunnelmode.IsUnknown() {
		data.Tunnelmode = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute (nodeid) - use plain value as ID (matches SDK v2)
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Nodeid.ValueInt64()))

	return data
}
