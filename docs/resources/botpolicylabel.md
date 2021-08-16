---
subcategory: "Bot"
---

# Resource: botpolicylabel

The botpolicylabel resource is used to create a user-defined bot policy label, to which you can bind policies.


## Example usage

```hcl
resource "citrixadc_botpolicylabel" "tf_botpolicylabel" {
        labelname = "tf_botpolicylabel"
}

```


## Argument Reference

* `labelname` - (Required) Name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder policy label" or my responder policy label').
* `comment` - (Optional) Any comments to preserve information about this bot policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botpolicylabel. It has the same value as the `labelname` attribute.


## Import

A botpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_botpolicylabel.tf_botpolicylabel tf_botpolicylabel
```
