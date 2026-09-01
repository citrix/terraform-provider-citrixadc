package sslocspresponder

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslocspresponderDataSourceModel is the data-source-specific model, decoupled
// from SslocspresponderResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslocspresponderDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Batchingdelay         types.Int64  `tfsdk:"batchingdelay"`
	Batchingdepth         types.Int64  `tfsdk:"batchingdepth"`
	Cache                 types.String `tfsdk:"cache"`
	Cachetimeout          types.Int64  `tfsdk:"cachetimeout"`
	Httpmethod            types.String `tfsdk:"httpmethod"`
	Insertclientcert      types.String `tfsdk:"insertclientcert"`
	Name                  types.String `tfsdk:"name"` // Required lookup key
	Ocspurlresolvetimeout types.Int64  `tfsdk:"ocspurlresolvetimeout"`
	Producedattimeskew    types.Int64  `tfsdk:"producedattimeskew"`
	Respondercert         types.String `tfsdk:"respondercert"`
	Resptimeout           types.Int64  `tfsdk:"resptimeout"`
	Signingcert           types.String `tfsdk:"signingcert"`
	Trustresponder        types.Bool   `tfsdk:"trustresponder"`
	Url                   types.String `tfsdk:"url"`
	Usenonce              types.String `tfsdk:"usenonce"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslocspresponder.json). Never settable; populated from GET.
	Ocspaiarefcount types.Int64  `tfsdk:"ocspaiarefcount"`
	Ocspipaddrstr   types.String `tfsdk:"ocspipaddrstr"`
	Port            types.Int64  `tfsdk:"port"`
}

func SslocspresponderDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"batchingdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch.  Does not apply if the Batching Depth is 1.",
			},
			"batchingdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of client certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.",
			},
			"cache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable caching of responses. Caching of responses received from the OCSP responder enables faster responses to the clients and reduces the load on the OCSP responder.",
			},
			"cachetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for caching the OCSP response. After the timeout, the Citrix ADC sends a fresh request to the OCSP responder for the certificate status. If a timeout is not specified, the timeout provided in the OCSP response applies.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod",
			},
			"insertclientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the complete client certificate in the OCSP request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the OCSP responder. Cannot begin with a hash (#) or space character and must contain only ASCII alphanumeric, underscore (_), hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the responder is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder\" or 'my responder').",
			},
			"ocspurlresolvetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.",
			},
			"producedattimeskew": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.",
			},
			"respondercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"resptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.",
			},
			"signingcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate-key pair that is used to sign OCSP requests. If this parameter is not set, the requests are not signed.",
			},
			"trustresponder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A certificate to use to validate OCSP responses.  Alternatively, if -trustResponder is specified, no verification will be done on the reponse.  If both are omitted, only the response times (producedAt, lastUpdate, nextUpdate) will be verified.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the OCSP responder.",
			},
			"usenonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the OCSP nonce extension, which is designed to prevent replay attacks.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"ocspaiarefcount": schema.Int64Attribute{
				Computed:    true,
				Description: "No of CA certs referencing this AIA responder.",
			},
			"ocspipaddrstr": schema.StringAttribute{
				Computed:    true,
				Description: "DNS resolved IP address.",
			},
			"port": schema.Int64Attribute{
				Computed:    true,
				Description: "Port number on which OCSP Server listens. Range 1 - 65535. * in CLI is represented as 65535 in NITRO API.",
			},
		},
	}
}

// sslocspresponderDataSourceSetAttrFromGet projects a NITRO sslocspresponder GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func sslocspresponderDataSourceSetAttrFromGet(ctx context.Context, data *SslocspresponderDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslocspresponderDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Batchingdelay = utils.MapGetInt64(g, "batchingdelay")
	data.Batchingdepth = utils.MapGetInt64(g, "batchingdepth")
	data.Cache = utils.MapGetString(g, "cache")
	data.Cachetimeout = utils.MapGetInt64(g, "cachetimeout")
	data.Httpmethod = utils.MapGetString(g, "httpmethod")
	data.Insertclientcert = utils.MapGetString(g, "insertclientcert")
	data.Ocspurlresolvetimeout = utils.MapGetInt64(g, "ocspurlresolvetimeout")
	data.Producedattimeskew = utils.MapGetInt64(g, "producedattimeskew")
	data.Respondercert = utils.MapGetString(g, "respondercert")
	data.Resptimeout = utils.MapGetInt64(g, "resptimeout")
	data.Signingcert = utils.MapGetString(g, "signingcert")
	data.Trustresponder = utils.MapGetBool(g, "trustresponder")
	data.Url = utils.MapGetString(g, "url")
	data.Usenonce = utils.MapGetString(g, "usenonce")

	// Read-only attributes.
	data.Ocspaiarefcount = utils.MapGetInt64(g, "ocspaiarefcount")
	data.Ocspipaddrstr = utils.MapGetString(g, "ocspipaddrstr")
	data.Port = utils.MapGetInt64(g, "port")
}
