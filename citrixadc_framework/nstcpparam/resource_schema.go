package nstcpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstcpparam resource.",
			},
			"ackonpush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"autosyncookietimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.",
			},
			"compacttcpoptionnoop": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, non-negotiated TCP options are removed from the received packet while proxying it. By default, non-negotiated TCP options would be replaced by NOPs in the proxied packets. This option is not applicable for Citrix ADC generated packets.",
			},
			"connflushifnomem": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NSA_CONNFLUSH_NONE"),
				Description: "Flush an existing connection if no memory can be obtained for new connection.\n\nHALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time.  New connection fails if no such connection can be found.\n\nFIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active.  New connection fails if the oldest few connections are management connections.\n\nNote: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout.\n\nSee Also: connFlushThres argument below.",
			},
			"connflushthres": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established.  Note: This value may be rounded down to be a whole multiple of the number of packet engines running.",
			},
			"delayedack": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"delinkclientserveronrst": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, Delink client and server connection, when there is outstanding data to be sent to the other side.",
			},
			"downstaterst": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Flag to switch on RST on down services.",
			},
			"enhancedisngeneration": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, increase the ISN variation in SYN-ACKs sent by the NetScaler",
			},
			"initialcwnd": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.",
			},
			"kaprobeupdatelastactivity": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Update last activity for KA probes",
			},
			"learnvsvrmss": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable maximum segment size (MSS) learning for virtual servers.",
			},
			"limitedpersist": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Limit the number of persist (zero window) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(6),
				Description: "Maximum number of TCP segments allowed in a burst.",
			},
			"maxdynserverprobes": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(7),
				Description: "Maximum number of probes that Citrix ADC can send out in 10 milliseconds, to dynamically learn a service. Citrix ADC probes for the existence of the origin in case of wildcard virtual server or services.",
			},
			"maxpktpermss": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP packets allowed per maximum segment size (MSS).",
			},
			"maxsynackretx": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.",
			},
			"maxsynhold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(16384),
				Description: "Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.",
			},
			"maxsynholdperprobe": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.",
			},
			"maxtimewaitconn": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(7000),
				Description: "Maximum number of connections to hold in the TCP TIME_WAIT state on a packet engine. New connections entering TIME_WAIT state are proactively cleaned up.",
			},
			"minrto": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by 10).",
			},
			"mptcpchecksum": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Use MPTCP DSS checksum",
			},
			"mptcpclosemptcpsessiononlastsfclose": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow.",
			},
			"mptcpconcloseonpassivesf": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Accept DATA_FIN/FAST_CLOSE on passive subflow",
			},
			"mptcpfastcloseoption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ACK"),
				Description: "Allow to select option ACK or RESET to force the closure of an MPTCP connection abruptly.",
			},
			"mptcpimmediatesfcloseonfin": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level.",
			},
			"mptcpmaxpendingsf": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Maximum number of subflow connections supported in pending join state per mptcp connection.",
			},
			"mptcpmaxsf": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Maximum number of subflow connections supported in established state per mptcp connection.",
			},
			"mptcppendingjointhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum system level pending join connections allowed.",
			},
			"mptcpreliableaddaddr": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, Citrix ADC retransmits MPTCP ADD-ADDR option if echo response is not received within the timeout interval. The retransmission is attempted only once.",
			},
			"mptcprtostoswitchsf": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Number of RTO's at subflow level, after which MPCTP should start using other subflow.",
			},
			"mptcpsendsfresetoption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow MPTCP subflows to send TCP RST Reason (MP_TCPRST) Option while sending TCP RST.",
			},
			"mptcpsfreplacetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction",
			},
			"mptcpsftimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver",
			},
			"mptcpusebackupondss": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable.",
			},
			"msslearndelay": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"msslearninterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(180),
				Description: "Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"nagle": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"oooqsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Maximum size of out-of-order packets queue. A value of 0 means no limit.",
			},
			"pktperretx": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.",
			},
			"recvbuffsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8190),
				Description: "TCP Receive buffer size",
			},
			"rfc5961chlgacklimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limits number of Challenge ACK sent per second, as recommended in RFC 5961(Improving TCP's Robustness to Blind In-Window Attacks)",
			},
			"sack": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable Selective ACKnowledgement (SACK).",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"synattackdetection": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Detect TCP SYN packet flood and send an SNMP trap.",
			},
			"synholdfastgiveup": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1024),
				Description: "Maximum threshold. After crossing this threshold number of outstanding probes for origin, the Citrix ADC reduces the number of connection retries for probe connections.",
			},
			"tcpfastopencookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.",
			},
			"tcpfintimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(40),
				Description: "The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed.",
			},
			"tcpmaxretries": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(7),
				Description: "Number of RTO's after which a connection should be freed.",
			},
			"ws": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable window scaling.",
			},
			"wsval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8),
				Description: "Factor used to calculate the new window size.\nThis argument is needed only when the window scaling is enabled.",
			},
		},
	}
}

func nstcpparamGetThePayloadFromtheConfig(ctx context.Context, data *NstcpparamResourceModel) ns.Nstcpparam {
	tflog.Debug(ctx, "In nstcpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstcpparam := ns.Nstcpparam{}
	if !data.Ackonpush.IsNull() {
		nstcpparam.Ackonpush = data.Ackonpush.ValueString()
	}
	if !data.Autosyncookietimeout.IsNull() {
		nstcpparam.Autosyncookietimeout = utils.IntPtr(int(data.Autosyncookietimeout.ValueInt64()))
	}
	if !data.Compacttcpoptionnoop.IsNull() {
		nstcpparam.Compacttcpoptionnoop = data.Compacttcpoptionnoop.ValueString()
	}
	if !data.Connflushifnomem.IsNull() {
		nstcpparam.Connflushifnomem = data.Connflushifnomem.ValueString()
	}
	if !data.Connflushthres.IsNull() {
		nstcpparam.Connflushthres = utils.IntPtr(int(data.Connflushthres.ValueInt64()))
	}
	if !data.Delayedack.IsNull() {
		nstcpparam.Delayedack = utils.IntPtr(int(data.Delayedack.ValueInt64()))
	}
	if !data.Delinkclientserveronrst.IsNull() {
		nstcpparam.Delinkclientserveronrst = data.Delinkclientserveronrst.ValueString()
	}
	if !data.Downstaterst.IsNull() {
		nstcpparam.Downstaterst = data.Downstaterst.ValueString()
	}
	if !data.Enhancedisngeneration.IsNull() {
		nstcpparam.Enhancedisngeneration = data.Enhancedisngeneration.ValueString()
	}
	if !data.Initialcwnd.IsNull() {
		nstcpparam.Initialcwnd = utils.IntPtr(int(data.Initialcwnd.ValueInt64()))
	}
	if !data.Kaprobeupdatelastactivity.IsNull() {
		nstcpparam.Kaprobeupdatelastactivity = data.Kaprobeupdatelastactivity.ValueString()
	}
	if !data.Learnvsvrmss.IsNull() {
		nstcpparam.Learnvsvrmss = data.Learnvsvrmss.ValueString()
	}
	if !data.Limitedpersist.IsNull() {
		nstcpparam.Limitedpersist = data.Limitedpersist.ValueString()
	}
	if !data.Maxburst.IsNull() {
		nstcpparam.Maxburst = utils.IntPtr(int(data.Maxburst.ValueInt64()))
	}
	if !data.Maxdynserverprobes.IsNull() {
		nstcpparam.Maxdynserverprobes = utils.IntPtr(int(data.Maxdynserverprobes.ValueInt64()))
	}
	if !data.Maxpktpermss.IsNull() {
		nstcpparam.Maxpktpermss = utils.IntPtr(int(data.Maxpktpermss.ValueInt64()))
	}
	if !data.Maxsynackretx.IsNull() {
		nstcpparam.Maxsynackretx = utils.IntPtr(int(data.Maxsynackretx.ValueInt64()))
	}
	if !data.Maxsynhold.IsNull() {
		nstcpparam.Maxsynhold = utils.IntPtr(int(data.Maxsynhold.ValueInt64()))
	}
	if !data.Maxsynholdperprobe.IsNull() {
		nstcpparam.Maxsynholdperprobe = utils.IntPtr(int(data.Maxsynholdperprobe.ValueInt64()))
	}
	if !data.Maxtimewaitconn.IsNull() {
		nstcpparam.Maxtimewaitconn = utils.IntPtr(int(data.Maxtimewaitconn.ValueInt64()))
	}
	if !data.Minrto.IsNull() {
		nstcpparam.Minrto = utils.IntPtr(int(data.Minrto.ValueInt64()))
	}
	if !data.Mptcpchecksum.IsNull() {
		nstcpparam.Mptcpchecksum = data.Mptcpchecksum.ValueString()
	}
	if !data.Mptcpclosemptcpsessiononlastsfclose.IsNull() {
		nstcpparam.Mptcpclosemptcpsessiononlastsfclose = data.Mptcpclosemptcpsessiononlastsfclose.ValueString()
	}
	if !data.Mptcpconcloseonpassivesf.IsNull() {
		nstcpparam.Mptcpconcloseonpassivesf = data.Mptcpconcloseonpassivesf.ValueString()
	}
	if !data.Mptcpfastcloseoption.IsNull() {
		nstcpparam.Mptcpfastcloseoption = data.Mptcpfastcloseoption.ValueString()
	}
	if !data.Mptcpimmediatesfcloseonfin.IsNull() {
		nstcpparam.Mptcpimmediatesfcloseonfin = data.Mptcpimmediatesfcloseonfin.ValueString()
	}
	if !data.Mptcpmaxpendingsf.IsNull() {
		nstcpparam.Mptcpmaxpendingsf = utils.IntPtr(int(data.Mptcpmaxpendingsf.ValueInt64()))
	}
	if !data.Mptcpmaxsf.IsNull() {
		nstcpparam.Mptcpmaxsf = utils.IntPtr(int(data.Mptcpmaxsf.ValueInt64()))
	}
	if !data.Mptcppendingjointhreshold.IsNull() {
		nstcpparam.Mptcppendingjointhreshold = utils.IntPtr(int(data.Mptcppendingjointhreshold.ValueInt64()))
	}
	if !data.Mptcpreliableaddaddr.IsNull() {
		nstcpparam.Mptcpreliableaddaddr = data.Mptcpreliableaddaddr.ValueString()
	}
	if !data.Mptcprtostoswitchsf.IsNull() {
		nstcpparam.Mptcprtostoswitchsf = utils.IntPtr(int(data.Mptcprtostoswitchsf.ValueInt64()))
	}
	if !data.Mptcpsendsfresetoption.IsNull() {
		nstcpparam.Mptcpsendsfresetoption = data.Mptcpsendsfresetoption.ValueString()
	}
	if !data.Mptcpsfreplacetimeout.IsNull() {
		nstcpparam.Mptcpsfreplacetimeout = utils.IntPtr(int(data.Mptcpsfreplacetimeout.ValueInt64()))
	}
	if !data.Mptcpsftimeout.IsNull() {
		nstcpparam.Mptcpsftimeout = utils.IntPtr(int(data.Mptcpsftimeout.ValueInt64()))
	}
	if !data.Mptcpusebackupondss.IsNull() {
		nstcpparam.Mptcpusebackupondss = data.Mptcpusebackupondss.ValueString()
	}
	if !data.Msslearndelay.IsNull() {
		nstcpparam.Msslearndelay = utils.IntPtr(int(data.Msslearndelay.ValueInt64()))
	}
	if !data.Msslearninterval.IsNull() {
		nstcpparam.Msslearninterval = utils.IntPtr(int(data.Msslearninterval.ValueInt64()))
	}
	if !data.Nagle.IsNull() {
		nstcpparam.Nagle = data.Nagle.ValueString()
	}
	if !data.Oooqsize.IsNull() {
		nstcpparam.Oooqsize = utils.IntPtr(int(data.Oooqsize.ValueInt64()))
	}
	if !data.Pktperretx.IsNull() {
		nstcpparam.Pktperretx = utils.IntPtr(int(data.Pktperretx.ValueInt64()))
	}
	if !data.Recvbuffsize.IsNull() {
		nstcpparam.Recvbuffsize = utils.IntPtr(int(data.Recvbuffsize.ValueInt64()))
	}
	if !data.Rfc5961chlgacklimit.IsNull() {
		nstcpparam.Rfc5961chlgacklimit = utils.IntPtr(int(data.Rfc5961chlgacklimit.ValueInt64()))
	}
	if !data.Sack.IsNull() {
		nstcpparam.Sack = data.Sack.ValueString()
	}
	if !data.Slowstartincr.IsNull() {
		nstcpparam.Slowstartincr = utils.IntPtr(int(data.Slowstartincr.ValueInt64()))
	}
	if !data.Synattackdetection.IsNull() {
		nstcpparam.Synattackdetection = data.Synattackdetection.ValueString()
	}
	if !data.Synholdfastgiveup.IsNull() {
		nstcpparam.Synholdfastgiveup = utils.IntPtr(int(data.Synholdfastgiveup.ValueInt64()))
	}
	if !data.Tcpfastopencookietimeout.IsNull() {
		nstcpparam.Tcpfastopencookietimeout = utils.IntPtr(int(data.Tcpfastopencookietimeout.ValueInt64()))
	}
	if !data.Tcpfintimeout.IsNull() {
		nstcpparam.Tcpfintimeout = utils.IntPtr(int(data.Tcpfintimeout.ValueInt64()))
	}
	if !data.Tcpmaxretries.IsNull() {
		nstcpparam.Tcpmaxretries = utils.IntPtr(int(data.Tcpmaxretries.ValueInt64()))
	}
	if !data.Ws.IsNull() {
		nstcpparam.Ws = data.Ws.ValueString()
	}
	if !data.Wsval.IsNull() {
		nstcpparam.Wsval = utils.IntPtr(int(data.Wsval.ValueInt64()))
	}

	return nstcpparam
}

func nstcpparamSetAttrFromGet(ctx context.Context, data *NstcpparamResourceModel, getResponseData map[string]interface{}) *NstcpparamResourceModel {
	tflog.Debug(ctx, "In nstcpparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ackonpush"]; ok && val != nil {
		data.Ackonpush = types.StringValue(val.(string))
	} else {
		data.Ackonpush = types.StringNull()
	}
	if val, ok := getResponseData["autosyncookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Autosyncookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Autosyncookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["compacttcpoptionnoop"]; ok && val != nil {
		data.Compacttcpoptionnoop = types.StringValue(val.(string))
	} else {
		data.Compacttcpoptionnoop = types.StringNull()
	}
	if val, ok := getResponseData["connflushifnomem"]; ok && val != nil {
		data.Connflushifnomem = types.StringValue(val.(string))
	} else {
		data.Connflushifnomem = types.StringNull()
	}
	if val, ok := getResponseData["connflushthres"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Connflushthres = types.Int64Value(intVal)
		}
	} else {
		data.Connflushthres = types.Int64Null()
	}
	if val, ok := getResponseData["delayedack"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delayedack = types.Int64Value(intVal)
		}
	} else {
		data.Delayedack = types.Int64Null()
	}
	if val, ok := getResponseData["delinkclientserveronrst"]; ok && val != nil {
		data.Delinkclientserveronrst = types.StringValue(val.(string))
	} else {
		data.Delinkclientserveronrst = types.StringNull()
	}
	if val, ok := getResponseData["downstaterst"]; ok && val != nil {
		data.Downstaterst = types.StringValue(val.(string))
	} else {
		data.Downstaterst = types.StringNull()
	}
	if val, ok := getResponseData["enhancedisngeneration"]; ok && val != nil {
		data.Enhancedisngeneration = types.StringValue(val.(string))
	} else {
		data.Enhancedisngeneration = types.StringNull()
	}
	if val, ok := getResponseData["initialcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialcwnd = types.Int64Value(intVal)
		}
	} else {
		data.Initialcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeupdatelastactivity"]; ok && val != nil {
		data.Kaprobeupdatelastactivity = types.StringValue(val.(string))
	} else {
		data.Kaprobeupdatelastactivity = types.StringNull()
	}
	if val, ok := getResponseData["learnvsvrmss"]; ok && val != nil {
		data.Learnvsvrmss = types.StringValue(val.(string))
	} else {
		data.Learnvsvrmss = types.StringNull()
	}
	if val, ok := getResponseData["limitedpersist"]; ok && val != nil {
		data.Limitedpersist = types.StringValue(val.(string))
	} else {
		data.Limitedpersist = types.StringNull()
	}
	if val, ok := getResponseData["maxburst"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxburst = types.Int64Value(intVal)
		}
	} else {
		data.Maxburst = types.Int64Null()
	}
	if val, ok := getResponseData["maxdynserverprobes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxdynserverprobes = types.Int64Value(intVal)
		}
	} else {
		data.Maxdynserverprobes = types.Int64Null()
	}
	if val, ok := getResponseData["maxpktpermss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpktpermss = types.Int64Value(intVal)
		}
	} else {
		data.Maxpktpermss = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynackretx"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynackretx = types.Int64Value(intVal)
		}
	} else {
		data.Maxsynackretx = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynhold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynhold = types.Int64Value(intVal)
		}
	} else {
		data.Maxsynhold = types.Int64Null()
	}
	if val, ok := getResponseData["maxsynholdperprobe"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsynholdperprobe = types.Int64Value(intVal)
		}
	} else {
		data.Maxsynholdperprobe = types.Int64Null()
	}
	if val, ok := getResponseData["maxtimewaitconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxtimewaitconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxtimewaitconn = types.Int64Null()
	}
	if val, ok := getResponseData["minrto"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrto = types.Int64Value(intVal)
		}
	} else {
		data.Minrto = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpchecksum"]; ok && val != nil {
		data.Mptcpchecksum = types.StringValue(val.(string))
	} else {
		data.Mptcpchecksum = types.StringNull()
	}
	if val, ok := getResponseData["mptcpclosemptcpsessiononlastsfclose"]; ok && val != nil {
		data.Mptcpclosemptcpsessiononlastsfclose = types.StringValue(val.(string))
	} else {
		data.Mptcpclosemptcpsessiononlastsfclose = types.StringNull()
	}
	if val, ok := getResponseData["mptcpconcloseonpassivesf"]; ok && val != nil {
		data.Mptcpconcloseonpassivesf = types.StringValue(val.(string))
	} else {
		data.Mptcpconcloseonpassivesf = types.StringNull()
	}
	if val, ok := getResponseData["mptcpfastcloseoption"]; ok && val != nil {
		data.Mptcpfastcloseoption = types.StringValue(val.(string))
	} else {
		data.Mptcpfastcloseoption = types.StringNull()
	}
	if val, ok := getResponseData["mptcpimmediatesfcloseonfin"]; ok && val != nil {
		data.Mptcpimmediatesfcloseonfin = types.StringValue(val.(string))
	} else {
		data.Mptcpimmediatesfcloseonfin = types.StringNull()
	}
	if val, ok := getResponseData["mptcpmaxpendingsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpmaxpendingsf = types.Int64Value(intVal)
		}
	} else {
		data.Mptcpmaxpendingsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpmaxsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpmaxsf = types.Int64Value(intVal)
		}
	} else {
		data.Mptcpmaxsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcppendingjointhreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcppendingjointhreshold = types.Int64Value(intVal)
		}
	} else {
		data.Mptcppendingjointhreshold = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpreliableaddaddr"]; ok && val != nil {
		data.Mptcpreliableaddaddr = types.StringValue(val.(string))
	} else {
		data.Mptcpreliableaddaddr = types.StringNull()
	}
	if val, ok := getResponseData["mptcprtostoswitchsf"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcprtostoswitchsf = types.Int64Value(intVal)
		}
	} else {
		data.Mptcprtostoswitchsf = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpsendsfresetoption"]; ok && val != nil {
		data.Mptcpsendsfresetoption = types.StringValue(val.(string))
	} else {
		data.Mptcpsendsfresetoption = types.StringNull()
	}
	if val, ok := getResponseData["mptcpsfreplacetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsfreplacetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Mptcpsfreplacetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpsftimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsftimeout = types.Int64Value(intVal)
		}
	} else {
		data.Mptcpsftimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mptcpusebackupondss"]; ok && val != nil {
		data.Mptcpusebackupondss = types.StringValue(val.(string))
	} else {
		data.Mptcpusebackupondss = types.StringNull()
	}
	if val, ok := getResponseData["msslearndelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Msslearndelay = types.Int64Value(intVal)
		}
	} else {
		data.Msslearndelay = types.Int64Null()
	}
	if val, ok := getResponseData["msslearninterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Msslearninterval = types.Int64Value(intVal)
		}
	} else {
		data.Msslearninterval = types.Int64Null()
	}
	if val, ok := getResponseData["nagle"]; ok && val != nil {
		data.Nagle = types.StringValue(val.(string))
	} else {
		data.Nagle = types.StringNull()
	}
	if val, ok := getResponseData["oooqsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Oooqsize = types.Int64Value(intVal)
		}
	} else {
		data.Oooqsize = types.Int64Null()
	}
	if val, ok := getResponseData["pktperretx"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pktperretx = types.Int64Value(intVal)
		}
	} else {
		data.Pktperretx = types.Int64Null()
	}
	if val, ok := getResponseData["recvbuffsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recvbuffsize = types.Int64Value(intVal)
		}
	} else {
		data.Recvbuffsize = types.Int64Null()
	}
	if val, ok := getResponseData["rfc5961chlgacklimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rfc5961chlgacklimit = types.Int64Value(intVal)
		}
	} else {
		data.Rfc5961chlgacklimit = types.Int64Null()
	}
	if val, ok := getResponseData["sack"]; ok && val != nil {
		data.Sack = types.StringValue(val.(string))
	} else {
		data.Sack = types.StringNull()
	}
	if val, ok := getResponseData["slowstartincr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartincr = types.Int64Value(intVal)
		}
	} else {
		data.Slowstartincr = types.Int64Null()
	}
	if val, ok := getResponseData["synattackdetection"]; ok && val != nil {
		data.Synattackdetection = types.StringValue(val.(string))
	} else {
		data.Synattackdetection = types.StringNull()
	}
	if val, ok := getResponseData["synholdfastgiveup"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Synholdfastgiveup = types.Int64Value(intVal)
		}
	} else {
		data.Synholdfastgiveup = types.Int64Null()
	}
	if val, ok := getResponseData["tcpfastopencookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfastopencookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Tcpfastopencookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpfintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfintimeout = types.Int64Value(intVal)
		}
	} else {
		data.Tcpfintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpmaxretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpmaxretries = types.Int64Value(intVal)
		}
	} else {
		data.Tcpmaxretries = types.Int64Null()
	}
	if val, ok := getResponseData["ws"]; ok && val != nil {
		data.Ws = types.StringValue(val.(string))
	} else {
		data.Ws = types.StringNull()
	}
	if val, ok := getResponseData["wsval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Wsval = types.Int64Value(intVal)
		}
	} else {
		data.Wsval = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nstcpparam-config")

	return data
}
