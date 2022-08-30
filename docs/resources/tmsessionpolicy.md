---
subcategory: "Traffic Management"
---

# Resource: tmsessionpolicy

The tmsessionpolicy resource is used to create tmsessionpolicy.


## Example usage

```hcl
resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
  name   = "my_tmsession_policy"
  rule   = "true"
  action = "tf_tmsessaction"
}
```


## Argument Reference

* `name` - (Required) Name for the session policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a session policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy'). Minimum length =  1
* `rule` - (Required) Expression, against which traffic is evaluated. Both classic and advance expressions are supported in default partition but only advance expressions in non-default partition. The following requirements apply only to the Citrix ADC CLI: *  If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Action to be applied to connections that match this policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmsessionpolicy. It has the same value as the `name` attribute.


## Import

A tmsessionpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsessionpolicy.tf_tmsessionpolicy my_tmsession_policy
```
