package transformpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformpolicylabelResourceModel describes the resource data model.
type TransformpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *TransformpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"policylabeltype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).",
			},
		},
	}
}

func transformpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *TransformpolicylabelResourceModel) transform.Transformpolicylabel {
	tflog.Debug(ctx, "In transformpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	transformpolicylabel := transform.Transformpolicylabel{}
	if !data.Labelname.IsNull() {
		transformpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		transformpolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Policylabeltype.IsNull() {
		transformpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return transformpolicylabel
}

func transformpolicylabelSetAttrFromGet(ctx context.Context, data *TransformpolicylabelResourceModel, getResponseData map[string]interface{}) *TransformpolicylabelResourceModel {
	tflog.Debug(ctx, "In transformpolicylabelSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	} else {
		data.Policylabeltype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
