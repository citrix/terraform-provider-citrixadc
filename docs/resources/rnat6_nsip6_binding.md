---
subcategory: "Network"
---

# Resource: rnat6_nsip6_binding

The rnat6_nsip6_binding resource is used to create rnat6_nsip6_binding.


## Example usage

```hcl
resource "citrixadc_rnat6_nsip6_binding" "tf_rnat6_nsip6_binding" {
  name   = "my_rnat6"
  natip6 = "2001:db8:85a3::8a2e:370:7334"
}
```


## Argument Reference

* `name` - (Required) Name of the RNAT6 rule to which to bind NAT IPs.
* `natip6` - (Required) Nat IP Address.
* `ownergroup` - (Optional) The owner node group in a Cluster for this rnat rule. Defaults to `"DEFAULT_NG"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat6_nsip6_binding. It is the concatenation of the `name` and `natip6` attributes separated by a comma.


## Import

A rnat6_nsip6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_rnat6_nsip6_binding.tfrnat6_nsip6_binding my_rnat6,2001:db8:85a3::8a2e:370:7334
```
