package authenticationvserver_authenticationloginschemapolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverAuthenticationloginschemapolicyBindingDataSourceModel is
// the data-source-specific model, decoupled from the resource model. A data
// source is a pure read surface (Read only; no plan/apply lifecycle), so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AuthenticationvserverAuthenticationloginschemapolicyBindingDataSourceModel struct {
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
	// (zion73x_readonly/authenticationvserver_authenticationloginschemapolicy_binding.json).
	// Never settable; populated from GET.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AuthenticationvserverAuthenticationloginschemapolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
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
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
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

// authenticationvserver_authenticationloginschemapolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func authenticationvserver_authenticationloginschemapolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationloginschemapolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserver_authenticationloginschemapolicy_bindingDataSourceSetAttrFromGet Function")

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
