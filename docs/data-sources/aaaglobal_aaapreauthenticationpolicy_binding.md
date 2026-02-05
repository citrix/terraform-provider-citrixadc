---
subcategory: "AAA"
---

# Data Source `aaaglobal_aaapreauthenticationpolicy_binding`

The aaaglobal_aaapreauthenticationpolicy_binding data source allows you to retrieve information about global AAA preauthentication policy bindings.


## Example usage

```terraform
data "citrixadc_aaaglobal_aaapreauthenticationpolicy_binding" "tf_binding" {
  policy = "tf_aaapreauthenticationpolicy"
}

output "priority" {
  value = data.citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `policy` - (Required) Name of the AAA preauthentication policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `priority` - The priority of the policy binding.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `type` - The type of the policy binding.

## Attribute Reference

* `id` - The id of the aaaglobal_aaapreauthenticationpolicy_binding. It has the same value as the `policy` attribute.


## Import

A aaaglobal_aaapreauthenticationpolicy_binding can be imported using its policy name, e.g.

```shell
terraform import citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_binding tf_aaapreauthenticationpolicy
```
