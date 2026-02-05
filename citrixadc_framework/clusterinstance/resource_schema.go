package clusterinstance

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ClusterinstanceResourceModel describes the resource data model.
type ClusterinstanceResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Backplanebasedview         types.String `tfsdk:"backplanebasedview"`
	Clid                       types.Int64  `tfsdk:"clid"`
	Clusterproxyarp            types.String `tfsdk:"clusterproxyarp"`
	Deadinterval               types.Int64  `tfsdk:"deadinterval"`
	Dfdretainl2params          types.String `tfsdk:"dfdretainl2params"`
	Hellointerval              types.Int64  `tfsdk:"hellointerval"`
	Inc                        types.String `tfsdk:"inc"`
	Nodegroup                  types.String `tfsdk:"nodegroup"`
	Preemption                 types.String `tfsdk:"preemption"`
	Processlocal               types.String `tfsdk:"processlocal"`
	Quorumtype                 types.String `tfsdk:"quorumtype"`
	Retainconnectionsoncluster types.String `tfsdk:"retainconnectionsoncluster"`
	Secureheartbeats           types.String `tfsdk:"secureheartbeats"`
	Syncstatusstrictmode       types.String `tfsdk:"syncstatusstrictmode"`
}

func (r *ClusterinstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusterinstance resource.",
			},
			"backplanebasedview": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "View based on heartbeat only on bkplane interface",
			},
			"clid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster.",
			},
			"clusterproxyarp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "This field controls the proxy arp feature in cluster. By default the flag is enabled.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec",
			},
			"dfdretainl2params": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "flag to add ext l2 header during steering. By default the flag is disabled.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec",
			},
			"inc": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option is required if the cluster nodes reside on different networks.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The node group in a Cluster system used for transition from L2 to L3.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering.",
			},
			"quorumtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("MAJORITY"),
				Description: "Quorum Configuration Choices  - \"Majority\" (recommended) requires majority of nodes to be online for the cluster to be UP. \"None\" relaxes this requirement.",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.",
			},
			"secureheartbeats": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By turning on this option cluster heartbeats will have security enabled.",
			},
			"syncstatusstrictmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled.",
			},
		},
	}
}

func clusterinstanceGetThePayloadFromtheConfig(ctx context.Context, data *ClusterinstanceResourceModel) cluster.Clusterinstance {
	tflog.Debug(ctx, "In clusterinstanceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	clusterinstance := cluster.Clusterinstance{}
	if !data.Backplanebasedview.IsNull() {
		clusterinstance.Backplanebasedview = data.Backplanebasedview.ValueString()
	}
	if !data.Clid.IsNull() {
		clusterinstance.Clid = utils.IntPtr(int(data.Clid.ValueInt64()))
	}
	if !data.Clusterproxyarp.IsNull() {
		clusterinstance.Clusterproxyarp = data.Clusterproxyarp.ValueString()
	}
	if !data.Deadinterval.IsNull() {
		clusterinstance.Deadinterval = utils.IntPtr(int(data.Deadinterval.ValueInt64()))
	}
	if !data.Dfdretainl2params.IsNull() {
		clusterinstance.Dfdretainl2params = data.Dfdretainl2params.ValueString()
	}
	if !data.Hellointerval.IsNull() {
		clusterinstance.Hellointerval = utils.IntPtr(int(data.Hellointerval.ValueInt64()))
	}
	if !data.Inc.IsNull() {
		clusterinstance.Inc = data.Inc.ValueString()
	}
	if !data.Nodegroup.IsNull() {
		clusterinstance.Nodegroup = data.Nodegroup.ValueString()
	}
	if !data.Preemption.IsNull() {
		clusterinstance.Preemption = data.Preemption.ValueString()
	}
	if !data.Processlocal.IsNull() {
		clusterinstance.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Quorumtype.IsNull() {
		clusterinstance.Quorumtype = data.Quorumtype.ValueString()
	}
	if !data.Retainconnectionsoncluster.IsNull() {
		clusterinstance.Retainconnectionsoncluster = data.Retainconnectionsoncluster.ValueString()
	}
	if !data.Secureheartbeats.IsNull() {
		clusterinstance.Secureheartbeats = data.Secureheartbeats.ValueString()
	}
	if !data.Syncstatusstrictmode.IsNull() {
		clusterinstance.Syncstatusstrictmode = data.Syncstatusstrictmode.ValueString()
	}

	return clusterinstance
}

func clusterinstanceSetAttrFromGet(ctx context.Context, data *ClusterinstanceResourceModel, getResponseData map[string]interface{}) *ClusterinstanceResourceModel {
	tflog.Debug(ctx, "In clusterinstanceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["backplanebasedview"]; ok && val != nil {
		data.Backplanebasedview = types.StringValue(val.(string))
	} else {
		data.Backplanebasedview = types.StringNull()
	}
	if val, ok := getResponseData["clid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clid = types.Int64Value(intVal)
		}
	} else {
		data.Clid = types.Int64Null()
	}
	if val, ok := getResponseData["clusterproxyarp"]; ok && val != nil {
		data.Clusterproxyarp = types.StringValue(val.(string))
	} else {
		data.Clusterproxyarp = types.StringNull()
	}
	if val, ok := getResponseData["deadinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deadinterval = types.Int64Value(intVal)
		}
	} else {
		data.Deadinterval = types.Int64Null()
	}
	if val, ok := getResponseData["dfdretainl2params"]; ok && val != nil {
		data.Dfdretainl2params = types.StringValue(val.(string))
	} else {
		data.Dfdretainl2params = types.StringNull()
	}
	if val, ok := getResponseData["hellointerval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hellointerval = types.Int64Value(intVal)
		}
	} else {
		data.Hellointerval = types.Int64Null()
	}
	if val, ok := getResponseData["inc"]; ok && val != nil {
		data.Inc = types.StringValue(val.(string))
	} else {
		data.Inc = types.StringNull()
	}
	if val, ok := getResponseData["nodegroup"]; ok && val != nil {
		data.Nodegroup = types.StringValue(val.(string))
	} else {
		data.Nodegroup = types.StringNull()
	}
	if val, ok := getResponseData["preemption"]; ok && val != nil {
		data.Preemption = types.StringValue(val.(string))
	} else {
		data.Preemption = types.StringNull()
	}
	if val, ok := getResponseData["processlocal"]; ok && val != nil {
		data.Processlocal = types.StringValue(val.(string))
	} else {
		data.Processlocal = types.StringNull()
	}
	if val, ok := getResponseData["quorumtype"]; ok && val != nil {
		data.Quorumtype = types.StringValue(val.(string))
	} else {
		data.Quorumtype = types.StringNull()
	}
	if val, ok := getResponseData["retainconnectionsoncluster"]; ok && val != nil {
		data.Retainconnectionsoncluster = types.StringValue(val.(string))
	} else {
		data.Retainconnectionsoncluster = types.StringNull()
	}
	if val, ok := getResponseData["secureheartbeats"]; ok && val != nil {
		data.Secureheartbeats = types.StringValue(val.(string))
	} else {
		data.Secureheartbeats = types.StringNull()
	}
	if val, ok := getResponseData["syncstatusstrictmode"]; ok && val != nil {
		data.Syncstatusstrictmode = types.StringValue(val.(string))
	} else {
		data.Syncstatusstrictmode = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Clid.ValueInt64()))

	return data
}
