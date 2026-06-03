---
subcategory: "VideoOptimization"
---

# Data Source: videooptimizationparameter

The videooptimizationparameter data source allows you to retrieve the global video optimization parameters configured on the Citrix ADC. Because this is a singleton resource, no lookup attribute is required — there is exactly one instance per ADC.


## Example usage

```hcl
data "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
}

output "randomsamplingpercentage" {
  value = data.citrixadc_videooptimizationparameter.tf_videooptimizationparameter.randomsamplingpercentage
}
```


## Argument Reference

This data source takes no arguments. It is a singleton, so the single global instance is always returned.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the videooptimizationparameter resource. It is the static string `videooptimizationparameter-config`.
* `randomsamplingpercentage` - Random sampling percentage applied to video traffic for optimization decisions.
* `quicpacingrate` - QUIC video pacing rate, in Kbps.
