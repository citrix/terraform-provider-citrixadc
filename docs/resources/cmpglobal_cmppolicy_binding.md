---
subcategory: "Compression"
---

# Resource: cmpglobal_cmppolicy_binding

The cmpglobal_cmppolicy_binding resource is used to bind an HTTP compression policy to the Citrix ADC global bind point.


## Example usage

```hcl
resource "citrixadc_cmpglobal_cmppolicy_binding" "tf_cmpglobal_cmppolicy_binding" {
  policyname     = citrixadc_cmppolicy.tf_cmppolicy.name
  priority       = 50
  type           = "RES_DEFAULT"
  globalbindtype = "SYSTEM_GLOBAL"
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
  name      = "tf_cmppolicy"
  rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
  resaction = "COMPRESS"
}
```


## Argument Reference

* `policyname` - (Required) The name of the globally bound HTTP compression policy.
* `priority` - (Required) Positive integer specifying the priority of the policy. The lower the number, the higher the priority. By default, polices within a label are evaluated in the order of their priority numbers.
* `type` - (Optional) Bind point to which the policy is bound.
* `globalbindtype` - (Optional) The global bind type. Defaults to `"SYSTEM_GLOBAL"`.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the priority of the next policy, within the policy label, to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher numbered priority. END - Stop evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number, whose result determines the next policy to evaluate.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `labeltype` - (Optional) Type of policy label invocation.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmpglobal_cmppolicy_binding. It is the concatenation of `policyname` and `type` attributes separated by a comma.


## Import

A cmpglobal_cmppolicy_binding can be imported using the concatenation of the `policyname` and `type` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding tf_cmppolicy,RES_DEFAULT
```
