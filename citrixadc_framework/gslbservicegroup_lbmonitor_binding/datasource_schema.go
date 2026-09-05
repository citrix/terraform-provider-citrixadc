package gslbservicegroup_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbservicegroupLbmonitorBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read
// surface, so it can expose the FULL GET projection: the read/write attributes
// (as Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type GslbservicegroupLbmonitorBindingDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	MonitorName      types.String `tfsdk:"monitor_name"`
	Monstate         types.String `tfsdk:"monstate"`
	Order            types.Int64  `tfsdk:"order"`
	Passive          types.Bool   `tfsdk:"passive"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbservicegroup_lbmonitor_binding.json). Never settable;
	// populated from GET; null when the appliance omits them.
	Monweight types.Int64 `tfsdk:"monweight"`
}

func GslbservicegroupLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique numerical identifier used by hash based load balancing methods to identify a service.",
			},
			"monitor_name": schema.StringAttribute{
				Required:    true,
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor state.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the GSLB service. Each service must have a unique port number.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service group.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the service after binding.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"monweight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight of the monitor that is bound to the GSLB service group.",
			},
		},
	}
}

// gslbservicegroup_lbmonitor_bindingDataSourceSetAttrFromGet projects a NITRO
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers. The
// composite ID is built exactly as the resource Create emits it.
func gslbservicegroup_lbmonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *GslbservicegroupLbmonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbservicegroup_lbmonitor_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.MonitorName = utils.MapGetString(g, "monitor_name")
	data.Monstate = utils.MapGetString(g, "monstate")
	data.Order = utils.MapGetInt64(g, "order")
	data.Passive = utils.MapGetBool(g, "passive")
	data.Port = utils.MapGetInt64(g, "port")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Publicport = utils.MapGetInt64(g, "publicport")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.Siteprefix = utils.MapGetString(g, "siteprefix")
	data.State = utils.MapGetString(g, "state")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) attributes.
	data.Monweight = utils.MapGetInt64(g, "monweight")

	// Composite ID: servicegroupname,monitor_name (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
