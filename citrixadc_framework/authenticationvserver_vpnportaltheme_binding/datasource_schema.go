package authenticationvserver_vpnportaltheme_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationvserverVpnportalthemeBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source
// is a pure read surface, so it exposes the full GET projection: the lookup
// keys AND the read-only attributes the resource deliberately omits.
type AuthenticationvserverVpnportalthemeBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`        // Required lookup key
	Portaltheme types.String `tfsdk:"portaltheme"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationvserver_vpnportaltheme_binding.json).
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AuthenticationvserverVpnportalthemeBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"portaltheme": schema.StringAttribute{
				Required:    true,
				Description: "Theme for Authentication virtual server Login portal",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Type of the bound portal theme action, as reported by the appliance.",
			},
		},
	}
}

// authenticationvserver_vpnportaltheme_bindingDataSourceSetAttrFromGet projects a
// NITRO GET response onto the data-source model. Attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers.
func authenticationvserver_vpnportaltheme_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationvserverVpnportalthemeBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationvserver_vpnportaltheme_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Portaltheme = utils.MapGetString(g, "portaltheme")

	// Read-only (GET-only) attributes.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Compose the ID for the datasource (no Create step). ID = name,portaltheme
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("portaltheme:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Portaltheme.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
