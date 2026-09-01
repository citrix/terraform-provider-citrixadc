package sslservicegroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslservicegroupDataSourceModel is the data-source-specific model, decoupled
// from SslservicegroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the resource attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslservicegroupDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Commonname           types.String `tfsdk:"commonname"`
	Ocspstapling         types.String `tfsdk:"ocspstapling"`
	Sendclosenotify      types.String `tfsdk:"sendclosenotify"`
	Serverauth           types.String `tfsdk:"serverauth"`
	Servicegroupname     types.String `tfsdk:"servicegroupname"` // Required lookup key
	Sessreuse            types.String `tfsdk:"sessreuse"`
	Sesstimeout          types.Int64  `tfsdk:"sesstimeout"`
	Snienable            types.String `tfsdk:"snienable"`
	Ssl3                 types.String `tfsdk:"ssl3"`
	Sslclientlogs        types.String `tfsdk:"sslclientlogs"`
	Sslprofile           types.String `tfsdk:"sslprofile"`
	Strictsigdigestcheck types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                 types.String `tfsdk:"tls1"`
	Tls11                types.String `tfsdk:"tls11"`
	Tls12                types.String `tfsdk:"tls12"`
	Tls13                types.String `tfsdk:"tls13"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslservicegroup.json). Never settable; populated from GET.
	Dh                  types.String `tfsdk:"dh"`
	Dhfile              types.String `tfsdk:"dhfile"`
	Dhcount             types.Int64  `tfsdk:"dhcount"`
	Dhkeyexpsizelimit   types.String `tfsdk:"dhkeyexpsizelimit"`
	Ersa                types.String `tfsdk:"ersa"`
	Ersacount           types.Int64  `tfsdk:"ersacount"`
	Cipherredirect      types.String `tfsdk:"cipherredirect"`
	Cipherurl           types.String `tfsdk:"cipherurl"`
	Sslv2redirect       types.String `tfsdk:"sslv2redirect"`
	Sslv2url            types.String `tfsdk:"sslv2url"`
	Clientauth          types.String `tfsdk:"clientauth"`
	Clientcert          types.String `tfsdk:"clientcert"`
	Sslredirect         types.String `tfsdk:"sslredirect"`
	Redirectportrewrite types.String `tfsdk:"redirectportrewrite"`
	Nonfipsciphers      types.String `tfsdk:"nonfipsciphers"`
	Ssl2                types.String `tfsdk:"ssl2"`
	Ocspcheck           types.String `tfsdk:"ocspcheck"`
	Crlcheck            types.String `tfsdk:"crlcheck"`
	Cleartextport       types.Int64  `tfsdk:"cleartextport"`
	Servicename         types.String `tfsdk:"servicename"`
	Ca                  types.Bool   `tfsdk:"ca"`
	Snicert             types.Bool   `tfsdk:"snicert"`
	Quicflag            types.Bool   `tfsdk:"quicflag"`
}

func SslservicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server",
			},
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:\nENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.\nDISABLED: The appliance does not check the status of the server certificate.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of server authentication support for the SSL service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service group for which to set advanced configuration.",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the Server Name Indication (SNI) feature on the service. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv3 protocol support for the SSL service group.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the Service Group.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter indicating to check whether peer's certificate is signed with one of signature-hash combination supported by Citrix ADC",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.0 protocol support for the SSL service group.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.1 protocol support for the SSL service group.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.2 protocol support for the SSL service group.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.3 protocol support for the SSL service group.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"dh": schema.StringAttribute{
				Computed:    true,
				Description: "The state of DH key exchange support for the SSL service group.",
			},
			"dhfile": schema.StringAttribute{
				Computed:    true,
				Description: "The file name and path for the DH parameter.",
			},
			"dhcount": schema.Int64Attribute{
				Computed:    true,
				Description: "The refresh count for the re-generation of DH public-key and private-key from the DH parameter.",
			},
			"dhkeyexpsizelimit": schema.StringAttribute{
				Computed:    true,
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size.",
			},
			"ersa": schema.StringAttribute{
				Computed:    true,
				Description: "The state of Ephemeral RSA key exchange support for the SSL service group. Ephemeral RSA is used for export ciphers.",
			},
			"ersacount": schema.Int64Attribute{
				Computed:    true,
				Description: "The refresh count for the re-generation of RSA public-key and private-key pair.",
			},
			"cipherredirect": schema.StringAttribute{
				Computed:    true,
				Description: "The state of Cipher Redirect feature.",
			},
			"cipherurl": schema.StringAttribute{
				Computed:    true,
				Description: "The redirect URL to be used with the Cipher Redirect feature.",
			},
			"sslv2redirect": schema.StringAttribute{
				Computed:    true,
				Description: "The state of SSLv2 Redirect feature.",
			},
			"sslv2url": schema.StringAttribute{
				Computed:    true,
				Description: "The redirect URL to be used with SSLv2 Redirect feature.",
			},
			"clientauth": schema.StringAttribute{
				Computed:    true,
				Description: "The state of Client-Authentication support for the SSL service group.",
			},
			"clientcert": schema.StringAttribute{
				Computed:    true,
				Description: "The rule for client certificate requirement in client authentication.",
			},
			"sslredirect": schema.StringAttribute{
				Computed:    true,
				Description: "The state of HTTPS redirects for the SSL service group.",
			},
			"redirectportrewrite": schema.StringAttribute{
				Computed:    true,
				Description: "The state of port-rewrite feature.",
			},
			"nonfipsciphers": schema.StringAttribute{
				Computed:    true,
				Description: "The state of usage of non FIPS approved ciphers.",
			},
			"ssl2": schema.StringAttribute{
				Computed:    true,
				Description: "The state of SSLv2 protocol support for the SSL service group.",
			},
			"ocspcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional).",
			},
			"crlcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional).",
			},
			"cleartextport": schema.Int64Attribute{
				Computed:    true,
				Description: "The port on the back-end web-servers where the clear-text data is sent by system.",
			},
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "The service name.",
			},
			"ca": schema.BoolAttribute{
				Computed:    true,
				Description: "CA certificate.",
			},
			"snicert": schema.BoolAttribute{
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"quicflag": schema.BoolAttribute{
				Computed:    true,
				Description: "This flag is used to indicate the use of the QUIC transport protocol by a virtual server or service.",
			},
		},
	}
}

// sslservicegroupDataSourceSetAttrFromGet projects a NITRO sslservicegroup GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func sslservicegroupDataSourceSetAttrFromGet(ctx context.Context, data *SslservicegroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslservicegroupDataSourceSetAttrFromGet Function")

	if v, ok := g["servicegroupname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Servicegroupname = types.StringValue(utils.AnyToString(v))
	}

	data.Commonname = utils.MapGetString(g, "commonname")
	data.Ocspstapling = utils.MapGetString(g, "ocspstapling")
	data.Sendclosenotify = utils.MapGetString(g, "sendclosenotify")
	data.Serverauth = utils.MapGetString(g, "serverauth")
	data.Sessreuse = utils.MapGetString(g, "sessreuse")
	data.Sesstimeout = utils.MapGetInt64(g, "sesstimeout")
	data.Snienable = utils.MapGetString(g, "snienable")
	data.Ssl3 = utils.MapGetString(g, "ssl3")
	data.Sslclientlogs = utils.MapGetString(g, "sslclientlogs")
	data.Sslprofile = utils.MapGetString(g, "sslprofile")
	data.Strictsigdigestcheck = utils.MapGetString(g, "strictsigdigestcheck")
	data.Tls1 = utils.MapGetString(g, "tls1")
	data.Tls11 = utils.MapGetString(g, "tls11")
	data.Tls12 = utils.MapGetString(g, "tls12")
	data.Tls13 = utils.MapGetString(g, "tls13")

	// Read-only attributes.
	data.Dh = utils.MapGetString(g, "dh")
	data.Dhfile = utils.MapGetString(g, "dhfile")
	data.Dhcount = utils.MapGetInt64(g, "dhcount")
	data.Dhkeyexpsizelimit = utils.MapGetString(g, "dhkeyexpsizelimit")
	data.Ersa = utils.MapGetString(g, "ersa")
	data.Ersacount = utils.MapGetInt64(g, "ersacount")
	data.Cipherredirect = utils.MapGetString(g, "cipherredirect")
	data.Cipherurl = utils.MapGetString(g, "cipherurl")
	data.Sslv2redirect = utils.MapGetString(g, "sslv2redirect")
	data.Sslv2url = utils.MapGetString(g, "sslv2url")
	data.Clientauth = utils.MapGetString(g, "clientauth")
	data.Clientcert = utils.MapGetString(g, "clientcert")
	data.Sslredirect = utils.MapGetString(g, "sslredirect")
	data.Redirectportrewrite = utils.MapGetString(g, "redirectportrewrite")
	data.Nonfipsciphers = utils.MapGetString(g, "nonfipsciphers")
	data.Ssl2 = utils.MapGetString(g, "ssl2")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Ca = utils.MapGetBool(g, "ca")
	data.Snicert = utils.MapGetBool(g, "snicert")
	data.Quicflag = utils.MapGetBool(g, "quicflag")
}
