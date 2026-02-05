package contentinspectionaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ContentinspectionactionResourceModel describes the resource data model.
type ContentinspectionactionResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Icapprofilename types.String `tfsdk:"icapprofilename"`
	Ifserverdown    types.String `tfsdk:"ifserverdown"`
	Name            types.String `tfsdk:"name"`
	Serverip        types.String `tfsdk:"serverip"`
	Servername      types.String `tfsdk:"servername"`
	Serverport      types.Int64  `tfsdk:"serverport"`
	Type            types.String `tfsdk:"type"`
}

func (r *ContentinspectionactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionaction resource.",
			},
			"icapprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ICAP profile to be attached to the contentInspection action.",
			},
			"ifserverdown": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RESET"),
				Description: "Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.\n* CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remoteService",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver or service",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1344),
				Description: "Port of remoteService",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of operation this action is going to perform. following actions are available to configure:\n* ICAP - forward the incoming request or response to an ICAP server for modification.\n* INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.\n* MIRROR - Forwards cloned packets for Intrusion Detection.\n* NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.\n* NSTRACE - capture current and further incoming packets on this transaction.",
			},
		},
	}
}

func contentinspectionactionGetThePayloadFromtheConfig(ctx context.Context, data *ContentinspectionactionResourceModel) contentinspection.Contentinspectionaction {
	tflog.Debug(ctx, "In contentinspectionactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	contentinspectionaction := contentinspection.Contentinspectionaction{}
	if !data.Icapprofilename.IsNull() {
		contentinspectionaction.Icapprofilename = data.Icapprofilename.ValueString()
	}
	if !data.Ifserverdown.IsNull() {
		contentinspectionaction.Ifserverdown = data.Ifserverdown.ValueString()
	}
	if !data.Name.IsNull() {
		contentinspectionaction.Name = data.Name.ValueString()
	}
	if !data.Serverip.IsNull() {
		contentinspectionaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() {
		contentinspectionaction.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() {
		contentinspectionaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Type.IsNull() {
		contentinspectionaction.Type = data.Type.ValueString()
	}

	return contentinspectionaction
}

func contentinspectionactionSetAttrFromGet(ctx context.Context, data *ContentinspectionactionResourceModel, getResponseData map[string]interface{}) *ContentinspectionactionResourceModel {
	tflog.Debug(ctx, "In contentinspectionactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["icapprofilename"]; ok && val != nil {
		data.Icapprofilename = types.StringValue(val.(string))
	} else {
		data.Icapprofilename = types.StringNull()
	}
	if val, ok := getResponseData["ifserverdown"]; ok && val != nil {
		data.Ifserverdown = types.StringValue(val.(string))
	} else {
		data.Ifserverdown = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
