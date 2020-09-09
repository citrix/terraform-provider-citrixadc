---
subcategory: "System"
---

# Resource: systemcmdpolicy

The systemcmdpolicy resource is used to create command policies.


## Example usage

```hcl
resource "citrixadc_systemcmdpolicy" "tf_policy" {
    policyname = "tf_policy"
    action = "DENY"
    cmdspec = "add.*"
}
```


## Argument Reference

* `policyname` - (Optional) Name for a command policy. 
* `action` - (Optional) Action to perform when a request matches the policy. Possible values: [ ALLOW, DENY ]
* `cmdspec` - (Optional) Regular expression specifying the data that matches the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemcmdpolicy. It has the same value as the `policyname` attribute.


## Import

A systemcmdpolicy can be imported using its policyname, e.g.

```shell
terraform import citrixadc_systemcmdpolicy.tf_policy tf_policy
```
