package botprofile_whitelist_binding

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

// BotprofileWhitelistBindingResourceModel describes the resource data model.
type BotprofileWhitelistBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	BotBindComment      types.String `tfsdk:"bot_bind_comment"`
	BotWhitelist        types.Bool   `tfsdk:"bot_whitelist"`
	BotWhitelistEnabled types.String `tfsdk:"bot_whitelist_enabled"`
	BotWhitelistType    types.String `tfsdk:"bot_whitelist_type"`
	BotWhitelistValue   types.String `tfsdk:"bot_whitelist_value"`
	Log                 types.String `tfsdk:"log"`
	Logmessage          types.String `tfsdk:"logmessage"`
	Name                types.String `tfsdk:"name"`
}

func (r *BotprofileWhitelistBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_whitelist_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about this binding.",
			},
			"bot_whitelist": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Whitelist binding. Maximum 32 bindings can be configured per profile for Whitelist detection.",
			},
			"bot_whitelist_enabled": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled or disabled white-list binding.",
			},
			"bot_whitelist_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the white-list entry.",
			},
			"bot_whitelist_value": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Value of bot white-list entry.",
			},
			"log": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enable logging for Whitelist binding.",
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

func botprofile_whitelist_bindingGetThePayloadFromthePlan(ctx context.Context, data *BotprofileWhitelistBindingResourceModel) bot.Botprofilewhitelistbinding {
	tflog.Debug(ctx, "In botprofile_whitelist_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botprofile_whitelist_binding := bot.Botprofilewhitelistbinding{}
	if !data.BotBindComment.IsNull() && !data.BotBindComment.IsUnknown() {
		botprofile_whitelist_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotWhitelist.IsNull() && !data.BotWhitelist.IsUnknown() {
		botprofile_whitelist_binding.Botwhitelist = data.BotWhitelist.ValueBool()
	}
	if !data.BotWhitelistEnabled.IsNull() && !data.BotWhitelistEnabled.IsUnknown() {
		botprofile_whitelist_binding.Botwhitelistenabled = data.BotWhitelistEnabled.ValueString()
	}
	if !data.BotWhitelistType.IsNull() && !data.BotWhitelistType.IsUnknown() {
		botprofile_whitelist_binding.Botwhitelisttype = data.BotWhitelistType.ValueString()
	}
	if !data.BotWhitelistValue.IsNull() && !data.BotWhitelistValue.IsUnknown() {
		botprofile_whitelist_binding.Botwhitelistvalue = data.BotWhitelistValue.ValueString()
	}
	if !data.Log.IsNull() && !data.Log.IsUnknown() {
		botprofile_whitelist_binding.Log = data.Log.ValueString()
	}
	if !data.Logmessage.IsNull() && !data.Logmessage.IsUnknown() {
		botprofile_whitelist_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botprofile_whitelist_binding.Name = data.Name.ValueString()
	}

	return botprofile_whitelist_binding
}

func botprofile_whitelist_bindingSetAttrFromGet(ctx context.Context, data *BotprofileWhitelistBindingResourceModel, getResponseData map[string]interface{}) *BotprofileWhitelistBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_whitelist_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_whitelist"]; ok && val != nil {
		data.BotWhitelist = types.BoolValue(val.(bool))
	} else {
		data.BotWhitelist = types.BoolNull()
	}
	if val, ok := getResponseData["bot_whitelist_enabled"]; ok && val != nil {
		data.BotWhitelistEnabled = types.StringValue(val.(string))
	} else {
		data.BotWhitelistEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_whitelist_type"]; ok && val != nil {
		data.BotWhitelistType = types.StringValue(val.(string))
	} else {
		data.BotWhitelistType = types.StringNull()
	}
	if val, ok := getResponseData["bot_whitelist_value"]; ok && val != nil {
		data.BotWhitelistValue = types.StringValue(val.(string))
	} else {
		data.BotWhitelistValue = types.StringNull()
	}
	if val, ok := getResponseData["log"]; ok && val != nil {
		data.Log = types.StringValue(val.(string))
	} else {
		data.Log = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("bot_whitelist:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotWhitelist.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("bot_whitelist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotWhitelistValue.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
