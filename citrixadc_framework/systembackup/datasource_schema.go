package systembackup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystembackupDataSourceModel is the data-source-specific model, decoupled from
// SystembackupResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only metadata the appliance returns on GET (size,
// creationtime, version, createdby, ipaddress) that the resource omits.
type SystembackupDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Filename         types.String `tfsdk:"filename"` // Required lookup key
	Includekernel    types.String `tfsdk:"includekernel"`
	Level            types.String `tfsdk:"level"`
	Skipbackup       types.Bool   `tfsdk:"skipbackup"`
	Uselocaltimezone types.Bool   `tfsdk:"uselocaltimezone"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/systembackup.json). Never settable; populated from GET.
	Size         types.Int64  `tfsdk:"size"`
	Creationtime types.String `tfsdk:"creationtime"`
	Version      types.String `tfsdk:"version"`
	Createdby    types.String `tfsdk:"createdby"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
}

func SystembackupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment specified at the time of creation of the backup file(*.tgz).",
			},
			"filename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the backup file(*.tgz) to be restored.",
			},
			"includekernel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to add kernel in the backup file",
			},
			"level": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Level of data to be backed up.",
			},
			"skipbackup": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to skip taking backup during restore operation",
			},
			"uselocaltimezone": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option will create backup file with local timezone timestamp",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"size": schema.Int64Attribute{
				Computed:    true,
				Description: "Size of the backup file(*.tgz) in KB.",
			},
			"creationtime": schema.StringAttribute{
				Computed:    true,
				Description: "Creation time of the backup file(*.tgz).",
			},
			"version": schema.StringAttribute{
				Computed:    true,
				Description: "Build version of the backup file(*.tgz).",
			},
			"createdby": schema.StringAttribute{
				Computed:    true,
				Description: "Name of user who created the backup file(*.tgz).",
			},
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "Ip of Citrix ADC box where the backup file(*.tgz) was created.",
			},
		},
	}
}

// systembackupDataSourceSetAttrFromGet projects a NITRO systembackup GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func systembackupDataSourceSetAttrFromGet(ctx context.Context, data *SystembackupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systembackupDataSourceSetAttrFromGet Function")

	if v, ok := g["filename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Filename = types.StringValue(utils.AnyToString(v))
	}

	// Configurable attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Includekernel = utils.MapGetString(g, "includekernel")
	data.Level = utils.MapGetString(g, "level")
	data.Skipbackup = utils.MapGetBool(g, "skipbackup")
	data.Uselocaltimezone = utils.MapGetBool(g, "uselocaltimezone")

	// Read-only metadata.
	data.Size = utils.MapGetInt64(g, "size")
	data.Creationtime = utils.MapGetString(g, "creationtime")
	data.Version = utils.MapGetString(g, "version")
	data.Createdby = utils.MapGetString(g, "createdby")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
}
