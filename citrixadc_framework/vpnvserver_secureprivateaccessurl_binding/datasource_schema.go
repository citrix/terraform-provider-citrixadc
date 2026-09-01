package vpnvserver_secureprivateaccessurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverSecureprivateaccessurlBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnvserverSecureprivateaccessurlBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`                   // Required lookup key
	Secureprivateaccessurl types.String `tfsdk:"secureprivateaccessurl"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_secureprivateaccessurl_binding.json). Never
	// settable; populated from GET, Null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverSecureprivateaccessurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"secureprivateaccessurl": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Private Access URL",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
		},
	}
}

// vpnvserver_secureprivateaccessurl_bindingDataSourceSetAttrFromGet projects a
// NITRO vpnvserver_secureprivateaccessurl_binding GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func vpnvserver_secureprivateaccessurl_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverSecureprivateaccessurlBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_secureprivateaccessurl_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Secureprivateaccessurl = utils.MapGetString(g, "secureprivateaccessurl")

	// Read-only attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Build the composite id using the legacy attribute order.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
