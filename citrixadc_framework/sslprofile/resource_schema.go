package sslprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

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

// SslprofileResourceModel describes the resource data model.
type SslprofileResourceModel struct {
	Id                                types.String `tfsdk:"id"`
	Allowextendedmastersecret         types.String `tfsdk:"allowextendedmastersecret"`
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
	Encryptedclienthello              types.String `tfsdk:"encryptedclienthello"`
	Encrypttriggerpktcount            types.Int64  `tfsdk:"encrypttriggerpktcount"`
	Ersa                              types.String `tfsdk:"ersa"`
	Ersacount                         types.Int64  `tfsdk:"ersacount"`
	Hsts                              types.String `tfsdk:"hsts"`
	Includesubdomains                 types.String `tfsdk:"includesubdomains"`
	Insertionencoding                 types.String `tfsdk:"insertionencoding"`
	Maxage                            types.Int64  `tfsdk:"maxage"`
	Maxrenegrate                      types.Int64  `tfsdk:"maxrenegrate"`
	Name                              types.String `tfsdk:"name"`
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
	Strictsigdigestcheck              types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                              types.String `tfsdk:"tls1"`
	Tls11                             types.String `tfsdk:"tls11"`
	Tls12                             types.String `tfsdk:"tls12"`
	Tls13                             types.String `tfsdk:"tls13"`
	Tls13sessionticketsperauthcontext types.Int64  `tfsdk:"tls13sessionticketsperauthcontext"`
	Zerorttearlydata                  types.String `tfsdk:"zerorttearlydata"`
}

func (r *SslprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslprofile resource.",
			},
			"allowextendedmastersecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to YES, attempt to use the TLS Extended Master Secret (EMS, as\ndescribed in RFC 7627) when negotiating TLS 1.0, TLS 1.1 and TLS 1.2\nconnection parameters. EMS must be supported by both the TLS client and server\nin order to be enabled during a handshake. This setting applies to both\nfrontend and backend SSL profiles.",
			},
			"allowunknownsni": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Controls how the handshake is handled when the server name extension does not match any of the bound certificates. These checks are performed only if the session is SNI enabled (i.e. when profile bound to vserver has SNIEnable and Client Hello arrived with SNI extension). Available settings function as follows :\nENABLED   - handshakes with an unknown SNI are allowed to continue, if a default cert is bound.\nDISLABED  - handshakes with an unknown SNI are not allowed to continue.",
			},
			"alpnprotocol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
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
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"clientauthuseboundcachain": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("ALL"),
				Description: "Deny renegotiation in specified circumstances. Available settings function as follows:\n* NO - Allow SSL renegotiation.\n* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.\n* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.\n* ALL - Deny all secure and nonsecure SSL renegotiation.\n* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.",
			},
			"dh": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.",
			},
			"dropreqwithnohostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.",
			},
			"encryptedclienthello": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLS 1.3 Encrypted Client Hello Support",
			},
			"encrypttriggerpktcount": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(45),
				Description: "Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.",
			},
			"ersa": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"ersacount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  refresh  count  for the re-generation of RSA public-key and private-key pair.",
			},
			"hsts": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client",
			},
			"includesubdomains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains.",
			},
			"insertionencoding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Unicode"),
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
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     int64default.StaticInt64(1),
				Description: "PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.",
			},
			"pushflag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:\n0 - Auto (PUSH flag is not set.)\n1 - Insert PUSH flag into every decrypted record.\n2 -Insert PUSH flag into every encrypted record.\n3 - Insert PUSH flag into every decrypted and encrypted record.",
			},
			"quantumsize": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("8192"),
				Description: "Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable sending SSL Close-Notify at the end of a transaction.",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of server authentication support for the SSL Backend profile.",
			},
			"sessionkeylifetime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3000),
				Description: "This option sets the life time of symm key used to generate session tickets issued by NS in secs",
			},
			"sessionticket": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option enables the use of session tickets, as per the RFC 5077",
			},
			"sessionticketkeydata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Session ticket enc/dec key , admin can set it",
			},
			"sessionticketkeyrefresh": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "This option enables the use of session tickets, as per the RFC 5077",
			},
			"sessionticketlifetime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "This option sets the life time of session tickets issued by NS in secs",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Session timeout value in seconds.",
			},
			"skipclientcertpolicycheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"snihttphostmatch": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CERT"),
				Description: "Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.\nAvailable settings function as follows:\nCERT   - Request is forwarded if the 'Host' value is covered\n         by the certificate used to establish this SSL session.\n         Note: 'CERT' matching mode cannot be applied in\n         TLS 1.3 connections established by resuming from a\n         previous TLS 1.3 session. On these connections, 'STRICT'\n         matching mode will be used instead.\nSTRICT - Request is forwarded only if value of 'Host' header\n         in HTTP is identical to the 'Server name' value passed\n         in 'Client Hello' of the SSL connection.\nNO     - No validation is performed on the HTTP 'Host'\n         header value.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of SSLv3 protocol support for the SSL profile.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "When enabled, NetScaler will log the session ID and SNI name during SSL handshakes on both the external and internal interfaces.",
			},
			"sslimaxsessperserver": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Maximum ssl session to be cached per dynamic origin server. A unique ssl session is created for each SNI received from the client on ClientHello and the matching session is used for server session reuse.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable transparent interception of SSL sessions.",
			},
			"ssliocspcheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable OCSP check for origin server certificate.",
			},
			"sslireneg": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server.",
			},
			"ssllogprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the ssllogprofile.",
			},
			"sslprofiletype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("FrontEnd"),
				Description: "Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.",
			},
			"sslredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of HTTPS redirects for the SSL service.\nFor an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.\nIf SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.\nThis parameter is not applicable when configuring a backend profile.",
			},
			"ssltriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.",
			},
			"strictcachecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable strict CA certificate checks on the appliance.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.0 protocol support for the SSL profile.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.1 protocol support for the SSL profile.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.2 protocol support for the SSL profile.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLSv1.3 protocol support for the SSL profile.",
			},
			"tls13sessionticketsperauthcontext": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes.\nThis value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection.\nNo tickets are sent if resumption is disabled.",
			},
			"zerorttearlydata": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake.\nEarly application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.",
			},
		},
	}
}

func sslprofileGetThePayloadFromtheConfig(ctx context.Context, data *SslprofileResourceModel) ssl.Sslprofile {
	tflog.Debug(ctx, "In sslprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslprofile := ssl.Sslprofile{}
	if !data.Allowextendedmastersecret.IsNull() {
		sslprofile.Allowextendedmastersecret = data.Allowextendedmastersecret.ValueString()
	}
	if !data.Allowunknownsni.IsNull() {
		sslprofile.Allowunknownsni = data.Allowunknownsni.ValueString()
	}
	if !data.Alpnprotocol.IsNull() {
		sslprofile.Alpnprotocol = data.Alpnprotocol.ValueString()
	}
	if !data.Ciphername.IsNull() {
		sslprofile.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipherpriority.IsNull() {
		sslprofile.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Cipherredirect.IsNull() {
		sslprofile.Cipherredirect = data.Cipherredirect.ValueString()
	}
	if !data.Cipherurl.IsNull() {
		sslprofile.Cipherurl = data.Cipherurl.ValueString()
	}
	if !data.Cleartextport.IsNull() {
		sslprofile.Cleartextport = utils.IntPtr(int(data.Cleartextport.ValueInt64()))
	}
	if !data.Clientauth.IsNull() {
		sslprofile.Clientauth = data.Clientauth.ValueString()
	}
	if !data.Clientauthuseboundcachain.IsNull() {
		sslprofile.Clientauthuseboundcachain = data.Clientauthuseboundcachain.ValueString()
	}
	if !data.Clientcert.IsNull() {
		sslprofile.Clientcert = data.Clientcert.ValueString()
	}
	if !data.Commonname.IsNull() {
		sslprofile.Commonname = data.Commonname.ValueString()
	}
	if !data.Defaultsni.IsNull() {
		sslprofile.Defaultsni = data.Defaultsni.ValueString()
	}
	if !data.Denysslreneg.IsNull() {
		sslprofile.Denysslreneg = data.Denysslreneg.ValueString()
	}
	if !data.Dh.IsNull() {
		sslprofile.Dh = data.Dh.ValueString()
	}
	if !data.Dhcount.IsNull() {
		sslprofile.Dhcount = utils.IntPtr(int(data.Dhcount.ValueInt64()))
	}
	if !data.Dhekeyexchangewithpsk.IsNull() {
		sslprofile.Dhekeyexchangewithpsk = data.Dhekeyexchangewithpsk.ValueString()
	}
	if !data.Dhfile.IsNull() {
		sslprofile.Dhfile = data.Dhfile.ValueString()
	}
	if !data.Dhkeyexpsizelimit.IsNull() {
		sslprofile.Dhkeyexpsizelimit = data.Dhkeyexpsizelimit.ValueString()
	}
	if !data.Dropreqwithnohostheader.IsNull() {
		sslprofile.Dropreqwithnohostheader = data.Dropreqwithnohostheader.ValueString()
	}
	if !data.Encryptedclienthello.IsNull() {
		sslprofile.Encryptedclienthello = data.Encryptedclienthello.ValueString()
	}
	if !data.Encrypttriggerpktcount.IsNull() {
		sslprofile.Encrypttriggerpktcount = utils.IntPtr(int(data.Encrypttriggerpktcount.ValueInt64()))
	}
	if !data.Ersa.IsNull() {
		sslprofile.Ersa = data.Ersa.ValueString()
	}
	if !data.Ersacount.IsNull() {
		sslprofile.Ersacount = utils.IntPtr(int(data.Ersacount.ValueInt64()))
	}
	if !data.Hsts.IsNull() {
		sslprofile.Hsts = data.Hsts.ValueString()
	}
	if !data.Includesubdomains.IsNull() {
		sslprofile.Includesubdomains = data.Includesubdomains.ValueString()
	}
	if !data.Insertionencoding.IsNull() {
		sslprofile.Insertionencoding = data.Insertionencoding.ValueString()
	}
	if !data.Maxage.IsNull() {
		sslprofile.Maxage = utils.IntPtr(int(data.Maxage.ValueInt64()))
	}
	if !data.Maxrenegrate.IsNull() {
		sslprofile.Maxrenegrate = utils.IntPtr(int(data.Maxrenegrate.ValueInt64()))
	}
	if !data.Name.IsNull() {
		sslprofile.Name = data.Name.ValueString()
	}
	if !data.Ocspstapling.IsNull() {
		sslprofile.Ocspstapling = data.Ocspstapling.ValueString()
	}
	if !data.Preload.IsNull() {
		sslprofile.Preload = data.Preload.ValueString()
	}
	if !data.Prevsessionkeylifetime.IsNull() {
		sslprofile.Prevsessionkeylifetime = utils.IntPtr(int(data.Prevsessionkeylifetime.ValueInt64()))
	}
	if !data.Pushenctrigger.IsNull() {
		sslprofile.Pushenctrigger = data.Pushenctrigger.ValueString()
	}
	if !data.Pushenctriggertimeout.IsNull() {
		sslprofile.Pushenctriggertimeout = utils.IntPtr(int(data.Pushenctriggertimeout.ValueInt64()))
	}
	if !data.Pushflag.IsNull() {
		sslprofile.Pushflag = utils.IntPtr(int(data.Pushflag.ValueInt64()))
	}
	if !data.Quantumsize.IsNull() {
		sslprofile.Quantumsize = data.Quantumsize.ValueString()
	}
	if !data.Redirectportrewrite.IsNull() {
		sslprofile.Redirectportrewrite = data.Redirectportrewrite.ValueString()
	}
	if !data.Sendclosenotify.IsNull() {
		sslprofile.Sendclosenotify = data.Sendclosenotify.ValueString()
	}
	if !data.Serverauth.IsNull() {
		sslprofile.Serverauth = data.Serverauth.ValueString()
	}
	if !data.Sessionkeylifetime.IsNull() {
		sslprofile.Sessionkeylifetime = utils.IntPtr(int(data.Sessionkeylifetime.ValueInt64()))
	}
	if !data.Sessionticket.IsNull() {
		sslprofile.Sessionticket = data.Sessionticket.ValueString()
	}
	if !data.Sessionticketkeydata.IsNull() {
		sslprofile.Sessionticketkeydata = data.Sessionticketkeydata.ValueString()
	}
	if !data.Sessionticketkeyrefresh.IsNull() {
		sslprofile.Sessionticketkeyrefresh = data.Sessionticketkeyrefresh.ValueString()
	}
	if !data.Sessionticketlifetime.IsNull() {
		sslprofile.Sessionticketlifetime = utils.IntPtr(int(data.Sessionticketlifetime.ValueInt64()))
	}
	if !data.Sessreuse.IsNull() {
		sslprofile.Sessreuse = data.Sessreuse.ValueString()
	}
	if !data.Sesstimeout.IsNull() {
		sslprofile.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Skipclientcertpolicycheck.IsNull() {
		sslprofile.Skipclientcertpolicycheck = data.Skipclientcertpolicycheck.ValueString()
	}
	if !data.Snienable.IsNull() {
		sslprofile.Snienable = data.Snienable.ValueString()
	}
	if !data.Snihttphostmatch.IsNull() {
		sslprofile.Snihttphostmatch = data.Snihttphostmatch.ValueString()
	}
	if !data.Ssl3.IsNull() {
		sslprofile.Ssl3 = data.Ssl3.ValueString()
	}
	if !data.Sslclientlogs.IsNull() {
		sslprofile.Sslclientlogs = data.Sslclientlogs.ValueString()
	}
	if !data.Sslimaxsessperserver.IsNull() {
		sslprofile.Sslimaxsessperserver = utils.IntPtr(int(data.Sslimaxsessperserver.ValueInt64()))
	}
	if !data.Sslinterception.IsNull() {
		sslprofile.Sslinterception = data.Sslinterception.ValueString()
	}
	if !data.Ssliocspcheck.IsNull() {
		sslprofile.Ssliocspcheck = data.Ssliocspcheck.ValueString()
	}
	if !data.Sslireneg.IsNull() {
		sslprofile.Sslireneg = data.Sslireneg.ValueString()
	}
	if !data.Ssllogprofile.IsNull() {
		sslprofile.Ssllogprofile = data.Ssllogprofile.ValueString()
	}
	if !data.Sslprofiletype.IsNull() {
		sslprofile.Sslprofiletype = data.Sslprofiletype.ValueString()
	}
	if !data.Sslredirect.IsNull() {
		sslprofile.Sslredirect = data.Sslredirect.ValueString()
	}
	if !data.Ssltriggertimeout.IsNull() {
		sslprofile.Ssltriggertimeout = utils.IntPtr(int(data.Ssltriggertimeout.ValueInt64()))
	}
	if !data.Strictcachecks.IsNull() {
		sslprofile.Strictcachecks = data.Strictcachecks.ValueString()
	}
	if !data.Strictsigdigestcheck.IsNull() {
		sslprofile.Strictsigdigestcheck = data.Strictsigdigestcheck.ValueString()
	}
	if !data.Tls1.IsNull() {
		sslprofile.Tls1 = data.Tls1.ValueString()
	}
	if !data.Tls11.IsNull() {
		sslprofile.Tls11 = data.Tls11.ValueString()
	}
	if !data.Tls12.IsNull() {
		sslprofile.Tls12 = data.Tls12.ValueString()
	}
	if !data.Tls13.IsNull() {
		sslprofile.Tls13 = data.Tls13.ValueString()
	}
	if !data.Tls13sessionticketsperauthcontext.IsNull() {
		sslprofile.Tls13sessionticketsperauthcontext = utils.IntPtr(int(data.Tls13sessionticketsperauthcontext.ValueInt64()))
	}
	if !data.Zerorttearlydata.IsNull() {
		sslprofile.Zerorttearlydata = data.Zerorttearlydata.ValueString()
	}

	return sslprofile
}

func sslprofileSetAttrFromGet(ctx context.Context, data *SslprofileResourceModel, getResponseData map[string]interface{}) *SslprofileResourceModel {
	tflog.Debug(ctx, "In sslprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allowextendedmastersecret"]; ok && val != nil {
		data.Allowextendedmastersecret = types.StringValue(val.(string))
	} else {
		data.Allowextendedmastersecret = types.StringNull()
	}
	if val, ok := getResponseData["allowunknownsni"]; ok && val != nil {
		data.Allowunknownsni = types.StringValue(val.(string))
	} else {
		data.Allowunknownsni = types.StringNull()
	}
	if val, ok := getResponseData["alpnprotocol"]; ok && val != nil {
		data.Alpnprotocol = types.StringValue(val.(string))
	} else {
		data.Alpnprotocol = types.StringNull()
	}
	if val, ok := getResponseData["ciphername"]; ok && val != nil {
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Ciphername = types.StringNull()
	}
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["cipherredirect"]; ok && val != nil {
		data.Cipherredirect = types.StringValue(val.(string))
	} else {
		data.Cipherredirect = types.StringNull()
	}
	if val, ok := getResponseData["cipherurl"]; ok && val != nil {
		data.Cipherurl = types.StringValue(val.(string))
	} else {
		data.Cipherurl = types.StringNull()
	}
	if val, ok := getResponseData["cleartextport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cleartextport = types.Int64Value(intVal)
		}
	} else {
		data.Cleartextport = types.Int64Null()
	}
	if val, ok := getResponseData["clientauth"]; ok && val != nil {
		data.Clientauth = types.StringValue(val.(string))
	} else {
		data.Clientauth = types.StringNull()
	}
	if val, ok := getResponseData["clientauthuseboundcachain"]; ok && val != nil {
		data.Clientauthuseboundcachain = types.StringValue(val.(string))
	} else {
		data.Clientauthuseboundcachain = types.StringNull()
	}
	if val, ok := getResponseData["clientcert"]; ok && val != nil {
		data.Clientcert = types.StringValue(val.(string))
	} else {
		data.Clientcert = types.StringNull()
	}
	if val, ok := getResponseData["commonname"]; ok && val != nil {
		data.Commonname = types.StringValue(val.(string))
	} else {
		data.Commonname = types.StringNull()
	}
	if val, ok := getResponseData["defaultsni"]; ok && val != nil {
		data.Defaultsni = types.StringValue(val.(string))
	} else {
		data.Defaultsni = types.StringNull()
	}
	if val, ok := getResponseData["denysslreneg"]; ok && val != nil {
		data.Denysslreneg = types.StringValue(val.(string))
	} else {
		data.Denysslreneg = types.StringNull()
	}
	if val, ok := getResponseData["dh"]; ok && val != nil {
		data.Dh = types.StringValue(val.(string))
	} else {
		data.Dh = types.StringNull()
	}
	if val, ok := getResponseData["dhcount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dhcount = types.Int64Value(intVal)
		}
	} else {
		data.Dhcount = types.Int64Null()
	}
	if val, ok := getResponseData["dhekeyexchangewithpsk"]; ok && val != nil {
		data.Dhekeyexchangewithpsk = types.StringValue(val.(string))
	} else {
		data.Dhekeyexchangewithpsk = types.StringNull()
	}
	if val, ok := getResponseData["dhfile"]; ok && val != nil {
		data.Dhfile = types.StringValue(val.(string))
	} else {
		data.Dhfile = types.StringNull()
	}
	if val, ok := getResponseData["dhkeyexpsizelimit"]; ok && val != nil {
		data.Dhkeyexpsizelimit = types.StringValue(val.(string))
	} else {
		data.Dhkeyexpsizelimit = types.StringNull()
	}
	if val, ok := getResponseData["dropreqwithnohostheader"]; ok && val != nil {
		data.Dropreqwithnohostheader = types.StringValue(val.(string))
	} else {
		data.Dropreqwithnohostheader = types.StringNull()
	}
	if val, ok := getResponseData["encryptedclienthello"]; ok && val != nil {
		data.Encryptedclienthello = types.StringValue(val.(string))
	} else {
		data.Encryptedclienthello = types.StringNull()
	}
	if val, ok := getResponseData["encrypttriggerpktcount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Encrypttriggerpktcount = types.Int64Value(intVal)
		}
	} else {
		data.Encrypttriggerpktcount = types.Int64Null()
	}
	if val, ok := getResponseData["ersa"]; ok && val != nil {
		data.Ersa = types.StringValue(val.(string))
	} else {
		data.Ersa = types.StringNull()
	}
	if val, ok := getResponseData["ersacount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ersacount = types.Int64Value(intVal)
		}
	} else {
		data.Ersacount = types.Int64Null()
	}
	if val, ok := getResponseData["hsts"]; ok && val != nil {
		data.Hsts = types.StringValue(val.(string))
	} else {
		data.Hsts = types.StringNull()
	}
	if val, ok := getResponseData["includesubdomains"]; ok && val != nil {
		data.Includesubdomains = types.StringValue(val.(string))
	} else {
		data.Includesubdomains = types.StringNull()
	}
	if val, ok := getResponseData["insertionencoding"]; ok && val != nil {
		data.Insertionencoding = types.StringValue(val.(string))
	} else {
		data.Insertionencoding = types.StringNull()
	}
	if val, ok := getResponseData["maxage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxage = types.Int64Value(intVal)
		}
	} else {
		data.Maxage = types.Int64Null()
	}
	if val, ok := getResponseData["maxrenegrate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxrenegrate = types.Int64Value(intVal)
		}
	} else {
		data.Maxrenegrate = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ocspstapling"]; ok && val != nil {
		data.Ocspstapling = types.StringValue(val.(string))
	} else {
		data.Ocspstapling = types.StringNull()
	}
	if val, ok := getResponseData["preload"]; ok && val != nil {
		data.Preload = types.StringValue(val.(string))
	} else {
		data.Preload = types.StringNull()
	}
	if val, ok := getResponseData["prevsessionkeylifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prevsessionkeylifetime = types.Int64Value(intVal)
		}
	} else {
		data.Prevsessionkeylifetime = types.Int64Null()
	}
	if val, ok := getResponseData["pushenctrigger"]; ok && val != nil {
		data.Pushenctrigger = types.StringValue(val.(string))
	} else {
		data.Pushenctrigger = types.StringNull()
	}
	if val, ok := getResponseData["pushenctriggertimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pushenctriggertimeout = types.Int64Value(intVal)
		}
	} else {
		data.Pushenctriggertimeout = types.Int64Null()
	}
	if val, ok := getResponseData["pushflag"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pushflag = types.Int64Value(intVal)
		}
	} else {
		data.Pushflag = types.Int64Null()
	}
	if val, ok := getResponseData["quantumsize"]; ok && val != nil {
		data.Quantumsize = types.StringValue(val.(string))
	} else {
		data.Quantumsize = types.StringNull()
	}
	if val, ok := getResponseData["redirectportrewrite"]; ok && val != nil {
		data.Redirectportrewrite = types.StringValue(val.(string))
	} else {
		data.Redirectportrewrite = types.StringNull()
	}
	if val, ok := getResponseData["sendclosenotify"]; ok && val != nil {
		data.Sendclosenotify = types.StringValue(val.(string))
	} else {
		data.Sendclosenotify = types.StringNull()
	}
	if val, ok := getResponseData["serverauth"]; ok && val != nil {
		data.Serverauth = types.StringValue(val.(string))
	} else {
		data.Serverauth = types.StringNull()
	}
	if val, ok := getResponseData["sessionkeylifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionkeylifetime = types.Int64Value(intVal)
		}
	} else {
		data.Sessionkeylifetime = types.Int64Null()
	}
	if val, ok := getResponseData["sessionticket"]; ok && val != nil {
		data.Sessionticket = types.StringValue(val.(string))
	} else {
		data.Sessionticket = types.StringNull()
	}
	if val, ok := getResponseData["sessionticketkeydata"]; ok && val != nil {
		data.Sessionticketkeydata = types.StringValue(val.(string))
	} else {
		data.Sessionticketkeydata = types.StringNull()
	}
	if val, ok := getResponseData["sessionticketkeyrefresh"]; ok && val != nil {
		data.Sessionticketkeyrefresh = types.StringValue(val.(string))
	} else {
		data.Sessionticketkeyrefresh = types.StringNull()
	}
	if val, ok := getResponseData["sessionticketlifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionticketlifetime = types.Int64Value(intVal)
		}
	} else {
		data.Sessionticketlifetime = types.Int64Null()
	}
	if val, ok := getResponseData["sessreuse"]; ok && val != nil {
		data.Sessreuse = types.StringValue(val.(string))
	} else {
		data.Sessreuse = types.StringNull()
	}
	if val, ok := getResponseData["sesstimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sesstimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sesstimeout = types.Int64Null()
	}
	if val, ok := getResponseData["skipclientcertpolicycheck"]; ok && val != nil {
		data.Skipclientcertpolicycheck = types.StringValue(val.(string))
	} else {
		data.Skipclientcertpolicycheck = types.StringNull()
	}
	if val, ok := getResponseData["snienable"]; ok && val != nil {
		data.Snienable = types.StringValue(val.(string))
	} else {
		data.Snienable = types.StringNull()
	}
	if val, ok := getResponseData["snihttphostmatch"]; ok && val != nil {
		data.Snihttphostmatch = types.StringValue(val.(string))
	} else {
		data.Snihttphostmatch = types.StringNull()
	}
	if val, ok := getResponseData["ssl3"]; ok && val != nil {
		data.Ssl3 = types.StringValue(val.(string))
	} else {
		data.Ssl3 = types.StringNull()
	}
	if val, ok := getResponseData["sslclientlogs"]; ok && val != nil {
		data.Sslclientlogs = types.StringValue(val.(string))
	} else {
		data.Sslclientlogs = types.StringNull()
	}
	if val, ok := getResponseData["sslimaxsessperserver"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sslimaxsessperserver = types.Int64Value(intVal)
		}
	} else {
		data.Sslimaxsessperserver = types.Int64Null()
	}
	if val, ok := getResponseData["sslinterception"]; ok && val != nil {
		data.Sslinterception = types.StringValue(val.(string))
	} else {
		data.Sslinterception = types.StringNull()
	}
	if val, ok := getResponseData["ssliocspcheck"]; ok && val != nil {
		data.Ssliocspcheck = types.StringValue(val.(string))
	} else {
		data.Ssliocspcheck = types.StringNull()
	}
	if val, ok := getResponseData["sslireneg"]; ok && val != nil {
		data.Sslireneg = types.StringValue(val.(string))
	} else {
		data.Sslireneg = types.StringNull()
	}
	if val, ok := getResponseData["ssllogprofile"]; ok && val != nil {
		data.Ssllogprofile = types.StringValue(val.(string))
	} else {
		data.Ssllogprofile = types.StringNull()
	}
	if val, ok := getResponseData["sslprofiletype"]; ok && val != nil {
		data.Sslprofiletype = types.StringValue(val.(string))
	} else {
		data.Sslprofiletype = types.StringNull()
	}
	if val, ok := getResponseData["sslredirect"]; ok && val != nil {
		data.Sslredirect = types.StringValue(val.(string))
	} else {
		data.Sslredirect = types.StringNull()
	}
	if val, ok := getResponseData["ssltriggertimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ssltriggertimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ssltriggertimeout = types.Int64Null()
	}
	if val, ok := getResponseData["strictcachecks"]; ok && val != nil {
		data.Strictcachecks = types.StringValue(val.(string))
	} else {
		data.Strictcachecks = types.StringNull()
	}
	if val, ok := getResponseData["strictsigdigestcheck"]; ok && val != nil {
		data.Strictsigdigestcheck = types.StringValue(val.(string))
	} else {
		data.Strictsigdigestcheck = types.StringNull()
	}
	if val, ok := getResponseData["tls1"]; ok && val != nil {
		data.Tls1 = types.StringValue(val.(string))
	} else {
		data.Tls1 = types.StringNull()
	}
	if val, ok := getResponseData["tls11"]; ok && val != nil {
		data.Tls11 = types.StringValue(val.(string))
	} else {
		data.Tls11 = types.StringNull()
	}
	if val, ok := getResponseData["tls12"]; ok && val != nil {
		data.Tls12 = types.StringValue(val.(string))
	} else {
		data.Tls12 = types.StringNull()
	}
	if val, ok := getResponseData["tls13"]; ok && val != nil {
		data.Tls13 = types.StringValue(val.(string))
	} else {
		data.Tls13 = types.StringNull()
	}
	if val, ok := getResponseData["tls13sessionticketsperauthcontext"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tls13sessionticketsperauthcontext = types.Int64Value(intVal)
		}
	} else {
		data.Tls13sessionticketsperauthcontext = types.Int64Null()
	}
	if val, ok := getResponseData["zerorttearlydata"]; ok && val != nil {
		data.Zerorttearlydata = types.StringValue(val.(string))
	} else {
		data.Zerorttearlydata = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
