package servicegroup_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServicegroupLbmonitorBindingDataSourceModel describes the datasource data model.
// The datasource exposes the monitor name under the "monitor_name" attribute (the
// NITRO wire name) while the resource keeps the legacy SDK v2 "monitorname" attribute,
// so the two need separate models even though they share most fields.
//
// A data source is a pure read surface, so it additionally exposes the read-only
// (GET-only) NITRO attributes that the resource intentionally omits.
type ServicegroupLbmonitorBindingDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Aigwprofilename  types.String `tfsdk:"aigwprofilename"`
	Customserverid   types.String `tfsdk:"customserverid"`
	Dbsttl           types.Int64  `tfsdk:"dbsttl"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	MonitorName      types.String `tfsdk:"monitor_name"`
	Monstate         types.String `tfsdk:"monstate"`
	Nameserver       types.String `tfsdk:"nameserver"`
	Order            types.Int64  `tfsdk:"order"`
	Passive          types.Bool   `tfsdk:"passive"`
	Port             types.Int64  `tfsdk:"port"`
	Serverid         types.Int64  `tfsdk:"serverid"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	State            types.String `tfsdk:"state"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/servicegroup_lbmonitor_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Monweight types.Int64 `tfsdk:"monweight"`
}

func ServicegroupLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aigwprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backend AIGW Profile which will be attached to the servicegroup. This parameter enables the servicegroup to process the LLM request/response based on the profile config. Any service item bound to the servicegroup will inherit the backend AIGW Profile bound at the servicegroup level, if it does not have an explicit AIGW Profile given at bind time.",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
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
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the service. Each service must have a unique port number.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service group.",
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

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"monweight": schema.Int64Attribute{
				Computed:    true,
				Description: "weight of the monitor that is bound to servicegroup.",
			},
		},
	}
}

// servicegroup_lbmonitor_bindingDataSourceSetAttrFromGet projects a NITRO
// servicegroup_lbmonitor_binding GET response onto the data-source model and sets
// the composite ID. Attributes are filled from the GET (or left Null when the GET
// omits them) via the shared utils.MapGet* helpers.
func servicegroup_lbmonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ServicegroupLbmonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In servicegroup_lbmonitor_bindingDataSourceSetAttrFromGet Function")

	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.Customserverid = utils.MapGetString(g, "customserverid")
	data.Dbsttl = utils.MapGetInt64(g, "dbsttl")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.MonitorName = utils.MapGetString(g, "monitor_name")
	data.Monstate = utils.MapGetString(g, "monstate")
	data.Nameserver = utils.MapGetString(g, "nameserver")
	data.Order = utils.MapGetInt64(g, "order")
	data.Passive = utils.MapGetBool(g, "passive")
	data.Port = utils.MapGetInt64(g, "port")
	data.Serverid = utils.MapGetInt64(g, "serverid")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.State = utils.MapGetString(g, "state")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) attributes.
	data.Monweight = utils.MapGetInt64(g, "monweight")

	// Datasource has no Create — set the composite ID here.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
