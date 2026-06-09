---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_auditnslogpolicy_binding

The lbvserver_auditnslogpolicy_binding data source allows you to retrieve information about an audit nslog policy that is bound to a load balancing virtual server, including the priority and invocation settings of the binding.


## Example usage

```terraform
data "citrixadc_lbvserver_auditnslogpolicy_binding" "tf_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_auditnslogpolicy"
}

output "binding_priority" {
  value = data.citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding.priority
}
```


## Argument Reference

* `name` - (Required) Name of the load balancing virtual server whose audit nslog policy binding you want to look up.
* `policyname` - (Required) Name of the audit nslog policy bound to the LB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `lbvserver_auditnslogpolicy_binding`. It is a composite, comma-separated key of `key:value` pairs (URL-encoded): `name:<name>,policyname:<policyname>`.
* `priority` - Priority that determines the order in which this policy is evaluated relative to other policies bound to the vserver.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE (for example, NEXT, END, USE_INVOCATION_RESULT, or a numeric expression).
* `invoke` - Whether policies bound to a virtual server or policy label are invoked.
* `labeltype` - Type of policy label to invoke (`reqvserver`, `resvserver`, or `policylabel`). Applicable only when `invoke` is `true`.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE. Applicable only when `invoke` is `true`.
* `order` - Integer specifying the order of the service relative to the other services in the load balancing vserver's bindings.
