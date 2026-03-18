---
subcategory: "System"
---

# Data Source: systemglobal_authenticationradiuspolicy_binding

The systemglobal_authenticationradiuspolicy_binding data source allows you to retrieve information about the binding between the system global configuration and authentication RADIUS policies.


## Example Usage

```terraform
data "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_bind" {
  policyname = "tf_radiuspolicy"
}

output "priority" {
  value = data.citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_bind.priority
}

output "globalbindtype" {
  value = data.citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_bind.globalbindtype
}
```


## Argument Reference

* `policyname` - (Required) The name of the command policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation.
* `id` - The id of the systemglobal_authenticationradiuspolicy_binding. It is a system-generated identifier.
* `nextfactor` - On success invoke label. Applicable for advanced authentication policy binding.
* `priority` - The priority of the command policy.
