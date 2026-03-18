package contentinspectioncallout

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ContentinspectioncalloutResourceModel describes the resource data model.
type ContentinspectioncalloutResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
	Profilename types.String `tfsdk:"profilename"`
	Resultexpr  types.String `tfsdk:"resultexpr"`
	Returntype  types.String `tfsdk:"returntype"`
	Serverip    types.String `tfsdk:"serverip"`
	Servername  types.String `tfsdk:"servername"`
	Serverport  types.Int64  `tfsdk:"serverport"`
	Type        types.String `tfsdk:"type"`
}

func (r *ContentinspectioncalloutResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectioncallout resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this Content Inspection callout.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Content Inspection callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or callout.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Content Inspection profile. The type of the configured profile must match the type specified using -type argument.",
			},
			"resultexpr": schema.StringAttribute{
				Required:    true,
				Description: "Expression that extracts the callout results from the response sent by the CI callout agent. Must be a response based expression, that is, it must begin with ICAP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression, as in the following example: icap.res.header(\"ISTag\")",
			},
			"returntype": schema.StringAttribute{
				Required:    true,
				Description: "Type of data that the target callout agent returns in response to the callout.\nAvailable settings function as follows:\n* TEXT - Treat the returned value as a text string.\n* NUM - Treat the returned value as a number.\n* BOOL - Treat the returned value as a Boolean value.\nNote: You cannot change the return type after it is set.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of Content Inspection server. Mutually exclusive with the server name parameter.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing or content switching virtual server or service to which the Content Inspection request is issued. Mutually exclusive with server IP address and port parameters. The service type must be TCP or SSL_TCP. If there are vservers and services with the same name, then vserver is selected.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1344),
				Description: "Port of the Content Inspection server.",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the Content Inspection callout. It must be one of the following:\n* ICAP - Sends ICAP request to the configured ICAP server.",
			},
		},
	}
}

func contentinspectioncalloutGetThePayloadFromtheConfig(ctx context.Context, data *ContentinspectioncalloutResourceModel) contentinspection.Contentinspectioncallout {
	tflog.Debug(ctx, "In contentinspectioncalloutGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	contentinspectioncallout := contentinspection.Contentinspectioncallout{}
	if !data.Comment.IsNull() {
		contentinspectioncallout.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		contentinspectioncallout.Name = data.Name.ValueString()
	}
	if !data.Profilename.IsNull() {
		contentinspectioncallout.Profilename = data.Profilename.ValueString()
	}
	if !data.Resultexpr.IsNull() {
		contentinspectioncallout.Resultexpr = data.Resultexpr.ValueString()
	}
	if !data.Returntype.IsNull() {
		contentinspectioncallout.Returntype = data.Returntype.ValueString()
	}
	if !data.Serverip.IsNull() {
		contentinspectioncallout.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() {
		contentinspectioncallout.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() {
		contentinspectioncallout.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Type.IsNull() {
		contentinspectioncallout.Type = data.Type.ValueString()
	}

	return contentinspectioncallout
}

func contentinspectioncalloutSetAttrFromGet(ctx context.Context, data *ContentinspectioncalloutResourceModel, getResponseData map[string]interface{}) *ContentinspectioncalloutResourceModel {
	tflog.Debug(ctx, "In contentinspectioncalloutSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["resultexpr"]; ok && val != nil {
		data.Resultexpr = types.StringValue(val.(string))
	} else {
		data.Resultexpr = types.StringNull()
	}
	if val, ok := getResponseData["returntype"]; ok && val != nil {
		data.Returntype = types.StringValue(val.(string))
	} else {
		data.Returntype = types.StringNull()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
