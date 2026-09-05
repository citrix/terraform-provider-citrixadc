package fis

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FisDataSourceModel is the data-source-specific model, decoupled from
// FisResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (ifaces). Every non-key attribute is Computed.
type FisDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Ownernode types.Int64  `tfsdk:"ownernode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/fis.json). Never settable; populated from GET.
	Ifaces types.String `tfsdk:"ifaces"`
}

func FisDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the FIS to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Note: In a cluster setup, the FIS name on each node must be unique.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Interfaces to be bound to the FIS, in slot/port notation (for example, 1/3).",
			},
		},
	}
}

// fisDataSourceSetAttrFromGet projects a NITRO fis GET response onto the
// data-source model. Attributes are filled from the GET (or left Null when the
// GET omits them) via the shared utils.MapGet* helpers.
func fisDataSourceSetAttrFromGet(ctx context.Context, data *FisDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In fisDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ownernode = utils.MapGetInt64(g, "ownernode")

	// Read-only attributes.
	data.Ifaces = utils.MapGetString(g, "ifaces")
}
