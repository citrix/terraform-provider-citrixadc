---
subcategory: "Content Inspection"
---

# Data Source `contentinspectionpolicylabel`

The contentinspectionpolicylabel data source allows you to retrieve information about a Content Inspection policy label.


## Example usage

```terraform
data "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
  labelname = "my_ci_policylabel"
}

output "type" {
  value = data.citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.type
}

output "comment" {
  value = data.citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.comment
}
```


## Argument Reference

* `labelname` - (Required) Name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this contentInspection policy label.
* `type` - Type of packets (request or response packets) against which to match the policies bound to this policy label.
* `newname` - New name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

* `id` - The id of the contentinspectionpolicylabel. It has the same value as the `labelname` attribute.


## Import

A contentinspectionpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel my_ci_policylabel
```
