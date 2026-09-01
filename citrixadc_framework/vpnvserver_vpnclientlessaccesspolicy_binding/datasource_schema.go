package vpnvserver_vpnclientlessaccesspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverVpnclientlessaccesspolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnvserverVpnclientlessaccesspolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`   // Required lookup key
	Policy                 types.String `tfsdk:"policy"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_vpnclientlessaccesspolicy_binding.json). Never
	// settable; populated from GET, Null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverVpnclientlessaccesspolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bindpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Next priority expression.",
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

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
		},
	}
}

// vpnvserver_vpnclientlessaccesspolicy_bindingDataSourceSetAttrFromGet projects a
// NITRO vpnvserver_vpnclientlessaccesspolicy_binding GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func vpnvserver_vpnclientlessaccesspolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverVpnclientlessaccesspolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_vpnclientlessaccesspolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
	data.Name = utils.MapGetString(g, "name")
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Secondary = utils.MapGetBool(g, "secondary")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Build the composite id using the legacy attribute order (bindpoint, name, policy).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bindpoint:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Bindpoint.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
