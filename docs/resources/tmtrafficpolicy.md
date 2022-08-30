---
subcategory: "Traffic Management"
---

# Resource: tmtrafficpolicy

The tmtrafficpolicy resource is used to create tmtrafficpolicy.


## Example usage

```hcl
resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
  name   = "my_tmtraffic_policy"
  rule   = "true"
  action = "my_tmtraffic_action"
}

```


## Argument Reference

* `name` - (Required) Name for the traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy'). Minimum length =  1
* `rule` - (Required) Name of the Citrix ADC named expression, or an expression, that the policy uses to determine whether to apply certain action on the current traffic.
* `action` - (Required) Name of the action to apply to requests or connections that match this policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmtrafficpolicy. It has the same value as the `name` attribute.


## Import

A tmtrafficpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy my_tmtraffic_policy
```
