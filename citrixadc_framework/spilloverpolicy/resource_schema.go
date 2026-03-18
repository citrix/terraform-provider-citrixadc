package spilloverpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SpilloverpolicyResourceModel describes the resource data model.
type SpilloverpolicyResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Rule    types.String `tfsdk:"rule"`
}

func (r *SpilloverpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the spilloverpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action for the spillover policy. Action is created using add spillover action command",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the spillover policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the spillover policy.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression to be used by the spillover policy.",
			},
		},
	}
}

func spilloverpolicyGetThePayloadFromtheConfig(ctx context.Context, data *SpilloverpolicyResourceModel) spillover.Spilloverpolicy {
	tflog.Debug(ctx, "In spilloverpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	spilloverpolicy := spillover.Spilloverpolicy{}
	if !data.Action.IsNull() {
		spilloverpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		spilloverpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		spilloverpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		spilloverpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		spilloverpolicy.Rule = data.Rule.ValueString()
	}

	return spilloverpolicy
}

func spilloverpolicySetAttrFromGet(ctx context.Context, data *SpilloverpolicyResourceModel, getResponseData map[string]interface{}) *SpilloverpolicyResourceModel {
	tflog.Debug(ctx, "In spilloverpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
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
