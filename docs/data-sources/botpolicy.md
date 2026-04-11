---
subcategory: "Bot"
---

# Data Source: botpolicy

Use this data source to retrieve information about an existing Bot Policy.

The `citrixadc_botpolicy` data source allows you to retrieve details of a bot policy by its name. This is useful for referencing existing bot policies in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing bot policy
data "citrixadc_botpolicy" "example" {
  name = "my_bot_policy"
}

# Use the retrieved policy data in a binding
resource "citrixadc_lbvserver_botpolicy_binding" "example_binding" {
  name         = citrixadc_lbvserver.main.name
  policyname   = data.citrixadc_botpolicy.example.name
  bindpoint    = "REQUEST"
  priority     = 100
}

# Reference policy attributes
output "policy_rule" {
  value = data.citrixadc_botpolicy.example.rule
}

output "policy_profile" {
  value = data.citrixadc_botpolicy.example.profilename
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the bot policy is added.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the bot policy (same as name).

* `comment` - Any type of information about this bot policy.

* `logaction` - Name of the messagelog action to use for requests that match this policy.

* `newname` - New name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

* `profilename` - Name of the bot profile to apply if the request matches this bot policy.

* `rule` - Expression that the policy uses to determine whether to apply bot profile on the specified request.

* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition.
