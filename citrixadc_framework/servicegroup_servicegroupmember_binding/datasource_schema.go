package servicegroup_servicegroupmember_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServicegroupServicegroupmemberBindingDataSourceModel is the data-source-specific
// model. A data source is a pure read surface, so in addition to the read/write
// attributes (surfaced as Computed outputs) it exposes the read-only (GET-only)
// NITRO attributes that the resource intentionally omits.
type ServicegroupServicegroupmemberBindingDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	DisableRead      types.Bool   `tfsdk:"disable_read"`
	Aigwprofilename  types.String `tfsdk:"aigwprofilename"`
	Customserverid   types.String `tfsdk:"customserverid"`
	Dbsttl           types.Int64  `tfsdk:"dbsttl"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Ip               types.String `tfsdk:"ip"`
	Nameserver       types.String `tfsdk:"nameserver"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Serverid         types.Int64  `tfsdk:"serverid"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/servicegroup_servicegroupmember_binding.json). Never
	// settable; populated from GET, null when the appliance omits them.
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Trofsreason               types.String `tfsdk:"trofsreason"`
	Svrstate                  types.String `tfsdk:"svrstate"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Trofsdelay                types.Int64  `tfsdk:"trofsdelay"`
	Orderstr                  types.String `tfsdk:"orderstr"`
	Graceful                  types.String `tfsdk:"graceful"`
	Svcitmpriority            types.Int64  `tfsdk:"svcitmpriority"`
	Delay                     types.Int64  `tfsdk:"delay"`
}

func ServicegroupServicegroupmemberBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"disable_read": schema.BoolAttribute{
				Computed:    true,
				Description: "Skip reading the resource attributes from the NetScaler during refresh (resource-only flag; always reported false for the data source).",
			},
			"aigwprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backend AIGW Profile which will be attached to the servicegroup. This parameter enables the servicegroup to process the LLM request/response based on the profile config. Any service item bound to the servicegroup will inherit the backend AIGW Profile bound at the servicegroup level, if it does not have an explicit AIGW Profile given at bind time.",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
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
			"nameserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service group.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the service group.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
			"trofsreason": schema.StringAttribute{
				Computed:    true,
				Description: "Specify reason if service group member in TROFS.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the service.",
			},
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change occurred. Seconds part.",
			},
			"trofsdelay": schema.Int64Attribute{
				Computed:    true,
				Description: "Delay before moving to TROFS.",
			},
			"orderstr": schema.StringAttribute{
				Computed:    true,
				Description: "Order number in string form to be assigned to the servicegroup member.",
			},
			"graceful": schema.StringAttribute{
				Computed:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"svcitmpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "This gives the priority of the FQDN service items for SRV server binding.",
			},
			"delay": schema.Int64Attribute{
				Computed:    true,
				Description: "Time, in seconds, allocated for a shutdown of the services in the service group.",
			},
		},
	}
}

// servicegroup_servicegroupmember_bindingDataSourceSetAttrFromGet projects a NITRO
// servicegroup_servicegroupmember_binding GET response onto the data-source model
// and sets the composite ID. Attributes are filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func servicegroup_servicegroupmember_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ServicegroupServicegroupmemberBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In servicegroup_servicegroupmember_bindingDataSourceSetAttrFromGet Function")

	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.Customserverid = utils.MapGetString(g, "customserverid")
	data.Dbsttl = utils.MapGetInt64(g, "dbsttl")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Ip = utils.MapGetString(g, "ip")
	data.Nameserver = utils.MapGetString(g, "nameserver")
	data.Order = utils.MapGetInt64(g, "order")
	data.Port = utils.MapGetInt64(g, "port")
	data.Serverid = utils.MapGetInt64(g, "serverid")
	data.Servername = utils.MapGetString(g, "servername")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.State = utils.MapGetString(g, "state")
	data.Weight = utils.MapGetInt64(g, "weight")

	// disable_read is a resource-only convenience flag with no datasource meaning;
	// give it a known value so the model is fully populated.
	data.DisableRead = types.BoolValue(false)

	// Read-only (GET-only) attributes.
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Trofsreason = utils.MapGetString(g, "trofsreason")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Trofsdelay = utils.MapGetInt64(g, "trofsdelay")
	data.Orderstr = utils.MapGetString(g, "orderstr")
	data.Graceful = utils.MapGetString(g, "graceful")
	data.Svcitmpriority = utils.MapGetInt64(g, "svcitmpriority")
	data.Delay = utils.MapGetInt64(g, "delay")

	// Datasource has no Create — set the composite ID here.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
