package vpnvserver_sharefileserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverSharefileserverBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnvserverSharefileserverBindingDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`      // Required lookup key
	Sharefile types.String `tfsdk:"sharefile"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_sharefileserver_binding.json). Never
	// settable; populated from GET, Null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverSharefileserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"sharefile": schema.StringAttribute{
				Required:    true,
				Description: "Configured ShareFile server in XenMobile deployment. Format IP:PORT / FQDN:PORT",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
		},
	}
}

// vpnvserver_sharefileserver_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_sharefileserver_binding GET response onto the data-source model.
// Because a data source has no plan/apply reconciliation, attributes are simply
// filled from the GET (or left Null when the GET omits them). The shared
// utils.MapGet* helpers implement that projection.
func vpnvserver_sharefileserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverSharefileserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_sharefileserver_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Sharefile = utils.MapGetString(g, "sharefile")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Build the composite id using the legacy attribute order.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sharefile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sharefile.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
