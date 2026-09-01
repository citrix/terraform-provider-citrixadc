package nstrafficdomain

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NstrafficdomainDataSourceModel is the data-source-specific model, decoupled
// from NstrafficdomainResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type NstrafficdomainDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Aliasname types.String `tfsdk:"aliasname"`
	Td        types.Int64  `tfsdk:"td"` // Required lookup key
	Vmac      types.String `tfsdk:"vmac"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nstrafficdomain.json). Never settable; populated from GET.
	State types.String `tfsdk:"state"`
}

func NstrafficdomainDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aliasname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of traffic domain  being added.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Associate the traffic domain with a VMAC address instead of with VLANs. The Citrix ADC then sends the VMAC address of the traffic domain in all responses to ARP queries for network entities in that domain. As a result, the ADC can segregate subsequent incoming traffic for this traffic domain on the basis of the destination MAC address, because the destination MAC address is the VMAC address of the traffic domain. After creating entities on a traffic domain, you can easily manage and monitor them by performing traffic domain level operations.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the traffic domain.",
			},
		},
	}
}

// nstrafficdomainDataSourceSetAttrFromGet projects a NITRO nstrafficdomain GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nstrafficdomainDataSourceSetAttrFromGet(ctx context.Context, data *NstrafficdomainDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nstrafficdomainDataSourceSetAttrFromGet Function")

	if v, ok := g["td"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		// td is a config-supplied key; NITRO omits it for the default traffic
		// domain (0), so preserve the configured value instead of nulling it.
		if tdv, tdok := g["td"]; tdok && tdv != nil {
			if iv, err := utils.ConvertToInt64(tdv); err == nil {
				data.Td = types.Int64Value(iv)
			}
		}
	}

	// Read/write attributes as read-back outputs.
	data.Aliasname = utils.MapGetString(g, "aliasname")
	data.Vmac = utils.MapGetString(g, "vmac")

	// Read-only attributes.
	data.State = utils.MapGetString(g, "state")
}
