package sslvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslvserverResourceModel describes the resource data model.
type SslvserverResourceModel struct {
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
	Strictsigdigestcheck              types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                              types.String `tfsdk:"tls1"`
	Tls11                             types.String `tfsdk:"tls11"`
	Tls12                             types.String `tfsdk:"tls12"`
	Tls13                             types.String `tfsdk:"tls13"`
	Tls13sessionticketsperauthcontext types.Int64  `tfsdk:"tls13sessionticketsperauthcontext"`
	Vservername                       types.String `tfsdk:"vservername"`
	Zerorttearlydata                  types.String `tfsdk:"zerorttearlydata"`
}

func (r *SslvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver resource.",
			},
			"cipherredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.",
			},
			"dtls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of DTLSv1.0 protocol support for the SSL Virtual Server.",
			},
			"dtls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of DTLSv1.2 protocol support for the SSL Virtual Server.",
			},
			"dtlsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DTLS profile whose settings are to be applied to the virtual server.",
			},
			"ersa": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.",
			},
			"ersacount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Refresh count for regeneration of the RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).",
			},
			"hsts": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of the port rewrite while performing HTTPS redirect. If this parameter is ENABLED and the URL from the server does not contain the standard port, the port is rewritten to the standard.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"ssl2": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of SSLv2 protocol support for the SSL Virtual Server.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of SSLv3 protocol support for the SSL Virtual Server.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the virtual server.",
			},
			"sslredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of HTTPS redirects for the SSL virtual server.\n\nFor an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.\nIf SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.",
			},
			"sslv2redirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of SSLv2 Redirect. If SSLv2 redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client.",
			},
			"sslv2url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.0 protocol support for the SSL Virtual Server.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.1 protocol support for the SSL Virtual Server.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.2 protocol support for the SSL Virtual Server.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLSv1.3 protocol support for the SSL Virtual Server.",
			},
			"tls13sessionticketsperauthcontext": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes.\nThis value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection.\nNo tickets are sent if resumption is disabled.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server for which to set advanced configuration.",
			},
			"zerorttearlydata": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake.\nEarly application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.",
			},
		},
	}
}

func sslvserverGetThePayloadFromtheConfig(ctx context.Context, data *SslvserverResourceModel) ssl.Sslvserver {
	tflog.Debug(ctx, "In sslvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslvserver := ssl.Sslvserver{}
	if !data.Cipherredirect.IsNull() {
		sslvserver.Cipherredirect = data.Cipherredirect.ValueString()
	}
	if !data.Cipherurl.IsNull() {
		sslvserver.Cipherurl = data.Cipherurl.ValueString()
	}
	if !data.Cleartextport.IsNull() {
		sslvserver.Cleartextport = utils.IntPtr(int(data.Cleartextport.ValueInt64()))
	}
	if !data.Clientauth.IsNull() {
		sslvserver.Clientauth = data.Clientauth.ValueString()
	}
	if !data.Clientcert.IsNull() {
		sslvserver.Clientcert = data.Clientcert.ValueString()
	}
	if !data.Defaultsni.IsNull() {
		sslvserver.Defaultsni = data.Defaultsni.ValueString()
	}
	if !data.Dh.IsNull() {
		sslvserver.Dh = data.Dh.ValueString()
	}
	if !data.Dhcount.IsNull() {
		sslvserver.Dhcount = utils.IntPtr(int(data.Dhcount.ValueInt64()))
	}
	if !data.Dhekeyexchangewithpsk.IsNull() {
		sslvserver.Dhekeyexchangewithpsk = data.Dhekeyexchangewithpsk.ValueString()
	}
	if !data.Dhfile.IsNull() {
		sslvserver.Dhfile = data.Dhfile.ValueString()
	}
	if !data.Dhkeyexpsizelimit.IsNull() {
		sslvserver.Dhkeyexpsizelimit = data.Dhkeyexpsizelimit.ValueString()
	}
	if !data.Dtls1.IsNull() {
		sslvserver.Dtls1 = data.Dtls1.ValueString()
	}
	if !data.Dtls12.IsNull() {
		sslvserver.Dtls12 = data.Dtls12.ValueString()
	}
	if !data.Dtlsprofilename.IsNull() {
		sslvserver.Dtlsprofilename = data.Dtlsprofilename.ValueString()
	}
	if !data.Ersa.IsNull() {
		sslvserver.Ersa = data.Ersa.ValueString()
	}
	if !data.Ersacount.IsNull() {
		sslvserver.Ersacount = utils.IntPtr(int(data.Ersacount.ValueInt64()))
	}
	if !data.Hsts.IsNull() {
		sslvserver.Hsts = data.Hsts.ValueString()
	}
	if !data.Includesubdomains.IsNull() {
		sslvserver.Includesubdomains = data.Includesubdomains.ValueString()
	}
	if !data.Maxage.IsNull() {
		sslvserver.Maxage = utils.IntPtr(int(data.Maxage.ValueInt64()))
	}
	if !data.Ocspstapling.IsNull() {
		sslvserver.Ocspstapling = data.Ocspstapling.ValueString()
	}
	if !data.Preload.IsNull() {
		sslvserver.Preload = data.Preload.ValueString()
	}
	if !data.Pushenctrigger.IsNull() {
		sslvserver.Pushenctrigger = data.Pushenctrigger.ValueString()
	}
	if !data.Redirectportrewrite.IsNull() {
		sslvserver.Redirectportrewrite = data.Redirectportrewrite.ValueString()
	}
	if !data.Sendclosenotify.IsNull() {
		sslvserver.Sendclosenotify = data.Sendclosenotify.ValueString()
	}
	if !data.Sessreuse.IsNull() {
		sslvserver.Sessreuse = data.Sessreuse.ValueString()
	}
	if !data.Sesstimeout.IsNull() {
		sslvserver.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Snienable.IsNull() {
		sslvserver.Snienable = data.Snienable.ValueString()
	}
	if !data.Ssl2.IsNull() {
		sslvserver.Ssl2 = data.Ssl2.ValueString()
	}
	if !data.Ssl3.IsNull() {
		sslvserver.Ssl3 = data.Ssl3.ValueString()
	}
	if !data.Sslclientlogs.IsNull() {
		sslvserver.Sslclientlogs = data.Sslclientlogs.ValueString()
	}
	if !data.Sslprofile.IsNull() {
		sslvserver.Sslprofile = data.Sslprofile.ValueString()
	}
	if !data.Sslredirect.IsNull() {
		sslvserver.Sslredirect = data.Sslredirect.ValueString()
	}
	if !data.Sslv2redirect.IsNull() {
		sslvserver.Sslv2redirect = data.Sslv2redirect.ValueString()
	}
	if !data.Sslv2url.IsNull() {
		sslvserver.Sslv2url = data.Sslv2url.ValueString()
	}
	if !data.Strictsigdigestcheck.IsNull() {
		sslvserver.Strictsigdigestcheck = data.Strictsigdigestcheck.ValueString()
	}
	if !data.Tls1.IsNull() {
		sslvserver.Tls1 = data.Tls1.ValueString()
	}
	if !data.Tls11.IsNull() {
		sslvserver.Tls11 = data.Tls11.ValueString()
	}
	if !data.Tls12.IsNull() {
		sslvserver.Tls12 = data.Tls12.ValueString()
	}
	if !data.Tls13.IsNull() {
		sslvserver.Tls13 = data.Tls13.ValueString()
	}
	if !data.Tls13sessionticketsperauthcontext.IsNull() {
		sslvserver.Tls13sessionticketsperauthcontext = utils.IntPtr(int(data.Tls13sessionticketsperauthcontext.ValueInt64()))
	}
	if !data.Vservername.IsNull() {
		sslvserver.Vservername = data.Vservername.ValueString()
	}
	if !data.Zerorttearlydata.IsNull() {
		sslvserver.Zerorttearlydata = data.Zerorttearlydata.ValueString()
	}

	return sslvserver
}

func sslvserverSetAttrFromGet(ctx context.Context, data *SslvserverResourceModel, getResponseData map[string]interface{}) *SslvserverResourceModel {
	tflog.Debug(ctx, "In sslvserverSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["clientcert"]; ok && val != nil {
		data.Clientcert = types.StringValue(val.(string))
	} else {
		data.Clientcert = types.StringNull()
	}
	if val, ok := getResponseData["defaultsni"]; ok && val != nil {
		data.Defaultsni = types.StringValue(val.(string))
	} else {
		data.Defaultsni = types.StringNull()
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
	if val, ok := getResponseData["dtls1"]; ok && val != nil {
		data.Dtls1 = types.StringValue(val.(string))
	} else {
		data.Dtls1 = types.StringNull()
	}
	if val, ok := getResponseData["dtls12"]; ok && val != nil {
		data.Dtls12 = types.StringValue(val.(string))
	} else {
		data.Dtls12 = types.StringNull()
	}
	if val, ok := getResponseData["dtlsprofilename"]; ok && val != nil {
		data.Dtlsprofilename = types.StringValue(val.(string))
	} else {
		data.Dtlsprofilename = types.StringNull()
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
	if val, ok := getResponseData["maxage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxage = types.Int64Value(intVal)
		}
	} else {
		data.Maxage = types.Int64Null()
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
	if val, ok := getResponseData["pushenctrigger"]; ok && val != nil {
		data.Pushenctrigger = types.StringValue(val.(string))
	} else {
		data.Pushenctrigger = types.StringNull()
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
	if val, ok := getResponseData["snienable"]; ok && val != nil {
		data.Snienable = types.StringValue(val.(string))
	} else {
		data.Snienable = types.StringNull()
	}
	if val, ok := getResponseData["ssl2"]; ok && val != nil {
		data.Ssl2 = types.StringValue(val.(string))
	} else {
		data.Ssl2 = types.StringNull()
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
	if val, ok := getResponseData["sslprofile"]; ok && val != nil {
		data.Sslprofile = types.StringValue(val.(string))
	} else {
		data.Sslprofile = types.StringNull()
	}
	if val, ok := getResponseData["sslredirect"]; ok && val != nil {
		data.Sslredirect = types.StringValue(val.(string))
	} else {
		data.Sslredirect = types.StringNull()
	}
	if val, ok := getResponseData["sslv2redirect"]; ok && val != nil {
		data.Sslv2redirect = types.StringValue(val.(string))
	} else {
		data.Sslv2redirect = types.StringNull()
	}
	if val, ok := getResponseData["sslv2url"]; ok && val != nil {
		data.Sslv2url = types.StringValue(val.(string))
	} else {
		data.Sslv2url = types.StringNull()
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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}
	if val, ok := getResponseData["zerorttearlydata"]; ok && val != nil {
		data.Zerorttearlydata = types.StringValue(val.(string))
	} else {
		data.Zerorttearlydata = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Vservername.ValueString())

	return data
}
