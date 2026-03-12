---
subcategory: "System"
---

# Data Source: systemglobal_authenticationtacacspolicy_binding

The systemglobal_authenticationtacacspolicy_binding data source allows you to retrieve information about a specific global TACACS authentication policy binding to the system entity.

## Example Usage

```terraform
data "citrixadc_systemglobal_authenticationtacacspolicy_binding" "tf_binding" {
  policyname = "tf_tacacspolicy"
}

output "priority" {
  value = data.citrixadc_systemglobal_authenticationtacacspolicy_binding.tf_binding.priority
}

output "feature" {
  value = data.citrixadc_systemglobal_authenticationtacacspolicy_binding.tf_binding.feature
}
```

## Argument Reference

* `policyname` - (Required) The name of the command policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_authenticationtacacspolicy_binding. It has the same value as the `policyname` attribute.
* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - The global bind type. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:
  * NEXT - Evaluate the policy with the next higher priority number.
  * END - End policy evaluation.
* `nextfactor` - On success invoke label. Applicable for advanced authentication policy binding.
* `priority` - The priority of the command policy.
