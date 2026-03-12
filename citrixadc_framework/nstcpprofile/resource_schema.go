package nstcpprofile

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

// NstcpprofileResourceModel describes the resource data model.
type NstcpprofileResourceModel struct {
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
}

func (r *NstcpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstcpprofile resource.",
			},
			"ackaggregation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable ACK Aggregation.",
			},
			"ackonpush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"applyadaptivetcp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Apply Adaptive TCP optimizations",
			},
			"buffersize": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_BUFFSIZE),
				Description: "TCP buffering size, in bytes.",
			},
			"burstratecontrol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set.",
			},
			"clientiptcpoption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Client IP in TCP options",
			},
			"clientiptcpoptionnumber": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(DISABLED),
				Description: "ClientIP TCP Option number",
			},
			"delayedack": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"dropestconnontimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Silently drop tcp established connections on idle timeout",
			},
			"drophalfclosedconnontimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Silently drop tcp half closed connections on idle timeout",
			},
			"dsack": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable DSACK.",
			},
			"dupackthresh": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "TCP dupack threshold.",
			},
			"dynamicreceivebuffering": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions.\nNote: The buffer size argument must be set for dynamic adjustments to take place.",
			},
			"ecn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable TCP Explicit Congestion Notification.",
			},
			"establishclientconn": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AUTOMATIC"),
				Description: "Establishing Client Client connection on First data/ Final-ACK / Automatic",
			},
			"fack": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable FACK (Forward ACK).",
			},
			"flavor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("BIC"),
				Description: "Set TCP congestion control algorithm.",
			},
			"frto": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable FRTO (Forward RTO-Recovery).",
			},
			"hystart": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable CUBIC Hystart",
			},
			"initialcwnd": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_INITIALCWND),
				Description: "Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.",
			},
			"ka": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Update last activity for the connection after receiving keep-alive (KA) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(6),
				Description: "Maximum number of TCP segments allowed in a burst.",
			},
			"maxcwnd": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(524288),
				Description: "TCP Maximum Congestion Window.",
			},
			"maxpktpermss": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of TCP packets allowed per maximum segment size (MSS).",
			},
			"minrto": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by  10).",
			},
			"mpcapablecbit": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Set C bit in MP-CAPABLE Syn-Ack sent by Citrix ADC",
			},
			"mptcp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable Multipath TCP.",
			},
			"mptcpdropdataonpreestsf": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow.",
			},
			"mptcpfastopen": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake.",
			},
			"mptcpsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MPTCP session timeout in seconds. If this value is not set, idle MPTCP sessions are flushed after vserver's client idle timeout.",
			},
			"mss": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_CLIENT_MSS),
				Description: "Maximum number of octets to allow in a TCP data segment.",
			},
			"nagle": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for a TCP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a TCP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my tcp profile\" or 'my tcp profile'\\).",
			},
			"oooqsize": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_MAX_OOO_PKTS),
				Description: "Maximum size of out-of-order packets queue. A value of 0 means no limit.",
			},
			"pktperretx": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.",
			},
			"rateqmax": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum connection queue size in bytes, when BurstRateControl is used",
			},
			"rfc5961compliance": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable RFC 5961 compliance to protect against tcp spoofing(RST/SYN/Data). When enabled, will be compliant with RFC 5961.",
			},
			"rstmaxack": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode.",
			},
			"rstwindowattenuate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid.",
			},
			"sack": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable Selective ACKnowledgement (SACK).",
			},
			"sendbuffsize": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_SENDBUFFSIZE),
				Description: "TCP Send Buffer Size",
			},
			"sendclientportintcpoption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send Client Port number along with Client IP in TCP-Options. ClientIpTcpOption must be ENABLED",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"slowstartthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(524288),
				Description: "TCP Slow Start Threhsold Value.",
			},
			"spoofsyndrop": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received.",
			},
			"syncookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC.",
			},
			"taillossprobe": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "TCP tail loss probe optimizations",
			},
			"tcpfastopen": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets.",
			},
			"tcpfastopencookiesize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8),
				Description: "TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.",
			},
			"tcpmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("TRANSPARENT"),
				Description: "TCP Optimization modes TRANSPARENT / ENDPOINT.",
			},
			"tcprate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP connection payload send rate in Kb/s",
			},
			"tcpsegoffload": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AUTOMATIC"),
				Description: "Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it.",
			},
			"timestamp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or Disable TCP Timestamp option (RFC 1323)",
			},
			"ws": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable window scaling.",
			},
			"wsval": schema.Int64Attribute{
				Optional: true,
				// Default:     int64default.StaticInt64(TCP_DEFAULT_WSVAL),
				Description: "Factor used to calculate the new window size.\nThis argument is needed only when window scaling is enabled.",
			},
		},
	}
}

func nstcpprofileGetThePayloadFromtheConfig(ctx context.Context, data *NstcpprofileResourceModel) ns.Nstcpprofile {
	tflog.Debug(ctx, "In nstcpprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstcpprofile := ns.Nstcpprofile{}
	if !data.Ackaggregation.IsNull() {
		nstcpprofile.Ackaggregation = data.Ackaggregation.ValueString()
	}
	if !data.Ackonpush.IsNull() {
		nstcpprofile.Ackonpush = data.Ackonpush.ValueString()
	}
	if !data.Applyadaptivetcp.IsNull() {
		nstcpprofile.Applyadaptivetcp = data.Applyadaptivetcp.ValueString()
	}
	if !data.Buffersize.IsNull() {
		nstcpprofile.Buffersize = utils.IntPtr(int(data.Buffersize.ValueInt64()))
	}
	if !data.Burstratecontrol.IsNull() {
		nstcpprofile.Burstratecontrol = data.Burstratecontrol.ValueString()
	}
	if !data.Clientiptcpoption.IsNull() {
		nstcpprofile.Clientiptcpoption = data.Clientiptcpoption.ValueString()
	}
	if !data.Clientiptcpoptionnumber.IsNull() {
		nstcpprofile.Clientiptcpoptionnumber = utils.IntPtr(int(data.Clientiptcpoptionnumber.ValueInt64()))
	}
	if !data.Delayedack.IsNull() {
		nstcpprofile.Delayedack = utils.IntPtr(int(data.Delayedack.ValueInt64()))
	}
	if !data.Dropestconnontimeout.IsNull() {
		nstcpprofile.Dropestconnontimeout = data.Dropestconnontimeout.ValueString()
	}
	if !data.Drophalfclosedconnontimeout.IsNull() {
		nstcpprofile.Drophalfclosedconnontimeout = data.Drophalfclosedconnontimeout.ValueString()
	}
	if !data.Dsack.IsNull() {
		nstcpprofile.Dsack = data.Dsack.ValueString()
	}
	if !data.Dupackthresh.IsNull() {
		nstcpprofile.Dupackthresh = utils.IntPtr(int(data.Dupackthresh.ValueInt64()))
	}
	if !data.Dynamicreceivebuffering.IsNull() {
		nstcpprofile.Dynamicreceivebuffering = data.Dynamicreceivebuffering.ValueString()
	}
	if !data.Ecn.IsNull() {
		nstcpprofile.Ecn = data.Ecn.ValueString()
	}
	if !data.Establishclientconn.IsNull() {
		nstcpprofile.Establishclientconn = data.Establishclientconn.ValueString()
	}
	if !data.Fack.IsNull() {
		nstcpprofile.Fack = data.Fack.ValueString()
	}
	if !data.Flavor.IsNull() {
		nstcpprofile.Flavor = data.Flavor.ValueString()
	}
	if !data.Frto.IsNull() {
		nstcpprofile.Frto = data.Frto.ValueString()
	}
	if !data.Hystart.IsNull() {
		nstcpprofile.Hystart = data.Hystart.ValueString()
	}
	if !data.Initialcwnd.IsNull() {
		nstcpprofile.Initialcwnd = utils.IntPtr(int(data.Initialcwnd.ValueInt64()))
	}
	if !data.Ka.IsNull() {
		nstcpprofile.Ka = data.Ka.ValueString()
	}
	if !data.Kaconnidletime.IsNull() {
		nstcpprofile.Kaconnidletime = utils.IntPtr(int(data.Kaconnidletime.ValueInt64()))
	}
	if !data.Kamaxprobes.IsNull() {
		nstcpprofile.Kamaxprobes = utils.IntPtr(int(data.Kamaxprobes.ValueInt64()))
	}
	if !data.Kaprobeinterval.IsNull() {
		nstcpprofile.Kaprobeinterval = utils.IntPtr(int(data.Kaprobeinterval.ValueInt64()))
	}
	if !data.Kaprobeupdatelastactivity.IsNull() {
		nstcpprofile.Kaprobeupdatelastactivity = data.Kaprobeupdatelastactivity.ValueString()
	}
	if !data.Maxburst.IsNull() {
		nstcpprofile.Maxburst = utils.IntPtr(int(data.Maxburst.ValueInt64()))
	}
	if !data.Maxcwnd.IsNull() {
		nstcpprofile.Maxcwnd = utils.IntPtr(int(data.Maxcwnd.ValueInt64()))
	}
	if !data.Maxpktpermss.IsNull() {
		nstcpprofile.Maxpktpermss = utils.IntPtr(int(data.Maxpktpermss.ValueInt64()))
	}
	if !data.Minrto.IsNull() {
		nstcpprofile.Minrto = utils.IntPtr(int(data.Minrto.ValueInt64()))
	}
	if !data.Mpcapablecbit.IsNull() {
		nstcpprofile.Mpcapablecbit = data.Mpcapablecbit.ValueString()
	}
	if !data.Mptcp.IsNull() {
		nstcpprofile.Mptcp = data.Mptcp.ValueString()
	}
	if !data.Mptcpdropdataonpreestsf.IsNull() {
		nstcpprofile.Mptcpdropdataonpreestsf = data.Mptcpdropdataonpreestsf.ValueString()
	}
	if !data.Mptcpfastopen.IsNull() {
		nstcpprofile.Mptcpfastopen = data.Mptcpfastopen.ValueString()
	}
	if !data.Mptcpsessiontimeout.IsNull() {
		nstcpprofile.Mptcpsessiontimeout = utils.IntPtr(int(data.Mptcpsessiontimeout.ValueInt64()))
	}
	if !data.Mss.IsNull() {
		nstcpprofile.Mss = utils.IntPtr(int(data.Mss.ValueInt64()))
	}
	if !data.Nagle.IsNull() {
		nstcpprofile.Nagle = data.Nagle.ValueString()
	}
	if !data.Name.IsNull() {
		nstcpprofile.Name = data.Name.ValueString()
	}
	if !data.Oooqsize.IsNull() {
		nstcpprofile.Oooqsize = utils.IntPtr(int(data.Oooqsize.ValueInt64()))
	}
	if !data.Pktperretx.IsNull() {
		nstcpprofile.Pktperretx = utils.IntPtr(int(data.Pktperretx.ValueInt64()))
	}
	if !data.Rateqmax.IsNull() {
		nstcpprofile.Rateqmax = utils.IntPtr(int(data.Rateqmax.ValueInt64()))
	}
	if !data.Rfc5961compliance.IsNull() {
		nstcpprofile.Rfc5961compliance = data.Rfc5961compliance.ValueString()
	}
	if !data.Rstmaxack.IsNull() {
		nstcpprofile.Rstmaxack = data.Rstmaxack.ValueString()
	}
	if !data.Rstwindowattenuate.IsNull() {
		nstcpprofile.Rstwindowattenuate = data.Rstwindowattenuate.ValueString()
	}
	if !data.Sack.IsNull() {
		nstcpprofile.Sack = data.Sack.ValueString()
	}
	if !data.Sendbuffsize.IsNull() {
		nstcpprofile.Sendbuffsize = utils.IntPtr(int(data.Sendbuffsize.ValueInt64()))
	}
	if !data.Sendclientportintcpoption.IsNull() {
		nstcpprofile.Sendclientportintcpoption = data.Sendclientportintcpoption.ValueString()
	}
	if !data.Slowstartincr.IsNull() {
		nstcpprofile.Slowstartincr = utils.IntPtr(int(data.Slowstartincr.ValueInt64()))
	}
	if !data.Slowstartthreshold.IsNull() {
		nstcpprofile.Slowstartthreshold = utils.IntPtr(int(data.Slowstartthreshold.ValueInt64()))
	}
	if !data.Spoofsyndrop.IsNull() {
		nstcpprofile.Spoofsyndrop = data.Spoofsyndrop.ValueString()
	}
	if !data.Syncookie.IsNull() {
		nstcpprofile.Syncookie = data.Syncookie.ValueString()
	}
	if !data.Taillossprobe.IsNull() {
		nstcpprofile.Taillossprobe = data.Taillossprobe.ValueString()
	}
	if !data.Tcpfastopen.IsNull() {
		nstcpprofile.Tcpfastopen = data.Tcpfastopen.ValueString()
	}
	if !data.Tcpfastopencookiesize.IsNull() {
		nstcpprofile.Tcpfastopencookiesize = utils.IntPtr(int(data.Tcpfastopencookiesize.ValueInt64()))
	}
	if !data.Tcpmode.IsNull() {
		nstcpprofile.Tcpmode = data.Tcpmode.ValueString()
	}
	if !data.Tcprate.IsNull() {
		nstcpprofile.Tcprate = utils.IntPtr(int(data.Tcprate.ValueInt64()))
	}
	if !data.Tcpsegoffload.IsNull() {
		nstcpprofile.Tcpsegoffload = data.Tcpsegoffload.ValueString()
	}
	if !data.Timestamp.IsNull() {
		nstcpprofile.Timestamp = data.Timestamp.ValueString()
	}
	if !data.Ws.IsNull() {
		nstcpprofile.Ws = data.Ws.ValueString()
	}
	if !data.Wsval.IsNull() {
		nstcpprofile.Wsval = utils.IntPtr(int(data.Wsval.ValueInt64()))
	}

	return nstcpprofile
}

func nstcpprofileSetAttrFromGet(ctx context.Context, data *NstcpprofileResourceModel, getResponseData map[string]interface{}) *NstcpprofileResourceModel {
	tflog.Debug(ctx, "In nstcpprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ackaggregation"]; ok && val != nil {
		data.Ackaggregation = types.StringValue(val.(string))
	} else {
		data.Ackaggregation = types.StringNull()
	}
	if val, ok := getResponseData["ackonpush"]; ok && val != nil {
		data.Ackonpush = types.StringValue(val.(string))
	} else {
		data.Ackonpush = types.StringNull()
	}
	if val, ok := getResponseData["applyadaptivetcp"]; ok && val != nil {
		data.Applyadaptivetcp = types.StringValue(val.(string))
	} else {
		data.Applyadaptivetcp = types.StringNull()
	}
	if val, ok := getResponseData["buffersize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Buffersize = types.Int64Value(intVal)
		}
	} else {
		data.Buffersize = types.Int64Null()
	}
	if val, ok := getResponseData["burstratecontrol"]; ok && val != nil {
		data.Burstratecontrol = types.StringValue(val.(string))
	} else {
		data.Burstratecontrol = types.StringNull()
	}
	if val, ok := getResponseData["clientiptcpoption"]; ok && val != nil {
		data.Clientiptcpoption = types.StringValue(val.(string))
	} else {
		data.Clientiptcpoption = types.StringNull()
	}
	if val, ok := getResponseData["clientiptcpoptionnumber"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clientiptcpoptionnumber = types.Int64Value(intVal)
		}
	} else {
		data.Clientiptcpoptionnumber = types.Int64Null()
	}
	if val, ok := getResponseData["delayedack"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delayedack = types.Int64Value(intVal)
		}
	} else {
		data.Delayedack = types.Int64Null()
	}
	if val, ok := getResponseData["dropestconnontimeout"]; ok && val != nil {
		data.Dropestconnontimeout = types.StringValue(val.(string))
	} else {
		data.Dropestconnontimeout = types.StringNull()
	}
	if val, ok := getResponseData["drophalfclosedconnontimeout"]; ok && val != nil {
		data.Drophalfclosedconnontimeout = types.StringValue(val.(string))
	} else {
		data.Drophalfclosedconnontimeout = types.StringNull()
	}
	if val, ok := getResponseData["dsack"]; ok && val != nil {
		data.Dsack = types.StringValue(val.(string))
	} else {
		data.Dsack = types.StringNull()
	}
	if val, ok := getResponseData["dupackthresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dupackthresh = types.Int64Value(intVal)
		}
	} else {
		data.Dupackthresh = types.Int64Null()
	}
	if val, ok := getResponseData["dynamicreceivebuffering"]; ok && val != nil {
		data.Dynamicreceivebuffering = types.StringValue(val.(string))
	} else {
		data.Dynamicreceivebuffering = types.StringNull()
	}
	if val, ok := getResponseData["ecn"]; ok && val != nil {
		data.Ecn = types.StringValue(val.(string))
	} else {
		data.Ecn = types.StringNull()
	}
	if val, ok := getResponseData["establishclientconn"]; ok && val != nil {
		data.Establishclientconn = types.StringValue(val.(string))
	} else {
		data.Establishclientconn = types.StringNull()
	}
	if val, ok := getResponseData["fack"]; ok && val != nil {
		data.Fack = types.StringValue(val.(string))
	} else {
		data.Fack = types.StringNull()
	}
	if val, ok := getResponseData["flavor"]; ok && val != nil {
		data.Flavor = types.StringValue(val.(string))
	} else {
		data.Flavor = types.StringNull()
	}
	if val, ok := getResponseData["frto"]; ok && val != nil {
		data.Frto = types.StringValue(val.(string))
	} else {
		data.Frto = types.StringNull()
	}
	if val, ok := getResponseData["hystart"]; ok && val != nil {
		data.Hystart = types.StringValue(val.(string))
	} else {
		data.Hystart = types.StringNull()
	}
	if val, ok := getResponseData["initialcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialcwnd = types.Int64Value(intVal)
		}
	} else {
		data.Initialcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["ka"]; ok && val != nil {
		data.Ka = types.StringValue(val.(string))
	} else {
		data.Ka = types.StringNull()
	}
	if val, ok := getResponseData["kaconnidletime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kaconnidletime = types.Int64Value(intVal)
		}
	} else {
		data.Kaconnidletime = types.Int64Null()
	}
	if val, ok := getResponseData["kamaxprobes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kamaxprobes = types.Int64Value(intVal)
		}
	} else {
		data.Kamaxprobes = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kaprobeinterval = types.Int64Value(intVal)
		}
	} else {
		data.Kaprobeinterval = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeupdatelastactivity"]; ok && val != nil {
		data.Kaprobeupdatelastactivity = types.StringValue(val.(string))
	} else {
		data.Kaprobeupdatelastactivity = types.StringNull()
	}
	if val, ok := getResponseData["maxburst"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxburst = types.Int64Value(intVal)
		}
	} else {
		data.Maxburst = types.Int64Null()
	}
	if val, ok := getResponseData["maxcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxcwnd = types.Int64Value(intVal)
		}
	} else {
		data.Maxcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["maxpktpermss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpktpermss = types.Int64Value(intVal)
		}
	} else {
		data.Maxpktpermss = types.Int64Null()
	}
	if val, ok := getResponseData["minrto"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrto = types.Int64Value(intVal)
		}
	} else {
		data.Minrto = types.Int64Null()
	}
	if val, ok := getResponseData["mpcapablecbit"]; ok && val != nil {
		data.Mpcapablecbit = types.StringValue(val.(string))
	} else {
		data.Mpcapablecbit = types.StringNull()
	}
	if val, ok := getResponseData["mptcp"]; ok && val != nil {
		data.Mptcp = types.StringValue(val.(string))
	} else {
		data.Mptcp = types.StringNull()
	}
	if val, ok := getResponseData["mptcpdropdataonpreestsf"]; ok && val != nil {
		data.Mptcpdropdataonpreestsf = types.StringValue(val.(string))
	} else {
		data.Mptcpdropdataonpreestsf = types.StringNull()
	}
	if val, ok := getResponseData["mptcpfastopen"]; ok && val != nil {
		data.Mptcpfastopen = types.StringValue(val.(string))
	} else {
		data.Mptcpfastopen = types.StringNull()
	}
	if val, ok := getResponseData["mptcpsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Mptcpsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mss = types.Int64Value(intVal)
		}
	} else {
		data.Mss = types.Int64Null()
	}
	if val, ok := getResponseData["nagle"]; ok && val != nil {
		data.Nagle = types.StringValue(val.(string))
	} else {
		data.Nagle = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	if val, ok := getResponseData["rateqmax"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rateqmax = types.Int64Value(intVal)
		}
	} else {
		data.Rateqmax = types.Int64Null()
	}
	if val, ok := getResponseData["rfc5961compliance"]; ok && val != nil {
		data.Rfc5961compliance = types.StringValue(val.(string))
	} else {
		data.Rfc5961compliance = types.StringNull()
	}
	if val, ok := getResponseData["rstmaxack"]; ok && val != nil {
		data.Rstmaxack = types.StringValue(val.(string))
	} else {
		data.Rstmaxack = types.StringNull()
	}
	if val, ok := getResponseData["rstwindowattenuate"]; ok && val != nil {
		data.Rstwindowattenuate = types.StringValue(val.(string))
	} else {
		data.Rstwindowattenuate = types.StringNull()
	}
	if val, ok := getResponseData["sack"]; ok && val != nil {
		data.Sack = types.StringValue(val.(string))
	} else {
		data.Sack = types.StringNull()
	}
	if val, ok := getResponseData["sendbuffsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sendbuffsize = types.Int64Value(intVal)
		}
	} else {
		data.Sendbuffsize = types.Int64Null()
	}
	if val, ok := getResponseData["sendclientportintcpoption"]; ok && val != nil {
		data.Sendclientportintcpoption = types.StringValue(val.(string))
	} else {
		data.Sendclientportintcpoption = types.StringNull()
	}
	if val, ok := getResponseData["slowstartincr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartincr = types.Int64Value(intVal)
		}
	} else {
		data.Slowstartincr = types.Int64Null()
	}
	if val, ok := getResponseData["slowstartthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Slowstartthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["spoofsyndrop"]; ok && val != nil {
		data.Spoofsyndrop = types.StringValue(val.(string))
	} else {
		data.Spoofsyndrop = types.StringNull()
	}
	if val, ok := getResponseData["syncookie"]; ok && val != nil {
		data.Syncookie = types.StringValue(val.(string))
	} else {
		data.Syncookie = types.StringNull()
	}
	if val, ok := getResponseData["taillossprobe"]; ok && val != nil {
		data.Taillossprobe = types.StringValue(val.(string))
	} else {
		data.Taillossprobe = types.StringNull()
	}
	if val, ok := getResponseData["tcpfastopen"]; ok && val != nil {
		data.Tcpfastopen = types.StringValue(val.(string))
	} else {
		data.Tcpfastopen = types.StringNull()
	}
	if val, ok := getResponseData["tcpfastopencookiesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfastopencookiesize = types.Int64Value(intVal)
		}
	} else {
		data.Tcpfastopencookiesize = types.Int64Null()
	}
	if val, ok := getResponseData["tcpmode"]; ok && val != nil {
		data.Tcpmode = types.StringValue(val.(string))
	} else {
		data.Tcpmode = types.StringNull()
	}
	if val, ok := getResponseData["tcprate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcprate = types.Int64Value(intVal)
		}
	} else {
		data.Tcprate = types.Int64Null()
	}
	if val, ok := getResponseData["tcpsegoffload"]; ok && val != nil {
		data.Tcpsegoffload = types.StringValue(val.(string))
	} else {
		data.Tcpsegoffload = types.StringNull()
	}
	if val, ok := getResponseData["timestamp"]; ok && val != nil {
		data.Timestamp = types.StringValue(val.(string))
	} else {
		data.Timestamp = types.StringNull()
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
