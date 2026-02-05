---
subcategory: "AppFlow"
---

# Data Source `appflowaction`

The appflowaction data source allows you to retrieve information about an existing appflowaction.


## Example usage

```terraform
data "citrixadc_appflowaction" "tf_appflowaction" {
  name = "test_action"
}

output "name" {
  value = data.citrixadc_appflowaction.tf_appflowaction.name
}

output "securityinsight" {
  value = data.citrixadc_appflowaction.tf_appflowaction.securityinsight
}
```


## Argument Reference

* `name` - (Required) Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow action" or 'my appflow action').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowaction. It has the same value as the `name` attribute.
* `botinsight` - On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.
* `ciinsight` - On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.
* `clientsidemeasurements` - On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.
* `collectors` - Name(s) of collector(s) to be associated with the AppFlow action.
* `comment` - Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)
* `distributionalgorithm` - On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.
* `metricslog` - If only the stats records are to be exported, turn on this option.
* `newname` - New name for the AppFlow action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `pagetracking` - On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.
* `securityinsight` - On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.
* `transactionlog` - Log ANOMALOUS or ALL transactions
* `videoanalytics` - On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.
* `webinsight` - On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.
