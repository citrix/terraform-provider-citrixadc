package appfwjsonerrorpage

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwjsonerrorpageDataSourceModel is the data-source-specific model,
// decoupled from AppfwjsonerrorpageResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AppfwjsonerrorpageDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwjsonerrorpage.json). Never settable; populated
	// from GET.
	Response types.String `tfsdk:"response"`
}

func AppfwjsonerrorpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the JSON error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the imported json error page to be removed.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing JSON error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported JSON error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "JSON error object contents.",
			},
		},
	}
}

// appfwjsonerrorpageDataSourceSetAttrFromGet projects a NITRO appfwjsonerrorpage
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func appfwjsonerrorpageDataSourceSetAttrFromGet(ctx context.Context, data *AppfwjsonerrorpageDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwjsonerrorpageDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only metadata.
	data.Response = utils.MapGetString(g, "response")
}
