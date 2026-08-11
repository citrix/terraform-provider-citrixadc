package nstcpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstcpparamResourceModel describes the resource data model.
type NstcpparamResourceModel struct {
	Id                                  types.String `tfsdk:"id"`
	Ackonpush                           types.String `tfsdk:"ackonpush"`
	Autosyncookietimeout                types.Int64  `tfsdk:"autosyncookietimeout"`
	Compacttcpoptionnoop                types.String `tfsdk:"compacttcpoptionnoop"`
	Connflushifnomem                    types.String `tfsdk:"connflushifnomem"`
	Connflushthres                      types.Int64  `tfsdk:"connflushthres"`
	Delayedack                          types.Int64  `tfsdk:"delayedack"`
	Delinkclientserveronrst             types.String `tfsdk:"delinkclientserveronrst"`
	Downstaterst                        types.String `tfsdk:"downstaterst"`
	Enhancedisngeneration               types.String `tfsdk:"enhancedisngeneration"`
	Initialcwnd                         types.Int64  `tfsdk:"initialcwnd"`
	Kaprobeupdatelastactivity           types.String `tfsdk:"kaprobeupdatelastactivity"`
	Learnvsvrmss                        types.String `tfsdk:"learnvsvrmss"`
	Limitedpersist                      types.String `tfsdk:"limitedpersist"`
	Maxburst                            types.Int64  `tfsdk:"maxburst"`
	Maxdynserverprobes                  types.Int64  `tfsdk:"maxdynserverprobes"`
	Maxpktpermss                        types.Int64  `tfsdk:"maxpktpermss"`
	Maxsynackretx                       types.Int64  `tfsdk:"maxsynackretx"`
	Maxsynhold                          types.Int64  `tfsdk:"maxsynhold"`
	Maxsynholdperprobe                  types.Int64  `tfsdk:"maxsynholdperprobe"`
	Maxtimewaitconn                     types.Int64  `tfsdk:"maxtimewaitconn"`
	Minrto                              types.Int64  `tfsdk:"minrto"`
	Mptcpchecksum                       types.String `tfsdk:"mptcpchecksum"`
	Mptcpclosemptcpsessiononlastsfclose types.String `tfsdk:"mptcpclosemptcpsessiononlastsfclose"`
	Mptcpconcloseonpassivesf            types.String `tfsdk:"mptcpconcloseonpassivesf"`
	Mptcpfastcloseoption                types.String `tfsdk:"mptcpfastcloseoption"`
	Mptcpimmediatesfcloseonfin          types.String `tfsdk:"mptcpimmediatesfcloseonfin"`
	Mptcpmaxpendingsf                   types.Int64  `tfsdk:"mptcpmaxpendingsf"`
	Mptcpmaxsf                          types.Int64  `tfsdk:"mptcpmaxsf"`
	Mptcppendingjointhreshold           types.Int64  `tfsdk:"mptcppendingjointhreshold"`
	Mptcpreliableaddaddr                types.String `tfsdk:"mptcpreliableaddaddr"`
	Mptcprtostoswitchsf                 types.Int64  `tfsdk:"mptcprtostoswitchsf"`
	Mptcpsendsfresetoption              types.String `tfsdk:"mptcpsendsfresetoption"`
	Mptcpsfreplacetimeout               types.Int64  `tfsdk:"mptcpsfreplacetimeout"`
	Mptcpsftimeout                      types.Int64  `tfsdk:"mptcpsftimeout"`
	Mptcpusebackupondss                 types.String `tfsdk:"mptcpusebackupondss"`
	Msslearndelay                       types.Int64  `tfsdk:"msslearndelay"`
	Msslearninterval                    types.Int64  `tfsdk:"msslearninterval"`
	Nagle                               types.String `tfsdk:"nagle"`
	Oooqsize                            types.Int64  `tfsdk:"oooqsize"`
	Pktperretx                          types.Int64  `tfsdk:"pktperretx"`
	Recvbuffsize                        types.Int64  `tfsdk:"recvbuffsize"`
	Rfc5961chlgacklimit                 types.Int64  `tfsdk:"rfc5961chlgacklimit"`
	Sack                                types.String `tfsdk:"sack"`
	Slowstartincr                       types.Int64  `tfsdk:"slowstartincr"`
	Synattackdetection                  types.String `tfsdk:"synattackdetection"`
	Synholdfastgiveup                   types.Int64  `tfsdk:"synholdfastgiveup"`
	Tcpfastopencookietimeout            types.Int64  `tfsdk:"tcpfastopencookietimeout"`
	Tcpfintimeout                       types.Int64  `tfsdk:"tcpfintimeout"`
	Tcpmaxretries                       types.Int64  `tfsdk:"tcpmaxretries"`
	Ws                                  types.String `tfsdk:"ws"`
	Wsval                               types.Int64  `tfsdk:"wsval"`
}

func (r *NstcpparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// All configurable attributes were Optional+Computed+ForceNew in the SDK v2
	// resource. ForceNew is reproduced with RequiresReplaceIfConfigured (only a
	// configured change forces recreation), and UseStateForUnknown avoids
	// perpetual "known after apply" churn on the Computed values.
	stringPM := []planmodifier.String{
		stringplanmodifier.UseStateForUnknown(),
		stringplanmodifier.RequiresReplaceIfConfigured(),
	}
	int64PM := []planmodifier.Int64{
		int64planmodifier.UseStateForUnknown(),
		int64planmodifier.RequiresReplaceIfConfigured(),
	}

	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstcpparam resource.",
			},
			"ackonpush": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"autosyncookietimeout": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.",
			},
			"compacttcpoptionnoop": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "If enabled, non-negotiated TCP options are removed from the received packet while proxying it. By default, non-negotiated TCP options would be replaced by NOPs in the proxied packets. This option is not applicable for Citrix ADC generated packets.",
			},
			"connflushifnomem": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Flush an existing connection if no memory can be obtained for new connection.\n\nHALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time.  New connection fails if no such connection can be found.\n\nFIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active.  New connection fails if the oldest few connections are management connections.\n\nNote: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout.\n\nSee Also: connFlushThres argument below.",
			},
			"connflushthres": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established.  Note: This value may be rounded down to be a whole multiple of the number of packet engines running.",
			},
			"delayedack": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"delinkclientserveronrst": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "If enabled, Delink client and server connection, when there is outstanding data to be sent to the other side.",
			},
			"downstaterst": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Flag to switch on RST on down services.",
			},
			"enhancedisngeneration": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "If enabled, increase the ISN variation in SYN-ACKs sent by the NetScaler",
			},
			"initialcwnd": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.",
			},
			"kaprobeupdatelastactivity": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Update last activity for KA probes",
			},
			"learnvsvrmss": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Enable or disable maximum segment size (MSS) learning for virtual servers.",
			},
			"limitedpersist": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Limit the number of persist (zero window) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum number of TCP segments allowed in a burst.",
			},
			"maxdynserverprobes": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum number of probes that Citrix ADC can send out in 10 milliseconds, to dynamically learn a service. Citrix ADC probes for the existence of the origin in case of wildcard virtual server or services.",
			},
			"maxpktpermss": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum number of TCP packets allowed per maximum segment size (MSS).",
			},
			"maxsynackretx": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.",
			},
			"maxsynhold": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.",
			},
			"maxsynholdperprobe": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.",
			},
			"maxtimewaitconn": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum number of connections to hold in the TCP TIME_WAIT state on a packet engine. New connections entering TIME_WAIT state are proactively cleaned up.",
			},
			"minrto": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by 10).",
			},
			"mptcpchecksum": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Use MPTCP DSS checksum",
			},
			"mptcpclosemptcpsessiononlastsfclose": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow.",
			},
			"mptcpconcloseonpassivesf": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Accept DATA_FIN/FAST_CLOSE on passive subflow",
			},
			"mptcpfastcloseoption": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Allow to select option ACK or RESET to force the closure of an MPTCP connection abruptly.",
			},
			"mptcpimmediatesfcloseonfin": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level.",
			},
			"mptcpmaxpendingsf": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(4),
				PlanModifiers: int64PM,
				Description:   "Maximum number of subflow connections supported in pending join state per mptcp connection.",
			},
			"mptcpmaxsf": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum number of subflow connections supported in established state per mptcp connection.",
			},
			"mptcppendingjointhreshold": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(0),
				PlanModifiers: int64PM,
				Description:   "Maximum system level pending join connections allowed.",
			},
			"mptcpreliableaddaddr": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "If enabled, Citrix ADC retransmits MPTCP ADD-ADDR option if echo response is not received within the timeout interval. The retransmission is attempted only once.",
			},
			"mptcprtostoswitchsf": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Number of RTO's at subflow level, after which MPCTP should start using other subflow.",
			},
			"mptcpsendsfresetoption": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Allow MPTCP subflows to send TCP RST Reason (MP_TCPRST) Option while sending TCP RST.",
			},
			"mptcpsfreplacetimeout": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(10),
				PlanModifiers: int64PM,
				Description:   "The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction",
			},
			"mptcpsftimeout": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(0),
				PlanModifiers: int64PM,
				Description:   "The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver",
			},
			"mptcpusebackupondss": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable.",
			},
			"msslearndelay": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"msslearninterval": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"nagle": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"oooqsize": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(300),
				PlanModifiers: int64PM,
				Description:   "Maximum size of out-of-order packets queue. A value of 0 means no limit.",
			},
			"pktperretx": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.",
			},
			"recvbuffsize": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "TCP Receive buffer size",
			},
			"rfc5961chlgacklimit": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(0),
				PlanModifiers: int64PM,
				Description:   "Limits number of Challenge ACK sent per second, as recommended in RFC 5961(Improving TCP's Robustness to Blind In-Window Attacks)",
			},
			"sack": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Enable or disable Selective ACKnowledgement (SACK).",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"synattackdetection": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Detect TCP SYN packet flood and send an SNMP trap.",
			},
			"synholdfastgiveup": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Maximum threshold. After crossing this threshold number of outstanding probes for origin, the Citrix ADC reduces the number of connection retries for probe connections.",
			},
			"tcpfastopencookietimeout": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(0),
				PlanModifiers: int64PM,
				Description:   "Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.",
			},
			"tcpfintimeout": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed.",
			},
			"tcpmaxretries": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: int64PM,
				Description:   "Number of RTO's after which a connection should be freed.",
			},
			"ws": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: stringPM,
				Description:   "Enable or disable window scaling.",
			},
			"wsval": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Default:       int64default.StaticInt64(8),
				PlanModifiers: int64PM,
				Description:   "Factor used to calculate the new window size.\nThis argument is needed only when the window scaling is enabled.",
			},
		},
	}
}

func nstcpparamGetThePayloadFromtheConfig(ctx context.Context, data *NstcpparamResourceModel) ns.Nstcpparam {
	tflog.Debug(ctx, "In nstcpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model. Only push attributes that are
	// both known and configured. Unconfigured Optional+Computed attributes are
	// Unknown during Create (and resolved by the read-back), so they must be
	// skipped here to avoid pushing spurious 0/"" values to the ADC.
	nstcpparam := ns.Nstcpparam{}
	if !data.Ackonpush.IsNull() && !data.Ackonpush.IsUnknown() {
		nstcpparam.Ackonpush = data.Ackonpush.ValueString()
	}
	if !data.Autosyncookietimeout.IsNull() && !data.Autosyncookietimeout.IsUnknown() {
		nstcpparam.Autosyncookietimeout = utils.IntPtr(int(data.Autosyncookietimeout.ValueInt64()))
	}
	if !data.Compacttcpoptionnoop.IsNull() && !data.Compacttcpoptionnoop.IsUnknown() {
		nstcpparam.Compacttcpoptionnoop = data.Compacttcpoptionnoop.ValueString()
	}
	if !data.Connflushifnomem.IsNull() && !data.Connflushifnomem.IsUnknown() {
		nstcpparam.Connflushifnomem = data.Connflushifnomem.ValueString()
	}
	if !data.Connflushthres.IsNull() && !data.Connflushthres.IsUnknown() {
		nstcpparam.Connflushthres = utils.IntPtr(int(data.Connflushthres.ValueInt64()))
	}
	if !data.Delayedack.IsNull() && !data.Delayedack.IsUnknown() {
		nstcpparam.Delayedack = utils.IntPtr(int(data.Delayedack.ValueInt64()))
	}
	if !data.Delinkclientserveronrst.IsNull() && !data.Delinkclientserveronrst.IsUnknown() {
		nstcpparam.Delinkclientserveronrst = data.Delinkclientserveronrst.ValueString()
	}
	if !data.Downstaterst.IsNull() && !data.Downstaterst.IsUnknown() {
		nstcpparam.Downstaterst = data.Downstaterst.ValueString()
	}
	if !data.Enhancedisngeneration.IsNull() && !data.Enhancedisngeneration.IsUnknown() {
		nstcpparam.Enhancedisngeneration = data.Enhancedisngeneration.ValueString()
	}
	if !data.Initialcwnd.IsNull() && !data.Initialcwnd.IsUnknown() {
		nstcpparam.Initialcwnd = utils.IntPtr(int(data.Initialcwnd.ValueInt64()))
	}
	if !data.Kaprobeupdatelastactivity.IsNull() && !data.Kaprobeupdatelastactivity.IsUnknown() {
		nstcpparam.Kaprobeupdatelastactivity = data.Kaprobeupdatelastactivity.ValueString()
	}
	if !data.Learnvsvrmss.IsNull() && !data.Learnvsvrmss.IsUnknown() {
		nstcpparam.Learnvsvrmss = data.Learnvsvrmss.ValueString()
	}
	if !data.Limitedpersist.IsNull() && !data.Limitedpersist.IsUnknown() {
		nstcpparam.Limitedpersist = data.Limitedpersist.ValueString()
	}
	if !data.Maxburst.IsNull() && !data.Maxburst.IsUnknown() {
		nstcpparam.Maxburst = utils.IntPtr(int(data.Maxburst.ValueInt64()))
	}
	if !data.Maxdynserverprobes.IsNull() && !data.Maxdynserverprobes.IsUnknown() {
		nstcpparam.Maxdynserverprobes = utils.IntPtr(int(data.Maxdynserverprobes.ValueInt64()))
	}
	if !data.Maxpktpermss.IsNull() && !data.Maxpktpermss.IsUnknown() {
		nstcpparam.Maxpktpermss = utils.IntPtr(int(data.Maxpktpermss.ValueInt64()))
	}
	if !data.Maxsynackretx.IsNull() && !data.Maxsynackretx.IsUnknown() {
		nstcpparam.Maxsynackretx = utils.IntPtr(int(data.Maxsynackretx.ValueInt64()))
	}
	if !data.Maxsynhold.IsNull() && !data.Maxsynhold.IsUnknown() {
		nstcpparam.Maxsynhold = utils.IntPtr(int(data.Maxsynhold.ValueInt64()))
	}
	if !data.Maxsynholdperprobe.IsNull() && !data.Maxsynholdperprobe.IsUnknown() {
		nstcpparam.Maxsynholdperprobe = utils.IntPtr(int(data.Maxsynholdperprobe.ValueInt64()))
	}
	if !data.Maxtimewaitconn.IsNull() && !data.Maxtimewaitconn.IsUnknown() {
		nstcpparam.Maxtimewaitconn = utils.IntPtr(int(data.Maxtimewaitconn.ValueInt64()))
	}
	if !data.Minrto.IsNull() && !data.Minrto.IsUnknown() {
		nstcpparam.Minrto = utils.IntPtr(int(data.Minrto.ValueInt64()))
	}
	if !data.Mptcpchecksum.IsNull() && !data.Mptcpchecksum.IsUnknown() {
		nstcpparam.Mptcpchecksum = data.Mptcpchecksum.ValueString()
	}
	if !data.Mptcpclosemptcpsessiononlastsfclose.IsNull() && !data.Mptcpclosemptcpsessiononlastsfclose.IsUnknown() {
		nstcpparam.Mptcpclosemptcpsessiononlastsfclose = data.Mptcpclosemptcpsessiononlastsfclose.ValueString()
	}
	if !data.Mptcpconcloseonpassivesf.IsNull() && !data.Mptcpconcloseonpassivesf.IsUnknown() {
		nstcpparam.Mptcpconcloseonpassivesf = data.Mptcpconcloseonpassivesf.ValueString()
	}
	if !data.Mptcpfastcloseoption.IsNull() && !data.Mptcpfastcloseoption.IsUnknown() {
		nstcpparam.Mptcpfastcloseoption = data.Mptcpfastcloseoption.ValueString()
	}
	if !data.Mptcpimmediatesfcloseonfin.IsNull() && !data.Mptcpimmediatesfcloseonfin.IsUnknown() {
		nstcpparam.Mptcpimmediatesfcloseonfin = data.Mptcpimmediatesfcloseonfin.ValueString()
	}
	if !data.Mptcpmaxpendingsf.IsNull() && !data.Mptcpmaxpendingsf.IsUnknown() {
		nstcpparam.Mptcpmaxpendingsf = utils.IntPtr(int(data.Mptcpmaxpendingsf.ValueInt64()))
	}
	if !data.Mptcpmaxsf.IsNull() && !data.Mptcpmaxsf.IsUnknown() {
		nstcpparam.Mptcpmaxsf = utils.IntPtr(int(data.Mptcpmaxsf.ValueInt64()))
	}
	if !data.Mptcppendingjointhreshold.IsNull() && !data.Mptcppendingjointhreshold.IsUnknown() {
		nstcpparam.Mptcppendingjointhreshold = utils.IntPtr(int(data.Mptcppendingjointhreshold.ValueInt64()))
	}
	if !data.Mptcpreliableaddaddr.IsNull() && !data.Mptcpreliableaddaddr.IsUnknown() {
		nstcpparam.Mptcpreliableaddaddr = data.Mptcpreliableaddaddr.ValueString()
	}
	if !data.Mptcprtostoswitchsf.IsNull() && !data.Mptcprtostoswitchsf.IsUnknown() {
		nstcpparam.Mptcprtostoswitchsf = utils.IntPtr(int(data.Mptcprtostoswitchsf.ValueInt64()))
	}
	if !data.Mptcpsendsfresetoption.IsNull() && !data.Mptcpsendsfresetoption.IsUnknown() {
		nstcpparam.Mptcpsendsfresetoption = data.Mptcpsendsfresetoption.ValueString()
	}
	if !data.Mptcpsfreplacetimeout.IsNull() && !data.Mptcpsfreplacetimeout.IsUnknown() {
		nstcpparam.Mptcpsfreplacetimeout = utils.IntPtr(int(data.Mptcpsfreplacetimeout.ValueInt64()))
	}
	if !data.Mptcpsftimeout.IsNull() && !data.Mptcpsftimeout.IsUnknown() {
		nstcpparam.Mptcpsftimeout = utils.IntPtr(int(data.Mptcpsftimeout.ValueInt64()))
	}
	if !data.Mptcpusebackupondss.IsNull() && !data.Mptcpusebackupondss.IsUnknown() {
		nstcpparam.Mptcpusebackupondss = data.Mptcpusebackupondss.ValueString()
	}
	if !data.Msslearndelay.IsNull() && !data.Msslearndelay.IsUnknown() {
		nstcpparam.Msslearndelay = utils.IntPtr(int(data.Msslearndelay.ValueInt64()))
	}
	if !data.Msslearninterval.IsNull() && !data.Msslearninterval.IsUnknown() {
		nstcpparam.Msslearninterval = utils.IntPtr(int(data.Msslearninterval.ValueInt64()))
	}
	if !data.Nagle.IsNull() && !data.Nagle.IsUnknown() {
		nstcpparam.Nagle = data.Nagle.ValueString()
	}
	if !data.Oooqsize.IsNull() && !data.Oooqsize.IsUnknown() {
		nstcpparam.Oooqsize = utils.IntPtr(int(data.Oooqsize.ValueInt64()))
	}
	if !data.Pktperretx.IsNull() && !data.Pktperretx.IsUnknown() {
		nstcpparam.Pktperretx = utils.IntPtr(int(data.Pktperretx.ValueInt64()))
	}
	if !data.Recvbuffsize.IsNull() && !data.Recvbuffsize.IsUnknown() {
		nstcpparam.Recvbuffsize = utils.IntPtr(int(data.Recvbuffsize.ValueInt64()))
	}
	if !data.Rfc5961chlgacklimit.IsNull() && !data.Rfc5961chlgacklimit.IsUnknown() {
		nstcpparam.Rfc5961chlgacklimit = utils.IntPtr(int(data.Rfc5961chlgacklimit.ValueInt64()))
	}
	if !data.Sack.IsNull() && !data.Sack.IsUnknown() {
		nstcpparam.Sack = data.Sack.ValueString()
	}
	if !data.Slowstartincr.IsNull() && !data.Slowstartincr.IsUnknown() {
		nstcpparam.Slowstartincr = utils.IntPtr(int(data.Slowstartincr.ValueInt64()))
	}
	if !data.Synattackdetection.IsNull() && !data.Synattackdetection.IsUnknown() {
		nstcpparam.Synattackdetection = data.Synattackdetection.ValueString()
	}
	if !data.Synholdfastgiveup.IsNull() && !data.Synholdfastgiveup.IsUnknown() {
		nstcpparam.Synholdfastgiveup = utils.IntPtr(int(data.Synholdfastgiveup.ValueInt64()))
	}
	if !data.Tcpfastopencookietimeout.IsNull() && !data.Tcpfastopencookietimeout.IsUnknown() {
		nstcpparam.Tcpfastopencookietimeout = utils.IntPtr(int(data.Tcpfastopencookietimeout.ValueInt64()))
	}
	if !data.Tcpfintimeout.IsNull() && !data.Tcpfintimeout.IsUnknown() {
		nstcpparam.Tcpfintimeout = utils.IntPtr(int(data.Tcpfintimeout.ValueInt64()))
	}
	if !data.Tcpmaxretries.IsNull() && !data.Tcpmaxretries.IsUnknown() {
		nstcpparam.Tcpmaxretries = utils.IntPtr(int(data.Tcpmaxretries.ValueInt64()))
	}
	if !data.Ws.IsNull() && !data.Ws.IsUnknown() {
		nstcpparam.Ws = data.Ws.ValueString()
	}
	if !data.Wsval.IsNull() && !data.Wsval.IsUnknown() {
		nstcpparam.Wsval = utils.IntPtr(int(data.Wsval.ValueInt64()))
	}

	return nstcpparam
}

func nstcpparamSetAttrFromGet(ctx context.Context, data *NstcpparamResourceModel, getResponseData map[string]interface{}) *NstcpparamResourceModel {
	tflog.Debug(ctx, "In nstcpparamSetAttrFromGet Function")

	// Convert API response to model. The else-branches only null a value when it
	// is still Unknown, so a known configured value that NITRO omits from GET
	// (omit-on-default) is preserved instead of being clobbered.
	if val, ok := getResponseData["ackonpush"]; ok && val != nil {
		data.Ackonpush = types.StringValue(val.(string))
	} else if data.Ackonpush.IsUnknown() {
		data.Ackonpush = types.StringNull()
	}
	if val, ok := getResponseData["autosyncookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Autosyncookietimeout = types.Int64Value(intVal)
		}
	} else if data.Autosyncookietimeout.IsUnknown() {
		data.Autosyncookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["compacttcpoptionnoop"]; ok && val != nil {
		data.Compacttcpoptionnoop = types.StringValue(val.(string))
	} else if data.Compacttcpoptionnoop.IsUnknown() {
		data.Compacttcpoptionnoop = types.StringNull()
	}
	if val, ok := getResponseData["connflushifnomem"]; ok && val != nil {
		data.Connflushifnomem = types.StringValue(val.(string))
	} else if data.Connflushifnomem.IsUnknown() {
		data.Connflushifnomem = types.StringNull()
	}
	if val, ok := getResponseData["connflushthres"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Connflushthres = types.Int64Value(intVal)
		}
	} else if data.Connflushthres.IsUnknown() {
		data.Connflushthres = types.Int64Null()
	}
	if val, ok := getResponseData["delayedack"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delayedack = types.Int64Value(intVal)
		}
	} else if data.Delayedack.IsUnknown() {
		data.Delayedack = types.Int64Null()
	}
	if val, ok := getResponseData["delinkclientserveronrst"]; ok && val != nil {
		data.Delinkclientserveronrst = types.StringValue(val.(string))
	} else if data.Delinkclientserveronrst.IsUnknown() {
		data.Delinkclientserveronrst = types.StringNull()
	}
	if val, ok := getResponseData["downstaterst"]; ok && val != nil {
		data.Downstaterst = types.StringValue(val.(string))
	} else if data.Downstaterst.IsUnknown() {
		data.Downstaterst = types.StringNull()
	}
	if val, ok := getResponseData["enhancedisngeneration"]; ok && val != nil {
		data.Enhancedisngeneration = types.StringValue(val.(string))
	} else if data.Enhancedisngeneration.IsUnknown() {
		data.Enhancedisngeneration = types.StringNull()
	}
	if val, ok := getResponseData["initialcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialcwnd = types.Int64Value(intVal)
		}
	} else if data.Initialcwnd.IsUnknown() {
		data.Initialcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeupdatelastactivity"]; ok && val != nil {
		data.Kaprobeupdatelastactivity = types.StringValue(val.(string))
	} else if data.Kaprobeupdatelastactivity.IsUnknown() {
		data.Kaprobeupdatelastactivity = types.StringNull()
	}
	if val, ok := getResponseData["learnvsvrmss"]; ok && val != nil {
		data.Learnvsvrmss = types.StringValue(val.(string))
	} else if data.Learnvsvrmss.IsUnknown() {
		data.Learnvsvrmss = types.StringNull()
	}
	if val, ok := getResponseData["limitedpersist"]; ok && val != nil {
		data.Limitedpersist = types.StringValue(val.(string))
	} else if data.Limitedpersist.IsUnknown() {
		data.Limitedpersist = types.StringNull()
	}
	if val, ok := getResponseData["maxburst"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxburst = types.Int64Value(intVal)
		}
	} else if data.Maxburst.IsUnknown() {
		data.Maxburst = types.Int64Null()
	}
	if val, ok := getResponseData["maxdynserverprobes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxdynserverprobes = types.Int64Value(intVal)
		}
	} else if data.Maxdynserverprobes.IsUnknown() {
		data.Maxdynserverprobes = types.Int64Null()
	}
	if val, ok := getResponseData["maxpktpermss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpktpermss = types.Int64Value(intVal)
		}
	} else if data.Maxpktpermss.IsUnknown() {
		data.Maxpktpermss = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynackretx"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynackretx = types.Int64Value(intVal)
		}
	} else if data.Maxsynackretx.IsUnknown() {
		data.Maxsynackretx = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynhold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynhold = types.Int64Value(intVal)
		}
	} else if data.Maxsynhold.IsUnknown() {
		data.Maxsynhold = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynholdperprobe"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynholdperprobe = types.Int64Value(intVal)
		}
	} else if data.Maxsynholdperprobe.IsUnknown() {
		data.Maxsynholdperprobe = types.Int64Null()
	}
	if val, ok := getResponseData["maxtimewaitconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxtimewaitconn = types.Int64Value(intVal)
		}
	} else if data.Maxtimewaitconn.IsUnknown() {
		data.Maxtimewaitconn = types.Int64Null()
	}
	if val, ok := getResponseData["minrto"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrto = types.Int64Value(intVal)
		}
	} else if data.Minrto.IsUnknown() {
		data.Minrto = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpchecksum"]; ok && val != nil {
		data.Mptcpchecksum = types.StringValue(val.(string))
	} else if data.Mptcpchecksum.IsUnknown() {
		data.Mptcpchecksum = types.StringNull()
	}
	if val, ok := getResponseData["mptcpclosemptcpsessiononlastsfclose"]; ok && val != nil {
		data.Mptcpclosemptcpsessiononlastsfclose = types.StringValue(val.(string))
	} else if data.Mptcpclosemptcpsessiononlastsfclose.IsUnknown() {
		data.Mptcpclosemptcpsessiononlastsfclose = types.StringNull()
	}
	if val, ok := getResponseData["mptcpconcloseonpassivesf"]; ok && val != nil {
		data.Mptcpconcloseonpassivesf = types.StringValue(val.(string))
	} else if data.Mptcpconcloseonpassivesf.IsUnknown() {
		data.Mptcpconcloseonpassivesf = types.StringNull()
	}
	if val, ok := getResponseData["mptcpfastcloseoption"]; ok && val != nil {
		data.Mptcpfastcloseoption = types.StringValue(val.(string))
	} else if data.Mptcpfastcloseoption.IsUnknown() {
		data.Mptcpfastcloseoption = types.StringNull()
	}
	if val, ok := getResponseData["mptcpimmediatesfcloseonfin"]; ok && val != nil {
		data.Mptcpimmediatesfcloseonfin = types.StringValue(val.(string))
	} else if data.Mptcpimmediatesfcloseonfin.IsUnknown() {
		data.Mptcpimmediatesfcloseonfin = types.StringNull()
	}
	if val, ok := getResponseData["mptcpmaxpendingsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpmaxpendingsf = types.Int64Value(intVal)
		}
	} else if data.Mptcpmaxpendingsf.IsUnknown() {
		data.Mptcpmaxpendingsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpmaxsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpmaxsf = types.Int64Value(intVal)
		}
	} else if data.Mptcpmaxsf.IsUnknown() {
		data.Mptcpmaxsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcppendingjointhreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcppendingjointhreshold = types.Int64Value(intVal)
		}
	} else if data.Mptcppendingjointhreshold.IsUnknown() {
		data.Mptcppendingjointhreshold = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpreliableaddaddr"]; ok && val != nil {
		data.Mptcpreliableaddaddr = types.StringValue(val.(string))
	} else if data.Mptcpreliableaddaddr.IsUnknown() {
		data.Mptcpreliableaddaddr = types.StringNull()
	}
	if val, ok := getResponseData["mptcprtostoswitchsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcprtostoswitchsf = types.Int64Value(intVal)
		}
	} else if data.Mptcprtostoswitchsf.IsUnknown() {
		data.Mptcprtostoswitchsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpsendsfresetoption"]; ok && val != nil {
		data.Mptcpsendsfresetoption = types.StringValue(val.(string))
	} else if data.Mptcpsendsfresetoption.IsUnknown() {
		data.Mptcpsendsfresetoption = types.StringNull()
	}
	if val, ok := getResponseData["mptcpsfreplacetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsfreplacetimeout = types.Int64Value(intVal)
		}
	} else if data.Mptcpsfreplacetimeout.IsUnknown() {
		data.Mptcpsfreplacetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpsftimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsftimeout = types.Int64Value(intVal)
		}
	} else if data.Mptcpsftimeout.IsUnknown() {
		data.Mptcpsftimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpusebackupondss"]; ok && val != nil {
		data.Mptcpusebackupondss = types.StringValue(val.(string))
	} else if data.Mptcpusebackupondss.IsUnknown() {
		data.Mptcpusebackupondss = types.StringNull()
	}
	if val, ok := getResponseData["msslearndelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Msslearndelay = types.Int64Value(intVal)
		}
	} else if data.Msslearndelay.IsUnknown() {
		data.Msslearndelay = types.Int64Null()
	}
	if val, ok := getResponseData["msslearninterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Msslearninterval = types.Int64Value(intVal)
		}
	} else if data.Msslearninterval.IsUnknown() {
		data.Msslearninterval = types.Int64Null()
	}
	if val, ok := getResponseData["nagle"]; ok && val != nil {
		data.Nagle = types.StringValue(val.(string))
	} else if data.Nagle.IsUnknown() {
		data.Nagle = types.StringNull()
	}
	if val, ok := getResponseData["oooqsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Oooqsize = types.Int64Value(intVal)
		}
	} else if data.Oooqsize.IsUnknown() {
		data.Oooqsize = types.Int64Null()
	}
	if val, ok := getResponseData["pktperretx"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pktperretx = types.Int64Value(intVal)
		}
	} else if data.Pktperretx.IsUnknown() {
		data.Pktperretx = types.Int64Null()
	}
	if val, ok := getResponseData["recvbuffsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recvbuffsize = types.Int64Value(intVal)
		}
	} else if data.Recvbuffsize.IsUnknown() {
		data.Recvbuffsize = types.Int64Null()
	}
	if val, ok := getResponseData["rfc5961chlgacklimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rfc5961chlgacklimit = types.Int64Value(intVal)
		}
	} else if data.Rfc5961chlgacklimit.IsUnknown() {
		data.Rfc5961chlgacklimit = types.Int64Null()
	}
	if val, ok := getResponseData["sack"]; ok && val != nil {
		data.Sack = types.StringValue(val.(string))
	} else if data.Sack.IsUnknown() {
		data.Sack = types.StringNull()
	}
	if val, ok := getResponseData["slowstartincr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartincr = types.Int64Value(intVal)
		}
	} else if data.Slowstartincr.IsUnknown() {
		data.Slowstartincr = types.Int64Null()
	}
	if val, ok := getResponseData["synattackdetection"]; ok && val != nil {
		data.Synattackdetection = types.StringValue(val.(string))
	} else if data.Synattackdetection.IsUnknown() {
		data.Synattackdetection = types.StringNull()
	}
	if val, ok := getResponseData["synholdfastgiveup"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Synholdfastgiveup = types.Int64Value(intVal)
		}
	} else if data.Synholdfastgiveup.IsUnknown() {
		data.Synholdfastgiveup = types.Int64Null()
	}
	if val, ok := getResponseData["tcpfastopencookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfastopencookietimeout = types.Int64Value(intVal)
		}
	} else if data.Tcpfastopencookietimeout.IsUnknown() {
		data.Tcpfastopencookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpfintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfintimeout = types.Int64Value(intVal)
		}
	} else if data.Tcpfintimeout.IsUnknown() {
		data.Tcpfintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpmaxretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpmaxretries = types.Int64Value(intVal)
		}
	} else if data.Tcpmaxretries.IsUnknown() {
		data.Tcpmaxretries = types.Int64Null()
	}
	if val, ok := getResponseData["ws"]; ok && val != nil {
		data.Ws = types.StringValue(val.(string))
	} else if data.Ws.IsUnknown() {
		data.Ws = types.StringNull()
	}
	if val, ok := getResponseData["wsval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Wsval = types.Int64Value(intVal)
		}
	} else if data.Wsval.IsUnknown() {
		data.Wsval = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nstcpparam-config")

	return data
}
