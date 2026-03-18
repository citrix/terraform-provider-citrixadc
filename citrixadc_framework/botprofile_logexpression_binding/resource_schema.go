package botprofile_logexpression_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileLogexpressionBindingResourceModel describes the resource data model.
type BotprofileLogexpressionBindingResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	BotBindComment          types.String `tfsdk:"bot_bind_comment"`
	BotLogExpressionEnabled types.String `tfsdk:"bot_log_expression_enabled"`
	BotLogExpressionName    types.String `tfsdk:"bot_log_expression_name"`
	BotLogExpressionValue   types.String `tfsdk:"bot_log_expression_value"`
	Logexpression           types.Bool   `tfsdk:"logexpression"`
	Logmessage              types.String `tfsdk:"logmessage"`
	Name                    types.String `tfsdk:"name"`
}

func (r *BotprofileLogexpressionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_logexpression_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_log_expression_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the log expression binding.",
			},
			"bot_log_expression_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the log expression object.",
			},
			"bot_log_expression_value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression whose result to be logged when violation happened on the bot profile.",
			},
			"logexpression": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log expression binding.",
			},
			"logmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to be logged for this binding.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
		},
	}
}

func botprofile_logexpression_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileLogexpressionBindingResourceModel) bot.Botprofilelogexpressionbinding {
	tflog.Debug(ctx, "In botprofile_logexpression_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_logexpression_binding := bot.Botprofilelogexpressionbinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_logexpression_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotLogExpressionEnabled.IsNull() {
		botprofile_logexpression_binding.Botlogexpressionenabled = data.BotLogExpressionEnabled.ValueString()
	}
	if !data.BotLogExpressionName.IsNull() {
		botprofile_logexpression_binding.Botlogexpressionname = data.BotLogExpressionName.ValueString()
	}
	if !data.BotLogExpressionValue.IsNull() {
		botprofile_logexpression_binding.Botlogexpressionvalue = data.BotLogExpressionValue.ValueString()
	}
	if !data.Logexpression.IsNull() {
		botprofile_logexpression_binding.Logexpression = data.Logexpression.ValueBool()
	}
	if !data.Logmessage.IsNull() {
		botprofile_logexpression_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() {
		botprofile_logexpression_binding.Name = data.Name.ValueString()
	}

	return botprofile_logexpression_binding
}

func botprofile_logexpression_bindingSetAttrFromGet(ctx context.Context, data *BotprofileLogexpressionBindingResourceModel, getResponseData map[string]interface{}) *BotprofileLogexpressionBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_logexpression_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_log_expression_enabled"]; ok && val != nil {
		data.BotLogExpressionEnabled = types.StringValue(val.(string))
	} else {
		data.BotLogExpressionEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_log_expression_name"]; ok && val != nil {
		data.BotLogExpressionName = types.StringValue(val.(string))
	} else {
		data.BotLogExpressionName = types.StringNull()
	}
	if val, ok := getResponseData["bot_log_expression_value"]; ok && val != nil {
		data.BotLogExpressionValue = types.StringValue(val.(string))
	} else {
		data.BotLogExpressionValue = types.StringNull()
	}
	if val, ok := getResponseData["logexpression"]; ok && val != nil {
		data.Logexpression = types.BoolValue(val.(bool))
	} else {
		data.Logexpression = types.BoolNull()
	}
	if val, ok := getResponseData["logmessage"]; ok && val != nil {
		data.Logmessage = types.StringValue(val.(string))
	} else {
		data.Logmessage = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_log_expression_name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.BotLogExpressionName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
