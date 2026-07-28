---
subcategory: "Load Balancing"
---

# Resource: lbpolicylabel_lbpolicy_binding

This resource is used to manage the binding between an LB policy and an LB policy label.


## Example usage

```hcl
resource "citrixadc_lbaction" "tf_lbaction" {
  name  = "selectvserver1"
  type  = "SELECTIONORDER"
  value = ["1"]
}

resource "citrixadc_lbpolicy" "tf_lbpolicy" {
  name   = "lbpolicy1"
  rule   = "CLIENT.IP.SRC.IN_SUBNET(10.0.0.0/8)"
  action = citrixadc_lbaction.tf_lbaction.name
}

resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname = "lbpolabel1"
  policylabeltype = "HTTP"
}

resource "citrixadc_lbpolicylabel_lbpolicy_binding" "tf_binding" {
  labelname  = citrixadc_lbpolicylabel.tf_lbpolicylabel.labelname
  policyname = citrixadc_lbpolicy.tf_lbpolicy.name
  priority   = 100
}
```

To chain evaluation to another policy label when the bound policy matches, set the invoke block:

```hcl
resource "citrixadc_lbpolicylabel_lbpolicy_binding" "tf_binding_invoke" {
  labelname        = citrixadc_lbpolicylabel.tf_lbpolicylabel.labelname
  policyname       = citrixadc_lbpolicy.tf_lbpolicy.name
  priority         = 110
  invoke           = true
  labeltype        = "policylabel"
  invoke_labelname = "nextlabel1"
}
```


## Argument Reference

* `labelname` - (Required) Name for the LB policy label to which the policy is bound. Must begin with a letter, number, or the underscore character (`_`), and must contain only letters, numbers, and the hyphen (`-`), period (`.`), hash (`#`), space (` `), at (`@`), equals (`=`), colon (`:`), and underscore characters. Changing this forces a new resource to be created.
* `policyname` - (Required) Name of the LB policy to bind to the policy label. Changing this forces a new resource to be created.
* `priority` - (Required) Specifies the priority of the policy. The priority determines the order in which the bound policies are evaluated within the policy label. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. The Citrix ADC server assigns a default when omitted. Changing this forces a new resource to be created.
* `invoke` - (Optional) Boolean. If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label. When set, `labeltype` and `invoke_labelname` define the target to invoke. The Citrix ADC server assigns a default when omitted. Changing this forces a new resource to be created.
* `labeltype` - (Optional) Type of policy label to invoke. Applies only when `invoke` is `true`. The Citrix ADC server assigns a default when omitted. Changing this forces a new resource to be created. Available settings function as follows:
  * `vserver` - Invokes the unnamed policy label associated with the specified virtual server.
  * `policylabel` - Invoke a user-defined policy label.
* `invoke_labelname` - (Optional) Applies only when `invoke` is `true`. If `labeltype` is `policylabel`, name of the policy label to invoke; if `labeltype` is `reqvserver`, name of the virtual server. The Citrix ADC server assigns a default when omitted. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lbpolicylabel_lbpolicy_binding resource. It is a composite, comma-separated string of `key:value` pairs in the form `labelname:<labelname>,policyname:<policyname>` (each value is URL-encoded).


## Import

A lbpolicylabel_lbpolicy_binding can be imported using its composite ID, e.g.

```shell
terraform import citrixadc_lbpolicylabel_lbpolicy_binding.tf_binding labelname:lbpolabel1,policyname:lbpolicy1
```
