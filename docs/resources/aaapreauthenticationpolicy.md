---
subcategory: "AAA"
---

# Resource: aaapreauthenticationpolicy

The aaapreauthenticationpolicy resource is used to create aaapreauthenticationpolicy.


## Example usage

```hcl
resource "citrixadc_aaapreauthenticationpolicy" "tf_aaapreauthenticationpolicy" {
  name = "my_policy"
  rule = "REQ.VLANID == 5"
  reqaction = "my_action"
}
```


## Argument Reference

* `name` - (Required) Name for the preauthentication policy. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the preauthentication policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy'). Minimum length =  1
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, defining connections that match the policy.
* `reqaction` - (Required) Name of the action that the policy is to invoke when a connection matches the policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaapreauthenticationpolicy. It has the same value as the `name` attribute.


## Import

A aaapreauthenticationpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_aaapreauthenticationpolicy.tf_aaapreauthenticationpolicy my_policy
```
