---
subcategory: "Authentication"
---

# Data Source: authenticationpolicylabel_authenticationpolicy_binding

The authenticationpolicylabel_authenticationpolicy_binding data source allows you to retrieve information about a binding between an authentication policy label and an authentication policy.

## Example Usage

```terraform
data "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
  labelname  = "tf_authenticationpolicylabel"
  policyname = "tf_authenticationpolicy"
}

output "labelname" {
  value = data.citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind.labelname
}

output "priority" {
  value = data.citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind.priority
}
```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name of the authentication policy label to which to bind the policy.
* `policyname` - (Required) Name of the authentication policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpolicylabel_authenticationpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `nextfactor` - On success invoke label.
* `priority` - Specifies the priority of the policy.
