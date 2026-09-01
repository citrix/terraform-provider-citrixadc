package appfwarchive

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwarchiveDataSourceModel is the data-source-specific model, decoupled from
// AppfwarchiveResourceModel. A data source is a pure read surface, so it exposes
// the existing lookup/config attributes (as Computed outputs) PLUS the read-only
// (GET-only) attributes the resource deliberately omits.
type AppfwarchiveDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"` // Required lookup key
	Src     types.String `tfsdk:"src"`
	Target  types.String `tfsdk:"target"`

	// Read-only (GET-only) attribute from the NITRO read-only set
	// (zion73x_readonly/appfwarchive.json). Populated from GET; never settable.
	Response types.String `tfsdk:"response"`
}

func AppfwarchiveDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this archive.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of tar archive",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates the source of the tar archive file as a URL\nof the form\n\n    <protocol>://<host>[:<port>][/<path>]\n\n<protocol> is http or https.\n<host> is the DNS name or IP address of the http or https server.\n<port> is the port number of the server. If omitted, the\ndefault port for http or https will be used.\n<path> is the path of the file on the server.\n\nImport will fail if an https server requires client\ncertificate authentication.",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the file to be exported",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Response returned by the appliance for the appfwarchive operation.",
			},
		},
	}
}

// appfwarchiveDataSourceSetAttrFromGet projects a NITRO appfwarchive GET response
// onto the data-source model. The appfwarchive GET (all) response carries no
// per-archive identifying fields, so the caller-supplied name is preserved and
// mirrored to id; the write-only Import/export inputs are Null.
func appfwarchiveDataSourceSetAttrFromGet(ctx context.Context, data *AppfwarchiveDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwarchiveDataSourceSetAttrFromGet Function")

	// name is the required lookup key supplied via config; the appfwarchive
	// GET (all) response echoes no per-archive fields, so mirror it to id.
	data.Id = types.StringValue(utils.AnyToString(data.Name.ValueString()))

	// comment / src / target are write-only Import/export inputs the GET never
	// returns -> Null.
	data.Comment = types.StringNull()
	data.Src = types.StringNull()
	data.Target = types.StringNull()

	// Read-only (GET-only) attribute.
	data.Response = utils.MapGetString(g, "response")
}
