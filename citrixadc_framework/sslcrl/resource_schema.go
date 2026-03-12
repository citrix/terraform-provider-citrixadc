package sslcrl

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslcrlResourceModel describes the resource data model.
type SslcrlResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Basedn     types.String `tfsdk:"basedn"`
	Binary     types.String `tfsdk:"binary"`
	Binddn     types.String `tfsdk:"binddn"`
	Cacert     types.String `tfsdk:"cacert"`
	Cacertfile types.String `tfsdk:"cacertfile"`
	Cakeyfile  types.String `tfsdk:"cakeyfile"`
	Crlname    types.String `tfsdk:"crlname"`
	Crlpath    types.String `tfsdk:"crlpath"`
	Day        types.Int64  `tfsdk:"day"`
	Gencrl     types.String `tfsdk:"gencrl"`
	Indexfile  types.String `tfsdk:"indexfile"`
	Inform     types.String `tfsdk:"inform"`
	Interval   types.String `tfsdk:"interval"`
	Method     types.String `tfsdk:"method"`
	Password   types.String `tfsdk:"password"`
	Port       types.Int64  `tfsdk:"port"`
	Refresh    types.String `tfsdk:"refresh"`
	Revoke     types.String `tfsdk:"revoke"`
	Scope      types.String `tfsdk:"scope"`
	Server     types.String `tfsdk:"server"`
	Time       types.String `tfsdk:"time"`
	Url        types.String `tfsdk:"url"`
}

func (r *SslcrlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcrl resource.",
			},
			"basedn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN.",
			},
			"binary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the LDAP-based CRL retrieval mode to binary.",
			},
			"binddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL.",
			},
			"cacertfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the CA certificate file.\n/nsconfig/ssl/ is the default path.",
			},
			"cakeyfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path",
			},
			"crlname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my crl\" or 'my crl').",
			},
			"crlpath": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Path to the CRL file. /var/netscaler/ssl/ is the default path.",
			},
			"day": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY.",
			},
			"gencrl": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path.",
			},
			"indexfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path",
			},
			"inform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Input format of the CRL file. The two formats supported on the appliance are:\nPEM - Privacy Enhanced Mail.\nDER - Distinguished Encoding Rule.",
			},
			"interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CRL refresh interval. Use the NONE setting to unset this parameter.",
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port for the LDAP server.",
			},
			"refresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set CRL auto refresh.",
			},
			"revoke": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("One"),
				Description: "Extent of the search operation on the LDAP server. Available settings function as follows:\nOne - One level below Base DN.\nBase - Exactly the same level as Base DN.",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the LDAP server from which to fetch the CRLs.",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the CRL distribution point.",
			},
		},
	}
}

func sslcrlGetThePayloadFromtheConfig(ctx context.Context, data *SslcrlResourceModel) ssl.Sslcrl {
	tflog.Debug(ctx, "In sslcrlGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcrl := ssl.Sslcrl{}
	if !data.Basedn.IsNull() {
		sslcrl.Basedn = data.Basedn.ValueString()
	}
	if !data.Binary.IsNull() {
		sslcrl.Binary = data.Binary.ValueString()
	}
	if !data.Binddn.IsNull() {
		sslcrl.Binddn = data.Binddn.ValueString()
	}
	if !data.Cacert.IsNull() {
		sslcrl.Cacert = data.Cacert.ValueString()
	}
	if !data.Cacertfile.IsNull() {
		sslcrl.Cacertfile = data.Cacertfile.ValueString()
	}
	if !data.Cakeyfile.IsNull() {
		sslcrl.Cakeyfile = data.Cakeyfile.ValueString()
	}
	if !data.Crlname.IsNull() {
		sslcrl.Crlname = data.Crlname.ValueString()
	}
	if !data.Crlpath.IsNull() {
		sslcrl.Crlpath = data.Crlpath.ValueString()
	}
	if !data.Day.IsNull() {
		sslcrl.Day = utils.IntPtr(int(data.Day.ValueInt64()))
	}
	if !data.Gencrl.IsNull() {
		sslcrl.Gencrl = data.Gencrl.ValueString()
	}
	if !data.Indexfile.IsNull() {
		sslcrl.Indexfile = data.Indexfile.ValueString()
	}
	if !data.Inform.IsNull() {
		sslcrl.Inform = data.Inform.ValueString()
	}
	if !data.Interval.IsNull() {
		sslcrl.Interval = data.Interval.ValueString()
	}
	if !data.Method.IsNull() {
		sslcrl.Method = data.Method.ValueString()
	}
	if !data.Password.IsNull() {
		sslcrl.Password = data.Password.ValueString()
	}
	if !data.Port.IsNull() {
		sslcrl.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Refresh.IsNull() {
		sslcrl.Refresh = data.Refresh.ValueString()
	}
	if !data.Revoke.IsNull() {
		sslcrl.Revoke = data.Revoke.ValueString()
	}
	if !data.Scope.IsNull() {
		sslcrl.Scope = data.Scope.ValueString()
	}
	if !data.Server.IsNull() {
		sslcrl.Server = data.Server.ValueString()
	}
	if !data.Time.IsNull() {
		sslcrl.Time = data.Time.ValueString()
	}
	if !data.Url.IsNull() {
		sslcrl.Url = data.Url.ValueString()
	}

	return sslcrl
}

func sslcrlSetAttrFromGet(ctx context.Context, data *SslcrlResourceModel, getResponseData map[string]interface{}) *SslcrlResourceModel {
	tflog.Debug(ctx, "In sslcrlSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["basedn"]; ok && val != nil {
		data.Basedn = types.StringValue(val.(string))
	} else {
		data.Basedn = types.StringNull()
	}
	if val, ok := getResponseData["binary"]; ok && val != nil {
		data.Binary = types.StringValue(val.(string))
	} else {
		data.Binary = types.StringNull()
	}
	if val, ok := getResponseData["binddn"]; ok && val != nil {
		data.Binddn = types.StringValue(val.(string))
	} else {
		data.Binddn = types.StringNull()
	}
	if val, ok := getResponseData["cacert"]; ok && val != nil {
		data.Cacert = types.StringValue(val.(string))
	} else {
		data.Cacert = types.StringNull()
	}
	if val, ok := getResponseData["cacertfile"]; ok && val != nil {
		data.Cacertfile = types.StringValue(val.(string))
	} else {
		data.Cacertfile = types.StringNull()
	}
	if val, ok := getResponseData["cakeyfile"]; ok && val != nil {
		data.Cakeyfile = types.StringValue(val.(string))
	} else {
		data.Cakeyfile = types.StringNull()
	}
	if val, ok := getResponseData["crlname"]; ok && val != nil {
		data.Crlname = types.StringValue(val.(string))
	} else {
		data.Crlname = types.StringNull()
	}
	if val, ok := getResponseData["crlpath"]; ok && val != nil {
		data.Crlpath = types.StringValue(val.(string))
	} else {
		data.Crlpath = types.StringNull()
	}
	if val, ok := getResponseData["day"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Day = types.Int64Value(intVal)
		}
	} else {
		data.Day = types.Int64Null()
	}
	if val, ok := getResponseData["gencrl"]; ok && val != nil {
		data.Gencrl = types.StringValue(val.(string))
	} else {
		data.Gencrl = types.StringNull()
	}
	if val, ok := getResponseData["indexfile"]; ok && val != nil {
		data.Indexfile = types.StringValue(val.(string))
	} else {
		data.Indexfile = types.StringNull()
	}
	if val, ok := getResponseData["inform"]; ok && val != nil {
		data.Inform = types.StringValue(val.(string))
	} else {
		data.Inform = types.StringNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		data.Interval = types.StringValue(val.(string))
	} else {
		data.Interval = types.StringNull()
	}
	if val, ok := getResponseData["method"]; ok && val != nil {
		data.Method = types.StringValue(val.(string))
	} else {
		data.Method = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["refresh"]; ok && val != nil {
		data.Refresh = types.StringValue(val.(string))
	} else {
		data.Refresh = types.StringNull()
	}
	if val, ok := getResponseData["revoke"]; ok && val != nil {
		data.Revoke = types.StringValue(val.(string))
	} else {
		data.Revoke = types.StringNull()
	}
	if val, ok := getResponseData["scope"]; ok && val != nil {
		data.Scope = types.StringValue(val.(string))
	} else {
		data.Scope = types.StringNull()
	}
	if val, ok := getResponseData["server"]; ok && val != nil {
		data.Server = types.StringValue(val.(string))
	} else {
		data.Server = types.StringNull()
	}
	if val, ok := getResponseData["time"]; ok && val != nil {
		data.Time = types.StringValue(val.(string))
	} else {
		data.Time = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Crlname.ValueString())

	return data
}
