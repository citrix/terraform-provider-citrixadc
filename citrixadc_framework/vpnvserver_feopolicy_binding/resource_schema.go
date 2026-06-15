package vpnvserver_feopolicy_binding

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

// VpnvserverFeopolicyBindingResourceModel describes the resource data model.
type VpnvserverFeopolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverFeopolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_feopolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				// Not Computed: the server returns a normalized default ("END") that
				// differs from user input and is not adopted into state (Pattern 13).
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Next priority expression.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				// Not Computed: NITRO GET does not echo this field back (Pattern 13).
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
				Description: "Name of a policy to bind to the virtual server (for example, the name of an authentication, session, or endpoint analysis policy).",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				// Not Computed: priority is a user-driven input that is not adopted from
				// the GET response into resource state (Pattern 13).
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				// Not Computed: NITRO GET does not echo this field back (Pattern 13).
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserver_feopolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverFeopolicyBindingResourceModel) vpn.Vpnvserverfeopolicybinding {
	tflog.Debug(ctx, "In vpnvserver_feopolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_feopolicy_binding := vpn.Vpnvserverfeopolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		vpnvserver_feopolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnvserver_feopolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnvserver_feopolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_feopolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		vpnvserver_feopolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnvserver_feopolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnvserver_feopolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserver_feopolicy_binding
}

// vpnvserver_feopolicy_bindingSetAttrFromGet is the RESOURCE-side state setter.
// All attributes on this binding are RequiresReplace (none updateable), so the only
// values that ever reach the resource state are those the user configured. The NITRO
// GET response does NOT echo back several inputs (secondary, groupextraction) and
// normalizes others (gotopriorityexpression returns the server default "END", priority
// comes back as a string). To avoid "inconsistent result after apply" / perpetual
// diffs (Pattern 7 + Pattern 13), this setter PRESERVES the user's plan/state values
// for those non-echoed / server-overridden fields and only adopts the parent/identity
// keys (name, policy, bindpoint) which the server reliably returns. NITRO returns the
// bound policy under the "policyname" key (not "policy").
func vpnvserver_feopolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverFeopolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverFeopolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_feopolicy_bindingSetAttrFromGet Function")

	// Identity keys reliably returned by the GET response.
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// NITRO returns the policy name under "policyname".
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}

	// gotopriorityexpression, groupextraction, secondary, priority are either not
	// echoed by the server or server-normalized. Preserve the configured plan/state
	// value rather than overwriting/nulling it. (No-op here — leaving data.* untouched
	// keeps the value that was read from plan/state.)

	return data
}

// vpnvserver_feopolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter (Pattern 7 split). A datasource has no prior plan/state to preserve, so it
// faithfully copies every field the GET response returns and sets the composite ID.
func vpnvserver_feopolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverFeopolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverFeopolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_feopolicy_bindingSetAttrFromGetForDatasource Function")

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
	// NITRO returns the policy name under "policyname".
	if val, ok := getResponseData["policyname"]; ok && val != nil {
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bindpoint:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Bindpoint.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
