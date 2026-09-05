package nsweblogparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsweblogparamDataSourceModel is the data-source-specific model, decoupled from
// NsweblogparamResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) plus
// the read-only attributes the resource deliberately omits.
type NsweblogparamDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Buffersizemb  types.Int64  `tfsdk:"buffersizemb"`
	Customreqhdrs types.List   `tfsdk:"customreqhdrs"`
	Customrsphdrs types.List   `tfsdk:"customrsphdrs"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsweblogparam.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func NsweblogparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"buffersizemb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system.",
			},
			"customreqhdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP request headers whose values should be exported by the Web Logging feature.",
			},
			"customrsphdrs": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of HTTP response headers whose values should be exported by the Web Logging feature.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"builtin": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Flag to determine if log bufs is built-in or not (MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nsweblogparamDataSourceSetAttrFromGet projects a NITRO nsweblogparam GET response
// onto the data-source model. Attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func nsweblogparamDataSourceSetAttrFromGet(ctx context.Context, data *NsweblogparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsweblogparamDataSourceSetAttrFromGet Function")

	data.Buffersizemb = utils.MapGetInt64(g, "buffersizemb")
	data.Customreqhdrs = utils.MapGetStringList(g, "customreqhdrs")
	data.Customrsphdrs = utils.MapGetStringList(g, "customrsphdrs")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Singleton resource - static ID.
	data.Id = types.StringValue("nsweblogparam-config")
}
