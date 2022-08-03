---
subcategory: "AppFlow"
---

# Resource: appflowaction

The appflowaction resource is used to create appflowaction.


## Example usage

```hcl
resource "citrixadc_appflowaction" "tf_appflowaction" {
  name = "test_action"
  collectors = ["tf_collector", "tf2_collector" ]
  securityinsight = "ENABLED"
  botinsight      = "ENABLED"
  videoanalytics  = "ENABLED"
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf2_collector -IPAddress 192.168.2.3
#add appflow collector tf_collector -IPAddress 192.168.2.2

# ----------------- NOT YET IMPLEMENTED -----------------------
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
#   name      = "tf2_collector"
#   ipaddress = "192.168.2.3"
#   port      = 80
# }
```


## Argument Reference

* `name` - (Required) Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow action" or 'my appflow action').
* `botinsight` - (Optional) On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.
* `ciinsight` - (Optional) On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.
* `clientsidemeasurements` - (Optional) On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.
* `collectors` - (Optional) Name(s) of collector(s) to be associated with the AppFlow action.
* `comment` - (Optional) Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)
* `distributionalgorithm` - (Optional) On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.
* `metricslog` - (Optional) If only the stats records are to be exported, turn on this option.
* `pagetracking` - (Optional) On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.
* `securityinsight` - (Optional) On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.
* `transactionlog` - (Optional) Log ANOMALOUS or ALL transactions
* `videoanalytics` - (Optional) On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.
* `webinsight` - (Optional) On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowaction. It has the same value as the `name` attribute.


## Import

A appflowaction can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowaction.tf_appflowaction test_action
```