package responderhtmlpage

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderhtmlpageDataSourceModel is the data-source-specific model, decoupled
// from ResponderhtmlpageResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type ResponderhtmlpageDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Cacertfile types.String `tfsdk:"cacertfile"`
	Comment    types.String `tfsdk:"comment"`
	Overwrite  types.Bool   `tfsdk:"overwrite"`
	Src        types.String `tfsdk:"src"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/responderhtmlpage.json). Never settable; from GET.
	Response types.String `tfsdk:"response"`
}

func ResponderhtmlpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate file name which will be used to verify the peer's certificate. The certificate should be imported using \"import ssl certfile\" CLI command or equivalent in API or GUI. If certificate name is not configured, then default root CA certificates are used for peer's certificate verification.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the HTML page object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the HTML page object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported HTML page.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "The imported HTML page response content returned by the appliance.",
			},
		},
	}
}

// responderhtmlpageDataSourceSetAttrFromGet projects a NITRO responderhtmlpage
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func responderhtmlpageDataSourceSetAttrFromGet(ctx context.Context, data *ResponderhtmlpageDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In responderhtmlpageDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// cacertfile/comment/overwrite/src are write-only inputs to the
	// ?action=Import call; the GET never returns them -> Null.
	data.Cacertfile = utils.MapGetString(g, "cacertfile")
	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only attributes.
	data.Response = utils.MapGetString(g, "response")
}
