---
subcategory: "Ica"
---

# Resource: Icaaction

The Icaaction resource is used to create Icaaction.


## Example usage

```hcl
resource "citrixadc_icaaction" "tf_icaaction" {
  name              = "my_ica_action"
  accessprofilename = "default_ica_accessprofile"
}

```


## Argument Reference

* `name` - (Required) Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA action is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica action" or 'my ica action'). Minimum length =  1
* `accessprofilename` - (Optional) Name of the ica accessprofile to be associated with this action.
* `latencyprofilename` - (Optional) Name of the ica latencyprofile to be associated with this action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the Icaaction. It has the same value as the `name` attribute.


## Import

A Icaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_icaaction.tf_icaaction my_ica_action
```
