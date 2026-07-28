---
subcategory: "Network"
---

# Data Source: rnat_retainsourceportset_binding

The rnat_retainsourceportset_binding data source allows you to retrieve information about a source port range that is bound to an RNAT (Reverse NAT) rule.


## Example usage

```terraform
data "citrixadc_rnat_retainsourceportset_binding" "tf_binding" {
  name                  = "rnat1"
  retainsourceportrange = "1024-2048"
}

output "bound_retainsourceportrange" {
  value = data.citrixadc_rnat_retainsourceportset_binding.tf_binding.retainsourceportrange
}
```


## Argument Reference

* `name` - (Required) Name of the RNAT rule whose binding you want to look up.
* `retainsourceportrange` - (Required) The source port range bound to the specified RNAT rule, specified as a single port or a port range (for example, `"1024-2048"`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the rnat_retainsourceportset_binding resource. It is a comma-separated string of `key:value` pairs (URL-encoded), composed of the `name` and `retainsourceportrange` attributes — for example, `name:rnat1,retainsourceportrange:1024-2048`.
* `name` - Name of the RNAT rule.
* `retainsourceportrange` - The source port range retained for the RNAT rule.
