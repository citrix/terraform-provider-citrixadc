package authenticationvserver_authenticationoauthidppolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverAuthenticationoauthidppolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source is
// a pure read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits. Every non-key attribute is Computed.
type AuthenticationvserverAuthenticationoauthidppolicyBindingDataSourceModel struct {
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
	// (zion73x_readonly/authenticationvserver_authenticationoauthidppolicy_binding.json).
	// Never settable; populated from GET, null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AuthenticationvserverAuthenticationoauthidppolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bindpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind point to which to bind the policy.",
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
				Description: "On success invoke label.",
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

			// Read-only (GET-only) attributes surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "The type of the authentication action associated with this policy binding.",
			},
		},
	}
}

// authenticationvserver_authenticationoauthidppolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model. Because a data source
// has no plan/apply reconciliation, attributes are simply filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func authenticationvserver_authenticationoauthidppolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationoauthidppolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserver_authenticationoauthidppolicy_bindingDataSourceSetAttrFromGet Function")

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

	// Composite ID (name,policy) mirroring the resource id format.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
