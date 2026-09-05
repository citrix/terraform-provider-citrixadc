package sslvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslvserverDataSourceModel is the data-source-specific model, decoupled from
// SslvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the resource attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslvserverDataSourceModel struct {
	Id                                types.String `tfsdk:"id"`
	Cipherredirect                    types.String `tfsdk:"cipherredirect"`
	Cipherurl                         types.String `tfsdk:"cipherurl"`
	Cleartextport                     types.Int64  `tfsdk:"cleartextport"`
	Clientauth                        types.String `tfsdk:"clientauth"`
	Clientcert                        types.String `tfsdk:"clientcert"`
	Defaultsni                        types.String `tfsdk:"defaultsni"`
	Dh                                types.String `tfsdk:"dh"`
	Dhcount                           types.Int64  `tfsdk:"dhcount"`
	Dhekeyexchangewithpsk             types.String `tfsdk:"dhekeyexchangewithpsk"`
	Dhfile                            types.String `tfsdk:"dhfile"`
	Dhkeyexpsizelimit                 types.String `tfsdk:"dhkeyexpsizelimit"`
	Dtls1                             types.String `tfsdk:"dtls1"`
	Dtls12                            types.String `tfsdk:"dtls12"`
	Dtlsprofilename                   types.String `tfsdk:"dtlsprofilename"`
	Ersa                              types.String `tfsdk:"ersa"`
	Ersacount                         types.Int64  `tfsdk:"ersacount"`
	Hsts                              types.String `tfsdk:"hsts"`
	Includesubdomains                 types.String `tfsdk:"includesubdomains"`
	Maxage                            types.Int64  `tfsdk:"maxage"`
	Ocspstapling                      types.String `tfsdk:"ocspstapling"`
	Preload                           types.String `tfsdk:"preload"`
	Pushenctrigger                    types.String `tfsdk:"pushenctrigger"`
	Redirectportrewrite               types.String `tfsdk:"redirectportrewrite"`
	Sendclosenotify                   types.String `tfsdk:"sendclosenotify"`
	Sessreuse                         types.String `tfsdk:"sessreuse"`
	Sesstimeout                       types.Int64  `tfsdk:"sesstimeout"`
	Snienable                         types.String `tfsdk:"snienable"`
	Ssl2                              types.String `tfsdk:"ssl2"`
	Ssl3                              types.String `tfsdk:"ssl3"`
	Sslclientlogs                     types.String `tfsdk:"sslclientlogs"`
	Sslprofile                        types.String `tfsdk:"sslprofile"`
	Sslredirect                       types.String `tfsdk:"sslredirect"`
	Sslv2redirect                     types.String `tfsdk:"sslv2redirect"`
	Sslv2url                          types.String `tfsdk:"sslv2url"`
	Strictclientekucheck              types.String `tfsdk:"strictclientekucheck"`
	Strictsigdigestcheck              types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                              types.String `tfsdk:"tls1"`
	Tls11                             types.String `tfsdk:"tls11"`
	Tls12                             types.String `tfsdk:"tls12"`
	Tls13                             types.String `tfsdk:"tls13"`
	Tls13sessionticketsperauthcontext types.Int64  `tfsdk:"tls13sessionticketsperauthcontext"`
	Vservername                       types.String `tfsdk:"vservername"` // Required lookup key
	Zerorttearlydata                  types.String `tfsdk:"zerorttearlydata"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslvserver.json). Never settable; populated from GET.
	Crlcheck         types.String `tfsdk:"crlcheck"`
	Nonfipsciphers   types.String `tfsdk:"nonfipsciphers"`
	Service          types.Int64  `tfsdk:"service"`
	Ocspcheck        types.String `tfsdk:"ocspcheck"`
	Ca               types.Bool   `tfsdk:"ca"`
	Snicert          types.Bool   `tfsdk:"snicert"`
	Skipcaname       types.Bool   `tfsdk:"skipcaname"`
	Dtlsflag         types.Bool   `tfsdk:"dtlsflag"`
	Quicflag         types.Bool   `tfsdk:"quicflag"`
	Skipcacertbundle types.Bool   `tfsdk:"skipcacertbundle"`
}

func SslvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipherredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Cipher Redirect. If cipher redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.",
			},
			"cipherurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The redirect URL to be used with the Cipher Redirect feature.",
			},
			"cleartextport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.",
			},
			"clientauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of client authentication. If client authentication is enabled, the virtual server terminates the SSL handshake if the SSL client does not provide a valid certificate.",
			},
			"clientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate.\nCaution: Define proper access control policies before changing this setting to Optional.",
			},
			"defaultsni": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default domain name supported by the SSL virtual server. The parameter is effective, when zero touch certificate management is active for the SSL virtual server i.e. no manual SNI cert or default server cert is bound to the v-server.\nFor SSL transactions, when SNI is not presented by the client, server-certificate corresponding to the default SNI, if available in the cert-store, is selected else connection is terminated.",
			},
			"dh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Diffie-Hellman (DH) key exchange.",
			},
			"dhcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time.",
			},
			"dhekeyexchangewithpsk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake.\nA DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange.\nIf disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client.\nIf enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled.",
			},
			"dhfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the DH parameter file, in PEM format, to be installed. /nsconfig/ssl/ is the default path.",
			},
			"dhkeyexpsizelimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.",
			},
			"dtls1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of DTLSv1.0 protocol support for the SSL Virtual Server.",
			},
			"dtls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of DTLSv1.2 protocol support for the SSL Virtual Server.",
			},
			"dtlsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DTLS profile whose settings are to be applied to the virtual server.",
			},
			"ersa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.",
			},
			"ersacount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Refresh count for regeneration of the RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).",
			},
			"hsts": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of HSTS protocol support for the SSL Virtual Server. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client",
			},
			"includesubdomains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains.",
			},
			"maxage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server",
			},
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:\nENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.\nDISABLED: The appliance does not check the status of the server certificate.",
			},
			"preload": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag indicates the consent of the site owner to have their domain preloaded.",
			},
			"pushenctrigger": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:\n* ALWAYS - Any PUSH packet triggers encryption.\n* IGNORE - Ignore PUSH packet for triggering encryption.\n* MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption.\n* TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the port rewrite while performing HTTPS redirect. If this parameter is ENABLED and the URL from the server does not contain the standard port, the port is rewritten to the standard.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
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
				Description: "State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"ssl2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv2 protocol support for the SSL Virtual Server.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv3 protocol support for the SSL Virtual Server.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the virtual server.",
			},
			"sslredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of HTTPS redirects for the SSL virtual server.\n\nFor an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.\nIf SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.",
			},
			"sslv2redirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv2 Redirect. If SSLv2 redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client.",
			},
			"sslv2url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.",
			},
			"strictclientekucheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable strict EKU extension check during client authentication.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.0 protocol support for the SSL Virtual Server.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.1 protocol support for the SSL Virtual Server.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.2 protocol support for the SSL Virtual Server.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.3 protocol support for the SSL Virtual Server.",
			},
			"tls13sessionticketsperauthcontext": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes.\nThis value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection.\nNo tickets are sent if resumption is disabled.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server for which to set advanced configuration.",
			},
			"zerorttearlydata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake.\nEarly application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"crlcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional).",
			},
			"nonfipsciphers": schema.StringAttribute{
				Computed:    true,
				Description: "The state of usage of non FIPS approved ciphers.",
			},
			"service": schema.Int64Attribute{
				Computed:    true,
				Description: "Service.",
			},
			"ocspcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional).",
			},
			"ca": schema.BoolAttribute{
				Computed:    true,
				Description: "CA certificate.",
			},
			"snicert": schema.BoolAttribute{
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"skipcaname": schema.BoolAttribute{
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.",
			},
			"dtlsflag": schema.BoolAttribute{
				Computed:    true,
				Description: "The flag is used to indicate whether DTLS is set or not.",
			},
			"quicflag": schema.BoolAttribute{
				Computed:    true,
				Description: "This flag is used to indicate the use of the QUIC transport protocol by an SSL virtual server or service.",
			},
			"skipcacertbundle": schema.BoolAttribute{
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.",
			},
		},
	}
}

// sslvserverDataSourceSetAttrFromGet projects a NITRO sslvserver GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func sslvserverDataSourceSetAttrFromGet(ctx context.Context, data *SslvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["vservername"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Vservername = types.StringValue(utils.AnyToString(v))
	}

	data.Cipherredirect = utils.MapGetString(g, "cipherredirect")
	data.Cipherurl = utils.MapGetString(g, "cipherurl")
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")
	data.Clientauth = utils.MapGetString(g, "clientauth")
	data.Clientcert = utils.MapGetString(g, "clientcert")
	data.Defaultsni = utils.MapGetString(g, "defaultsni")
	data.Dh = utils.MapGetString(g, "dh")
	data.Dhcount = utils.MapGetInt64(g, "dhcount")
	data.Dhekeyexchangewithpsk = utils.MapGetString(g, "dhekeyexchangewithpsk")
	data.Dhfile = utils.MapGetString(g, "dhfile")
	data.Dhkeyexpsizelimit = utils.MapGetString(g, "dhkeyexpsizelimit")
	data.Dtls1 = utils.MapGetString(g, "dtls1")
	data.Dtls12 = utils.MapGetString(g, "dtls12")
	data.Dtlsprofilename = utils.MapGetString(g, "dtlsprofilename")
	data.Ersa = utils.MapGetString(g, "ersa")
	data.Ersacount = utils.MapGetInt64(g, "ersacount")
	data.Hsts = utils.MapGetString(g, "hsts")
	data.Includesubdomains = utils.MapGetString(g, "includesubdomains")
	data.Maxage = utils.MapGetInt64(g, "maxage")
	data.Ocspstapling = utils.MapGetString(g, "ocspstapling")
	data.Preload = utils.MapGetString(g, "preload")
	data.Pushenctrigger = utils.MapGetString(g, "pushenctrigger")
	data.Redirectportrewrite = utils.MapGetString(g, "redirectportrewrite")
	data.Sendclosenotify = utils.MapGetString(g, "sendclosenotify")
	data.Sessreuse = utils.MapGetString(g, "sessreuse")
	data.Sesstimeout = utils.MapGetInt64(g, "sesstimeout")
	data.Snienable = utils.MapGetString(g, "snienable")
	data.Ssl2 = utils.MapGetString(g, "ssl2")
	data.Ssl3 = utils.MapGetString(g, "ssl3")
	data.Sslclientlogs = utils.MapGetString(g, "sslclientlogs")
	data.Sslprofile = utils.MapGetString(g, "sslprofile")
	data.Sslredirect = utils.MapGetString(g, "sslredirect")
	data.Sslv2redirect = utils.MapGetString(g, "sslv2redirect")
	data.Sslv2url = utils.MapGetString(g, "sslv2url")
	data.Strictclientekucheck = utils.MapGetString(g, "strictclientekucheck")
	data.Strictsigdigestcheck = utils.MapGetString(g, "strictsigdigestcheck")
	data.Tls1 = utils.MapGetString(g, "tls1")
	data.Tls11 = utils.MapGetString(g, "tls11")
	data.Tls12 = utils.MapGetString(g, "tls12")
	data.Tls13 = utils.MapGetString(g, "tls13")
	data.Tls13sessionticketsperauthcontext = utils.MapGetInt64(g, "tls13sessionticketsperauthcontext")
	data.Zerorttearlydata = utils.MapGetString(g, "zerorttearlydata")

	// Read-only attributes.
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
	data.Nonfipsciphers = utils.MapGetString(g, "nonfipsciphers")
	data.Service = utils.MapGetInt64(g, "service")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Ca = utils.MapGetBool(g, "ca")
	data.Snicert = utils.MapGetBool(g, "snicert")
	data.Skipcaname = utils.MapGetBool(g, "skipcaname")
	data.Dtlsflag = utils.MapGetBool(g, "dtlsflag")
	data.Quicflag = utils.MapGetBool(g, "quicflag")
	data.Skipcacertbundle = utils.MapGetBool(g, "skipcacertbundle")
}
