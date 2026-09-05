package servicegroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServicegroupDataSourceModel is the data-source-specific model, decoupled from
// ServicegroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/state metadata that the resource
// deliberately omits (svrstate, numofconnections, effective state, monitor probe
// counters, ...). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares.
type ServicegroupDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Aigwprofilename      types.String `tfsdk:"aigwprofilename"`
	Appflowlog           types.String `tfsdk:"appflowlog"`
	Autodelayedtrofs     types.String `tfsdk:"autodelayedtrofs"`
	Autodisabledelay     types.Int64  `tfsdk:"autodisabledelay"`
	Autodisablegraceful  types.String `tfsdk:"autodisablegraceful"`
	Autoscale            types.String `tfsdk:"autoscale"`
	Bootstrap            types.String `tfsdk:"bootstrap"`
	Cacheable            types.String `tfsdk:"cacheable"`
	Cachetype            types.String `tfsdk:"cachetype"`
	Cip                  types.String `tfsdk:"cip"`
	Cipheader            types.String `tfsdk:"cipheader"`
	Cka                  types.String `tfsdk:"cka"`
	Clttimeout           types.Int64  `tfsdk:"clttimeout"`
	Cmp                  types.String `tfsdk:"cmp"`
	Comment              types.String `tfsdk:"comment"`
	Customserverid       types.String `tfsdk:"customserverid"`
	Dbsttl               types.Int64  `tfsdk:"dbsttl"`
	Delay                types.Int64  `tfsdk:"delay"`
	Downstateflush       types.String `tfsdk:"downstateflush"`
	Dupweight            types.Int64  `tfsdk:"dupweight"`
	Graceful             types.String `tfsdk:"graceful"`
	Hashid               types.Int64  `tfsdk:"hashid"`
	Healthmonitor        types.String `tfsdk:"healthmonitor"`
	Httpprofilename      types.String `tfsdk:"httpprofilename"`
	Includemembers       types.Bool   `tfsdk:"includemembers"`
	Maxbandwidth         types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient            types.Int64  `tfsdk:"maxclient"`
	Maxreq               types.Int64  `tfsdk:"maxreq"`
	Mcpprofilename       types.String `tfsdk:"mcpprofilename"`
	Memberport           types.Int64  `tfsdk:"memberport"`
	Monconnectionclose   types.String `tfsdk:"monconnectionclose"`
	Monitornamesvc       types.String `tfsdk:"monitornamesvc"`
	Monthreshold         types.Int64  `tfsdk:"monthreshold"`
	Nameserver           types.String `tfsdk:"nameserver"`
	Netprofile           types.String `tfsdk:"netprofile"`
	Pathmonitor          types.String `tfsdk:"pathmonitor"`
	Pathmonitorindv      types.String `tfsdk:"pathmonitorindv"`
	Port                 types.Int64  `tfsdk:"port"`
	Quicprofilename      types.String `tfsdk:"quicprofilename"`
	Riseapbrstatsmsgcode types.Int64  `tfsdk:"riseapbrstatsmsgcode"`
	Rtspsessionidremap   types.String `tfsdk:"rtspsessionidremap"`
	Serverid             types.Int64  `tfsdk:"serverid"`
	Servername           types.String `tfsdk:"servername"`
	Servicegroupname     types.String `tfsdk:"servicegroupname"` // Required lookup key
	Servicetype          types.String `tfsdk:"servicetype"`
	Sp                   types.String `tfsdk:"sp"`
	State                types.String `tfsdk:"state"`
	Svrtimeout           types.Int64  `tfsdk:"svrtimeout"`
	Tcpb                 types.String `tfsdk:"tcpb"`
	Tcpprofilename       types.String `tfsdk:"tcpprofilename"`
	Td                   types.Int64  `tfsdk:"td"`
	Topicname            types.String `tfsdk:"topicname"`
	Useproxyport         types.String `tfsdk:"useproxyport"`
	Usip                 types.String `tfsdk:"usip"`
	Wasmmodule           types.String `tfsdk:"wasmmodule"`
	Weight               types.Int64  `tfsdk:"weight"`

	// Convenience blocks (shared model with the resource); action-only inputs the
	// GET never returns.
	Lbvservers                      types.Set    `tfsdk:"lbvservers"`
	Lbmonitor                       types.String `tfsdk:"lbmonitor"`
	Servicegroupmembers             types.Set    `tfsdk:"servicegroupmembers"`
	ServicegroupmembersByServername types.Set    `tfsdk:"servicegroupmembers_by_servername"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/servicegroup.json). Never settable; populated from GET.
	Numofconnections           types.Int64  `tfsdk:"numofconnections"`
	Serviceconftype            types.Bool   `tfsdk:"serviceconftype"`
	Value                      types.String `tfsdk:"value"`
	Svrstate                   types.String `tfsdk:"svrstate"`
	Ip                         types.String `tfsdk:"ip"`
	Monstatcode                types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1              types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2              types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3              types.Int64  `tfsdk:"monstatparam3"`
	Statechangetimemsec        types.Int64  `tfsdk:"statechangetimemsec"`
	Stateupdatereason          types.Int64  `tfsdk:"stateupdatereason"`
	Clmonowner                 types.Int64  `tfsdk:"clmonowner"`
	Clmonview                  types.Int64  `tfsdk:"clmonview"`
	Groupcount                 types.Int64  `tfsdk:"groupcount"`
	Serviceipstr               types.String `tfsdk:"serviceipstr"`
	Servicegroupeffectivestate types.String `tfsdk:"servicegroupeffectivestate"`
	Nodefaultbindings          types.String `tfsdk:"nodefaultbindings"`
	Svcitmactsvcs              types.Int64  `tfsdk:"svcitmactsvcs"`
	Svcitmboundsvcs            types.Int64  `tfsdk:"svcitmboundsvcs"`
	Monuserstatusmesg          types.String `tfsdk:"monuserstatusmesg"`
}

func ServicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aigwprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autodisabledelay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"autodisablegraceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autoscale": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"bootstrap": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cacheable": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cachetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cip": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cipheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cka": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"clttimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"cmp": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"customserverid": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"downstateflush": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dupweight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"hashid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"healthmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"httpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"includemembers": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"maxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"maxreq": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"mcpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"memberport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"monconnectionclose": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monthreshold": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"nameserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"netprofile": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"pathmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"quicprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"riseapbrstatsmsgcode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"serverid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
			},
			"servicetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sp": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"svrtimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"tcpb": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"topicname": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useproxyport": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"usip": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"wasmmodule": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},

			// Convenience blocks (shared model with the resource).
			"lbvservers": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"lbmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"servicegroupmembers": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"servicegroupmembers_by_servername": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numofconnections": schema.Int64Attribute{
				Computed:    true,
				Description: "This will tell the number of client side connections are still open.",
			},
			"serviceconftype": schema.BoolAttribute{
				Computed:    true,
				Description: "The configuration type of the service group.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "SSL Status. Possible values = Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the service. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "IP Address.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
			"statechangetimemsec": schema.Int64Attribute{
				Computed:    true,
				Description: "Time when last state change occurred. Milliseconds part.",
			},
			"stateupdatereason": schema.Int64Attribute{
				Computed:    true,
				Description: "Checks state update reason on the secondary node.",
			},
			"clmonowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the mon owner of the service.",
			},
			"clmonview": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the view id of the monitoring owner.",
			},
			"groupcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Servicegroup Count.",
			},
			"serviceipstr": schema.StringAttribute{
				Computed:    true,
				Description: "This field has been introduced to show the dbs services ip.",
			},
			"servicegroupeffectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates the effective servicegroup state based on the state of the bound service items. Possible values = UP, DOWN, OUT OF SERVICE, PARTIAL-UP, PARTIAL-DOWN.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration is from stylebooks. Possible values = YES, NO.",
			},
			"svcitmactsvcs": schema.Int64Attribute{
				Computed:    true,
				Description: "This gives the total active service items for an FQDN for SRV type server binding.",
			},
			"svcitmboundsvcs": schema.Int64Attribute{
				Computed:    true,
				Description: "This gives the total bound items for an FQDN for SRV type server binding.",
			},
			"monuserstatusmesg": schema.StringAttribute{
				Computed:    true,
				Description: "This field has been introduced to show user monitor failure reasons.",
			},
		},
	}
}

// servicegroupDataSourceSetAttrFromGet projects a NITRO servicegroup GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection. A few NITRO keys differ from the TF attribute names
// (dup_weight -> dupweight, monitor_name_svc -> monitornamesvc) and sp is
// normalized as in the resource read.
func servicegroupDataSourceSetAttrFromGet(ctx context.Context, data *ServicegroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In servicegroupDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Autodelayedtrofs = utils.MapGetString(g, "autodelayedtrofs")
	data.Autodisabledelay = utils.MapGetInt64(g, "autodisabledelay")
	data.Autodisablegraceful = utils.MapGetString(g, "autodisablegraceful")
	data.Autoscale = utils.MapGetString(g, "autoscale")
	data.Bootstrap = utils.MapGetString(g, "bootstrap")
	data.Cacheable = utils.MapGetString(g, "cacheable")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Cip = utils.MapGetString(g, "cip")
	data.Cipheader = utils.MapGetString(g, "cipheader")
	data.Cka = utils.MapGetString(g, "cka")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Cmp = utils.MapGetString(g, "cmp")
	data.Comment = utils.MapGetString(g, "comment")
	data.Customserverid = utils.MapGetString(g, "customserverid")
	data.Dbsttl = utils.MapGetInt64(g, "dbsttl")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Dupweight = utils.MapGetInt64(g, "dup_weight")
	data.Graceful = utils.MapGetString(g, "graceful")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Healthmonitor = utils.MapGetString(g, "healthmonitor")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Includemembers = utils.MapGetBool(g, "includemembers")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Maxclient = utils.MapGetInt64(g, "maxclient")
	data.Maxreq = utils.MapGetInt64(g, "maxreq")
	data.Mcpprofilename = utils.MapGetString(g, "mcpprofilename")
	data.Memberport = utils.MapGetInt64(g, "memberport")
	data.Monconnectionclose = utils.MapGetString(g, "monconnectionclose")
	data.Monitornamesvc = utils.MapGetString(g, "monitor_name_svc")
	data.Monthreshold = utils.MapGetInt64(g, "monthreshold")
	data.Nameserver = utils.MapGetString(g, "nameserver")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Pathmonitor = utils.MapGetString(g, "pathmonitor")
	data.Pathmonitorindv = utils.MapGetString(g, "pathmonitorindv")
	data.Port = utils.MapGetInt64(g, "port")
	data.Quicprofilename = utils.MapGetString(g, "quicprofilename")
	data.Riseapbrstatsmsgcode = utils.MapGetInt64(g, "riseapbrstatsmsgcode")
	data.Rtspsessionidremap = utils.MapGetString(g, "rtspsessionidremap")
	data.Serverid = utils.MapGetInt64(g, "serverid")
	data.Servername = utils.MapGetString(g, "servername")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	// sp: NITRO may report "ON (but effectively OFF)"; normalize to "ON".
	data.Sp = utils.MapGetString(g, "sp")
	if data.Sp.ValueString() == "ON (but effectively OFF)" {
		data.Sp = types.StringValue("ON")
	}
	data.State = utils.MapGetString(g, "state")
	data.Svrtimeout = utils.MapGetInt64(g, "svrtimeout")
	data.Tcpb = utils.MapGetString(g, "tcpb")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Topicname = utils.MapGetString(g, "topicname")
	data.Useproxyport = utils.MapGetString(g, "useproxyport")
	data.Usip = utils.MapGetString(g, "usip")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Convenience blocks are action-only inputs the GET never returns -> Null.
	data.Lbvservers = types.SetNull(types.StringType)
	data.Lbmonitor = types.StringNull()
	data.Servicegroupmembers = types.SetNull(types.StringType)
	data.ServicegroupmembersByServername = types.SetNull(types.StringType)

	// Read-only metadata.
	data.Numofconnections = utils.MapGetInt64(g, "numofconnections")
	data.Serviceconftype = utils.MapGetBool(g, "serviceconftype")
	data.Value = utils.MapGetString(g, "value")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Ip = utils.MapGetString(g, "ip")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Stateupdatereason = utils.MapGetInt64(g, "stateupdatereason")
	data.Clmonowner = utils.MapGetInt64(g, "clmonowner")
	data.Clmonview = utils.MapGetInt64(g, "clmonview")
	data.Groupcount = utils.MapGetInt64(g, "groupcount")
	data.Serviceipstr = utils.MapGetString(g, "serviceipstr")
	data.Servicegroupeffectivestate = utils.MapGetString(g, "servicegroupeffectivestate")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Svcitmactsvcs = utils.MapGetInt64(g, "svcitmactsvcs")
	data.Svcitmboundsvcs = utils.MapGetInt64(g, "svcitmboundsvcs")
	data.Monuserstatusmesg = utils.MapGetString(g, "monuserstatusmesg")

	// Set ID for the datasource (named resource - plain servicegroupname value).
	data.Id = types.StringValue(data.Servicegroupname.ValueString())
}
