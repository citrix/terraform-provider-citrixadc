---
subcategory: "Compression"
---

# Data Source `cmppolicylabel`

The cmppolicylabel data source allows you to retrieve information about an existing cmppolicylabel.


## Example usage

```terraform
data "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
  labelname = "my_cmppolicy_label"
}

output "labelname" {
  value = data.citrixadc_cmppolicylabel.tf_cmppolicylabel.labelname
}

output "type" {
  value = data.citrixadc_cmppolicylabel.tf_cmppolicylabel.type
}
```


## Argument Reference

* `labelname` - (Required) Name of the HTTP compression policy label. Must begin with a letter, number, or the underscore character (_). Additional characters allowed, after the first character, are the hyphen (-), period (.) pound sign (#), space ( ), at sign (@), equals (=), and colon (:). The name must be unique within the list of policy labels for compression policies. Can be renamed after the policy label is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmppolicylabel. It has the same value as the `labelname` attribute.
* `newname` - New name for the compression policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `type` - Type of packets (request packets or response) against which to match the policies bound to this policy label.
