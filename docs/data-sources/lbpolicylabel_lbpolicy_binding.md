---
subcategory: "Load Balancing"
---

# Data Source: lbpolicylabel_lbpolicy_binding

The lbpolicylabel_lbpolicy_binding data source allows you to retrieve information about an LB policy bound to an LB policy label.


## Example usage

```terraform
data "citrixadc_lbpolicylabel_lbpolicy_binding" "example" {
  labelname  = "lbpolabel1"
  policyname = "lbpolicy1"
}

output "binding_priority" {
  value = data.citrixadc_lbpolicylabel_lbpolicy_binding.example.priority
}
```


## Argument Reference

* `labelname` - (Required) Name of the LB policy label whose binding you want to look up.
* `policyname` - (Required) Name of the LB policy bound to the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding. It is a composite, comma-separated string of `key:value` pairs in the form `labelname:<labelname>,policyname:<policyname>` (each value is URL-encoded).
* `priority` - Specifies the priority of the policy, which determines its evaluation order within the policy label.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Boolean. If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `labeltype` - Type of policy label to invoke. Possible values:
  * `vserver` - Invokes the unnamed policy label associated with the specified virtual server.
  * `policylabel` - Invoke a user-defined policy label.
* `invoke_labelname` - If `labeltype` is `policylabel`, name of the policy label to invoke; if `labeltype` is `reqvserver`, name of the virtual server.
