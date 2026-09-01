package vpnportaltheme

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnportalthemeDataSourceModel is the data-source-specific model, decoupled
// from VpnportalthemeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnportalthemeDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Basetheme types.String `tfsdk:"basetheme"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnportaltheme.json). Never settable; populated from GET.
	Feature types.String `tfsdk:"feature"`
}

func VpnportalthemeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basetheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the uitheme",
			},

			// Read-only (GET-only) metadata surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// vpnportalthemeDataSourceSetAttrFromGet projects a NITRO vpnportaltheme GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func vpnportalthemeDataSourceSetAttrFromGet(ctx context.Context, data *VpnportalthemeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnportalthemeDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Basetheme = utils.MapGetString(g, "basetheme")

	// Read-only metadata.
	data.Feature = utils.MapGetString(g, "feature")
}
