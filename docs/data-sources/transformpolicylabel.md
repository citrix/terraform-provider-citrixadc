---
subcategory: "Transform"
---

# Data Source `transformpolicylabel`

The transformpolicylabel data source allows you to retrieve information about a URL Transformation policy label.


## Example usage

```terraform
data "citrixadc_transformpolicylabel" "transformpolicylabel" {
  labelname = "label_1"
}

output "policylabeltype" {
  value = data.citrixadc_transformpolicylabel.transformpolicylabel.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `newname` - New name for the policy label.
* `policylabeltype` - Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).

## Attribute Reference

* `id` - The id of the transformpolicylabel. It has the same value as the `labelname` attribute.


## Import

A transformpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_transformpolicylabel.transformpolicylabel label_1
```
