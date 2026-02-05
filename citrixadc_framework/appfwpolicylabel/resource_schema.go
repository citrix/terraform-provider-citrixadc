package appfwpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwpolicylabelResourceModel describes the resource data model.
type AppfwpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *AppfwpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the policy label is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the application firewall policylabel.",
			},
			"policylabeltype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of transformations allowed by the policies bound to the label. Always http_req for application firewall policy labels.",
			},
		},
	}
}

func appfwpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *AppfwpolicylabelResourceModel) appfw.Appfwpolicylabel {
	tflog.Debug(ctx, "In appfwpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwpolicylabel := appfw.Appfwpolicylabel{}
	if !data.Labelname.IsNull() {
		appfwpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		appfwpolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Policylabeltype.IsNull() {
		appfwpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return appfwpolicylabel
}

func appfwpolicylabelSetAttrFromGet(ctx context.Context, data *AppfwpolicylabelResourceModel, getResponseData map[string]interface{}) *AppfwpolicylabelResourceModel {
	tflog.Debug(ctx, "In appfwpolicylabelSetAttrFromGet Function")

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
