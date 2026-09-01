package appfwhtmlerrorpage

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwhtmlerrorpageDataSourceModel is the data-source-specific model, decoupled
// from AppfwhtmlerrorpageResourceModel. A data source is a pure read surface, so
// it exposes the existing lookup/config attributes (as Computed outputs) PLUS the
// read-only (GET-only) attributes the resource deliberately omits.
type AppfwhtmlerrorpageDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attribute from the NITRO read-only set
	// (zion73x_readonly/appfwhtmlerrorpage.json). Populated from GET; never settable.
	Response types.String `tfsdk:"response"`
}

func AppfwhtmlerrorpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the HTML error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the XML error object to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing HTML error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported HTML error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Response returned by the appliance for the appfwhtmlerrorpage operation.",
			},
		},
	}
}

// appfwhtmlerrorpageDataSourceSetAttrFromGet projects a NITRO appfwhtmlerrorpage
// GET response onto the data-source model. Every attribute is filled from the GET
// (or left Null when the GET omits it); id mirrors the returned name key.
func appfwhtmlerrorpageDataSourceSetAttrFromGet(ctx context.Context, data *AppfwhtmlerrorpageDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwhtmlerrorpageDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only (GET-only) attribute.
	data.Response = utils.MapGetString(g, "response")
}
