---
subcategory: "SSL"
---

# Resource: sslglobal_sslpolicy_binding

This resource is used to bind an SSL policy to the global SSL bind point.


## Example usage

```hcl
resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "client.ssl.is_ssl"
  action = "NOOP"
}

resource "citrixadc_sslglobal_sslpolicy_binding" "tf_binding" {
  policyname             = citrixadc_sslpolicy.tf_sslpolicy.name
  priority               = 100
  type                   = "CONTROL_OVERRIDE"
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `policyname` - (Required) The name for the SSL policy. Changing this attribute forces a new resource to be created.
* `priority` - (Optional) The priority of the policy binding. Changing this attribute forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: `NEXT` (evaluate the policy with the next higher priority number), `END` (end policy evaluation), `USE_INVOCATION_RESULT` (applicable if this policy invokes another policy label), or an expression that evaluates to a number. Defaults to `"END"`. Changing this attribute forces a new resource to be created.
* `type` - (Optional) Global bind point to which the policy is bound. Changing this attribute forces a new resource to be created.
* `invoke` - (Optional) Invoke policies bound to a virtual server, service, or policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Changing this attribute forces a new resource to be created.
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE. Only meaningful when `invoke` is set. Changing this attribute forces a new resource to be created.
* `labeltype` - (Optional) Type of policy label to invoke. Specify virtual server for a policy label associated with a virtual server, or policy label for a user-defined policy label. Only meaningful when `invoke` is set. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslglobal_sslpolicy_binding. It is the concatenation of the `policyname`, `type`, and `priority` attributes, formatted as `policyname:<policyname>,type:<type>,priority:<priority>`.
* `globalbindtype` - (Read-only) The global bind point type returned by the appliance.


## Import

A sslglobal_sslpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslglobal_sslpolicy_binding.tf_binding policyname:tf_sslpolicy,type:CONTROL_OVERRIDE,priority:100
```
