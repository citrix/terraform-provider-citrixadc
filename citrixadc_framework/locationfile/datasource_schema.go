package locationfile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationfileDataSourceModel is the data-source-specific model, decoupled from
// LocationfileResourceModel. A data source is a pure read surface, so it exposes
// the existing datasource attributes as Computed outputs PLUS the read-only
// attributes the NITRO GET returns (zion73x_readonly/locationfile.json) that the
// resource intentionally omits.
type LocationfileDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Locationfile types.String `tfsdk:"locationfile"`
	Format       types.String `tfsdk:"format"`
	Src          types.String `tfsdk:"src"`

	// Read-only (GET-only) attributes from the NITRO read-only set.
	Curlocfilestatus  types.String `tfsdk:"curlocfilestatus"`
	Prevlocationfile  types.String `tfsdk:"prevlocationfile"`
	Prevlocfileformat types.String `tfsdk:"prevlocfileformat"`
	Prevlocfilestatus types.String `tfsdk:"prevlocfilestatus"`
	Locfilestatusstr  types.String `tfsdk:"locfilestatusstr"`
}

func LocationfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"locationfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			"format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of the location file. Required for the NetScaler to identify how to read the location file.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"curlocfilestatus": schema.StringAttribute{
				Computed:    true,
				Description: "The status of the current location file (for example Not Loaded, Active, In Progress, Failed).",
			},
			"prevlocationfile": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the previous location file.",
			},
			"prevlocfileformat": schema.StringAttribute{
				Computed:    true,
				Description: "The format of the previous location file.",
			},
			"prevlocfilestatus": schema.StringAttribute{
				Computed:    true,
				Description: "The status of the previous location file (for example Not Loaded, Active, In Progress, Failed).",
			},
			"locfilestatusstr": schema.StringAttribute{
				Computed:    true,
				Description: "Status string of the location file.",
			},
		},
	}
}

// locationfileDataSourceSetAttrFromGet projects a NITRO locationfile GET response
// onto the data-source model. Unlike the resource setter it copies every value
// straight from the GET (including the read-only status metadata) and assigns the
// datasource ID from the location file name. Absent attributes resolve to Null.
// The NITRO GET returns the location file name under the "Locationfile" key.
func locationfileDataSourceSetAttrFromGet(ctx context.Context, data *LocationfileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In locationfileDataSourceSetAttrFromGet Function")

	data.Locationfile = utils.MapGetString(g, "Locationfile")
	data.Format = utils.MapGetString(g, "format")
	// src is not returned by GET (only the separate Import action carries it).
	data.Src = types.StringNull()

	// Read-only status metadata.
	data.Curlocfilestatus = utils.MapGetString(g, "curlocfilestatus")
	data.Prevlocationfile = utils.MapGetString(g, "prevlocationfile")
	data.Prevlocfileformat = utils.MapGetString(g, "prevlocfileformat")
	data.Prevlocfilestatus = utils.MapGetString(g, "prevlocfilestatus")
	data.Locfilestatusstr = utils.MapGetString(g, "locfilestatusstr")

	// Datasource ID mirrors the SDK v2 resource ID scheme (the location file name),
	// falling back to a static handle if the ADC returns no name.
	if v, ok := g["Locationfile"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = types.StringValue("locationfile-config")
	}
}
