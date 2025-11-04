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
* `dynamic` - (Optional) This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name "allow_test" it can be used dynamically as client.ip.src.equals_any("allow_" + http.req.url.path.get(1)). This cannot be used with default datasets.
* `patsetfile` - (Optional) File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policydataset. It has the same value as the `name` attribute.


## Import

A policydataset can be imported using its name, e.g.

```shell
terraform import citrixadc_policydataset.tf_dataset tf_dataset
```
