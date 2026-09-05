package crvserver_lbvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrvserverLbvserverBindingDataSourceModel is the data-source-specific model,
// decoupled from CrvserverLbvserverBindingResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes AND the read-only attributes that
// the resource deliberately omits.
type CrvserverLbvserverBindingDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Lbvserver types.String `tfsdk:"lbvserver"`
	Name      types.String `tfsdk:"name"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/crvserver_lbvserver_binding.json). Never settable;
	// populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func CrvserverLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"lbvserver": schema.StringAttribute{
				Required:    true,
				Description: "The Default target server name.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
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

// crvserver_lbvserver_bindingDataSourceSetAttrFromGet projects a NITRO
// crvserver_lbvserver_binding GET response onto the data-source model via the
// shared utils.MapGet* helpers.
func crvserver_lbvserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CrvserverLbvserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In crvserver_lbvserver_bindingDataSourceSetAttrFromGet Function")

	data.Lbvserver = utils.MapGetString(g, "lbvserver")
	data.Name = utils.MapGetString(g, "name")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")

	// Set ID for the data source (composite key: lbvserver,name).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("lbvserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Lbvserver.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
