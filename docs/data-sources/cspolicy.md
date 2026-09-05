---
subcategory: "Content Switching"
---

# Data Source: cspolicy

The cspolicy data source allows you to retrieve information about content switching policies.


## Example usage

```terraform
data "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "my_cspolicy"
}

output "rule" {
  value = data.citrixadc_cspolicy.tf_cspolicy.rule
}

output "action" {
  value = data.citrixadc_cspolicy.tf_cspolicy.action
}
```


## Argument Reference

* `policyname` - (Required) Name for the content switching policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Content switching action that names the target load balancing virtual server to which the traffic is switched.
* `logaction` - The log action associated with the content switching policy.
* `newname` - The new name of the content switching policy.
* `rule` - Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: If the expression includes one or more spaces, enclose the entire expression in double quotation marks. If the expression itself includes double quotation marks, escape the quotations by using the  character. Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `id` - The id of the cspolicy. It has the same value as the `policyname` attribute.

### Read-only cspolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_cspolicy` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `vstype` - Virtual server type.
* `hits` - Total number of hits.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type (for example `reqvserver`, `resvserver`, `policylabel`).
* `activepolicy` - Indicates whether the policy is bound or not.
* `boundto` - Location where the policy is bound.
