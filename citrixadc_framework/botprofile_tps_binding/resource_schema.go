package botprofile_tps_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileTpsBindingResourceModel describes the resource data model.
type BotprofileTpsBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	BotBindComment types.String `tfsdk:"bot_bind_comment"`
	BotTps         types.Bool   `tfsdk:"bot_tps"`
	BotTpsAction   types.List   `tfsdk:"bot_tps_action"`
	BotTpsEnabled  types.String `tfsdk:"bot_tps_enabled"`
	BotTpsType     types.String `tfsdk:"bot_tps_type"`
	Logmessage     types.String `tfsdk:"logmessage"`
	Name           types.String `tfsdk:"name"`
	Percentage     types.Int64  `tfsdk:"percentage"`
	Threshold      types.Int64  `tfsdk:"threshold"`
}

func (r *BotprofileTpsBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_tps_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_tps": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TPS binding. For each type only binding can be configured. To  update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.",
			},
			"bot_tps_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One to more actions to be taken if bot is detected based on this TPS binding. Only LOG action can be combined with DROP, RESET, REDIRECT, or MITIGIATION action.",
			},
			"bot_tps_enabled": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enabled or disabled TPS binding.",
			},
			"bot_tps_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of TPS binding.",
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
			"percentage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum percentage increase in the requests from (or to) a IP, Geolocation, URL or Host in 30 minutes interval.",
			},
			"threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that are allowed from (or to) a IP, Geolocation, URL or Host in 1 second time interval.",
			},
		},
	}
}

func botprofile_tps_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileTpsBindingResourceModel) bot.Botprofiletpsbinding {
	tflog.Debug(ctx, "In botprofile_tps_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_tps_binding := bot.Botprofiletpsbinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_tps_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotTps.IsNull() {
		botprofile_tps_binding.Bottps = data.BotTps.ValueBool()
	}
	if !data.BotTpsEnabled.IsNull() {
		botprofile_tps_binding.Bottpsenabled = data.BotTpsEnabled.ValueString()
	}
	if !data.BotTpsType.IsNull() {
		botprofile_tps_binding.Bottpstype = data.BotTpsType.ValueString()
	}
	if !data.Logmessage.IsNull() {
		botprofile_tps_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() {
		botprofile_tps_binding.Name = data.Name.ValueString()
	}
	if !data.Percentage.IsNull() {
		botprofile_tps_binding.Percentage = utils.IntPtr(int(data.Percentage.ValueInt64()))
	}
	if !data.Threshold.IsNull() {
		botprofile_tps_binding.Threshold = utils.IntPtr(int(data.Threshold.ValueInt64()))
	}

	return botprofile_tps_binding
}

func botprofile_tps_bindingSetAttrFromGet(ctx context.Context, data *BotprofileTpsBindingResourceModel, getResponseData map[string]interface{}) *BotprofileTpsBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_tps_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_tps"]; ok && val != nil {
		data.BotTps = types.BoolValue(val.(bool))
	} else {
		data.BotTps = types.BoolNull()
	}
	if val, ok := getResponseData["bot_tps_enabled"]; ok && val != nil {
		data.BotTpsEnabled = types.StringValue(val.(string))
	} else {
		data.BotTpsEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_tps_type"]; ok && val != nil {
		data.BotTpsType = types.StringValue(val.(string))
	} else {
		data.BotTpsType = types.StringNull()
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
	if val, ok := getResponseData["percentage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Percentage = types.Int64Value(intVal)
		}
	} else {
		data.Percentage = types.Int64Null()
	}
	if val, ok := getResponseData["threshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Threshold = types.Int64Value(intVal)
		}
	} else {
		data.Threshold = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_tps_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotTpsType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
