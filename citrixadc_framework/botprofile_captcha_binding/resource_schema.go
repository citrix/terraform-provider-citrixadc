package botprofile_captcha_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileCaptchaBindingResourceModel describes the resource data model.
type BotprofileCaptchaBindingResourceModel struct {
	Id                types.String `tfsdk:"id"`
	BotBindComment    types.String `tfsdk:"bot_bind_comment"`
	BotCaptchaAction  types.List   `tfsdk:"bot_captcha_action"`
	BotCaptchaEnabled types.String `tfsdk:"bot_captcha_enabled"`
	BotCaptchaUrl     types.String `tfsdk:"bot_captcha_url"`
	Captcharesource   types.Bool   `tfsdk:"captcharesource"`
	Graceperiod       types.Int64  `tfsdk:"graceperiod"`
	Logmessage        types.String `tfsdk:"logmessage"`
	Muteperiod        types.Int64  `tfsdk:"muteperiod"`
	Name              types.String `tfsdk:"name"`
	Requestsizelimit  types.Int64  `tfsdk:"requestsizelimit"`
	Retryattempts     types.Int64  `tfsdk:"retryattempts"`
	Waittime          types.Int64  `tfsdk:"waittime"`
}

func (r *BotprofileCaptchaBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_captcha_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_captcha_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more actions to be taken when client fails captcha challenge. Only, log action can be configured with DROP, REDIRECT or RESET action.",
			},
			"bot_captcha_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the captcha binding.",
			},
			"bot_captcha_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL for which the Captcha action, if configured under IP reputation, TPS or device fingerprint, need to be applied.",
			},
			"captcharesource": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Captcha action binding. For each URL, only one binding is allowed. To update the values of an existing URL binding, user has to first unbind that binding, and then needs to bind the URL again with new values. Maximum 30 bindings can be configured per profile.",
			},
			"graceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(900),
				Description: "Time (in seconds) duration for which no new captcha challenge is sent after current captcha challenge has been answered successfully.",
			},
			"logmessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to be logged for this binding.",
			},
			"muteperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time (in seconds) duration for which client which failed captcha need to wait until allowed to try again. The requests from this client are silently dropped during the mute period.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"requestsizelimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8000),
				Description: "Length of body request (in Bytes) up to (equal or less than) which captcha challenge will be provided to client. Above this length threshold the request will be dropped. This is to avoid DOS and DDOS attacks.",
			},
			"retryattempts": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of times client can retry solving the captcha.",
			},
			"waittime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(15),
				Description: "Wait time in seconds for which ADC needs to wait for the Captcha response. This is to avoid DOS attacks.",
			},
		},
	}
}

func botprofile_captcha_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileCaptchaBindingResourceModel) bot.Botprofilecaptchabinding {
	tflog.Debug(ctx, "In botprofile_captcha_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_captcha_binding := bot.Botprofilecaptchabinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_captcha_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotCaptchaEnabled.IsNull() {
		botprofile_captcha_binding.Botcaptchaenabled = data.BotCaptchaEnabled.ValueString()
	}
	if !data.BotCaptchaUrl.IsNull() {
		botprofile_captcha_binding.Botcaptchaurl = data.BotCaptchaUrl.ValueString()
	}
	if !data.Captcharesource.IsNull() {
		botprofile_captcha_binding.Captcharesource = data.Captcharesource.ValueBool()
	}
	if !data.Graceperiod.IsNull() {
		botprofile_captcha_binding.Graceperiod = utils.IntPtr(int(data.Graceperiod.ValueInt64()))
	}
	if !data.Logmessage.IsNull() {
		botprofile_captcha_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Muteperiod.IsNull() {
		botprofile_captcha_binding.Muteperiod = utils.IntPtr(int(data.Muteperiod.ValueInt64()))
	}
	if !data.Name.IsNull() {
		botprofile_captcha_binding.Name = data.Name.ValueString()
	}
	if !data.Requestsizelimit.IsNull() {
		botprofile_captcha_binding.Requestsizelimit = utils.IntPtr(int(data.Requestsizelimit.ValueInt64()))
	}
	if !data.Retryattempts.IsNull() {
		botprofile_captcha_binding.Retryattempts = utils.IntPtr(int(data.Retryattempts.ValueInt64()))
	}
	if !data.Waittime.IsNull() {
		botprofile_captcha_binding.Waittime = utils.IntPtr(int(data.Waittime.ValueInt64()))
	}

	return botprofile_captcha_binding
}

func botprofile_captcha_bindingSetAttrFromGet(ctx context.Context, data *BotprofileCaptchaBindingResourceModel, getResponseData map[string]interface{}) *BotprofileCaptchaBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_captcha_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_captcha_enabled"]; ok && val != nil {
		data.BotCaptchaEnabled = types.StringValue(val.(string))
	} else {
		data.BotCaptchaEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_captcha_url"]; ok && val != nil {
		data.BotCaptchaUrl = types.StringValue(val.(string))
	} else {
		data.BotCaptchaUrl = types.StringNull()
	}
	if val, ok := getResponseData["captcharesource"]; ok && val != nil {
		data.Captcharesource = types.BoolValue(val.(bool))
	} else {
		data.Captcharesource = types.BoolNull()
	}
	if val, ok := getResponseData["graceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Graceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Graceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["logmessage"]; ok && val != nil {
		data.Logmessage = types.StringValue(val.(string))
	} else {
		data.Logmessage = types.StringNull()
	}
	if val, ok := getResponseData["muteperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Muteperiod = types.Int64Value(intVal)
		}
	} else {
		data.Muteperiod = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["requestsizelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Requestsizelimit = types.Int64Value(intVal)
		}
	} else {
		data.Requestsizelimit = types.Int64Null()
	}
	if val, ok := getResponseData["retryattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retryattempts = types.Int64Value(intVal)
		}
	} else {
		data.Retryattempts = types.Int64Null()
	}
	if val, ok := getResponseData["waittime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Waittime = types.Int64Value(intVal)
		}
	} else {
		data.Waittime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_captcha_url:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotCaptchaUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
