package vpnvserver_appcontroller_binding

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverAppcontrollerBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. It exposes the read/write attributes (as
// Computed outputs) AND the read-only attributes the appliance returns on a GET
// (zion73x_readonly/vpnvserver_appcontroller_binding.json) that the resource omits.
type VpnvserverAppcontrollerBindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Appcontroller types.String `tfsdk:"appcontroller"` // Required lookup key
	Name          types.String `tfsdk:"name"`          // Required lookup key (parent)

	// Read-only (GET-only) attributes surfaced only by the data source.
	Acttype types.Int64 `tfsdk:"acttype"`
}

func VpnvserverAppcontrollerBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appcontroller": schema.StringAttribute{
				Required:    true,
				Description: "Configured App Controller server in XenMobile deployment.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},

			// Read-only (GET-only) attributes exposed only by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "The bound entity (action) type, as returned by the appliance. GET-only; null when the appliance omits it.",
			},
		},
	}
}

// vpnvserver_appcontroller_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model via the shared utils.MapGet* helpers. A data
// source has no plan/apply reconciliation, so attributes are simply filled from the
// GET (or left Null when the GET omits them).
func vpnvserver_appcontroller_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverAppcontrollerBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_appcontroller_bindingDataSourceSetAttrFromGet Function")

	data.Appcontroller = utils.MapGetString(g, "appcontroller")
	data.Name = utils.MapGetString(g, "name")

	// Read-only attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID (legacy SDK v2 order: appcontroller,name).
	data.Id = types.StringValue(fmt.Sprintf("appcontroller:%s,name:%s", utils.UrlEncode(data.Appcontroller.ValueString()), utils.UrlEncode(data.Name.ValueString())))
}
