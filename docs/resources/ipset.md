---
subcategory: "Network"
---

# Resource: ipset

The ipset resource is used to create ip set resources.


## Example usage

```hcl
resource "citrixadc_ipset" "tf_ipset" {
    name = "tf_ipset"
    nsipbinding = [
        "10.78.60.10",
        "10.78.50.10",
    ]
}
```


## Argument Reference

* `name` - (Optional) Name for the IP set. Cannot be changed after the IP set is created. Choose a name that helps identify the IP set.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `nsipbinding` - (Optional) A set of ipv4 addresses that will be bound to this ipset.
* `nsip6binding` - (Optional) A set of ipv6 addresses that will be bound to this ipset.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipset. It has the same value as the `name` attribute.


## Import

A ipset can be imported using its name, e.g.

```shell
terraform import citrixadc_ipset.tf_ipset tf_ipset
```
