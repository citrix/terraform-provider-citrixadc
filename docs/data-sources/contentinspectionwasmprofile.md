---
subcategory: "Content Inspection"
---

# Data Source: contentinspectionwasmprofile

The contentinspectionwasmprofile data source allows you to retrieve information about a content inspection WASM profile.


## Example usage

```terraform
data "citrixadc_contentinspectionwasmprofile" "tf_contentinspectionwasmprofile" {
  name = "my_ci_wasmprofile"
}

output "timeout" {
  value = data.citrixadc_contentinspectionwasmprofile.tf_contentinspectionwasmprofile.timeout
}

output "timeoutaction" {
  value = data.citrixadc_contentinspectionwasmprofile.tf_contentinspectionwasmprofile.timeoutaction
}
```


## Argument Reference

* `name` - (Required) Name of CI WASM profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `timeout` - Timeout (in milliseconds) for the connection with the CI WASM agent.
* `timeoutaction` - Timeout action for the connection with the CI agent. Either the original request can be bypassed i.e. request/response is forwarded to the endpoint or the transaction is dropped/reset. Possible values = BYPASS, DROP, RESET
* `maxbodylen` - Max data size (in KB) that will be sent to the CI Agent. Default is 16KB. Maximum value that can be configured is 32KB.
* `anomalousdatasize` - Transaction data size (in KB) greater than which a transaction is considered as anomalous. Default is 512KB.
* `anomalousttfbtime` - Transaction time (in milliseconds) above which a transaction is considered as anomalous. Default is 1 seconds.
* `wasmmodule` - Name of the WASM Module.
* `id` - The id of the contentinspectionwasmprofile. It has the same value as the `name` attribute.

## Import

A contentinspectionwasmprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionwasmprofile.tf_contentinspectionwasmprofile my_ci_wasmprofile
```
