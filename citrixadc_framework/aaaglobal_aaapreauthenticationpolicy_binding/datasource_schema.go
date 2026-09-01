package aaaglobal_aaapreauthenticationpolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaaglobalAaapreauthenticationpolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model so the data
// source can expose the FULL GET projection: the existing lookup/config
// attributes (as Computed outputs) PLUS the read-only attributes the resource
// intentionally omits. Every non-key attribute is Computed.
type AaaglobalAaapreauthenticationpolicyBindingDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Policy   types.String `tfsdk:"policy"` // Required lookup key
	Priority types.Int64  `tfsdk:"priority"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaaglobal_aaapreauthenticationpolicy_binding.json).
	Bindpolicytype types.Int64 `tfsdk:"bindpolicytype"`
}

func AaaglobalAaapreauthenticationpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy to be unbound.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the bound policy",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"bindpolicytype": schema.Int64Attribute{
				Computed:    true,
				Description: "Bound policy type.",
			},
		},
	}
}

// aaaglobal_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet projects a
// NITRO GET response onto the data-source model. Attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaaglobal_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaaglobalAaapreauthenticationpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaaglobal_aaapreauthenticationpolicy_bindingDataSourceSetAttrFromGet Function")

	if v, ok := g["policy"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policy = types.StringValue(utils.AnyToString(v))
	}

	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only attributes.
	data.Bindpolicytype = utils.MapGetInt64(g, "bindpolicytype")
}
