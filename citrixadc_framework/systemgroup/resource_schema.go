package systemgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CmdpolicybindingModel is one element of the cmdpolicybinding set (convenience
// block that manages systemgroup_systemcmdpolicy_binding entries).
type CmdpolicybindingModel struct {
	Policyname types.String `tfsdk:"policyname"`
	Priority   types.Int64  `tfsdk:"priority"`
}

var cmdpolicybindingAttrTypes = map[string]attr.Type{
	"policyname": types.StringType,
	"priority":   types.Int64Type,
}

// SystemgroupResourceModel describes the resource data model.
type SystemgroupResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedmanagementinterface types.List   `tfsdk:"allowedmanagementinterface"`
	Daystoexpire               types.Int64  `tfsdk:"daystoexpire"`
	Groupname                  types.String `tfsdk:"groupname"`
	Promptstring               types.String `tfsdk:"promptstring"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Warnpriorndays             types.Int64  `tfsdk:"warnpriorndays"`
	Systemusers                types.Set    `tfsdk:"systemusers"`
	Cmdpolicybinding           types.Set    `tfsdk:"cmdpolicybinding"`
}

func (r *SystemgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemgroup resource.",
			},
			"allowedmanagementinterface": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Allowed Management interfaces of the system users in the group. By default allowed from both API and CLI interfaces. If management interface for a group is set to API, then all users under this group will not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"daystoexpire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Password days to expire for system groups. The daystoexpire value ranges from 30 to 255.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the group. Must begin with a letter, number, hash(#) or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the group is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my group\" or 'my group').",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds.If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"warnpriorndays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.",
			},
			"systemusers": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Convenience attribute that manages the systemgroup_systemuser_binding entries for this group. Each element is the name of a system user to bind to the group.",
			},
		},
		Blocks: map[string]schema.Block{
			"cmdpolicybinding": schema.SetNestedBlock{
				Description: "Convenience block that manages the systemgroup_systemcmdpolicy_binding entries for this group.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"policyname": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The name of command policy.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "The priority of the command policy.",
						},
					},
				},
			},
		},
	}
}

// systemgroupGetThePayloadFromthePlan builds the systemgroup add/create payload.
func systemgroupGetThePayloadFromthePlan(ctx context.Context, data *SystemgroupResourceModel) system.Systemgroup {
	tflog.Debug(ctx, "In systemgroupGetThePayloadFromthePlan Function")

	systemgroup := system.Systemgroup{}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		systemgroup.Groupname = data.Groupname.ValueString()
	}
	if !data.Promptstring.IsNull() && !data.Promptstring.IsUnknown() {
		systemgroup.Promptstring = data.Promptstring.ValueString()
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		systemgroup.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Daystoexpire.IsNull() && !data.Daystoexpire.IsUnknown() {
		systemgroup.Daystoexpire = utils.IntPtr(int(data.Daystoexpire.ValueInt64()))
	}
	if !data.Warnpriorndays.IsNull() && !data.Warnpriorndays.IsUnknown() {
		systemgroup.Warnpriorndays = utils.IntPtr(int(data.Warnpriorndays.ValueInt64()))
	}
	if !data.Allowedmanagementinterface.IsNull() && !data.Allowedmanagementinterface.IsUnknown() {
		var iface []string
		data.Allowedmanagementinterface.ElementsAs(ctx, &iface, false)
		systemgroup.Allowedmanagementinterface = iface
	}

	return systemgroup
}

// systemgroupSetAttrFromGet copies the scalar systemgroup attributes from the
// NITRO GET response into the model. It deliberately does NOT touch the
// convenience collections (systemusers / cmdpolicybinding) — those are refreshed
// by their own binding read helpers only when the resource manages them.
func systemgroupSetAttrFromGet(ctx context.Context, data *SystemgroupResourceModel, getResponseData map[string]interface{}) *SystemgroupResourceModel {
	tflog.Debug(ctx, "In systemgroupSetAttrFromGet Function")

	if val, ok := getResponseData["daystoexpire"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Daystoexpire = types.Int64Value(intVal)
		}
	} else {
		data.Daystoexpire = types.Int64Null()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["promptstring"]; ok && val != nil {
		data.Promptstring = types.StringValue(val.(string))
	} else {
		data.Promptstring = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["warnpriorndays"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Warnpriorndays = types.Int64Value(intVal)
		}
	} else {
		data.Warnpriorndays = types.Int64Null()
	}
	if val, ok := getResponseData["allowedmanagementinterface"]; ok && val != nil {
		if lv, ok2 := stringListFromGet(ctx, val); ok2 {
			data.Allowedmanagementinterface = lv
		} else {
			data.Allowedmanagementinterface = types.ListNull(types.StringType)
		}
	} else {
		data.Allowedmanagementinterface = types.ListNull(types.StringType)
	}

	// Set ID for the resource (single unique attribute -> plain value).
	data.Id = types.StringValue(data.Groupname.ValueString())

	return data
}

// stringListFromGet converts a NITRO GET response value (either a []interface{}
// of strings or a single string) into a types.List of strings.
func stringListFromGet(ctx context.Context, val interface{}) (types.List, bool) {
	switch v := val.(type) {
	case []interface{}:
		strs := make([]string, 0, len(v))
		for _, e := range v {
			if s, ok := e.(string); ok {
				strs = append(strs, s)
			}
		}
		lv, d := types.ListValueFrom(ctx, types.StringType, strs)
		if d.HasError() {
			return types.ListNull(types.StringType), false
		}
		return lv, true
	case string:
		lv, d := types.ListValueFrom(ctx, types.StringType, []string{v})
		if d.HasError() {
			return types.ListNull(types.StringType), false
		}
		return lv, true
	}
	return types.ListNull(types.StringType), false
}
