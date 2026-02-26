---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_auditnslogpolicy_binding

The authenticationvserver_auditnslogpolicy_binding data source allows you to retrieve information about a specific binding between an authentication virtual server and an audit nslog policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_auditnslogpolicy_binding" "example" {
  name   = "tf_authenticationvserver"
  policy = "my_auditnslogpolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_auditnslogpolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_authenticationvserver_auditnslogpolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
* `id` - The id of the authenticationvserver_auditnslogpolicy_binding. It is a system-generated identifier.
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor.
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
