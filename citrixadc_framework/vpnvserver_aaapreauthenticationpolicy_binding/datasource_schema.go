package vpnvserver_aaapreauthenticationpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverAaapreauthenticationpolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source
// is a pure read surface, so it can expose the read/write attributes (as Computed
// outputs) AND the read-only attributes the appliance returns on a GET
// (zion73x_readonly/vpnvserver_aaapreauthenticationpolicy_binding.json) that the
// resource intentionally omits.
type VpnvserverAaapreauthenticationpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`   // Required lookup key (parent)
	Policy                 types.String `tfsdk:"policy"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`

	// Read-only (GET-only) attributes surfaced only by the data source.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverAaapreauthenticationpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bindpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy, if any, bound to the VPN virtual server.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},

			// Read-only (GET-only) attributes exposed only by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "The bound entity (action) type, as returned by the appliance. GET-only; null when the appliance omits it.",
			},
		},
	}
}

// vpnvserver_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet projects a
// NITRO GET response onto the data-source model. A data source has no plan/apply
// reconciliation, so attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func vpnvserver_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverAaapreauthenticationpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
	data.Name = utils.MapGetString(g, "name")
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Secondary = utils.MapGetBool(g, "secondary")

	// Read-only attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID (legacy SDK v2 order: name,policy).
	data.Id = types.StringValue(fmt.Sprintf("name:%s,policy:%s", utils.UrlEncode(data.Name.ValueString()), utils.UrlEncode(data.Policy.ValueString())))
}
