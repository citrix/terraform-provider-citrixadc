package systemparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SystemparameterResourceModel describes the resource data model.
type SystemparameterResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Basicauth               types.String `tfsdk:"basicauth"`
	Cliloglevel             types.String `tfsdk:"cliloglevel"`
	Daystoexpire            types.Int64  `tfsdk:"daystoexpire"`
	Doppler                 types.String `tfsdk:"doppler"`
	Fipsusermode            types.String `tfsdk:"fipsusermode"`
	Forcepasswordchange     types.String `tfsdk:"forcepasswordchange"`
	Googleanalytics         types.String `tfsdk:"googleanalytics"`
	Localauth               types.String `tfsdk:"localauth"`
	Maxsessionperuser       types.Int64  `tfsdk:"maxsessionperuser"`
	Minpasswordlen          types.Int64  `tfsdk:"minpasswordlen"`
	Natpcbforceflushlimit   types.Int64  `tfsdk:"natpcbforceflushlimit"`
	Natpcbrstontimeout      types.String `tfsdk:"natpcbrstontimeout"`
	Passwordhistorycontrol  types.String `tfsdk:"passwordhistorycontrol"`
	Promptstring            types.String `tfsdk:"promptstring"`
	Pwdhistorycount         types.Int64  `tfsdk:"pwdhistorycount"`
	Rbaonresponse           types.String `tfsdk:"rbaonresponse"`
	Reauthonauthparamchange types.String `tfsdk:"reauthonauthparamchange"`
	Removesensitivefiles    types.String `tfsdk:"removesensitivefiles"`
	Restrictedtimeout       types.String `tfsdk:"restrictedtimeout"`
	Strongpassword          types.String `tfsdk:"strongpassword"`
	Timeout                 types.Int64  `tfsdk:"timeout"`
	Totalauthtimeout        types.Int64  `tfsdk:"totalauthtimeout"`
	Wafprotection           types.List   `tfsdk:"wafprotection"`
	Warnpriorndays          types.Int64  `tfsdk:"warnpriorndays"`
}

func (r *SystemparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemparameter resource.",
			},
			"basicauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable basic authentication for Nitro API.",
			},
			"cliloglevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("INFORMATIONAL"),
				Description: "Audit log level, which specifies the types of events to log for cli executed commands.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"daystoexpire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Password expiry days for all the system users. The daystoexpire value ranges from 30 to 255.",
			},
			"doppler": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable Doppler",
			},
			"fipsusermode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this option to set the FIPS mode for key user-land processes. When enabled, these user-land processes will operate in FIPS mode. In this mode, these processes will use FIPS 140-2 certified crypto algorithms.\nWith a FIPS license, it is enabled by default and cannot be disabled.\nWithout a FIPS license, it is disabled by default, wherein these user-land processes will not operate in FIPS mode.",
			},
			"forcepasswordchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable force password change for nsroot user",
			},
			"googleanalytics": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable Google analytics",
			},
			"localauth": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "When enabled, local users can access Citrix ADC even when external authentication is configured. When disabled, local users are not allowed to access the Citrix ADC, Local users can access the Citrix ADC only when the configured external authentication servers are unavailable. This parameter is not applicable to SSH Key-based authentication",
			},
			"maxsessionperuser": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user.The maxsessionperuser value ranges from 1 to 40",
			},
			"minpasswordlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length of system user password. When strong password is enabled default minimum length is 8. User entered value can be greater than or equal to 8. Default mininum value is 1 when strong password is disabled. Maximum value is 127 in both cases.",
			},
			"natpcbforceflushlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2147483647),
				Description: "Flush the system if the number of Network Address Translation Protocol Control Blocks (NATPCBs) exceeds this value.",
			},
			"natpcbrstontimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send a reset signal to client and server connections when their NATPCBs time out. Avoids the buildup of idle TCP connections on both the sides.",
			},
			"passwordhistorycontrol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enables or disable password expiry feature for system users.\nIf the feature is ENABLED, by default the last 6 passwords of users will be maintained and will not be allowed to reuse same.\nWhen the feature is enabled the daystoexpire, warnpriorndays and pwdhistoryCount will be set with default values. The values can only be set in system\nfor system parameter. It cannot be unset. It is possible to set and unset the values for daytoexpire and warnpriorndays in system groups.\nDefault values if feature is ENABLED:\ndaystoexpire: 30\nwarnpriorndays: 5\npwdhistoryCount: 6\nIf the feature is DISABLED the values cannot be set or unset in system parameter and system groups",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"pwdhistorycount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of passwords to be maintained as history for system users. The pwdhistorycount value ranges from 1 to 10.",
			},
			"rbaonresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable Role-Based Authentication (RBA) on responses.",
			},
			"reauthonauthparamchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable External user reauthentication when authentication parameter changes",
			},
			"removesensitivefiles": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this option to remove the sensitive files from the system like authorise keys, public keys etc. The commands which will remove sensitive files when this system paramter is enabled are rm cluster instance, rm cluster node, rm ha node, clear config full, join cluster and add cluster instance.",
			},
			"restrictedtimeout": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/Disable the restricted timeout behaviour. When enabled, timeout cannot be configured beyond admin configured timeout  and also it will have the [minimum - maximum] range check. When disabled, timeout will have the old behaviour. By default the value is disabled",
			},
			"strongpassword": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("disabled"),
				Description: "After enabling strong password (enableall / enablelocal - not included in exclude list), all the passwords / sensitive information must have - Atleast 1 Lower case character, Atleast 1 Upper case character, Atleast 1 numeric character, Atleast 1 special character ( ~, `, !, @, #, $, %, ^, &, *, -, _, =, +, {, }, [, ], |, \\, :, <, >, /, ., ,, \" \"). Exclude list in case of enablelocal is - NS_FIPS, NS_CRL, NS_RSAKEY, NS_PKCS12, NS_PKCS8, NS_LDAP, NS_TACACS, NS_TACACSACTION, NS_RADIUS, NS_RADIUSACTION, NS_ENCRYPTION_PARAMS. So no Strong Password checks will be performed on these ObjectType commands for enablelocal case.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument is enabled, Timeout can have values in the range [300-86400] seconds.\nIf Restrictedtimeout argument is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"totalauthtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20),
				Description: "Total time a request can take for authentication/authorization",
			},
			"wafprotection": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Configure WAF protection for endpoints used by NetScaler management interfaces. The available options are:\n* DEFAULT - NetScaler determines which endpoints have WAF protection enabled or disabled. In the current release, WAF protection is disabled for all endpoints when this option is used. The behavior of this option may change in future releases.\n* GUI - Endpoints used by the Management GUI Interface are WAF protected.\n* DISABLED - WAF protection is disabled for all endpoints.",
			},
			"warnpriorndays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.",
			},
		},
	}
}

func systemparameterGetThePayloadFromtheConfig(ctx context.Context, data *SystemparameterResourceModel) system.Systemparameter {
	tflog.Debug(ctx, "In systemparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemparameter := system.Systemparameter{}
	if !data.Basicauth.IsNull() {
		systemparameter.Basicauth = data.Basicauth.ValueString()
	}
	if !data.Cliloglevel.IsNull() {
		systemparameter.Cliloglevel = data.Cliloglevel.ValueString()
	}
	if !data.Daystoexpire.IsNull() {
		systemparameter.Daystoexpire = utils.IntPtr(int(data.Daystoexpire.ValueInt64()))
	}
	if !data.Doppler.IsNull() {
		systemparameter.Doppler = data.Doppler.ValueString()
	}
	if !data.Fipsusermode.IsNull() {
		systemparameter.Fipsusermode = data.Fipsusermode.ValueString()
	}
	if !data.Forcepasswordchange.IsNull() {
		systemparameter.Forcepasswordchange = data.Forcepasswordchange.ValueString()
	}
	if !data.Googleanalytics.IsNull() {
		systemparameter.Googleanalytics = data.Googleanalytics.ValueString()
	}
	if !data.Localauth.IsNull() {
		systemparameter.Localauth = data.Localauth.ValueString()
	}
	if !data.Maxsessionperuser.IsNull() {
		systemparameter.Maxsessionperuser = utils.IntPtr(int(data.Maxsessionperuser.ValueInt64()))
	}
	if !data.Minpasswordlen.IsNull() {
		systemparameter.Minpasswordlen = utils.IntPtr(int(data.Minpasswordlen.ValueInt64()))
	}
	if !data.Natpcbforceflushlimit.IsNull() {
		systemparameter.Natpcbforceflushlimit = utils.IntPtr(int(data.Natpcbforceflushlimit.ValueInt64()))
	}
	if !data.Natpcbrstontimeout.IsNull() {
		systemparameter.Natpcbrstontimeout = data.Natpcbrstontimeout.ValueString()
	}
	if !data.Passwordhistorycontrol.IsNull() {
		systemparameter.Passwordhistorycontrol = data.Passwordhistorycontrol.ValueString()
	}
	if !data.Promptstring.IsNull() {
		systemparameter.Promptstring = data.Promptstring.ValueString()
	}
	if !data.Pwdhistorycount.IsNull() {
		systemparameter.Pwdhistorycount = utils.IntPtr(int(data.Pwdhistorycount.ValueInt64()))
	}
	if !data.Rbaonresponse.IsNull() {
		systemparameter.Rbaonresponse = data.Rbaonresponse.ValueString()
	}
	if !data.Reauthonauthparamchange.IsNull() {
		systemparameter.Reauthonauthparamchange = data.Reauthonauthparamchange.ValueString()
	}
	if !data.Removesensitivefiles.IsNull() {
		systemparameter.Removesensitivefiles = data.Removesensitivefiles.ValueString()
	}
	if !data.Restrictedtimeout.IsNull() {
		systemparameter.Restrictedtimeout = data.Restrictedtimeout.ValueString()
	}
	if !data.Strongpassword.IsNull() {
		systemparameter.Strongpassword = data.Strongpassword.ValueString()
	}
	if !data.Timeout.IsNull() {
		systemparameter.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Totalauthtimeout.IsNull() {
		systemparameter.Totalauthtimeout = utils.IntPtr(int(data.Totalauthtimeout.ValueInt64()))
	}
	if !data.Warnpriorndays.IsNull() {
		systemparameter.Warnpriorndays = utils.IntPtr(int(data.Warnpriorndays.ValueInt64()))
	}

	return systemparameter
}

func systemparameterSetAttrFromGet(ctx context.Context, data *SystemparameterResourceModel, getResponseData map[string]interface{}) *SystemparameterResourceModel {
	tflog.Debug(ctx, "In systemparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["basicauth"]; ok && val != nil {
		data.Basicauth = types.StringValue(val.(string))
	} else {
		data.Basicauth = types.StringNull()
	}
	if val, ok := getResponseData["cliloglevel"]; ok && val != nil {
		data.Cliloglevel = types.StringValue(val.(string))
	} else {
		data.Cliloglevel = types.StringNull()
	}
	if val, ok := getResponseData["daystoexpire"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Daystoexpire = types.Int64Value(intVal)
		}
	} else {
		data.Daystoexpire = types.Int64Null()
	}
	if val, ok := getResponseData["doppler"]; ok && val != nil {
		data.Doppler = types.StringValue(val.(string))
	} else {
		data.Doppler = types.StringNull()
	}
	if val, ok := getResponseData["fipsusermode"]; ok && val != nil {
		data.Fipsusermode = types.StringValue(val.(string))
	} else {
		data.Fipsusermode = types.StringNull()
	}
	if val, ok := getResponseData["forcepasswordchange"]; ok && val != nil {
		data.Forcepasswordchange = types.StringValue(val.(string))
	} else {
		data.Forcepasswordchange = types.StringNull()
	}
	if val, ok := getResponseData["googleanalytics"]; ok && val != nil {
		data.Googleanalytics = types.StringValue(val.(string))
	} else {
		data.Googleanalytics = types.StringNull()
	}
	if val, ok := getResponseData["localauth"]; ok && val != nil {
		data.Localauth = types.StringValue(val.(string))
	} else {
		data.Localauth = types.StringNull()
	}
	if val, ok := getResponseData["maxsessionperuser"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxsessionperuser = types.Int64Value(intVal)
		}
	} else {
		data.Maxsessionperuser = types.Int64Null()
	}
	if val, ok := getResponseData["minpasswordlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minpasswordlen = types.Int64Value(intVal)
		}
	} else {
		data.Minpasswordlen = types.Int64Null()
	}
	if val, ok := getResponseData["natpcbforceflushlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Natpcbforceflushlimit = types.Int64Value(intVal)
		}
	} else {
		data.Natpcbforceflushlimit = types.Int64Null()
	}
	if val, ok := getResponseData["natpcbrstontimeout"]; ok && val != nil {
		data.Natpcbrstontimeout = types.StringValue(val.(string))
	} else {
		data.Natpcbrstontimeout = types.StringNull()
	}
	if val, ok := getResponseData["passwordhistorycontrol"]; ok && val != nil {
		data.Passwordhistorycontrol = types.StringValue(val.(string))
	} else {
		data.Passwordhistorycontrol = types.StringNull()
	}
	if val, ok := getResponseData["promptstring"]; ok && val != nil {
		data.Promptstring = types.StringValue(val.(string))
	} else {
		data.Promptstring = types.StringNull()
	}
	if val, ok := getResponseData["pwdhistorycount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pwdhistorycount = types.Int64Value(intVal)
		}
	} else {
		data.Pwdhistorycount = types.Int64Null()
	}
	if val, ok := getResponseData["rbaonresponse"]; ok && val != nil {
		data.Rbaonresponse = types.StringValue(val.(string))
	} else {
		data.Rbaonresponse = types.StringNull()
	}
	if val, ok := getResponseData["reauthonauthparamchange"]; ok && val != nil {
		data.Reauthonauthparamchange = types.StringValue(val.(string))
	} else {
		data.Reauthonauthparamchange = types.StringNull()
	}
	if val, ok := getResponseData["removesensitivefiles"]; ok && val != nil {
		data.Removesensitivefiles = types.StringValue(val.(string))
	} else {
		data.Removesensitivefiles = types.StringNull()
	}
	if val, ok := getResponseData["restrictedtimeout"]; ok && val != nil {
		data.Restrictedtimeout = types.StringValue(val.(string))
	} else {
		data.Restrictedtimeout = types.StringNull()
	}
	if val, ok := getResponseData["strongpassword"]; ok && val != nil {
		data.Strongpassword = types.StringValue(val.(string))
	} else {
		data.Strongpassword = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["totalauthtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Totalauthtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Totalauthtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["warnpriorndays"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Warnpriorndays = types.Int64Value(intVal)
		}
	} else {
		data.Warnpriorndays = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("systemparameter-config")

	return data
}
