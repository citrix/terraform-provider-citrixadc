package servicegroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ServicegroupResourceModel describes the resource data model.
//
// Attribute names/types mirror the legacy SDK v2 resource
// (citrixadc/resource_citrixadc_servicegroup.go) for backward compatibility,
// including the convenience blocks lbvservers/lbmonitor/servicegroupmembers/
// servicegroupmembers_by_servername that manage the associated bindings.
type ServicegroupResourceModel struct {
	Id                   types.String `tfsdk:"id"`
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
	Servicegroupname     types.String `tfsdk:"servicegroupname"`
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
	Weight               types.Int64  `tfsdk:"weight"`

	// Convenience blocks (SDK v2 backward compatibility). These are not NITRO
	// servicegroup attributes; they manage the associated bindings.
	Lbvservers                      types.Set    `tfsdk:"lbvservers"`
	Lbmonitor                       types.String `tfsdk:"lbmonitor"`
	Servicegroupmembers             types.Set    `tfsdk:"servicegroupmembers"`
	ServicegroupmembersByServername types.Set    `tfsdk:"servicegroupmembers_by_servername"`
}

func (r *ServicegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the servicegroup resource (the servicegroupname).",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging of AppFlow information for the specified service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful movement of IP-Port binding/s to TROFS when IP addresses are removed from DNS response. System will wait for monitor response timeout period before moving to TROFS .",
			},
			"autodisabledelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown.",
			},
			"autodisablegraceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful shutdown of the service.",
			},
			"autoscale": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Auto scale option for a servicegroup",
			},
			"bootstrap": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Flag to check if kafka broker servicegroup is of type bootstrap or not.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Use the transparent cache redirection virtual server to forward the request to the cache server.",
			},
			"cachetype": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Cache type supported by the cache server.",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client.",
			},
			"cka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable client keep-alive for the service group.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"cmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable compression for the specified service.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the service group.",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service.",
			},
			"delay": schema.Int64Attribute{
				// SDK v2: Optional+Computed (NOT ForceNew). Used as a disable-action
				// parameter, not sent in the add/set payload.
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, allocated for a shutdown of the services in the service group.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN.",
			},
			"dupweight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to servicegroup.",
			},
			"graceful": schema.StringAttribute{
				// SDK v2: Optional+Computed (NOT ForceNew). Used as a disable-action
				// parameter, not sent in the add/set payload.
				Optional:    true,
				Computed:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "Monitor the health of this service.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service group.",
			},
			"includemembers": schema.BoolAttribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Display the members of the listed service groups in addition to their settings.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the service group.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the service group.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that can be sent on a persistent connection to the service group.",
			},
			"memberport": schema.Int64Attribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "member port",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitornamesvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service.",
			},
			"nameserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the nameserver to which the query for bound domain needs to be sent.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile for the service group.",
			},
			"pathmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path monitoring for clustering",
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Individual Path monitoring decisions.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service group.",
			},
			"riseapbrstatsmsgcode": schema.Int64Attribute{
				// SDK v2: Optional+Computed+ForceNew. Read-only from the ADC (no NITRO
				// struct field), so it is never sent in the add/set payload.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The code indicating the rise apbr status.",
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service group.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service group. SDK v2 ForceNew.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used to exchange data with the service. SDK v2 ForceNew.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable surge protection for the service group.",
			},
			"state": schema.StringAttribute{
				// SDK v2: Optional+Computed (NOT ForceNew). Toggled in-place via the
				// enable/disable NITRO actions (see Update).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Initial state of the service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"tcpb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP buffering for the service group.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile that contains TCP configuration settings for the service group.",
			},
			"td": schema.Int64Attribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain.",
			},
			"topicname": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name of the Kafka topic.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client's IP address as the source IP address when initiating connection to the server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group.",
			},

			// Convenience blocks (SDK v2 backward compatibility).
			"lbvservers": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Set of lb vserver names to bind this service group to.",
			},
			"lbmonitor": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the lb monitor to bind to this service group.",
			},
			"servicegroupmembers": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Set of service group members in the form ip:port:weight.",
			},
			"servicegroupmembers_by_servername": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Set of service group members in the form servername:port:weight.",
			},
		},
	}
}

// servicegroupGetThePayloadFromthePlan builds the ADD (create) payload. It carries
// the full create-time attribute set including create-only (ForceNew) attrs and
// state. It EXCLUDES: riseapbrstatsmsgcode (no NITRO struct field, read-only) and
// the disable-action-only delay/graceful attrs (SDK v2 excludes them from the add
// payload; they are used only by the disable action).
func servicegroupGetThePayloadFromthePlan(ctx context.Context, data *ServicegroupResourceModel) basic.Servicegroup {
	tflog.Debug(ctx, "In servicegroupGetThePayloadFromthePlan Function")

	servicegroup := basic.Servicegroup{}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		servicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Autodelayedtrofs.IsNull() && !data.Autodelayedtrofs.IsUnknown() {
		servicegroup.Autodelayedtrofs = data.Autodelayedtrofs.ValueString()
	}
	if !data.Autodisabledelay.IsNull() && !data.Autodisabledelay.IsUnknown() {
		servicegroup.Autodisabledelay = utils.IntPtr(int(data.Autodisabledelay.ValueInt64()))
	}
	if !data.Autodisablegraceful.IsNull() && !data.Autodisablegraceful.IsUnknown() {
		servicegroup.Autodisablegraceful = data.Autodisablegraceful.ValueString()
	}
	if !data.Autoscale.IsNull() && !data.Autoscale.IsUnknown() {
		servicegroup.Autoscale = data.Autoscale.ValueString()
	}
	if !data.Bootstrap.IsNull() && !data.Bootstrap.IsUnknown() {
		servicegroup.Bootstrap = data.Bootstrap.ValueString()
	}
	if !data.Cacheable.IsNull() && !data.Cacheable.IsUnknown() {
		servicegroup.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cachetype.IsNull() && !data.Cachetype.IsUnknown() {
		servicegroup.Cachetype = data.Cachetype.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		servicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		servicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() && !data.Cka.IsUnknown() {
		servicegroup.Cka = data.Cka.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		servicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() && !data.Cmp.IsUnknown() {
		servicegroup.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		servicegroup.Comment = data.Comment.ValueString()
	}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		servicegroup.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dbsttl.IsNull() && !data.Dbsttl.IsUnknown() {
		servicegroup.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	// delay is a disable-action-only parameter - excluded from the add payload.
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		servicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Dupweight.IsNull() && !data.Dupweight.IsUnknown() {
		servicegroup.Dupweight = utils.IntPtr(int(data.Dupweight.ValueInt64()))
	}
	// graceful is a disable-action-only parameter - excluded from the add payload.
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		servicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		servicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() && !data.Httpprofilename.IsUnknown() {
		servicegroup.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Includemembers.IsNull() && !data.Includemembers.IsUnknown() {
		servicegroup.Includemembers = data.Includemembers.ValueBool()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		servicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		servicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		servicegroup.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Memberport.IsNull() && !data.Memberport.IsUnknown() {
		servicegroup.Memberport = utils.IntPtr(int(data.Memberport.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() && !data.Monconnectionclose.IsUnknown() {
		servicegroup.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.Monitornamesvc.IsNull() && !data.Monitornamesvc.IsUnknown() {
		servicegroup.Monitornamesvc = data.Monitornamesvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		servicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Nameserver.IsNull() && !data.Nameserver.IsUnknown() {
		servicegroup.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Netprofile.IsNull() && !data.Netprofile.IsUnknown() {
		servicegroup.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Pathmonitor.IsNull() && !data.Pathmonitor.IsUnknown() {
		servicegroup.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() && !data.Pathmonitorindv.IsUnknown() {
		servicegroup.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		servicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Quicprofilename.IsNull() && !data.Quicprofilename.IsUnknown() {
		servicegroup.Quicprofilename = data.Quicprofilename.ValueString()
	}
	// riseapbrstatsmsgcode has no NITRO struct field (read-only) - excluded.
	if !data.Rtspsessionidremap.IsNull() && !data.Rtspsessionidremap.IsUnknown() {
		servicegroup.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		servicegroup.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		servicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		servicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		servicegroup.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sp.IsNull() && !data.Sp.IsUnknown() {
		servicegroup.Sp = data.Sp.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		servicegroup.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		servicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() && !data.Tcpb.IsUnknown() {
		servicegroup.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() && !data.Tcpprofilename.IsUnknown() {
		servicegroup.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		servicegroup.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Topicname.IsNull() && !data.Topicname.IsUnknown() {
		servicegroup.Topicname = data.Topicname.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		servicegroup.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		servicegroup.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		servicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return servicegroup
}

// servicegroupGetTheUpdatePayloadFromthePlan builds the UPDATE (set) payload. It
// carries only the non-ForceNew, updateable attributes. It EXCLUDES the ForceNew
// attrs (autoscale, bootstrap, cachetype, topicname, includemembers, memberport,
// td, servicetype - they trigger recreation and never reach Update), state (toggled
// via the enable/disable actions), the disable-action-only delay/graceful, and
// riseapbrstatsmsgcode (no NITRO struct field).
func servicegroupGetTheUpdatePayloadFromthePlan(ctx context.Context, data *ServicegroupResourceModel) basic.Servicegroup {
	tflog.Debug(ctx, "In servicegroupGetTheUpdatePayloadFromthePlan Function")

	servicegroup := basic.Servicegroup{}
	// servicegroupname is the name key.
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		servicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		servicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Autodelayedtrofs.IsNull() && !data.Autodelayedtrofs.IsUnknown() {
		servicegroup.Autodelayedtrofs = data.Autodelayedtrofs.ValueString()
	}
	if !data.Autodisabledelay.IsNull() && !data.Autodisabledelay.IsUnknown() {
		servicegroup.Autodisabledelay = utils.IntPtr(int(data.Autodisabledelay.ValueInt64()))
	}
	if !data.Autodisablegraceful.IsNull() && !data.Autodisablegraceful.IsUnknown() {
		servicegroup.Autodisablegraceful = data.Autodisablegraceful.ValueString()
	}
	if !data.Cacheable.IsNull() && !data.Cacheable.IsUnknown() {
		servicegroup.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		servicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		servicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() && !data.Cka.IsUnknown() {
		servicegroup.Cka = data.Cka.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		servicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() && !data.Cmp.IsUnknown() {
		servicegroup.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		servicegroup.Comment = data.Comment.ValueString()
	}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		servicegroup.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dbsttl.IsNull() && !data.Dbsttl.IsUnknown() {
		servicegroup.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		servicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Dupweight.IsNull() && !data.Dupweight.IsUnknown() {
		servicegroup.Dupweight = utils.IntPtr(int(data.Dupweight.ValueInt64()))
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		servicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		servicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() && !data.Httpprofilename.IsUnknown() {
		servicegroup.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		servicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		servicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		servicegroup.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() && !data.Monconnectionclose.IsUnknown() {
		servicegroup.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.Monitornamesvc.IsNull() && !data.Monitornamesvc.IsUnknown() {
		servicegroup.Monitornamesvc = data.Monitornamesvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		servicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Nameserver.IsNull() && !data.Nameserver.IsUnknown() {
		servicegroup.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Netprofile.IsNull() && !data.Netprofile.IsUnknown() {
		servicegroup.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Pathmonitor.IsNull() && !data.Pathmonitor.IsUnknown() {
		servicegroup.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() && !data.Pathmonitorindv.IsUnknown() {
		servicegroup.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		servicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Quicprofilename.IsNull() && !data.Quicprofilename.IsUnknown() {
		servicegroup.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Rtspsessionidremap.IsNull() && !data.Rtspsessionidremap.IsUnknown() {
		servicegroup.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		servicegroup.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		servicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Sp.IsNull() && !data.Sp.IsUnknown() {
		servicegroup.Sp = data.Sp.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		servicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() && !data.Tcpb.IsUnknown() {
		servicegroup.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() && !data.Tcpprofilename.IsUnknown() {
		servicegroup.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		servicegroup.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		servicegroup.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		servicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return servicegroup
}

// servicegroupSetAttrFromGet updates the resource model from a NITRO GET response.
//
// Each else-branch is guarded with IsUnknown() (the omit-on-default fix): a value
// is only nulled when it is currently unknown (an unconfigured Computed attr on
// create), never when it holds a known/configured value that NITRO happens to omit
// from GET. This avoids both "inconsistent result after apply: unknown value" and
// clobbering a configured value that the ADC returns as a default. The convenience
// blocks (lbvservers/lbmonitor/servicegroupmembers*) are handled separately in the
// read helper.
func servicegroupSetAttrFromGet(ctx context.Context, data *ServicegroupResourceModel, getResponseData map[string]interface{}) *ServicegroupResourceModel {
	tflog.Debug(ctx, "In servicegroupSetAttrFromGet Function")

	setStr := func(key string, cur *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			*cur = types.StringValue(val.(string))
		} else if cur.IsUnknown() {
			*cur = types.StringNull()
		}
	}
	setInt := func(key string, cur *types.Int64) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				*cur = types.Int64Value(intVal)
			}
		} else if cur.IsUnknown() {
			*cur = types.Int64Null()
		}
	}

	setStr("appflowlog", &data.Appflowlog)
	setStr("autodelayedtrofs", &data.Autodelayedtrofs)
	setInt("autodisabledelay", &data.Autodisabledelay)
	setStr("autodisablegraceful", &data.Autodisablegraceful)
	setStr("autoscale", &data.Autoscale)
	setStr("bootstrap", &data.Bootstrap)
	setStr("cacheable", &data.Cacheable)
	setStr("cachetype", &data.Cachetype)
	setStr("cip", &data.Cip)
	setStr("cipheader", &data.Cipheader)
	setStr("cka", &data.Cka)
	setInt("clttimeout", &data.Clttimeout)
	setStr("cmp", &data.Cmp)
	setStr("comment", &data.Comment)
	setStr("customserverid", &data.Customserverid)
	setInt("dbsttl", &data.Dbsttl)
	// delay & graceful are disable-action-only inputs. SDK v2's read never reads
	// them back, and NITRO GET always returns the ADC defaults (delay=0,
	// graceful="NO") regardless of what was configured (they are not persisted in
	// the servicegroup config object). Adopt the GET value ONLY when the model
	// value is unknown (unconfigured Computed attr on create); otherwise preserve
	// the user's configured value to avoid "inconsistent result after apply".
	if data.Delay.IsUnknown() {
		setInt("delay", &data.Delay)
	}
	setStr("downstateflush", &data.Downstateflush)
	setInt("dup_weight", &data.Dupweight)
	if data.Graceful.IsUnknown() {
		setStr("graceful", &data.Graceful)
	}
	setInt("hashid", &data.Hashid)
	setStr("healthmonitor", &data.Healthmonitor)
	setStr("httpprofilename", &data.Httpprofilename)
	if val, ok := getResponseData["includemembers"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Includemembers = types.BoolValue(b)
		}
	} else if data.Includemembers.IsUnknown() {
		data.Includemembers = types.BoolNull()
	}
	setInt("maxbandwidth", &data.Maxbandwidth)
	setInt("maxclient", &data.Maxclient)
	setInt("maxreq", &data.Maxreq)
	setInt("memberport", &data.Memberport)
	setStr("monconnectionclose", &data.Monconnectionclose)
	setStr("monitor_name_svc", &data.Monitornamesvc)
	setInt("monthreshold", &data.Monthreshold)
	setStr("nameserver", &data.Nameserver)
	setStr("netprofile", &data.Netprofile)
	setStr("pathmonitor", &data.Pathmonitor)
	setStr("pathmonitorindv", &data.Pathmonitorindv)
	setInt("port", &data.Port)
	setStr("quicprofilename", &data.Quicprofilename)
	setInt("riseapbrstatsmsgcode", &data.Riseapbrstatsmsgcode)
	setStr("rtspsessionidremap", &data.Rtspsessionidremap)
	setInt("serverid", &data.Serverid)
	setStr("servername", &data.Servername)
	setStr("servicegroupname", &data.Servicegroupname)
	setStr("servicetype", &data.Servicetype)
	// sp: NITRO may report "ON (but effectively OFF)"; SDK v2 normalizes to "ON".
	if val, ok := getResponseData["sp"]; ok && val != nil {
		if val.(string) == "ON (but effectively OFF)" {
			data.Sp = types.StringValue("ON")
		} else {
			data.Sp = types.StringValue(val.(string))
		}
	} else if data.Sp.IsUnknown() {
		data.Sp = types.StringNull()
	}
	setStr("state", &data.State)
	setInt("svrtimeout", &data.Svrtimeout)
	setStr("tcpb", &data.Tcpb)
	setStr("tcpprofilename", &data.Tcpprofilename)
	setInt("td", &data.Td)
	setStr("topicname", &data.Topicname)
	setStr("useproxyport", &data.Useproxyport)
	setStr("usip", &data.Usip)
	setInt("weight", &data.Weight)

	return data
}

// servicegroupSetAttrFromGetForDatasource copies every attribute from the NITRO GET
// response (no config guards) and sets the datasource ID. It is the datasource-only
// counterpart of servicegroupSetAttrFromGet (Pattern 7).
func servicegroupSetAttrFromGetForDatasource(ctx context.Context, data *ServicegroupResourceModel, getResponseData map[string]interface{}) *ServicegroupResourceModel {
	tflog.Debug(ctx, "In servicegroupSetAttrFromGetForDatasource Function")

	setStr := func(key string, cur *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			*cur = types.StringValue(val.(string))
		} else {
			*cur = types.StringNull()
		}
	}
	setInt := func(key string, cur *types.Int64) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				*cur = types.Int64Value(intVal)
			} else {
				*cur = types.Int64Null()
			}
		} else {
			*cur = types.Int64Null()
		}
	}

	setStr("appflowlog", &data.Appflowlog)
	setStr("autodelayedtrofs", &data.Autodelayedtrofs)
	setInt("autodisabledelay", &data.Autodisabledelay)
	setStr("autodisablegraceful", &data.Autodisablegraceful)
	setStr("autoscale", &data.Autoscale)
	setStr("bootstrap", &data.Bootstrap)
	setStr("cacheable", &data.Cacheable)
	setStr("cachetype", &data.Cachetype)
	setStr("cip", &data.Cip)
	setStr("cipheader", &data.Cipheader)
	setStr("cka", &data.Cka)
	setInt("clttimeout", &data.Clttimeout)
	setStr("cmp", &data.Cmp)
	setStr("comment", &data.Comment)
	setStr("customserverid", &data.Customserverid)
	setInt("dbsttl", &data.Dbsttl)
	setInt("delay", &data.Delay)
	setStr("downstateflush", &data.Downstateflush)
	setInt("dup_weight", &data.Dupweight)
	setStr("graceful", &data.Graceful)
	setInt("hashid", &data.Hashid)
	setStr("healthmonitor", &data.Healthmonitor)
	setStr("httpprofilename", &data.Httpprofilename)
	if val, ok := getResponseData["includemembers"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Includemembers = types.BoolValue(b)
		} else {
			data.Includemembers = types.BoolNull()
		}
	} else {
		data.Includemembers = types.BoolNull()
	}
	setInt("maxbandwidth", &data.Maxbandwidth)
	setInt("maxclient", &data.Maxclient)
	setInt("maxreq", &data.Maxreq)
	setInt("memberport", &data.Memberport)
	setStr("monconnectionclose", &data.Monconnectionclose)
	setStr("monitor_name_svc", &data.Monitornamesvc)
	setInt("monthreshold", &data.Monthreshold)
	setStr("nameserver", &data.Nameserver)
	setStr("netprofile", &data.Netprofile)
	setStr("pathmonitor", &data.Pathmonitor)
	setStr("pathmonitorindv", &data.Pathmonitorindv)
	setInt("port", &data.Port)
	setStr("quicprofilename", &data.Quicprofilename)
	setInt("riseapbrstatsmsgcode", &data.Riseapbrstatsmsgcode)
	setStr("rtspsessionidremap", &data.Rtspsessionidremap)
	setInt("serverid", &data.Serverid)
	setStr("servername", &data.Servername)
	setStr("servicegroupname", &data.Servicegroupname)
	setStr("servicetype", &data.Servicetype)
	if val, ok := getResponseData["sp"]; ok && val != nil {
		if val.(string) == "ON (but effectively OFF)" {
			data.Sp = types.StringValue("ON")
		} else {
			data.Sp = types.StringValue(val.(string))
		}
	} else {
		data.Sp = types.StringNull()
	}
	setStr("state", &data.State)
	setInt("svrtimeout", &data.Svrtimeout)
	setStr("tcpb", &data.Tcpb)
	setStr("tcpprofilename", &data.Tcpprofilename)
	setInt("td", &data.Td)
	setStr("topicname", &data.Topicname)
	setStr("useproxyport", &data.Useproxyport)
	setStr("usip", &data.Usip)
	setInt("weight", &data.Weight)

	// Convenience blocks are not populated for the datasource.
	data.Lbmonitor = types.StringNull()

	// Set ID for the datasource (named resource - plain servicegroupname value).
	data.Id = types.StringValue(data.Servicegroupname.ValueString())

	return data
}
