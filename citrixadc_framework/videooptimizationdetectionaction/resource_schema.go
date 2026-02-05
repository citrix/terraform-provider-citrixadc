package videooptimizationdetectionaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationdetectionactionResourceModel describes the resource data model.
type VideooptimizationdetectionactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Type    types.String `tfsdk:"type"`
}

func (r *VideooptimizationdetectionactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationdetectionaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the videooptimization detection action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of video optimization action. Available settings function as follows:\n* clear_text_pd - Cleartext PD type is detected.\n* clear_text_abr - Cleartext ABR is detected.\n* encrypted_abr - Encrypted ABR is detected.\n* trigger_enc_abr - Possible encrypted ABR is detected.\n* trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.",
			},
		},
	}
}

func videooptimizationdetectionactionGetThePayloadFromtheConfig(ctx context.Context, data *VideooptimizationdetectionactionResourceModel) videooptimization.Videooptimizationdetectionaction {
	tflog.Debug(ctx, "In videooptimizationdetectionactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	videooptimizationdetectionaction := videooptimization.Videooptimizationdetectionaction{}
	if !data.Comment.IsNull() {
		videooptimizationdetectionaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		videooptimizationdetectionaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		videooptimizationdetectionaction.Newname = data.Newname.ValueString()
	}
	if !data.Type.IsNull() {
		videooptimizationdetectionaction.Type = data.Type.ValueString()
	}

	return videooptimizationdetectionaction
}

func videooptimizationdetectionactionSetAttrFromGet(ctx context.Context, data *VideooptimizationdetectionactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationdetectionactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationdetectionactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
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
