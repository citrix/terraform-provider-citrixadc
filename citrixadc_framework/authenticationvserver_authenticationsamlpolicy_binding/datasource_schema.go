package authenticationvserver_authenticationsamlpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverAuthenticationsamlpolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source is
// a pure read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits. Every non-key attribute is Computed.
type AuthenticationvserverAuthenticationsamlpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationvserver_authenticationsamlpolicy_binding.json).
	// Never settable; populated from GET, null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AuthenticationvserverAuthenticationsamlpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"nextfactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy, if any, bound to the authentication vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority, if any, of the vpn vserver policy.",
			},
			"secondary": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the authentication policy to the secondary chain.\nProvides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method.\nBecause user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "The type of the authentication action associated with this policy binding.",
			},
		},
	}
}

// authenticationvserver_authenticationsamlpolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model. Because a data source
// has no plan/apply reconciliation, attributes are simply filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func authenticationvserver_authenticationsamlpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationsamlpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserver_authenticationsamlpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
	data.Name = utils.MapGetString(g, "name")
	data.Nextfactor = utils.MapGetString(g, "nextfactor")
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Secondary = utils.MapGetBool(g, "secondary")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID (groupextraction,name,policy,secondary) mirroring the resource id format.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupextraction:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secondary:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secondary.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
