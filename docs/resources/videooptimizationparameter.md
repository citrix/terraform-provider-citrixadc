---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationparameter

This resource is used to manage the global video optimization parameters.


## Example usage

```hcl
resource "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
  randomsamplingpercentage = 25
  quicpacingrate           = 1024
}
```


## Argument Reference

* `randomsamplingpercentage` - (Optional) Random sampling percentage applied to video traffic for optimization decisions. Minimum value = `0` Maximum value = `100`. Defaults to `0`.
* `quicpacingrate` - (Optional) QUIC video pacing rate, in Kbps.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the videooptimizationparameter. It is set to `videooptimizationparameter-config`.

~> **Note** This is a singleton resource — only one instance can exist per Citrix ADC. There is no delete operation: removing the resource from your configuration only removes it from Terraform state; the parameters remain on the ADC at their last-applied values.


## Import

A videooptimizationparameter can be imported using its id (the static string `videooptimizationparameter-config`), e.g.

```shell
terraform import citrixadc_videooptimizationparameter.tf_videooptimizationparameter videooptimizationparameter-config
```
