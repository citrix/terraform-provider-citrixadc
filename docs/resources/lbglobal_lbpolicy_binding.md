---
subcategory: "Load Balancing"
---

# Resource: lbglobal_lbpolicy_binding

This resource is used to bind an LB policy to the global LB bind point.


## Example usage

```hcl
resource "citrixadc_lbaction" "lb_action" {
  name = "lbact1"
  type = "NOLBACTION"
}

resource "citrixadc_lbpolicy" "lb_policy" {
  name   = "lbpol1"
  rule   = "client.IP.SRC.IN_SUBNET(10.0.0.0/8)"
  action = citrixadc_lbaction.lb_action.name
}

resource "citrixadc_lbglobal_lbpolicy_binding" "lb_global_bind" {
  policyname = citrixadc_lbpolicy.lb_policy.name
  priority   = 100
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the LB policy.
* `priority` - (Required) Specifies the priority of the policy. A lower number is evaluated before a higher number.
* `type` - (Optional) The bind point at which the policy is evaluated globally. Forms part of the resource ID together with `policyname`. Possible values = REQ_OVERRIDE, REQ_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: `vserver` - Invokes the unnamed policy label associated with the specified virtual server; `policylabel` - Invoke the specified policy label. Possible values = vserver, policylabel


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lbglobal_lbpolicy_binding. It is a comma-separated list of `key:value` pairs (URL-encoded) composed of the `policyname` and `type` attributes — for example, `policyname:lbpol1,type:REQ_DEFAULT`.
* `globalbindtype` - Global bind type discriminator (defaults to `SYSTEM_GLOBAL`). It is not a configurable input.


## Import

A lbglobal_lbpolicy_binding can be imported using its id, in the format `policyname:<value>,type:<value>`, e.g.

```shell
terraform import citrixadc_lbglobal_lbpolicy_binding.lb_global_bind "policyname:lbpol1,type:REQ_DEFAULT"
```
