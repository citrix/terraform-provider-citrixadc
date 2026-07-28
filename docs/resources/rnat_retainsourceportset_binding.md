---
subcategory: "Network"
---

# Resource: rnat_retainsourceportset_binding

This resource is used to bind a source port range to an RNAT rule.


## Example usage

```hcl
resource "citrixadc_rnat" "tf_rnat" {
  name       = "rnat1"
  network    = "192.168.10.0"
  netmask    = "255.255.255.0"
  natip      = "10.0.0.10"
}

resource "citrixadc_rnat_retainsourceportset_binding" "tf_binding" {
  name                  = citrixadc_rnat.tf_rnat.name
  retainsourceportrange = "1024-2048"
}
```


## Argument Reference

* `name` - (Required) Name of the RNAT rule to which to bind the source port range. Changing this value forces a new resource to be created.
* `retainsourceportrange` - (Required) The source ports to retain, specified as a single port or a port range (for example, `"1024-2048"`). When this source port range is associated with the RNAT rule, Citrix ADC chooses a port from the specified range when establishing connections to the backend servers. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the rnat_retainsourceportset_binding resource. It is a comma-separated string of `key:value` pairs (URL-encoded), composed of the `name` and `retainsourceportrange` attributes — for example, `name:rnat1,retainsourceportrange:1024-2048`.


## Import

A rnat_retainsourceportset_binding can be imported using its ID, which is the concatenation of the `name` and `retainsourceportrange` attributes formatted as `key:value` pairs separated by a comma, e.g.

```shell
terraform import citrixadc_rnat_retainsourceportset_binding.tf_binding "name:rnat1,retainsourceportrange:1024-2048"
```
