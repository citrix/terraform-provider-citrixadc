---
subcategory: "Policy"
---

# Resource: policydataset

The policydataset resource is used to create policy data sets.


## Example usage

```hcl
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
}
```


## Argument Reference

* `name` - (Optional) Name of the dataset. Must not exceed 127 characters.
* `type` - (Optional) Type of value to bind to the dataset. Possible values: [ ipv4, number, ipv6, ulong, double, mac ]
* `indextype` - (Optional) Index type. Possible values: [ Auto-generated, User-defined ]
* `comment` - (Optional) Any comments to preserve information about this dataset or a data bound to this dataset.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policydataset. It has the same value as the `name` attribute.


## Import

A policydataset can be imported using its name, e.g.

```shell
terraform import citrixadc_policydataset.tf_dataset tf_dataset
```
