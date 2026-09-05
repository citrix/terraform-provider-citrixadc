package nstcpprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NstcpprofileDataSourceModel is the data-source-specific model, decoupled from
// NstcpprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type NstcpprofileDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Ackaggregation              types.String `tfsdk:"ackaggregation"`
	Ackonpush                   types.String `tfsdk:"ackonpush"`
	Applyadaptivetcp            types.String `tfsdk:"applyadaptivetcp"`
	Buffersize                  types.Int64  `tfsdk:"buffersize"`
	Burstratecontrol            types.String `tfsdk:"burstratecontrol"`
	Clientiptcpoption           types.String `tfsdk:"clientiptcpoption"`
	Clientiptcpoptionnumber     types.Int64  `tfsdk:"clientiptcpoptionnumber"`
	Delayedack                  types.Int64  `tfsdk:"delayedack"`
	Dropestconnontimeout        types.String `tfsdk:"dropestconnontimeout"`
	Drophalfclosedconnontimeout types.String `tfsdk:"drophalfclosedconnontimeout"`
	Dsack                       types.String `tfsdk:"dsack"`
	Dupackthresh                types.Int64  `tfsdk:"dupackthresh"`
	Dynamicreceivebuffering     types.String `tfsdk:"dynamicreceivebuffering"`
	Ecn                         types.String `tfsdk:"ecn"`
	Establishclientconn         types.String `tfsdk:"establishclientconn"`
	Fack                        types.String `tfsdk:"fack"`
	Flavor                      types.String `tfsdk:"flavor"`
	Frto                        types.String `tfsdk:"frto"`
	Hystart                     types.String `tfsdk:"hystart"`
	Initialcwnd                 types.Int64  `tfsdk:"initialcwnd"`
	Ka                          types.String `tfsdk:"ka"`
	Kaconnidletime              types.Int64  `tfsdk:"kaconnidletime"`
	Kamaxprobes                 types.Int64  `tfsdk:"kamaxprobes"`
	Kaprobeinterval             types.Int64  `tfsdk:"kaprobeinterval"`
	Kaprobeupdatelastactivity   types.String `tfsdk:"kaprobeupdatelastactivity"`
	Maxburst                    types.Int64  `tfsdk:"maxburst"`
	Maxcwnd                     types.Int64  `tfsdk:"maxcwnd"`
	Maxpktpermss                types.Int64  `tfsdk:"maxpktpermss"`
	Minrto                      types.Int64  `tfsdk:"minrto"`
	Mpcapablecbit               types.String `tfsdk:"mpcapablecbit"`
	Mptcp                       types.String `tfsdk:"mptcp"`
	Mptcpdropdataonpreestsf     types.String `tfsdk:"mptcpdropdataonpreestsf"`
	Mptcpfastopen               types.String `tfsdk:"mptcpfastopen"`
	Mptcpsessiontimeout         types.Int64  `tfsdk:"mptcpsessiontimeout"`
	Mss                         types.Int64  `tfsdk:"mss"`
	Nagle                       types.String `tfsdk:"nagle"`
	Name                        types.String `tfsdk:"name"`
	Oooqsize                    types.Int64  `tfsdk:"oooqsize"`
	Pktperretx                  types.Int64  `tfsdk:"pktperretx"`
	Rateqmax                    types.Int64  `tfsdk:"rateqmax"`
	Rfc5961compliance           types.String `tfsdk:"rfc5961compliance"`
	Rstmaxack                   types.String `tfsdk:"rstmaxack"`
	Rstwindowattenuate          types.String `tfsdk:"rstwindowattenuate"`
	Sack                        types.String `tfsdk:"sack"`
	Sendbuffsize                types.Int64  `tfsdk:"sendbuffsize"`
	Sendclientportintcpoption   types.String `tfsdk:"sendclientportintcpoption"`
	Slowstartincr               types.Int64  `tfsdk:"slowstartincr"`
	Slowstartthreshold          types.Int64  `tfsdk:"slowstartthreshold"`
	Spoofsyndrop                types.String `tfsdk:"spoofsyndrop"`
	Syncookie                   types.String `tfsdk:"syncookie"`
	Taillossprobe               types.String `tfsdk:"taillossprobe"`
	Tcpfastopen                 types.String `tfsdk:"tcpfastopen"`
	Tcpfastopencookiesize       types.Int64  `tfsdk:"tcpfastopencookiesize"`
	Tcpmode                     types.String `tfsdk:"tcpmode"`
	Tcprate                     types.Int64  `tfsdk:"tcprate"`
	Tcpsegoffload               types.String `tfsdk:"tcpsegoffload"`
	Timestamp                   types.String `tfsdk:"timestamp"`
	Ws                          types.String `tfsdk:"ws"`
	Wsval                       types.Int64  `tfsdk:"wsval"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nstcpprofile.json). Never settable; populated from GET.
	Refcnt  types.Int64  `tfsdk:"refcnt"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this profile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if tcp profile is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nstcpprofileDataSourceSetAttrFromGet projects a NITRO nstcpprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nstcpprofileDataSourceSetAttrFromGet(ctx context.Context, data *NstcpprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nstcpprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ackaggregation = utils.MapGetString(g, "ackaggregation")
	data.Ackonpush = utils.MapGetString(g, "ackonpush")
	data.Applyadaptivetcp = utils.MapGetString(g, "applyadaptivetcp")
	data.Buffersize = utils.MapGetInt64(g, "buffersize")
	data.Burstratecontrol = utils.MapGetString(g, "burstratecontrol")
	data.Clientiptcpoption = utils.MapGetString(g, "clientiptcpoption")
	data.Clientiptcpoptionnumber = utils.MapGetInt64(g, "clientiptcpoptionnumber")
	data.Delayedack = utils.MapGetInt64(g, "delayedack")
	data.Dropestconnontimeout = utils.MapGetString(g, "dropestconnontimeout")
	data.Drophalfclosedconnontimeout = utils.MapGetString(g, "drophalfclosedconnontimeout")
	data.Dsack = utils.MapGetString(g, "dsack")
	data.Dupackthresh = utils.MapGetInt64(g, "dupackthresh")
	data.Dynamicreceivebuffering = utils.MapGetString(g, "dynamicreceivebuffering")
	data.Ecn = utils.MapGetString(g, "ecn")
	data.Establishclientconn = utils.MapGetString(g, "establishclientconn")
	data.Fack = utils.MapGetString(g, "fack")
	data.Flavor = utils.MapGetString(g, "flavor")
	data.Frto = utils.MapGetString(g, "frto")
	data.Hystart = utils.MapGetString(g, "hystart")
	data.Initialcwnd = utils.MapGetInt64(g, "initialcwnd")
	data.Ka = utils.MapGetString(g, "ka")
	data.Kaconnidletime = utils.MapGetInt64(g, "kaconnidletime")
	data.Kamaxprobes = utils.MapGetInt64(g, "kamaxprobes")
	data.Kaprobeinterval = utils.MapGetInt64(g, "kaprobeinterval")
	data.Kaprobeupdatelastactivity = utils.MapGetString(g, "kaprobeupdatelastactivity")
	data.Maxburst = utils.MapGetInt64(g, "maxburst")
	data.Maxcwnd = utils.MapGetInt64(g, "maxcwnd")
	data.Maxpktpermss = utils.MapGetInt64(g, "maxpktpermss")
	data.Minrto = utils.MapGetInt64(g, "minrto")
	data.Mpcapablecbit = utils.MapGetString(g, "mpcapablecbit")
	data.Mptcp = utils.MapGetString(g, "mptcp")
	data.Mptcpdropdataonpreestsf = utils.MapGetString(g, "mptcpdropdataonpreestsf")
	data.Mptcpfastopen = utils.MapGetString(g, "mptcpfastopen")
	data.Mptcpsessiontimeout = utils.MapGetInt64(g, "mptcpsessiontimeout")
	data.Mss = utils.MapGetInt64(g, "mss")
	data.Nagle = utils.MapGetString(g, "nagle")
	data.Oooqsize = utils.MapGetInt64(g, "oooqsize")
	data.Pktperretx = utils.MapGetInt64(g, "pktperretx")
	data.Rateqmax = utils.MapGetInt64(g, "rateqmax")
	data.Rfc5961compliance = utils.MapGetString(g, "rfc5961compliance")
	data.Rstmaxack = utils.MapGetString(g, "rstmaxack")
	data.Rstwindowattenuate = utils.MapGetString(g, "rstwindowattenuate")
	data.Sack = utils.MapGetString(g, "sack")
	data.Sendbuffsize = utils.MapGetInt64(g, "sendbuffsize")
	data.Sendclientportintcpoption = utils.MapGetString(g, "sendclientportintcpoption")
	data.Slowstartincr = utils.MapGetInt64(g, "slowstartincr")
	data.Slowstartthreshold = utils.MapGetInt64(g, "slowstartthreshold")
	data.Spoofsyndrop = utils.MapGetString(g, "spoofsyndrop")
	data.Syncookie = utils.MapGetString(g, "syncookie")
	data.Taillossprobe = utils.MapGetString(g, "taillossprobe")
	data.Tcpfastopen = utils.MapGetString(g, "tcpfastopen")
	data.Tcpfastopencookiesize = utils.MapGetInt64(g, "tcpfastopencookiesize")
	data.Tcpmode = utils.MapGetString(g, "tcpmode")
	data.Tcprate = utils.MapGetInt64(g, "tcprate")
	data.Tcpsegoffload = utils.MapGetString(g, "tcpsegoffload")
	data.Timestamp = utils.MapGetString(g, "timestamp")
	data.Ws = utils.MapGetString(g, "ws")
	data.Wsval = utils.MapGetInt64(g, "wsval")

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
