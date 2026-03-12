package auditnslogparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuditnslogparamsResourceModel describes the resource data model.
type AuditnslogparamsResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Tcp                  types.String `tfsdk:"tcp"`
	Timezone             types.String `tfsdk:"timezone"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`
}

func (r *AuditnslogparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditnslogparams resource.",
			},
			"acl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure auditing to log access control list (ACL) messages.",
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
			"logfacility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Facility value, as defined in RFC 3164, assigned to the log message.\nLog facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.",
			},
			"loglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Types of information to be logged.\nAvailable settings function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log the LSN messages",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
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
				Description: "Configure auditing to log TCP messages.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nSupported settings are:\n* GMT_TIME - Coordinated Universal Time.\n* LOCAL_TIME - Use the server's timezone setting.",
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

func auditnslogparamsGetThePayloadFromtheConfig(ctx context.Context, data *AuditnslogparamsResourceModel) audit.Auditnslogparams {
	tflog.Debug(ctx, "In auditnslogparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditnslogparams := audit.Auditnslogparams{}
	if !data.Acl.IsNull() {
		auditnslogparams.Acl = data.Acl.ValueString()
	}
	if !data.Alg.IsNull() {
		auditnslogparams.Alg = data.Alg.ValueString()
	}
	if !data.Appflowexport.IsNull() {
		auditnslogparams.Appflowexport = data.Appflowexport.ValueString()
	}
	if !data.Contentinspectionlog.IsNull() {
		auditnslogparams.Contentinspectionlog = data.Contentinspectionlog.ValueString()
	}
	if !data.Dateformat.IsNull() {
		auditnslogparams.Dateformat = data.Dateformat.ValueString()
	}
	if !data.Logfacility.IsNull() {
		auditnslogparams.Logfacility = data.Logfacility.ValueString()
	}
	if !data.Lsn.IsNull() {
		auditnslogparams.Lsn = data.Lsn.ValueString()
	}
	if !data.Protocolviolations.IsNull() {
		auditnslogparams.Protocolviolations = data.Protocolviolations.ValueString()
	}
	if !data.Serverip.IsNull() {
		auditnslogparams.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		auditnslogparams.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Sslinterception.IsNull() {
		auditnslogparams.Sslinterception = data.Sslinterception.ValueString()
	}
	if !data.Subscriberlog.IsNull() {
		auditnslogparams.Subscriberlog = data.Subscriberlog.ValueString()
	}
	if !data.Tcp.IsNull() {
		auditnslogparams.Tcp = data.Tcp.ValueString()
	}
	if !data.Timezone.IsNull() {
		auditnslogparams.Timezone = data.Timezone.ValueString()
	}
	if !data.Urlfiltering.IsNull() {
		auditnslogparams.Urlfiltering = data.Urlfiltering.ValueString()
	}
	if !data.Userdefinedauditlog.IsNull() {
		auditnslogparams.Userdefinedauditlog = data.Userdefinedauditlog.ValueString()
	}

	return auditnslogparams
}

func auditnslogparamsSetAttrFromGet(ctx context.Context, data *AuditnslogparamsResourceModel, getResponseData map[string]interface{}) *AuditnslogparamsResourceModel {
	tflog.Debug(ctx, "In auditnslogparamsSetAttrFromGet Function")

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
	if val, ok := getResponseData["protocolviolations"]; ok && val != nil {
		data.Protocolviolations = types.StringValue(val.(string))
	} else {
		data.Protocolviolations = types.StringNull()
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
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("auditnslogparams-config")

	return data
}
