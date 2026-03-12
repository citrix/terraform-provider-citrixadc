---
subcategory: "System"
---

# Data Source: systemcmdpolicy

The systemcmdpolicy data source allows you to retrieve information about command policies.

## Example usage

```terraform
data "citrixadc_systemcmdpolicy" "tf_policy" {
  policyname = "tf_policy"
}

output "action" {
  value = data.citrixadc_systemcmdpolicy.tf_policy.action
}

output "cmdspec" {
  value = data.citrixadc_systemcmdpolicy.tf_policy.cmdspec
}
```

## Argument Reference

* `policyname` - (Required) Name for a command policy. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to perform when a request matches the policy.
* `cmdspec` - Regular expression specifying the data that matches the policy.
* `id` - The id of the systemcmdpolicy. It has the same value as the `policyname` attribute.

## Import

A systemcmdpolicy can be imported using its policyname, e.g.

```shell
terraform import citrixadc_systemcmdpolicy.tf_policy tf_policy
```
