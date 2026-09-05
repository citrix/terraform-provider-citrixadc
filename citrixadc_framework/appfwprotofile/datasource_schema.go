package appfwprotofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwprotofileDataSourceModel is the data-source-specific model, decoupled
// from AppfwprotofileResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (response). Every non-key attribute is Computed.
type AppfwprotofileDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwprotofile.json). Never settable; populated from
	// GET, Null when the appliance omits them.
	Response types.String `tfsdk:"response"`
}

func AppfwprotofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this gRPC schema file.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the gRPC schema object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing gRPC schema object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates source path of the gRPC schema file.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "gRPC import object response contents.",
			},
		},
	}
}

// appfwprotofileDataSourceSetAttrFromGet projects a NITRO appfwprotofile GET
// response onto the data-source model via the shared utils.MapGet* helpers.
func appfwprotofileDataSourceSetAttrFromGet(ctx context.Context, data *AppfwprotofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwprotofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// comment / overwrite are write-only Import inputs the GET never echoes -> Null.
	data.Comment = types.StringNull()
	data.Overwrite = types.BoolNull()
	data.Src = utils.MapGetString(g, "src")

	// Read-only attribute.
	data.Response = utils.MapGetString(g, "response")
}
