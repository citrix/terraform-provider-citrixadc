---
subcategory: "SSL"
---

# Resource: sslpolicy

The sslpolicy resource is used to create SSL policies.


## Example usage

```hcl
resource "citrixadc_sslpolicy" "tf_sslpolicy" {
name   = "tf_sslpolicy"
rule   = "false"
action = citrixadc_sslaction.foo.name
}
```


## Argument Reference

* `name` - (Optional) Name for the new SSL policy.
* `rule` - (Optional) Expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the  character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `reqaction` - (Optional) The name of the action to be performed on the request. Refer to 'add ssl action' command to add a new action. Builtin actions like NOOP, RESET, DROP, CLIENTAUTH and NOCLIENTAUTH are also allowed.
* `action` - (Optional) Name of the built-in or user-defined action to perform on the request. Available built-in actions are NOOP, RESET, DROP, CLIENTAUTH, NOCLIENTAUTH, INTERCEPT AND BYPASS.
* `undefaction` - (Optional) Name of the action to be performed when the result of rule evaluation is undefined. Possible values for control policies: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, DROP. Possible values for data policies: NOOP, RESET, DROP and BYPASS.
* `comment` - (Optional) Any comments associated with this policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicy. It has the same value as the `name` attribute.


## Import

A sslpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_sslpolicy.tf_sslpolicy tf_sslpolicy
```
