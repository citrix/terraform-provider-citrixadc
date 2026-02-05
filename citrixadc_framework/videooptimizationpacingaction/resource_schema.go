package videooptimizationpacingaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VideooptimizationpacingactionResourceModel describes the resource data model.
type VideooptimizationpacingactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rate    types.Int64  `tfsdk:"rate"`
}

func (r *VideooptimizationpacingactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationpacingaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the video optimization pacing action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the videooptimization pacing action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rate": schema.Int64Attribute{
				Required:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "ABR Video Optimization Pacing Rate (in Kbps)",
			},
		},
	}
}

func videooptimizationpacingactionGetThePayloadFromtheConfig(ctx context.Context, data *VideooptimizationpacingactionResourceModel) videooptimization.Videooptimizationpacingaction {
	tflog.Debug(ctx, "In videooptimizationpacingactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	videooptimizationpacingaction := videooptimization.Videooptimizationpacingaction{}
	if !data.Comment.IsNull() {
		videooptimizationpacingaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		videooptimizationpacingaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		videooptimizationpacingaction.Newname = data.Newname.ValueString()
	}
	if !data.Rate.IsNull() {
		videooptimizationpacingaction.Rate = utils.IntPtr(int(data.Rate.ValueInt64()))
	}

	return videooptimizationpacingaction
}

func videooptimizationpacingactionSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingactionResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingactionResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["rate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rate = types.Int64Value(intVal)
		}
	} else {
		data.Rate = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
