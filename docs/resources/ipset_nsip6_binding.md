---
subcategory: "Network"
---

# Resource: ipset_nsip6_binding

The ipset_nsip6_binding resource is used to create ipset_nsip6_binding.


## Example usage

```hcl
resource "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
  name      = "tf_ipset"
  ipaddress = "2003:db8:100::fb/64"
}


```


## Argument Reference

* `ipaddress` - (Required) One or more IP addresses bound to the IP set. Minimum length =  1
* `name` - (Required) Name of the IP set to which to bind IP addresses. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipset_nsip6_binding. It has the same value as the `name` and `ipaddress` attributes separated by a comma.


## Import

A ipset_nsip6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding tf_ipset,2003:db8:100::fb/64
```
