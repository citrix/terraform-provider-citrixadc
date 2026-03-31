package botprofile_blacklist_binding

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

// BotprofileBlacklistBindingResourceModel describes the resource data model.
type BotprofileBlacklistBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	BotBindComment      types.String `tfsdk:"bot_bind_comment"`
	BotBlacklist        types.Bool   `tfsdk:"bot_blacklist"`
	BotBlacklistAction  types.List   `tfsdk:"bot_blacklist_action"`
	BotBlacklistEnabled types.String `tfsdk:"bot_blacklist_enabled"`
	BotBlacklistType    types.String `tfsdk:"bot_blacklist_type"`
	BotBlacklistValue   types.String `tfsdk:"bot_blacklist_value"`
	Logmessage          types.String `tfsdk:"logmessage"`
	Name                types.String `tfsdk:"name"`
}

func (r *BotprofileBlacklistBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_blacklist_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_blacklist": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.",
			},
			"bot_blacklist_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more actions to be taken if  bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.",
			},
			"bot_blacklist_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled or disbaled black-list binding.",
			},
			"bot_blacklist_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the black-list entry.",
			},
			"bot_blacklist_value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value of the bot black-list entry.",
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

func botprofile_blacklist_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileBlacklistBindingResourceModel) bot.Botprofileblacklistbinding {
	tflog.Debug(ctx, "In botprofile_blacklist_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_blacklist_binding := bot.Botprofileblacklistbinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_blacklist_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotBlacklist.IsNull() {
		botprofile_blacklist_binding.Botblacklist = data.BotBlacklist.ValueBool()
	}
	if !data.BotBlacklistEnabled.IsNull() {
		botprofile_blacklist_binding.Botblacklistenabled = data.BotBlacklistEnabled.ValueString()
	}
	if !data.BotBlacklistType.IsNull() {
		botprofile_blacklist_binding.Botblacklisttype = data.BotBlacklistType.ValueString()
	}
	if !data.BotBlacklistValue.IsNull() {
		botprofile_blacklist_binding.Botblacklistvalue = data.BotBlacklistValue.ValueString()
	}
	if !data.Logmessage.IsNull() {
		botprofile_blacklist_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() {
		botprofile_blacklist_binding.Name = data.Name.ValueString()
	}

	return botprofile_blacklist_binding
}

func botprofile_blacklist_bindingSetAttrFromGet(ctx context.Context, data *BotprofileBlacklistBindingResourceModel, getResponseData map[string]interface{}) *BotprofileBlacklistBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_blacklist_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_blacklist"]; ok && val != nil {
		data.BotBlacklist = types.BoolValue(val.(bool))
	} else {
		data.BotBlacklist = types.BoolNull()
	}
	if val, ok := getResponseData["bot_blacklist_enabled"]; ok && val != nil {
		data.BotBlacklistEnabled = types.StringValue(val.(string))
	} else {
		data.BotBlacklistEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_blacklist_type"]; ok && val != nil {
		data.BotBlacklistType = types.StringValue(val.(string))
	} else {
		data.BotBlacklistType = types.StringNull()
	}
	if val, ok := getResponseData["bot_blacklist_value"]; ok && val != nil {
		data.BotBlacklistValue = types.StringValue(val.(string))
	} else {
		data.BotBlacklistValue = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("bot_blacklist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotBlacklistValue.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
