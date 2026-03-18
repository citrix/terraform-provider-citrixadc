package botprofile_ipreputation_binding

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

// BotprofileIpreputationBindingResourceModel describes the resource data model.
type BotprofileIpreputationBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	BotBindComment  types.String `tfsdk:"bot_bind_comment"`
	BotIprepAction  types.List   `tfsdk:"bot_iprep_action"`
	BotIprepEnabled types.String `tfsdk:"bot_iprep_enabled"`
	BotIpreputation types.Bool   `tfsdk:"bot_ipreputation"`
	Category        types.String `tfsdk:"category"`
	Logmessage      types.String `tfsdk:"logmessage"`
	Name            types.String `tfsdk:"name"`
}

func (r *BotprofileIpreputationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile_ipreputation_binding resource.",
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_iprep_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "One or more actions to be taken if bot is detected based on this IP Reputation binding. Only LOG action can be combinded with DROP, RESET, REDIRECT or MITIGATION action.",
			},
			"bot_iprep_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled or disabled IP-repuation binding.",
			},
			"bot_ipreputation": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP reputation binding. For each category, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with the new values.",
			},
			"category": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Repuation category. Following IP Reuputation categories are allowed:\n*IP_BASED - This category checks whether client IP is malicious or not.\n*BOTNET - This category includes Botnet C&C channels, and infected zombie machines controlled by Bot master.\n*SPAM_SOURCES - This category includes tunneling spam messages through a proxy, anomalous SMTP activities, and forum spam activities.\n*SCANNERS - This category includes all reconnaissance such as probes, host scan, domain scan, and password brute force attack.\n*DOS - This category includes DOS, DDOS, anomalous sync flood, and anomalous traffic detection.\n*REPUTATION - This category denies access from IP addresses currently known to be infected with malware. This category also includes IPs with average low Webroot Reputation Index score. Enabling this category will prevent access from sources identified to contact malware distribution points.\n*PHISHING - This category includes IP addresses hosting phishing sites and other kinds of fraud activities such as ad click fraud or gaming fraud.\n*PROXY - This category includes IP addresses providing proxy services.\n*NETWORK - IPs providing proxy and anonymization services including The Onion Router aka TOR or darknet.\n*MOBILE_THREATS - This category checks client IP with the list of IPs harmful for mobile devices.\n*WINDOWS_EXPLOITS - This category includes active IP address offering or distributig malware, shell code, rootkits, worms or viruses.\n*WEB_ATTACKS - This category includes cross site scripting, iFrame injection, SQL injection, cross domain injection or domain password brute force attack.\n*TOR_PROXY - This category includes IP address acting as exit nodes for the Tor Network.\n*CLOUD - This category checks client IP with list of public cloud IPs.\n*CLOUD_AWS - This category checks client IP with list of public cloud IPs from Amazon Web Services.\n*CLOUD_GCP - This category checks client IP with list of public cloud IPs from Google Cloud Platform.\n*CLOUD_AZURE - This category checks client IP with list of public cloud IPs from Azure.\n*CLOUD_ORACLE - This category checks client IP with list of public cloud IPs from Oracle.\n*CLOUD_IBM - This category checks client IP with list of public cloud IPs from IBM.\n*CLOUD_SALESFORCE - This category checks client IP with list of public cloud IPs from Salesforce.",
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

func botprofile_ipreputation_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotprofileIpreputationBindingResourceModel) bot.Botprofileipreputationbinding {
	tflog.Debug(ctx, "In botprofile_ipreputation_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botprofile_ipreputation_binding := bot.Botprofileipreputationbinding{}
	if !data.BotBindComment.IsNull() {
		botprofile_ipreputation_binding.Botbindcomment = data.BotBindComment.ValueString()
	}
	if !data.BotIprepEnabled.IsNull() {
		botprofile_ipreputation_binding.Botiprepenabled = data.BotIprepEnabled.ValueString()
	}
	if !data.BotIpreputation.IsNull() {
		botprofile_ipreputation_binding.Botipreputation = data.BotIpreputation.ValueBool()
	}
	if !data.Category.IsNull() {
		botprofile_ipreputation_binding.Category = data.Category.ValueString()
	}
	if !data.Logmessage.IsNull() {
		botprofile_ipreputation_binding.Logmessage = data.Logmessage.ValueString()
	}
	if !data.Name.IsNull() {
		botprofile_ipreputation_binding.Name = data.Name.ValueString()
	}

	return botprofile_ipreputation_binding
}

func botprofile_ipreputation_bindingSetAttrFromGet(ctx context.Context, data *BotprofileIpreputationBindingResourceModel, getResponseData map[string]interface{}) *BotprofileIpreputationBindingResourceModel {
	tflog.Debug(ctx, "In botprofile_ipreputation_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bot_bind_comment"]; ok && val != nil {
		data.BotBindComment = types.StringValue(val.(string))
	} else {
		data.BotBindComment = types.StringNull()
	}
	if val, ok := getResponseData["bot_iprep_enabled"]; ok && val != nil {
		data.BotIprepEnabled = types.StringValue(val.(string))
	} else {
		data.BotIprepEnabled = types.StringNull()
	}
	if val, ok := getResponseData["bot_ipreputation"]; ok && val != nil {
		data.BotIpreputation = types.BoolValue(val.(bool))
	} else {
		data.BotIpreputation = types.BoolNull()
	}
	if val, ok := getResponseData["category"]; ok && val != nil {
		data.Category = types.StringValue(val.(string))
	} else {
		data.Category = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("category:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Category.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
