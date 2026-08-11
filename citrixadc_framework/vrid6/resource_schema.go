package vrid6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Vrid6ResourceModel describes the resource data model.
type Vrid6ResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	All                  types.Bool   `tfsdk:"all"`
	Vrid6_id             types.Int64  `tfsdk:"vrid6_id"`
	Ownernode            types.Int64  `tfsdk:"ownernode"`
	Preemption           types.String `tfsdk:"preemption"`
	Preemptiondelaytimer types.Int64  `tfsdk:"preemptiondelaytimer"`
	Priority             types.Int64  `tfsdk:"priority"`
	Sharing              types.String `tfsdk:"sharing"`
	Trackifnumpriority   types.Int64  `tfsdk:"trackifnumpriority"`
	Tracking             types.String `tfsdk:"tracking"`
}

func (r *Vrid6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid6 resource.",
			},
			// SDK v2: Optional + Computed, NOT ForceNew and updateable.
			// The auto-gen wrongly added RequiresReplace() here; removed for
			// backward-compat with the SDK v2 schema.
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all configured VMAC6 addresses from the Citrix ADC.",
			},
			// SDK v2: Required + ForceNew. This is the named-resource key.
			"vrid6_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, owner node is displayed as ALL and one node is dynamically elected as the owner.",
			},
			// SDK v2 had NO Default for preemption (Optional + Computed only).
			// A Default is invalid without Computed and would diverge from SDK v2,
			// so read the effective value from the ADC instead.
			// Default added (NITRO spec default ENABLED) so removing this attr
			// from config produces a plan diff that fires Update -> unset.
			"preemption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.\n             If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.",
			},
			// Default added (NITRO spec default 0) to enable unset-on-removal.
			"preemptiondelaytimer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.",
			},
			// Default added (NITRO spec default 255) to enable unset-on-removal.
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(255),
				Description: "Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.",
			},
			// Default added (NITRO spec default DISABLED) to enable unset-on-removal.
			"sharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.",
			},
			// Default added (NITRO spec default 0) to enable unset-on-removal.
			"trackifnumpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.",
			},
			// Default added (NITRO spec default NONE) to enable unset-on-removal.
			"tracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address.\nAvailable settings function as follows:\n* NONE - No tracking. EP = BP\n* ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0.\n* ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0.\n* PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN.\nDefault: NONE.",
			},
		},
	}
}

func vrid6GetThePayloadFromthePlan(ctx context.Context, data *Vrid6ResourceModel) network.Vrid6 {
	tflog.Debug(ctx, "In vrid6GetThePayloadFromthePlan Function")

	// Create API request body from the model
	vrid6 := network.Vrid6{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		vrid6.All = data.All.ValueBool()
	}
	if !data.Vrid6_id.IsNull() && !data.Vrid6_id.IsUnknown() {
		vrid6.Id = utils.IntPtr(int(data.Vrid6_id.ValueInt64()))
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		vrid6.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Preemption.IsNull() && !data.Preemption.IsUnknown() {
		vrid6.Preemption = data.Preemption.ValueString()
	}
	if !data.Preemptiondelaytimer.IsNull() && !data.Preemptiondelaytimer.IsUnknown() {
		vrid6.Preemptiondelaytimer = utils.IntPtr(int(data.Preemptiondelaytimer.ValueInt64()))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vrid6.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Sharing.IsNull() && !data.Sharing.IsUnknown() {
		vrid6.Sharing = data.Sharing.ValueString()
	}
	if !data.Trackifnumpriority.IsNull() && !data.Trackifnumpriority.IsUnknown() {
		vrid6.Trackifnumpriority = utils.IntPtr(int(data.Trackifnumpriority.ValueInt64()))
	}
	if !data.Tracking.IsNull() && !data.Tracking.IsUnknown() {
		vrid6.Tracking = data.Tracking.ValueString()
	}

	return vrid6
}

func vrid6SetAttrFromGet(ctx context.Context, data *Vrid6ResourceModel, getResponseData map[string]interface{}) *Vrid6ResourceModel {
	tflog.Debug(ctx, "In vrid6SetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches so a known (configured/prior-state) value is not
	// clobbered to null when NITRO omits an attribute from GET (omit-on-default
	// trap). Only null a value that is still Unknown (never configured).
	if val, ok := getResponseData["all"]; ok && val != nil {
		data.All = types.BoolValue(val.(bool))
	} else if data.All.IsUnknown() {
		data.All = types.BoolNull()
	}
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid6_id = types.Int64Value(intVal)
		}
	} else if data.Vrid6_id.IsUnknown() {
		data.Vrid6_id = types.Int64Null()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["preemption"]; ok && val != nil {
		data.Preemption = types.StringValue(val.(string))
	} else if data.Preemption.IsUnknown() {
		data.Preemption = types.StringNull()
	}
	if val, ok := getResponseData["preemptiondelaytimer"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Preemptiondelaytimer = types.Int64Value(intVal)
		}
	} else if data.Preemptiondelaytimer.IsUnknown() {
		data.Preemptiondelaytimer = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["sharing"]; ok && val != nil {
		data.Sharing = types.StringValue(val.(string))
	} else if data.Sharing.IsUnknown() {
		data.Sharing = types.StringNull()
	}
	if val, ok := getResponseData["trackifnumpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trackifnumpriority = types.Int64Value(intVal)
		}
	} else if data.Trackifnumpriority.IsUnknown() {
		data.Trackifnumpriority = types.Int64Null()
	}
	if val, ok := getResponseData["tracking"]; ok && val != nil {
		data.Tracking = types.StringValue(val.(string))
	} else if data.Tracking.IsUnknown() {
		data.Tracking = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vrid6_id.ValueInt64()))

	return data
}
