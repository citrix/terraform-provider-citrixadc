package systemfile

import (
	"context"
	stdpath "path"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemfileDataSourceModel is the data-source-specific model, decoupled from
// SystemfileResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only file metadata the appliance returns on GET
// (fileaccesstime, filemodifiedtime, filemode, filesize) that the resource omits.
type SystemfileDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Filecontent     types.String `tfsdk:"filecontent"`
	Fileencoding    types.String `tfsdk:"fileencoding"`
	Filelocation    types.String `tfsdk:"filelocation"` // Required lookup key
	Filename        types.String `tfsdk:"filename"`     // Required lookup key
	IsBase64Encoded types.Bool   `tfsdk:"is_base64_encoded"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/systemfile.json). Never settable; populated from GET.
	Fileaccesstime   types.String `tfsdk:"fileaccesstime"`
	Filemodifiedtime types.String `tfsdk:"filemodifiedtime"`
	Filemode         types.List   `tfsdk:"filemode"`
	Filesize         types.Int64  `tfsdk:"filesize"`
}

func SystemfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"filecontent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "file content in Base64 format.",
			},
			"fileencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "encoding type of the file content.",
			},
			"filelocation": schema.StringAttribute{
				Required:    true,
				Description: "location of the file on Citrix ADC.",
			},
			"filename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the file. It should not include filepath.",
			},
			// Present in the shared model; not meaningful as a datasource lookup input.
			"is_base64_encoded": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set to true when filecontent is already base64 encoded.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"fileaccesstime": schema.StringAttribute{
				Computed:    true,
				Description: "Last access time of the file.",
			},
			"filemodifiedtime": schema.StringAttribute{
				Computed:    true,
				Description: "Last modified time of the file.",
			},
			"filemode": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "File mode. Possible values: DIRECTORY.",
			},
			"filesize": schema.Int64Attribute{
				Computed:    true,
				Description: "Size of the file in BYTES.",
			},
		},
	}
}

// systemfileDataSourceSetAttrFromGet projects a NITRO systemfile GET response
// onto the data-source model. filecontent is stored as returned by NITRO
// (base64); the remaining attributes are filled via the shared utils.MapGet*
// helpers (Null when the GET omits them).
func systemfileDataSourceSetAttrFromGet(ctx context.Context, data *SystemfileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemfileDataSourceSetAttrFromGet Function")

	// Configurable attributes as read-back outputs. The datasource has no
	// is_base64_encoded decode step, so filecontent is stored as returned (base64).
	data.Filecontent = utils.MapGetString(g, "filecontent")

	if v, ok := g["fileencoding"]; ok && v != nil {
		data.Fileencoding = types.StringValue(utils.AnyToString(v))
	} else {
		data.Fileencoding = types.StringValue("BASE64")
	}

	data.Filelocation = utils.MapGetString(g, "filelocation")
	data.Filename = utils.MapGetString(g, "filename")

	// is_base64_encoded is a client-side-only flag never returned by GET.
	if data.IsBase64Encoded.IsNull() || data.IsBase64Encoded.IsUnknown() {
		data.IsBase64Encoded = types.BoolValue(false)
	}

	// ID scheme matches the resource: path.Join(filelocation, filename).
	data.Id = types.StringValue(stdpath.Join(data.Filelocation.ValueString(), data.Filename.ValueString()))

	// Read-only metadata.
	data.Fileaccesstime = utils.MapGetString(g, "fileaccesstime")
	data.Filemodifiedtime = utils.MapGetString(g, "filemodifiedtime")
	data.Filemode = utils.MapGetStringList(g, "filemode")
	data.Filesize = utils.MapGetInt64(g, "filesize")
}
