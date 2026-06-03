package videooptimizationpacingpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// VideooptimizationpacingpolicylabelResourceModel describes the resource data model.
type VideooptimizationpacingpolicylabelResourceModel struct {
	Id types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *VideooptimizationpacingpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationpacingpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this videooptimization pacing policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Video optimization pacing policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (\n.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the videooptimization pacing policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my videooptimization pacing policy label\" or my videooptimization pacing policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the videooptimization pacing policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (\n-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.",
			},
		},
	}
}

func videooptimizationpacingpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationpacingpolicylabelResourceModel) videooptimization.Videooptimizationpacingpolicylabel {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	videooptimizationpacingpolicylabel := videooptimization.Videooptimizationpacingpolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		videooptimizationpacingpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		videooptimizationpacingpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname belongs to the `rename` action only and is not accepted by the add
	// payload; it is intentionally excluded from the create body.
	if !data.Policylabeltype.IsNull() && !data.Policylabeltype.IsUnknown() {
		videooptimizationpacingpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return videooptimizationpacingpolicylabel
}

func videooptimizationpacingpolicylabelSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingpolicylabelResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingpolicylabelResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabelSetAttrFromGet Function")

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
	// newname is a rename-only write parameter that the NITRO GET response never
	// echoes back. Preserve the existing plan/state value instead of nulling it,
	// otherwise a configured value would be wiped on every Read (Pattern 7).
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	} else {
		data.Policylabeltype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}