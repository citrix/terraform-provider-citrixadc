package vpnvserver_authenticationsamlidppolicy_binding

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

// VpnvserverAuthenticationsamlidppolicyBindingResourceModel describes the resource data model.
type VpnvserverAuthenticationsamlidppolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAuthenticationsamlidppolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_authenticationsamlidppolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Optional only (no Computed): bindpoint is not echoed back by the NITRO
				// GET response, so a Computed flag would leave it "known after apply"
				// forever and fail with an unknown-value error (Pattern 13).
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
				// Optional only (no Computed): not echoed back by NITRO GET (Pattern 13).
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
				// Optional only (no Computed): not echoed back by NITRO GET (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_authenticationsamlidppolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverAuthenticationsamlidppolicyBindingResourceModel) vpn.Vpnvserverauthenticationsamlidppolicybinding {
	tflog.Debug(ctx, "In vpnvserver_authenticationsamlidppolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_authenticationsamlidppolicy_binding := vpn.Vpnvserverauthenticationsamlidppolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_authenticationsamlidppolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_authenticationsamlidppolicy_binding
}

// vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGet updates the resource
// model from a GET response while PRESERVING the prior plan/state values for the
// attributes the NITRO server does not echo back (bindpoint, secondary,
// groupextraction). The live GET only returns name, policy, priority and
// gotopriorityexpression, so blindly nulling the non-echoed fields would wipe the
// user-configured values and cause an "inconsistent result after apply" error
// (Patterns 7 and 13). The ID is composed once in Create (name,policy) and is not
// recomputed here (Pattern 6).
func vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAuthenticationsamlidppolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationsamlidppolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGet Function")

	// bindpoint is not echoed back by the NITRO GET response - preserve the
	// existing plan/state value rather than nulling it (Pattern 7).
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	// groupextraction is not echoed back by the NITRO GET response - preserve.
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
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
	// secondary is not echoed back by the NITRO GET response - preserve.
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGetForDatasource faithfully
// copies every field from the GET response and sets the datasource ID. The datasource
// has no prior plan/state to preserve, so unlike the resource setter it must read all
// values directly from the response (Pattern 7 datasource split).
func vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverAuthenticationsamlidppolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationsamlidppolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationsamlidppolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
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
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	// Set ID for the datasource (legacy order name,policy)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
