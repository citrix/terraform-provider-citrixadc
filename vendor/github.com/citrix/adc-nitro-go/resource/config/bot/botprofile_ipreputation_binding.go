/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package bot

/**
* Binding class showing the ipreputation that can be bound to botprofile.
*/
type Botprofileipreputationbinding struct {
	/**
	* IP reputation binding. For each category, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with the new values.
	*/
	Botipreputation bool `json:"bot_ipreputation,omitempty"`
	/**
	* IP Repuation category. Following IP Reuputation categories are allowed:
		*IP_BASED - This category checks whether client IP is malicious or not.
		*BOTNET - This category includes Botnet C&C channels, and infected zombie machines controlled by Bot master.
		*SPAM_SOURCES - This category includes tunneling spam messages through a proxy, anomalous SMTP activities, and forum spam activities.
		*SCANNERS - This category includes all reconnaissance such as probes, host scan, domain scan, and password brute force attack.
		*DOS - This category includes DOS, DDOS, anomalous sync flood, and anomalous traffic detection.
		*REPUTATION - This category denies access from IP addresses currently known to be infected with malware. This category also includes IPs with average low Webroot Reputation Index score. Enabling this category will prevent access from sources identified to contact malware distribution points.
		*PHISHING - This category includes IP addresses hosting phishing sites and other kinds of fraud activities such as ad click fraud or gaming fraud.
		*PROXY - This category includes IP addresses providing proxy services.
		*NETWORK - IPs providing proxy and anonymization services including The Onion Router aka TOR or darknet.
		*MOBILE_THREATS - This category checks client IP with the list of IPs harmful for mobile devices.
		*WINDOWS_EXPLOITS - This category includes active IP address offering or distributig malware, shell code, rootkits, worms or viruses.
		*WEB_ATTACKS - This category includes cross site scripting, iFrame injection, SQL injection, cross domain injection or domain password brute force attack.
		*TOR_PROXY - This category includes IP address acting as exit nodes for the Tor Network.
		*CLOUD - This category checks client IP with list of public cloud IPs.
		*CLOUD_AWS - This category checks client IP with list of public cloud IPs from Amazon Web Services.
		*CLOUD_GCP - This category checks client IP with list of public cloud IPs from Google Cloud Platform.
		*CLOUD_AZURE - This category checks client IP with list of public cloud IPs from Azure.
		*CLOUD_ORACLE - This category checks client IP with list of public cloud IPs from Oracle.
		*CLOUD_IBM - This category checks client IP with list of public cloud IPs from IBM.
		*CLOUD_SALESFORCE - This category checks client IP with list of public cloud IPs from Salesforce.
	*/
	Category string `json:"category,omitempty"`
	/**
	* Enabled or disabled IP-repuation binding.
	*/
	Botiprepenabled string `json:"bot_iprep_enabled,omitempty"`
	/**
	* One or more actions to be taken if bot is detected based on this IP Reputation binding. Only LOG action can be combinded with DROP, RESET, REDIRECT or MITIGATION action.
	*/
	Botiprepaction []string `json:"bot_iprep_action,omitempty"`
	/**
	* Message to be logged for this binding.
	*/
	Logmessage string `json:"logmessage,omitempty"`
	/**
	* Any comments about this binding.
	*/
	Botbindcomment string `json:"bot_bind_comment,omitempty"`
	/**
	* Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
	*/
	Name string `json:"name,omitempty"`


}