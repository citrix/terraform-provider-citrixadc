package vpnglobal_authenticationtacacspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalAuthenticationtacacspolicyBindingResourceModel describes the resource data model.
type VpnglobalAuthenticationtacacspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_authenticationtacacspolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
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
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.",
			},
		},
	}
}

func vpnglobal_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalAuthenticationtacacspolicyBindingResourceModel) vpn.Vpnglobalauthenticationtacacspolicybinding {
	tflog.Debug(ctx, "In vpnglobal_authenticationtacacspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_authenticationtacacspolicy_binding := vpn.Vpnglobalauthenticationtacacspolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_authenticationtacacspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnglobal_authenticationtacacspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		vpnglobal_authenticationtacacspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnglobal_authenticationtacacspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnglobal_authenticationtacacspolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnglobal_authenticationtacacspolicy_binding
}

// vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGet is the RESOURCE-side setter.
// All schema attributes are RequiresReplace (pure user inputs). The NITRO GET for this
// binding does NOT echo every field back (e.g. groupextraction and gotopriorityexpression
// are omitted when unset, and secondary is omitted when false), so we must NOT null out
// the configured plan/state values for fields the GET does not return — that would cause
// "inconsistent result after apply" / perpetual diffs. We only adopt a GET value when the
// API actually returns it; otherwise we preserve whatever is already in data (plan/state).
// Pattern 7 (server non-echoed inputs) + Pattern 13.
func vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAuthenticationtacacspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationtacacspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGet Function")

	// Convert API response to model — preserve existing value if the field is not echoed.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

// vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter. A datasource has no prior plan/state to preserve, so it must
// faithfully copy every field from the GET response (nulling fields the API omits) and
// set its own ID. Pattern 7 datasource split.
func vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalAuthenticationtacacspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationtacacspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource Function")

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
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Set ID for the datasource (no Create to set it)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
