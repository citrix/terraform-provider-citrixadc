package appfwwsdl

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwwsdlDataSourceModel is the data-source-specific model, decoupled from
// AppfwwsdlResourceModel. A data source is a pure read surface (Read only; no
// plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes the
// resource deliberately omits. The Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares.
type AppfwwsdlDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/appfwwsdl.json). Never settable; populated from GET.
	Response types.String `tfsdk:"response"`
}

func AppfwwsdlDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the WSDL.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the WSDL file to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing WSDL of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and name) of the WSDL file to be imported is stored.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "WSDL response returned by the appliance.",
			},
		},
	}
}

// appfwwsdlDataSourceSetAttrFromGet projects a NITRO appfwwsdl GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func appfwwsdlDataSourceSetAttrFromGet(ctx context.Context, data *AppfwwsdlDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwwsdlDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs (Null when the GET omits them).
	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only (GET-only) attribute.
	data.Response = utils.MapGetString(g, "response")
}
