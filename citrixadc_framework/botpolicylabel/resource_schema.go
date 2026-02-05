package botpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotpolicylabelResourceModel describes the resource data model.
type BotpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *BotpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this bot policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy label\" or my responder policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
		},
	}
}

func botpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *BotpolicylabelResourceModel) bot.Botpolicylabel {
	tflog.Debug(ctx, "In botpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botpolicylabel := bot.Botpolicylabel{}
	if !data.Comment.IsNull() {
		botpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() {
		botpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		botpolicylabel.Newname = data.Newname.ValueString()
	}

	return botpolicylabel
}

func botpolicylabelSetAttrFromGet(ctx context.Context, data *BotpolicylabelResourceModel, getResponseData map[string]interface{}) *BotpolicylabelResourceModel {
	tflog.Debug(ctx, "In botpolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
