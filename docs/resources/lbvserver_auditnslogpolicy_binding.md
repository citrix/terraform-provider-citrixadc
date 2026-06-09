---
subcategory: "Load Balancing"
---

# Resource: lbvserver_auditnslogpolicy_binding

Binds an audit nslog policy to a load balancing virtual server so that traffic processed by the vserver is logged to an external nslog (syslog-style) audit server. Use this binding to enable per-vserver audit logging and to control, via priority and goto-priority expressions, the order in which audit policies are evaluated for that vserver.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  servicetype = "HTTP"
  ipv46       = "10.10.10.10"
  port        = 80
  lbmethod    = "ROUNDROBIN"
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
  name   = "tf_auditnslogpolicy"
  rule   = "true"
  action = "SETASLEARNNSLOG_ACT"
}

resource "citrixadc_lbvserver_auditnslogpolicy_binding" "tf_binding" {
  name       = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
  priority   = 100
}
```


## Argument Reference

* `name` - (Required) Name of the load balancing virtual server to which the audit nslog policy is bound. Changing this forces a new resource to be created.
* `policyname` - (Required) Name of the audit nslog policy bound to the LB vserver. Changing this forces a new resource to be created.
* `priority` - (Optional, Computed) Priority that determines the order in which this policy is evaluated relative to other policies bound to the vserver. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional, Computed) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT (evaluate the policy with the next higher priority number), END (end policy evaluation), USE_INVOCATION_RESULT (applicable when this policy invokes another policy label), or an expression that evaluates to a number identifying the next policy to evaluate. Changing this forces a new resource to be created.
* `invoke` - (Optional, Computed) Whether to invoke policies bound to a virtual server or policy label. When set to `true`, the `labeltype` and `labelname` attributes specify which label to invoke. Changing this forces a new resource to be created.
* `labeltype` - (Optional, Computed) Type of policy label to invoke. Applicable only when `invoke` is `true`. Available settings function as follows: `reqvserver` - evaluate the request against the request-based policies bound to the specified virtual server; `resvserver` - evaluate the response against the response-based policies bound to the specified virtual server; `policylabel` - invoke the request or response against the specified user-defined policy label. Changing this forces a new resource to be created.
* `labelname` - (Optional, Computed) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE. Applicable only when `invoke` is `true`. Changing this forces a new resource to be created.
* `order` - (Optional, Computed) Integer specifying the order of the service relative to the other services in the load balancing vserver's bindings. Note: This is a service-branch attribute that is not valid for a policy-name binding; the provider does not send it on create and only reads it back when the NITRO server echoes a value. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `lbvserver_auditnslogpolicy_binding`. It is a composite, comma-separated key of `key:value` pairs (URL-encoded): `name:<name>,policyname:<policyname>`.


## Import

A lbvserver_auditnslogpolicy_binding can be imported using its composite ID, which is the URL-encoded `name:<name>,policyname:<policyname>` key, e.g.

```shell
terraform import citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding "name:tf_lbvserver,policyname:tf_auditnslogpolicy"
```
