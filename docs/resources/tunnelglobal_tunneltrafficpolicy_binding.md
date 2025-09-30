---
subcategory: "Tunnel"
---

# Resource: tunnelglobal_tunneltrafficpolicy_binding

The tunnelglobal_tunneltrafficpolicy_binding resource is used to create tunnelglobal_tunneltrafficpolicy_binding.


## Example usage

```hcl
resource "citrixadc_tunnelglobal_tunneltrafficpolicy_binding" "tf_tunnelglobal_tunneltrafficpolicy_binding" {
  priority   = 50
  policyname = "my_tunneltrafficpolicy"
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Policy name.
* `priority` - (Required) Priority.
* `type` - (Optional) Bind point to which the policy is bound. Possible values: [REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, NONE]
* `feature` - (Optional) The feature to be checked while applying this config
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `state` - (Optional) Current state of the binding. If the binding is enabled, the policy is active.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tunnelglobal_tunneltrafficpolicy_binding. It has the same value as the `name` attribute.


## Import

A tunnelglobal_tunneltrafficpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_tunnelglobal_tunneltrafficpolicy_binding.tf_tunnelglobal_tunneltrafficpolicy_binding my_tunneltrafficpolicy
```
