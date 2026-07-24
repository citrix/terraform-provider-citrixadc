---
subcategory: "Metrics"
---

# Resource: metricsprofile

This resource is used to manage metrics profiles on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "prometheus_profile"
  outputmode             = "prometheus"
  metrics                = "ENABLED"
  servemode              = "Pull"
  metricsexportfrequency = 60
}
```

### Using metricsauthtoken (sensitive attribute - persisted in state)

```hcl
variable "metricsprofile_metricsauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "splunk_profile"
  outputmode             = "json"
  metrics                = "ENABLED"
  servemode              = "Push"
  collector              = "https://collector.example.com:8088"
  metricsendpointurl     = "/services/collector/event"
  metricsexportfrequency = 60
  metricsauthtoken       = var.metricsprofile_metricsauthtoken
}
```

### Using metricsauthtoken_wo (write-only/ephemeral - NOT persisted in state)

The `metricsauthtoken_wo` attribute provides an ephemeral path for the endpoint authentication token. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `metricsauthtoken_wo_version`.

```hcl
variable "metricsprofile_metricsauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                        = "splunk_profile"
  outputmode                  = "json"
  metrics                     = "ENABLED"
  servemode                   = "Push"
  collector                   = "https://collector.example.com:8088"
  metricsendpointurl          = "/services/collector/event"
  metricsexportfrequency      = 60
  metricsauthtoken_wo         = var.metricsprofile_metricsauthtoken
  metricsauthtoken_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                        = "splunk_profile"
  outputmode                  = "json"
  metrics                     = "ENABLED"
  servemode                   = "Push"
  collector                   = "https://collector.example.com:8088"
  metricsendpointurl          = "/services/collector/event"
  metricsexportfrequency      = 60
  metricsauthtoken_wo         = var.metricsprofile_metricsauthtoken
  metricsauthtoken_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Changing this attribute forces a new resource to be created.
* `collector` - (Optional) The collector should be a HTTP/HTTPS service. Not required when `servemode` is `Pull`.
* `metrics` - (Optional) This option is used to enable or disable metrics. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `outputmode` - (Optional) This option indicates the format in which metrics data is generated. Possible values: [ avro, prometheus, influx, json ]. Defaults to `"avro"`.
* `servemode` - (Optional) This option is to configure metrics pull or push mode. In push mode metricscollector exports metrics to the configured collector. In pull mode, metricscollector only generates the metrics which will be pulled by an external agent. No collector configuration is required in pull mode and it is applicable only for output mode Prometheus. Possible values: [ Push, Pull ]. Defaults to `"Push"`.
* `schemafile` - (Optional) This option is for configuring the json schema file containing a list of counters to be exported by metricscollector. Schema file should be present under the `/var/metrics_conf` path.
* `metricsexportfrequency` - (Optional) This option is for configuring the metrics export frequency in seconds. The frequency value must be in the [30, 300] seconds range. Minimum value = 30. Maximum value = 300. Defaults to `30`.
* `metricsendpointurl` - (Optional) The URL at which to upload the metrics data on the endpoint.
* `metricsauthtoken` - (Optional, Sensitive) Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For example, in the case of Splunk, the Authorization header is required to be of the form `Splunk <auth-token>`. The value is persisted in Terraform state (encrypted). See also `metricsauthtoken_wo` for an ephemeral alternative.
* `metricsauthtoken_wo` - (Optional, Sensitive, WriteOnly) Same as `metricsauthtoken`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `metricsauthtoken_wo_version`. If both `metricsauthtoken` and `metricsauthtoken_wo` are set, `metricsauthtoken_wo` takes precedence.
* `metricsauthtoken_wo_version` - (Optional) An integer version tracker for `metricsauthtoken_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile. It has the same value as the `name` attribute.


## Import

A metricsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_metricsprofile.tf_metricsprofile splunk_profile
```
