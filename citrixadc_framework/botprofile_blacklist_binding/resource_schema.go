package botprofile_blacklist_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about this binding.",
			},
			"bot_blacklist": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.",
			},
			"bot_blacklist_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "One or more actions to be taken if  bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.",
			},
			"bot_blacklist_enabled": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled or disbaled black-list binding.",
			},
			"bot_blacklist_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the black-list entry.",
			},
			"bot_blacklist_value": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Value of the bot black-list entry.",
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

func botprofile_blacklist_bindingGetThePayloadFromthePlan(ctx context.Context, data *BotprofileBlacklistBindingResourceModel) bot.Botprofileblacklistbinding {
	tflog.Debug(ctx, "In botprofile_blacklist_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botprofile_blacklist_binding := bot.Botprofileblacklistbinding{}
	if !data.BotBindComment.IsNull() && !data.BotBindComment.IsUnknown() {
		botprofile_blacklist_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotBlacklist.IsNull() && !data.BotBlacklist.IsUnknown() {
		botprofile_blacklist_binding.Botblacklist = data.BotBlacklist.ValueBool()
	}
	if !data.BotBlacklistAction.IsNull() && !data.BotBlacklistAction.IsUnknown() {
		var bot_blacklist_actionList []string
		data.BotBlacklistAction.ElementsAs(ctx, &bot_blacklist_actionList, false)
		botprofile_blacklist_binding.Botblacklistaction = bot_blacklist_actionList
	}
	if !data.BotBlacklistEnabled.IsNull() && !data.BotBlacklistEnabled.IsUnknown() {
		botprofile_blacklist_binding.Botblacklistenabled = data.BotBlacklistEnabled.ValueString()
	}
	if !data.BotBlacklistType.IsNull() && !data.BotBlacklistType.IsUnknown() {
		botprofile_blacklist_binding.Botblacklisttype = data.BotBlacklistType.ValueString()
	}
	if !data.BotBlacklistValue.IsNull() && !data.BotBlacklistValue.IsUnknown() {
		botprofile_blacklist_binding.Botblacklistvalue = data.BotBlacklistValue.ValueString()
	}
	if !data.Logmessage.IsNull() && !data.Logmessage.IsUnknown() {
		botprofile_blacklist_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
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
	if val, ok := getResponseData["bot_blacklist_action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.BotBlacklistAction = listValue
		} else {
			data.BotBlacklistAction = types.ListNull(types.StringType)
		}
	} else {
		data.BotBlacklistAction = types.ListNull(types.StringType)
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

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_blacklist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotBlacklistValue.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// Datasource variant: faithfully copy the GET response and set the ID (the
// datasource has no Create to set it). ID uses the legacy attribute order
// (name, bot_blacklist_value) so it matches resource_id_mapping.json.
func botprofile_blacklist_bindingSetAttrFromGetForDatasource(ctx context.Context, data *BotprofileBlacklistBindingResourceModel, getResponseData map[string]interface{}) *BotprofileBlacklistBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_blacklist_bindingSetAttrFromGetForDatasource Function")

	botprofile_blacklist_bindingSetAttrFromGet(ctx, data, getResponseData)

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_blacklist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotBlacklistValue.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
