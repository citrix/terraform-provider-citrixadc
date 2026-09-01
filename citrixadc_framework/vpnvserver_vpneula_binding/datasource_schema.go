package vpnvserver_vpneula_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverVpneulaBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the FULL GET
// projection: the lookup keys plus the read-only attributes the resource
// deliberately omits. Every non-key attribute is Computed.
type VpnvserverVpneulaBindingDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Eula types.String `tfsdk:"eula"` // Required lookup key
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_vpneula_binding.json). Never settable;
	// populated from GET.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverVpneulaBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"eula": schema.StringAttribute{
				Required:    true,
				Description: "Name of the EULA bound to VPN vserver",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (GET-only).",
			},
		},
	}
}

// vpnvserver_vpneula_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_vpneula_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func vpnvserver_vpneula_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverVpneulaBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_vpneula_bindingDataSourceSetAttrFromGet Function")

	data.Eula = utils.MapGetString(g, "eula")
	data.Name = utils.MapGetString(g, "name")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID: comma-separated key:UrlEncode(value) pairs, matching the resource.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("eula:%s", utils.UrlEncode(utils.AnyToString(data.Eula.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(utils.AnyToString(data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
