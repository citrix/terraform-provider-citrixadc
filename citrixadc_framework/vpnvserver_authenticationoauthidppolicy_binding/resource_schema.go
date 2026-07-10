package vpnvserver_authenticationoauthidppolicy_binding

import (
	"context"
	"fmt"
	"strings"

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

// VpnvserverAuthenticationoauthidppolicyBindingResourceModel describes the resource data model.
type VpnvserverAuthenticationoauthidppolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_authenticationoauthidppolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// NITRO GET for this binding does not echo bindpoint back, so it must not be
				// Computed (Computed would cause "inconsistent result"/known-after-apply churn).
				Optional: true,
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
				Description: "Next priority expression.",
			},
			"groupextraction": schema.BoolAttribute{
				// NITRO GET for this binding does not echo groupextraction back, so it must
				// not be Computed (would cause "inconsistent result" after apply).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server.",
			},
			"policy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy, if any, bound to the VPN virtual server.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				// NITRO GET for this binding does not echo secondary back, so it must not be
				// Computed (would cause "inconsistent result" after apply).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel) vpn.Vpnvserverauthenticationoauthidppolicybinding {
	tflog.Debug(ctx, "In vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_authenticationoauthidppolicy_binding := vpn.Vpnvserverauthenticationoauthidppolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_authenticationoauthidppolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_authenticationoauthidppolicy_binding
}

// Resource-side state setter. The NITRO GET for this binding only echoes back
// name, policy, priority and gotopriorityexpression; it does NOT return bindpoint,
// secondary or groupextraction. Those non-echoed user inputs are PRESERVED from the
// existing plan/state (Pattern 7) so the post-apply state matches the configuration
// and there is no "inconsistent result after apply" error.
func vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationoauthidppolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGet Function")

	// bindpoint, secondary, groupextraction are NOT returned by GET -> preserve existing value.

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// Set ID for the resource - composite of legacy unique attrs (name,policy),
	// matching resource_id_mapping.json so legacy SDK v2 state imports resolve.
	data.Id = types.StringValue(vpnvserver_authenticationoauthidppolicy_bindingComposeId(data))

	return data
}

// Datasource-side state setter. Datasources have no prior plan/state to preserve,
// so this faithfully copies every field the GET response provides and sets the ID.
func vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationoauthidppolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue(vpnvserver_authenticationoauthidppolicy_bindingComposeId(data))

	return data
}

// vpnvserver_authenticationoauthidppolicy_bindingComposeId builds the composite ID
// from the legacy unique attrs (name,policy) in resource_id_mapping.json order.
func vpnvserver_authenticationoauthidppolicy_bindingComposeId(data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	return strings.Join(idParts, ",")
}
