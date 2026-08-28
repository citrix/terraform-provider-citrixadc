---
subcategory: "System"
---

# Data Source: systemscalablemgmtthreads

The systemscalablemgmtthreads data source allows you to retrieve the state of the Scalable Management Threads feature.


## Example usage

```terraform
data "citrixadc_systemscalablemgmtthreads" "tf_systemscalablemgmtthreads" {
}

output "configuredstate" {
  value = data.citrixadc_systemscalablemgmtthreads.tf_systemscalablemgmtthreads.configuredstate
}

output "effectivestate" {
  value = data.citrixadc_systemscalablemgmtthreads.tf_systemscalablemgmtthreads.effectivestate
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. Minimum value = 0, Maximum value = 31.

## Attribute Reference

The following attributes are available:

* `id` - The ID of the systemscalablemgmtthreads resource.
* `configuredstate` - The configured state of the Scalable Management Threads feature. Possible values: [ ENABLED, DISABLED ]
* `effectivestate` - The current running state of the Scalable Management Threads feature. Possible values: [ ENABLED, DISABLED ]
