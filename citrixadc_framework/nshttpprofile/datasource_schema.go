package nshttpprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NshttpprofileDataSourceModel is the data-source-specific model, decoupled from
// NshttpprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (refcnt, builtin, apdexsvrresptimethreshold, dropinvalreqswarning,
// feature). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type NshttpprofileDataSourceModel struct {
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
	Http2smallwndtimeout             types.Int64  `tfsdk:"http2smallwndtimeout"`
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
	Name                             types.String `tfsdk:"name"` // Required lookup key
	Normalizeurl                     types.String `tfsdk:"normalizeurl"`
	Normalizeurltoorigin             types.String `tfsdk:"normalizeurltoorigin"`
	Passprotocolupgrade              types.String `tfsdk:"passprotocolupgrade"`
	Persistentetag                   types.String `tfsdk:"persistentetag"`
	Reqtimeout                       types.Int64  `tfsdk:"reqtimeout"`
	Reqtimeoutaction                 types.String `tfsdk:"reqtimeoutaction"`
	Reusepooltimeout                 types.Int64  `tfsdk:"reusepooltimeout"`
	Rtsptunnel                       types.String `tfsdk:"rtsptunnel"`
	Weblog                           types.String `tfsdk:"weblog"`
	Websocket                        types.String `tfsdk:"websocket"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/nshttpprofile.json). Never settable; populated from GET.
	Refcnt                    types.Int64  `tfsdk:"refcnt"`
	Builtin                   types.List   `tfsdk:"builtin"`
	Apdexsvrresptimethreshold types.Int64  `tfsdk:"apdexsvrresptimethreshold"`
	Dropinvalreqswarning      types.String `tfsdk:"dropinvalreqswarning"`
	Feature                   types.String `tfsdk:"feature"`
}

func NshttpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"adpttimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.",
			},
			"allowonlywordcharactersandhyphen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled allows only the word characters [A-Za-z0-9_] and hyphen [-] in the request/response header names and the connection will be reset for the other characters. When disabled allows any visible (printing) characters (%21-%7E) except delimiters (double quotes and \"(),/:;<=>?@[]{}\").",
			},
			"altsvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for Alternative Services.",
			},
			"altsvcvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure a custom Alternative Services header value that should be inserted in the response to advertise a HTTP/SSL/HTTP_QUIC vserver.",
			},
			"apdexcltresptimethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4*T is considered tolerable. Any transaction responding in more than 4*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.",
			},
			"clientiphdrexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header that contains the real client IP address.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Start data compression on receiving a TCP packet with PUSH flag set.",
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropextracrlf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any extra 'CR' and 'LF' characters present after the header.",
			},
			"dropextradata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any extra data when server sends more data than the specified content-length.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop invalid HTTP requests or responses.",
			},
			"grpcholdlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size in bytes allowed to buffer gRPC packets till trailer is received",
			},
			"grpcholdtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time in milliseconds allowed to buffer gRPC packets till trailer is received. The value should be in multiples of 100",
			},
			"grpclengthdelimitation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set to DISABLED for gRPC without a length delimitation.",
			},
			"hostheadervalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validates the length of the Host header and its syntax. Also includes validation of the port number if specified",
			},
			"http2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for HTTP/2.",
			},
			"http2altsvcframe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for sending HTTP/2 ALTSVC frames. When enabled, the ADC sends HTTP/2 ALTSVC frames to HTTP/2 clients, instead of the Alt-Svc response header field. Not applicable to servers.",
			},
			"http2direct": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for Direct HTTP/2.",
			},
			"http2extendedconnect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable HTTP/2 Extended CONNECT mechanism.",
			},
			"http2headertablesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the header compression table used to decode header blocks, in bytes.",
			},
			"http2initialconnwindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial window size for connection level flow control, in bytes.",
			},
			"http2initialwindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial window size for stream level flow control, in bytes.",
			},
			"http2maxconcurrentstreams": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent streams that is allowed per connection.",
			},
			"http2maxemptyframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of empty frames allowed in HTTP2 connection per minute",
			},
			"http2maxframesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the frame payload that the Citrix ADC is willing to receive, in bytes.",
			},
			"http2maxheaderlistsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of header list that the Citrix ADC is prepared to accept, in bytes. NOTE: The actual plain text header size that the Citrix ADC accepts is limited by maxHeaderLen. Please change maxHeaderLen parameter as well when modifying http2MaxHeaderListSize.",
			},
			"http2maxpingframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of PING frames allowed in HTTP2 connection per minute",
			},
			"http2maxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of outgoing RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2maxrxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of incoming RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2smallwndtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout (in seconds) for HTTP/2 small-window stalled streams. Required to mitigate CVE-2026-13474.",
			},
			"http2maxsettingsframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SETTINGS frames allowed in HTTP2 connection per minute",
			},
			"http2minseverconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of HTTP2 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/2 connections.",
			},
			"http2strictcipher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable strict HTTP/2 cipher selection",
			},
			"http3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for HTTP/3.",
			},
			"http3maxheaderblockedstreams": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of HTTP/3 streams that can be blocked while HTTP/3 headers are being decoded.",
			},
			"http3maxheaderfieldsectionsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the HTTP/3 header field section, in bytes.",
			},
			"http3maxheadertablesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the HTTP/3 QPACK dynamic header table, in bytes.",
			},
			"http3minseverconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of HTTP/3 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/3 connections.",
			},
			"http3webtransport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for WebTransport over HTTP/3.",
			},
			"httppipelinebuffsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Application pipeline request buffering size, in bytes.",
			},
			"incomphdrdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time to wait, in milliseconds, between incomplete header packets. If the header packets take longer to arrive at Citrix ADC, the connection is silently dropped.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"markhttpheaderextrawserror": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark Http header with extra white space as invalid",
			},
			"markrfc7230noncompliantinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark RFC7230 non-compliant transaction as invalid",
			},
			"marktracereqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark TRACE requests as invalid.",
			},
			"maxduplicateheaderfields": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of allowed occurrences of header fields that share the same field name. This threshold is enforced only for well-known header field names recognized by NetScaler. If the value is set to 0, then it will be similar to previous behavior, Where we store only 15 duplicate headers and rest are parsed and send to the server.",
			},
			"maxheaderfieldlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bytes allowed for header field for HTTP header. If number of bytes exceeds beyond configured value, then request will be marked invalid",
			},
			"maxheaderlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
			"normalizeurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable RFC 3986 normalization of incoming URL before validation or consumption.",
			},
			"normalizeurltoorigin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable RFC 3986 URL normalization for request sent to the origin server.",
			},
			"passprotocolupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pass protocol upgrade request to the server.",
			},
			"persistentetag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection.",
			},
			"weblog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable web logging.",
			},
			"websocket": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this profile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if http profile is built-in or not. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"apdexsvrresptimethreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Satisfactory threshold (T) for server response time in milliseconds to be used for APDEX calculations.",
			},
			"dropinvalreqswarning": schema.StringAttribute{
				Computed:    true,
				Description: "Display warning if Drop invalid reqs is disabled in the profile.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nshttpprofileDataSourceSetAttrFromGet projects a NITRO nshttpprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func nshttpprofileDataSourceSetAttrFromGet(ctx context.Context, data *NshttpprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nshttpprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Adpttimeout = utils.MapGetString(g, "adpttimeout")
	data.Allowonlywordcharactersandhyphen = utils.MapGetString(g, "allowonlywordcharactersandhyphen")
	data.Altsvc = utils.MapGetString(g, "altsvc")
	data.Altsvcvalue = utils.MapGetString(g, "altsvcvalue")
	data.Apdexcltresptimethreshold = utils.MapGetInt64(g, "apdexcltresptimethreshold")
	data.Clientiphdrexpr = utils.MapGetString(g, "clientiphdrexpr")
	data.Cmponpush = utils.MapGetString(g, "cmponpush")
	data.Conmultiplex = utils.MapGetString(g, "conmultiplex")
	data.Dropextracrlf = utils.MapGetString(g, "dropextracrlf")
	data.Dropextradata = utils.MapGetString(g, "dropextradata")
	data.Dropinvalreqs = utils.MapGetString(g, "dropinvalreqs")
	data.Grpcholdlimit = utils.MapGetInt64(g, "grpcholdlimit")
	data.Grpcholdtimeout = utils.MapGetInt64(g, "grpcholdtimeout")
	data.Grpclengthdelimitation = utils.MapGetString(g, "grpclengthdelimitation")
	data.Hostheadervalidation = utils.MapGetString(g, "hostheadervalidation")
	data.Http2 = utils.MapGetString(g, "http2")
	data.Http2altsvcframe = utils.MapGetString(g, "http2altsvcframe")
	data.Http2direct = utils.MapGetString(g, "http2direct")
	data.Http2extendedconnect = utils.MapGetString(g, "http2extendedconnect")
	data.Http2headertablesize = utils.MapGetInt64(g, "http2headertablesize")
	data.Http2initialconnwindowsize = utils.MapGetInt64(g, "http2initialconnwindowsize")
	data.Http2initialwindowsize = utils.MapGetInt64(g, "http2initialwindowsize")
	data.Http2maxconcurrentstreams = utils.MapGetInt64(g, "http2maxconcurrentstreams")
	data.Http2maxemptyframespermin = utils.MapGetInt64(g, "http2maxemptyframespermin")
	data.Http2maxframesize = utils.MapGetInt64(g, "http2maxframesize")
	data.Http2maxheaderlistsize = utils.MapGetInt64(g, "http2maxheaderlistsize")
	data.Http2maxpingframespermin = utils.MapGetInt64(g, "http2maxpingframespermin")
	data.Http2maxresetframespermin = utils.MapGetInt64(g, "http2maxresetframespermin")
	data.Http2maxrxresetframespermin = utils.MapGetInt64(g, "http2maxrxresetframespermin")
	data.Http2maxsettingsframespermin = utils.MapGetInt64(g, "http2maxsettingsframespermin")
	data.Http2minseverconn = utils.MapGetInt64(g, "http2minseverconn")
	data.Http2smallwndtimeout = utils.MapGetInt64(g, "http2smallwndtimeout")
	data.Http2strictcipher = utils.MapGetString(g, "http2strictcipher")
	data.Http3 = utils.MapGetString(g, "http3")
	data.Http3maxheaderblockedstreams = utils.MapGetInt64(g, "http3maxheaderblockedstreams")
	data.Http3maxheaderfieldsectionsize = utils.MapGetInt64(g, "http3maxheaderfieldsectionsize")
	data.Http3maxheadertablesize = utils.MapGetInt64(g, "http3maxheadertablesize")
	data.Http3minseverconn = utils.MapGetInt64(g, "http3minseverconn")
	data.Http3webtransport = utils.MapGetString(g, "http3webtransport")
	data.Httppipelinebuffsize = utils.MapGetInt64(g, "httppipelinebuffsize")
	data.Incomphdrdelay = utils.MapGetInt64(g, "incomphdrdelay")
	data.Markconnreqinval = utils.MapGetString(g, "markconnreqinval")
	data.Markhttp09inval = utils.MapGetString(g, "markhttp09inval")
	data.Markhttpheaderextrawserror = utils.MapGetString(g, "markhttpheaderextrawserror")
	data.Markrfc7230noncompliantinval = utils.MapGetString(g, "markrfc7230noncompliantinval")
	data.Marktracereqinval = utils.MapGetString(g, "marktracereqinval")
	data.Maxduplicateheaderfields = utils.MapGetInt64(g, "maxduplicateheaderfields")
	data.Maxheaderfieldlen = utils.MapGetInt64(g, "maxheaderfieldlen")
	data.Maxheaderlen = utils.MapGetInt64(g, "maxheaderlen")
	data.Maxreq = utils.MapGetInt64(g, "maxreq")
	data.Maxreusepool = utils.MapGetInt64(g, "maxreusepool")
	data.Minreusepool = utils.MapGetInt64(g, "minreusepool")
	data.Normalizeurl = utils.MapGetString(g, "normalizeurl")
	data.Normalizeurltoorigin = utils.MapGetString(g, "normalizeurltoorigin")
	data.Passprotocolupgrade = utils.MapGetString(g, "passprotocolupgrade")
	data.Persistentetag = utils.MapGetString(g, "persistentetag")
	data.Reqtimeout = utils.MapGetInt64(g, "reqtimeout")
	data.Reqtimeoutaction = utils.MapGetString(g, "reqtimeoutaction")
	data.Reusepooltimeout = utils.MapGetInt64(g, "reusepooltimeout")
	data.Rtsptunnel = utils.MapGetString(g, "rtsptunnel")
	data.Weblog = utils.MapGetString(g, "weblog")
	data.Websocket = utils.MapGetString(g, "websocket")

	// Read-only (GET-only) metadata.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Apdexsvrresptimethreshold = utils.MapGetInt64(g, "apdexsvrresptimethreshold")
	data.Dropinvalreqswarning = utils.MapGetString(g, "dropinvalreqswarning")
	data.Feature = utils.MapGetString(g, "feature")
}
