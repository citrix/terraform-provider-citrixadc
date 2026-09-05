package sslprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslprofileDataSourceModel is the data-source-specific model, decoupled from
// SslprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the resource attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslprofileDataSourceModel struct {
	Id                                types.String `tfsdk:"id"`
	Allowextendedmastersecret         types.String `tfsdk:"allowextendedmastersecret"`
	Allowlegacykdf                    types.String `tfsdk:"allowlegacykdf"`
	Allowunknownsni                   types.String `tfsdk:"allowunknownsni"`
	Alpnprotocol                      types.String `tfsdk:"alpnprotocol"`
	Ciphername                        types.String `tfsdk:"ciphername"`
	Cipherpriority                    types.Int64  `tfsdk:"cipherpriority"`
	Cipherredirect                    types.String `tfsdk:"cipherredirect"`
	Cipherurl                         types.String `tfsdk:"cipherurl"`
	Cleartextport                     types.Int64  `tfsdk:"cleartextport"`
	Clientauth                        types.String `tfsdk:"clientauth"`
	Clientauthuseboundcachain         types.String `tfsdk:"clientauthuseboundcachain"`
	Clientcert                        types.String `tfsdk:"clientcert"`
	Commonname                        types.String `tfsdk:"commonname"`
	Defaultsni                        types.String `tfsdk:"defaultsni"`
	Denysslreneg                      types.String `tfsdk:"denysslreneg"`
	Dh                                types.String `tfsdk:"dh"`
	Dhcount                           types.Int64  `tfsdk:"dhcount"`
	Dhekeyexchangewithpsk             types.String `tfsdk:"dhekeyexchangewithpsk"`
	Dhfile                            types.String `tfsdk:"dhfile"`
	Dhkeyexpsizelimit                 types.String `tfsdk:"dhkeyexpsizelimit"`
	Dropreqwithnohostheader           types.String `tfsdk:"dropreqwithnohostheader"`
	Dynamicclientcert                 types.String `tfsdk:"dynamicclientcert"`
	Ecccurvebindings                  types.Set    `tfsdk:"ecccurvebindings"`
	Encryptedclienthello              types.String `tfsdk:"encryptedclienthello"`
	Encrypttriggerpktcount            types.Int64  `tfsdk:"encrypttriggerpktcount"`
	Ersa                              types.String `tfsdk:"ersa"`
	Ersacount                         types.Int64  `tfsdk:"ersacount"`
	Hsts                              types.String `tfsdk:"hsts"`
	Includesubdomains                 types.String `tfsdk:"includesubdomains"`
	Insertionencoding                 types.String `tfsdk:"insertionencoding"`
	Maxage                            types.Int64  `tfsdk:"maxage"`
	Maxrenegrate                      types.Int64  `tfsdk:"maxrenegrate"`
	Name                              types.String `tfsdk:"name"` // Required lookup key
	Nodefaultbindings                 types.String `tfsdk:"nodefaultbindings"`
	Nodefaultcipherbindings           types.Bool   `tfsdk:"nodefaultcipherbindings"`
	Nodefaultecccurvebindings         types.Bool   `tfsdk:"nodefaultecccurvebindings"`
	Ocspstapling                      types.String `tfsdk:"ocspstapling"`
	Preload                           types.String `tfsdk:"preload"`
	Prevsessionkeylifetime            types.Int64  `tfsdk:"prevsessionkeylifetime"`
	Pushenctrigger                    types.String `tfsdk:"pushenctrigger"`
	Pushenctriggertimeout             types.Int64  `tfsdk:"pushenctriggertimeout"`
	Pushflag                          types.Int64  `tfsdk:"pushflag"`
	Quantumsize                       types.String `tfsdk:"quantumsize"`
	Redirectportrewrite               types.String `tfsdk:"redirectportrewrite"`
	Sendclosenotify                   types.String `tfsdk:"sendclosenotify"`
	Serverauth                        types.String `tfsdk:"serverauth"`
	Sessionkeylifetime                types.Int64  `tfsdk:"sessionkeylifetime"`
	Sessionticket                     types.String `tfsdk:"sessionticket"`
	Sessionticketkeydata              types.String `tfsdk:"sessionticketkeydata"`
	Sessionticketkeyrefresh           types.String `tfsdk:"sessionticketkeyrefresh"`
	Sessionticketlifetime             types.Int64  `tfsdk:"sessionticketlifetime"`
	Sessreuse                         types.String `tfsdk:"sessreuse"`
	Sesstimeout                       types.Int64  `tfsdk:"sesstimeout"`
	Skipclientcertpolicycheck         types.String `tfsdk:"skipclientcertpolicycheck"`
	Snienable                         types.String `tfsdk:"snienable"`
	Snihttphostmatch                  types.String `tfsdk:"snihttphostmatch"`
	Ssl3                              types.String `tfsdk:"ssl3"`
	Sslclientlogs                     types.String `tfsdk:"sslclientlogs"`
	Sslimaxsessperserver              types.Int64  `tfsdk:"sslimaxsessperserver"`
	Sslinterception                   types.String `tfsdk:"sslinterception"`
	Ssliocspcheck                     types.String `tfsdk:"ssliocspcheck"`
	Sslireneg                         types.String `tfsdk:"sslireneg"`
	Ssllogprofile                     types.String `tfsdk:"ssllogprofile"`
	Sslprofiletype                    types.String `tfsdk:"sslprofiletype"`
	Sslredirect                       types.String `tfsdk:"sslredirect"`
	Ssltriggertimeout                 types.Int64  `tfsdk:"ssltriggertimeout"`
	Strictcachecks                    types.String `tfsdk:"strictcachecks"`
	Strictclientekucheck              types.String `tfsdk:"strictclientekucheck"`
	Strictsigdigestcheck              types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                              types.String `tfsdk:"tls1"`
	Tls11                             types.String `tfsdk:"tls11"`
	Tls12                             types.String `tfsdk:"tls12"`
	Tls13                             types.String `tfsdk:"tls13"`
	Tls13sessionticketsperauthcontext types.Int64  `tfsdk:"tls13sessionticketsperauthcontext"`
	Zerorttearlydata                  types.String `tfsdk:"zerorttearlydata"`
	Cipherbindings                    types.Set    `tfsdk:"cipherbindings"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslprofile.json). Never settable; populated from GET.
	Nonfipsciphers               types.String `tfsdk:"nonfipsciphers"`
	Crlcheck                     types.String `tfsdk:"crlcheck"`
	Ocspcheck                    types.String `tfsdk:"ocspcheck"`
	Snicert                      types.Bool   `tfsdk:"snicert"`
	Skipcaname                   types.Bool   `tfsdk:"skipcaname"`
	Invoke                       types.Bool   `tfsdk:"invoke"`
	Labeltype                    types.String `tfsdk:"labeltype"`
	Service                      types.Int64  `tfsdk:"service"`
	Builtin                      types.List   `tfsdk:"builtin"`
	Feature                      types.String `tfsdk:"feature"`
	Sslpfobjecttype              types.Int64  `tfsdk:"sslpfobjecttype"`
	Ssliverifyservercertforreuse types.String `tfsdk:"ssliverifyservercertforreuse"`
}

func SslprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowextendedmastersecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to YES, attempt to use the TLS Extended Master Secret (EMS, as\ndescribed in RFC 7627) when negotiating TLS 1.0, TLS 1.1 and TLS 1.2\nconnection parameters. EMS must be supported by both the TLS client and server\nin order to be enabled during a handshake. This setting applies to both\nfrontend and backend SSL profiles.",
			},
			"allowlegacykdf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FIPS 140-3 certification requires all handshakes without EMS be blocked. Such KDFs are allowed by default. This setting is to allow/disallow such legacy KDFs when needed. This setting applies to both frontend and backend SSL profiles.",
			},
			"allowunknownsni": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls how the handshake is handled when the server name extension does not match any of the bound certificates. These checks are performed only if the session is SNI enabled (i.e. when profile bound to vserver has SNIEnable and Client Hello arrived with SNI extension). Available settings function as follows :\nENABLED   - handshakes with an unknown SNI are allowed to continue, if a default cert is bound.\nDISLABED  - handshakes with an unknown SNI are not allowed to continue.",
			},
			"alpnprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Application protocol supported by the server and used in negotiation of the protocol with the client. Possible values are HTTP1.1, HTTP2 and NONE. Default value is NONE which implies application protocol is not enabled hence remain unknown to the TLS layer. This parameter is relevant only if SSL connection is handled by the virtual server of the type SSL_TCP.",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cipher group/alias/individual cipher configuration",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "cipher priority",
			},
			"cipherredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.\nThis parameter is not applicable when configuring a backend profile.",
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
				Description: "State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"clientauthuseboundcachain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certficates bound on the VIP are used for validating the client cert. Certficates came along with client cert are not used for validating the client cert",
			},
			"clientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule for client certificate requirement in client authentication.",
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.",
			},
			"defaultsni": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default domain name supported by the SSL virtual server. The parameter is effective, when zero touch certificate management is active for the SSL virtual server i.e. no manual SNI cert or default server cert is bound to the v-server. For SSL transactions, when SNI is not presented by the client, server-certificate corresponding to the default SNI, if available in the cert-store, is selected else connection is terminated.",
			},
			"denysslreneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny renegotiation in specified circumstances. Available settings function as follows:\n* NO - Allow SSL renegotiation.\n* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.\n* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.\n* ALL - Deny all secure and nonsecure SSL renegotiation.\n* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.",
			},
			"dh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Diffie-Hellman (DH) key exchange.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"dhcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time.\nThis parameter is not applicable when configuring a backend profile. Allowed DH count values are 0 and >= 500.",
			},
			"dhekeyexchangewithpsk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake.\nA DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange.\nIf disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client.\nIf enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled.",
			},
			"dhfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The file name and path for the DH parameter.",
			},
			"dhkeyexpsizelimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.",
			},
			"dropreqwithnohostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.",
			},
			"dynamicclientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Dynamic Client Certificate Generation for SSL sessions.",
			},
			"ecccurvebindings": schema.SetAttribute{
				Computed:    true,
				Description: "Set of ECC curve names bound to the SSL profile.",
				ElementType: types.StringType,
			},
			"encryptedclienthello": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLS 1.3 Encrypted Client Hello Support",
			},
			"encrypttriggerpktcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.",
			},
			"ersa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"ersacount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  refresh  count  for the re-generation of RSA public-key and private-key pair.",
			},
			"hsts": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client",
			},
			"includesubdomains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains.",
			},
			"insertionencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encoding method used to insert the subject or issuer's name in HTTP requests to servers.",
			},
			"maxage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server",
			},
			"maxrenegrate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of renegotiation requests allowed, in one second, to each SSL entity to which this profile is bound. When set to 0, an unlimited number of renegotiation requests are allowed. Applicable only when Deny SSL renegotiation is set to a value other than ALL.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SSL profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "Control default bindings for the SSL profile.",
			},
			"nodefaultcipherbindings": schema.BoolAttribute{
				Computed:    true,
				Description: "When set to true, removes the default cipher bindings from the SSL profile.",
			},
			"nodefaultecccurvebindings": schema.BoolAttribute{
				Computed:    true,
				Description: "When set to true, removes the default ECC curve bindings from the SSL profile.",
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
			"prevsessionkeylifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option sets the life time of symm key used to generate session tickets issued by NS in secs",
			},
			"pushenctrigger": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:\n* ALWAYS - Any PUSH packet triggers encryption.\n* IGNORE - Ignore PUSH packet for triggering encryption.\n* MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption.\n* TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.",
			},
			"pushenctriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.",
			},
			"pushflag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:\n0 - Auto (PUSH flag is not set.)\n1 - Insert PUSH flag into every decrypted record.\n2 -Insert PUSH flag into every encrypted record.\n3 - Insert PUSH flag into every decrypted and encrypted record.",
			},
			"quantumsize": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable sending SSL Close-Notify at the end of a transaction.",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of server authentication support for the SSL Backend profile.",
			},
			"sessionkeylifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option sets the life time of symm key used to generate session tickets issued by NS in secs",
			},
			"sessionticket": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables the use of session tickets, as per the RFC 5077",
			},
			"sessionticketkeydata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Session ticket enc/dec key , admin can set it",
			},
			"sessionticketkeyrefresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables the use of session tickets, as per the RFC 5077",
			},
			"sessionticketlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option sets the life time of session tickets issued by NS in secs",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Session timeout value in seconds.",
			},
			"skipclientcertpolicycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"snihttphostmatch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.\nAvailable settings function as follows:\nCERT   - Request is forwarded if the 'Host' value is covered\n         by the certificate used to establish this SSL session.\n         Note: 'CERT' matching mode cannot be applied in\n         TLS 1.3 connections established by resuming from a\n         previous TLS 1.3 session. On these connections, 'STRICT'\n         matching mode will be used instead.\nSTRICT - Request is forwarded only if value of 'Host' header\n         in HTTP is identical to the 'Server name' value passed\n         in 'Client Hello' of the SSL connection.\nNO     - No validation is performed on the HTTP 'Host'\n         header value.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv3 protocol support for the SSL profile.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled, NetScaler will log the session ID and SNI name during SSL handshakes on both the external and internal interfaces.",
			},
			"sslimaxsessperserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum ssl session to be cached per dynamic origin server. A unique ssl session is created for each SNI received from the client on ClientHello and the matching session is used for server session reuse.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable transparent interception of SSL sessions.",
			},
			"ssliocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable OCSP check for origin server certificate.",
			},
			"sslireneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server.",
			},
			"ssllogprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the ssllogprofile.",
			},
			"sslprofiletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.",
			},
			"sslredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of HTTPS redirects for the SSL service.\nFor an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.\nIf SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"ssltriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.",
			},
			"strictcachecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable strict CA certificate checks on the appliance.",
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
				Description: "State of TLSv1.0 protocol support for the SSL profile.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.1 protocol support for the SSL profile.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.2 protocol support for the SSL profile.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.3 protocol support for the SSL profile.",
			},
			"tls13sessionticketsperauthcontext": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes.\nThis value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection.\nNo tickets are sent if resumption is disabled.",
			},
			"zerorttearlydata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake.\nEarly application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"nonfipsciphers": schema.StringAttribute{
				Computed:    true,
				Description: "State of usage of ciphers that are not FIPS approved. Valid only for an SSL service bound with a FIPS key and certificate.",
			},
			"crlcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional).",
			},
			"ocspcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional).",
			},
			"snicert": schema.BoolAttribute{
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"skipcaname": schema.BoolAttribute{
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.",
			},
			"invoke": schema.BoolAttribute{
				Computed:    true,
				Description: "Invoke flag. This attribute is relevant only for ADVANCED policies.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy label invocation. Possible values = vserver, service, policylabel",
			},
			"service": schema.Int64Attribute{
				Computed:    true,
				Description: "Service.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether ssl profile is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"sslpfobjecttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Internal flag to indicate what type of object binds this profile: monitor or service.",
			},
			"ssliverifyservercertforreuse": schema.StringAttribute{
				Computed:    true,
				Description: "Verify the origin server's certificate before reusing the front-end SSL session.",
			},
		},
		Blocks: map[string]schema.Block{
			"cipherbindings": schema.SetNestedBlock{
				Description: "Set of cipher bindings bound to the SSL profile.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"ciphername": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the cipher.",
						},
						"cipherpriority": schema.Int64Attribute{
							Computed:    true,
							Description: "Priority of the cipher binding.",
						},
					},
				},
			},
		},
	}
}

// sslprofileDataSourceSetAttrFromGet projects a NITRO sslprofile GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func sslprofileDataSourceSetAttrFromGet(ctx context.Context, data *SslprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Allowextendedmastersecret = utils.MapGetString(g, "allowextendedmastersecret")
	data.Allowlegacykdf = utils.MapGetString(g, "allowlegacykdf")
	data.Allowunknownsni = utils.MapGetString(g, "allowunknownsni")
	data.Alpnprotocol = utils.MapGetString(g, "alpnprotocol")
	data.Ciphername = utils.MapGetString(g, "ciphername")
	data.Cipherpriority = utils.MapGetInt64(g, "cipherpriority")
	data.Cipherredirect = utils.MapGetString(g, "cipherredirect")
	data.Cipherurl = utils.MapGetString(g, "cipherurl")
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")
	data.Clientauth = utils.MapGetString(g, "clientauth")
	data.Clientauthuseboundcachain = utils.MapGetString(g, "clientauthuseboundcachain")
	data.Clientcert = utils.MapGetString(g, "clientcert")
	data.Commonname = utils.MapGetString(g, "commonname")
	data.Defaultsni = utils.MapGetString(g, "defaultsni")
	data.Denysslreneg = utils.MapGetString(g, "denysslreneg")
	data.Dh = utils.MapGetString(g, "dh")
	data.Dhcount = utils.MapGetInt64(g, "dhcount")
	data.Dhekeyexchangewithpsk = utils.MapGetString(g, "dhekeyexchangewithpsk")
	data.Dhfile = utils.MapGetString(g, "dhfile")
	data.Dhkeyexpsizelimit = utils.MapGetString(g, "dhkeyexpsizelimit")
	data.Dropreqwithnohostheader = utils.MapGetString(g, "dropreqwithnohostheader")
	data.Dynamicclientcert = utils.MapGetString(g, "dynamicclientcert")
	data.Encryptedclienthello = utils.MapGetString(g, "encryptedclienthello")
	data.Encrypttriggerpktcount = utils.MapGetInt64(g, "encrypttriggerpktcount")
	data.Ersa = utils.MapGetString(g, "ersa")
	data.Ersacount = utils.MapGetInt64(g, "ersacount")
	data.Hsts = utils.MapGetString(g, "hsts")
	data.Includesubdomains = utils.MapGetString(g, "includesubdomains")
	data.Insertionencoding = utils.MapGetString(g, "insertionencoding")
	data.Maxage = utils.MapGetInt64(g, "maxage")
	data.Maxrenegrate = utils.MapGetInt64(g, "maxrenegrate")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Nodefaultcipherbindings = utils.MapGetBool(g, "nodefaultcipherbindings")
	data.Nodefaultecccurvebindings = utils.MapGetBool(g, "nodefaultecccurvebindings")
	data.Ocspstapling = utils.MapGetString(g, "ocspstapling")
	data.Preload = utils.MapGetString(g, "preload")
	data.Prevsessionkeylifetime = utils.MapGetInt64(g, "prevsessionkeylifetime")
	data.Pushenctrigger = utils.MapGetString(g, "pushenctrigger")
	data.Pushenctriggertimeout = utils.MapGetInt64(g, "pushenctriggertimeout")
	data.Pushflag = utils.MapGetInt64(g, "pushflag")
	data.Quantumsize = utils.MapGetString(g, "quantumsize")
	data.Redirectportrewrite = utils.MapGetString(g, "redirectportrewrite")
	data.Sendclosenotify = utils.MapGetString(g, "sendclosenotify")
	data.Serverauth = utils.MapGetString(g, "serverauth")
	data.Sessionkeylifetime = utils.MapGetInt64(g, "sessionkeylifetime")
	data.Sessionticket = utils.MapGetString(g, "sessionticket")
	data.Sessionticketkeyrefresh = utils.MapGetString(g, "sessionticketkeyrefresh")
	data.Sessionticketlifetime = utils.MapGetInt64(g, "sessionticketlifetime")
	data.Sessreuse = utils.MapGetString(g, "sessreuse")
	data.Sesstimeout = utils.MapGetInt64(g, "sesstimeout")
	data.Skipclientcertpolicycheck = utils.MapGetString(g, "skipclientcertpolicycheck")
	data.Snienable = utils.MapGetString(g, "snienable")
	data.Snihttphostmatch = utils.MapGetString(g, "snihttphostmatch")
	data.Ssl3 = utils.MapGetString(g, "ssl3")
	data.Sslclientlogs = utils.MapGetString(g, "sslclientlogs")
	data.Sslimaxsessperserver = utils.MapGetInt64(g, "sslimaxsessperserver")
	data.Sslinterception = utils.MapGetString(g, "sslinterception")
	data.Ssliocspcheck = utils.MapGetString(g, "ssliocspcheck")
	data.Sslireneg = utils.MapGetString(g, "sslireneg")
	data.Ssllogprofile = utils.MapGetString(g, "ssllogprofile")
	data.Sslprofiletype = utils.MapGetString(g, "sslprofiletype")
	data.Sslredirect = utils.MapGetString(g, "sslredirect")
	data.Ssltriggertimeout = utils.MapGetInt64(g, "ssltriggertimeout")
	data.Strictcachecks = utils.MapGetString(g, "strictcachecks")
	data.Strictclientekucheck = utils.MapGetString(g, "strictclientekucheck")
	data.Strictsigdigestcheck = utils.MapGetString(g, "strictsigdigestcheck")
	data.Tls1 = utils.MapGetString(g, "tls1")
	data.Tls11 = utils.MapGetString(g, "tls11")
	data.Tls12 = utils.MapGetString(g, "tls12")
	data.Tls13 = utils.MapGetString(g, "tls13")
	data.Tls13sessionticketsperauthcontext = utils.MapGetInt64(g, "tls13sessionticketsperauthcontext")
	data.Zerorttearlydata = utils.MapGetString(g, "zerorttearlydata")

	// sessionticketkeydata is a secret input the GET never returns -> Null.
	data.Sessionticketkeydata = types.StringNull()

	// ecccurvebindings and the cipherbindings block are not projected from the
	// scalar GET response; leave them as typed Null.
	data.Ecccurvebindings = types.SetNull(types.StringType)
	data.Cipherbindings = types.SetNull(types.ObjectType{AttrTypes: map[string]attr.Type{
		"ciphername":     types.StringType,
		"cipherpriority": types.Int64Type,
	}})

	// Read-only attributes.
	data.Nonfipsciphers = utils.MapGetString(g, "nonfipsciphers")
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Snicert = utils.MapGetBool(g, "snicert")
	data.Skipcaname = utils.MapGetBool(g, "skipcaname")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Service = utils.MapGetInt64(g, "service")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Sslpfobjecttype = utils.MapGetInt64(g, "sslpfobjecttype")
	data.Ssliverifyservercertforreuse = utils.MapGetString(g, "ssliverifyservercertforreuse")
}
