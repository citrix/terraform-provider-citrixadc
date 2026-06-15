package vpnvserver_authenticationloginschemapolicy_binding

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

// VpnvserverAuthenticationloginschemapolicyBindingResourceModel describes the resource data model.
type VpnvserverAuthenticationloginschemapolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAuthenticationloginschemapolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_authenticationloginschemapolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Optional only (no Computed): NITRO GET does not echo bindpoint, so it can
				// never be resolved server-side; Computed would cause "unknown after apply".
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
				// Optional only (no Computed): NITRO GET does not echo groupextraction.
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
				// Optional only (no Computed): NITRO GET does not echo secondary.
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_authenticationloginschemapolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverAuthenticationloginschemapolicyBindingResourceModel) vpn.Vpnvserverauthenticationloginschemapolicybinding {
	tflog.Debug(ctx, "In vpnvserver_authenticationloginschemapolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_authenticationloginschemapolicy_binding := vpn.Vpnvserverauthenticationloginschemapolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_authenticationloginschemapolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_authenticationloginschemapolicy_binding
}

// vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGet is the RESOURCE-side
// setter. NITRO GET for this binding does NOT echo back bindpoint, secondary or
// groupextraction (they are write-only inputs the server normalizes/consumes). Setting
// them from the (absent) GET response would null the user-configured values and trigger
// an "inconsistent result after apply" error. So preserve the existing plan/state value
// for those three fields and only adopt fields the server actually returns. The ID is set
// once in Create (and preserved in Update); it must NOT be recomputed here. (Patterns 7, 13)
func vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAuthenticationloginschemapolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationloginschemapolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGet Function")

	// bindpoint: not echoed by GET -> preserve existing plan/state value (do not null).
	// secondary: not echoed by GET -> preserve existing plan/state value (do not null).
	// groupextraction: not echoed by GET -> preserve existing plan/state value (do not null).

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

	return data
}

// vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter. A datasource has no prior plan/state to preserve, so it copies
// every field faithfully from the GET response (defaulting the bool fields that the server
// does not echo to false, the NITRO default) and sets the ID. (Pattern 7)
func vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverAuthenticationloginschemapolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAuthenticationloginschemapolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_authenticationloginschemapolicy_bindingSetAttrFromGetForDatasource Function")

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
		data.Groupextraction = types.BoolValue(false)
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
		data.Secondary = types.BoolValue(false)
	}

	// Set ID for the datasource (no Create runs). Composite "name,policy" order.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
