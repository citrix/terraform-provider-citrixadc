package vpnglobal_vpnclientlessaccesspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalVpnclientlessaccesspolicyBindingResourceModel describes the resource data model.
type VpnglobalVpnclientlessaccesspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Builtin                types.List   `tfsdk:"builtin"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
	Type                   types.String `tfsdk:"type"`
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnclientlessaccesspolicy_binding resource.",
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
					listplanmodifier.UseStateForUnknown(),
				},
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
			},
			"globalbindtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound",
			},
		},
	}
}

func vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel) vpn.Vpnglobalvpnclientlessaccesspolicybinding {
	tflog.Debug(ctx, "In vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_vpnclientlessaccesspolicy_binding := vpn.Vpnglobalvpnclientlessaccesspolicybinding{}
	if !data.Builtin.IsNull() && !data.Builtin.IsUnknown() {
		builtinList := make([]string, 0, len(data.Builtin.Elements()))
		for _, elem := range data.Builtin.Elements() {
			if strElem, ok := elem.(types.String); ok {
				builtinList = append(builtinList, strElem.ValueString())
			}
		}
		vpnglobal_vpnclientlessaccesspolicy_binding.Builtin = builtinList
	}
	if !data.Feature.IsNull() && !data.Feature.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Secondary = data.Secondary.ValueBool()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Type = data.Type.ValueString()
	}

	return vpnglobal_vpnclientlessaccesspolicy_binding
}

// vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet is the RESOURCE-side
// state setter. The NITRO server overrides several user-supplied inputs for this
// global binding (globalbindtype, type, feature, gotopriorityexpression are
// normalized server-side; secondary/groupextraction/builtin are not echoed back),
// mirroring the SDK v2 resource which deliberately did NOT write globalbindtype/type
// back to state. To avoid Terraform "inconsistent result after apply" errors, this
// setter PRESERVES the configured plan/state value for those fields and only adopts
// the GET value when the model value is null/unknown (covers import, where state
// carries only the ID, and Computed resolution when the attribute was omitted).
// Pattern 7 (server overrides user input) + Pattern 13.
func vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnclientlessaccesspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet Function")

	// builtin: not echoed back by the binding GET; preserve plan/state value.
	if data.Builtin.IsNull() || data.Builtin.IsUnknown() {
		if val, ok := getResponseData["builtin"]; ok && val != nil {
			if rawList, ok := val.([]interface{}); ok {
				elems := make([]attr.Value, 0, len(rawList))
				for _, item := range rawList {
					elems = append(elems, types.StringValue(fmt.Sprintf("%v", item)))
				}
				if listVal, diags := types.ListValue(types.StringType, elems); !diags.HasError() {
					data.Builtin = listVal
				} else {
					data.Builtin = types.ListNull(types.StringType)
				}
			} else {
				data.Builtin = types.ListNull(types.StringType)
			}
		} else {
			data.Builtin = types.ListNull(types.StringType)
		}
	}
	// feature: server-normalized (returns "SYSTEM"); preserve plan/state value.
	if data.Feature.IsNull() || data.Feature.IsUnknown() {
		if val, ok := getResponseData["feature"]; ok && val != nil {
			data.Feature = types.StringValue(val.(string))
		} else {
			data.Feature = types.StringNull()
		}
	}
	// globalbindtype: server-normalized (e.g. "SYSTEM_GLOBAL"); preserve plan/state value.
	if data.Globalbindtype.IsNull() || data.Globalbindtype.IsUnknown() {
		if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
			data.Globalbindtype = types.StringValue(val.(string))
		} else {
			data.Globalbindtype = types.StringNull()
		}
	}
	// gotopriorityexpression: server-normalized (e.g. "END"); preserve plan/state value.
	if data.Gotopriorityexpression.IsNull() || data.Gotopriorityexpression.IsUnknown() {
		if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
			data.Gotopriorityexpression = types.StringValue(val.(string))
		} else {
			data.Gotopriorityexpression = types.StringNull()
		}
	}
	// groupextraction: not echoed back; preserve plan/state value.
	if data.Groupextraction.IsNull() || data.Groupextraction.IsUnknown() {
		if val, ok := getResponseData["groupextraction"]; ok && val != nil {
			data.Groupextraction = types.BoolValue(val.(bool))
		} else {
			data.Groupextraction = types.BoolNull()
		}
	}
	// policyname is the identity key; adopt from GET (matches configured value).
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	// priority: honored by the server; preserve plan/state value, fall back to GET.
	if data.Priority.IsNull() || data.Priority.IsUnknown() {
		if val, ok := getResponseData["priority"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Priority = types.Int64Value(intVal)
			} else {
				data.Priority = types.Int64Null()
			}
		} else {
			data.Priority = types.Int64Null()
		}
	}
	// secondary: not echoed back when false; preserve plan/state value.
	if data.Secondary.IsNull() || data.Secondary.IsUnknown() {
		if val, ok := getResponseData["secondary"]; ok && val != nil {
			data.Secondary = types.BoolValue(val.(bool))
		} else {
			data.Secondary = types.BoolNull()
		}
	}
	// type: server-normalized (e.g. "REQ_DEFAULT"); preserve plan/state value.
	if data.Type.IsNull() || data.Type.IsUnknown() {
		if val, ok := getResponseData["type"]; ok && val != nil {
			data.Type = types.StringValue(val.(string))
		} else {
			data.Type = types.StringNull()
		}
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

// vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter. The datasource has no prior plan/state, so it faithfully
// copies every field from the GET response (Pattern 7 datasource split).
func vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnclientlessaccesspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["builtin"]; ok && val != nil {
		if rawList, ok := val.([]interface{}); ok {
			elems := make([]attr.Value, 0, len(rawList))
			for _, item := range rawList {
				elems = append(elems, types.StringValue(fmt.Sprintf("%v", item)))
			}
			if listVal, diags := types.ListValue(types.StringType, elems); !diags.HasError() {
				data.Builtin = listVal
			} else {
				data.Builtin = types.ListNull(types.StringType)
			}
		} else {
			data.Builtin = types.ListNull(types.StringType)
		}
	} else {
		data.Builtin = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	} else {
		data.Globalbindtype = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the datasource (no Create runs for a datasource).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
