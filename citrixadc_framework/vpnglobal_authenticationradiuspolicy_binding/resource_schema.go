package vpnglobal_authenticationradiuspolicy_binding

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

// VpnglobalAuthenticationradiuspolicyBindingResourceModel describes the resource data model.
type VpnglobalAuthenticationradiuspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_authenticationradiuspolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Optional only (no Computed): the NITRO binding GET never echoes this
				// field, so a Computed flag would leave an unset value Unknown after
				// apply (Pattern 13 / Pattern 8). SDK v2 contract was Optional.
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

func vpnglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalAuthenticationradiuspolicyBindingResourceModel) vpn.Vpnglobalauthenticationradiuspolicybinding {
	tflog.Debug(ctx, "In vpnglobal_authenticationradiuspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_authenticationradiuspolicy_binding := vpn.Vpnglobalauthenticationradiuspolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_authenticationradiuspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnglobal_authenticationradiuspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		vpnglobal_authenticationradiuspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnglobal_authenticationradiuspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnglobal_authenticationradiuspolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnglobal_authenticationradiuspolicy_binding
}

func vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAuthenticationradiuspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationradiuspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGet Function")

	// Resource setter: every attribute is RequiresReplace (write-once). The NITRO
	// binding GET omits non-echoed / false-boolean / default fields (omitempty), so
	// only overwrite a model field when the GET response actually carries a value;
	// otherwise preserve the prior plan/state value to avoid "inconsistent result
	// after apply" diffs (Pattern 7 / Pattern 13).
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

// vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource faithfully
// copies every field from the GET response (the datasource has no prior plan/state to
// preserve) and sets the datasource ID. (Pattern 7 datasource split.)
func vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalAuthenticationradiuspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationradiuspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (no Create runs)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
