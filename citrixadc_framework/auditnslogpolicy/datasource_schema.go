package auditnslogpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditnslogpolicyDataSourceModel is the data-source-specific model, decoupled
// from AuditnslogpolicyResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes the resource deliberately omits.
type AuditnslogpolicyDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"` // Required lookup key
	Rule   types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditnslogpolicy.json). Never settable; populated from GET.
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
	Expressiontype types.String `tfsdk:"expressiontype"`
}

func AuditnslogpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Nslog server action that is performed when this policy matches.\nNOTE: An nslog server action must be associated with an nslog audit policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy.\nMust begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my nslog policy\" or 'my nslog policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the nslog server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"expressiontype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy (Classic/Advanced). Possible values = Classic Policy, Advanced Policy.",
			},
		},
	}
}

// auditnslogpolicyDataSourceSetAttrFromGet projects a NITRO auditnslogpolicy GET
// response onto the data-source model. Attributes are filled from the GET (or left
// Null when the GET omits them) using the shared utils.MapGet* helpers.
func auditnslogpolicyDataSourceSetAttrFromGet(ctx context.Context, data *AuditnslogpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditnslogpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Expressiontype = utils.MapGetString(g, "expressiontype")
}
