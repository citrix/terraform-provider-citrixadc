package routerdynamicrouting

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RouterdynamicroutingDataSourceModel describes the datasource data model.
//
// The datasource is a read-only "show command" query and is intentionally
// decoupled from the resource model: the resource is action-only and carries a
// commandlines list, while the datasource queries by a single commandstring. It
// additionally exposes the read-only "output" field returned by the appliance.
type RouterdynamicroutingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Commandstring types.String `tfsdk:"commandstring"` // Required lookup key
	Nodeid        types.Int64  `tfsdk:"nodeid"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/routerdynamicrouting.json). Never settable; populated from GET.
	Output types.String `tfsdk:"output"`
}

func RouterdynamicroutingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"commandstring": schema.StringAttribute{
				Required:    true,
				Description: "command to be executed",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). Computed.
			"output": schema.StringAttribute{
				Computed:    true,
				Description: "command output.",
			},
		},
	}
}

// routerdynamicroutingDataSourceSetAttrFromGet projects a NITRO
// routerdynamicrouting GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func routerdynamicroutingDataSourceSetAttrFromGet(ctx context.Context, data *RouterdynamicroutingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In routerdynamicroutingDataSourceSetAttrFromGet Function")

	data.Commandstring = utils.MapGetString(g, "commandstring")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")

	// Read-only metadata.
	data.Output = utils.MapGetString(g, "output")

	// Set ID for the datasource.
	data.Id = types.StringValue(data.Commandstring.ValueString())
}
