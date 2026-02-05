package aaaparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaaparameterResourceModel describes the resource data model.
type AaaparameterResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Aaadloglevel               types.String `tfsdk:"aaadloglevel"`
	Aaadnatip                  types.String `tfsdk:"aaadnatip"`
	Aaasessionloglevel         types.String `tfsdk:"aaasessionloglevel"`
	Apitokencache              types.String `tfsdk:"apitokencache"`
	Defaultauthtype            types.String `tfsdk:"defaultauthtype"`
	Defaultcspheader           types.String `tfsdk:"defaultcspheader"`
	Dynaddr                    types.String `tfsdk:"dynaddr"`
	Enableenhancedauthfeedback types.String `tfsdk:"enableenhancedauthfeedback"`
	Enablesessionstickiness    types.String `tfsdk:"enablesessionstickiness"`
	Enablestaticpagecaching    types.String `tfsdk:"enablestaticpagecaching"`
	Enhancedepa                types.String `tfsdk:"enhancedepa"`
	Failedlogintimeout         types.Int64  `tfsdk:"failedlogintimeout"`
	Ftmode                     types.String `tfsdk:"ftmode"`
	Httponlycookie             types.String `tfsdk:"httponlycookie"`
	Loginencryption            types.String `tfsdk:"loginencryption"`
	Maxaaausers                types.Int64  `tfsdk:"maxaaausers"`
	Maxkbquestions             types.Int64  `tfsdk:"maxkbquestions"`
	Maxloginattempts           types.Int64  `tfsdk:"maxloginattempts"`
	Maxsamldeflatesize         types.Int64  `tfsdk:"maxsamldeflatesize"`
	Persistentloginattempts    types.String `tfsdk:"persistentloginattempts"`
	Pwdexpirynotificationdays  types.Int64  `tfsdk:"pwdexpirynotificationdays"`
	Samesite                   types.String `tfsdk:"samesite"`
	Securityinsights           types.String `tfsdk:"securityinsights"`
	Tokenintrospectioninterval types.Int64  `tfsdk:"tokenintrospectioninterval"`
	Wafprotection              types.List   `tfsdk:"wafprotection"`
}

func (r *AaaparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaparameter resource.",
			},
			"aaadloglevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("INFORMATIONAL"),
				Description: "AAAD log level, which specifies the types of AAAD events to log in nsvpn.log.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"aaadnatip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to use for traffic that is sent to the authentication server.",
			},
			"aaasessionloglevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DEFAULT_LOGLEVEL_AAA"),
				Description: "Audit log level, which specifies the types of events to log for cli executed commands.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"apitokencache": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Option to enable/disable API cache feature.",
			},
			"defaultauthtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("LOCAL"),
				Description: "The default authentication server type.",
			},
			"defaultcspheader": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Parameter to enable/disable default CSP header",
			},
			"dynaddr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set by the DHCP client when the IP address was fetched dynamically.",
			},
			"enableenhancedauthfeedback": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enhanced auth feedback provides more information to the end user about the reason for an authentication failure.  The default value is set to NO.",
			},
			"enablesessionstickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables/Disables stickiness to authentication servers",
			},
			"enablestaticpagecaching": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "The default state of VPN Static Page caching. Static Page caching is enabled by default.",
			},
			"enhancedepa": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter to enable/disable EPA v2 functionality",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"ftmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "First time user mode determines which configuration options are shown by default when logging in to the GUI. This setting is controlled by the GUI.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Parameter to set/reset HttpOnly Flag for NSC_AAAC/NSC_TMAS cookies in nfactor",
			},
			"loginencryption": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Parameter to encrypt login information for nFactor flow",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent users allowed to log on to VPN simultaneously.",
			},
			"maxkbquestions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum Number of login Attempts",
			},
			"maxsamldeflatesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the maximum deflate size in case of SAML Redirect binding.",
			},
			"persistentloginattempts": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Persistent storage of unsuccessful user login attempts",
			},
			"pwdexpirynotificationdays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"securityinsights": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors when request comes to Authentication endpoint.\n* If cs vserver is frontend with Authentication vserver as target for cs action, then record is sent using Authentication vserver name.\n* If vpn/lb/cs vserver are configured with Authentication ON, then then record is sent using vpn/lb/cs vserver name accordingly.\n* If authentication vserver is frontend, then record is sent using Authentication vserver name.",
			},
			"tokenintrospectioninterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.",
			},
			"wafprotection": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Entities for which WAF Protection need to be applied.\nAvailable settings function as follows:\n* DEFAULT - No Endpoint WAF protection.\n* AUTH - Endpoints used for Authentication applicable for both AAATM, IDP, GATEWAY use cases.\n* VPN - Endpoints used for Gateway use cases.\n* PORTAL - Endpoints related to web portal.\n* DISABLED - No Endpoint WAF protection.\nCurrently supported only in default partition",
			},
		},
	}
}

func aaaparameterGetThePayloadFromtheConfig(ctx context.Context, data *AaaparameterResourceModel) aaa.Aaaparameter {
	tflog.Debug(ctx, "In aaaparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaaparameter := aaa.Aaaparameter{}
	if !data.Aaadloglevel.IsNull() {
		aaaparameter.Aaadloglevel = data.Aaadloglevel.ValueString()
	}
	if !data.Aaadnatip.IsNull() {
		aaaparameter.Aaadnatip = data.Aaadnatip.ValueString()
	}
	if !data.Aaasessionloglevel.IsNull() {
		aaaparameter.Aaasessionloglevel = data.Aaasessionloglevel.ValueString()
	}
	if !data.Apitokencache.IsNull() {
		aaaparameter.Apitokencache = data.Apitokencache.ValueString()
	}
	if !data.Defaultauthtype.IsNull() {
		aaaparameter.Defaultauthtype = data.Defaultauthtype.ValueString()
	}
	if !data.Defaultcspheader.IsNull() {
		aaaparameter.Defaultcspheader = data.Defaultcspheader.ValueString()
	}
	if !data.Dynaddr.IsNull() {
		aaaparameter.Dynaddr = data.Dynaddr.ValueString()
	}
	if !data.Enableenhancedauthfeedback.IsNull() {
		aaaparameter.Enableenhancedauthfeedback = data.Enableenhancedauthfeedback.ValueString()
	}
	if !data.Enablesessionstickiness.IsNull() {
		aaaparameter.Enablesessionstickiness = data.Enablesessionstickiness.ValueString()
	}
	if !data.Enablestaticpagecaching.IsNull() {
		aaaparameter.Enablestaticpagecaching = data.Enablestaticpagecaching.ValueString()
	}
	if !data.Enhancedepa.IsNull() {
		aaaparameter.Enhancedepa = data.Enhancedepa.ValueString()
	}
	if !data.Failedlogintimeout.IsNull() {
		aaaparameter.Failedlogintimeout = utils.IntPtr(int(data.Failedlogintimeout.ValueInt64()))
	}
	if !data.Ftmode.IsNull() {
		aaaparameter.Ftmode = data.Ftmode.ValueString()
	}
	if !data.Httponlycookie.IsNull() {
		aaaparameter.Httponlycookie = data.Httponlycookie.ValueString()
	}
	if !data.Loginencryption.IsNull() {
		aaaparameter.Loginencryption = data.Loginencryption.ValueString()
	}
	if !data.Maxaaausers.IsNull() {
		aaaparameter.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
	}
	if !data.Maxkbquestions.IsNull() {
		aaaparameter.Maxkbquestions = utils.IntPtr(int(data.Maxkbquestions.ValueInt64()))
	}
	if !data.Maxloginattempts.IsNull() {
		aaaparameter.Maxloginattempts = utils.IntPtr(int(data.Maxloginattempts.ValueInt64()))
	}
	if !data.Maxsamldeflatesize.IsNull() {
		aaaparameter.Maxsamldeflatesize = utils.IntPtr(int(data.Maxsamldeflatesize.ValueInt64()))
	}
	if !data.Persistentloginattempts.IsNull() {
		aaaparameter.Persistentloginattempts = data.Persistentloginattempts.ValueString()
	}
	if !data.Pwdexpirynotificationdays.IsNull() {
		aaaparameter.Pwdexpirynotificationdays = utils.IntPtr(int(data.Pwdexpirynotificationdays.ValueInt64()))
	}
	if !data.Samesite.IsNull() {
		aaaparameter.Samesite = data.Samesite.ValueString()
	}
	if !data.Securityinsights.IsNull() {
		aaaparameter.Securityinsights = data.Securityinsights.ValueString()
	}
	if !data.Tokenintrospectioninterval.IsNull() {
		aaaparameter.Tokenintrospectioninterval = utils.IntPtr(int(data.Tokenintrospectioninterval.ValueInt64()))
	}

	return aaaparameter
}

func aaaparameterSetAttrFromGet(ctx context.Context, data *AaaparameterResourceModel, getResponseData map[string]interface{}) *AaaparameterResourceModel {
	tflog.Debug(ctx, "In aaaparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aaadloglevel"]; ok && val != nil {
		data.Aaadloglevel = types.StringValue(val.(string))
	} else {
		data.Aaadloglevel = types.StringNull()
	}
	if val, ok := getResponseData["aaadnatip"]; ok && val != nil {
		data.Aaadnatip = types.StringValue(val.(string))
	} else {
		data.Aaadnatip = types.StringNull()
	}
	if val, ok := getResponseData["aaasessionloglevel"]; ok && val != nil {
		data.Aaasessionloglevel = types.StringValue(val.(string))
	} else {
		data.Aaasessionloglevel = types.StringNull()
	}
	if val, ok := getResponseData["apitokencache"]; ok && val != nil {
		data.Apitokencache = types.StringValue(val.(string))
	} else {
		data.Apitokencache = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthtype"]; ok && val != nil {
		data.Defaultauthtype = types.StringValue(val.(string))
	} else {
		data.Defaultauthtype = types.StringNull()
	}
	if val, ok := getResponseData["defaultcspheader"]; ok && val != nil {
		data.Defaultcspheader = types.StringValue(val.(string))
	} else {
		data.Defaultcspheader = types.StringNull()
	}
	if val, ok := getResponseData["dynaddr"]; ok && val != nil {
		data.Dynaddr = types.StringValue(val.(string))
	} else {
		data.Dynaddr = types.StringNull()
	}
	if val, ok := getResponseData["enableenhancedauthfeedback"]; ok && val != nil {
		data.Enableenhancedauthfeedback = types.StringValue(val.(string))
	} else {
		data.Enableenhancedauthfeedback = types.StringNull()
	}
	if val, ok := getResponseData["enablesessionstickiness"]; ok && val != nil {
		data.Enablesessionstickiness = types.StringValue(val.(string))
	} else {
		data.Enablesessionstickiness = types.StringNull()
	}
	if val, ok := getResponseData["enablestaticpagecaching"]; ok && val != nil {
		data.Enablestaticpagecaching = types.StringValue(val.(string))
	} else {
		data.Enablestaticpagecaching = types.StringNull()
	}
	if val, ok := getResponseData["enhancedepa"]; ok && val != nil {
		data.Enhancedepa = types.StringValue(val.(string))
	} else {
		data.Enhancedepa = types.StringNull()
	}
	if val, ok := getResponseData["failedlogintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Failedlogintimeout = types.Int64Value(intVal)
		}
	} else {
		data.Failedlogintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ftmode"]; ok && val != nil {
		data.Ftmode = types.StringValue(val.(string))
	} else {
		data.Ftmode = types.StringNull()
	}
	if val, ok := getResponseData["httponlycookie"]; ok && val != nil {
		data.Httponlycookie = types.StringValue(val.(string))
	} else {
		data.Httponlycookie = types.StringNull()
	}
	if val, ok := getResponseData["loginencryption"]; ok && val != nil {
		data.Loginencryption = types.StringValue(val.(string))
	} else {
		data.Loginencryption = types.StringNull()
	}
	if val, ok := getResponseData["maxaaausers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaaausers = types.Int64Value(intVal)
		}
	} else {
		data.Maxaaausers = types.Int64Null()
	}
	if val, ok := getResponseData["maxkbquestions"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxkbquestions = types.Int64Value(intVal)
		}
	} else {
		data.Maxkbquestions = types.Int64Null()
	}
	if val, ok := getResponseData["maxloginattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxloginattempts = types.Int64Value(intVal)
		}
	} else {
		data.Maxloginattempts = types.Int64Null()
	}
	if val, ok := getResponseData["maxsamldeflatesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsamldeflatesize = types.Int64Value(intVal)
		}
	} else {
		data.Maxsamldeflatesize = types.Int64Null()
	}
	if val, ok := getResponseData["persistentloginattempts"]; ok && val != nil {
		data.Persistentloginattempts = types.StringValue(val.(string))
	} else {
		data.Persistentloginattempts = types.StringNull()
	}
	if val, ok := getResponseData["pwdexpirynotificationdays"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pwdexpirynotificationdays = types.Int64Value(intVal)
		}
	} else {
		data.Pwdexpirynotificationdays = types.Int64Null()
	}
	if val, ok := getResponseData["samesite"]; ok && val != nil {
		data.Samesite = types.StringValue(val.(string))
	} else {
		data.Samesite = types.StringNull()
	}
	if val, ok := getResponseData["securityinsights"]; ok && val != nil {
		data.Securityinsights = types.StringValue(val.(string))
	} else {
		data.Securityinsights = types.StringNull()
	}
	if val, ok := getResponseData["tokenintrospectioninterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tokenintrospectioninterval = types.Int64Value(intVal)
		}
	} else {
		data.Tokenintrospectioninterval = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaaparameter-config")

	return data
}
