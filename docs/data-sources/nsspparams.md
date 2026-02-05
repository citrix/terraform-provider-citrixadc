---
subcategory: "NS"
---

# Data Source `nsspparams`

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
