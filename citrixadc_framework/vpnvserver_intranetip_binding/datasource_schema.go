package vpnvserver_intranetip_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverIntranetipBindingDataSourceModel is the data-source-specific model,
// decoupled from VpnvserverIntranetipBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnvserverIntranetipBindingDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Intranetip types.String `tfsdk:"intranetip"` // Required lookup key
	Name       types.String `tfsdk:"name"`       // Required lookup key
	Netmask    types.String `tfsdk:"netmask"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_intranetip_binding.json). Never settable;
	// populated from GET, Null when the appliance omits them.
	Acttype types.Int64  `tfsdk:"acttype"`
	Map     types.String `tfsdk:"map"`
}

func VpnvserverIntranetipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"intranetip": schema.StringAttribute{
				Required:    true,
				Description: "The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the intranet IP address or range.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
			"map": schema.StringAttribute{
				Computed:    true,
				Description: "Whether or not mapped IP addresses are ON or OFF. Mapped IP addresses are source IP addresses for the virtual servers running on the Citrix ADC. Possible values: [ ON, OFF ]",
			},
		},
	}
}

// vpnvserver_intranetip_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_intranetip_binding GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func vpnvserver_intranetip_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverIntranetipBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingDataSourceSetAttrFromGet Function")

	data.Intranetip = utils.MapGetString(g, "intranetip")
	data.Name = utils.MapGetString(g, "name")
	data.Netmask = utils.MapGetString(g, "netmask")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")
	data.Map = utils.MapGetString(g, "map")

	// Build the composite id using the legacy attribute order (name, intranetip)
	// matching resource_id_mapping.json.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
