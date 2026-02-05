package auditnslogaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuditnslogactionResourceModel describes the resource data model.
type AuditnslogactionResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Domainresolvenow     types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry   types.Int64  `tfsdk:"domainresolveretry"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Name                 types.String `tfsdk:"name"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverdomainname     types.String `tfsdk:"serverdomainname"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Tcp                  types.String `tfsdk:"tcp"`
	Timezone             types.String `tfsdk:"timezone"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`
}

func (r *AuditnslogactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditnslogaction resource.",
			},
			"acl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log access control list (ACL) messages.",
			},
			"alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log the ALG messages",
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
				Description: "Format of dates in the logs.\nSupported formats are:\n* MMDDYYYY - U.S. style month/date/year format.\n* DDMMYYYY - European style date/month/year format.\n* YYYYMMDD - ISO style year/month/date format.",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the audit server if the last query failed.",
			},
			"logfacility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Facility value, as defined in RFC 3164, assigned to the log message.\nLog facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.",
			},
			"loglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Audit log level, which specifies the types of events to log.\nAvailable settings function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log the LSN messages",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my nslog action\" or 'my nslog action').",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
			},
			"serverdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auditserver name as a FQDN. Mutually exclusive with serverIP",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the nslog server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the nslog server accepts connections.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log SSL Interception event information",
			},
			"subscriberlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log subscriber session event information",
			},
			"tcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log TCP messages.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nAvailable settings function as follows:\n* GMT_TIME. Coordinated Universal Time.\n* LOCAL_TIME. The server's timezone setting.",
			},
			"urlfiltering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log URL filtering event information",
			},
			"userdefinedauditlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log user-configurable log messages to nslog.\nSetting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.",
			},
		},
	}
}

func auditnslogactionGetThePayloadFromtheConfig(ctx context.Context, data *AuditnslogactionResourceModel) audit.Auditnslogaction {
	tflog.Debug(ctx, "In auditnslogactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditnslogaction := audit.Auditnslogaction{}
	if !data.Acl.IsNull() {
		auditnslogaction.Acl = data.Acl.ValueString()
	}
	if !data.Alg.IsNull() {
		auditnslogaction.Alg = data.Alg.ValueString()
	}
	if !data.Appflowexport.IsNull() {
		auditnslogaction.Appflowexport = data.Appflowexport.ValueString()
	}
	if !data.Contentinspectionlog.IsNull() {
		auditnslogaction.Contentinspectionlog = data.Contentinspectionlog.ValueString()
	}
	if !data.Dateformat.IsNull() {
		auditnslogaction.Dateformat = data.Dateformat.ValueString()
	}
	if !data.Domainresolvenow.IsNull() {
		auditnslogaction.Domainresolvenow = data.Domainresolvenow.ValueBool()
	}
	if !data.Domainresolveretry.IsNull() {
		auditnslogaction.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
	}
	if !data.Logfacility.IsNull() {
		auditnslogaction.Logfacility = data.Logfacility.ValueString()
	}
	if !data.Lsn.IsNull() {
		auditnslogaction.Lsn = data.Lsn.ValueString()
	}
	if !data.Name.IsNull() {
		auditnslogaction.Name = data.Name.ValueString()
	}
	if !data.Protocolviolations.IsNull() {
		auditnslogaction.Protocolviolations = data.Protocolviolations.ValueString()
	}
	if !data.Serverdomainname.IsNull() {
		auditnslogaction.Serverdomainname = data.Serverdomainname.ValueString()
	}
	if !data.Serverip.IsNull() {
		auditnslogaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		auditnslogaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Sslinterception.IsNull() {
		auditnslogaction.Sslinterception = data.Sslinterception.ValueString()
	}
	if !data.Subscriberlog.IsNull() {
		auditnslogaction.Subscriberlog = data.Subscriberlog.ValueString()
	}
	if !data.Tcp.IsNull() {
		auditnslogaction.Tcp = data.Tcp.ValueString()
	}
	if !data.Timezone.IsNull() {
		auditnslogaction.Timezone = data.Timezone.ValueString()
	}
	if !data.Urlfiltering.IsNull() {
		auditnslogaction.Urlfiltering = data.Urlfiltering.ValueString()
	}
	if !data.Userdefinedauditlog.IsNull() {
		auditnslogaction.Userdefinedauditlog = data.Userdefinedauditlog.ValueString()
	}

	return auditnslogaction
}

func auditnslogactionSetAttrFromGet(ctx context.Context, data *AuditnslogactionResourceModel, getResponseData map[string]interface{}) *AuditnslogactionResourceModel {
	tflog.Debug(ctx, "In auditnslogactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	if val, ok := getResponseData["subscriberlog"]; ok && val != nil {
		data.Subscriberlog = types.StringValue(val.(string))
	} else {
		data.Subscriberlog = types.StringNull()
	}
	if val, ok := getResponseData["tcp"]; ok && val != nil {
		data.Tcp = types.StringValue(val.(string))
	} else {
		data.Tcp = types.StringNull()
	}
	if val, ok := getResponseData["timezone"]; ok && val != nil {
		data.Timezone = types.StringValue(val.(string))
	} else {
		data.Timezone = types.StringNull()
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
