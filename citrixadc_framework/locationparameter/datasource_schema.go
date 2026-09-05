package locationparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationparameterDataSourceModel is the data-source-specific model, decoupled
// from LocationparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as
// Computed outputs) AND the read-only attributes the resource deliberately omits
// (loaded database file, counters, status flags, ...). Every non-key attribute
// is Computed.
type LocationparameterDataSourceModel struct {
	Id types.String `tfsdk:"id"`

	// Configurable attributes, surfaced here as Computed outputs.
	Context            types.String `tfsdk:"context"`
	Matchwildcardtoany types.String `tfsdk:"matchwildcardtoany"`
	Q1label            types.String `tfsdk:"q1label"`
	Q2label            types.String `tfsdk:"q2label"`
	Q3label            types.String `tfsdk:"q3label"`
	Q4label            types.String `tfsdk:"q4label"`
	Q5label            types.String `tfsdk:"q5label"`
	Q6label            types.String `tfsdk:"q6label"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/locationparameter.json). Never settable; populated from
	// GET.
	Locationfile  types.String `tfsdk:"locationfile"`
	Format        types.String `tfsdk:"format"`
	Custom        types.Int64  `tfsdk:"custom"`
	Static        types.Int64  `tfsdk:"static"`
	Lines         types.Int64  `tfsdk:"lines"`
	Errors        types.Int64  `tfsdk:"errors"`
	Warnings      types.Int64  `tfsdk:"warnings"`
	Entries       types.Int64  `tfsdk:"entries"`
	Locationfile6 types.String `tfsdk:"locationfile6"`
	Format6       types.String `tfsdk:"format6"`
	Custom6       types.Int64  `tfsdk:"custom6"`
	Static6       types.Int64  `tfsdk:"static6"`
	Lines6        types.Int64  `tfsdk:"lines6"`
	Errors6       types.Int64  `tfsdk:"errors6"`
	Warnings6     types.Int64  `tfsdk:"warnings6"`
	Entries6      types.Int64  `tfsdk:"entries6"`
	Flags         types.Int64  `tfsdk:"flags"`
	Status        types.Int64  `tfsdk:"status"`
	Databasemode  types.String `tfsdk:"databasemode"`
	Flushing      types.String `tfsdk:"flushing"`
	Loading       types.String `tfsdk:"loading"`
	Builtin       types.List   `tfsdk:"builtin"`
	Feature       types.String `tfsdk:"feature"`
}

func LocationparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read the global location parameters configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"context": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate.",
			},
			"matchwildcardtoany": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether wildcard qualifiers should match any other\nqualifier including non-wildcard while evaluating\nlocation based expressions.\nPossible values: Yes, No, Expression.\n    Yes - Wildcard qualifiers match any other qualifiers.\n    No  - Wildcard qualifiers do not match non-wildcard\n          qualifiers, but match other wildcard qualifiers.\n    Expression - Wildcard qualifiers in an expression\n          match any qualifier in an LDNS location,\n          wildcard qualifiers in the LDNS location do not match\n          non-wildcard qualifiers in an expression",
			},
			"q1label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the first qualifier. Can be specified for custom context only.",
			},
			"q2label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the second qualifier. Can be specified for custom context only.",
			},
			"q3label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the third qualifier. Can be specified for custom context only.",
			},
			"q4label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fourth qualifier. Can be specified for custom context only.",
			},
			"q5label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fifth qualifier. Can be specified for custom context only.",
			},
			"q6label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the sixth qualifier. Can be specified for custom context only.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"locationfile": schema.StringAttribute{
				Computed:    true,
				Description: "Currently loaded location database file.",
			},
			"format": schema.StringAttribute{
				Computed:    true,
				Description: "Location file format.",
			},
			"custom": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of configured custom locations.",
			},
			"static": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of configured locations in the database file (static locations).",
			},
			"lines": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of lines in the database files.",
			},
			"errors": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of errors encountered while reading the database file.",
			},
			"warnings": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of warnings encountered while reading the database file.",
			},
			"entries": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successfully added entries.",
			},
			"locationfile6": schema.StringAttribute{
				Computed:    true,
				Description: "Currently loaded location database file (IPv6).",
			},
			"format6": schema.StringAttribute{
				Computed:    true,
				Description: "Location file format (IPv6).",
			},
			"custom6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of configured custom locations (IPv6).",
			},
			"static6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of configured locations in the database file (static locations, IPv6).",
			},
			"lines6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of lines in the database files (IPv6).",
			},
			"errors6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of errors encountered while reading the database file (IPv6).",
			},
			"warnings6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of warnings encountered while reading the database file (IPv6).",
			},
			"entries6": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successfully added entries (IPv6).",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Information needed for display. This argument passes information from the kernel to the user space.",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "This argument displays the status (success or failure) of database loading.",
			},
			"databasemode": schema.StringAttribute{
				Computed:    true,
				Description: "This argument displays the database mode. Possible values: File, Internal, Not applicable.",
			},
			"flushing": schema.StringAttribute{
				Computed:    true,
				Description: "This argument displays the state of flushing. Possible values: In progress, Idle.",
			},
			"loading": schema.StringAttribute{
				Computed:    true,
				Description: "This argument displays the state of loading. Possible values: In progress, Idle.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flags indicating the built-in nature of the configuration. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
		},
	}
}

// locationparameterDataSourceSetAttrFromGet projects a NITRO locationparameter
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func locationparameterDataSourceSetAttrFromGet(ctx context.Context, data *LocationparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In locationparameterDataSourceSetAttrFromGet Function")

	// locationparameter is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("locationparameter-config")

	// Configurable attributes as read-back outputs.
	data.Context = utils.MapGetString(g, "context")
	data.Matchwildcardtoany = utils.MapGetString(g, "matchwildcardtoany")
	data.Q1label = utils.MapGetString(g, "q1label")
	data.Q2label = utils.MapGetString(g, "q2label")
	data.Q3label = utils.MapGetString(g, "q3label")
	data.Q4label = utils.MapGetString(g, "q4label")
	data.Q5label = utils.MapGetString(g, "q5label")
	data.Q6label = utils.MapGetString(g, "q6label")

	// Read-only attributes.
	data.Locationfile = utils.MapGetString(g, "Locationfile")
	data.Format = utils.MapGetString(g, "format")
	data.Custom = utils.MapGetInt64(g, "custom")
	data.Static = utils.MapGetInt64(g, "Static")
	data.Lines = utils.MapGetInt64(g, "lines")
	data.Errors = utils.MapGetInt64(g, "errors")
	data.Warnings = utils.MapGetInt64(g, "warnings")
	data.Entries = utils.MapGetInt64(g, "entries")
	data.Locationfile6 = utils.MapGetString(g, "locationfile6")
	data.Format6 = utils.MapGetString(g, "format6")
	data.Custom6 = utils.MapGetInt64(g, "custom6")
	data.Static6 = utils.MapGetInt64(g, "static6")
	data.Lines6 = utils.MapGetInt64(g, "lines6")
	data.Errors6 = utils.MapGetInt64(g, "errors6")
	data.Warnings6 = utils.MapGetInt64(g, "warnings6")
	data.Entries6 = utils.MapGetInt64(g, "entries6")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Status = utils.MapGetInt64(g, "status")
	data.Databasemode = utils.MapGetString(g, "databasemode")
	data.Flushing = utils.MapGetString(g, "flushing")
	data.Loading = utils.MapGetString(g, "loading")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
