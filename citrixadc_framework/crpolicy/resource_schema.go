package crpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrpolicyResourceModel describes the resource data model.
type CrpolicyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Logaction  types.String `tfsdk:"logaction"`
	Newname    types.String `tfsdk:"newname"`
	Policyname types.String `tfsdk:"policyname"`
	Rule       types.String `tfsdk:"rule"`
}

func (r *CrpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crpolicy resource.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in cache redirection action: CACHE/ORIGIN.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the cache redirection policy",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the content switching policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func crpolicyGetThePayloadFromtheConfig(ctx context.Context, data *CrpolicyResourceModel) cr.Crpolicy {
	tflog.Debug(ctx, "In crpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	crpolicy := cr.Crpolicy{}
	if !data.Action.IsNull() {
		crpolicy.Action = data.Action.ValueString()
	}
	if !data.Logaction.IsNull() {
		crpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Newname.IsNull() {
		crpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Policyname.IsNull() {
		crpolicy.Policyname = data.Policyname.ValueString()
	}
	if !data.Rule.IsNull() {
		crpolicy.Rule = data.Rule.ValueString()
	}

	return crpolicy
}

func crpolicySetAttrFromGet(ctx context.Context, data *CrpolicyResourceModel, getResponseData map[string]interface{}) *CrpolicyResourceModel {
	tflog.Debug(ctx, "In crpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
