---
subcategory: "Metrics"
---

# Data Source: metricsprofile

The metricsprofile data source allows you to retrieve information about an existing metrics profile configured on the Citrix ADC, such as its output mode, collector, serve mode, and export frequency.


## Example usage

```terraform
data "citrixadc_metricsprofile" "example" {
  name = "splunk_profile"
}

output "metricsprofile_outputmode" {
  value = data.citrixadc_metricsprofile.example.outputmode
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile. It has the same value as the `name` attribute.
* `collector` - The collector HTTP/HTTPS service to which metrics are exported.
* `metrics` - Whether metrics collection is enabled or disabled. Possible values: [ ENABLED, DISABLED ].
* `outputmode` - The format in which metrics data is generated. Possible values: [ avro, prometheus, influx, json ].
* `servemode` - The metrics serve mode. Possible values: [ Push, Pull ].
* `schemafile` - The json schema file containing the list of counters exported by metricscollector.
* `metricsexportfrequency` - The metrics export frequency in seconds.
* `metricsendpointurl` - The URL at which the metrics data is uploaded on the endpoint.

Note: The `metricsauthtoken` token is a secret and is not returned by the NITRO API, so it is not populated by this data source.
