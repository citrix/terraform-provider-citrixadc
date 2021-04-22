---
subcategory: "Audit"
---

# Resource: auditsyslogpolicy

The resource is used to create audit syslog policies


## Example usage

```hcl
resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "ns_true"
    action = "tf_syslogaction"

    globalbinding {
        priority = 120
        feature = "SYSTEM"
        globalbindtype = "SYSTEM_GLOBAL"
    }
}
```


## Argument Reference

* `name` - (Required) Name for the policy.
* `rule` - (Optional) Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.
* `action` - (Optional) Syslog server action to perform when this policy matches traffic.
* `globalbinding` - (Optional) A single global binding block. Documented below.

Global binding supports the following:

* `priority` - (Optional) The priority of the command policy.
* `globalbindtype` - (Optional) Possible values: [ SYSTEM\_GLOBAL, VPN\_GLOBAL, RNAT\_GLOBAL ]
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
    Specify one of the following values:
    * NEXT - Evaluate the policy with the next higher priority number.
    * END - End policy evaluation.
* `feature` - (Optional) The feature to be checked while applying this config.


## Attribute Reference


In addition to the arguments, the following attributes are available:

* `id` - The id of the policy. It has the same value as the `name` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy tf_auditsyslogpolicy
```
