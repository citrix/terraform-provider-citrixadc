---
subcategory: "Content Switching"
---

# Resource: csvserver_auditnslogpolicy_binding

The csvserver_auditnslogpolicy_binding resource is used to bind an audit nslog policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
  priority   = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver"
  ipv46       = "10.202.11.11"
  port        = 8080
  servicetype = "HTTP"
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
  name   = "tf_auditnslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditnslogaction.tf_nslogaction.name
}

resource "citrixadc_auditnslogaction" "tf_nslogaction" {
  name       = "tf_nslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number, where the resulting number determines the next policy to evaluate.
* `bindpoint` - (Optional) Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labeltype` - (Optional) Type of label to be invoked.
* `labelname` - (Optional) Name of the label to be invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1. Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_auditnslogpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding tf_csvserver,tf_auditnslogpolicy
```
