---
subcategory: "Cache Redirection"
---

# Data Source: crpolicy

The crpolicy data source allows you to retrieve information about cache redirection policies.


## Example usage

```terraform
data "citrixadc_crpolicy" "tf_crpolicy" {
  policyname = "my_crpolicy"
}

output "action" {
  value = data.citrixadc_crpolicy.tf_crpolicy.action
}

output "rule" {
  value = data.citrixadc_crpolicy.tf_crpolicy.rule
}
```


## Argument Reference

* `policyname` - (Required) Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the built-in cache redirection action: CACHE/ORIGIN.
* `logaction` - The log action associated with the cache redirection policy
* `newname` - The new name of the content switching policy.
* `rule` - Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `id` - The id of the crpolicy. It has the same value as the `policyname` attribute.

### Read-only crpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_crpolicy` resource) and are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `boundto` - Domain name.
* `vstype` - Virtual server type.
* `hits` - Total number of hits.
* `priority` - Priority of the bound policy.
* `activepolicy` - Indicates whether the policy is bound or not.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type (for example `reqvserver`, `resvserver`, `policylabel`).
* `builtin` - Flag to determine if the cr policy is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
* `isdefault` - A value of true is returned if it is a default cr policy.
