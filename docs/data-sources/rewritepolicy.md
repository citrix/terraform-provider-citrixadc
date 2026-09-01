---
subcategory: "Rewrite"
---

# Data Source: rewritepolicy

This data source retrieves information about a specific rewrite policy.

## Example Usage

```hcl
data "citrixadc_rewritepolicy" "example" {
  name = "my_rewrite_policy"
}

output "policy_action" {
  value = data.citrixadc_rewritepolicy.example.action
}
```

## Argument Reference

* `name` - (Required) Name of the rewrite policy.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the rewrite policy.
* `action` - Name of the rewrite action to perform if the request or response matches this rewrite policy. Built-in actions include: NOREWRITE, RESET, DROP.
* `rule` - Expression against which traffic is evaluated.
* `comment` - Any comments to preserve information about this rewrite policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF).
* `newname` - New name for the rewrite policy.

### Read-only rewritepolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_rewritepolicy` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `hits` - Number of hits.
* `undefhits` - Number of Undef hits.
* `description` - Description of the policy.
* `isdefault` - A value of true is returned if it is a default rewritepolicy.
* `builtin` - Flag to determine if the rewrite policy is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this configuration.
