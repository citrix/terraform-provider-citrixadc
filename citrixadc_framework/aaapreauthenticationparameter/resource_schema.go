package aaapreauthenticationparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaapreauthenticationparameterResourceModel describes the resource data model.
type AaapreauthenticationparameterResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Deletefiles             types.String `tfsdk:"deletefiles"`
	Killprocess             types.String `tfsdk:"killprocess"`
	Preauthenticationaction types.String `tfsdk:"preauthenticationaction"`
	Rule                    types.String `tfsdk:"rule"`
}

func (r *AaapreauthenticationparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaapreauthenticationparameter resource.",
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) to and name(s) of the files to be deleted by the EPA tool, as a string of between 1 and 1023 characters.",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the EPA tool.",
			},
			"preauthenticationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny or allow login on the basis of end point analysis results.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, to be evaluated by the EPA tool.",
			},
		},
	}
}

func aaapreauthenticationparameterGetThePayloadFromtheConfig(ctx context.Context, data *AaapreauthenticationparameterResourceModel) aaa.Aaapreauthenticationparameter {
	tflog.Debug(ctx, "In aaapreauthenticationparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaapreauthenticationparameter := aaa.Aaapreauthenticationparameter{}
	if !data.Deletefiles.IsNull() {
		aaapreauthenticationparameter.Deletefiles = data.Deletefiles.ValueString()
	}
	if !data.Killprocess.IsNull() {
		aaapreauthenticationparameter.Killprocess = data.Killprocess.ValueString()
	}
	if !data.Preauthenticationaction.IsNull() {
		aaapreauthenticationparameter.Preauthenticationaction = data.Preauthenticationaction.ValueString()
	}
	if !data.Rule.IsNull() {
		aaapreauthenticationparameter.Rule = data.Rule.ValueString()
	}

	return aaapreauthenticationparameter
}

func aaapreauthenticationparameterSetAttrFromGet(ctx context.Context, data *AaapreauthenticationparameterResourceModel, getResponseData map[string]interface{}) *AaapreauthenticationparameterResourceModel {
	tflog.Debug(ctx, "In aaapreauthenticationparameterSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["preauthenticationaction"]; ok && val != nil {
		data.Preauthenticationaction = types.StringValue(val.(string))
	} else {
		data.Preauthenticationaction = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaapreauthenticationparameter-config")

	return data
}
