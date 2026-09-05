package appfwxmlerrorpage

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwxmlerrorpageDataSourceModel is the data-source-specific model, decoupled
// from AppfwxmlerrorpageResourceModel. A data source is a pure read surface, so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AppfwxmlerrorpageDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/appfwxmlerrorpage.json). Never settable; populated from GET.
	Response types.String `tfsdk:"response"`
}

func AppfwxmlerrorpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the XML error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the imported xml error page to be removed.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing XML error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported XML error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Response returned by the appliance for the imported XML error object.",
			},
		},
	}
}

// appfwxmlerrorpageDataSourceSetAttrFromGet projects a NITRO appfwxmlerrorpage
// GET response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func appfwxmlerrorpageDataSourceSetAttrFromGet(ctx context.Context, data *AppfwxmlerrorpageDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwxmlerrorpageDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only (GET-only) attribute.
	data.Response = utils.MapGetString(g, "response")
}
