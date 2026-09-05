package vpnvserver_intranetip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverIntranetip6BindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface (Read
// only; no plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes the
// resource deliberately omits (acttype). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type VpnvserverIntranetip6BindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Intranetip6 types.String `tfsdk:"intranetip6"`
	Name        types.String `tfsdk:"name"`
	Numaddr     types.Int64  `tfsdk:"numaddr"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_intranetip6_binding.json).
	// Never settable; populated from GET.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverIntranetip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"intranetip6": schema.StringAttribute{
				Required:    true,
				Description: "The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"numaddr": schema.Int64Attribute{
				// Optional lookup filter; returned by GET when omitted.
				Optional:    true,
				Computed:    true,
				Description: "The number of ipv6 addresses",
			},

			// Read-only (GET-only) attribute surfaced by the data source
			// (intentionally NOT modeled on the resource). Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Type of the bound action. Returned by the appliance on a GET; null when the appliance omits it.",
			},
		},
	}
}

// vpnvserver_intranetip6_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response element onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func vpnvserver_intranetip6_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverIntranetip6BindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_intranetip6_bindingDataSourceSetAttrFromGet Function")

	data.Intranetip6 = utils.MapGetString(g, "intranetip6")
	data.Name = utils.MapGetString(g, "name")
	data.Numaddr = utils.MapGetInt64(g, "numaddr")

	// Read-only GET-only attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set ID: multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	// (matching the SDK v2 / resource getter order: intranetip6, name, numaddr).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.UrlEncode(utils.AnyToString(data.Intranetip6.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(utils.AnyToString(data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("numaddr:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Numaddr.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
