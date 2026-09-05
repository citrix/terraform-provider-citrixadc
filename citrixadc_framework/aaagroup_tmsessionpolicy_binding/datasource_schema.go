package aaagroup_tmsessionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaagroupTmsessionpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model so the data source can expose the
// FULL GET projection: the existing lookup/config attributes (as Computed
// outputs) PLUS the read-only attributes the resource intentionally omits.
type AaagroupTmsessionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"` // Required lookup key
	Policy                 types.String `tfsdk:"policy"`    // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaagroup_tmsessionpolicy_binding.json).
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaagroupTmsessionpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
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
				Computed:    true,
				Description: "Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Bound policy action type.",
			},
		},
	}
}

// aaagroup_tmsessionpolicy_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func aaagroup_tmsessionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaagroupTmsessionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaagroup_tmsessionpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	// groupname is the parent lookup key; preserve the config value when the
	// binding GET response does not echo it back.
	if v := utils.MapGetString(g, "groupname"); !v.IsNull() {
		data.Groupname = v
	}
	data.Policy = utils.MapGetString(g, "policy")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set ID: comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
