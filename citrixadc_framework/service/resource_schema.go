package service

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	Monitornamesvc               types.String `tfsdk:"monitornamesvc"`
	Monthreshold                 types.Int64  `tfsdk:"monthreshold"`
	Name                         types.String `tfsdk:"name"`
	Netprofile                   types.String `tfsdk:"netprofile"`
	Pathmonitor                  types.String `tfsdk:"pathmonitor"`
	Pathmonitorindv              types.String `tfsdk:"pathmonitorindv"`
	Port                         types.Int64  `tfsdk:"port"`
	Processlocal                 types.String `tfsdk:"processlocal"`
	Quicprofilename              types.String `tfsdk:"quicprofilename"`
	Riseapbrstatsmsgcode         types.Int64  `tfsdk:"riseapbrstatsmsgcode"`
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

	// Convenience blocks preserved from the SDK v2 resource.
	Lbvserver            types.String `tfsdk:"lbvserver"`
	Lbmonitor            types.String `tfsdk:"lbmonitor"`
	Snienable            types.String `tfsdk:"snienable"`
	Commonname           types.String `tfsdk:"commonname"`
	WaitUntilDisabled    types.Bool   `tfsdk:"wait_until_disabled"`
	DisabledTimeout      types.String `tfsdk:"disabled_timeout"`
	DisabledPollDelay    types.String `tfsdk:"disabled_poll_delay"`
	DisabledPollInterval types.String `tfsdk:"disabled_poll_interval"`
}

func (r *ServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the service resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"internal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display only dynamically learned services.",
			},
			"accessdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.",
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display both user-configured and dynamically learned services.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging of AppFlow information.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Use the transparent cache redirection virtual server to forward requests to the cache server.\nNote: Do not specify this parameter if you set the Cache Type parameter.",
			},
			"cachetype": schema.StringAttribute{
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
				Description: "Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header whose value must be set to the IP address of the client.",
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
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
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
				Computed:    true,
				Description: "Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				// Not Computed: this is a disable-action input that the base service
				// GET never returns, so it can never be resolved after apply.
				Description: "Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the service.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with a service whose state transitions from UP to DOWN.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				// Not Computed: this is a disable-action input that the base service
				// GET never returns, so it can never be resolved after apply.
				Description: "Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.",
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
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service.",
			},
			"ip": schema.StringAttribute{
				Optional: true,
				// Not Computed: the ADC never returns "ip" in the service GET
				// (write-only, per SDK v2), so a Computed value can never be
				// resolved after apply. ForceNew parity preserved.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
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
				Description: "Maximum number of requests that can be sent on a persistent connection to the service.",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitornamesvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the specified service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name for the service. Cannot be changed after the service has been created.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile to use for the service.",
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Port number of the service.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service.",
			},
			"riseapbrstatsmsgcode": schema.Int64Attribute{
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
				Description: "Enable RTSP session ID mapping for the service.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name of the server that hosts the service.",
			},
			"servicetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
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
				Computed: true,
				// Derived from svrstate on read; keep the prior value when the user
				// does not configure it so an update never re-issues enable/disable
				// with an unknown/empty state (SDK v2 parity: state change was only
				// driven by an explicit config change).
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address when initiating a connection to the server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding.",
			},

			// Convenience blocks preserved from the SDK v2 resource.
			"lbvserver": schema.StringAttribute{
				Optional:    true,
				Description: "The name of the lb vserver to which the service is bound.",
			},
			"lbmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the lb monitor to bind to the service.",
			},
			"snienable": schema.StringAttribute{
				Optional: true,
				// Not Computed: the SSL convenience block is only read back when in
				// use; otherwise the base service GET never resolves it after apply.
				Description: "State of the Server Name Indication (SNI) feature on the service (SSL services only).",
			},
			"commonname": schema.StringAttribute{
				Optional: true,
				// Not Computed: the SSL convenience block is only read back when in
				// use; otherwise the base service GET never resolves it after apply.
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL service.",
			},
			"wait_until_disabled": schema.BoolAttribute{
				Optional: true,
				// Not Computed: this is a provider-side wait flag that the ADC never
				// returns, so it can never be resolved after apply.
				Description: "When set, the provider waits until the service reaches the DISABLED state before returning.",
			},
			"disabled_timeout": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum duration to wait for the service to reach the DISABLED state.",
			},
			"disabled_poll_delay": schema.StringAttribute{
				Optional:    true,
				Description: "Delay before the first poll while waiting for the DISABLED state.",
			},
			"disabled_poll_interval": schema.StringAttribute{
				Optional:    true,
				Description: "Interval between polls while waiting for the DISABLED state.",
			},
		},
	}
}

// serviceGetThePayloadFromthePlan builds the full add payload (used on Create),
// mirroring the attribute set the SDK v2 resource pushed on create. Action-only
// attributes (delay, graceful), convenience blocks, and read-only attributes
// (riseapbrstatsmsgcode) are intentionally excluded.
func serviceGetThePayloadFromthePlan(ctx context.Context, data *ServiceResourceModel) basic.Service {
	tflog.Debug(ctx, "In serviceGetThePayloadFromthePlan Function")

	svc := basic.Service{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		svc.Name = data.Name.ValueString()
	}
	if !data.Internal.IsNull() && !data.Internal.IsUnknown() {
		svc.Internal = data.Internal.ValueBool()
	}
	if !data.Accessdown.IsNull() && !data.Accessdown.IsUnknown() {
		svc.Accessdown = data.Accessdown.ValueString()
	}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		svc.All = data.All.ValueBool()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		svc.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cacheable.IsNull() && !data.Cacheable.IsUnknown() {
		svc.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cachetype.IsNull() && !data.Cachetype.IsUnknown() {
		svc.Cachetype = data.Cachetype.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		svc.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		svc.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() && !data.Cka.IsUnknown() {
		svc.Cka = data.Cka.ValueString()
	}
	if !data.Cleartextport.IsNull() && !data.Cleartextport.IsUnknown() {
		svc.Cleartextport = utils.IntPtr(int(data.Cleartextport.ValueInt64()))
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		svc.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() && !data.Cmp.IsUnknown() {
		svc.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		svc.Comment = data.Comment.ValueString()
	}
	if !data.Contentinspectionprofilename.IsNull() && !data.Contentinspectionprofilename.IsUnknown() {
		svc.Contentinspectionprofilename = data.Contentinspectionprofilename.ValueString()
	}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		svc.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dnsprofilename.IsNull() && !data.Dnsprofilename.IsUnknown() {
		svc.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		svc.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		svc.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		svc.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() && !data.Httpprofilename.IsUnknown() {
		svc.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		svc.Ip = data.Ip.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		svc.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		svc.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		svc.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		svc.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() && !data.Monconnectionclose.IsUnknown() {
		svc.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.Monitornamesvc.IsNull() && !data.Monitornamesvc.IsUnknown() {
		svc.Monitornamesvc = data.Monitornamesvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		svc.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Netprofile.IsNull() && !data.Netprofile.IsUnknown() {
		svc.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Pathmonitor.IsNull() && !data.Pathmonitor.IsUnknown() {
		svc.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() && !data.Pathmonitorindv.IsUnknown() {
		svc.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		svc.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Processlocal.IsNull() && !data.Processlocal.IsUnknown() {
		svc.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Quicprofilename.IsNull() && !data.Quicprofilename.IsUnknown() {
		svc.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Rtspsessionidremap.IsNull() && !data.Rtspsessionidremap.IsUnknown() {
		svc.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		svc.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		svc.Servername = data.Servername.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		svc.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sp.IsNull() && !data.Sp.IsUnknown() {
		svc.Sp = data.Sp.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		svc.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		svc.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() && !data.Tcpb.IsUnknown() {
		svc.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() && !data.Tcpprofilename.IsUnknown() {
		svc.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		svc.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		svc.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		svc.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		svc.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return svc
}

// serviceGetTheUpdatablePayloadFromThePlan builds the payload used on Update.
// ForceNew attributes (name is the identity, ip/cachetype/cleartextport/port/
// servername/servicetype/td/riseapbrstatsmsgcode trigger replacement) and
// action-only attributes (state via enable/disable, delay, graceful) are
// intentionally excluded so an in-place update never carries them.
func serviceGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *ServiceResourceModel) basic.Service {
	tflog.Debug(ctx, "In serviceGetTheUpdatablePayloadFromThePlan Function")

	svc := basic.Service{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		svc.Name = data.Name.ValueString()
	}
	if !data.Internal.IsNull() && !data.Internal.IsUnknown() {
		svc.Internal = data.Internal.ValueBool()
	}
	if !data.Accessdown.IsNull() && !data.Accessdown.IsUnknown() {
		svc.Accessdown = data.Accessdown.ValueString()
	}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		svc.All = data.All.ValueBool()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		svc.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cacheable.IsNull() && !data.Cacheable.IsUnknown() {
		svc.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		svc.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		svc.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Cka.IsNull() && !data.Cka.IsUnknown() {
		svc.Cka = data.Cka.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		svc.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cmp.IsNull() && !data.Cmp.IsUnknown() {
		svc.Cmp = data.Cmp.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		svc.Comment = data.Comment.ValueString()
	}
	if !data.Contentinspectionprofilename.IsNull() && !data.Contentinspectionprofilename.IsUnknown() {
		svc.Contentinspectionprofilename = data.Contentinspectionprofilename.ValueString()
	}
	if !data.Customserverid.IsNull() && !data.Customserverid.IsUnknown() {
		svc.Customserverid = data.Customserverid.ValueString()
	}
	if !data.Dnsprofilename.IsNull() && !data.Dnsprofilename.IsUnknown() {
		svc.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		svc.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		svc.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		svc.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Httpprofilename.IsNull() && !data.Httpprofilename.IsUnknown() {
		svc.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		svc.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		svc.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		svc.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Maxreq.IsNull() && !data.Maxreq.IsUnknown() {
		svc.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Monconnectionclose.IsNull() && !data.Monconnectionclose.IsUnknown() {
		svc.Monconnectionclose = data.Monconnectionclose.ValueString()
	}
	if !data.Monitornamesvc.IsNull() && !data.Monitornamesvc.IsUnknown() {
		svc.Monitornamesvc = data.Monitornamesvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		svc.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Netprofile.IsNull() && !data.Netprofile.IsUnknown() {
		svc.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Pathmonitor.IsNull() && !data.Pathmonitor.IsUnknown() {
		svc.Pathmonitor = data.Pathmonitor.ValueString()
	}
	if !data.Pathmonitorindv.IsNull() && !data.Pathmonitorindv.IsUnknown() {
		svc.Pathmonitorindv = data.Pathmonitorindv.ValueString()
	}
	if !data.Processlocal.IsNull() && !data.Processlocal.IsUnknown() {
		svc.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Quicprofilename.IsNull() && !data.Quicprofilename.IsUnknown() {
		svc.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Rtspsessionidremap.IsNull() && !data.Rtspsessionidremap.IsUnknown() {
		svc.Rtspsessionidremap = data.Rtspsessionidremap.ValueString()
	}
	if !data.Serverid.IsNull() && !data.Serverid.IsUnknown() {
		svc.Serverid = utils.IntPtr(int(data.Serverid.ValueInt64()))
	}
	if !data.Sp.IsNull() && !data.Sp.IsUnknown() {
		svc.Sp = data.Sp.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		svc.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Tcpb.IsNull() && !data.Tcpb.IsUnknown() {
		svc.Tcpb = data.Tcpb.ValueString()
	}
	if !data.Tcpprofilename.IsNull() && !data.Tcpprofilename.IsUnknown() {
		svc.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		svc.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		svc.Usip = data.Usip.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		svc.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return svc
}

// serviceSetStringFromGet reads a string attribute from the GET response, only
// nulling the model field when the current value is unknown (omit-on-default
// guard) so a known configured value NITRO omits is never clobbered.
func serviceSetStringFromGet(getResponseData map[string]interface{}, key string, current types.String) types.String {
	if val, ok := getResponseData[key]; ok && val != nil {
		if s, ok2 := val.(string); ok2 {
			return types.StringValue(s)
		}
	}
	if current.IsUnknown() {
		return types.StringNull()
	}
	return current
}

// serviceSetInt64FromGet reads an int attribute from the GET response, guarding
// the else-branch like serviceSetStringFromGet.
func serviceSetInt64FromGet(getResponseData map[string]interface{}, key string, current types.Int64) types.Int64 {
	if val, ok := getResponseData[key]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			return types.Int64Value(intVal)
		}
	}
	if current.IsUnknown() {
		return types.Int64Null()
	}
	return current
}

// serviceSetBoolFromGet reads a bool attribute from the GET response, guarding
// the else-branch like serviceSetStringFromGet.
func serviceSetBoolFromGet(getResponseData map[string]interface{}, key string, current types.Bool) types.Bool {
	if val, ok := getResponseData[key]; ok && val != nil {
		if b, ok2 := val.(bool); ok2 {
			return types.BoolValue(b)
		}
	}
	if current.IsUnknown() {
		return types.BoolNull()
	}
	return current
}

// serviceApplyGetToModel maps the raw service GET response onto the model,
// applying the SDK v2 read semantics. Convenience blocks (lbvserver, lbmonitor,
// snienable, commonname, wait_until_disabled, disabled_*), action-only fields
// (delay, graceful) and the write-only ip attribute are intentionally left
// untouched here — they are handled by the read helper or preserved from state.
func serviceApplyGetToModel(ctx context.Context, data *ServiceResourceModel, getResponseData map[string]interface{}) {
	tflog.Debug(ctx, "In serviceApplyGetToModel Function")

	data.Internal = serviceSetBoolFromGet(getResponseData, "Internal", data.Internal)
	data.All = serviceSetBoolFromGet(getResponseData, "all", data.All)
	data.Accessdown = serviceSetStringFromGet(getResponseData, "accessdown", data.Accessdown)
	data.Appflowlog = serviceSetStringFromGet(getResponseData, "appflowlog", data.Appflowlog)
	data.Cacheable = serviceSetStringFromGet(getResponseData, "cacheable", data.Cacheable)
	data.Cachetype = serviceSetStringFromGet(getResponseData, "cachetype", data.Cachetype)
	data.Cip = serviceSetStringFromGet(getResponseData, "cip", data.Cip)
	data.Cipheader = serviceSetStringFromGet(getResponseData, "cipheader", data.Cipheader)
	data.Cka = serviceSetStringFromGet(getResponseData, "cka", data.Cka)
	data.Cleartextport = serviceSetInt64FromGet(getResponseData, "cleartextport", data.Cleartextport)
	data.Clttimeout = serviceSetInt64FromGet(getResponseData, "clttimeout", data.Clttimeout)
	data.Cmp = serviceSetStringFromGet(getResponseData, "cmp", data.Cmp)
	data.Comment = serviceSetStringFromGet(getResponseData, "comment", data.Comment)
	data.Contentinspectionprofilename = serviceSetStringFromGet(getResponseData, "contentinspectionprofilename", data.Contentinspectionprofilename)
	data.Customserverid = serviceSetStringFromGet(getResponseData, "customserverid", data.Customserverid)
	data.Dnsprofilename = serviceSetStringFromGet(getResponseData, "dnsprofilename", data.Dnsprofilename)
	data.Downstateflush = serviceSetStringFromGet(getResponseData, "downstateflush", data.Downstateflush)
	data.Hashid = serviceSetInt64FromGet(getResponseData, "hashid", data.Hashid)
	data.Healthmonitor = serviceSetStringFromGet(getResponseData, "healthmonitor", data.Healthmonitor)
	data.Httpprofilename = serviceSetStringFromGet(getResponseData, "httpprofilename", data.Httpprofilename)
	// ip is not read back from GET (write-only on the ADC) - preserve state/plan.
	data.Ipaddress = serviceSetStringFromGet(getResponseData, "ipaddress", data.Ipaddress)
	data.Maxbandwidth = serviceSetInt64FromGet(getResponseData, "maxbandwidth", data.Maxbandwidth)
	data.Maxclient = serviceSetInt64FromGet(getResponseData, "maxclient", data.Maxclient)
	data.Maxreq = serviceSetInt64FromGet(getResponseData, "maxreq", data.Maxreq)
	data.Monconnectionclose = serviceSetStringFromGet(getResponseData, "monconnectionclose", data.Monconnectionclose)
	data.Monitornamesvc = serviceSetStringFromGet(getResponseData, "monitor_name_svc", data.Monitornamesvc)
	data.Monthreshold = serviceSetInt64FromGet(getResponseData, "monthreshold", data.Monthreshold)
	data.Name = serviceSetStringFromGet(getResponseData, "name", data.Name)
	data.Netprofile = serviceSetStringFromGet(getResponseData, "netprofile", data.Netprofile)
	data.Pathmonitor = serviceSetStringFromGet(getResponseData, "pathmonitor", data.Pathmonitor)
	data.Pathmonitorindv = serviceSetStringFromGet(getResponseData, "pathmonitorindv", data.Pathmonitorindv)
	data.Port = serviceSetInt64FromGet(getResponseData, "port", data.Port)
	data.Processlocal = serviceSetStringFromGet(getResponseData, "processlocal", data.Processlocal)
	data.Quicprofilename = serviceSetStringFromGet(getResponseData, "quicprofilename", data.Quicprofilename)
	data.Riseapbrstatsmsgcode = serviceSetInt64FromGet(getResponseData, "riseapbrstatsmsgcode", data.Riseapbrstatsmsgcode)
	data.Rtspsessionidremap = serviceSetStringFromGet(getResponseData, "rtspsessionidremap", data.Rtspsessionidremap)
	data.Serverid = serviceSetInt64FromGet(getResponseData, "serverid", data.Serverid)
	data.Servername = serviceSetStringFromGet(getResponseData, "servername", data.Servername)
	data.Servicetype = serviceSetStringFromGet(getResponseData, "servicetype", data.Servicetype)
	data.Svrtimeout = serviceSetInt64FromGet(getResponseData, "svrtimeout", data.Svrtimeout)
	data.Tcpb = serviceSetStringFromGet(getResponseData, "tcpb", data.Tcpb)
	data.Tcpprofilename = serviceSetStringFromGet(getResponseData, "tcpprofilename", data.Tcpprofilename)
	data.Td = serviceSetInt64FromGet(getResponseData, "td", data.Td)
	data.Useproxyport = serviceSetStringFromGet(getResponseData, "useproxyport", data.Useproxyport)
	data.Usip = serviceSetStringFromGet(getResponseData, "usip", data.Usip)
	data.Weight = serviceSetInt64FromGet(getResponseData, "weight", data.Weight)

	// sp: the ADC reports "ON (but effectively OFF)"; normalise to "ON" (SDK v2 parity).
	if val, ok := getResponseData["sp"]; ok && val != nil {
		if s, ok2 := val.(string); ok2 {
			if s == "ON (but effectively OFF)" {
				data.Sp = types.StringValue("ON")
			} else {
				data.Sp = types.StringValue(s)
			}
		}
	} else if data.Sp.IsUnknown() {
		data.Sp = types.StringNull()
	}

	// state: derive from svrstate (SDK v2 parity).
	if val, ok := getResponseData["svrstate"]; ok && val != nil {
		if s, ok2 := val.(string); ok2 {
			if s == "OUT OF SERVICE" {
				data.State = types.StringValue("DISABLED")
			} else {
				data.State = types.StringValue("ENABLED")
			}
		}
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
}

// serviceSetAttrFromGet applies the GET response to the resource model
// (preserving convenience/action-only fields) and sets the resource ID.
func serviceSetAttrFromGet(ctx context.Context, data *ServiceResourceModel, getResponseData map[string]interface{}) *ServiceResourceModel {
	tflog.Debug(ctx, "In serviceSetAttrFromGet Function")

	serviceApplyGetToModel(ctx, data, getResponseData)

	// The ID is the service name.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// serviceSetAttrFromGetForDatasource applies the GET response to the datasource
// model, copying every returned attribute and setting the ID.
func serviceSetAttrFromGetForDatasource(ctx context.Context, data *ServiceResourceModel, getResponseData map[string]interface{}) *ServiceResourceModel {
	tflog.Debug(ctx, "In serviceSetAttrFromGetForDatasource Function")

	serviceApplyGetToModel(ctx, data, getResponseData)

	// Convenience blocks are not part of the plain service GET.
	data.Lbvserver = types.StringNull()
	data.Lbmonitor = types.StringNull()
	data.Snienable = types.StringNull()
	data.Commonname = types.StringNull()
	data.WaitUntilDisabled = types.BoolNull()
	data.DisabledTimeout = types.StringNull()
	data.DisabledPollDelay = types.StringNull()
	data.DisabledPollInterval = types.StringNull()
	data.Ip = types.StringNull()
	data.Delay = types.Int64Null()
	data.Graceful = types.StringNull()

	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
