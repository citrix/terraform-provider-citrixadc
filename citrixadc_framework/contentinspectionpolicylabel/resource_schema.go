package contentinspectionpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectionpolicylabelResourceModel describes the resource data model.
type ContentinspectionpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
	Type      types.String `tfsdk:"type"`
}

func (r *ContentinspectionpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this contentInspection policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the contentInspection policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my contentInspection policy label\" or 'my contentInspection policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the contentInspection policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of packets (request or response packets) against which to match the policies bound to this policy label.",
			},
		},
	}
}

func contentinspectionpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *ContentinspectionpolicylabelResourceModel) contentinspection.Contentinspectionpolicylabel {
	tflog.Debug(ctx, "In contentinspectionpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	contentinspectionpolicylabel := contentinspection.Contentinspectionpolicylabel{}
	if !data.Comment.IsNull() {
		contentinspectionpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() {
		contentinspectionpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		contentinspectionpolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Type.IsNull() {
		contentinspectionpolicylabel.Type = data.Type.ValueString()
	}

	return contentinspectionpolicylabel
}

func contentinspectionpolicylabelSetAttrFromGet(ctx context.Context, data *ContentinspectionpolicylabelResourceModel, getResponseData map[string]interface{}) *ContentinspectionpolicylabelResourceModel {
	tflog.Debug(ctx, "In contentinspectionpolicylabelSetAttrFromGet Function")

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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
