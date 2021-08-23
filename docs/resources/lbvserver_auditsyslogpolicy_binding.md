---
subcategory: "Load Balancing"
---

# Resource: lbvserver_auditsyslogpolicy_binding

The lbvserver_auditsyslogpolicy_binding resource is used to bind load balancing virtual servers with audit syslog policies.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver1" {
  name        = "tf_lbvserver1"
  servicetype = "HTTP"
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction1" {
    name = "tf_syslogaction1"
    serverip = "10.124.67.92"
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy1" {
    name = "tf_syslogpolicy"
    rule = "true"
    action = citrixadc_auditsyslogaction.tf_syslogaction1.name

}

resource "citrixadc_lbvserver_auditsyslogpolicy_binding" "demo" {
    name = citrixadc_lbvserver.tf_lbvserver1.name
    policyname = citrixadc_auditsyslogpolicy.tf_syslogpolicy1.name
    priority = 100
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: * reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. * resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. * policylabel - invoke the request or response against the specified user-defined policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_auditsyslogpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.

## Import

A lbvserver_auditsyslogpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_auditsyslogpolicy_binding.tf_lbvserver_auditsyslogpolicy_binding tf_lbvserver_auditsyslogpolicy_binding
```
