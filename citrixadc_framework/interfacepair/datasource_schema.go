package interfacepair

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InterfacepairDataSourceModel is the data-source-specific model, decoupled from
// InterfacepairResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type InterfacepairDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Interfaceid types.Int64  `tfsdk:"interface_id"` // Required lookup key
	Ifnum       types.List   `tfsdk:"ifnum"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/interfacepair.json). Never settable; populated from GET.
	Ifaces types.String `tfsdk:"ifaces"`
}

func InterfacepairDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"interface_id": schema.Int64Attribute{
				Required:    true,
				Description: "The Interface pair id",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The constituent interfaces in the interface pair",
			},

			// Read-only (GET-only) metadata surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed; null when the
			// appliance omits it.
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Names of all member interfaces of this Interface Pair.",
			},
		},
	}
}

// interfacepairDataSourceSetAttrFromGet projects a NITRO interfacepair GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func interfacepairDataSourceSetAttrFromGet(ctx context.Context, data *InterfacepairDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In interfacepairDataSourceSetAttrFromGet Function")

	// The NITRO key attribute is "id".
	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Interfaceid = utils.MapGetInt64(g, "id")
	} else {
		data.Interfaceid = types.Int64Null()
	}

	data.Ifnum = utils.MapGetStringList(g, "ifnum")

	// Read-only metadata.
	data.Ifaces = utils.MapGetString(g, "ifaces")
}
