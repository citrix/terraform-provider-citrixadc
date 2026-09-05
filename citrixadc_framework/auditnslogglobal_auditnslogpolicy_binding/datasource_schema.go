package auditnslogglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditnslogglobalAuditnslogpolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from
// AuditnslogglobalAuditnslogpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attribute the resource deliberately omits (numpol).
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type AuditnslogglobalAuditnslogpolicyBindingDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Builtin        types.List   `tfsdk:"builtin"`
	Globalbindtype types.String `tfsdk:"globalbindtype"`
	Policyname     types.String `tfsdk:"policyname"`
	Priority       types.Int64  `tfsdk:"priority"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/auditnslogglobal_auditnslogpolicy_binding.json). Never
	// settable; populated from GET.
	Numpol types.Int64 `tfsdk:"numpol"`
}

func AuditnslogglobalAuditnslogpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"globalbindtype": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit nslog policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},

			// Read-only (GET-only) attribute surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "number of polices bound to label.",
			},
		},
	}
}

// auditnslogglobal_auditnslogpolicy_bindingDataSourceSetAttrFromGet projects a
// NITRO auditnslogglobal_auditnslogpolicy_binding GET response onto the
// data-source model. A data source has no plan/apply reconciliation, so
// attributes are simply filled from the GET (or left Null when the GET omits
// them) via the shared utils.MapGet* helpers, and the composite ID is set.
func auditnslogglobal_auditnslogpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuditnslogglobalAuditnslogpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditnslogglobal_auditnslogpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only (GET-only) attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set composite ID for the datasource.
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("globalbindtype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Globalbindtype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
