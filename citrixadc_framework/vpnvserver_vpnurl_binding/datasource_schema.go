package vpnvserver_vpnurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverVpnurlBindingDataSourceModel is the data-source-specific model,
// decoupled from VpnvserverVpnurlBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the lookup keys AND the read-only
// attributes that the resource deliberately omits. Every non-key attribute is
// Computed.
type VpnvserverVpnurlBindingDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`    // Required lookup key (parent)
	Urlname types.String `tfsdk:"urlname"` // Required lookup key (bound entity)

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_vpnurl_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverVpnurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"urlname": schema.StringAttribute{
				Required:    true,
				Description: "The intranet URL.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Type of action associated with the bound URL.",
			},
		},
	}
}

// vpnvserver_vpnurl_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_vpnurl_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func vpnvserver_vpnurl_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverVpnurlBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_vpnurl_bindingDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	if v, ok := g["urlname"]; ok && v != nil {
		data.Urlname = types.StringValue(utils.AnyToString(v))
	}

	// Read-only (GET-only) attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Set the composite ID (legacy key order: name,urlname).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("urlname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Urlname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
