---
subcategory: "System"
---

# Resource: systemgroup_systemuser_binding

This resource is used to bind a system user to a system group on the Citrix ADC.

~> **Note:** When using this resource to bind `systemuser` to a `systemgroup`, do not also define the `systemusers` attribute on the `systemgroup` resource.


## Example usage

```hcl
resource "citrixadc_systemgroup" "tf_systemgroup" {
  groupname    = "tf_systemgroup"
  timeout      = 999
  promptstring = "bye>"
}

resource "citrixadc_systemuser" "tf_user" {
  username = "tf_user"
  password = "tf_password"
  timeout  = 200
}

resource "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
  groupname = citrixadc_systemgroup.tf_systemgroup.groupname
  username  = citrixadc_systemuser.tf_user.username
}

```


## Argument Reference

* `username` - (Required) The system user.
* `groupname` - (Required) Name of the system group. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup_systemuser_binding. It is the concatenation of the `groupname` and `username` attributes separated by a comma.


## Import

A systemgroup_systemuser_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_systemgroup_systemuser_binding.tf_bind tf_systemgroup,tf_user
```
