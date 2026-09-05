package gslbservicegroup_gslbservicegroupmember_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbservicegroupGslbservicegroupmemberBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model. A data source
// is a pure read surface, so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits. Every non-key attribute is Computed.
type GslbservicegroupGslbservicegroupmemberBindingDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Ip               types.String `tfsdk:"ip"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbservicegroup_gslbservicegroupmember_binding.json).
	// Never settable; populated from GET; null when the appliance omits them.
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Preferredlocation         types.String `tfsdk:"preferredlocation"`
	Trofsdelay                types.Int64  `tfsdk:"trofsdelay"`
	Delay                     types.Int64  `tfsdk:"delay"`
	Gslbthreshold             types.Int64  `tfsdk:"gslbthreshold"`
	Orderstr                  types.String `tfsdk:"orderstr"`
	Graceful                  types.String `tfsdk:"graceful"`
	Threshold                 types.String `tfsdk:"threshold"`
	Svrstate                  types.String `tfsdk:"svrstate"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
}

func GslbservicegroupGslbservicegroupmemberBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Address.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Required:    true,
				Description: "Server port number.",
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
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
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
				Description: "Initial state of the GSLB service group.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change occurred. Seconds part.",
			},
			"preferredlocation": schema.StringAttribute{
				Computed:    true,
				Description: "Prefered location.",
			},
			"trofsdelay": schema.Int64Attribute{
				Computed:    true,
				Description: "Delay before moving to TROFS.",
			},
			"delay": schema.Int64Attribute{
				Computed:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown.",
			},
			"gslbthreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates if gslb svc has reached threshold.",
			},
			"orderstr": schema.StringAttribute{
				Computed:    true,
				Description: "Order number in string form assigned to the gslb servicegroup member.",
			},
			"graceful": schema.StringAttribute{
				Computed:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"threshold": schema.StringAttribute{
				Computed:    true,
				Description: "Threshold indicator. Possible values = ABOVE, BELOW.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the GSLB service.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
		},
	}
}

// gslbservicegroup_gslbservicegroupmember_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers. The composite ID is built exactly as the resource Create emits it.
func gslbservicegroup_gslbservicegroupmember_bindingDataSourceSetAttrFromGet(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbservicegroup_gslbservicegroupmember_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Ip = utils.MapGetString(g, "ip")
	data.Order = utils.MapGetInt64(g, "order")
	data.Port = utils.MapGetInt64(g, "port")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Publicport = utils.MapGetInt64(g, "publicport")
	data.Servername = utils.MapGetString(g, "servername")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.Siteprefix = utils.MapGetString(g, "siteprefix")
	data.State = utils.MapGetString(g, "state")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) attributes.
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")
	data.Trofsdelay = utils.MapGetInt64(g, "trofsdelay")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Gslbthreshold = utils.MapGetInt64(g, "gslbthreshold")
	data.Orderstr = utils.MapGetString(g, "orderstr")
	data.Graceful = utils.MapGetString(g, "graceful")
	data.Threshold = utils.MapGetString(g, "threshold")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")

	// Composite ID (comma-separated key:UrlEncode(value) pairs), same order Create emits.
	data.Id = types.StringValue(gslbservicegroup_gslbservicegroupmember_bindingBuildId(
		data.Servicegroupname.ValueString(),
		data.Servername.ValueString(),
		data.Ip.ValueString(),
		data.Port.ValueInt64(),
	))
}
