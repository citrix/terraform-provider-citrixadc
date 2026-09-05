package vpnvserver_vpnnexthopserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverVpnnexthopserverBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the FULL GET projection:
// the lookup keys plus the read-only attributes the resource deliberately omits.
// Every non-key attribute is Computed.
type VpnvserverVpnnexthopserverBindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`          // Required lookup key
	Nexthopserver types.String `tfsdk:"nexthopserver"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_vpnnexthopserver_binding.json). Never
	// settable; populated from GET.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverVpnnexthopserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"nexthopserver": schema.StringAttribute{
				Required:    true,
				Description: "The name of the next hop server bound to the VPN virtual server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (GET-only).",
			},
		},
	}
}

// vpnvserver_vpnnexthopserver_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_vpnnexthopserver_binding GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers.
func vpnvserver_vpnnexthopserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverVpnnexthopserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_vpnnexthopserver_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Nexthopserver = utils.MapGetString(g, "nexthopserver")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID: comma-separated key:UrlEncode(value) pairs, matching the resource.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(utils.AnyToString(data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("nexthopserver:%s", utils.UrlEncode(utils.AnyToString(data.Nexthopserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
