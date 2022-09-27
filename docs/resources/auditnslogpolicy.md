---
subcategory: "Audit"
---

# Resource: auditnslogpolicy

The auditnslogpolicy resource is used to create auditnslogpolicy.


## Example usage

```hcl
resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
  name   = "my_auditnslogpolicy"
  rule   = "true"
  action = "SETASLEARNNSLOG_ACT"
}
```


## Argument Reference

* `action` - (Required) Nslog server action that is performed when this policy matches. NOTE: An nslog server action must be associated with an nslog audit policy.
* `name` - (Required) Name for the policy.  Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog policy is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my nslog policy" or 'my nslog policy').
* `rule` - (Required) Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the nslog server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogpolicy. It has the same value as the `name` attribute.


## Import

A auditnslogpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_auditnslogpolicy.tf_auditnslogpolicy my_auditnslogpolicy
```
