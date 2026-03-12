package aaapreauthenticationaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaapreauthenticationactionResourceModel describes the resource data model.
type AaapreauthenticationactionResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Defaultepagroup         types.String `tfsdk:"defaultepagroup"`
	Deletefiles             types.String `tfsdk:"deletefiles"`
	Killprocess             types.String `tfsdk:"killprocess"`
	Name                    types.String `tfsdk:"name"`
	Preauthenticationaction types.String `tfsdk:"preauthenticationaction"`
}

func (r *AaapreauthenticationactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaapreauthenticationaction resource.",
			},
			"defaultepagroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the EPA check succeeds.",
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool.",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the preauthentication action. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after preauthentication action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my aaa action\" or 'my aaa action').",
			},
			"preauthenticationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow or deny logon after endpoint analysis (EPA) results.",
			},
		},
	}
}

func aaapreauthenticationactionGetThePayloadFromtheConfig(ctx context.Context, data *AaapreauthenticationactionResourceModel) aaa.Aaapreauthenticationaction {
	tflog.Debug(ctx, "In aaapreauthenticationactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaapreauthenticationaction := aaa.Aaapreauthenticationaction{}
	if !data.Defaultepagroup.IsNull() {
		aaapreauthenticationaction.Defaultepagroup = data.Defaultepagroup.ValueString()
	}
	if !data.Deletefiles.IsNull() {
		aaapreauthenticationaction.Deletefiles = data.Deletefiles.ValueString()
	}
	if !data.Killprocess.IsNull() {
		aaapreauthenticationaction.Killprocess = data.Killprocess.ValueString()
	}
	if !data.Name.IsNull() {
		aaapreauthenticationaction.Name = data.Name.ValueString()
	}
	if !data.Preauthenticationaction.IsNull() {
		aaapreauthenticationaction.Preauthenticationaction = data.Preauthenticationaction.ValueString()
	}

	return aaapreauthenticationaction
}

func aaapreauthenticationactionSetAttrFromGet(ctx context.Context, data *AaapreauthenticationactionResourceModel, getResponseData map[string]interface{}) *AaapreauthenticationactionResourceModel {
	tflog.Debug(ctx, "In aaapreauthenticationactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultepagroup"]; ok && val != nil {
		data.Defaultepagroup = types.StringValue(val.(string))
	} else {
		data.Defaultepagroup = types.StringNull()
	}
	if val, ok := getResponseData["deletefiles"]; ok && val != nil {
		data.Deletefiles = types.StringValue(val.(string))
	} else {
		data.Deletefiles = types.StringNull()
	}
	if val, ok := getResponseData["killprocess"]; ok && val != nil {
		data.Killprocess = types.StringValue(val.(string))
	} else {
		data.Killprocess = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["preauthenticationaction"]; ok && val != nil {
		data.Preauthenticationaction = types.StringValue(val.(string))
	} else {
		data.Preauthenticationaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
