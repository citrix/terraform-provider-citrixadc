---
subcategory: "Stream"
---

# Resource: streamidentifier

The streamidentifier resource is used to create streamidentifier.


## Example usage

```hcl
resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name         = "my_streamidentifier"
  selectorname = "my_streamselector"
  samplecount  = 10
  sort         = "CONNECTIONS"
  snmptrap     = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) The name of stream identifier.
* `acceptancethreshold` - (Optional) Non-Breaching transactions to Total transactions threshold expressed in percent. Maximum of 6 decimal places is supported.
* `appflowlog` - (Optional) Enable/disable Appflow logging for stream identifier
* `breachthreshold` - (Optional) Breaching transactions threshold calculated over interval.
* `interval` - (Optional) Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.
* `maxtransactionthreshold` - (Optional) Maximum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
* `mintransactionthreshold` - (Optional) Minimum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.
* `samplecount` - (Optional) Size of the sample from which to select a request for evaluation. The smaller the sample count, the more accurate is the statistical data. To evaluate all requests, set the sample count to 1. However, such a low setting can result in excessive consumption of memory and processing resources.
* `selectorname` - (Optional) Name of the selector to use with the stream identifier.
* `snmptrap` - (Optional) Enable/disable SNMP trap for stream identifier
* `sort` - (Optional) Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).
* `trackackonlypackets` - (Optional) Track ack only packets as well. This setting is applicable only when packet rate limiting is being used.
* `tracktransactions` - (Optional) Track transactions exceeding configured threshold. Transaction tracking can be enabled for following metric: ResponseTime. By default transaction tracking is disabled


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the streamidentifier. It has the same value as the `name` attribute.


## Import

A streamidentifier can be imported using its name, e.g.

```shell
terraform import citrixadc_streamidentifier.tf_streamidentifier my_streamidentifier
```
