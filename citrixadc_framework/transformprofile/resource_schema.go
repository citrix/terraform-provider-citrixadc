package transformprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformprofileResourceModel describes the resource data model.
type TransformprofileResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Comment                   types.String `tfsdk:"comment"`
	Name                      types.String `tfsdk:"name"`
	Onlytransformabsurlinbody types.String `tfsdk:"onlytransformabsurlinbody"`
	Type                      types.String `tfsdk:"type"`
}

func (r *TransformprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformprofile resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform profile or my transform profile).",
			},
			"onlytransformabsurlinbody": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In the HTTP body, transform only absolute URLs. Relative URLs are ignored.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of transformation. Always URL for URL Transformation profiles.",
			},
		},
	}
}

func transformprofileGetThePayloadFromtheConfig(ctx context.Context, data *TransformprofileResourceModel) transform.Transformprofile {
	tflog.Debug(ctx, "In transformprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	transformprofile := transform.Transformprofile{}
	if !data.Comment.IsNull() {
		transformprofile.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		transformprofile.Name = data.Name.ValueString()
	}
	if !data.Onlytransformabsurlinbody.IsNull() {
		transformprofile.Onlytransformabsurlinbody = data.Onlytransformabsurlinbody.ValueString()
	}
	if !data.Type.IsNull() {
		transformprofile.Type = data.Type.ValueString()
	}

	return transformprofile
}

func transformprofileSetAttrFromGet(ctx context.Context, data *TransformprofileResourceModel, getResponseData map[string]interface{}) *TransformprofileResourceModel {
	tflog.Debug(ctx, "In transformprofileSetAttrFromGet Function")

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
	if val, ok := getResponseData["onlytransformabsurlinbody"]; ok && val != nil {
		data.Onlytransformabsurlinbody = types.StringValue(val.(string))
	} else {
		data.Onlytransformabsurlinbody = types.StringNull()
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
