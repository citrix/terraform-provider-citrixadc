---
subcategory: "Policy"
---

# Data Source `policydataset`

The policydataset data source allows you to retrieve information about an existing policy dataset configuration.

## Example usage

```terraform
data "citrixadc_policydataset" "tf_dataset" {
  name = "my_dataset"
}

output "dataset_name" {
  value = data.citrixadc_policydataset.tf_dataset.name
}

output "dataset_type" {
  value = data.citrixadc_policydataset.tf_dataset.type
}
```

## Argument Reference

* `name` - (Required) Name of the dataset. Must not exceed 127 characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policydataset resource. It has the same value as the `name` attribute.
* `comment` - Any comments to preserve information about this dataset or a data bound to this dataset.
* `dynamic` - This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression.
* `dynamiconly` - Shows only dynamic datasets when set true.
* `patsetfile` - File which contains list of patterns that needs to be bound to the dataset.
* `type` - Type of value to bind to the dataset. Possible values: [ ipv4, number, ipv6, ulong, double, mac ]
