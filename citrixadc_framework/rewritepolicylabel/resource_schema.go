package rewritepolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicylabelResourceModel describes the resource data model.
type RewritepolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Transform types.String `tfsdk:"transform"`
}

func (r *RewritepolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewritepolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this rewrite policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the rewrite policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rewrite policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy label\" or 'my rewrite policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the rewrite policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"transform": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:\n* http_req - HTTP requests\n* http_res - HTTP responses\n* othertcp_req - Non-HTTP TCP requests\n* othertcp_res - Non-HTTP TCP responses\n* url - URLs\n* text - Text strings\n* clientless_vpn_req - Citrix ADC clientless VPN requests\n* clientless_vpn_res - Citrix ADC clientless VPN responses\n* sipudp_req - SIP requests\n* sipudp_res - SIP responses\n* diameter_req - DIAMETER requests\n* diameter_res - DIAMETER responses\n* radius_req - RADIUS requests\n* radius_res - RADIUS responses\n* dns_req - DNS requests\n* dns_res - DNS responses\n* mqtt_req - MQTT requests\n* mqtt_res - MQTT responses",
			},
		},
	}
}

func rewritepolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *RewritepolicylabelResourceModel) rewrite.Rewritepolicylabel {
	tflog.Debug(ctx, "In rewritepolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rewritepolicylabel := rewrite.Rewritepolicylabel{}
	if !data.Comment.IsNull() {
		rewritepolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() {
		rewritepolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		rewritepolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Transform.IsNull() {
		rewritepolicylabel.Transform = data.Transform.ValueString()
	}

	return rewritepolicylabel
}

func rewritepolicylabelSetAttrFromGet(ctx context.Context, data *RewritepolicylabelResourceModel, getResponseData map[string]interface{}) *RewritepolicylabelResourceModel {
	tflog.Debug(ctx, "In rewritepolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
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
	if val, ok := getResponseData["transform"]; ok && val != nil {
		data.Transform = types.StringValue(val.(string))
	} else {
		data.Transform = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
