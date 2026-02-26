package botprofile_trapinsertionurl_binding

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

// BotprofileTrapinsertionurlBindingResourceModel describes the resource data model.
type BotprofileTrapinsertionurlBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	BotBindComment             types.String `tfsdk:"bot_bind_comment"`
	BotTrapUrl                 types.String `tfsdk:"bot_trap_url"`
	BotTrapUrlInsertionEnabled types.String `tfsdk:"bot_trap_url_insertion_enabled"`
	Logmessage                 types.String `tfsdk:"logmessage"`
	Name                       types.String `tfsdk:"name"`
	Trapinsertionurl           types.Bool   `tfsdk:"trapinsertionurl"`
}

func (r *BotprofileTrapinsertionurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_trapinsertionurl_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_trap_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Request URL regex pattern for which Trap URL is inserted.",
			},
			"bot_trap_url_insertion_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the request URL pattern.",
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
			"trapinsertionurl": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the trap URL for the configured request URLs. Maximum 30 bindings can be configured per profile.",
			},
		},
	}
}

func botprofile_trapinsertionurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileTrapinsertionurlBindingResourceModel) bot.Botprofiletrapinsertionurlbinding {
	tflog.Debug(ctx, "In botprofile_trapinsertionurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_trapinsertionurl_binding := bot.Botprofiletrapinsertionurlbinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_trapinsertionurl_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotTrapUrl.IsNull() {
		botprofile_trapinsertionurl_binding.Bottrapurl = data.BotTrapUrl.ValueString()
	}
	if !data.BotTrapUrlInsertionEnabled.IsNull() {
		botprofile_trapinsertionurl_binding.Bottrapurlinsertionenabled = data.BotTrapUrlInsertionEnabled.ValueString()
	}
	if !data.Logmessage.IsNull() {
		botprofile_trapinsertionurl_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() {
		botprofile_trapinsertionurl_binding.Name = data.Name.ValueString()
	}
	if !data.Trapinsertionurl.IsNull() {
		botprofile_trapinsertionurl_binding.Trapinsertionurl = data.Trapinsertionurl.ValueBool()
	}

	return botprofile_trapinsertionurl_binding
}

func botprofile_trapinsertionurl_bindingSetAttrFromGet(ctx context.Context, data *BotprofileTrapinsertionurlBindingResourceModel, getResponseData map[string]interface{}) *BotprofileTrapinsertionurlBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_trapinsertionurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_trap_url"]; ok && val != nil {
		data.BotTrapUrl = types.StringValue(val.(string))
	} else {
		data.BotTrapUrl = types.StringNull()
	}
	if val, ok := getResponseData["bot_trap_url_insertion_enabled"]; ok && val != nil {
		data.BotTrapUrlInsertionEnabled = types.StringValue(val.(string))
	} else {
		data.BotTrapUrlInsertionEnabled = types.StringNull()
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
	if val, ok := getResponseData["trapinsertionurl"]; ok && val != nil {
		data.Trapinsertionurl = types.BoolValue(val.(bool))
	} else {
		data.Trapinsertionurl = types.BoolNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_trap_url:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.BotTrapUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
