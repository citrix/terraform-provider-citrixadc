---
subcategory: "TRANSFORM"
---

# Resource: transformpolicylabel

The transformpolicylabel resource is used to create Transform policylabel.


## Example usage

```hcl
resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
  labelname = "label_1"
  policylabeltype = "httpquic_req"
}
```


## Argument Reference

* `labelname` - (Required) Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy label is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).
* `policylabeltype` - (Required) Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformpolicylabel. It has the same value as the `labelname` attribute.


## Import

A transformpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_transformpolicylabel.transformpolicylabel label_1
```
