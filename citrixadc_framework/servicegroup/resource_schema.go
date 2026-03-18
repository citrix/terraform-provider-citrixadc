package servicegroup

import (
	"context"
	"fmt"

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
type ServicegroupResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Appflowlog          types.String `tfsdk:"appflowlog"`
	Autodelayedtrofs    types.String `tfsdk:"autodelayedtrofs"`
	Autodisabledelay    types.Int64  `tfsdk:"autodisabledelay"`
	Autodisablegraceful types.String `tfsdk:"autodisablegraceful"`
	Autoscale           types.String `tfsdk:"autoscale"`
	Bootstrap           types.String `tfsdk:"bootstrap"`
	Cacheable           types.String `tfsdk:"cacheable"`
	Cachetype           types.String `tfsdk:"cachetype"`
	Cip                 types.String `tfsdk:"cip"`
	Cipheader           types.String `tfsdk:"cipheader"`
	Cka                 types.String `tfsdk:"cka"`
	Clttimeout          types.Int64  `tfsdk:"clttimeout"`
	Cmp                 types.String `tfsdk:"cmp"`
	Comment             types.String `tfsdk:"comment"`
	Customserverid      types.String `tfsdk:"customserverid"`
	Dbsttl              types.Int64  `tfsdk:"dbsttl"`
	Delay               types.Int64  `tfsdk:"delay"`
	Downstateflush      types.String `tfsdk:"downstateflush"`
	DupWeight           types.Int64  `tfsdk:"dup_weight"`
	Graceful            types.String `tfsdk:"graceful"`
	Hashid              types.Int64  `tfsdk:"hashid"`
	Healthmonitor       types.String `tfsdk:"healthmonitor"`
	Httpprofilename     types.String `tfsdk:"httpprofilename"`
	Includemembers      types.Bool   `tfsdk:"includemembers"`
	Maxbandwidth        types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient           types.Int64  `tfsdk:"maxclient"`
	Maxreq              types.Int64  `tfsdk:"maxreq"`
	Memberport          types.Int64  `tfsdk:"memberport"`
	Monconnectionclose  types.String `tfsdk:"monconnectionclose"`
	MonitorNameSvc      types.String `tfsdk:"monitor_name_svc"`
	Monthreshold        types.Int64  `tfsdk:"monthreshold"`
	Nameserver          types.String `tfsdk:"nameserver"`
	Netprofile          types.String `tfsdk:"netprofile"`
	Newname             types.String `tfsdk:"newname"`
	Order               types.Int64  `tfsdk:"order"`
	Pathmonitor         types.String `tfsdk:"pathmonitor"`
	Pathmonitorindv     types.String `tfsdk:"pathmonitorindv"`
	Port                types.Int64  `tfsdk:"port"`
	Quicprofilename     types.String `tfsdk:"quicprofilename"`
	Rtspsessionidremap  types.String `tfsdk:"rtspsessionidremap"`
	Serverid            types.Int64  `tfsdk:"serverid"`
	Servername          types.String `tfsdk:"servername"`
	Servicegroupname    types.String `tfsdk:"servicegroupname"`
	Servicetype         types.String `tfsdk:"servicetype"`
	Sp                  types.String `tfsdk:"sp"`
	State               types.String `tfsdk:"state"`
	Svrtimeout          types.Int64  `tfsdk:"svrtimeout"`
	Tcpb                types.String `tfsdk:"tcpb"`
	Tcpprofilename      types.String `tfsdk:"tcpprofilename"`
	Td                  types.Int64  `tfsdk:"td"`
	Topicname           types.String `tfsdk:"topicname"`
	Useproxyport        types.String `tfsdk:"useproxyport"`
	Usip                types.String `tfsdk:"usip"`
	Weight              types.Int64  `tfsdk:"weight"`
}

func (r *ServicegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the servicegroup resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
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
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"autodisablegraceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.",
			},
			"autoscale": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Auto scale option for a servicegroup",
			},
			"bootstrap": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Flag to check if kafka broker servicegroup is of type bootstrap or not.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward the request to the cache server.\nNote: Do not set this parameter if you set the Cache Type.",
			},
			"cachetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
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
				Default:     stringdefault.StaticString("None"),
				Description: "The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.",
			},
			"dbsttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to servicegroup.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Monitor the health of this service.  Available settings function as follows:\nYES - Send probes to check the health of the service.\nNO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service group.",
			},
			"includemembers": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display the members of the listed service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
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
				Description: "Maximum number of requests that can be sent on a persistent connection to the service group.\nNote: Connection requests beyond this value are rejected.",
			},
			"memberport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "member port",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.",
			},
			"nameserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile for the service group.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the servicegroup member",
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
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service group.",
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
				Description: "Name of the service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used to exchange data with the service.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable surge protection for the service group.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"topicname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Kafka topic.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.\nNote: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client's IP address as the source IP address when initiating connection to the server. With the NO setting, which is the default, a mapped IP (MIP) address or subnet IP (SNIP) address is used as the source IP address to initiate server side connections.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

func servicegroupGetThePayloadFromtheConfig(ctx context.Context, data *ServicegroupResourceModel) basic.Servicegroup {
	tflog.Debug(ctx, "In servicegroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	servicegroup := basic.Servicegroup{}
	if !data.Appflowlog.IsNull() {
		servicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Autodelayedtrofs.IsNull() {
		servicegroup.Autodelayedtrofs = data.Autodelayedtrofs.ValueString()
	}
	if !data.Autodisabledelay.IsNull() {
		servicegroup.Autodisabledelay = utils.IntPtr(int(data.Autodisabledelay.ValueInt64()))
	}
	if !data.Autodisablegraceful.IsNull() {
		servicegroup.Autodisablegraceful = data.Autodisablegraceful.ValueString()
	}
	if !data.Autoscale.IsNull() {
		servicegroup.Autoscale = data.Autoscale.ValueString()
	}
	if !data.Bootstrap.IsNull() {
		servicegroup.Bootstrap = data.Bootstrap.ValueString()
	}
	if !data.Cacheable.IsNull() {
		servicegroup.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cachetype.IsNull() {
		servicegroup.Cachetype = data.Cachetype.ValueString()
	}
	if !data.Cip.IsNull() {
		servicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() {
		servicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() {
		servicegroup.Cka = data.Cka.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		servicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() {
		servicegroup.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() {
		servicegroup.Comment = data.Comment.ValueString()
	}
	if !data.Customserverid.IsNull() {
		servicegroup.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dbsttl.IsNull() {
		servicegroup.Dbsttl = utils.IntPtr(int(data.Dbsttl.ValueInt64()))
	}
	if !data.Delay.IsNull() {
		servicegroup.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Downstateflush.IsNull() {
		servicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.DupWeight.IsNull() {
		servicegroup.Dupweight = utils.IntPtr(int(data.DupWeight.ValueInt64()))
	}
	if !data.Graceful.IsNull() {
		servicegroup.Graceful = data.Graceful.ValueString()
	}
	if !data.Hashid.IsNull() {
		servicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() {
		servicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() {
		servicegroup.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Includemembers.IsNull() {
		servicegroup.Includemembers = data.Includemembers.ValueBool()
	}
	if !data.Maxbandwidth.IsNull() {
		servicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() {
		servicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() {
		servicegroup.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Memberport.IsNull() {
		servicegroup.Memberport = utils.IntPtr(int(data.Memberport.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() {
		servicegroup.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.MonitorNameSvc.IsNull() {
		servicegroup.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() {
		servicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Nameserver.IsNull() {
		servicegroup.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Netprofile.IsNull() {
		servicegroup.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		servicegroup.Newname = data.Newname.ValueString()
	}
	if !data.Order.IsNull() {
		servicegroup.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Pathmonitor.IsNull() {
		servicegroup.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() {
		servicegroup.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Port.IsNull() {
		servicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Quicprofilename.IsNull() {
		servicegroup.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Rtspsessionidremap.IsNull() {
		servicegroup.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() {
		servicegroup.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() {
		servicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		servicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicetype.IsNull() {
		servicegroup.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sp.IsNull() {
		servicegroup.Sp = data.Sp.ValueString()
	}
	if !data.State.IsNull() {
		servicegroup.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() {
		servicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() {
		servicegroup.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() {
		servicegroup.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() {
		servicegroup.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Topicname.IsNull() {
		servicegroup.Topicname = data.Topicname.ValueString()
	}
	if !data.Useproxyport.IsNull() {
		servicegroup.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() {
		servicegroup.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() {
		servicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return servicegroup
}

func servicegroupSetAttrFromGet(ctx context.Context, data *ServicegroupResourceModel, getResponseData map[string]interface{}) *ServicegroupResourceModel {
	tflog.Debug(ctx, "In servicegroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["autodelayedtrofs"]; ok && val != nil {
		data.Autodelayedtrofs = types.StringValue(val.(string))
	} else {
		data.Autodelayedtrofs = types.StringNull()
	}
	if val, ok := getResponseData["autodisabledelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Autodisabledelay = types.Int64Value(intVal)
		}
	} else {
		data.Autodisabledelay = types.Int64Null()
	}
	if val, ok := getResponseData["autodisablegraceful"]; ok && val != nil {
		data.Autodisablegraceful = types.StringValue(val.(string))
	} else {
		data.Autodisablegraceful = types.StringNull()
	}
	if val, ok := getResponseData["autoscale"]; ok && val != nil {
		data.Autoscale = types.StringValue(val.(string))
	} else {
		data.Autoscale = types.StringNull()
	}
	if val, ok := getResponseData["bootstrap"]; ok && val != nil {
		data.Bootstrap = types.StringValue(val.(string))
	} else {
		data.Bootstrap = types.StringNull()
	}
	if val, ok := getResponseData["cacheable"]; ok && val != nil {
		data.Cacheable = types.StringValue(val.(string))
	} else {
		data.Cacheable = types.StringNull()
	}
	if val, ok := getResponseData["cachetype"]; ok && val != nil {
		data.Cachetype = types.StringValue(val.(string))
	} else {
		data.Cachetype = types.StringNull()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["cka"]; ok && val != nil {
		data.Cka = types.StringValue(val.(string))
	} else {
		data.Cka = types.StringNull()
	}
	if val, ok := getResponseData["clttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Clttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["cmp"]; ok && val != nil {
		data.Cmp = types.StringValue(val.(string))
	} else {
		data.Cmp = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["customserverid"]; ok && val != nil {
		data.Customserverid = types.StringValue(val.(string))
	} else {
		data.Customserverid = types.StringNull()
	}
	if val, ok := getResponseData["dbsttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dbsttl = types.Int64Value(intVal)
		}
	} else {
		data.Dbsttl = types.Int64Null()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["dup_weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.DupWeight = types.Int64Value(intVal)
		}
	} else {
		data.DupWeight = types.Int64Null()
	}
	if val, ok := getResponseData["graceful"]; ok && val != nil {
		data.Graceful = types.StringValue(val.(string))
	} else {
		data.Graceful = types.StringNull()
	}
	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	} else {
		data.Hashid = types.Int64Null()
	}
	if val, ok := getResponseData["healthmonitor"]; ok && val != nil {
		data.Healthmonitor = types.StringValue(val.(string))
	} else {
		data.Healthmonitor = types.StringNull()
	}
	if val, ok := getResponseData["httpprofilename"]; ok && val != nil {
		data.Httpprofilename = types.StringValue(val.(string))
	} else {
		data.Httpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["includemembers"]; ok && val != nil {
		data.Includemembers = types.BoolValue(val.(bool))
	} else {
		data.Includemembers = types.BoolNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["maxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxclient = types.Int64Value(intVal)
		}
	} else {
		data.Maxclient = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["memberport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memberport = types.Int64Value(intVal)
		}
	} else {
		data.Memberport = types.Int64Null()
	}
	if val, ok := getResponseData["monconnectionclose"]; ok && val != nil {
		data.Monconnectionclose = types.StringValue(val.(string))
	} else {
		data.Monconnectionclose = types.StringNull()
	}
	if val, ok := getResponseData["monitor_name_svc"]; ok && val != nil {
		data.MonitorNameSvc = types.StringValue(val.(string))
	} else {
		data.MonitorNameSvc = types.StringNull()
	}
	if val, ok := getResponseData["monthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Monthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Monthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	} else {
		data.Nameserver = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["pathmonitor"]; ok && val != nil {
		data.Pathmonitor = types.StringValue(val.(string))
	} else {
		data.Pathmonitor = types.StringNull()
	}
	if val, ok := getResponseData["pathmonitorindv"]; ok && val != nil {
		data.Pathmonitorindv = types.StringValue(val.(string))
	} else {
		data.Pathmonitorindv = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["quicprofilename"]; ok && val != nil {
		data.Quicprofilename = types.StringValue(val.(string))
	} else {
		data.Quicprofilename = types.StringNull()
	}
	if val, ok := getResponseData["rtspsessionidremap"]; ok && val != nil {
		data.Rtspsessionidremap = types.StringValue(val.(string))
	} else {
		data.Rtspsessionidremap = types.StringNull()
	}
	if val, ok := getResponseData["serverid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverid = types.Int64Value(intVal)
		}
	} else {
		data.Serverid = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sp"]; ok && val != nil {
		data.Sp = types.StringValue(val.(string))
	} else {
		data.Sp = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["svrtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Svrtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Svrtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpb"]; ok && val != nil {
		data.Tcpb = types.StringValue(val.(string))
	} else {
		data.Tcpb = types.StringNull()
	}
	if val, ok := getResponseData["tcpprofilename"]; ok && val != nil {
		data.Tcpprofilename = types.StringValue(val.(string))
	} else {
		data.Tcpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["topicname"]; ok && val != nil {
		data.Topicname = types.StringValue(val.(string))
	} else {
		data.Topicname = types.StringNull()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else {
		data.Useproxyport = types.StringNull()
	}
	if val, ok := getResponseData["usip"]; ok && val != nil {
		data.Usip = types.StringValue(val.(string))
	} else {
		data.Usip = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Servicegroupname.ValueString()))

	return data
}
