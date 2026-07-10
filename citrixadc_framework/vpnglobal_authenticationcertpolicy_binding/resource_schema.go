package vpnglobal_authenticationcertpolicy_binding

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

// VpnglobalAuthenticationcertpolicyBindingResourceModel describes the resource data model.
type VpnglobalAuthenticationcertpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnglobalAuthenticationcertpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_authenticationcertpolicy_binding resource.",
			},
			// Pattern 13: gotopriorityexpression is never echoed by the NITRO GET and
			// has no server-side default resolvable at apply time. Keeping Computed
			// would leave it "unknown after apply" when the user omits it. Drop
			// Computed so an unset value plans to null.
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			// Pattern 13: groupextraction is likewise not echoed by the NITRO GET;
			// drop Computed so an unset value plans to null instead of unknown.
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
		},
	}
}

func vpnglobal_authenticationcertpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalAuthenticationcertpolicyBindingResourceModel) vpn.Vpnglobalauthenticationcertpolicybinding {
	tflog.Debug(ctx, "In vpnglobal_authenticationcertpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_authenticationcertpolicy_binding := vpn.Vpnglobalauthenticationcertpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_authenticationcertpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnglobal_authenticationcertpolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		vpnglobal_authenticationcertpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnglobal_authenticationcertpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnglobal_authenticationcertpolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnglobal_authenticationcertpolicy_binding
}

func vpnglobal_authenticationcertpolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAuthenticationcertpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationcertpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationcertpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	// Pattern 7: the NITRO GET for this binding does NOT echo back several of the
	// configured inputs (gotopriorityexpression and groupextraction are never
	// returned by the appliance). Nulling them from a missing GET field wipes the
	// user-supplied value from state and produces an "inconsistent result after
	// apply" error. So in the resource setter we PRESERVE the existing plan/state
	// value when the field is absent, and only adopt the GET value when present.
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

	// ID is set once in Create (plain policyname value); do not recompute here.

	return data
}

// vpnglobal_authenticationcertpolicy_bindingSetAttrFromGetForDatasource is the
// datasource counterpart of the setter. The datasource has no prior plan/state to
// preserve, so it faithfully copies every field present in the GET response and
// sets the synthetic ID itself (datasources never call Create). Pattern 7 split.
func vpnglobal_authenticationcertpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalAuthenticationcertpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationcertpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationcertpolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (single unique attribute - plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
