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

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number, in which case the number to which it evaluates determines the next policy to evaluate.
* `bindpoint` - (Optional) Bind point to which to bind the policy.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. policylabel - invoke the request or response against the specified user-defined policy label.
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_auditsyslogpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.

## Import

A lbvserver_auditsyslogpolicy_binding can be imported using its id, which is the concatenation of the `name` and `policyname` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_lbvserver_auditsyslogpolicy_binding.demo tf_lbvserver1,tf_syslogpolicy
```
