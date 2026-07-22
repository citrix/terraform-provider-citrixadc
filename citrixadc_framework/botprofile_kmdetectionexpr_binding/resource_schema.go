package botprofile_kmdetectionexpr_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileKmdetectionexprBindingResourceModel describes the resource data model.
type BotprofileKmdetectionexprBindingResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	BotBindComment        types.String `tfsdk:"bot_bind_comment"`
	BotKmDetectionEnabled types.String `tfsdk:"bot_km_detection_enabled"`
	BotKmExpressionName   types.String `tfsdk:"bot_km_expression_name"`
	BotKmExpressionValue  types.String `tfsdk:"bot_km_expression_value"`
	Kmdetectionexpr       types.Bool   `tfsdk:"kmdetectionexpr"`
	Logmessage            types.String `tfsdk:"logmessage"`
	Name                  types.String `tfsdk:"name"`
}

func (r *BotprofileKmdetectionexprBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_kmdetectionexpr_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about this binding.",
			},
			"bot_km_detection_enabled": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enable or disable the keyboard-mouse based binding.",
			},
			"bot_km_expression_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the keyboard-mouse expression object.",
			},
			"bot_km_expression_value": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "JavaScript file for keyboard-mouse detection, would be inserted if the result of the expression is true.",
			},
			"kmdetectionexpr": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Keyboard-mouse based detection binding. For each name, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, then needs to bind again with new vlaues. Maximum 30 bindings can be configured per profile.",
			},
			"logmessage": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Message to be logged for this binding.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
		},
	}
}

func botprofile_kmdetectionexpr_bindingGetThePayloadFromthePlan(ctx context.Context, data *BotprofileKmdetectionexprBindingResourceModel) bot.Botprofilekmdetectionexprbinding {
	tflog.Debug(ctx, "In botprofile_kmdetectionexpr_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botprofile_kmdetectionexpr_binding := bot.Botprofilekmdetectionexprbinding{}
	if !data.BotBindComment.IsNull() && !data.BotBindComment.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotKmDetectionEnabled.IsNull() && !data.BotKmDetectionEnabled.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Botkmdetectionenabled = data.BotKmDetectionEnabled.ValueString()
	}
	if !data.BotKmExpressionName.IsNull() && !data.BotKmExpressionName.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Botkmexpressionname = data.BotKmExpressionName.ValueString()
	}
	if !data.BotKmExpressionValue.IsNull() && !data.BotKmExpressionValue.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Botkmexpressionvalue = data.BotKmExpressionValue.ValueString()
	}
	if !data.Kmdetectionexpr.IsNull() && !data.Kmdetectionexpr.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Kmdetectionexpr = data.Kmdetectionexpr.ValueBool()
	}
	if !data.Logmessage.IsNull() && !data.Logmessage.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botprofile_kmdetectionexpr_binding.Name = data.Name.ValueString()
	}

	return botprofile_kmdetectionexpr_binding
}

func botprofile_kmdetectionexpr_bindingSetAttrFromGet(ctx context.Context, data *BotprofileKmdetectionexprBindingResourceModel, getResponseData map[string]interface{}) *BotprofileKmdetectionexprBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_kmdetectionexpr_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_km_detection_enabled"]; ok && val != nil {
		data.BotKmDetectionEnabled = types.StringValue(val.(string))
	} else {
		data.BotKmDetectionEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_km_expression_name"]; ok && val != nil {
		data.BotKmExpressionName = types.StringValue(val.(string))
	} else {
		data.BotKmExpressionName = types.StringNull()
	}
	if val, ok := getResponseData["bot_km_expression_value"]; ok && val != nil {
		data.BotKmExpressionValue = types.StringValue(val.(string))
	} else {
		data.BotKmExpressionValue = types.StringNull()
	}
	if val, ok := getResponseData["kmdetectionexpr"]; ok && val != nil {
		data.Kmdetectionexpr = types.BoolValue(val.(bool))
	} else {
		data.Kmdetectionexpr = types.BoolNull()
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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_km_expression_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotKmExpressionName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("kmdetectionexpr:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Kmdetectionexpr.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
