---
subcategory: "Stream"
---

# Data Source: citrixadc_streamidentifier

The streamidentifier data source allows you to retrieve information about stream identifiers.

## Example usage

```terraform
data "citrixadc_streamidentifier" "tf_streamidentifier" {
  name = "my_streamidentifier"
}

output "selectorname" {
  value = data.citrixadc_streamidentifier.tf_streamidentifier.selectorname
}

output "samplecount" {
  value = data.citrixadc_streamidentifier.tf_streamidentifier.samplecount
}

output "sort" {
  value = data.citrixadc_streamidentifier.tf_streamidentifier.sort
}
```

## Argument Reference

* `name` - (Required) The name of stream identifier.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `acceptancethreshold` - Non-Breaching transactions to Total transactions threshold expressed in percent. Maximum of 6 decimal places is supported.
* `appflowlog` - Enable/disable Appflow logging for stream identifier
* `breachthreshold` - Breaching transactions threshold calculated over interval.
* `id` - The id of the streamidentifier. It has the same value as the `name` attribute.
* `interval` - Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.
* `log` - Location where objects collected on the identifier will be logged.
* `loginterval` - Time interval in minutes for logging the collected objects. Log interval should be greater than or equal to the inteval  of the stream identifier.
* `loglimit` - Maximum number of objects to be logged in the log interval.
* `maxtransactionthreshold` - Maximum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
* `mintransactionthreshold` - Minimum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
* `samplecount` - Size of the sample from which to select a request for evaluation. The smaller the sample count, the more accurate is the statistical data. To evaluate all requests, set the sample count to 1. However, such a low setting can result in excessive consumption of memory and processing resources.
* `selectorname` - Name of the selector to use with the stream identifier.
* `snmptrap` - Enable/disable SNMP trap for stream identifier
* `sort` - Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).
* `trackackonlypackets` - Track ack only packets as well. This setting is applicable only when packet rate limiting is being used.
* `tracktransactions` - Track transactions exceeding configured threshold. Transaction tracking can be enabled for following metric: ResponseTime. By default transaction tracking is disabled

## Import

A streamidentifier can be imported using its name, e.g.

```shell
terraform import citrixadc_streamidentifier.tf_streamidentifier my_streamidentifier
```
