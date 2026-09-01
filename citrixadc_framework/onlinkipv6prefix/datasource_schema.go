package onlinkipv6prefix

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Onlinkipv6prefixDataSourceModel is the data-source-specific model, decoupled
// from Onlinkipv6prefixResourceModel. A data source is a pure read surface, so it
// exposes the full GET projection: the read/write attributes (as Computed outputs)
// plus the read-only attributes the resource deliberately omits.
type Onlinkipv6prefixDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Autonomusprefix          types.String `tfsdk:"autonomusprefix"`
	Decrementprefixlifetimes types.String `tfsdk:"decrementprefixlifetimes"`
	Depricateprefix          types.String `tfsdk:"depricateprefix"`
	Ipv6prefix               types.String `tfsdk:"ipv6prefix"`
	Onlinkprefix             types.String `tfsdk:"onlinkprefix"`
	Prefixpreferredlifetime  types.Int64  `tfsdk:"prefixpreferredlifetime"`
	Prefixvalidelifetime     types.Int64  `tfsdk:"prefixvalidelifetime"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/onlinkipv6prefix.json). Never settable; populated from GET.
	Prefixcurrvalidelft    types.Int64 `tfsdk:"prefixcurrvalidelft"`
	Prefixcurrpreferredlft types.Int64 `tfsdk:"prefixcurrpreferredlft"`
}

func Onlinkipv6prefixDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"autonomusprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix Autonomus flag.",
			},
			"decrementprefixlifetimes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix Autonomus flag.",
			},
			"depricateprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Depricate the prefix.",
			},
			"ipv6prefix": schema.StringAttribute{
				Required:    true,
				Description: "Onlink prefixes for RA messages.",
			},
			"onlinkprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RA Prefix onlink flag.",
			},
			"prefixpreferredlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preferred life time of the prefix, in seconds.",
			},
			"prefixvalidelifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Valide life time of the prefix, in seconds.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"prefixcurrvalidelft": schema.Int64Attribute{
				Computed:    true,
				Description: "Prefix current valid life time.",
			},
			"prefixcurrpreferredlft": schema.Int64Attribute{
				Computed:    true,
				Description: "Prefix current prefered life time.",
			},
		},
	}
}

// onlinkipv6prefixDataSourceSetAttrFromGet projects a NITRO onlinkipv6prefix GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func onlinkipv6prefixDataSourceSetAttrFromGet(ctx context.Context, data *Onlinkipv6prefixDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In onlinkipv6prefixDataSourceSetAttrFromGet Function")

	if v, ok := g["ipv6prefix"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Ipv6prefix = types.StringValue(utils.AnyToString(v))
	}

	data.Autonomusprefix = utils.MapGetString(g, "autonomusprefix")
	data.Decrementprefixlifetimes = utils.MapGetString(g, "decrementprefixlifetimes")
	data.Depricateprefix = utils.MapGetString(g, "depricateprefix")
	data.Onlinkprefix = utils.MapGetString(g, "onlinkprefix")
	data.Prefixpreferredlifetime = utils.MapGetInt64(g, "prefixpreferredlifetime")
	data.Prefixvalidelifetime = utils.MapGetInt64(g, "prefixvalidelifetime")

	// Read-only attributes.
	data.Prefixcurrvalidelft = utils.MapGetInt64(g, "prefixcurrvalidelft")
	data.Prefixcurrpreferredlft = utils.MapGetInt64(g, "prefixcurrpreferredlft")
}
