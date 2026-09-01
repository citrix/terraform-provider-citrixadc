package extendedmemoryparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ExtendedmemoryparamDataSourceModel is the data-source-specific model, decoupled
// from ExtendedmemoryparamResourceModel. A data source is a pure read surface, so
// it can expose the FULL GET projection: the read/write attribute (as a Computed
// output) AND the read-only attributes the resource deliberately omits
// (memlimitactive, maxmemlimit, minrequiredmemory).
type ExtendedmemoryparamDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Memlimit types.Int64  `tfsdk:"memlimit"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/extendedmemoryparam.json). Never settable; populated from GET.
	Memlimitactive    types.Int64 `tfsdk:"memlimitactive"`
	Maxmemlimit       types.Int64 `tfsdk:"maxmemlimit"`
	Minrequiredmemory types.Int64 `tfsdk:"minrequiredmemory"`
}

func ExtendedmemoryparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of NetScaler memory to reserve for the memory used by LSN and Subscriber Session Store feature, in multiples of 2MB.\n\nNote: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"memlimitactive": schema.Int64Attribute{
				Computed:    true,
				Description: "The active memory limit for extendedmemory on the system. Active memory limit could be different from configured memory limit. This could happen when memory limit could not be increased due to unavailability, or could not be decreased as it is already in use. This active memory limit configures the current memory limit for LSN and Subscriber Session Store.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Computed:    true,
				Description: "The maximum value of memory limit for extendedmemory on the system. Actual available memory may be less. This is maximum memory that can be utilized by LSN and Subscriber Session Store modules.",
			},
			"minrequiredmemory": schema.Int64Attribute{
				Computed:    true,
				Description: "The minimum memory requirement for extendedmemory. This is minimum memory required for LSN and Subscriber Session Store Modules.",
			},
		},
	}
}

// extendedmemoryparamDataSourceSetAttrFromGet projects a NITRO extendedmemoryparam
// GET response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func extendedmemoryparamDataSourceSetAttrFromGet(ctx context.Context, data *ExtendedmemoryparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In extendedmemoryparamDataSourceSetAttrFromGet Function")

	// Singleton resource - static ID.
	data.Id = types.StringValue("extendedmemoryparam-config")

	data.Memlimit = utils.MapGetInt64(g, "memlimit")

	// Read-only metadata.
	data.Memlimitactive = utils.MapGetInt64(g, "memlimitactive")
	data.Maxmemlimit = utils.MapGetInt64(g, "maxmemlimit")
	data.Minrequiredmemory = utils.MapGetInt64(g, "minrequiredmemory")
}
