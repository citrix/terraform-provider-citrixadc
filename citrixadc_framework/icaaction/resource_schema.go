package icaaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaactionResourceModel describes the resource data model.
type IcaactionResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Accessprofilename  types.String `tfsdk:"accessprofilename"`
	Latencyprofilename types.String `tfsdk:"latencyprofilename"`
	Name               types.String `tfsdk:"name"`
	Newname            types.String `tfsdk:"newname"`
}

func (r *IcaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icaaction resource.",
			},
			"accessprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica accessprofile to be associated with this action.",
			},
			"latencyprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica latencyprofile to be associated with this action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica action\" or 'my ica action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the ICA action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#),period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks ( for example, \"my ica action\" or 'my ica action').",
			},
		},
	}
}

func icaactionGetThePayloadFromtheConfig(ctx context.Context, data *IcaactionResourceModel) ica.Icaaction {
	tflog.Debug(ctx, "In icaactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	icaaction := ica.Icaaction{}
	if !data.Accessprofilename.IsNull() {
		icaaction.Accessprofilename = data.Accessprofilename.ValueString()
	}
	if !data.Latencyprofilename.IsNull() {
		icaaction.Latencyprofilename = data.Latencyprofilename.ValueString()
	}
	if !data.Name.IsNull() {
		icaaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		icaaction.Newname = data.Newname.ValueString()
	}

	return icaaction
}

func icaactionSetAttrFromGet(ctx context.Context, data *IcaactionResourceModel, getResponseData map[string]interface{}) *IcaactionResourceModel {
	tflog.Debug(ctx, "In icaactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accessprofilename"]; ok && val != nil {
		data.Accessprofilename = types.StringValue(val.(string))
	} else {
		data.Accessprofilename = types.StringNull()
	}
	if val, ok := getResponseData["latencyprofilename"]; ok && val != nil {
		data.Latencyprofilename = types.StringValue(val.(string))
	} else {
		data.Latencyprofilename = types.StringNull()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
