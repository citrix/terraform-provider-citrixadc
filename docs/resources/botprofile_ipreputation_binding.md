---
subcategory: "Bot"
---

# Resource: botprofile_ipreputation_binding

The botprofile_ipreputation_binding resource is used to bind ipreputation to botprofile.


## Example usage

```hcl
resource "citrixadc_botprofile" "tf_botprofile" {
  name                     = "tf_botprofile"
  errorurl                 = "http://www.citrix.com"
  trapurl                  = "/http://www.citrix.com"
  comment                  = "tf_botprofile comment"
  bot_enable_white_list    = "ON"
  bot_enable_black_list    = "ON"
  bot_enable_rate_limit    = "ON"
  devicefingerprint        = "ON"
  devicefingerprintaction  = ["LOG", "RESET"]
  bot_enable_ip_reputation = "ON"
  trap                     = "ON"
  trapaction               = ["LOG", "RESET"]
  bot_enable_tps           = "ON"
}
resource "citrixadc_botprofile_ipreputation_binding" "tf_binding" {
  name              = citrixadc_botprofile.tf_botprofile.name
  bot_ipreputation  = "true"
  category          = "BOTNETS"
  bot_iprep_action  = ["LOG", "REDIRECT"]
  bot_bind_comment  = "TestingIpreputation"
  bot_iprep_enabled = "ON"
  logmessage        = "MessageTesting"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `category` - (Required) IP Repuation category. Following IP Reuputation categories are allowed: *IP_BASED - This category checks whether client IP is malicious or not. *BOTNET - This category includes Botnet C&C channels, and infected zombie machines controlled by Bot master. *SPAM_SOURCES - This category includes tunneling spam messages through a proxy, anomalous SMTP activities, and forum spam activities. *SCANNERS - This category includes all reconnaissance such as probes, host scan, domain scan, and password brute force attack. *DOS - This category includes DOS, DDOS, anomalous sync flood, and anomalous traffic detection. *REPUTATION - This category denies access from IP addresses currently known to be infected with malware. This category also includes IPs with average low Webroot Reputation Index score. Enabling this category will prevent access from sources identified to contact malware distribution points. *PHISHING - This category includes IP addresses hosting phishing sites and other kinds of fraud activities such as ad click fraud or gaming fraud. *PROXY - This category includes IP addresses providing proxy services. *NETWORK - IPs providing proxy and anonymization services including The Onion Router aka TOR or darknet. *MOBILE_THREATS - This category checks client IP with the list of IPs harmful for mobile devices.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_iprep_action` - (Optional) One or more actions to be taken if bot is detected based on this IP Reputation binding. Only LOG action can be combinded with DROP, RESET, REDIRECT or MITIGATION action.
* `bot_iprep_enabled` - (Optional) Enabled or disabled IP-repuation binding.
* `bot_ipreputation` - (Optional) IP reputation binding. For each category, only one binding is allowed. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with the new values.
* `logmessage` - (Optional) Message to be logged for this binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_ipreputation_binding. It is the concatenation of `name` and `category` attributes separated by comma.


## Import

A botprofile_ipreputation_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_ipreputation_binding.tf_binding tf_botprofile,BOTNETS
```
