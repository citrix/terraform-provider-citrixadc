package csvserver_vpnvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverVpnvserverBindingDataSourceModel is the data-source-specific model,
// decoupled from CsvserverVpnvserverBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares.
type CsvserverVpnvserverBindingDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver_vpnvserver_binding.json). Never settable;
	// populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func CsvserverVpnvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"vserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// csvserver_vpnvserver_bindingDataSourceSetAttrFromGet projects a NITRO
// csvserver_vpnvserver_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func csvserver_vpnvserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverVpnvserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserver_vpnvserver_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Name = utils.MapGetString(g, "name")
	data.Vserver = utils.MapGetString(g, "vserver")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")

	// Composite binding ID: comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.UrlEncode(data.Vserver.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
