package sslservicegroup

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

// SslservicegroupResourceModel describes the resource data model.
type SslservicegroupResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Commonname           types.String `tfsdk:"commonname"`
	Ocspstapling         types.String `tfsdk:"ocspstapling"`
	Sendclosenotify      types.String `tfsdk:"sendclosenotify"`
	Serverauth           types.String `tfsdk:"serverauth"`
	Servicegroupname     types.String `tfsdk:"servicegroupname"`
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
}

func (r *SslservicegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup resource.",
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server",
			},
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:\nENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.\nDISABLED: The appliance does not check the status of the server certificate.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of server authentication support for the SSL service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service group for which to set advanced configuration.",
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
				Description: "State of the Server Name Indication (SNI) feature on the service. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of SSLv3 protocol support for the SSL service group.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the Service Group.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter indicating to check whether peer's certificate is signed with one of signature-hash combination supported by Citrix ADC",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.0 protocol support for the SSL service group.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.1 protocol support for the SSL service group.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of TLSv1.2 protocol support for the SSL service group.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of TLSv1.3 protocol support for the SSL service group.",
			},
		},
	}
}

func sslservicegroupGetThePayloadFromtheConfig(ctx context.Context, data *SslservicegroupResourceModel) ssl.Sslservicegroup {
	tflog.Debug(ctx, "In sslservicegroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservicegroup := ssl.Sslservicegroup{}
	if !data.Commonname.IsNull() {
		sslservicegroup.Commonname = data.Commonname.ValueString()
	}
	if !data.Ocspstapling.IsNull() {
		sslservicegroup.Ocspstapling = data.Ocspstapling.ValueString()
	}
	if !data.Sendclosenotify.IsNull() {
		sslservicegroup.Sendclosenotify = data.Sendclosenotify.ValueString()
	}
	if !data.Serverauth.IsNull() {
		sslservicegroup.Serverauth = data.Serverauth.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		sslservicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Sessreuse.IsNull() {
		sslservicegroup.Sessreuse = data.Sessreuse.ValueString()
	}
	if !data.Sesstimeout.IsNull() {
		sslservicegroup.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Snienable.IsNull() {
		sslservicegroup.Snienable = data.Snienable.ValueString()
	}
	if !data.Ssl3.IsNull() {
		sslservicegroup.Ssl3 = data.Ssl3.ValueString()
	}
	if !data.Sslclientlogs.IsNull() {
		sslservicegroup.Sslclientlogs = data.Sslclientlogs.ValueString()
	}
	if !data.Sslprofile.IsNull() {
		sslservicegroup.Sslprofile = data.Sslprofile.ValueString()
	}
	if !data.Strictsigdigestcheck.IsNull() {
		sslservicegroup.Strictsigdigestcheck = data.Strictsigdigestcheck.ValueString()
	}
	if !data.Tls1.IsNull() {
		sslservicegroup.Tls1 = data.Tls1.ValueString()
	}
	if !data.Tls11.IsNull() {
		sslservicegroup.Tls11 = data.Tls11.ValueString()
	}
	if !data.Tls12.IsNull() {
		sslservicegroup.Tls12 = data.Tls12.ValueString()
	}
	if !data.Tls13.IsNull() {
		sslservicegroup.Tls13 = data.Tls13.ValueString()
	}

	return sslservicegroup
}

func sslservicegroupSetAttrFromGet(ctx context.Context, data *SslservicegroupResourceModel, getResponseData map[string]interface{}) *SslservicegroupResourceModel {
	tflog.Debug(ctx, "In sslservicegroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["commonname"]; ok && val != nil {
		data.Commonname = types.StringValue(val.(string))
	} else {
		data.Commonname = types.StringNull()
	}
	if val, ok := getResponseData["ocspstapling"]; ok && val != nil {
		data.Ocspstapling = types.StringValue(val.(string))
	} else {
		data.Ocspstapling = types.StringNull()
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
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
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
	data.Id = types.StringValue(data.Servicegroupname.ValueString())

	return data
}
