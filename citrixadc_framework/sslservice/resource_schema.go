package sslservice

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

// SslserviceResourceModel describes the resource data model.
type SslserviceResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Cipherredirect       types.String `tfsdk:"cipherredirect"`
	Cipherurl            types.String `tfsdk:"cipherurl"`
	Clientauth           types.String `tfsdk:"clientauth"`
	Clientcert           types.String `tfsdk:"clientcert"`
	Commonname           types.String `tfsdk:"commonname"`
	Dh                   types.String `tfsdk:"dh"`
	Dhcount              types.Int64  `tfsdk:"dhcount"`
	Dhfile               types.String `tfsdk:"dhfile"`
	Dhkeyexpsizelimit    types.String `tfsdk:"dhkeyexpsizelimit"`
	Dtls1                types.String `tfsdk:"dtls1"`
	Dtls12               types.String `tfsdk:"dtls12"`
	Dtlsprofilename      types.String `tfsdk:"dtlsprofilename"`
	Ersa                 types.String `tfsdk:"ersa"`
	Ersacount            types.Int64  `tfsdk:"ersacount"`
	Ocspstapling         types.String `tfsdk:"ocspstapling"`
	Pushenctrigger       types.String `tfsdk:"pushenctrigger"`
	Redirectportrewrite  types.String `tfsdk:"redirectportrewrite"`
	Sendclosenotify      types.String `tfsdk:"sendclosenotify"`
	Serverauth           types.String `tfsdk:"serverauth"`
	Servicename          types.String `tfsdk:"servicename"`
	Sessreuse            types.String `tfsdk:"sessreuse"`
	Sesstimeout          types.Int64  `tfsdk:"sesstimeout"`
	Snienable            types.String `tfsdk:"snienable"`
	Ssl2                 types.String `tfsdk:"ssl2"`
	Ssl3                 types.String `tfsdk:"ssl3"`
	Sslclientlogs        types.String `tfsdk:"sslclientlogs"`
	Sslprofile           types.String `tfsdk:"sslprofile"`
	Sslredirect          types.String `tfsdk:"sslredirect"`
	Sslv2redirect        types.String `tfsdk:"sslv2redirect"`
	Sslv2url             types.String `tfsdk:"sslv2url"`
	Strictsigdigestcheck types.String `tfsdk:"strictsigdigestcheck"`
	Tls1                 types.String `tfsdk:"tls1"`
	Tls11                types.String `tfsdk:"tls11"`
	Tls12                types.String `tfsdk:"tls12"`
	Tls13                types.String `tfsdk:"tls13"`
}

func (r *SslserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice resource.",
			},
			"cipherredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.\nThis parameter is not applicable when configuring a backend service.",
			},
			"cipherurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the page to which to redirect the client in case of a cipher mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.\nThis parameter is not applicable when configuring a backend service.",
			},
			"clientauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate.\nThis parameter is not applicable when configuring a backend service.",
			},
			"clientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate.\nThis parameter is not applicable when configuring a backend SSL service.\nCaution: Define proper access control policies before changing this setting to Optional.",
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server",
			},
			"dh": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend service.",
			},
			"dhcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time. This parameter is not applicable when configuring a backend service. Allowed DH count values are 0 and >= 500.",
			},
			"dhfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for and, optionally, path to the PEM-format DH parameter file to be installed. /nsconfig/ssl/ is the default path. This parameter is not applicable when configuring a backend service.",
			},
			"dhkeyexpsizelimit": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.",
			},
			"dtls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of DTLSv1.0 protocol support for the SSL service.",
			},
			"dtls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of DTLSv1.2 protocol support for the SSL service.",
			},
			"dtlsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DTLS profile that contains DTLS settings for the service.",
			},
			"ersa": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.\nThis parameter is not applicable when configuring a backend service.",
			},
			"ersacount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Refresh count for regeneration of RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).\nThis parameter is not applicable when configuring a backend service.",
			},
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:\nENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.\nDISABLED: The appliance does not check the status of the server certificate.",
			},
			"pushenctrigger": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:\n* ALWAYS - Any PUSH packet triggers encryption.\n* IGNORE - Ignore PUSH packet for triggering encryption.\n* MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption.\n* TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of server authentication support for the SSL service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service.",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
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
				Description: "State of SSLv2 protocol support for the SSL service.\nThis parameter is not applicable when configuring a backend service.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of SSLv3 protocol support for the SSL service.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI name, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the service.",
			},
			"sslredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of HTTPS redirects for the SSL service.\n\nFor an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.\nIf SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.\n\nThis parameter is not applicable when configuring a backend service.",
			},
			"sslv2redirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of SSLv2 Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client.\nThis parameter is not applicable when configuring a backend service.",
			},
			"sslv2url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.\nThis parameter is not applicable when configuring a backend service.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter indicating to check whether peer's certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.0 protocol support for the SSL service.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.1 protocol support for the SSL service.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.2 protocol support for the SSL service.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLSv1.3 protocol support for the SSL service.",
			},
		},
	}
}

func sslserviceGetThePayloadFromtheConfig(ctx context.Context, data *SslserviceResourceModel) ssl.Sslservice {
	tflog.Debug(ctx, "In sslserviceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservice := ssl.Sslservice{}
	if !data.Cipherredirect.IsNull() {
		sslservice.Cipherredirect = data.Cipherredirect.ValueString()
	}
	if !data.Cipherurl.IsNull() {
		sslservice.Cipherurl = data.Cipherurl.ValueString()
	}
	if !data.Clientauth.IsNull() {
		sslservice.Clientauth = data.Clientauth.ValueString()
	}
	if !data.Clientcert.IsNull() {
		sslservice.Clientcert = data.Clientcert.ValueString()
	}
	if !data.Commonname.IsNull() {
		sslservice.Commonname = data.Commonname.ValueString()
	}
	if !data.Dh.IsNull() {
		sslservice.Dh = data.Dh.ValueString()
	}
	if !data.Dhcount.IsNull() {
		sslservice.Dhcount = utils.IntPtr(int(data.Dhcount.ValueInt64()))
	}
	if !data.Dhfile.IsNull() {
		sslservice.Dhfile = data.Dhfile.ValueString()
	}
	if !data.Dhkeyexpsizelimit.IsNull() {
		sslservice.Dhkeyexpsizelimit = data.Dhkeyexpsizelimit.ValueString()
	}
	if !data.Dtls1.IsNull() {
		sslservice.Dtls1 = data.Dtls1.ValueString()
	}
	if !data.Dtls12.IsNull() {
		sslservice.Dtls12 = data.Dtls12.ValueString()
	}
	if !data.Dtlsprofilename.IsNull() {
		sslservice.Dtlsprofilename = data.Dtlsprofilename.ValueString()
	}
	if !data.Ersa.IsNull() {
		sslservice.Ersa = data.Ersa.ValueString()
	}
	if !data.Ersacount.IsNull() {
		sslservice.Ersacount = utils.IntPtr(int(data.Ersacount.ValueInt64()))
	}
	if !data.Ocspstapling.IsNull() {
		sslservice.Ocspstapling = data.Ocspstapling.ValueString()
	}
	if !data.Pushenctrigger.IsNull() {
		sslservice.Pushenctrigger = data.Pushenctrigger.ValueString()
	}
	if !data.Redirectportrewrite.IsNull() {
		sslservice.Redirectportrewrite = data.Redirectportrewrite.ValueString()
	}
	if !data.Sendclosenotify.IsNull() {
		sslservice.Sendclosenotify = data.Sendclosenotify.ValueString()
	}
	if !data.Serverauth.IsNull() {
		sslservice.Serverauth = data.Serverauth.ValueString()
	}
	if !data.Servicename.IsNull() {
		sslservice.Servicename = data.Servicename.ValueString()
	}
	if !data.Sessreuse.IsNull() {
		sslservice.Sessreuse = data.Sessreuse.ValueString()
	}
	if !data.Sesstimeout.IsNull() {
		sslservice.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Snienable.IsNull() {
		sslservice.Snienable = data.Snienable.ValueString()
	}
	if !data.Ssl2.IsNull() {
		sslservice.Ssl2 = data.Ssl2.ValueString()
	}
	if !data.Ssl3.IsNull() {
		sslservice.Ssl3 = data.Ssl3.ValueString()
	}
	if !data.Sslclientlogs.IsNull() {
		sslservice.Sslclientlogs = data.Sslclientlogs.ValueString()
	}
	if !data.Sslprofile.IsNull() {
		sslservice.Sslprofile = data.Sslprofile.ValueString()
	}
	if !data.Sslredirect.IsNull() {
		sslservice.Sslredirect = data.Sslredirect.ValueString()
	}
	if !data.Sslv2redirect.IsNull() {
		sslservice.Sslv2redirect = data.Sslv2redirect.ValueString()
	}
	if !data.Sslv2url.IsNull() {
		sslservice.Sslv2url = data.Sslv2url.ValueString()
	}
	if !data.Strictsigdigestcheck.IsNull() {
		sslservice.Strictsigdigestcheck = data.Strictsigdigestcheck.ValueString()
	}
	if !data.Tls1.IsNull() {
		sslservice.Tls1 = data.Tls1.ValueString()
	}
	if !data.Tls11.IsNull() {
		sslservice.Tls11 = data.Tls11.ValueString()
	}
	if !data.Tls12.IsNull() {
		sslservice.Tls12 = data.Tls12.ValueString()
	}
	if !data.Tls13.IsNull() {
		sslservice.Tls13 = data.Tls13.ValueString()
	}

	return sslservice
}

func sslserviceSetAttrFromGet(ctx context.Context, data *SslserviceResourceModel, getResponseData map[string]interface{}) *SslserviceResourceModel {
	tflog.Debug(ctx, "In sslserviceSetAttrFromGet Function")

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
	if val, ok := getResponseData["commonname"]; ok && val != nil {
		data.Commonname = types.StringValue(val.(string))
	} else {
		data.Commonname = types.StringNull()
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
	if val, ok := getResponseData["ocspstapling"]; ok && val != nil {
		data.Ocspstapling = types.StringValue(val.(string))
	} else {
		data.Ocspstapling = types.StringNull()
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
	if val, ok := getResponseData["serverauth"]; ok && val != nil {
		data.Serverauth = types.StringValue(val.(string))
	} else {
		data.Serverauth = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Servicename.ValueString())

	return data
}
