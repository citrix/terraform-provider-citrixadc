---
subcategory: "CI"
---

# Resource: contentinspectionwasmprofile

The contentinspectionwasmprofile resource is used to create contentinspectionwasmprofile.


## Example usage

```hcl
resource "citrixadc_contentinspectionwasmprofile" "tf_contentinspectionwasmprofile" {
  name              = "my_ci_wasmprofile"
  timeout           = 2000
  timeoutaction     = "BYPASS"
  maxbodylen        = 32
  anomalousdatasize = 256
  anomalousttfbtime = 2000
}
```


## Argument Reference

* `name` - (Required) Name of CI WASM profile.
* `timeout` - (Optional) Timeout (in milliseconds) for the connection with the CI WASM agent.
* `timeoutaction` - (Optional) Timeout action for the connection with the CI agent. Either the original request can be bypassed i.e. request/response is forwarded to the endpoint or the transaction is dropped/reset. Possible values = BYPASS, DROP, RESET
* `maxbodylen` - (Optional) Max data size (in KB) that will be sent to the CI Agent. Default is 16KB. Maximum value that can be configured is 32KB.
* `anomalousdatasize` - (Optional) Transaction data size (in KB) greater than which a transaction is considered as anomalous. Default is 512KB.
* `anomalousttfbtime` - (Optional) Transaction time (in milliseconds) above which a transaction is considered as anomalous. Default is 1 seconds.
* `wasmmodule` - (Optional) Name of the WASM Module.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionwasmprofile. It has the same value as the `name` attribute.


## Import

A contentinspectionwasmprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionwasmprofile.tf_contentinspectionwasmprofile my_ci_wasmprofile
```
