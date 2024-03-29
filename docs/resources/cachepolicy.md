---
subcategory: "Integrated Caching

"
---

# Resource: cachepolicy

The cachepolicy resource is used to create cachepolicy.


## Example usage

```hcl
resource "citrixadc_cachepolicy" "tf_cachepolicy" {
  policyname  = "my_cachepolicy"
  rule        = "true"
  action      = "CACHE"
  undefaction = "NOCACHE"
}
```


## Argument Reference

* `action` - (Required) Action to apply to content that matches the policy.  * CACHE or MAY_CACHE action - positive cachability policy * NOCACHE or MAY_NOCACHE action - negative cachability policy * INVAL action - Dynamic Invalidation Policy
* `rule` - (Required) Expression against which the traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `policyname` - (Required) Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.
* `invalgroups` - (Optional) Content group(s) to be invalidated when the INVAL action is applied. Maximum number of content groups that can be specified is 16.
* `invalobjects` - (Optional) Content groups(s) in which the objects will be invalidated if the action is INVAL.
* `storeingroup` - (Optional) Name of the content group in which to store the object when the final result of policy evaluation is CACHE. The content group must exist before being mentioned here. Use the "show cache contentgroup" command to view the list of existing content groups.
* `undefaction` - (Optional) Action to be performed when the result of rule evaluation is undefined.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachepolicy. It has the same value as the `policyname` attribute.


## Import

A cachepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_cachepolicy.tf_cachepolicy my_cachepolicy
```
