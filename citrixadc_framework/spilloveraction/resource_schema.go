package spilloveraction

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

// SpilloveractionResourceModel describes the resource data model.
type SpilloveractionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Action  types.String `tfsdk:"action"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
}

func (r *SpilloveractionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the spilloveraction resource.",
			},
			"action": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Spillover action. Currently only type SPILLOVER is supported",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the spillover action.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters. \nChoose a name that can be correlated with the function that the action performs. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
		},
	}
}

func spilloveractionGetThePayloadFromtheConfig(ctx context.Context, data *SpilloveractionResourceModel) spillover.Spilloveraction {
	tflog.Debug(ctx, "In spilloveractionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	spilloveraction := spillover.Spilloveraction{}
	if !data.Action.IsNull() {
		spilloveraction.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		spilloveraction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		spilloveraction.Newname = data.Newname.ValueString()
	}

	return spilloveraction
}

func spilloveractionSetAttrFromGet(ctx context.Context, data *SpilloveractionResourceModel, getResponseData map[string]interface{}) *SpilloveractionResourceModel {
	tflog.Debug(ctx, "In spilloveractionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
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
