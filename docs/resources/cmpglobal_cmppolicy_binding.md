---
subcategory: "Compression"
---

# Resource: cmpglobal_cmppolicy_binding

The cmpglobal_cmppolicy_binding resource is used to create cmpglobal_cmppolicy_binding.


## Example usage

```hcl
resource "citrixadc_cmpglobal_cmppolicy_binding" "tf_cmpglobal_cmppolicy_binding" {
  globalbindtype = "SYSTEM_GLOBAL"
  priority   = 50
  policyname =citrixadc_cmppolicy.tf_cmppolicy.name
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}
```


## Argument Reference

* `policyname` - (Required) The name of the globally bound HTTP compression policy.
* `priority` - (Required) Positive integer specifying the priority of the policy. The lower the number, the higher the priority. By default, polices within a label are evaluated in the order of their priority numbers. In the configuration utility, you can click the Priority field and edit the priority level or drag the entry to a new position in the list. If you drag the entry to a new position, the priority level is updated automatically.
* `state` - (Optional) The current state of the policy binding. This attribute is relevant only for CLASSIC policies. Possible values: [ ENABLED, DISABLED ]
* `type` - (Optional) Bind point to which the policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, HTTPQUIC_REQ_OVERRIDE, HTTPQUIC_REQ_DEFAULT, HTTPQUIC_RES_OVERRIDE, HTTPQUIC_RES_DEFAULT, NONE ]
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `gotopriorityexpression` - (Optional) Expression or other value specifying the priority of the next policy, within the policy label, to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher numbered priority. * END - Stop evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, it's evaluation result determines the next policy to evaluate, as follows:  * If the expression evaluates to a higher numbered priority, that policy is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher priority number is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest priority number, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only for default-syntax policies.
* `labeltype` - (Optional) Type of policy label invocation. This argument is relevant only for advanced (default-syntax) policies. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE. Applicable only to advanced (default-syntax) policies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmpglobal_cmppolicy_binding. It has the same value as the `policyname` attribute.


## Import

A cmpglobal_cmppolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding tf_cmppolicy
```
