package botprofile_ipreputation_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BotprofileIpreputationBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bot_bind_comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this binding.",
			},
			"bot_iprep_action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
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
				Required:    true,
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
