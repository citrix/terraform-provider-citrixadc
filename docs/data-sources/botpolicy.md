---
subcategory: "Bot"
---

# Data Source: citrixadc_botpolicy

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

## Common Use Cases

### Retrieve Policy for Load Balancing Virtual Server Binding

```hcl
data "citrixadc_botpolicy" "bot_protection" {
  name = "bot_protection_policy"
}

resource "citrixadc_lbvserver_botpolicy_binding" "lb_binding" {
  name       = citrixadc_lbvserver.app.name
  policyname = data.citrixadc_botpolicy.bot_protection.name
  bindpoint  = "REQUEST"
  priority   = 100
}
```

### Reference Policy in Multiple Bindings

```hcl
data "citrixadc_botpolicy" "common_bot_policy" {
  name = "common_bot_protection"
}

resource "citrixadc_lbvserver_botpolicy_binding" "lb_binding" {
  name       = citrixadc_lbvserver.app.name
  policyname = data.citrixadc_botpolicy.common_bot_policy.name
  bindpoint  = "REQUEST"
  priority   = 100
}

resource "citrixadc_csvserver_botpolicy_binding" "cs_binding" {
  name       = citrixadc_csvserver.app.name
  policyname = data.citrixadc_botpolicy.common_bot_policy.name
  bindpoint  = "REQUEST"
  priority   = 100
}
```

### Conditional Bot Protection

```hcl
data "citrixadc_botpolicy" "bot_policy" {
  name = "production_bot_policy"
}

# Output policy details for verification
output "bot_policy_rule" {
  value       = data.citrixadc_botpolicy.bot_policy.rule
  description = "The rule expression for bot policy evaluation"
}

output "bot_profile_used" {
  value       = data.citrixadc_botpolicy.bot_policy.profilename
  description = "The bot profile associated with this policy"
}

# Use in conditional binding
resource "citrixadc_lbvserver_botpolicy_binding" "conditional" {
  count      = var.enable_bot_protection ? 1 : 0
  name       = citrixadc_lbvserver.app.name
  policyname = data.citrixadc_botpolicy.bot_policy.name
  bindpoint  = "REQUEST"
  priority   = 100
}
```

## Notes

* The bot policy must already exist in the Citrix ADC configuration before it can be retrieved using this data source.
* Bot policies are typically bound to virtual servers to provide bot management and protection capabilities.
* The `rule` attribute contains the expression used to match requests for bot policy evaluation.
* Bot policies work in conjunction with bot profiles (specified by `profilename`) to define bot management behavior.
