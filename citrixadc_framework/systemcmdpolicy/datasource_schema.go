package systemcmdpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemcmdpolicyDataSourceModel is the data-source-specific model, decoupled
// from SystemcmdpolicyResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only metadata the appliance returns on GET (builtin,
// feature) that the resource omits.
type SystemcmdpolicyDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Cmdspec    types.String `tfsdk:"cmdspec"`
	Policyname types.String `tfsdk:"policyname"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/systemcmdpolicy.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func SystemcmdpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when a request matches the policy.",
			},
			"cmdspec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Regular expression specifying the data that matches the policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for a command policy. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the policy is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// systemcmdpolicyDataSourceSetAttrFromGet projects a NITRO systemcmdpolicy GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func systemcmdpolicyDataSourceSetAttrFromGet(ctx context.Context, data *SystemcmdpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemcmdpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	// Configurable attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Cmdspec = utils.MapGetString(g, "cmdspec")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
