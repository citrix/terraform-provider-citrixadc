---
subcategory: "Audit"
---

# Data Source: auditsyslogpolicy

The auditsyslogpolicy data source allows you to retrieve information about audit syslog policies.


## Example usage

```terraform
data "citrixadc_auditsyslogpolicy" "tf_syslogpolicy" {
  name = "tf_auditsyslogpolicy"
}

output "rule" {
  value = data.citrixadc_auditsyslogpolicy.tf_syslogpolicy.rule
}

output "action" {
  value = data.citrixadc_auditsyslogpolicy.tf_syslogpolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the policy. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog policy is added.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Syslog server action to perform when this policy matches traffic. NOTE: A syslog server action must be associated with a syslog audit policy.
* `rule` - Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.
* `id` - The id of the auditsyslogpolicy. It has the same value as the `name` attribute.

### Read-only auditsyslogpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_auditsyslogpolicy` resource) and are Computed-only. Any attribute the appliance does not return is `null`.

* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`. A list of strings.
* `feature` - The feature to be checked while applying this config.
* `expressiontype` - Type of policy (Classic/Advanced). Possible values = `Classic Policy`, `Advanced Policy`.
