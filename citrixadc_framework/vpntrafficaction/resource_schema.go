package vpntrafficaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpntrafficactionResourceModel describes the resource data model.
type VpntrafficactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Apptimeout       types.Int64  `tfsdk:"apptimeout"`
	Formssoaction    types.String `tfsdk:"formssoaction"`
	Fta              types.String `tfsdk:"fta"`
	Hdx              types.String `tfsdk:"hdx"`
	Kcdaccount       types.String `tfsdk:"kcdaccount"`
	Name             types.String `tfsdk:"name"`
	Passwdexpression types.String `tfsdk:"passwdexpression"`
	Proxy            types.String `tfsdk:"proxy"`
	Qual             types.String `tfsdk:"qual"`
	Samlssoprofile   types.String `tfsdk:"samlssoprofile"`
	Sso              types.String `tfsdk:"sso"`
	Userexpression   types.String `tfsdk:"userexpression"`
	Wanscaler        types.String `tfsdk:"wanscaler"`
}

func (r *VpntrafficactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpntrafficaction resource.",
			},
			"apptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum amount of time, in minutes, a user can stay logged on to the web application.",
			},
			"formssoaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.",
			},
			"fta": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify file type association, which is a list of file extensions that users are allowed to open.",
			},
			"hdx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provide hdx proxy to the ICA traffic",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Default"),
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain password for SingleSignOn",
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request.",
			},
			"qual": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol, either HTTP or TCP, to be used with the action.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO to remote relying party",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provide single sign-on to the web application.\n	    NOTE : Authentication mechanisms like Basic-authentication  require the user credentials to be sent in plaintext which is not secure if the server is running on HTTP (instead of HTTPS).",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain username for SingleSignOn",
			},
			"wanscaler": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the Repeater Plug-in to optimize network traffic.",
			},
		},
	}
}

func vpntrafficactionGetThePayloadFromtheConfig(ctx context.Context, data *VpntrafficactionResourceModel) vpn.Vpntrafficaction {
	tflog.Debug(ctx, "In vpntrafficactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpntrafficaction := vpn.Vpntrafficaction{}
	if !data.Apptimeout.IsNull() {
		vpntrafficaction.Apptimeout = utils.IntPtr(int(data.Apptimeout.ValueInt64()))
	}
	if !data.Formssoaction.IsNull() {
		vpntrafficaction.Formssoaction = data.Formssoaction.ValueString()
	}
	if !data.Fta.IsNull() {
		vpntrafficaction.Fta = data.Fta.ValueString()
	}
	if !data.Hdx.IsNull() {
		vpntrafficaction.Hdx = data.Hdx.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		vpntrafficaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() {
		vpntrafficaction.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() {
		vpntrafficaction.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Proxy.IsNull() {
		vpntrafficaction.Proxy = data.Proxy.ValueString()
	}
	if !data.Qual.IsNull() {
		vpntrafficaction.Qual = data.Qual.ValueString()
	}
	if !data.Samlssoprofile.IsNull() {
		vpntrafficaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Sso.IsNull() {
		vpntrafficaction.Sso = data.Sso.ValueString()
	}
	if !data.Userexpression.IsNull() {
		vpntrafficaction.Userexpression = data.Userexpression.ValueString()
	}
	if !data.Wanscaler.IsNull() {
		vpntrafficaction.Wanscaler = data.Wanscaler.ValueString()
	}

	return vpntrafficaction
}

func vpntrafficactionSetAttrFromGet(ctx context.Context, data *VpntrafficactionResourceModel, getResponseData map[string]interface{}) *VpntrafficactionResourceModel {
	tflog.Debug(ctx, "In vpntrafficactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Apptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["formssoaction"]; ok && val != nil {
		data.Formssoaction = types.StringValue(val.(string))
	} else {
		data.Formssoaction = types.StringNull()
	}
	if val, ok := getResponseData["fta"]; ok && val != nil {
		data.Fta = types.StringValue(val.(string))
	} else {
		data.Fta = types.StringNull()
	}
	if val, ok := getResponseData["hdx"]; ok && val != nil {
		data.Hdx = types.StringValue(val.(string))
	} else {
		data.Hdx = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passwdexpression"]; ok && val != nil {
		data.Passwdexpression = types.StringValue(val.(string))
	} else {
		data.Passwdexpression = types.StringNull()
	}
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["qual"]; ok && val != nil {
		data.Qual = types.StringValue(val.(string))
	} else {
		data.Qual = types.StringNull()
	}
	if val, ok := getResponseData["samlssoprofile"]; ok && val != nil {
		data.Samlssoprofile = types.StringValue(val.(string))
	} else {
		data.Samlssoprofile = types.StringNull()
	}
	if val, ok := getResponseData["sso"]; ok && val != nil {
		data.Sso = types.StringValue(val.(string))
	} else {
		data.Sso = types.StringNull()
	}
	if val, ok := getResponseData["userexpression"]; ok && val != nil {
		data.Userexpression = types.StringValue(val.(string))
	} else {
		data.Userexpression = types.StringNull()
	}
	if val, ok := getResponseData["wanscaler"]; ok && val != nil {
		data.Wanscaler = types.StringValue(val.(string))
	} else {
		data.Wanscaler = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
