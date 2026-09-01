package policypatsetfile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicypatsetfileDataSourceModel is the data-source-specific model, decoupled
// from PolicypatsetfileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (totalpatterns, boundpatterns, patsetname, bindstatuscode, bindstatus). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type PolicypatsetfileDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Charset   types.String `tfsdk:"charset"`
	Comment   types.String `tfsdk:"comment"`
	Delimiter types.String `tfsdk:"delimiter"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policypatsetfile.json). Never settable; populated from GET.
	Totalpatterns  types.Int64  `tfsdk:"totalpatterns"`
	Boundpatterns  types.Int64  `tfsdk:"boundpatterns"`
	Patsetname     types.String `tfsdk:"patsetname"`
	Bindstatuscode types.Int64  `tfsdk:"bindstatuscode"`
	Bindstatus     types.String `tfsdk:"bindstatus"`
}

func PolicypatsetfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"charset": schema.StringAttribute{
				Computed:    true,
				Description: "Character set associated with the characters in the string.",
			},
			"comment": schema.StringAttribute{
				Computed:    true,
				Description: "Any comments to preserve information about this patsetfile.",
			},
			"delimiter": schema.StringAttribute{
				Computed:    true,
				Description: "patset file patterns delimiter.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.",
			},
			"overwrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Computed:    true,
				Description: "URL in protocol, host, path, and file name format from where the patset file will be imported. If file is already present, then it can be imported using local keyword (import patsetfile local:filename patsetfile1)\n                      NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"totalpatterns": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of patterns in the patset file.",
			},
			"boundpatterns": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of patterns bound to a patset.",
			},
			"patsetname": schema.StringAttribute{
				Computed:    true,
				Description: "The patset with which the patsetfile is associated.",
			},
			"bindstatuscode": schema.Int64Attribute{
				Computed:    true,
				Description: "The status code of pattern bindings to patset.",
			},
			"bindstatus": schema.StringAttribute{
				Computed:    true,
				Description: "The status of pattern bindings to patset.",
			},
		},
	}
}

// policypatsetfileDataSourceSetAttrFromGet projects a NITRO policypatsetfile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection. Note the GET echoes "delimiter" as a number (e.g. 10 == newline);
// utils.MapGetString coerces it to a string.
func policypatsetfileDataSourceSetAttrFromGet(ctx context.Context, data *PolicypatsetfileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policypatsetfileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Charset = utils.MapGetString(g, "charset")
	data.Comment = utils.MapGetString(g, "comment")
	data.Delimiter = utils.MapGetString(g, "delimiter")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only metadata.
	data.Totalpatterns = utils.MapGetInt64(g, "totalpatterns")
	data.Boundpatterns = utils.MapGetInt64(g, "boundpatterns")
	data.Patsetname = utils.MapGetString(g, "patsetname")
	data.Bindstatuscode = utils.MapGetInt64(g, "bindstatuscode")
	data.Bindstatus = utils.MapGetString(g, "bindstatus")
}
