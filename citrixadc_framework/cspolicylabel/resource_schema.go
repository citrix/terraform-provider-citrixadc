package cspolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CspolicylabelResourceModel describes the resource data model.
type CspolicylabelResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Cspolicylabeltype types.String `tfsdk:"cspolicylabeltype"`
	Labelname         types.String `tfsdk:"labelname"`
	Newname           types.String `tfsdk:"newname"`
}

func (r *CspolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cspolicylabel resource.",
			},
			"cspolicylabeltype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol supported by the policy label. All policies bound to the policy label must either match the specified protocol or be a subtype of that protocol. Available settings function as follows:\n* HTTP - Supports policies that process HTTP traffic. Used to access unencrypted Web sites. (The default.)\n* SSL - Supports policies that process HTTPS/SSL encrypted traffic. Used to access encrypted Web sites.\n* TCP - Supports policies that process any type of TCP traffic, including HTTP.\n* SSL_TCP - Supports policies that process SSL-encrypted TCP traffic, including SSL.\n* UDP - Supports policies that process any type of UDP-based traffic, including DNS.\n* DNS - Supports policies that process DNS traffic.\n* ANY - Supports all types of policies except HTTP, SSL, and TCP.\n* SIP_UDP - Supports policies that process UDP based Session Initiation Protocol (SIP) traffic. SIP initiates, manages, and terminates multimedia communications sessions, and has emerged as the standard for Internet telephony (VoIP).\n* RTSP - Supports policies that process Real Time Streaming Protocol (RTSP) traffic. RTSP provides delivery of multimedia and other streaming data, such as audio, video, and other types of streamed media.\n* RADIUS - Supports policies that process Remote Authentication Dial In User Service (RADIUS) traffic. RADIUS supports combined authentication, authorization, and auditing services for network management.\n* MYSQL - Supports policies that process MYSQL traffic.\n* MSSQL - Supports policies that process Microsoft SQL traffic.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe label name must be unique within the list of policy labels for content switching.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policylabel\" or 'my policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the content switching policylabel.",
			},
		},
	}
}

func cspolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *CspolicylabelResourceModel) cs.Cspolicylabel {
	tflog.Debug(ctx, "In cspolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cspolicylabel := cs.Cspolicylabel{}
	if !data.Cspolicylabeltype.IsNull() {
		cspolicylabel.Cspolicylabeltype = data.Cspolicylabeltype.ValueString()
	}
	if !data.Labelname.IsNull() {
		cspolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		cspolicylabel.Newname = data.Newname.ValueString()
	}

	return cspolicylabel
}

func cspolicylabelSetAttrFromGet(ctx context.Context, data *CspolicylabelResourceModel, getResponseData map[string]interface{}) *CspolicylabelResourceModel {
	tflog.Debug(ctx, "In cspolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cspolicylabeltype"]; ok && val != nil {
		data.Cspolicylabeltype = types.StringValue(val.(string))
	} else {
		data.Cspolicylabeltype = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
