---
subcategory: "Load Balancing"
---

# Data Source: lbglobal_lbpolicy_binding

The lbglobal_lbpolicy_binding data source allows you to retrieve information about a load balancing policy that is bound to the global LB bind point on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_lbglobal_lbpolicy_binding" "lb_global_bind" {
  policyname = "lbpol1"
  type       = "REQ_DEFAULT"
}

output "lb_global_priority" {
  value = data.citrixadc_lbglobal_lbpolicy_binding.lb_global_bind.priority
}
```


## Argument Reference

* `policyname` - (Required) Name of the LB policy.
* `type` - (Required) The bind point at which the policy is evaluated globally. Possible values = REQ_OVERRIDE, REQ_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lbglobal_lbpolicy_binding. It is a comma-separated list of `key:value` pairs (URL-encoded) composed of the `policyname` and `type` attributes — for example, `policyname:lbpol1,type:REQ_DEFAULT`.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of invocation. Available settings function as follows: `vserver` - Invokes the unnamed policy label associated with the specified virtual server; `policylabel` - Invoke the specified policy label.
* `globalbindtype` - Server-managed global bind type discriminator echoed by the ADC (defaults to `SYSTEM_GLOBAL`).
