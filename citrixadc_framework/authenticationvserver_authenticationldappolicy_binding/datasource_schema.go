package authenticationvserver_authenticationldappolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverAuthenticationldappolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source
// is a pure read surface (Read only; no plan/apply lifecycle), so it can expose
// the FULL GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes that the resource deliberately omits. Every non-key
// attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AuthenticationvserverAuthenticationldappolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"` // Required lookup key
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policy                 types.String `tfsdk:"policy"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationvserver_authenticationldappolicy_binding.json).
	// Never settable; populated from GET.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AuthenticationvserverAuthenticationldappolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
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

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type associated with the policy binding. Returned by the appliance on a GET; null when omitted.",
			},
		},
	}
}

// authenticationvserver_authenticationldappolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func authenticationvserver_authenticationldappolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationldappolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserver_authenticationldappolicy_bindingDataSourceSetAttrFromGet Function")

	// Lookup keys — adopt from GET when present, else keep the configured value.
	if v, ok := g["name"]; ok && v != nil {
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	if v, ok := g["policy"]; ok && v != nil {
		data.Policy = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
	data.Nextfactor = utils.MapGetString(g, "nextfactor")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Secondary = utils.MapGetBool(g, "secondary")

	// Read-only (GET-only) metadata.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite id matching the resource Create ID format (name,policy).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
