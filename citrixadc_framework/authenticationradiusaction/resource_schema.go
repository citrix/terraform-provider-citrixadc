package authenticationradiusaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationradiusactionResourceModel describes the resource data model.
type AuthenticationradiusactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Authentication             types.String `tfsdk:"authentication"`
	Authservretry              types.Int64  `tfsdk:"authservretry"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Callingstationid           types.String `tfsdk:"callingstationid"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Ipattributetype            types.Int64  `tfsdk:"ipattributetype"`
	Ipvendorid                 types.Int64  `tfsdk:"ipvendorid"`
	Messageauthenticator       types.String `tfsdk:"messageauthenticator"`
	Name                       types.String `tfsdk:"name"`
	Passencoding               types.String `tfsdk:"passencoding"`
	Pwdattributetype           types.Int64  `tfsdk:"pwdattributetype"`
	Pwdvendorid                types.Int64  `tfsdk:"pwdvendorid"`
	Radattributetype           types.Int64  `tfsdk:"radattributetype"`
	Radgroupseparator          types.String `tfsdk:"radgroupseparator"`
	Radgroupsprefix            types.String `tfsdk:"radgroupsprefix"`
	Radkey                     types.String `tfsdk:"radkey"`
	Radnasid                   types.String `tfsdk:"radnasid"`
	Radnasip                   types.String `tfsdk:"radnasip"`
	Radvendorid                types.Int64  `tfsdk:"radvendorid"`
	Serverip                   types.String `tfsdk:"serverip"`
	Servername                 types.String `tfsdk:"servername"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Targetlbvserver            types.String `tfsdk:"targetlbvserver"`
	Transport                  types.String `tfsdk:"transport"`
	Tunnelendpointclientip     types.String `tfsdk:"tunnelendpointclientip"`
}

func (r *AuthenticationradiusactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationradiusaction resource.",
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the RADIUS server is currently accepting accounting messages.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Configure the RADIUS server state to accept or refuse authentication messages.",
			},
			"authservretry": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of retry by the Citrix ADC before getting response from the RADIUS server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of seconds the Citrix ADC waits for a response from the RADIUS server.",
			},
			"callingstationid": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"ipattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Remote IP address attribute type in a RADIUS response.",
			},
			"ipvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the intranet IP attribute in the RADIUS response.\nNOTE: A value of 0 indicates that the attribute is not vendor encoded.",
			},
			"messageauthenticator": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RADIUS action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.",
			},
			"passencoding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("pap"),
				Description: "Encoding type for passwords in RADIUS packets that the Citrix ADC sends to the RADIUS server.",
			},
			"pwdattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-specific password attribute type in a RADIUS response.",
			},
			"pwdvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the attribute, in the RADIUS response, used to extract the user password.",
			},
			"radattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS attribute type, used for RADIUS group extraction.",
			},
			"radgroupseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS group separator string\nThe group separator delimits group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radgroupsprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS groups prefix string.\nThis groups prefix precedes the group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radkey": schema.StringAttribute{
				Required:    true,
				Description: "Key shared between the RADIUS server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the RADIUS server.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If configured, this string is sent to the RADIUS server as the Network Access Server ID (NASID).",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, the Citrix ADC IP address (NSIP) is sent to the RADIUS server as the  Network Access Server IP (NASIP) address.\nThe RADIUS protocol defines the meaning and use of the NASIP address.",
			},
			"radvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS vendor ID attribute, used for RADIUS group extraction.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address assigned to the RADIUS server.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS server name as a FQDN.  Mutually exclusive with RADIUS IP address.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the RADIUS server listens for connections.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If transport mode is TLS, specify the name of LB vserver to associate. The LB vserver needs to be of type TCP and service associated needs to be SSL_TCP",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("UDP"),
				Description: "Transport mode to RADIUS server.",
			},
			"tunnelendpointclientip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send Tunnel Endpoint Client IP address to the RADIUS server.",
			},
		},
	}
}

func authenticationradiusactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationradiusactionResourceModel) authentication.Authenticationradiusaction {
	tflog.Debug(ctx, "In authenticationradiusactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationradiusaction := authentication.Authenticationradiusaction{}
	if !data.Accounting.IsNull() {
		authenticationradiusaction.Accounting = data.Accounting.ValueString()
	}
	if !data.Authentication.IsNull() {
		authenticationradiusaction.Authentication = data.Authentication.ValueString()
	}
	if !data.Authservretry.IsNull() {
		authenticationradiusaction.Authservretry = utils.IntPtr(int(data.Authservretry.ValueInt64()))
	}
	if !data.Authtimeout.IsNull() {
		authenticationradiusaction.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Callingstationid.IsNull() {
		authenticationradiusaction.Callingstationid = data.Callingstationid.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationradiusaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Ipattributetype.IsNull() {
		authenticationradiusaction.Ipattributetype = utils.IntPtr(int(data.Ipattributetype.ValueInt64()))
	}
	if !data.Ipvendorid.IsNull() {
		authenticationradiusaction.Ipvendorid = utils.IntPtr(int(data.Ipvendorid.ValueInt64()))
	}
	if !data.Messageauthenticator.IsNull() {
		authenticationradiusaction.Messageauthenticator = data.Messageauthenticator.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationradiusaction.Name = data.Name.ValueString()
	}
	if !data.Passencoding.IsNull() {
		authenticationradiusaction.Passencoding = data.Passencoding.ValueString()
	}
	if !data.Pwdattributetype.IsNull() {
		authenticationradiusaction.Pwdattributetype = utils.IntPtr(int(data.Pwdattributetype.ValueInt64()))
	}
	if !data.Pwdvendorid.IsNull() {
		authenticationradiusaction.Pwdvendorid = utils.IntPtr(int(data.Pwdvendorid.ValueInt64()))
	}
	if !data.Radattributetype.IsNull() {
		authenticationradiusaction.Radattributetype = utils.IntPtr(int(data.Radattributetype.ValueInt64()))
	}
	if !data.Radgroupseparator.IsNull() {
		authenticationradiusaction.Radgroupseparator = data.Radgroupseparator.ValueString()
	}
	if !data.Radgroupsprefix.IsNull() {
		authenticationradiusaction.Radgroupsprefix = data.Radgroupsprefix.ValueString()
	}
	if !data.Radkey.IsNull() {
		authenticationradiusaction.Radkey = data.Radkey.ValueString()
	}
	if !data.Radnasid.IsNull() {
		authenticationradiusaction.Radnasid = data.Radnasid.ValueString()
	}
	if !data.Radnasip.IsNull() {
		authenticationradiusaction.Radnasip = data.Radnasip.ValueString()
	}
	if !data.Radvendorid.IsNull() {
		authenticationradiusaction.Radvendorid = utils.IntPtr(int(data.Radvendorid.ValueInt64()))
	}
	if !data.Serverip.IsNull() {
		authenticationradiusaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() {
		authenticationradiusaction.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() {
		authenticationradiusaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Targetlbvserver.IsNull() {
		authenticationradiusaction.Targetlbvserver = data.Targetlbvserver.ValueString()
	}
	if !data.Transport.IsNull() {
		authenticationradiusaction.Transport = data.Transport.ValueString()
	}
	if !data.Tunnelendpointclientip.IsNull() {
		authenticationradiusaction.Tunnelendpointclientip = data.Tunnelendpointclientip.ValueString()
	}

	return authenticationradiusaction
}

func authenticationradiusactionSetAttrFromGet(ctx context.Context, data *AuthenticationradiusactionResourceModel, getResponseData map[string]interface{}) *AuthenticationradiusactionResourceModel {
	tflog.Debug(ctx, "In authenticationradiusactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accounting"]; ok && val != nil {
		data.Accounting = types.StringValue(val.(string))
	} else {
		data.Accounting = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authservretry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authservretry = types.Int64Value(intVal)
		}
	} else {
		data.Authservretry = types.Int64Null()
	}
	if val, ok := getResponseData["authtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Authtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["callingstationid"]; ok && val != nil {
		data.Callingstationid = types.StringValue(val.(string))
	} else {
		data.Callingstationid = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["ipattributetype"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ipattributetype = types.Int64Value(intVal)
		}
	} else {
		data.Ipattributetype = types.Int64Null()
	}
	if val, ok := getResponseData["ipvendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ipvendorid = types.Int64Value(intVal)
		}
	} else {
		data.Ipvendorid = types.Int64Null()
	}
	if val, ok := getResponseData["messageauthenticator"]; ok && val != nil {
		data.Messageauthenticator = types.StringValue(val.(string))
	} else {
		data.Messageauthenticator = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passencoding"]; ok && val != nil {
		data.Passencoding = types.StringValue(val.(string))
	} else {
		data.Passencoding = types.StringNull()
	}
	if val, ok := getResponseData["pwdattributetype"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pwdattributetype = types.Int64Value(intVal)
		}
	} else {
		data.Pwdattributetype = types.Int64Null()
	}
	if val, ok := getResponseData["pwdvendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pwdvendorid = types.Int64Value(intVal)
		}
	} else {
		data.Pwdvendorid = types.Int64Null()
	}
	if val, ok := getResponseData["radattributetype"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Radattributetype = types.Int64Value(intVal)
		}
	} else {
		data.Radattributetype = types.Int64Null()
	}
	if val, ok := getResponseData["radgroupseparator"]; ok && val != nil {
		data.Radgroupseparator = types.StringValue(val.(string))
	} else {
		data.Radgroupseparator = types.StringNull()
	}
	if val, ok := getResponseData["radgroupsprefix"]; ok && val != nil {
		data.Radgroupsprefix = types.StringValue(val.(string))
	} else {
		data.Radgroupsprefix = types.StringNull()
	}
	if val, ok := getResponseData["radkey"]; ok && val != nil {
		data.Radkey = types.StringValue(val.(string))
	} else {
		data.Radkey = types.StringNull()
	}
	if val, ok := getResponseData["radnasid"]; ok && val != nil {
		data.Radnasid = types.StringValue(val.(string))
	} else {
		data.Radnasid = types.StringNull()
	}
	if val, ok := getResponseData["radnasip"]; ok && val != nil {
		data.Radnasip = types.StringValue(val.(string))
	} else {
		data.Radnasip = types.StringNull()
	}
	if val, ok := getResponseData["radvendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Radvendorid = types.Int64Value(intVal)
		}
	} else {
		data.Radvendorid = types.Int64Null()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}
	if val, ok := getResponseData["transport"]; ok && val != nil {
		data.Transport = types.StringValue(val.(string))
	} else {
		data.Transport = types.StringNull()
	}
	if val, ok := getResponseData["tunnelendpointclientip"]; ok && val != nil {
		data.Tunnelendpointclientip = types.StringValue(val.(string))
	} else {
		data.Tunnelendpointclientip = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
