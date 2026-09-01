---
subcategory: "NS"
---

# Data Source: nsspparams

The nsspparams data source allows you to retrieve information about surge protection parameters.


## Example usage

```terraform
data "citrixadc_nsspparams" "my_nsspparams" {
}

output "basethreshold" {
  value = data.citrixadc_nsspparams.my_nsspparams.basethreshold
}

output "throttle" {
  value = data.citrixadc_nsspparams.my_nsspparams.throttle
}
```


## Argument Reference

This data source takes no arguments.

## Attribute Reference

The following attributes are available:

* `basethreshold` - Maximum number of server connections that can be opened before surge protection is activated.
* `throttle` - Rate at which the system opens connections to the server. Possible values: Normal, Aggressive, Relaxed.
* `id` - The id of the nsspparams resource.

### Read-only nsspparams metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nsspparams` resource) and are computed. Any attribute the appliance does not return is `null`.

* `table0` - Table. A list of strings.
* `builtin` - Flag to determine if sp param is built-in or not. A list of strings (possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).
* `feature` - The feature to be checked while applying this config.
