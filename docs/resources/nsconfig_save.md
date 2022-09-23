---
subcategory: "NS"
---

# Resource: nsconfig_save

The nsconfig_save resource is used to apply the save operation for ns config.


## Example usage

```hcl
resource "citrixadc_nsconfig_save" "tf_ns_save" {
    all        = true
    timestamp  = "2020-03-24T12:37:06Z"
}

resource "citrixadc_nsconfig_save" "tf_save" {
    timestamp  = "2020-03-24T12:37:06Z"

    # Will not error when save is already in progress
    concurrent_save_ok = true

    # Set to non zero value to retry the save config operation
    # Will throw error if limit is surpassed
    concurrent_save_retries = 1

    # Time interval between save retries
    concurrent_save_interval = "10s"

    # Total timeout for all retries
    concurrent_save_timeout = "5m"
}
```


## Argument Reference

* `all` - (Optional) Use this option to do saveconfig for all partitions.
* `timestamp` - (Required) the timestamp of the operation. Can be any string. Used to force the operation again if all other attributes have remained the same.
* `concurrent_save_ok` - (Optional) Boolean value signifying if a concurrent save error should be toleratted. When set to `true` a process of retries will take place waiting for the resource to return no error.
* `concurrent_save_retries` - (Optional) Number of retries after which we throw an error for the concurrent save error code.
* `concurrent_save_timeout` - (Optional) Time period after which we throw an error for the concurrent save error code.
* `concurrent_save_interval` - (Optional) Time period between tries to save the resource when processing the save error workflow.
* `save_on_destroy` - (Optional) Boolean flag. If set to `true` then the save configuration operation will be applied during the destroy operation. Defaults to `false`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconfig_save. It has the same value as the `timestamp` attribute.
