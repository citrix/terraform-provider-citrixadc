package auditsyslogaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuditsyslogactionResourceModel describes the resource data model.
type AuditsyslogactionResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Dns                  types.String `tfsdk:"dns"`
	Domainresolvenow     types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry   types.Int64  `tfsdk:"domainresolveretry"`
	Httpauthtoken        types.String `tfsdk:"httpauthtoken"`
	Httpendpointurl      types.String `tfsdk:"httpendpointurl"`
	Lbvservername        types.String `tfsdk:"lbvservername"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Managementlog        types.List   `tfsdk:"managementlog"`
	Maxlogdatasizetohold types.Int64  `tfsdk:"maxlogdatasizetohold"`
	Mgmtloglevel         types.List   `tfsdk:"mgmtloglevel"`
	Name                 types.String `tfsdk:"name"`
	Netprofile           types.String `tfsdk:"netprofile"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverdomainname     types.String `tfsdk:"serverdomainname"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Streamanalytics      types.String `tfsdk:"streamanalytics"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Syslogcompliance     types.String `tfsdk:"syslogcompliance"`
	Tcp                  types.String `tfsdk:"tcp"`
	Tcpprofilename       types.String `tfsdk:"tcpprofilename"`
	Timezone             types.String `tfsdk:"timezone"`
	Transport            types.String `tfsdk:"transport"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`
}

func (r *AuditsyslogactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditsyslogaction resource.",
			},
			"acl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log access control list (ACL) messages.",
			},
			"alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log alg info",
			},
			"appflowexport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log messages to AppFlow collectors.\nAppflow collectors are entities to which log messages can be sent so that some action can be performed on them.",
			},
			"contentinspectionlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log Content Inspection event information",
			},
			"dateformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of dates in the logs.\nSupported formats are:\n* MMDDYYYY. -U.S. style month/date/year format.\n* DDMMYYYY - European style date/month/year format.\n* YYYYMMDD - ISO style year/month/date format.",
			},
			"dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log DNS related syslog messages",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.",
			},
			"httpauthtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorization header is required to be of the form - Splunk <auth-token>.",
			},
			"httpendpointurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL at which to upload the logs messages on the endpoint",
			},
			"lbvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName",
			},
			"logfacility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Facility value, as defined in RFC 3164, assigned to the log message.\nLog facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.",
			},
			"loglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Audit log level, which specifies the types of events to log.\nAvailable values function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log lsn info",
			},
			"managementlog": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Management log specifies the categories of log files to be exported.\nIt use destination and transport from PE params.\nAvailable values function as follows:\n* ALL - All categories (SHELL, NSMGMT and ACCESS).\n* SHELL -  bash.log, and sh.log.\n* ACCESS - auth.log, nsvpn.log, httpaccess.log, httperror.log, httpaccess-vpn.log and httperror-vpn.log.\n* NSMGMT - notice.log and ns.log.\n* NONE - No logs.",
			},
			"maxlogdatasizetohold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(500),
				Description: "Max size of log data that can be held in NSB chain of server info.",
			},
			"mgmtloglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Management log level, which specifies the types of events to log.\nAvailable values function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the syslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my syslog action\" or 'my syslog action').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile.\nThe SNIP configured in the network profile will be used as source IP while sending log messages.",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
			},
			"serverdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SYSLOG server name as a FQDN.  Mutually exclusive with serverIP/lbVserverName",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the syslog server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the syslog server accepts connections.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log SSL Interception event information",
			},
			"streamanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log stream analytics statistics to syslog server.",
			},
			"subscriberlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log subscriber session event information",
			},
			"syslogcompliance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC. For example, set it to RFC5424 to ensure RFC 5424 compliance",
			},
			"tcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log TCP messages.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nSupported settings are:\n* GMT_TIME. Coordinated Universal time.\n* LOCAL_TIME. Use the server's timezone setting.",
			},
			"transport": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Transport type used to send auditlogs to syslog server. Default type is UDP.",
			},
			"urlfiltering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log URL filtering event information",
			},
			"userdefinedauditlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log user-configurable log messages to syslog.\nSetting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.",
			},
		},
	}
}

func auditsyslogactionGetThePayloadFromtheConfig(ctx context.Context, data *AuditsyslogactionResourceModel) audit.Auditsyslogaction {
	tflog.Debug(ctx, "In auditsyslogactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditsyslogaction := audit.Auditsyslogaction{}
	if !data.Acl.IsNull() {
		auditsyslogaction.Acl = data.Acl.ValueString()
	}
	if !data.Alg.IsNull() {
		auditsyslogaction.Alg = data.Alg.ValueString()
	}
	if !data.Appflowexport.IsNull() {
		auditsyslogaction.Appflowexport = data.Appflowexport.ValueString()
	}
	if !data.Contentinspectionlog.IsNull() {
		auditsyslogaction.Contentinspectionlog = data.Contentinspectionlog.ValueString()
	}
	if !data.Dateformat.IsNull() {
		auditsyslogaction.Dateformat = data.Dateformat.ValueString()
	}
	if !data.Dns.IsNull() {
		auditsyslogaction.Dns = data.Dns.ValueString()
	}
	if !data.Domainresolvenow.IsNull() {
		auditsyslogaction.Domainresolvenow = data.Domainresolvenow.ValueBool()
	}
	if !data.Domainresolveretry.IsNull() {
		auditsyslogaction.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
	}
	if !data.Httpauthtoken.IsNull() {
		auditsyslogaction.Httpauthtoken = data.Httpauthtoken.ValueString()
	}
	if !data.Httpendpointurl.IsNull() {
		auditsyslogaction.Httpendpointurl = data.Httpendpointurl.ValueString()
	}
	if !data.Lbvservername.IsNull() {
		auditsyslogaction.Lbvservername = data.Lbvservername.ValueString()
	}
	if !data.Logfacility.IsNull() {
		auditsyslogaction.Logfacility = data.Logfacility.ValueString()
	}
	if !data.Lsn.IsNull() {
		auditsyslogaction.Lsn = data.Lsn.ValueString()
	}
	if !data.Maxlogdatasizetohold.IsNull() {
		auditsyslogaction.Maxlogdatasizetohold = utils.IntPtr(int(data.Maxlogdatasizetohold.ValueInt64()))
	}
	if !data.Name.IsNull() {
		auditsyslogaction.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		auditsyslogaction.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Protocolviolations.IsNull() {
		auditsyslogaction.Protocolviolations = data.Protocolviolations.ValueString()
	}
	if !data.Serverdomainname.IsNull() {
		auditsyslogaction.Serverdomainname = data.Serverdomainname.ValueString()
	}
	if !data.Serverip.IsNull() {
		auditsyslogaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		auditsyslogaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Sslinterception.IsNull() {
		auditsyslogaction.Sslinterception = data.Sslinterception.ValueString()
	}
	if !data.Streamanalytics.IsNull() {
		auditsyslogaction.Streamanalytics = data.Streamanalytics.ValueString()
	}
	if !data.Subscriberlog.IsNull() {
		auditsyslogaction.Subscriberlog = data.Subscriberlog.ValueString()
	}
	if !data.Syslogcompliance.IsNull() {
		auditsyslogaction.Syslogcompliance = data.Syslogcompliance.ValueString()
	}
	if !data.Tcp.IsNull() {
		auditsyslogaction.Tcp = data.Tcp.ValueString()
	}
	if !data.Tcpprofilename.IsNull() {
		auditsyslogaction.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Timezone.IsNull() {
		auditsyslogaction.Timezone = data.Timezone.ValueString()
	}
	if !data.Transport.IsNull() {
		auditsyslogaction.Transport = data.Transport.ValueString()
	}
	if !data.Urlfiltering.IsNull() {
		auditsyslogaction.Urlfiltering = data.Urlfiltering.ValueString()
	}
	if !data.Userdefinedauditlog.IsNull() {
		auditsyslogaction.Userdefinedauditlog = data.Userdefinedauditlog.ValueString()
	}

	return auditsyslogaction
}

func auditsyslogactionSetAttrFromGet(ctx context.Context, data *AuditsyslogactionResourceModel, getResponseData map[string]interface{}) *AuditsyslogactionResourceModel {
	tflog.Debug(ctx, "In auditsyslogactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acl"]; ok && val != nil {
		data.Acl = types.StringValue(val.(string))
	} else {
		data.Acl = types.StringNull()
	}
	if val, ok := getResponseData["alg"]; ok && val != nil {
		data.Alg = types.StringValue(val.(string))
	} else {
		data.Alg = types.StringNull()
	}
	if val, ok := getResponseData["appflowexport"]; ok && val != nil {
		data.Appflowexport = types.StringValue(val.(string))
	} else {
		data.Appflowexport = types.StringNull()
	}
	if val, ok := getResponseData["contentinspectionlog"]; ok && val != nil {
		data.Contentinspectionlog = types.StringValue(val.(string))
	} else {
		data.Contentinspectionlog = types.StringNull()
	}
	if val, ok := getResponseData["dateformat"]; ok && val != nil {
		data.Dateformat = types.StringValue(val.(string))
	} else {
		data.Dateformat = types.StringNull()
	}
	if val, ok := getResponseData["dns"]; ok && val != nil {
		data.Dns = types.StringValue(val.(string))
	} else {
		data.Dns = types.StringNull()
	}
	if val, ok := getResponseData["domainresolvenow"]; ok && val != nil {
		data.Domainresolvenow = types.BoolValue(val.(bool))
	} else {
		data.Domainresolvenow = types.BoolNull()
	}
	if val, ok := getResponseData["domainresolveretry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Domainresolveretry = types.Int64Value(intVal)
		}
	} else {
		data.Domainresolveretry = types.Int64Null()
	}
	if val, ok := getResponseData["httpauthtoken"]; ok && val != nil {
		data.Httpauthtoken = types.StringValue(val.(string))
	} else {
		data.Httpauthtoken = types.StringNull()
	}
	if val, ok := getResponseData["httpendpointurl"]; ok && val != nil {
		data.Httpendpointurl = types.StringValue(val.(string))
	} else {
		data.Httpendpointurl = types.StringNull()
	}
	if val, ok := getResponseData["lbvservername"]; ok && val != nil {
		data.Lbvservername = types.StringValue(val.(string))
	} else {
		data.Lbvservername = types.StringNull()
	}
	if val, ok := getResponseData["logfacility"]; ok && val != nil {
		data.Logfacility = types.StringValue(val.(string))
	} else {
		data.Logfacility = types.StringNull()
	}
	if val, ok := getResponseData["lsn"]; ok && val != nil {
		data.Lsn = types.StringValue(val.(string))
	} else {
		data.Lsn = types.StringNull()
	}
	if val, ok := getResponseData["maxlogdatasizetohold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxlogdatasizetohold = types.Int64Value(intVal)
		}
	} else {
		data.Maxlogdatasizetohold = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}
	if val, ok := getResponseData["protocolviolations"]; ok && val != nil {
		data.Protocolviolations = types.StringValue(val.(string))
	} else {
		data.Protocolviolations = types.StringNull()
	}
	if val, ok := getResponseData["serverdomainname"]; ok && val != nil {
		data.Serverdomainname = types.StringValue(val.(string))
	} else {
		data.Serverdomainname = types.StringNull()
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
	if val, ok := getResponseData["sslinterception"]; ok && val != nil {
		data.Sslinterception = types.StringValue(val.(string))
	} else {
		data.Sslinterception = types.StringNull()
	}
	if val, ok := getResponseData["streamanalytics"]; ok && val != nil {
		data.Streamanalytics = types.StringValue(val.(string))
	} else {
		data.Streamanalytics = types.StringNull()
	}
	if val, ok := getResponseData["subscriberlog"]; ok && val != nil {
		data.Subscriberlog = types.StringValue(val.(string))
	} else {
		data.Subscriberlog = types.StringNull()
	}
	if val, ok := getResponseData["syslogcompliance"]; ok && val != nil {
		data.Syslogcompliance = types.StringValue(val.(string))
	} else {
		data.Syslogcompliance = types.StringNull()
	}
	if val, ok := getResponseData["tcp"]; ok && val != nil {
		data.Tcp = types.StringValue(val.(string))
	} else {
		data.Tcp = types.StringNull()
	}
	if val, ok := getResponseData["tcpprofilename"]; ok && val != nil {
		data.Tcpprofilename = types.StringValue(val.(string))
	} else {
		data.Tcpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else {
		data.Timezone = types.StringNull()
	}
	if val, ok := getResponseData["transport"]; ok && val != nil {
		data.Transport = types.StringValue(val.(string))
	} else {
		data.Transport = types.StringNull()
	}
	if val, ok := getResponseData["urlfiltering"]; ok && val != nil {
		data.Urlfiltering = types.StringValue(val.(string))
	} else {
		data.Urlfiltering = types.StringNull()
	}
	if val, ok := getResponseData["userdefinedauditlog"]; ok && val != nil {
		data.Userdefinedauditlog = types.StringValue(val.(string))
	} else {
		data.Userdefinedauditlog = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
