---
subcategory: "Network"
---

# Data Source: ipset

The `ipset` data source allows you to retrieve information about an IP set configuration.

## Example usage

```terraform
data "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
}

output "name" {
  value = data.citrixadc_ipset.tf_ipset.name
}

output "td" {
  value = data.citrixadc_ipset.tf_ipset.td
}
```

## Argument Reference

* `name` - (Required) Name for the IP set. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the IP set is created. Choose a name that helps identify the IP set.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipset. It has the same value as the `name` attribute.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Import

A ipset can be imported using its name, e.g.

```shell
terraform import citrixadc_ipset.tf_ipset tf_test_ipset
```
