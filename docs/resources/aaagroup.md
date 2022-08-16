---
subcategory: "AAA"
---

# Resource: aaagroup

The aaagroup resource is used to create aaagroup.


## Example usage

```hcl
resource "citrixadc_aaagroup" "tf_aaagroup" {
  groupname = "my_group"
  weight    = 100
  loggedin  = true
}
```


## Argument Reference

* `groupname` - (Required) Name for the group. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore  characters. Cannot be changed after the group is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa group" or 'my aaa group'). Minimum length =  1
* `weight` - (Optional) Weight of this group with respect to other configured aaa groups (lower the number higher the weight). Minimum value =  0 Maximum value =  65535
* `loggedin` - (Optional) Display only the group members who are currently logged in. If there are large number of sessions, this command may provide partial details.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup. It has the same value as the `groupname` attribute.


## Import

A aaagroup can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup.tf_aaagroup my_group
```
