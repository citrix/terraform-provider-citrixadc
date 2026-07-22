---
subcategory: "NS"
---

# Resource: nstimer_autoscalepolicy_binding

Binds an autoscale policy to an `nstimer` on the Citrix ADC. The timer drives periodic evaluation of the bound policy, so this binding is what actually schedules autoscale decisions: at every timer interval the appliance evaluates the policy's rule and, when the configured threshold is met across the sample window, triggers the policy's autoscale action (for example scaling a service group up or down).


## Example usage

```hcl
resource "citrixadc_nstimer" "tf_nstimer" {
  name     = "tf_autoscale_timer"
  interval = 60
  unit     = "SEC"
}

resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
  name   = "tf_autoscalepolicy"
  rule   = "SYS.VSERVER(\"vs1\").RESPTIME.GT(100)"
  action = "tf_autoscaleaction"
}

resource "citrixadc_nstimer_autoscalepolicy_binding" "tf_binding" {
  name       = citrixadc_nstimer.tf_nstimer.name
  policyname = citrixadc_autoscalepolicy.tf_autoscalepolicy.name
  priority   = 100
  samplesize = 5
  threshold  = 3
  vserver    = "vs1"
}
```


## Argument Reference

* `name` - (Required) Name of the timer (`nstimer`) to which the autoscale policy is bound. Changing this value forces a new resource to be created.
* `policyname` - (Required) Name of the autoscale (timer) policy to associate with the timer. Changing this value forces a new resource to be created.
* `priority` - (Required) Priority of the timer policy. Changing this value forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this value forces a new resource to be created.
* `samplesize` - (Optional) Denotes the sample size. A sample size of `x` means that the previous `(x - 1)` policy rule evaluation results and the current evaluation result are kept with the binding. For example, a sample size of 10 means the state of the previous 9 evaluations plus the current one is retained. Defaults to `3`. Changing this value forces a new resource to be created.
* `threshold` - (Optional) Denotes the threshold. If the policy rule evaluates to TRUE `threshold` number of times within `samplesize` evaluations, the corresponding action is taken. Its value must be less than or equal to the sample size value. Defaults to `3`. Changing this value forces a new resource to be created.
* `vserver` - (Optional) Name of the vserver which provides the context for the rule in the timer policy. When not specified it is treated as a Global Default context. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstimer_autoscalepolicy_binding. It is the concatenation of the `name` and `policyname` attributes (as `name:<name>,policyname:<policyname>`).


## Import

A nstimer_autoscalepolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nstimer_autoscalepolicy_binding.tf_binding "name:tf_autoscale_timer,policyname:tf_autoscalepolicy"
```
