---
subcategory: "System"
---

# Resource: systemuser_systemcmdpolicy_binding

The systemuser_systemcmdpolicy_binding resource is used to bind systemuser and systemcmdpolicy.

~>  If you are using this resource to bind systemcmdpolicy to a systemuser, do not define the `cmdpolicybinding` attribute in the systemuser resource.


## Example usage

```hcl

resource "citrixadc_systemuser" "tf_user" {
  username = "tf_user"
  password = "tf_password"
  timeout  = 200
}

resource "citrixadc_systemcmdpolicy" "tf_policy" {
  policyname = "tf_policy"
  action     = "DENY"
  cmdspec    = "add.*"
}

resource "citrixadc_systemuser_systemcmdpolicy_binding" "tf_bind" {
  username   = citrixadc_systemuser.tf_user.username
  policyname = citrixadc_systemcmdpolicy.tf_policy.policyname
  priority   = 100
}

```


## Argument Reference

* `policyname` - (Required) The name of command policy.
* `priority` - (Required) The priority of the policy.
* `username` - (Required) Name of the system-user entry to which to bind the command policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser_systemcmdpolicy_binding. It is the concatenation of the `username` and `policyname` attributes separated by a comma.


## Import

A systemuser_systemcmdpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_systemuser_systemcmdpolicy_binding.tf_bind tf_user,tf_policy
```
