package nshttpprofile

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

// NshttpprofileResourceModel describes the resource data model.
type NshttpprofileResourceModel struct {
	Id                               types.String `tfsdk:"id"`
	Adpttimeout                      types.String `tfsdk:"adpttimeout"`
	Allowonlywordcharactersandhyphen types.String `tfsdk:"allowonlywordcharactersandhyphen"`
	Altsvc                           types.String `tfsdk:"altsvc"`
	Altsvcvalue                      types.String `tfsdk:"altsvcvalue"`
	Apdexcltresptimethreshold        types.Int64  `tfsdk:"apdexcltresptimethreshold"`
	Clientiphdrexpr                  types.String `tfsdk:"clientiphdrexpr"`
	Cmponpush                        types.String `tfsdk:"cmponpush"`
	Conmultiplex                     types.String `tfsdk:"conmultiplex"`
	Dropextracrlf                    types.String `tfsdk:"dropextracrlf"`
	Dropextradata                    types.String `tfsdk:"dropextradata"`
	Dropinvalreqs                    types.String `tfsdk:"dropinvalreqs"`
	Grpcholdlimit                    types.Int64  `tfsdk:"grpcholdlimit"`
	Grpcholdtimeout                  types.Int64  `tfsdk:"grpcholdtimeout"`
	Grpclengthdelimitation           types.String `tfsdk:"grpclengthdelimitation"`
	Hostheadervalidation             types.String `tfsdk:"hostheadervalidation"`
	Http2                            types.String `tfsdk:"http2"`
	Http2altsvcframe                 types.String `tfsdk:"http2altsvcframe"`
	Http2direct                      types.String `tfsdk:"http2direct"`
	Http2extendedconnect             types.String `tfsdk:"http2extendedconnect"`
	Http2headertablesize             types.Int64  `tfsdk:"http2headertablesize"`
	Http2initialconnwindowsize       types.Int64  `tfsdk:"http2initialconnwindowsize"`
	Http2initialwindowsize           types.Int64  `tfsdk:"http2initialwindowsize"`
	Http2maxconcurrentstreams        types.Int64  `tfsdk:"http2maxconcurrentstreams"`
	Http2maxemptyframespermin        types.Int64  `tfsdk:"http2maxemptyframespermin"`
	Http2maxframesize                types.Int64  `tfsdk:"http2maxframesize"`
	Http2maxheaderlistsize           types.Int64  `tfsdk:"http2maxheaderlistsize"`
	Http2maxpingframespermin         types.Int64  `tfsdk:"http2maxpingframespermin"`
	Http2maxresetframespermin        types.Int64  `tfsdk:"http2maxresetframespermin"`
	Http2maxrxresetframespermin      types.Int64  `tfsdk:"http2maxrxresetframespermin"`
	Http2maxsettingsframespermin     types.Int64  `tfsdk:"http2maxsettingsframespermin"`
	Http2minseverconn                types.Int64  `tfsdk:"http2minseverconn"`
	Http2strictcipher                types.String `tfsdk:"http2strictcipher"`
	Http3                            types.String `tfsdk:"http3"`
	Http3maxheaderblockedstreams     types.Int64  `tfsdk:"http3maxheaderblockedstreams"`
	Http3maxheaderfieldsectionsize   types.Int64  `tfsdk:"http3maxheaderfieldsectionsize"`
	Http3maxheadertablesize          types.Int64  `tfsdk:"http3maxheadertablesize"`
	Http3minseverconn                types.Int64  `tfsdk:"http3minseverconn"`
	Http3webtransport                types.String `tfsdk:"http3webtransport"`
	Httppipelinebuffsize             types.Int64  `tfsdk:"httppipelinebuffsize"`
	Incomphdrdelay                   types.Int64  `tfsdk:"incomphdrdelay"`
	Markconnreqinval                 types.String `tfsdk:"markconnreqinval"`
	Markhttp09inval                  types.String `tfsdk:"markhttp09inval"`
	Markhttpheaderextrawserror       types.String `tfsdk:"markhttpheaderextrawserror"`
	Markrfc7230noncompliantinval     types.String `tfsdk:"markrfc7230noncompliantinval"`
	Marktracereqinval                types.String `tfsdk:"marktracereqinval"`
	Maxduplicateheaderfields         types.Int64  `tfsdk:"maxduplicateheaderfields"`
	Maxheaderfieldlen                types.Int64  `tfsdk:"maxheaderfieldlen"`
	Maxheaderlen                     types.Int64  `tfsdk:"maxheaderlen"`
	Maxreq                           types.Int64  `tfsdk:"maxreq"`
	Maxreusepool                     types.Int64  `tfsdk:"maxreusepool"`
	Minreusepool                     types.Int64  `tfsdk:"minreusepool"`
	Name                             types.String `tfsdk:"name"`
	Passprotocolupgrade              types.String `tfsdk:"passprotocolupgrade"`
	Persistentetag                   types.String `tfsdk:"persistentetag"`
	Reqtimeout                       types.Int64  `tfsdk:"reqtimeout"`
	Reqtimeoutaction                 types.String `tfsdk:"reqtimeoutaction"`
	Reusepooltimeout                 types.Int64  `tfsdk:"reusepooltimeout"`
	Rtsptunnel                       types.String `tfsdk:"rtsptunnel"`
	Weblog                           types.String `tfsdk:"weblog"`
	Websocket                        types.String `tfsdk:"websocket"`
}

func (r *NshttpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nshttpprofile resource.",
			},
			"adpttimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.",
			},
			"allowonlywordcharactersandhyphen": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "When enabled allows only the word characters [A-Za-z0-9_] and hyphen [-] in the request/response header names and the connection will be reset for the other characters. When disabled allows any visible (printing) characters (%21-%7E) except delimiters (double quotes and \"(),/:;<=>?@[]{}\").",
			},
			"altsvc": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for Alternative Services.",
			},
			"altsvcvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure a custom Alternative Services header value that should be inserted in the response to advertise a HTTP/SSL/HTTP_QUIC vserver.",
			},
			"apdexcltresptimethreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(500),
				Description: "This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4*T is considered tolerable. Any transaction responding in more than 4*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.",
			},
			"clientiphdrexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header that contains the real client IP address.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Start data compression on receiving a TCP packet with PUSH flag set.",
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropextracrlf": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Drop any extra 'CR' and 'LF' characters present after the header.",
			},
			"dropextradata": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Drop any extra data when server sends more data than the specified content-length.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Drop invalid HTTP requests or responses.",
			},
			"grpcholdlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(131072),
				Description: "Maximum size in bytes allowed to buffer gRPC packets till trailer is received",
			},
			"grpcholdtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Maximum time in milliseconds allowed to buffer gRPC packets till trailer is received. The value should be in multiples of 100",
			},
			"grpclengthdelimitation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Set to DISABLED for gRPC without a length delimitation.",
			},
			"hostheadervalidation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Validates the length of the Host header and its syntax. Also includes validation of the port number if specified",
			},
			"http2": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for HTTP/2.",
			},
			"http2altsvcframe": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for sending HTTP/2 ALTSVC frames. When enabled, the ADC sends HTTP/2 ALTSVC frames to HTTP/2 clients, instead of the Alt-Svc response header field. Not applicable to servers.",
			},
			"http2direct": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for Direct HTTP/2.",
			},
			"http2extendedconnect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Choose whether to enable HTTP/2 Extended CONNECT mechanism.",
			},
			"http2headertablesize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4096),
				Description: "Maximum size of the header compression table used to decode header blocks, in bytes.",
			},
			"http2initialconnwindowsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Initial window size for connection level flow control, in bytes.",
			},
			"http2initialwindowsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65535),
				Description: "Initial window size for stream level flow control, in bytes.",
			},
			"http2maxconcurrentstreams": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of concurrent streams that is allowed per connection.",
			},
			"http2maxemptyframespermin": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "Maximum number of empty frames allowed in HTTP2 connection per minute",
			},
			"http2maxframesize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(16384),
				Description: "Maximum size of the frame payload that the Citrix ADC is willing to receive, in bytes.",
			},
			"http2maxheaderlistsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(24576),
				Description: "Maximum size of header list that the Citrix ADC is prepared to accept, in bytes. NOTE: The actual plain text header size that the Citrix ADC accepts is limited by maxHeaderLen. Please change maxHeaderLen parameter as well when modifying http2MaxHeaderListSize.",
			},
			"http2maxpingframespermin": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "Maximum number of PING frames allowed in HTTP2 connection per minute",
			},
			"http2maxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(90),
				Description: "Maximum number of outgoing RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2maxrxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of incoming RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2maxsettingsframespermin": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(15),
				Description: "Maximum number of SETTINGS frames allowed in HTTP2 connection per minute",
			},
			"http2minseverconn": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20),
				Description: "Minimum number of HTTP2 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/2 connections.",
			},
			"http2strictcipher": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Choose whether to enable strict HTTP/2 cipher selection",
			},
			"http3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for HTTP/3.",
			},
			"http3maxheaderblockedstreams": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of HTTP/3 streams that can be blocked while HTTP/3 headers are being decoded.",
			},
			"http3maxheaderfieldsectionsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(24576),
				Description: "Maximum size of the HTTP/3 header field section, in bytes.",
			},
			"http3maxheadertablesize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4096),
				Description: "Maximum size of the HTTP/3 QPACK dynamic header table, in bytes.",
			},
			"http3minseverconn": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20),
				Description: "Minimum number of HTTP/3 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/3 connections.",
			},
			"http3webtransport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Choose whether to enable support for WebTransport over HTTP/3.",
			},
			"httppipelinebuffsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(131072),
				Description: "Application pipeline request buffering size, in bytes.",
			},
			"incomphdrdelay": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(7000),
				Description: "Maximum time to wait, in milliseconds, between incomplete header packets. If the header packets take longer to arrive at Citrix ADC, the connection is silently dropped.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"markhttpheaderextrawserror": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mark Http header with extra white space as invalid",
			},
			"markrfc7230noncompliantinval": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mark RFC7230 non-compliant transaction as invalid",
			},
			"marktracereqinval": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mark TRACE requests as invalid.",
			},
			"maxduplicateheaderfields": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of allowed occurrences of header fields that share the same field name. This threshold is enforced only for well-known header field names recognized by NetScaler. If the value is set to 0, then it will be similar to previous behavior, Where we store only 15 duplicate headers and rest are parsed and send to the server.",
			},
			"maxheaderfieldlen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(24820),
				Description: "Number of bytes allowed for header field for HTTP header. If number of bytes exceeds beyond configured value, then request will be marked invalid",
			},
			"maxheaderlen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(24820),
				Description: "Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests allowed on a single connection. Zero implies no limit on the number of requests.",
			},
			"maxreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size. If non-zero value is given, it has to be greater than or equal to the number of running Packet Engines.",
			},
			"minreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for an HTTP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a HTTP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my http profile\" or 'my http profile'\\).",
			},
			"passprotocolupgrade": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Pass protocol upgrade request to the server.",
			},
			"persistentetag": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header.",
			},
			"reqtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.",
			},
			"reqtimeoutaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions:\n* RESET - Send RST (reset) to client when timeout occurs.\n* DROP - Drop silently when timeout occurs.\n* Custom responder action - Name of the responder action to trigger when timeout occurs, used to send custom message.",
			},
			"reusepooltimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.",
			},
			"rtsptunnel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection.",
			},
			"weblog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable web logging.",
			},
			"websocket": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection.",
			},
		},
	}
}

func nshttpprofileGetThePayloadFromtheConfig(ctx context.Context, data *NshttpprofileResourceModel) ns.Nshttpprofile {
	tflog.Debug(ctx, "In nshttpprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nshttpprofile := ns.Nshttpprofile{}
	if !data.Adpttimeout.IsNull() {
		nshttpprofile.Adpttimeout = data.Adpttimeout.ValueString()
	}
	if !data.Allowonlywordcharactersandhyphen.IsNull() {
		nshttpprofile.Allowonlywordcharactersandhyphen = data.Allowonlywordcharactersandhyphen.ValueString()
	}
	if !data.Altsvc.IsNull() {
		nshttpprofile.Altsvc = data.Altsvc.ValueString()
	}
	if !data.Altsvcvalue.IsNull() {
		nshttpprofile.Altsvcvalue = data.Altsvcvalue.ValueString()
	}
	if !data.Apdexcltresptimethreshold.IsNull() {
		nshttpprofile.Apdexcltresptimethreshold = utils.IntPtr(int(data.Apdexcltresptimethreshold.ValueInt64()))
	}
	if !data.Clientiphdrexpr.IsNull() {
		nshttpprofile.Clientiphdrexpr = data.Clientiphdrexpr.ValueString()
	}
	if !data.Cmponpush.IsNull() {
		nshttpprofile.Cmponpush = data.Cmponpush.ValueString()
	}
	if !data.Conmultiplex.IsNull() {
		nshttpprofile.Conmultiplex = data.Conmultiplex.ValueString()
	}
	if !data.Dropextracrlf.IsNull() {
		nshttpprofile.Dropextracrlf = data.Dropextracrlf.ValueString()
	}
	if !data.Dropextradata.IsNull() {
		nshttpprofile.Dropextradata = data.Dropextradata.ValueString()
	}
	if !data.Dropinvalreqs.IsNull() {
		nshttpprofile.Dropinvalreqs = data.Dropinvalreqs.ValueString()
	}
	if !data.Grpcholdlimit.IsNull() {
		nshttpprofile.Grpcholdlimit = utils.IntPtr(int(data.Grpcholdlimit.ValueInt64()))
	}
	if !data.Grpcholdtimeout.IsNull() {
		nshttpprofile.Grpcholdtimeout = utils.IntPtr(int(data.Grpcholdtimeout.ValueInt64()))
	}
	if !data.Grpclengthdelimitation.IsNull() {
		nshttpprofile.Grpclengthdelimitation = data.Grpclengthdelimitation.ValueString()
	}
	if !data.Hostheadervalidation.IsNull() {
		nshttpprofile.Hostheadervalidation = data.Hostheadervalidation.ValueString()
	}
	if !data.Http2.IsNull() {
		nshttpprofile.Http2 = data.Http2.ValueString()
	}
	if !data.Http2altsvcframe.IsNull() {
		nshttpprofile.Http2altsvcframe = data.Http2altsvcframe.ValueString()
	}
	if !data.Http2direct.IsNull() {
		nshttpprofile.Http2direct = data.Http2direct.ValueString()
	}
	if !data.Http2extendedconnect.IsNull() {
		nshttpprofile.Http2extendedconnect = data.Http2extendedconnect.ValueString()
	}
	if !data.Http2headertablesize.IsNull() {
		nshttpprofile.Http2headertablesize = utils.IntPtr(int(data.Http2headertablesize.ValueInt64()))
	}
	if !data.Http2initialconnwindowsize.IsNull() {
		nshttpprofile.Http2initialconnwindowsize = utils.IntPtr(int(data.Http2initialconnwindowsize.ValueInt64()))
	}
	if !data.Http2initialwindowsize.IsNull() {
		nshttpprofile.Http2initialwindowsize = utils.IntPtr(int(data.Http2initialwindowsize.ValueInt64()))
	}
	if !data.Http2maxconcurrentstreams.IsNull() {
		nshttpprofile.Http2maxconcurrentstreams = utils.IntPtr(int(data.Http2maxconcurrentstreams.ValueInt64()))
	}
	if !data.Http2maxemptyframespermin.IsNull() {
		nshttpprofile.Http2maxemptyframespermin = utils.IntPtr(int(data.Http2maxemptyframespermin.ValueInt64()))
	}
	if !data.Http2maxframesize.IsNull() {
		nshttpprofile.Http2maxframesize = utils.IntPtr(int(data.Http2maxframesize.ValueInt64()))
	}
	if !data.Http2maxheaderlistsize.IsNull() {
		nshttpprofile.Http2maxheaderlistsize = utils.IntPtr(int(data.Http2maxheaderlistsize.ValueInt64()))
	}
	if !data.Http2maxpingframespermin.IsNull() {
		nshttpprofile.Http2maxpingframespermin = utils.IntPtr(int(data.Http2maxpingframespermin.ValueInt64()))
	}
	if !data.Http2maxresetframespermin.IsNull() {
		nshttpprofile.Http2maxresetframespermin = utils.IntPtr(int(data.Http2maxresetframespermin.ValueInt64()))
	}
	if !data.Http2maxrxresetframespermin.IsNull() {
		nshttpprofile.Http2maxrxresetframespermin = utils.IntPtr(int(data.Http2maxrxresetframespermin.ValueInt64()))
	}
	if !data.Http2maxsettingsframespermin.IsNull() {
		nshttpprofile.Http2maxsettingsframespermin = utils.IntPtr(int(data.Http2maxsettingsframespermin.ValueInt64()))
	}
	if !data.Http2minseverconn.IsNull() {
		nshttpprofile.Http2minseverconn = utils.IntPtr(int(data.Http2minseverconn.ValueInt64()))
	}
	if !data.Http2strictcipher.IsNull() {
		nshttpprofile.Http2strictcipher = data.Http2strictcipher.ValueString()
	}
	if !data.Http3.IsNull() {
		nshttpprofile.Http3 = data.Http3.ValueString()
	}
	if !data.Http3maxheaderblockedstreams.IsNull() {
		nshttpprofile.Http3maxheaderblockedstreams = utils.IntPtr(int(data.Http3maxheaderblockedstreams.ValueInt64()))
	}
	if !data.Http3maxheaderfieldsectionsize.IsNull() {
		nshttpprofile.Http3maxheaderfieldsectionsize = utils.IntPtr(int(data.Http3maxheaderfieldsectionsize.ValueInt64()))
	}
	if !data.Http3maxheadertablesize.IsNull() {
		nshttpprofile.Http3maxheadertablesize = utils.IntPtr(int(data.Http3maxheadertablesize.ValueInt64()))
	}
	if !data.Http3minseverconn.IsNull() {
		nshttpprofile.Http3minseverconn = utils.IntPtr(int(data.Http3minseverconn.ValueInt64()))
	}
	if !data.Http3webtransport.IsNull() {
		nshttpprofile.Http3webtransport = data.Http3webtransport.ValueString()
	}
	if !data.Httppipelinebuffsize.IsNull() {
		nshttpprofile.Httppipelinebuffsize = utils.IntPtr(int(data.Httppipelinebuffsize.ValueInt64()))
	}
	if !data.Incomphdrdelay.IsNull() {
		nshttpprofile.Incomphdrdelay = utils.IntPtr(int(data.Incomphdrdelay.ValueInt64()))
	}
	if !data.Markconnreqinval.IsNull() {
		nshttpprofile.Markconnreqinval = data.Markconnreqinval.ValueString()
	}
	if !data.Markhttp09inval.IsNull() {
		nshttpprofile.Markhttp09inval = data.Markhttp09inval.ValueString()
	}
	if !data.Markhttpheaderextrawserror.IsNull() {
		nshttpprofile.Markhttpheaderextrawserror = data.Markhttpheaderextrawserror.ValueString()
	}
	if !data.Markrfc7230noncompliantinval.IsNull() {
		nshttpprofile.Markrfc7230noncompliantinval = data.Markrfc7230noncompliantinval.ValueString()
	}
	if !data.Marktracereqinval.IsNull() {
		nshttpprofile.Marktracereqinval = data.Marktracereqinval.ValueString()
	}
	if !data.Maxduplicateheaderfields.IsNull() {
		nshttpprofile.Maxduplicateheaderfields = utils.IntPtr(int(data.Maxduplicateheaderfields.ValueInt64()))
	}
	if !data.Maxheaderfieldlen.IsNull() {
		nshttpprofile.Maxheaderfieldlen = utils.IntPtr(int(data.Maxheaderfieldlen.ValueInt64()))
	}
	if !data.Maxheaderlen.IsNull() {
		nshttpprofile.Maxheaderlen = utils.IntPtr(int(data.Maxheaderlen.ValueInt64()))
	}
	if !data.Maxreq.IsNull() {
		nshttpprofile.Maxreq = utils.IntPtr(int(data.Maxreq.ValueInt64()))
	}
	if !data.Maxreusepool.IsNull() {
		nshttpprofile.Maxreusepool = utils.IntPtr(int(data.Maxreusepool.ValueInt64()))
	}
	if !data.Minreusepool.IsNull() {
		nshttpprofile.Minreusepool = utils.IntPtr(int(data.Minreusepool.ValueInt64()))
	}
	if !data.Name.IsNull() {
		nshttpprofile.Name = data.Name.ValueString()
	}
	if !data.Passprotocolupgrade.IsNull() {
		nshttpprofile.Passprotocolupgrade = data.Passprotocolupgrade.ValueString()
	}
	if !data.Persistentetag.IsNull() {
		nshttpprofile.Persistentetag = data.Persistentetag.ValueString()
	}
	if !data.Reqtimeout.IsNull() {
		nshttpprofile.Reqtimeout = utils.IntPtr(int(data.Reqtimeout.ValueInt64()))
	}
	if !data.Reqtimeoutaction.IsNull() {
		nshttpprofile.Reqtimeoutaction = data.Reqtimeoutaction.ValueString()
	}
	if !data.Reusepooltimeout.IsNull() {
		nshttpprofile.Reusepooltimeout = utils.IntPtr(int(data.Reusepooltimeout.ValueInt64()))
	}
	if !data.Rtsptunnel.IsNull() {
		nshttpprofile.Rtsptunnel = data.Rtsptunnel.ValueString()
	}
	if !data.Weblog.IsNull() {
		nshttpprofile.Weblog = data.Weblog.ValueString()
	}
	if !data.Websocket.IsNull() {
		nshttpprofile.Websocket = data.Websocket.ValueString()
	}

	return nshttpprofile
}

func nshttpprofileSetAttrFromGet(ctx context.Context, data *NshttpprofileResourceModel, getResponseData map[string]interface{}) *NshttpprofileResourceModel {
	tflog.Debug(ctx, "In nshttpprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["adpttimeout"]; ok && val != nil {
		data.Adpttimeout = types.StringValue(val.(string))
	} else {
		data.Adpttimeout = types.StringNull()
	}
	if val, ok := getResponseData["allowonlywordcharactersandhyphen"]; ok && val != nil {
		data.Allowonlywordcharactersandhyphen = types.StringValue(val.(string))
	} else {
		data.Allowonlywordcharactersandhyphen = types.StringNull()
	}
	if val, ok := getResponseData["altsvc"]; ok && val != nil {
		data.Altsvc = types.StringValue(val.(string))
	} else {
		data.Altsvc = types.StringNull()
	}
	if val, ok := getResponseData["altsvcvalue"]; ok && val != nil {
		data.Altsvcvalue = types.StringValue(val.(string))
	} else {
		data.Altsvcvalue = types.StringNull()
	}
	if val, ok := getResponseData["apdexcltresptimethreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apdexcltresptimethreshold = types.Int64Value(intVal)
		}
	} else {
		data.Apdexcltresptimethreshold = types.Int64Null()
	}
	if val, ok := getResponseData["clientiphdrexpr"]; ok && val != nil {
		data.Clientiphdrexpr = types.StringValue(val.(string))
	} else {
		data.Clientiphdrexpr = types.StringNull()
	}
	if val, ok := getResponseData["cmponpush"]; ok && val != nil {
		data.Cmponpush = types.StringValue(val.(string))
	} else {
		data.Cmponpush = types.StringNull()
	}
	if val, ok := getResponseData["conmultiplex"]; ok && val != nil {
		data.Conmultiplex = types.StringValue(val.(string))
	} else {
		data.Conmultiplex = types.StringNull()
	}
	if val, ok := getResponseData["dropextracrlf"]; ok && val != nil {
		data.Dropextracrlf = types.StringValue(val.(string))
	} else {
		data.Dropextracrlf = types.StringNull()
	}
	if val, ok := getResponseData["dropextradata"]; ok && val != nil {
		data.Dropextradata = types.StringValue(val.(string))
	} else {
		data.Dropextradata = types.StringNull()
	}
	if val, ok := getResponseData["dropinvalreqs"]; ok && val != nil {
		data.Dropinvalreqs = types.StringValue(val.(string))
	} else {
		data.Dropinvalreqs = types.StringNull()
	}
	if val, ok := getResponseData["grpcholdlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grpcholdlimit = types.Int64Value(intVal)
		}
	} else {
		data.Grpcholdlimit = types.Int64Null()
	}
	if val, ok := getResponseData["grpcholdtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Grpcholdtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Grpcholdtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["grpclengthdelimitation"]; ok && val != nil {
		data.Grpclengthdelimitation = types.StringValue(val.(string))
	} else {
		data.Grpclengthdelimitation = types.StringNull()
	}
	if val, ok := getResponseData["hostheadervalidation"]; ok && val != nil {
		data.Hostheadervalidation = types.StringValue(val.(string))
	} else {
		data.Hostheadervalidation = types.StringNull()
	}
	if val, ok := getResponseData["http2"]; ok && val != nil {
		data.Http2 = types.StringValue(val.(string))
	} else {
		data.Http2 = types.StringNull()
	}
	if val, ok := getResponseData["http2altsvcframe"]; ok && val != nil {
		data.Http2altsvcframe = types.StringValue(val.(string))
	} else {
		data.Http2altsvcframe = types.StringNull()
	}
	if val, ok := getResponseData["http2direct"]; ok && val != nil {
		data.Http2direct = types.StringValue(val.(string))
	} else {
		data.Http2direct = types.StringNull()
	}
	if val, ok := getResponseData["http2extendedconnect"]; ok && val != nil {
		data.Http2extendedconnect = types.StringValue(val.(string))
	} else {
		data.Http2extendedconnect = types.StringNull()
	}
	if val, ok := getResponseData["http2headertablesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2headertablesize = types.Int64Value(intVal)
		}
	} else {
		data.Http2headertablesize = types.Int64Null()
	}
	if val, ok := getResponseData["http2initialconnwindowsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2initialconnwindowsize = types.Int64Value(intVal)
		}
	} else {
		data.Http2initialconnwindowsize = types.Int64Null()
	}
	if val, ok := getResponseData["http2initialwindowsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2initialwindowsize = types.Int64Value(intVal)
		}
	} else {
		data.Http2initialwindowsize = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxconcurrentstreams"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxconcurrentstreams = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxconcurrentstreams = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxemptyframespermin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxemptyframespermin = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxemptyframespermin = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxframesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxframesize = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxframesize = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxheaderlistsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxheaderlistsize = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxheaderlistsize = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxpingframespermin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxpingframespermin = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxpingframespermin = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxresetframespermin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxresetframespermin = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxresetframespermin = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxrxresetframespermin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxrxresetframespermin = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxrxresetframespermin = types.Int64Null()
	}
	if val, ok := getResponseData["http2maxsettingsframespermin"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2maxsettingsframespermin = types.Int64Value(intVal)
		}
	} else {
		data.Http2maxsettingsframespermin = types.Int64Null()
	}
	if val, ok := getResponseData["http2minseverconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http2minseverconn = types.Int64Value(intVal)
		}
	} else {
		data.Http2minseverconn = types.Int64Null()
	}
	if val, ok := getResponseData["http2strictcipher"]; ok && val != nil {
		data.Http2strictcipher = types.StringValue(val.(string))
	} else {
		data.Http2strictcipher = types.StringNull()
	}
	if val, ok := getResponseData["http3"]; ok && val != nil {
		data.Http3 = types.StringValue(val.(string))
	} else {
		data.Http3 = types.StringNull()
	}
	if val, ok := getResponseData["http3maxheaderblockedstreams"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http3maxheaderblockedstreams = types.Int64Value(intVal)
		}
	} else {
		data.Http3maxheaderblockedstreams = types.Int64Null()
	}
	if val, ok := getResponseData["http3maxheaderfieldsectionsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http3maxheaderfieldsectionsize = types.Int64Value(intVal)
		}
	} else {
		data.Http3maxheaderfieldsectionsize = types.Int64Null()
	}
	if val, ok := getResponseData["http3maxheadertablesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http3maxheadertablesize = types.Int64Value(intVal)
		}
	} else {
		data.Http3maxheadertablesize = types.Int64Null()
	}
	if val, ok := getResponseData["http3minseverconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Http3minseverconn = types.Int64Value(intVal)
		}
	} else {
		data.Http3minseverconn = types.Int64Null()
	}
	if val, ok := getResponseData["http3webtransport"]; ok && val != nil {
		data.Http3webtransport = types.StringValue(val.(string))
	} else {
		data.Http3webtransport = types.StringNull()
	}
	if val, ok := getResponseData["httppipelinebuffsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httppipelinebuffsize = types.Int64Value(intVal)
		}
	} else {
		data.Httppipelinebuffsize = types.Int64Null()
	}
	if val, ok := getResponseData["incomphdrdelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Incomphdrdelay = types.Int64Value(intVal)
		}
	} else {
		data.Incomphdrdelay = types.Int64Null()
	}
	if val, ok := getResponseData["markconnreqinval"]; ok && val != nil {
		data.Markconnreqinval = types.StringValue(val.(string))
	} else {
		data.Markconnreqinval = types.StringNull()
	}
	if val, ok := getResponseData["markhttp09inval"]; ok && val != nil {
		data.Markhttp09inval = types.StringValue(val.(string))
	} else {
		data.Markhttp09inval = types.StringNull()
	}
	if val, ok := getResponseData["markhttpheaderextrawserror"]; ok && val != nil {
		data.Markhttpheaderextrawserror = types.StringValue(val.(string))
	} else {
		data.Markhttpheaderextrawserror = types.StringNull()
	}
	if val, ok := getResponseData["markrfc7230noncompliantinval"]; ok && val != nil {
		data.Markrfc7230noncompliantinval = types.StringValue(val.(string))
	} else {
		data.Markrfc7230noncompliantinval = types.StringNull()
	}
	if val, ok := getResponseData["marktracereqinval"]; ok && val != nil {
		data.Marktracereqinval = types.StringValue(val.(string))
	} else {
		data.Marktracereqinval = types.StringNull()
	}
	if val, ok := getResponseData["maxduplicateheaderfields"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxduplicateheaderfields = types.Int64Value(intVal)
		}
	} else {
		data.Maxduplicateheaderfields = types.Int64Null()
	}
	if val, ok := getResponseData["maxheaderfieldlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxheaderfieldlen = types.Int64Value(intVal)
		}
	} else {
		data.Maxheaderfieldlen = types.Int64Null()
	}
	if val, ok := getResponseData["maxheaderlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxheaderlen = types.Int64Value(intVal)
		}
	} else {
		data.Maxheaderlen = types.Int64Null()
	}
	if val, ok := getResponseData["maxreq"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreq = types.Int64Value(intVal)
		}
	} else {
		data.Maxreq = types.Int64Null()
	}
	if val, ok := getResponseData["maxreusepool"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreusepool = types.Int64Value(intVal)
		}
	} else {
		data.Maxreusepool = types.Int64Null()
	}
	if val, ok := getResponseData["minreusepool"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minreusepool = types.Int64Value(intVal)
		}
	} else {
		data.Minreusepool = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passprotocolupgrade"]; ok && val != nil {
		data.Passprotocolupgrade = types.StringValue(val.(string))
	} else {
		data.Passprotocolupgrade = types.StringNull()
	}
	if val, ok := getResponseData["persistentetag"]; ok && val != nil {
		data.Persistentetag = types.StringValue(val.(string))
	} else {
		data.Persistentetag = types.StringNull()
	}
	if val, ok := getResponseData["reqtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reqtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Reqtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["reqtimeoutaction"]; ok && val != nil {
		data.Reqtimeoutaction = types.StringValue(val.(string))
	} else {
		data.Reqtimeoutaction = types.StringNull()
	}
	if val, ok := getResponseData["reusepooltimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reusepooltimeout = types.Int64Value(intVal)
		}
	} else {
		data.Reusepooltimeout = types.Int64Null()
	}
	if val, ok := getResponseData["rtsptunnel"]; ok && val != nil {
		data.Rtsptunnel = types.StringValue(val.(string))
	} else {
		data.Rtsptunnel = types.StringNull()
	}
	if val, ok := getResponseData["weblog"]; ok && val != nil {
		data.Weblog = types.StringValue(val.(string))
	} else {
		data.Weblog = types.StringNull()
	}
	if val, ok := getResponseData["websocket"]; ok && val != nil {
		data.Websocket = types.StringValue(val.(string))
	} else {
		data.Websocket = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
