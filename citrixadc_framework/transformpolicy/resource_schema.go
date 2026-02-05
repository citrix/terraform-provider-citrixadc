package transformpolicy

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

// TransformpolicyResourceModel describes the resource data model.
type TransformpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Profilename types.String `tfsdk:"profilename"`
	Rule        types.String `tfsdk:"rule"`
}

func (r *TransformpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformpolicy resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log server to use to log connections that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL Transformation policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policy or my transform policy).",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policy or my transform policy).",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the URL Transformation profile to use to transform requests and responses that match the policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, or name of a named expression, against which to evaluate traffic.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the \\ character. \n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func transformpolicyGetThePayloadFromtheConfig(ctx context.Context, data *TransformpolicyResourceModel) transform.Transformpolicy {
	tflog.Debug(ctx, "In transformpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	transformpolicy := transform.Transformpolicy{}
	if !data.Comment.IsNull() {
		transformpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		transformpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		transformpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		transformpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Profilename.IsNull() {
		transformpolicy.Profilename = data.Profilename.ValueString()
	}
	if !data.Rule.IsNull() {
		transformpolicy.Rule = data.Rule.ValueString()
	}

	return transformpolicy
}

func transformpolicySetAttrFromGet(ctx context.Context, data *TransformpolicyResourceModel, getResponseData map[string]interface{}) *TransformpolicyResourceModel {
	tflog.Debug(ctx, "In transformpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
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
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
