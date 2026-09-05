package authenticationvserver_authenticationtacacspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

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

// AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel describes the resource data model.
type AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_authenticationtacacspolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"nextfactor": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor",
			},
			"policy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy, if any, bound to the authentication vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "The priority, if any, of the vpn vserver policy.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
					boolplanmodifier.UseStateForUnknown(),
				},
				Description: "Bind the authentication policy to the secondary chain.\nProvides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method.\nBecause user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.",
			},
		},
	}
}

func authenticationvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel) authentication.Authenticationvserverauthenticationtacacspolicybinding {
	tflog.Debug(ctx, "In authenticationvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_authenticationtacacspolicy_binding := authentication.Authenticationvserverauthenticationtacacspolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		authenticationvserver_authenticationtacacspolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return authenticationvserver_authenticationtacacspolicy_binding
}

// authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGet is used by the resource Read.
// The NITRO GET for this binding does not faithfully echo every server-overridden / non-echoed input
// (notably bindpoint, which the SDK v2 resource also did NOT read back). To avoid "inconsistent result
// after apply" / perpetual-diff errors, we only adopt a value from the GET response when the current
// model value is null/unknown (e.g. on import); otherwise we preserve the existing plan/state value.
func authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGet Function")

	// Rule for every Optional+Computed attribute: if the model value is already known
	// (came from config/plan/state), PRESERVE it (so the apply result matches config and
	// there is no perpetual diff). Only when the model value is unknown (e.g. a Computed
	// attribute the user did not configure, or an import) do we resolve it from the GET
	// response — falling back to a known zero value (null for strings/ints, false for
	// bools, which is the NITRO default) when the field is absent from the GET response,
	// so the post-apply state never contains an unknown value.

	// bindpoint: NITRO does not echo this back (SDK v2 also skipped reading it).
	if data.Bindpoint.IsUnknown() {
		if val, ok := getResponseData["bindpoint"]; ok && val != nil {
			data.Bindpoint = types.StringValue(val.(string))
		} else {
			data.Bindpoint = types.StringNull()
		}
	}
	if data.Gotopriorityexpression.IsUnknown() {
		if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
			data.Gotopriorityexpression = types.StringValue(val.(string))
		} else {
			data.Gotopriorityexpression = types.StringNull()
		}
	}
	if data.Groupextraction.IsUnknown() {
		if val, ok := getResponseData["groupextraction"]; ok && val != nil {
			data.Groupextraction = types.BoolValue(val.(bool))
		} else {
			data.Groupextraction = types.BoolValue(false)
		}
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if data.Nextfactor.IsUnknown() {
		if val, ok := getResponseData["nextfactor"]; ok && val != nil {
			data.Nextfactor = types.StringValue(val.(string))
		} else {
			data.Nextfactor = types.StringNull()
		}
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	if data.Priority.IsUnknown() {
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
	if data.Secondary.IsUnknown() {
		if val, ok := getResponseData["secondary"]; ok && val != nil {
			data.Secondary = types.BoolValue(val.(bool))
		} else {
			data.Secondary = types.BoolValue(false)
		}
	}

	// Set ID for the resource (multiple unique attributes - comma-separated key:UrlEncode(value) pairs)
	data.Id = types.StringValue(authenticationvserver_authenticationtacacspolicy_bindingComposeId(data))

	return data
}

// authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource faithfully copies
// every field from the GET response (the datasource has no prior plan/state to preserve) and sets the ID.
func authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	} else {
		data.Nextfactor = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
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

	data.Id = types.StringValue(authenticationvserver_authenticationtacacspolicy_bindingComposeId(data))

	return data
}

// authenticationvserver_authenticationtacacspolicy_bindingComposeId builds the new-format composite ID
// from the unique attributes (groupextraction, name, policy, secondary).
func authenticationvserver_authenticationtacacspolicy_bindingComposeId(data *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupextraction:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secondary:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secondary.ValueBool()))))
	return strings.Join(idParts, ",")
}
