package aaagroup_vpnsessionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaagroupVpnsessionpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the FULL GET projection:
// the read/write attributes (as Computed outputs) AND the read-only attributes
// that the resource deliberately omits (e.g. acttype). Every non-key attribute is
// Computed.
type AaagroupVpnsessionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaagroup_vpnsessionpolicy_binding.json).
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaagroupVpnsessionpolicyBindingDataSourceSchema() schema.Schema {
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
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},

			// Read-only (GET-only) attribute surfaced only by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding. Read-only value returned by the appliance.",
			},
		},
	}
}

// aaagroup_vpnsessionpolicy_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when the
// GET omits them). The shared utils.MapGet* helpers implement that projection.
func aaagroup_vpnsessionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaagroupVpnsessionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaagroup_vpnsessionpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupname = utils.MapGetString(g, "groupname")
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID (groupname + policy).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
