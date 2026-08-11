package nstcpprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable ACK Aggregation.",
			},
			"ackonpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.",
			},
			"applyadaptivetcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set.",
			},
			"clientiptcpoption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     int64default.StaticInt64(100),
				Description: "Timeout for TCP delayed ACK, in milliseconds.",
			},
			"dropestconnontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Silently drop tcp established connections on idle timeout",
			},
			"drophalfclosedconnontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Silently drop tcp half closed connections on idle timeout",
			},
			"dsack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable DSACK.",
			},
			"dupackthresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3),
				Description: "TCP dupack threshold.",
			},
			"dynamicreceivebuffering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions.\nNote: The buffer size argument must be set for dynamic adjustments to take place.",
			},
			"ecn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable TCP Explicit Congestion Notification.",
			},
			"establishclientconn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("AUTOMATIC"),
				Description: "Establishing Client Client connection on First data/ Final-ACK / Automatic",
			},
			"fack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable FACK (Forward ACK).",
			},
			"flavor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("BIC"),
				Description: "Set TCP congestion control algorithm.",
			},
			"frto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable FRTO (Forward RTO-Recovery).",
			},
			"hystart": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Update last activity for the connection after receiving keep-alive (KA) probes.",
			},
			"maxburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(6),
				Description: "Maximum number of TCP segments allowed in a burst.",
			},
			"maxcwnd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by  10).",
			},
			"mpcapablecbit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Set C bit in MP-CAPABLE Syn-Ack sent by Citrix ADC",
			},
			"mptcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable Multipath TCP.",
			},
			"mptcpdropdataonpreestsf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow.",
			},
			"mptcpfastopen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake.",
			},
			"mptcpsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable the Nagle algorithm on TCP connections.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Default:     int64default.StaticInt64(1),
				Description: "Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.",
			},
			"rateqmax": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Maximum connection queue size in bytes, when BurstRateControl is used",
			},
			"rfc5961compliance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable RFC 5961 compliance to protect against tcp spoofing(RST/SYN/Data). When enabled, will be compliant with RFC 5961.",
			},
			"rstmaxack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode.",
			},
			"rstwindowattenuate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid.",
			},
			"sack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send Client Port number along with Client IP in TCP-Options. ClientIpTcpOption must be ENABLED",
			},
			"slowstartincr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.",
			},
			"slowstartthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(524288),
				Description: "TCP Slow Start Threhsold Value.",
			},
			"spoofsyndrop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received.",
			},
			"syncookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC.",
			},
			"taillossprobe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "TCP tail loss probe optimizations",
			},
			"tcpfastopen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets.",
			},
			"tcpfastopencookiesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(8),
				Description: "TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.",
			},
			"tcpmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("TRANSPARENT"),
				Description: "TCP Optimization modes TRANSPARENT / ENDPOINT.",
			},
			"tcprate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "TCP connection payload send rate in Kb/s",
			},
			"tcpsegoffload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("AUTOMATIC"),
				Description: "Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it.",
			},
			"timestamp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or Disable TCP Timestamp option (RFC 1323)",
			},
			"ws": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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

func nstcpprofileGetThePayloadFromthePlan(ctx context.Context, data *NstcpprofileResourceModel) ns.Nstcpprofile {
	tflog.Debug(ctx, "In nstcpprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nstcpprofile := ns.Nstcpprofile{}
	if !data.Ackaggregation.IsNull() && !data.Ackaggregation.IsUnknown() {
		nstcpprofile.Ackaggregation = data.Ackaggregation.ValueString()
	}
	if !data.Ackonpush.IsNull() && !data.Ackonpush.IsUnknown() {
		nstcpprofile.Ackonpush = data.Ackonpush.ValueString()
	}
	if !data.Applyadaptivetcp.IsNull() && !data.Applyadaptivetcp.IsUnknown() {
		nstcpprofile.Applyadaptivetcp = data.Applyadaptivetcp.ValueString()
	}
	if !data.Buffersize.IsNull() && !data.Buffersize.IsUnknown() {
		nstcpprofile.Buffersize = utils.IntPtr(int(data.Buffersize.ValueInt64()))
	}
	if !data.Burstratecontrol.IsNull() && !data.Burstratecontrol.IsUnknown() {
		nstcpprofile.Burstratecontrol = data.Burstratecontrol.ValueString()
	}
	if !data.Clientiptcpoption.IsNull() && !data.Clientiptcpoption.IsUnknown() {
		nstcpprofile.Clientiptcpoption = data.Clientiptcpoption.ValueString()
	}
	if !data.Clientiptcpoptionnumber.IsNull() && !data.Clientiptcpoptionnumber.IsUnknown() {
		nstcpprofile.Clientiptcpoptionnumber = utils.IntPtr(int(data.Clientiptcpoptionnumber.ValueInt64()))
	}
	if !data.Delayedack.IsNull() && !data.Delayedack.IsUnknown() {
		nstcpprofile.Delayedack = utils.IntPtr(int(data.Delayedack.ValueInt64()))
	}
	if !data.Dropestconnontimeout.IsNull() && !data.Dropestconnontimeout.IsUnknown() {
		nstcpprofile.Dropestconnontimeout = data.Dropestconnontimeout.ValueString()
	}
	if !data.Drophalfclosedconnontimeout.IsNull() && !data.Drophalfclosedconnontimeout.IsUnknown() {
		nstcpprofile.Drophalfclosedconnontimeout = data.Drophalfclosedconnontimeout.ValueString()
	}
	if !data.Dsack.IsNull() && !data.Dsack.IsUnknown() {
		nstcpprofile.Dsack = data.Dsack.ValueString()
	}
	if !data.Dupackthresh.IsNull() && !data.Dupackthresh.IsUnknown() {
		nstcpprofile.Dupackthresh = utils.IntPtr(int(data.Dupackthresh.ValueInt64()))
	}
	if !data.Dynamicreceivebuffering.IsNull() && !data.Dynamicreceivebuffering.IsUnknown() {
		nstcpprofile.Dynamicreceivebuffering = data.Dynamicreceivebuffering.ValueString()
	}
	if !data.Ecn.IsNull() && !data.Ecn.IsUnknown() {
		nstcpprofile.Ecn = data.Ecn.ValueString()
	}
	if !data.Establishclientconn.IsNull() && !data.Establishclientconn.IsUnknown() {
		nstcpprofile.Establishclientconn = data.Establishclientconn.ValueString()
	}
	if !data.Fack.IsNull() && !data.Fack.IsUnknown() {
		nstcpprofile.Fack = data.Fack.ValueString()
	}
	if !data.Flavor.IsNull() && !data.Flavor.IsUnknown() {
		nstcpprofile.Flavor = data.Flavor.ValueString()
	}
	if !data.Frto.IsNull() && !data.Frto.IsUnknown() {
		nstcpprofile.Frto = data.Frto.ValueString()
	}
	if !data.Hystart.IsNull() && !data.Hystart.IsUnknown() {
		nstcpprofile.Hystart = data.Hystart.ValueString()
	}
	if !data.Initialcwnd.IsNull() && !data.Initialcwnd.IsUnknown() {
		nstcpprofile.Initialcwnd = utils.IntPtr(int(data.Initialcwnd.ValueInt64()))
	}
	if !data.Ka.IsNull() && !data.Ka.IsUnknown() {
		nstcpprofile.Ka = data.Ka.ValueString()
	}
	if !data.Kaconnidletime.IsNull() && !data.Kaconnidletime.IsUnknown() {
		nstcpprofile.Kaconnidletime = utils.IntPtr(int(data.Kaconnidletime.ValueInt64()))
	}
	if !data.Kamaxprobes.IsNull() && !data.Kamaxprobes.IsUnknown() {
		nstcpprofile.Kamaxprobes = utils.IntPtr(int(data.Kamaxprobes.ValueInt64()))
	}
	if !data.Kaprobeinterval.IsNull() && !data.Kaprobeinterval.IsUnknown() {
		nstcpprofile.Kaprobeinterval = utils.IntPtr(int(data.Kaprobeinterval.ValueInt64()))
	}
	if !data.Kaprobeupdatelastactivity.IsNull() && !data.Kaprobeupdatelastactivity.IsUnknown() {
		nstcpprofile.Kaprobeupdatelastactivity = data.Kaprobeupdatelastactivity.ValueString()
	}
	if !data.Maxburst.IsNull() && !data.Maxburst.IsUnknown() {
		nstcpprofile.Maxburst = utils.IntPtr(int(data.Maxburst.ValueInt64()))
	}
	if !data.Maxcwnd.IsNull() && !data.Maxcwnd.IsUnknown() {
		nstcpprofile.Maxcwnd = utils.IntPtr(int(data.Maxcwnd.ValueInt64()))
	}
	if !data.Maxpktpermss.IsNull() && !data.Maxpktpermss.IsUnknown() {
		nstcpprofile.Maxpktpermss = utils.IntPtr(int(data.Maxpktpermss.ValueInt64()))
	}
	if !data.Minrto.IsNull() && !data.Minrto.IsUnknown() {
		nstcpprofile.Minrto = utils.IntPtr(int(data.Minrto.ValueInt64()))
	}
	if !data.Mpcapablecbit.IsNull() && !data.Mpcapablecbit.IsUnknown() {
		nstcpprofile.Mpcapablecbit = data.Mpcapablecbit.ValueString()
	}
	if !data.Mptcp.IsNull() && !data.Mptcp.IsUnknown() {
		nstcpprofile.Mptcp = data.Mptcp.ValueString()
	}
	if !data.Mptcpdropdataonpreestsf.IsNull() && !data.Mptcpdropdataonpreestsf.IsUnknown() {
		nstcpprofile.Mptcpdropdataonpreestsf = data.Mptcpdropdataonpreestsf.ValueString()
	}
	if !data.Mptcpfastopen.IsNull() && !data.Mptcpfastopen.IsUnknown() {
		nstcpprofile.Mptcpfastopen = data.Mptcpfastopen.ValueString()
	}
	if !data.Mptcpsessiontimeout.IsNull() && !data.Mptcpsessiontimeout.IsUnknown() {
		nstcpprofile.Mptcpsessiontimeout = utils.IntPtr(int(data.Mptcpsessiontimeout.ValueInt64()))
	}
	if !data.Mss.IsNull() && !data.Mss.IsUnknown() {
		nstcpprofile.Mss = utils.IntPtr(int(data.Mss.ValueInt64()))
	}
	if !data.Nagle.IsNull() && !data.Nagle.IsUnknown() {
		nstcpprofile.Nagle = data.Nagle.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nstcpprofile.Name = data.Name.ValueString()
	}
	if !data.Oooqsize.IsNull() && !data.Oooqsize.IsUnknown() {
		nstcpprofile.Oooqsize = utils.IntPtr(int(data.Oooqsize.ValueInt64()))
	}
	if !data.Pktperretx.IsNull() && !data.Pktperretx.IsUnknown() {
		nstcpprofile.Pktperretx = utils.IntPtr(int(data.Pktperretx.ValueInt64()))
	}
	if !data.Rateqmax.IsNull() && !data.Rateqmax.IsUnknown() {
		nstcpprofile.Rateqmax = utils.IntPtr(int(data.Rateqmax.ValueInt64()))
	}
	if !data.Rfc5961compliance.IsNull() && !data.Rfc5961compliance.IsUnknown() {
		nstcpprofile.Rfc5961compliance = data.Rfc5961compliance.ValueString()
	}
	if !data.Rstmaxack.IsNull() && !data.Rstmaxack.IsUnknown() {
		nstcpprofile.Rstmaxack = data.Rstmaxack.ValueString()
	}
	if !data.Rstwindowattenuate.IsNull() && !data.Rstwindowattenuate.IsUnknown() {
		nstcpprofile.Rstwindowattenuate = data.Rstwindowattenuate.ValueString()
	}
	if !data.Sack.IsNull() && !data.Sack.IsUnknown() {
		nstcpprofile.Sack = data.Sack.ValueString()
	}
	if !data.Sendbuffsize.IsNull() && !data.Sendbuffsize.IsUnknown() {
		nstcpprofile.Sendbuffsize = utils.IntPtr(int(data.Sendbuffsize.ValueInt64()))
	}
	if !data.Sendclientportintcpoption.IsNull() && !data.Sendclientportintcpoption.IsUnknown() {
		nstcpprofile.Sendclientportintcpoption = data.Sendclientportintcpoption.ValueString()
	}
	if !data.Slowstartincr.IsNull() && !data.Slowstartincr.IsUnknown() {
		nstcpprofile.Slowstartincr = utils.IntPtr(int(data.Slowstartincr.ValueInt64()))
	}
	if !data.Slowstartthreshold.IsNull() && !data.Slowstartthreshold.IsUnknown() {
		nstcpprofile.Slowstartthreshold = utils.IntPtr(int(data.Slowstartthreshold.ValueInt64()))
	}
	if !data.Spoofsyndrop.IsNull() && !data.Spoofsyndrop.IsUnknown() {
		nstcpprofile.Spoofsyndrop = data.Spoofsyndrop.ValueString()
	}
	if !data.Syncookie.IsNull() && !data.Syncookie.IsUnknown() {
		nstcpprofile.Syncookie = data.Syncookie.ValueString()
	}
	if !data.Taillossprobe.IsNull() && !data.Taillossprobe.IsUnknown() {
		nstcpprofile.Taillossprobe = data.Taillossprobe.ValueString()
	}
	if !data.Tcpfastopen.IsNull() && !data.Tcpfastopen.IsUnknown() {
		nstcpprofile.Tcpfastopen = data.Tcpfastopen.ValueString()
	}
	if !data.Tcpfastopencookiesize.IsNull() && !data.Tcpfastopencookiesize.IsUnknown() {
		nstcpprofile.Tcpfastopencookiesize = utils.IntPtr(int(data.Tcpfastopencookiesize.ValueInt64()))
	}
	if !data.Tcpmode.IsNull() && !data.Tcpmode.IsUnknown() {
		nstcpprofile.Tcpmode = data.Tcpmode.ValueString()
	}
	if !data.Tcprate.IsNull() && !data.Tcprate.IsUnknown() {
		nstcpprofile.Tcprate = utils.IntPtr(int(data.Tcprate.ValueInt64()))
	}
	if !data.Tcpsegoffload.IsNull() && !data.Tcpsegoffload.IsUnknown() {
		nstcpprofile.Tcpsegoffload = data.Tcpsegoffload.ValueString()
	}
	if !data.Timestamp.IsNull() && !data.Timestamp.IsUnknown() {
		nstcpprofile.Timestamp = data.Timestamp.ValueString()
	}
	if !data.Ws.IsNull() && !data.Ws.IsUnknown() {
		nstcpprofile.Ws = data.Ws.ValueString()
	}
	if !data.Wsval.IsNull() && !data.Wsval.IsUnknown() {
		nstcpprofile.Wsval = utils.IntPtr(int(data.Wsval.ValueInt64()))
	}

	return nstcpprofile
}

func nstcpprofileSetAttrFromGet(ctx context.Context, data *NstcpprofileResourceModel, getResponseData map[string]interface{}) *NstcpprofileResourceModel {
	tflog.Debug(ctx, "In nstcpprofileSetAttrFromGet Function")

	// Convert API response to model.
	// else-branches only null a value that is currently Unknown; a known
	// (configured) value that NITRO omits from GET is preserved to avoid the
	// omit-on-default inconsistent-result trap.
	if val, ok := getResponseData["ackaggregation"]; ok && val != nil {
		data.Ackaggregation = types.StringValue(val.(string))
	} else if data.Ackaggregation.IsUnknown() {
		data.Ackaggregation = types.StringNull()
	}
	if val, ok := getResponseData["ackonpush"]; ok && val != nil {
		data.Ackonpush = types.StringValue(val.(string))
	} else if data.Ackonpush.IsUnknown() {
		data.Ackonpush = types.StringNull()
	}
	if val, ok := getResponseData["applyadaptivetcp"]; ok && val != nil {
		data.Applyadaptivetcp = types.StringValue(val.(string))
	} else if data.Applyadaptivetcp.IsUnknown() {
		data.Applyadaptivetcp = types.StringNull()
	}
	if val, ok := getResponseData["buffersize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Buffersize = types.Int64Value(intVal)
		}
	} else if data.Buffersize.IsUnknown() {
		data.Buffersize = types.Int64Null()
	}
	if val, ok := getResponseData["burstratecontrol"]; ok && val != nil {
		data.Burstratecontrol = types.StringValue(val.(string))
	} else if data.Burstratecontrol.IsUnknown() {
		data.Burstratecontrol = types.StringNull()
	}
	if val, ok := getResponseData["clientiptcpoption"]; ok && val != nil {
		data.Clientiptcpoption = types.StringValue(val.(string))
	} else if data.Clientiptcpoption.IsUnknown() {
		data.Clientiptcpoption = types.StringNull()
	}
	if val, ok := getResponseData["clientiptcpoptionnumber"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clientiptcpoptionnumber = types.Int64Value(intVal)
		}
	} else if data.Clientiptcpoptionnumber.IsUnknown() {
		data.Clientiptcpoptionnumber = types.Int64Null()
	}
	if val, ok := getResponseData["delayedack"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delayedack = types.Int64Value(intVal)
		}
	} else if data.Delayedack.IsUnknown() {
		data.Delayedack = types.Int64Null()
	}
	if val, ok := getResponseData["dropestconnontimeout"]; ok && val != nil {
		data.Dropestconnontimeout = types.StringValue(val.(string))
	} else if data.Dropestconnontimeout.IsUnknown() {
		data.Dropestconnontimeout = types.StringNull()
	}
	if val, ok := getResponseData["drophalfclosedconnontimeout"]; ok && val != nil {
		data.Drophalfclosedconnontimeout = types.StringValue(val.(string))
	} else if data.Drophalfclosedconnontimeout.IsUnknown() {
		data.Drophalfclosedconnontimeout = types.StringNull()
	}
	if val, ok := getResponseData["dsack"]; ok && val != nil {
		data.Dsack = types.StringValue(val.(string))
	} else if data.Dsack.IsUnknown() {
		data.Dsack = types.StringNull()
	}
	if val, ok := getResponseData["dupackthresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dupackthresh = types.Int64Value(intVal)
		}
	} else if data.Dupackthresh.IsUnknown() {
		data.Dupackthresh = types.Int64Null()
	}
	if val, ok := getResponseData["dynamicreceivebuffering"]; ok && val != nil {
		data.Dynamicreceivebuffering = types.StringValue(val.(string))
	} else if data.Dynamicreceivebuffering.IsUnknown() {
		data.Dynamicreceivebuffering = types.StringNull()
	}
	if val, ok := getResponseData["ecn"]; ok && val != nil {
		data.Ecn = types.StringValue(val.(string))
	} else if data.Ecn.IsUnknown() {
		data.Ecn = types.StringNull()
	}
	if val, ok := getResponseData["establishclientconn"]; ok && val != nil {
		data.Establishclientconn = types.StringValue(val.(string))
	} else if data.Establishclientconn.IsUnknown() {
		data.Establishclientconn = types.StringNull()
	}
	if val, ok := getResponseData["fack"]; ok && val != nil {
		data.Fack = types.StringValue(val.(string))
	} else if data.Fack.IsUnknown() {
		data.Fack = types.StringNull()
	}
	if val, ok := getResponseData["flavor"]; ok && val != nil {
		data.Flavor = types.StringValue(val.(string))
	} else if data.Flavor.IsUnknown() {
		data.Flavor = types.StringNull()
	}
	if val, ok := getResponseData["frto"]; ok && val != nil {
		data.Frto = types.StringValue(val.(string))
	} else if data.Frto.IsUnknown() {
		data.Frto = types.StringNull()
	}
	if val, ok := getResponseData["hystart"]; ok && val != nil {
		data.Hystart = types.StringValue(val.(string))
	} else if data.Hystart.IsUnknown() {
		data.Hystart = types.StringNull()
	}
	if val, ok := getResponseData["initialcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Initialcwnd = types.Int64Value(intVal)
		}
	} else if data.Initialcwnd.IsUnknown() {
		data.Initialcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["ka"]; ok && val != nil {
		data.Ka = types.StringValue(val.(string))
	} else if data.Ka.IsUnknown() {
		data.Ka = types.StringNull()
	}
	if val, ok := getResponseData["kaconnidletime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kaconnidletime = types.Int64Value(intVal)
		}
	} else if data.Kaconnidletime.IsUnknown() {
		data.Kaconnidletime = types.Int64Null()
	}
	if val, ok := getResponseData["kamaxprobes"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kamaxprobes = types.Int64Value(intVal)
		}
	} else if data.Kamaxprobes.IsUnknown() {
		data.Kamaxprobes = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kaprobeinterval = types.Int64Value(intVal)
		}
	} else if data.Kaprobeinterval.IsUnknown() {
		data.Kaprobeinterval = types.Int64Null()
	}
	if val, ok := getResponseData["kaprobeupdatelastactivity"]; ok && val != nil {
		data.Kaprobeupdatelastactivity = types.StringValue(val.(string))
	} else if data.Kaprobeupdatelastactivity.IsUnknown() {
		data.Kaprobeupdatelastactivity = types.StringNull()
	}
	if val, ok := getResponseData["maxburst"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxburst = types.Int64Value(intVal)
		}
	} else if data.Maxburst.IsUnknown() {
		data.Maxburst = types.Int64Null()
	}
	if val, ok := getResponseData["maxcwnd"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxcwnd = types.Int64Value(intVal)
		}
	} else if data.Maxcwnd.IsUnknown() {
		data.Maxcwnd = types.Int64Null()
	}
	if val, ok := getResponseData["maxpktpermss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpktpermss = types.Int64Value(intVal)
		}
	} else if data.Maxpktpermss.IsUnknown() {
		data.Maxpktpermss = types.Int64Null()
	}
	if val, ok := getResponseData["minrto"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrto = types.Int64Value(intVal)
		}
	} else if data.Minrto.IsUnknown() {
		data.Minrto = types.Int64Null()
	}
	if val, ok := getResponseData["mpcapablecbit"]; ok && val != nil {
		data.Mpcapablecbit = types.StringValue(val.(string))
	} else if data.Mpcapablecbit.IsUnknown() {
		data.Mpcapablecbit = types.StringNull()
	}
	if val, ok := getResponseData["mptcp"]; ok && val != nil {
		data.Mptcp = types.StringValue(val.(string))
	} else if data.Mptcp.IsUnknown() {
		data.Mptcp = types.StringNull()
	}
	if val, ok := getResponseData["mptcpdropdataonpreestsf"]; ok && val != nil {
		data.Mptcpdropdataonpreestsf = types.StringValue(val.(string))
	} else if data.Mptcpdropdataonpreestsf.IsUnknown() {
		data.Mptcpdropdataonpreestsf = types.StringNull()
	}
	if val, ok := getResponseData["mptcpfastopen"]; ok && val != nil {
		data.Mptcpfastopen = types.StringValue(val.(string))
	} else if data.Mptcpfastopen.IsUnknown() {
		data.Mptcpfastopen = types.StringNull()
	}
	if val, ok := getResponseData["mptcpsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mptcpsessiontimeout = types.Int64Value(intVal)
		}
	} else if data.Mptcpsessiontimeout.IsUnknown() {
		data.Mptcpsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["mss"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mss = types.Int64Value(intVal)
		}
	} else if data.Mss.IsUnknown() {
		data.Mss = types.Int64Null()
	}
	if val, ok := getResponseData["nagle"]; ok && val != nil {
		data.Nagle = types.StringValue(val.(string))
	} else if data.Nagle.IsUnknown() {
		data.Nagle = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
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
	if val, ok := getResponseData["rateqmax"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rateqmax = types.Int64Value(intVal)
		}
	} else if data.Rateqmax.IsUnknown() {
		data.Rateqmax = types.Int64Null()
	}
	if val, ok := getResponseData["rfc5961compliance"]; ok && val != nil {
		data.Rfc5961compliance = types.StringValue(val.(string))
	} else if data.Rfc5961compliance.IsUnknown() {
		data.Rfc5961compliance = types.StringNull()
	}
	if val, ok := getResponseData["rstmaxack"]; ok && val != nil {
		data.Rstmaxack = types.StringValue(val.(string))
	} else if data.Rstmaxack.IsUnknown() {
		data.Rstmaxack = types.StringNull()
	}
	if val, ok := getResponseData["rstwindowattenuate"]; ok && val != nil {
		data.Rstwindowattenuate = types.StringValue(val.(string))
	} else if data.Rstwindowattenuate.IsUnknown() {
		data.Rstwindowattenuate = types.StringNull()
	}
	if val, ok := getResponseData["sack"]; ok && val != nil {
		data.Sack = types.StringValue(val.(string))
	} else if data.Sack.IsUnknown() {
		data.Sack = types.StringNull()
	}
	if val, ok := getResponseData["sendbuffsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sendbuffsize = types.Int64Value(intVal)
		}
	} else if data.Sendbuffsize.IsUnknown() {
		data.Sendbuffsize = types.Int64Null()
	}
	if val, ok := getResponseData["sendclientportintcpoption"]; ok && val != nil {
		data.Sendclientportintcpoption = types.StringValue(val.(string))
	} else if data.Sendclientportintcpoption.IsUnknown() {
		data.Sendclientportintcpoption = types.StringNull()
	}
	if val, ok := getResponseData["slowstartincr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartincr = types.Int64Value(intVal)
		}
	} else if data.Slowstartincr.IsUnknown() {
		data.Slowstartincr = types.Int64Null()
	}
	if val, ok := getResponseData["slowstartthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Slowstartthreshold = types.Int64Value(intVal)
		}
	} else if data.Slowstartthreshold.IsUnknown() {
		data.Slowstartthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["spoofsyndrop"]; ok && val != nil {
		data.Spoofsyndrop = types.StringValue(val.(string))
	} else if data.Spoofsyndrop.IsUnknown() {
		data.Spoofsyndrop = types.StringNull()
	}
	if val, ok := getResponseData["syncookie"]; ok && val != nil {
		data.Syncookie = types.StringValue(val.(string))
	} else if data.Syncookie.IsUnknown() {
		data.Syncookie = types.StringNull()
	}
	if val, ok := getResponseData["taillossprobe"]; ok && val != nil {
		data.Taillossprobe = types.StringValue(val.(string))
	} else if data.Taillossprobe.IsUnknown() {
		data.Taillossprobe = types.StringNull()
	}
	if val, ok := getResponseData["tcpfastopen"]; ok && val != nil {
		data.Tcpfastopen = types.StringValue(val.(string))
	} else if data.Tcpfastopen.IsUnknown() {
		data.Tcpfastopen = types.StringNull()
	}
	if val, ok := getResponseData["tcpfastopencookiesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpfastopencookiesize = types.Int64Value(intVal)
		}
	} else if data.Tcpfastopencookiesize.IsUnknown() {
		data.Tcpfastopencookiesize = types.Int64Null()
	}
	if val, ok := getResponseData["tcpmode"]; ok && val != nil {
		data.Tcpmode = types.StringValue(val.(string))
	} else if data.Tcpmode.IsUnknown() {
		data.Tcpmode = types.StringNull()
	}
	if val, ok := getResponseData["tcprate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcprate = types.Int64Value(intVal)
		}
	} else if data.Tcprate.IsUnknown() {
		data.Tcprate = types.Int64Null()
	}
	if val, ok := getResponseData["tcpsegoffload"]; ok && val != nil {
		data.Tcpsegoffload = types.StringValue(val.(string))
	} else if data.Tcpsegoffload.IsUnknown() {
		data.Tcpsegoffload = types.StringNull()
	}
	if val, ok := getResponseData["timestamp"]; ok && val != nil {
		data.Timestamp = types.StringValue(val.(string))
	} else if data.Timestamp.IsUnknown() {
		data.Timestamp = types.StringNull()
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
