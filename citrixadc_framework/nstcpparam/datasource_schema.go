package nstcpparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NstcpparamDataSourceModel is the data-source-specific model, decoupled from
// NstcpparamResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type NstcpparamDataSourceModel struct {
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
	Sendresetreasoncode                 types.String `tfsdk:"sendresetreasoncode"`
	Slowstartincr                       types.Int64  `tfsdk:"slowstartincr"`
	Synattackdetection                  types.String `tfsdk:"synattackdetection"`
	Synholdfastgiveup                   types.Int64  `tfsdk:"synholdfastgiveup"`
	Tcpfastopencookietimeout            types.Int64  `tfsdk:"tcpfastopencookietimeout"`
	Tcpfintimeout                       types.Int64  `tfsdk:"tcpfintimeout"`
	Tcpmaxretries                       types.Int64  `tfsdk:"tcpmaxretries"`
	Ws                                  types.String `tfsdk:"ws"`
	Wsval                               types.Int64  `tfsdk:"wsval"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nstcpparam.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

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
			"sendresetreasoncode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, NetScaler includes a debug code indicating the reason for the reset in the TCP Window header field of outgoing TCP RST segments.",
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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the tcp param is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nstcpparamDataSourceSetAttrFromGet projects a NITRO nstcpparam GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nstcpparamDataSourceSetAttrFromGet(ctx context.Context, data *NstcpparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nstcpparamDataSourceSetAttrFromGet Function")

	// nstcpparam is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("nstcpparam-config")

	// Read/write attributes as read-back outputs.
	data.Ackonpush = utils.MapGetString(g, "ackonpush")
	data.Autosyncookietimeout = utils.MapGetInt64(g, "autosyncookietimeout")
	data.Compacttcpoptionnoop = utils.MapGetString(g, "compacttcpoptionnoop")
	data.Connflushifnomem = utils.MapGetString(g, "connflushifnomem")
	data.Connflushthres = utils.MapGetInt64(g, "connflushthres")
	data.Delayedack = utils.MapGetInt64(g, "delayedack")
	data.Delinkclientserveronrst = utils.MapGetString(g, "delinkclientserveronrst")
	data.Downstaterst = utils.MapGetString(g, "downstaterst")
	data.Enhancedisngeneration = utils.MapGetString(g, "enhancedisngeneration")
	data.Initialcwnd = utils.MapGetInt64(g, "initialcwnd")
	data.Kaprobeupdatelastactivity = utils.MapGetString(g, "kaprobeupdatelastactivity")
	data.Learnvsvrmss = utils.MapGetString(g, "learnvsvrmss")
	data.Limitedpersist = utils.MapGetString(g, "limitedpersist")
	data.Maxburst = utils.MapGetInt64(g, "maxburst")
	data.Maxdynserverprobes = utils.MapGetInt64(g, "maxdynserverprobes")
	data.Maxpktpermss = utils.MapGetInt64(g, "maxpktpermss")
	data.Maxsynackretx = utils.MapGetInt64(g, "maxsynackretx")
	data.Maxsynhold = utils.MapGetInt64(g, "maxsynhold")
	data.Maxsynholdperprobe = utils.MapGetInt64(g, "maxsynholdperprobe")
	data.Maxtimewaitconn = utils.MapGetInt64(g, "maxtimewaitconn")
	data.Minrto = utils.MapGetInt64(g, "minrto")
	data.Mptcpchecksum = utils.MapGetString(g, "mptcpchecksum")
	data.Mptcpclosemptcpsessiononlastsfclose = utils.MapGetString(g, "mptcpclosemptcpsessiononlastsfclose")
	data.Mptcpconcloseonpassivesf = utils.MapGetString(g, "mptcpconcloseonpassivesf")
	data.Mptcpfastcloseoption = utils.MapGetString(g, "mptcpfastcloseoption")
	data.Mptcpimmediatesfcloseonfin = utils.MapGetString(g, "mptcpimmediatesfcloseonfin")
	data.Mptcpmaxpendingsf = utils.MapGetInt64(g, "mptcpmaxpendingsf")
	data.Mptcpmaxsf = utils.MapGetInt64(g, "mptcpmaxsf")
	data.Mptcppendingjointhreshold = utils.MapGetInt64(g, "mptcppendingjointhreshold")
	data.Mptcpreliableaddaddr = utils.MapGetString(g, "mptcpreliableaddaddr")
	data.Mptcprtostoswitchsf = utils.MapGetInt64(g, "mptcprtostoswitchsf")
	data.Mptcpsendsfresetoption = utils.MapGetString(g, "mptcpsendsfresetoption")
	data.Mptcpsfreplacetimeout = utils.MapGetInt64(g, "mptcpsfreplacetimeout")
	data.Mptcpsftimeout = utils.MapGetInt64(g, "mptcpsftimeout")
	data.Mptcpusebackupondss = utils.MapGetString(g, "mptcpusebackupondss")
	data.Msslearndelay = utils.MapGetInt64(g, "msslearndelay")
	data.Msslearninterval = utils.MapGetInt64(g, "msslearninterval")
	data.Nagle = utils.MapGetString(g, "nagle")
	data.Oooqsize = utils.MapGetInt64(g, "oooqsize")
	data.Pktperretx = utils.MapGetInt64(g, "pktperretx")
	data.Recvbuffsize = utils.MapGetInt64(g, "recvbuffsize")
	data.Rfc5961chlgacklimit = utils.MapGetInt64(g, "rfc5961chlgacklimit")
	data.Sack = utils.MapGetString(g, "sack")
	data.Sendresetreasoncode = utils.MapGetString(g, "sendresetreasoncode")
	data.Slowstartincr = utils.MapGetInt64(g, "slowstartincr")
	data.Synattackdetection = utils.MapGetString(g, "synattackdetection")
	data.Synholdfastgiveup = utils.MapGetInt64(g, "synholdfastgiveup")
	data.Tcpfastopencookietimeout = utils.MapGetInt64(g, "tcpfastopencookietimeout")
	data.Tcpfintimeout = utils.MapGetInt64(g, "tcpfintimeout")
	data.Tcpmaxretries = utils.MapGetInt64(g, "tcpmaxretries")
	data.Ws = utils.MapGetString(g, "ws")
	data.Wsval = utils.MapGetInt64(g, "wsval")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
