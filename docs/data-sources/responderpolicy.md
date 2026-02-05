---
subcategory: "Responder"
---

# Data Source: citrixadc_responderpolicy

The `citrixadc_responderpolicy` data source is used to retrieve information about an existing responder policy configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_responderpolicy" "example" {
  name = "my_responderpolicy"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the responder policy to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the responder policy (same as name).
* `action` - Name of the responder action to perform if the request matches this responder policy. Can be a user-defined action or built-in actions: `NOOP`, `RESET`, `DROP`.
* `appflowaction` - AppFlow action to invoke for requests that match this policy.
* `comment` - Any type of information about this responder policy.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the responder policy (if renamed).
* `rule` - Expression that the policy uses to determine whether to respond to the specified request. Written in the classic or default syntax.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition.
