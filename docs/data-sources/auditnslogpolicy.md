---
subcategory: "Audit"
---

# Data Source: citrixadc_auditnslogpolicy

Use this data source to retrieve information about an existing Audit NS Log Policy.

The `citrixadc_auditnslogpolicy` data source allows you to retrieve details of an audit nslog policy by its name. This is useful for referencing existing audit nslog policies in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing audit nslog policy
data "citrixadc_auditnslogpolicy" "example" {
  name = "my_nslog_policy"
}

# Use the retrieved policy data in a binding
resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "example_binding" {
  policyname     = data.citrixadc_auditnslogpolicy.example.name
  priority       = 100
  globalbindtype = "SYSTEM_GLOBAL"
}

# Reference policy attributes
output "policy_rule" {
  value = data.citrixadc_auditnslogpolicy.example.rule
}

output "policy_action" {
  value = data.citrixadc_auditnslogpolicy.example.action
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the audit nslog policy to retrieve. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the audit nslog policy (same as name).

* `action` - Nslog server action that is performed when this policy matches. NOTE: An nslog server action must be associated with an nslog audit policy.

* `rule` - Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the nslog server. The rule can be a simple expression like `true` or `false`, or a complex expression that evaluates request/response data.

## Common Use Cases

### Retrieve Policy for Binding

```hcl
data "citrixadc_auditnslogpolicy" "security_policy" {
  name = "security_audit_policy"
}

resource "citrixadc_csvserver_auditnslogpolicy_binding" "cs_binding" {
  name       = citrixadc_csvserver.main.name
  policyname = data.citrixadc_auditnslogpolicy.security_policy.name
  priority   = 100
}
```

### Reference Policy in Multiple Bindings

```hcl
data "citrixadc_auditnslogpolicy" "common_audit" {
  name = "common_audit_policy"
}

resource "citrixadc_lbvserver_auditnslogpolicy_binding" "lb_binding" {
  name       = citrixadc_lbvserver.app.name
  policyname = data.citrixadc_auditnslogpolicy.common_audit.name
  priority   = 100
}

resource "citrixadc_csvserver_auditnslogpolicy_binding" "cs_binding" {
  name       = citrixadc_csvserver.app.name
  policyname = data.citrixadc_auditnslogpolicy.common_audit.name
  priority   = 100
}
```

## Notes

* The policy must exist on the Citrix ADC before it can be retrieved using this data source.
* This data source is read-only and does not create or modify any resources on the Citrix ADC.
* Use the `citrixadc_auditnslogpolicy` resource if you need to manage the policy lifecycle with Terraform.
