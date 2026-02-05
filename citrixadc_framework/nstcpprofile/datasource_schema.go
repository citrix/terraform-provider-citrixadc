package nstcpprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NstcpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ackaggregation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable ACK Aggregation.",
			},
			"ackonpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"applyadaptivetcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Apply Adaptive TCP optimizations",
			},
			"buffersize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP buffering size, in bytes.",
			},
			"burstratecontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set.",
			},
			"clientiptcpoption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client IP in TCP options",
			},
			"clientiptcpoptionnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ClientIP TCP Option number",
			},
			"delayedack": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"dropestconnontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Silently drop tcp established connections on idle timeout",
			},
			"drophalfclosedconnontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Silently drop tcp half closed connections on idle timeout",
			},
			"dsack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable DSACK.",
			},
			"dupackthresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP dupack threshold.",
			},
			"dynamicreceivebuffering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions.\nNote: The buffer size argument must be set for dynamic adjustments to take place.",
			},
			"ecn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable TCP Explicit Congestion Notification.",
			},
			"establishclientconn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Establishing Client Client connection on First data/ Final-ACK / Automatic",
			},
			"fack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable FACK (Forward ACK).",
			},
			"flavor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set TCP congestion control algorithm.",
			},
			"frto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable FRTO (Forward RTO-Recovery).",
			},
			"hystart": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable CUBIC Hystart",
			},
			"initialcwnd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.",
			},
			"ka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send periodic TCP keep-alive (KA) probes to check if peer is still up.",
			},
			"kaconnidletime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Duration, in seconds, for the connection to be idle, before sending a keep-alive (KA) probe.",
			},
			"kamaxprobes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of keep-alive (KA) probes to be sent when not acknowledged, before assuming the peer to be down.",
			},
			"kaprobeinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in seconds, before the next keep-alive (KA) probe, if the peer does not respond.",
			},
			"kaprobeupdatelastactivity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Update last activity for the connection after receiving keep-alive (KA) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP segments allowed in a burst.",
			},
			"maxcwnd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Maximum Congestion Window.",
			},
			"maxpktpermss": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP packets allowed per maximum segment size (MSS).",
			},
			"minrto": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by  10).",
			},
			"mpcapablecbit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set C bit in MP-CAPABLE Syn-Ack sent by Citrix ADC",
			},
			"mptcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Multipath TCP.",
			},
			"mptcpdropdataonpreestsf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow.",
			},
			"mptcpfastopen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake.",
			},
			"mptcpsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MPTCP session timeout in seconds. If this value is not set, idle MPTCP sessions are flushed after vserver's client idle timeout.",
			},
			"mss": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of octets to allow in a TCP data segment.",
			},
			"nagle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for a TCP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a TCP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my tcp profile\" or 'my tcp profile'\\).",
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
			"rateqmax": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum connection queue size in bytes, when BurstRateControl is used",
			},
			"rfc5961compliance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable RFC 5961 compliance to protect against tcp spoofing(RST/SYN/Data). When enabled, will be compliant with RFC 5961.",
			},
			"rstmaxack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode.",
			},
			"rstwindowattenuate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid.",
			},
			"sack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Selective ACKnowledgement (SACK).",
			},
			"sendbuffsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Send Buffer Size",
			},
			"sendclientportintcpoption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send Client Port number along with Client IP in TCP-Options. ClientIpTcpOption must be ENABLED",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"slowstartthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Slow Start Threhsold Value.",
			},
			"spoofsyndrop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received.",
			},
			"syncookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC.",
			},
			"taillossprobe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP tail loss probe optimizations",
			},
			"tcpfastopen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets.",
			},
			"tcpfastopencookiesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.",
			},
			"tcpmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Optimization modes TRANSPARENT / ENDPOINT.",
			},
			"tcprate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP connection payload send rate in Kb/s",
			},
			"tcpsegoffload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it.",
			},
			"timestamp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or Disable TCP Timestamp option (RFC 1323)",
			},
			"ws": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable window scaling.",
			},
			"wsval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Factor used to calculate the new window size.\nThis argument is needed only when window scaling is enabled.",
			},
		},
	}
}
