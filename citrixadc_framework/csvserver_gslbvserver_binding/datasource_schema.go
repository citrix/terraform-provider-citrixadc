package csvserver_gslbvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverGslbvserverBindingDataSourceModel is the data-source-specific model,
// decoupled from CsvserverGslbvserverBindingResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes AND the read-only attributes that the
// resource deliberately omits.
type CsvserverGslbvserverBindingDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Vserver types.String `tfsdk:"vserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver_gslbvserver_binding.json). Never settable;
	// populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func CsvserverGslbvserverBindingDataSourceSchema() schema.Schema {
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

// csvserver_gslbvserver_bindingDataSourceSetAttrFromGet projects a NITRO
// csvserver_gslbvserver_binding GET response onto the data-source model via the
// shared utils.MapGet* helpers.
func csvserver_gslbvserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverGslbvserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserver_gslbvserver_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Vserver = utils.MapGetString(g, "vserver")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")

	// Set ID for the data source (composite key: name,vserver).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
