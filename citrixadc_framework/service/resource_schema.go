package service

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

// ServiceResourceModel describes the resource data model.
type ServiceResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Internal                     types.Bool   `tfsdk:"internal"`
	Accessdown                   types.String `tfsdk:"accessdown"`
	All                          types.Bool   `tfsdk:"all"`
	Appflowlog                   types.String `tfsdk:"appflowlog"`
	Cacheable                    types.String `tfsdk:"cacheable"`
	Cachetype                    types.String `tfsdk:"cachetype"`
	Cip                          types.String `tfsdk:"cip"`
	Cipheader                    types.String `tfsdk:"cipheader"`
	Cka                          types.String `tfsdk:"cka"`
	Cleartextport                types.Int64  `tfsdk:"cleartextport"`
	Clttimeout                   types.Int64  `tfsdk:"clttimeout"`
	Cmp                          types.String `tfsdk:"cmp"`
	Comment                      types.String `tfsdk:"comment"`
	Contentinspectionprofilename types.String `tfsdk:"contentinspectionprofilename"`
	Customserverid               types.String `tfsdk:"customserverid"`
	Delay                        types.Int64  `tfsdk:"delay"`
	Dnsprofilename               types.String `tfsdk:"dnsprofilename"`
	Downstateflush               types.String `tfsdk:"downstateflush"`
	Graceful                     types.String `tfsdk:"graceful"`
	Hashid                       types.Int64  `tfsdk:"hashid"`
	Healthmonitor                types.String `tfsdk:"healthmonitor"`
	Httpprofilename              types.String `tfsdk:"httpprofilename"`
	Ip                           types.String `tfsdk:"ip"`
	Ipaddress                    types.String `tfsdk:"ipaddress"`
	Maxbandwidth                 types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient                    types.Int64  `tfsdk:"maxclient"`
	Maxreq                       types.Int64  `tfsdk:"maxreq"`
	Monconnectionclose           types.String `tfsdk:"monconnectionclose"`
	MonitorNameSvc               types.String `tfsdk:"monitor_name_svc"`
	Monthreshold                 types.Int64  `tfsdk:"monthreshold"`
	Name                         types.String `tfsdk:"name"`
	Netprofile                   types.String `tfsdk:"netprofile"`
	Newname                      types.String `tfsdk:"newname"`
	Pathmonitor                  types.String `tfsdk:"pathmonitor"`
	Pathmonitorindv              types.String `tfsdk:"pathmonitorindv"`
	Port                         types.Int64  `tfsdk:"port"`
	Processlocal                 types.String `tfsdk:"processlocal"`
	Quicprofilename              types.String `tfsdk:"quicprofilename"`
	Rtspsessionidremap           types.String `tfsdk:"rtspsessionidremap"`
	Serverid                     types.Int64  `tfsdk:"serverid"`
	Servername                   types.String `tfsdk:"servername"`
	Servicetype                  types.String `tfsdk:"servicetype"`
	Sp                           types.String `tfsdk:"sp"`
	State                        types.String `tfsdk:"state"`
	Svrtimeout                   types.Int64  `tfsdk:"svrtimeout"`
	Tcpb                         types.String `tfsdk:"tcpb"`
	Tcpprofilename               types.String `tfsdk:"tcpprofilename"`
	Td                           types.Int64  `tfsdk:"td"`
	Useproxyport                 types.String `tfsdk:"useproxyport"`
	Usip                         types.String `tfsdk:"usip"`
	Weight                       types.Int64  `tfsdk:"weight"`
}

func (r *ServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the service resource.",
			},
			"internal": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display only dynamically learned services.",
			},
			"accessdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display both user-configured and dynamically learned services.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging of AppFlow information.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward requests to the cache server.\nNote: Do not specify this parameter if you set the Cache Type parameter.",
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
				Description: "Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name \"client-ip.\"",
			},
			"cka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable client keep-alive for the service.",
			},
			"cleartextport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"cmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable compression for the service.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the service.",
			},
			"contentinspectionprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS, ADNS-TCP and ADNS-DOT services.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Monitor the health of this service. Available settings function as follows:\nYES - Send probes to check the health of the service.\nNO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service.",
			},
			"ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP to assign to the service.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated to the service.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections to the service.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that can be sent on a persistent connection to the service.\nNote: Connection requests beyond this value are rejected.",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the specified service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the service has been created.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile to use for the service.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"pathmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path monitoring for clustering",
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Individual Path monitoring decisions",
			},
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number of the service.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service.",
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the server that hosts the service.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol in which data is exchanged with the service.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable surge protection for the service.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"tcpb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP buffering for the service.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile that contains TCP configuration settings for the service.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.\nNote: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},
		},
	}
}

func serviceGetThePayloadFromtheConfig(ctx context.Context, data *ServiceResourceModel) basic.Service {
	tflog.Debug(ctx, "In serviceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	service := basic.Service{}
	if !data.Internal.IsNull() {
		service.Internal = data.Internal.ValueBool()
	}
	if !data.Accessdown.IsNull() {
		service.Accessdown = data.Accessdown.ValueString()
	}
	if !data.All.IsNull() {
		service.All = data.All.ValueBool()
	}
	if !data.Appflowlog.IsNull() {
		service.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cacheable.IsNull() {
		service.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cachetype.IsNull() {
		service.Cachetype = data.Cachetype.ValueString()
	}
	if !data.Cip.IsNull() {
		service.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() {
		service.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() {
		service.Cka = data.Cka.ValueString()
	}
	if !data.Cleartextport.IsNull() {
		service.Cleartextport = utils.IntPtr(int(data.Cleartextport.ValueInt64()))
	}
	if !data.Clttimeout.IsNull() {
		service.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() {
		service.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() {
		service.Comment = data.Comment.ValueString()
	}
	if !data.Contentinspectionprofilename.IsNull() {
		service.Contentinspectionprofilename = data.Contentinspectionprofilename.ValueString()
	}
	if !data.Customserverid.IsNull() {
		service.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Delay.IsNull() {
		service.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Dnsprofilename.IsNull() {
		service.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Downstateflush.IsNull() {
		service.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Graceful.IsNull() {
		service.Graceful = data.Graceful.ValueString()
	}
	if !data.Hashid.IsNull() {
		service.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() {
		service.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() {
		service.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Ip.IsNull() {
		service.Ip = data.Ip.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		service.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxbandwidth.IsNull() {
		service.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() {
		service.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() {
		service.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() {
		service.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.MonitorNameSvc.IsNull() {
		service.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() {
		service.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Name.IsNull() {
		service.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		service.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		service.Newname = data.Newname.ValueString()
	}
	if !data.Pathmonitor.IsNull() {
		service.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() {
		service.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Port.IsNull() {
		service.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Processlocal.IsNull() {
		service.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Quicprofilename.IsNull() {
		service.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Rtspsessionidremap.IsNull() {
		service.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() {
		service.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() {
		service.Servername = data.Servername.ValueString()
	}
	if !data.Servicetype.IsNull() {
		service.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sp.IsNull() {
		service.Sp = data.Sp.ValueString()
	}
	if !data.State.IsNull() {
		service.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() {
		service.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() {
		service.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() {
		service.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() {
		service.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useproxyport.IsNull() {
		service.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() {
		service.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() {
		service.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return service
}

func serviceSetAttrFromGet(ctx context.Context, data *ServiceResourceModel, getResponseData map[string]interface{}) *ServiceResourceModel {
	tflog.Debug(ctx, "In serviceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Internal"]; ok && val != nil {
		data.Internal = types.BoolValue(val.(bool))
	} else {
		data.Internal = types.BoolNull()
	}
	if val, ok := getResponseData["accessdown"]; ok && val != nil {
		data.Accessdown = types.StringValue(val.(string))
	} else {
		data.Accessdown = types.StringNull()
	}
	if val, ok := getResponseData["all"]; ok && val != nil {
		data.All = types.BoolValue(val.(bool))
	} else {
		data.All = types.BoolNull()
	}
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
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
	if val, ok := getResponseData["cleartextport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cleartextport = types.Int64Value(intVal)
		}
	} else {
		data.Cleartextport = types.Int64Null()
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
	if val, ok := getResponseData["contentinspectionprofilename"]; ok && val != nil {
		data.Contentinspectionprofilename = types.StringValue(val.(string))
	} else {
		data.Contentinspectionprofilename = types.StringNull()
	}
	if val, ok := getResponseData["customserverid"]; ok && val != nil {
		data.Customserverid = types.StringValue(val.(string))
	} else {
		data.Customserverid = types.StringNull()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
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
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	if val, ok := getResponseData["processlocal"]; ok && val != nil {
		data.Processlocal = types.StringValue(val.(string))
	} else {
		data.Processlocal = types.StringNull()
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
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Name.ValueString()))

	return data
}
