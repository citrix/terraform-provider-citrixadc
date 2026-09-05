package nstcpbufparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NstcpbufparamDataSourceModel is the data-source-specific model, decoupled from
// NstcpbufparamResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type NstcpbufparamDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Memlimit types.Int64  `tfsdk:"memlimit"`
	Size     types.Int64  `tfsdk:"size"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nstcpbufparam.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func NstcpbufparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, that can be used for buffering.",
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP buffering size per connection, in kilobytes.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if TCP buffering is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nstcpbufparamDataSourceSetAttrFromGet projects a NITRO nstcpbufparam GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nstcpbufparamDataSourceSetAttrFromGet(ctx context.Context, data *NstcpbufparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nstcpbufparamDataSourceSetAttrFromGet Function")

	// nstcpbufparam is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("nstcpbufparam-config")

	// Read/write attributes as read-back outputs.
	data.Memlimit = utils.MapGetInt64(g, "memlimit")
	data.Size = utils.MapGetInt64(g, "size")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
