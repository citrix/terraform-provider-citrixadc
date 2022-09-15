---
subcategory: "NTP"
---

# Resource: ntpsync

The ntpsync resource is used to enable/disable ntpsync.


## Example usage

```hcl
resource "citrixadc_ntpsync" "tf_ntpsync" {
  state = "DISABLED"
}
```


## Argument Reference
* `state` - (Required) Enable or disable ntpsync. Possible values: ENABLED, DISABLED



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ntpsync. It is a unique string prefixed with `tf-ntpsync-`.
