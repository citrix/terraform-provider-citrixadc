---
subcategory: "VideoOptimization"
---

# Resource: videooptimizationparameter

Configures the global video optimization parameters on the Citrix ADC. These settings control how the ADC samples traffic for video optimization and the pacing rate applied to QUIC video streams. This is a singleton resource: a single instance exists per ADC, so creating it tunes the global parameters rather than provisioning a new object.


## Example usage

```hcl
resource "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
  randomsamplingpercentage = 25
  quicpacingrate           = 1024
}
```


## Argument Reference

* `randomsamplingpercentage` - (Optional, Computed) Random sampling percentage applied to video traffic for optimization decisions. Minimum value = `0` Maximum value = `100`. Defaults to `0`.
* `quicpacingrate` - (Optional, Computed) QUIC video pacing rate, in Kbps.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the videooptimizationparameter resource. Because this is a singleton, the ID is the static string `videooptimizationparameter-config`.

~> **Note** This is a singleton resource — only one instance can exist per Citrix ADC. There is no delete operation: removing the resource from your configuration only removes it from Terraform state; the parameters remain on the ADC at their last-applied values.
