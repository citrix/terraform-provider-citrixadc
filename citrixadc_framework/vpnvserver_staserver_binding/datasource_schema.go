package vpnvserver_staserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverStaserverBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type VpnvserverStaserverBindingDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"` // Required lookup key
	Staaddresstype types.String `tfsdk:"staaddresstype"`
	Staserver      types.String `tfsdk:"staserver"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver_staserver_binding.json). Never settable;
	// populated from GET, Null when the appliance omits them.
	Stastate  types.String `tfsdk:"stastate"`
	Staauthid types.String `tfsdk:"staauthid"`
	Acttype   types.Int64  `tfsdk:"acttype"`
}

func VpnvserverStaserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"stastate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the STA Server. If Authority ID is set then STA Server is UP else DOWN. Possible values: [ UP, DOWN ]",
			},
			"staauthid": schema.StringAttribute{
				Computed:    true,
				Description: "Authority ID of the STA Server. Authority ID is used to match incoming STA tickets in the SOCKS/CGP protocol with the right STA server.",
			},
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding (read-only).",
			},
		},
	}
}

// vpnvserver_staserver_bindingDataSourceSetAttrFromGet projects a NITRO
// vpnvserver_staserver_binding GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func vpnvserver_staserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverStaserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserver_staserver_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Staaddresstype = utils.MapGetString(g, "staaddresstype")
	data.Staserver = utils.MapGetString(g, "staserver")

	// Read-only attributes.
	data.Stastate = utils.MapGetString(g, "stastate")
	data.Staauthid = utils.MapGetString(g, "staauthid")
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Build the composite id using the legacy attribute order.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("staserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Staserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
