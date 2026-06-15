package vpnvserver_cspolicy_binding

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

// VpnvserverCspolicyBindingResourceModel describes the resource data model.
type VpnvserverCspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverCspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_cspolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Not echoed back by the NITRO GET response for this binding, so it
				// must not be Computed (Computed + null-from-GET => "inconsistent
				// result after apply"). Optional + RequiresReplace matches SDK v2.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Not echoed back by the NITRO GET response; not Computed.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Next priority expression.",
			},
			"groupextraction": schema.BoolAttribute{
				// Not echoed back by the NITRO GET response; not Computed.
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
				// Not echoed back by the NITRO GET response; not Computed.
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_cspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverCspolicyBindingResourceModel) vpn.Vpnvservercspolicybinding {
	tflog.Debug(ctx, "In vpnvserver_cspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_cspolicy_binding := vpn.Vpnvservercspolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_cspolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_cspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_cspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_cspolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_cspolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_cspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_cspolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_cspolicy_binding
}

// vpnvserverCspolicyBindingComposeId builds the composite ID from the two unique
// keys (name, policy) — matching the SDK v2 ID order and resource_id_mapping.json.
// bindpoint is NOT part of the ID: it is never echoed by the NITRO GET response and
// was never part of the SDK v2 ID.
func vpnvserverCspolicyBindingComposeId(name, policy string) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(name)))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(policy)))
	return strings.Join(idParts, ",")
}

// vpnvserver_cspolicy_bindingSetAttrFromGet is the RESOURCE-side setter. The NITRO
// GET response for this binding only echoes name, policy and priority. The other
// inputs (bindpoint, gotopriorityexpression, groupextraction, secondary) are NOT
// returned, so we must preserve the existing plan/state value instead of nulling
// them (Pattern 7 — otherwise Terraform reports "inconsistent result after apply").
func vpnvserver_cspolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_cspolicy_bindingSetAttrFromGet Function")

	// Echoed identity fields
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	// priority is echoed (as a string) by the GET response
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	// bindpoint, gotopriorityexpression, groupextraction and secondary are NOT
	// echoed by NITRO — preserve the existing plan/state value (do not touch).

	// Set ID for the resource (name,policy)
	data.Id = types.StringValue(vpnvserverCspolicyBindingComposeId(data.Name.ValueString(), data.Policy.ValueString()))

	return data
}

// vpnvserver_cspolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every field present in the GET response and sets the ID (Pattern 7 split).
func vpnvserver_cspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_cspolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (name,policy)
	data.Id = types.StringValue(vpnvserverCspolicyBindingComposeId(data.Name.ValueString(), data.Policy.ValueString()))

	return data
}
