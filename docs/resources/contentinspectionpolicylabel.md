---
subcategory: "CI"
---

# Resource: contentinspectionpolicylabel

The contentinspectionpolicylabel resource is used to create contentinspectionpolicylabel.


## Example usage

```hcl
resource "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
  labelname = "my_ci_policylabel"
  type      = "RES"
}

```


## Argument Reference

* `labelname` - (Required) Name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the contentInspection policy label is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my contentInspection policy label" or 'my contentInspection policy label').
* `type` - (Required) Type of packets (request or response packets) against which to match the policies bound to this policy label.
* `comment` - (Optional) Any comments to preserve information about this contentInspection policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionpolicylabel. It has the same value as the `labelname` attribute.


## Import

A contentinspectionpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel my_ci_policylabel
```
