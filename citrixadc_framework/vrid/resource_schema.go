package vrid

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VridResourceModel describes the resource data model.
type VridResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	All                  types.Bool   `tfsdk:"all"`
	Vrid_id              types.Int64  `tfsdk:"vrid_id"`
	Ownernode            types.Int64  `tfsdk:"ownernode"`
	Preemption           types.String `tfsdk:"preemption"`
	Preemptiondelaytimer types.Int64  `tfsdk:"preemptiondelaytimer"`
	Priority             types.Int64  `tfsdk:"priority"`
	Sharing              types.String `tfsdk:"sharing"`
	Trackifnumpriority   types.Int64  `tfsdk:"trackifnumpriority"`
	Tracking             types.String `tfsdk:"tracking"`
}

func (r *VridResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid resource.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Remove all the configured VMAC addresses from the Citrix ADC.",
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, owner node is displayed as ALL and one node is dynamically elected as the owner.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.\nIf you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.",
			},
			"preemptiondelaytimer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preemption delay time, in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.",
			},
			"sharing": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.",
			},
			"trackifnumpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.",
			},
			"tracking": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address.\nAvailable settings function as follows:\n* NONE - No tracking. EP = BP\n* ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0.\n* ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0.\n* PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN.\nDefault: NONE.",
			},
		},
	}
}

func vridGetThePayloadFromtheConfig(ctx context.Context, data *VridResourceModel) network.Vrid {
	tflog.Debug(ctx, "In vridGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vrid := network.Vrid{}
	if !data.All.IsNull() {
		vrid.All = data.All.ValueBool()
	}
	if !data.Vrid_id.IsNull() {
		vrid.Id = utils.IntPtr(int(data.Vrid_id.ValueInt64()))
	}
	if !data.Ownernode.IsNull() {
		vrid.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Preemption.IsNull() {
		vrid.Preemption = data.Preemption.ValueString()
	}
	if !data.Preemptiondelaytimer.IsNull() {
		vrid.Preemptiondelaytimer = utils.IntPtr(int(data.Preemptiondelaytimer.ValueInt64()))
	}
	if !data.Priority.IsNull() {
		vrid.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Sharing.IsNull() {
		vrid.Sharing = data.Sharing.ValueString()
	}
	if !data.Trackifnumpriority.IsNull() {
		vrid.Trackifnumpriority = utils.IntPtr(int(data.Trackifnumpriority.ValueInt64()))
	}
	if !data.Tracking.IsNull() {
		vrid.Tracking = data.Tracking.ValueString()
	}

	return vrid
}

func vridSetAttrFromGet(ctx context.Context, data *VridResourceModel, getResponseData map[string]interface{}) *VridResourceModel {
	tflog.Debug(ctx, "In vridSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["all"]; ok && val != nil {
		data.All = types.BoolValue(val.(bool))
	} else {
		data.All = types.BoolNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid_id = types.Int64Value(intVal)
		}
	} else {
		data.Vrid_id = types.Int64Null()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["preemption"]; ok && val != nil {
		data.Preemption = types.StringValue(val.(string))
	} else {
		data.Preemption = types.StringNull()
	}
	if val, ok := getResponseData["preemptiondelaytimer"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Preemptiondelaytimer = types.Int64Value(intVal)
		}
	} else {
		data.Preemptiondelaytimer = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["sharing"]; ok && val != nil {
		data.Sharing = types.StringValue(val.(string))
	} else {
		data.Sharing = types.StringNull()
	}
	if val, ok := getResponseData["trackifnumpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trackifnumpriority = types.Int64Value(intVal)
		}
	} else {
		data.Trackifnumpriority = types.Int64Null()
	}
	if val, ok := getResponseData["tracking"]; ok && val != nil {
		data.Tracking = types.StringValue(val.(string))
	} else {
		data.Tracking = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vrid_id.ValueInt64()))

	return data
}
