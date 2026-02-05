package nstcpparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstcpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ackonpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"autosyncookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.",
			},
			"compacttcpoptionnoop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, non-negotiated TCP options are removed from the received packet while proxying it. By default, non-negotiated TCP options would be replaced by NOPs in the proxied packets. This option is not applicable for Citrix ADC generated packets.",
			},
			"connflushifnomem": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush an existing connection if no memory can be obtained for new connection.\n\nHALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time.  New connection fails if no such connection can be found.\n\nFIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active.  New connection fails if the oldest few connections are management connections.\n\nNote: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout.\n\nSee Also: connFlushThres argument below.",
			},
			"connflushthres": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established.  Note: This value may be rounded down to be a whole multiple of the number of packet engines running.",
			},
			"delayedack": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"delinkclientserveronrst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, Delink client and server connection, when there is outstanding data to be sent to the other side.",
			},
			"downstaterst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to switch on RST on down services.",
			},
			"enhancedisngeneration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, increase the ISN variation in SYN-ACKs sent by the NetScaler",
			},
			"initialcwnd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.",
			},
			"kaprobeupdatelastactivity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Update last activity for KA probes",
			},
			"learnvsvrmss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable maximum segment size (MSS) learning for virtual servers.",
			},
			"limitedpersist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit the number of persist (zero window) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP segments allowed in a burst.",
			},
			"maxdynserverprobes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of probes that Citrix ADC can send out in 10 milliseconds, to dynamically learn a service. Citrix ADC probes for the existence of the origin in case of wildcard virtual server or services.",
			},
			"maxpktpermss": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP packets allowed per maximum segment size (MSS).",
			},
			"maxsynackretx": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.",
			},
			"maxsynhold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.",
			},
			"maxsynholdperprobe": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.",
			},
			"maxtimewaitconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of connections to hold in the TCP TIME_WAIT state on a packet engine. New connections entering TIME_WAIT state are proactively cleaned up.",
			},
			"minrto": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by 10).",
			},
			"mptcpchecksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use MPTCP DSS checksum",
			},
			"mptcpclosemptcpsessiononlastsfclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow.",
			},
			"mptcpconcloseonpassivesf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept DATA_FIN/FAST_CLOSE on passive subflow",
			},
			"mptcpfastcloseoption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow to select option ACK or RESET to force the closure of an MPTCP connection abruptly.",
			},
			"mptcpimmediatesfcloseonfin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level.",
			},
			"mptcpmaxpendingsf": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of subflow connections supported in pending join state per mptcp connection.",
			},
			"mptcpmaxsf": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of subflow connections supported in established state per mptcp connection.",
			},
			"mptcppendingjointhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum system level pending join connections allowed.",
			},
			"mptcpreliableaddaddr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, Citrix ADC retransmits MPTCP ADD-ADDR option if echo response is not received within the timeout interval. The retransmission is attempted only once.",
			},
			"mptcprtostoswitchsf": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of RTO's at subflow level, after which MPCTP should start using other subflow.",
			},
			"mptcpsendsfresetoption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow MPTCP subflows to send TCP RST Reason (MP_TCPRST) Option while sending TCP RST.",
			},
			"mptcpsfreplacetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction",
			},
			"mptcpsftimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver",
			},
			"mptcpusebackupondss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable.",
			},
			"msslearndelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"msslearninterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.",
			},
			"nagle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"oooqsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of out-of-order packets queue. A value of 0 means no limit.",
			},
			"pktperretx": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.",
			},
			"recvbuffsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Receive buffer size",
			},
			"rfc5961chlgacklimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limits number of Challenge ACK sent per second, as recommended in RFC 5961(Improving TCP's Robustness to Blind In-Window Attacks)",
			},
			"sack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Selective ACKnowledgement (SACK).",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"synattackdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Detect TCP SYN packet flood and send an SNMP trap.",
			},
			"synholdfastgiveup": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum threshold. After crossing this threshold number of outstanding probes for origin, the Citrix ADC reduces the number of connection retries for probe connections.",
			},
			"tcpfastopencookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.",
			},
			"tcpfintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed.",
			},
			"tcpmaxretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of RTO's after which a connection should be freed.",
			},
			"ws": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable window scaling.",
			},
			"wsval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Factor used to calculate the new window size.\nThis argument is needed only when the window scaling is enabled.",
			},
		},
	}
}
