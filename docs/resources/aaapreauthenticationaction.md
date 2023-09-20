---
subcategory: "AAA"
---

# Resource: aaapreauthenticationpolicy

The citrixadc_aaapreauthenticationaction resource is used to create aaapreauthenticationpolicy.


## Example usage

```hcl
resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
  name                    = "tf_aaapreauthenticationaction"
  preauthenticationaction = "ALLOW"
  deletefiles             = "/var/tmp/new/hello.txt"
}
```


## Argument Reference

* `name` - (Required) Name for the preauthentication action. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after preauthentication action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa action" or 'my aaa action'). Minimum length =  1
* `preauthenticationaction` - (Optional) Allow or deny logon after endpoint analysis (EPA) results. Possible values: [ ALLOW, DENY ]
* `killprocess` - (Optional) String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool.
* `deletefiles` - (Optional) String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool.
* `defaultepagroup` - (Optional) This is the default group that is chosen when the EPA check succeeds. Maximum length =  64


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaapreauthenticationpolicy. It has the same value as the `name` attribute.


## Import

A aaapreauthenticationpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction tf_aaapreauthenticationaction
```
