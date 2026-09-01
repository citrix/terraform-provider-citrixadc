package aaauser_vpnsessionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaauserVpnsessionpolicyBindingDataSourceModel is the data-source-specific model.
// It carries every attribute the data source already exposed PLUS the read-only
// (GET-only) attributes from the NITRO doc read-only set
// (zion73x_readonly/aaauser_vpnsessionpolicy_binding.json). All non-key
// attributes are Computed; the read-only attributes are never settable and are
// populated from the GET response (Null when the appliance omits them).
type AaauserVpnsessionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Username               types.String `tfsdk:"username"`

	// Read-only (GET-only) attributes surfaced by the data source.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaauserVpnsessionpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy Name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},

			// Read-only (GET-only) attributes. Computed; null when the
			// appliance omits them.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
		},
	}
}

// aaauser_vpnsessionpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// aaauser_vpnsessionpolicy_binding GET response onto the data-source model. A
// data source has no plan/apply reconciliation, so attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaauser_vpnsessionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaauserVpnsessionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaauser_vpnsessionpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")
	data.Username = utils.MapGetString(g, "username")

	// Read-only (GET-only) attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set the composite ID (comma-separated key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
