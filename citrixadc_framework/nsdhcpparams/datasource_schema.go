package nsdhcpparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsdhcpparamsDataSourceModel is the data-source-specific model, decoupled from
// NsdhcpparamsResourceModel. nsdhcpparams is a singleton, so this data source
// exposes the read/write attributes (as Computed outputs) AND the read-only
// DHCP-acquired metadata the resource deliberately omits (ipaddress, netmask,
// hostrtgw, running).
type NsdhcpparamsDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Dhcpclient      types.String `tfsdk:"dhcpclient"`
	Saveroute       types.String `tfsdk:"saveroute"`
	Subnetselection types.String `tfsdk:"subnetselection"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsdhcpparams.json). Never settable; populated from GET.
	Ipaddress types.String `tfsdk:"ipaddress"`
	Netmask   types.String `tfsdk:"netmask"`
	Hostrtgw  types.String `tfsdk:"hostrtgw"`
	Running   types.Bool   `tfsdk:"running"`
}

func NsdhcpparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dhcpclient": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot.",
			},
			"saveroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DHCP acquired routes are saved on the Citrix ADC.",
			},
			"subnetselection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet Selection option (RFC 3011) to request IP from a specific subnet.",
			},

			// Read-only (GET-only) DHCP-acquired metadata surfaced by the data source.
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "DHCP acquired IP.",
			},
			"netmask": schema.StringAttribute{
				Computed:    true,
				Description: "DHCP acquired Netmask.",
			},
			"hostrtgw": schema.StringAttribute{
				Computed:    true,
				Description: "DHCP acquired Gateway.",
			},
			"running": schema.BoolAttribute{
				Computed:    true,
				Description: "DHCP mode.",
			},
		},
	}
}

// nsdhcpparamsDataSourceSetAttrFromGet projects a NITRO nsdhcpparams GET response
// onto the data-source model. nsdhcpparams is a singleton, so the ID is static.
// Attributes are filled from the GET (or left Null when the GET omits them) via
// the shared utils.MapGet* helpers.
func nsdhcpparamsDataSourceSetAttrFromGet(ctx context.Context, data *NsdhcpparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsdhcpparamsDataSourceSetAttrFromGet Function")

	// Singleton resource - static ID.
	data.Id = types.StringValue("nsdhcpparams-config")

	data.Dhcpclient = utils.MapGetString(g, "dhcpclient")
	data.Saveroute = utils.MapGetString(g, "saveroute")
	data.Subnetselection = utils.MapGetString(g, "subnetselection")

	// Read-only DHCP-acquired metadata.
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Hostrtgw = utils.MapGetString(g, "hostrtgw")
	data.Running = utils.MapGetBool(g, "running")
}
