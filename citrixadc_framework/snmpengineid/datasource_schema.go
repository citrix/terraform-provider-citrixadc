package snmpengineid

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpengineidDataSourceModel is the data-source-specific model, decoupled from
// SnmpengineidResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type SnmpengineidDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Engineid  types.String `tfsdk:"engineid"`
	Ownernode types.Int64  `tfsdk:"ownernode"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/snmpengineid.json). Never settable; populated from GET.
	Defaultengineid types.String `tfsdk:"defaultengineid"`
}

func SnmpengineidDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"engineid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A hexadecimal value of at least 10 characters, uniquely identifying the engineid",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which you are setting the engineid",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"defaultengineid": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier to assign to the SNMPv3 engine. Should be a hexadecimal value with a minimum length of 10 hex characters.",
			},
		},
	}
}

// snmpengineidDataSourceSetAttrFromGet projects a NITRO snmpengineid GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func snmpengineidDataSourceSetAttrFromGet(ctx context.Context, data *SnmpengineidDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpengineidDataSourceSetAttrFromGet Function")

	data.Engineid = utils.MapGetString(g, "engineid")

	// ownernode is the Required lookup key supplied via config; only overwrite it
	// when the GET echoes it so a config value the appliance omits is preserved.
	if v, ok := g["ownernode"]; ok && v != nil {
		data.Ownernode = utils.MapGetInt64(g, "ownernode")
	}

	// Read-only (GET-only) metadata.
	data.Defaultengineid = utils.MapGetString(g, "defaultengineid")

	// Singleton resource - use a static ID (matches the resource getter).
	data.Id = types.StringValue("snmpengineid-config")
}
