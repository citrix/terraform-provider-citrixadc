package autoscalepolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AutoscalepolicyResourceModel describes the resource data model.
type AutoscalepolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AutoscalepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the autoscalepolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "The autoscale profile associated with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this autoscale policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the autoscale policy",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the autoscale policy.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the autoscale policy.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "The rule associated with the policy.",
			},
		},
	}
}

func autoscalepolicyGetThePayloadFromtheConfig(ctx context.Context, data *AutoscalepolicyResourceModel) autoscale.Autoscalepolicy {
	tflog.Debug(ctx, "In autoscalepolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	autoscalepolicy := autoscale.Autoscalepolicy{}
	if !data.Action.IsNull() {
		autoscalepolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		autoscalepolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		autoscalepolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		autoscalepolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		autoscalepolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		autoscalepolicy.Rule = data.Rule.ValueString()
	}

	return autoscalepolicy
}

func autoscalepolicySetAttrFromGet(ctx context.Context, data *AutoscalepolicyResourceModel, getResponseData map[string]interface{}) *AutoscalepolicyResourceModel {
	tflog.Debug(ctx, "In autoscalepolicySetAttrFromGet Function")

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
