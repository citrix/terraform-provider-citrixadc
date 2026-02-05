package clusternodegroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusternodegroupResourceModel describes the resource data model.
type ClusternodegroupResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Priority types.Int64  `tfsdk:"priority"`
	State    types.String `tfsdk:"state"`
	Sticky   types.String `tfsdk:"sticky"`
	Strict   types.String `tfsdk:"strict"`
}

func (r *ClusternodegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusternodegroup resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of Nodegroup. This priority is used for all the nodes bound to the nodegroup for Nodegroup selection",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the nodegroup. All the nodes binding to this nodegroup must have the same state. ACTIVE/SPARE/PASSIVE",
			},
			"sticky": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster.\n  * Enabled - When owner node goes down, backup node will become the owner node and takes the traffic for the entities bound to the nodegroup. When bound node rejoins the cluster, traffic for the entities bound to nodegroup will not be steered back to this bound node. Current owner will have the ownership till it goes down.\n  * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.",
			},
			"strict": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup.\n  * Enabled - When one of the nodes goes down, no other cluster node is picked up to replace it. When the node comes up, it will continue being part of the nodegroup.\n  * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.",
			},
		},
	}
}

func clusternodegroupGetThePayloadFromtheConfig(ctx context.Context, data *ClusternodegroupResourceModel) cluster.Clusternodegroup {
	tflog.Debug(ctx, "In clusternodegroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusternodegroup := cluster.Clusternodegroup{}
	if !data.Name.IsNull() {
		clusternodegroup.Name = data.Name.ValueString()
	}
	if !data.Priority.IsNull() {
		clusternodegroup.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.State.IsNull() {
		clusternodegroup.State = data.State.ValueString()
	}
	if !data.Sticky.IsNull() {
		clusternodegroup.Sticky = data.Sticky.ValueString()
	}
	if !data.Strict.IsNull() {
		clusternodegroup.Strict = data.Strict.ValueString()
	}

	return clusternodegroup
}

func clusternodegroupSetAttrFromGet(ctx context.Context, data *ClusternodegroupResourceModel, getResponseData map[string]interface{}) *ClusternodegroupResourceModel {
	tflog.Debug(ctx, "In clusternodegroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["sticky"]; ok && val != nil {
		data.Sticky = types.StringValue(val.(string))
	} else {
		data.Sticky = types.StringNull()
	}
	if val, ok := getResponseData["strict"]; ok && val != nil {
		data.Strict = types.StringValue(val.(string))
	} else {
		data.Strict = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
