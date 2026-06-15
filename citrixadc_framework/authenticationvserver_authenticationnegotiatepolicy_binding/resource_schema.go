package authenticationvserver_authenticationnegotiatepolicy_binding

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

// AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel describes the resource data model.
type AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel struct {
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

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_authenticationnegotiatepolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
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
				},
				Description: "The priority, if any, of the vpn vserver policy.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the authentication policy to the secondary chain.\nProvides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method.\nBecause user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.",
			},
		},
	}
}

func authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel) authentication.Authenticationvserverauthenticationnegotiatepolicybinding {
	tflog.Debug(ctx, "In authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_authenticationnegotiatepolicy_binding := authentication.Authenticationvserverauthenticationnegotiatepolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		authenticationvserver_authenticationnegotiatepolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return authenticationvserver_authenticationnegotiatepolicy_binding
}

// authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet is the
// RESOURCE setter. It preserves user-supplied / server-non-echoed inputs (bindpoint,
// gotopriorityexpression, priority) in plan/state instead of overwriting them from the
// GET response (the NITRO GET for this binding does not faithfully echo these back —
// the SDK v2 Read deliberately skipped bindpoint for the same reason). This avoids
// "inconsistent result after apply" / perpetual-diff errors. The ID is preserved
// (set once in Create / by import); this function does not recompute it.
func authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet Function")

	// Convert API response to model.
	// bindpoint is NOT overwritten from the GET response: the NITRO GET does not echo
	// it back consistently (the SDK v2 Read deliberately skipped bindpoint for the same
	// reason), so preserve the configured plan/state value to avoid an
	// "inconsistent result after apply" error.
	// gotopriorityexpression and nextfactor ARE resolved from the GET response (and
	// nulled when absent) because they are Computed and Terraform requires a known
	// value after apply.
	// priority is set from the GET response when present, otherwise preserved (the
	// configured value), since a missing priority in GET would otherwise null a
	// user-supplied value.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	} else {
		data.Nextfactor = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	return data
}

// authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGetForDatasource
// is the DATASOURCE setter. A datasource has no prior plan/state to preserve, so it
// faithfully copies every field from the GET response and sets the ID itself.
func authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGetForDatasource Function")

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
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Set ID for the datasource (composite key:UrlEncode(value), matching Create and
	// resource_id_mapping.json order: name,policy).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
