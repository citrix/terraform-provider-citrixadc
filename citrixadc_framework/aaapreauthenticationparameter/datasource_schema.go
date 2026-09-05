package aaapreauthenticationparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaapreauthenticationparameterDataSourceModel is the data-source-specific
// model, decoupled from AaapreauthenticationparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AaapreauthenticationparameterDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Deletefiles             types.String `tfsdk:"deletefiles"`
	Killprocess             types.String `tfsdk:"killprocess"`
	Preauthenticationaction types.String `tfsdk:"preauthenticationaction"`
	Rule                    types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaapreauthenticationparameter.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AaapreauthenticationparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) to and name(s) of the files to be deleted by the EPA tool, as a string of between 1 and 1023 characters.",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the EPA tool.",
			},
			"preauthenticationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny or allow login on the basis of end point analysis results.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, to be evaluated by the EPA tool.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// aaapreauthenticationparameterDataSourceSetAttrFromGet projects a NITRO
// aaapreauthenticationparameter GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func aaapreauthenticationparameterDataSourceSetAttrFromGet(ctx context.Context, data *AaapreauthenticationparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaapreauthenticationparameterDataSourceSetAttrFromGet Function")

	// aaapreauthenticationparameter is a singleton; use a static ID.
	data.Id = types.StringValue("aaapreauthenticationparameter-config")

	// Read/write attributes as read-back outputs.
	data.Deletefiles = utils.MapGetString(g, "deletefiles")
	data.Killprocess = utils.MapGetString(g, "killprocess")
	data.Preauthenticationaction = utils.MapGetString(g, "preauthenticationaction")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
