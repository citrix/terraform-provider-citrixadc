package aaaradiusparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaaradiusparamsResourceModel describes the resource data model.
type AaaradiusparamsResourceModel struct {
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
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tunnelendpointclientip     types.String `tfsdk:"tunnelendpointclientip"`
}

func (r *AaaradiusparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaradiusparams resource.",
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure the RADIUS server state to accept or refuse accounting messages.",
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
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server.",
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
				Description: "IP attribute type in the RADIUS response.",
			},
			"ipvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID attribute in the RADIUS response.\nIf the attribute is not vendor-encoded, it is set to 0.",
			},
			"messageauthenticator": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.",
			},
			"passencoding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("mschapv2"),
				Description: "Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server.",
			},
			"pwdattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute type of the Vendor ID in the RADIUS response.",
			},
			"pwdvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the password in the RADIUS response. Used to extract the user password.",
			},
			"radattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute type for RADIUS group extraction.",
			},
			"radgroupseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radgroupsprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radkey": schema.StringAttribute{
				Required:    true,
				Description: "The key shared between the RADIUS server and clients.\nRequired for allowing the Citrix ADC to communicate with the RADIUS server.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server as the nasid part of the Radius protocol.",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the Citrix ADC IP (NSIP) address to the RADIUS server as the Network Access Server IP (NASIP) part of the Radius protocol.",
			},
			"radvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID for RADIUS group extraction.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your RADIUS server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1812),
				Description: "Port number on which the RADIUS server listens for connections.",
			},
			"tunnelendpointclientip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send Tunnel Endpoint Client IP address to the RADIUS server.",
			},
		},
	}
}

func aaaradiusparamsGetThePayloadFromtheConfig(ctx context.Context, data *AaaradiusparamsResourceModel) aaa.Aaaradiusparams {
	tflog.Debug(ctx, "In aaaradiusparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaaradiusparams := aaa.Aaaradiusparams{}
	if !data.Accounting.IsNull() {
		aaaradiusparams.Accounting = data.Accounting.ValueString()
	}
	if !data.Authentication.IsNull() {
		aaaradiusparams.Authentication = data.Authentication.ValueString()
	}
	if !data.Authservretry.IsNull() {
		aaaradiusparams.Authservretry = utils.IntPtr(int(data.Authservretry.ValueInt64()))
	}
	if !data.Authtimeout.IsNull() {
		aaaradiusparams.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Callingstationid.IsNull() {
		aaaradiusparams.Callingstationid = data.Callingstationid.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		aaaradiusparams.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Ipattributetype.IsNull() {
		aaaradiusparams.Ipattributetype = utils.IntPtr(int(data.Ipattributetype.ValueInt64()))
	}
	if !data.Ipvendorid.IsNull() {
		aaaradiusparams.Ipvendorid = utils.IntPtr(int(data.Ipvendorid.ValueInt64()))
	}
	if !data.Messageauthenticator.IsNull() {
		aaaradiusparams.Messageauthenticator = data.Messageauthenticator.ValueString()
	}
	if !data.Passencoding.IsNull() {
		aaaradiusparams.Passencoding = data.Passencoding.ValueString()
	}
	if !data.Pwdattributetype.IsNull() {
		aaaradiusparams.Pwdattributetype = utils.IntPtr(int(data.Pwdattributetype.ValueInt64()))
	}
	if !data.Pwdvendorid.IsNull() {
		aaaradiusparams.Pwdvendorid = utils.IntPtr(int(data.Pwdvendorid.ValueInt64()))
	}
	if !data.Radattributetype.IsNull() {
		aaaradiusparams.Radattributetype = utils.IntPtr(int(data.Radattributetype.ValueInt64()))
	}
	if !data.Radgroupseparator.IsNull() {
		aaaradiusparams.Radgroupseparator = data.Radgroupseparator.ValueString()
	}
	if !data.Radgroupsprefix.IsNull() {
		aaaradiusparams.Radgroupsprefix = data.Radgroupsprefix.ValueString()
	}
	if !data.Radkey.IsNull() {
		aaaradiusparams.Radkey = data.Radkey.ValueString()
	}
	if !data.Radnasid.IsNull() {
		aaaradiusparams.Radnasid = data.Radnasid.ValueString()
	}
	if !data.Radnasip.IsNull() {
		aaaradiusparams.Radnasip = data.Radnasip.ValueString()
	}
	if !data.Radvendorid.IsNull() {
		aaaradiusparams.Radvendorid = utils.IntPtr(int(data.Radvendorid.ValueInt64()))
	}
	if !data.Serverip.IsNull() {
		aaaradiusparams.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		aaaradiusparams.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Tunnelendpointclientip.IsNull() {
		aaaradiusparams.Tunnelendpointclientip = data.Tunnelendpointclientip.ValueString()
	}

	return aaaradiusparams
}

func aaaradiusparamsSetAttrFromGet(ctx context.Context, data *AaaradiusparamsResourceModel, getResponseData map[string]interface{}) *AaaradiusparamsResourceModel {
	tflog.Debug(ctx, "In aaaradiusparamsSetAttrFromGet Function")

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
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["tunnelendpointclientip"]; ok && val != nil {
		data.Tunnelendpointclientip = types.StringValue(val.(string))
	} else {
		data.Tunnelendpointclientip = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaaradiusparams-config")

	return data
}
