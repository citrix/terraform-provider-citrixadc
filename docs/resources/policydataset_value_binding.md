---
subcategory: "Policy"
---

# Resource: policydataset_value_binding

The policydataset_value_binding resource is used to add values to a policydataset.


## Example usage

```hcl
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
}

resource "citrixadc_policydataset_value_binding" "tf_value2" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 200
}
```


## Argument Reference

* `value` - (Optional) Value of the specified type that is associated with the dataset.
* `index` - (Optional) The index of the value (ipv4, ipv6, number) associated with the set.
* `comment` - (Optional) Any comments to preserve information about this dataset or a data bound to this dataset.
* `name` - (Optional) Name of the dataset to which to bind the value.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policydataset_value_binding. Its value is the concatenation of the `name` and `value` attributes separated by a comma: `<name>,<value>`.
