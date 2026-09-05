package policydataset

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicydatasetDataSourceModel is the data-source-specific model, decoupled from
// PolicydatasetResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) plus
// the read-only attributes the resource deliberately omits.
type PolicydatasetDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Dynamic    types.String `tfsdk:"dynamic"`
	Name       types.String `tfsdk:"name"`
	Patsetfile types.String `tfsdk:"patsetfile"`
	Type       types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policydataset.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func PolicydatasetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"dynamic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name \"allow_test\" it can be used dynamically as client.ip.src.equals_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default datasets.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dataset. Must not exceed 127 characters.",
			},
			"patsetfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of value to bind to the dataset.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type (MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// policydatasetDataSourceSetAttrFromGet projects a NITRO policydataset GET response
// onto the data-source model. Attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func policydatasetDataSourceSetAttrFromGet(ctx context.Context, data *PolicydatasetDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policydatasetDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Dynamic = utils.MapGetString(g, "dynamic")
	data.Patsetfile = utils.MapGetString(g, "patsetfile")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
