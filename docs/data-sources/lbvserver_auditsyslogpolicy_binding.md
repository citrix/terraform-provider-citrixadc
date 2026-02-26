---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_auditsyslogpolicy_binding

The lbvserver_auditsyslogpolicy_binding data source allows you to retrieve information about the binding between a load balancing virtual server and an audit syslog policy.

## Example Usage

```terraform
data "citrixadc_lbvserver_auditsyslogpolicy_binding" "demo" {
  name       = "tf_lbvserver3"
  policyname = "tf_syslogpolicy2"
}

output "invoke" {
  value = data.citrixadc_lbvserver_auditsyslogpolicy_binding.demo.invoke
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_auditsyslogpolicy_binding.demo.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_auditsyslogpolicy_binding. It is a system-generated identifier.
* `priority` - Priority.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. policylabel - invoke the request or response against the specified user-defined policy label.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
