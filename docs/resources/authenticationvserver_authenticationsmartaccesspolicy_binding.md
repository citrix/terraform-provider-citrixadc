---
subcategory: "Authentication"
---

# Resource: authenticationvserver_authenticationsmartaccesspolicy_binding

This resource is used to bind a SmartAccess authentication policy to an authentication virtual server.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authvserver" {
  name        = "authvserver1"
  servicetype = "SSL"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_smartaccesspolicy" {
  name   = "smartaccesspolicy1"
  rule   = "true"
  action = "smartaccessprofile1"
}

resource "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" "tf_binding" {
  name     = citrixadc_authenticationvserver.tf_authvserver.name
  policy   = citrixadc_authenticationsmartaccesspolicy.tf_smartaccesspolicy.name
  priority = 100
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy. Changing this forces a new resource to be created.
* `policy` - (Required) The name of the policy bound to the authentication vserver. Changing this forces a new resource to be created.
* `priority` - (Optional) The priority of the policy within the authentication vserver's policy chain. Lower numbers are evaluated first. If not specified, the ADC assigns a server-side default, which is then reflected in state. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT (evaluate the policy with the next higher priority number), END (end policy evaluation), USE_INVOCATION_RESULT (applicable if this policy invokes another policy label; if the final goto in the invoked policy label has a value of END, the evaluation stops, otherwise the current policy label performs a NEXT), or an expression that evaluates to a number. If not specified, the ADC assigns a server-side default, which is then reflected in state. Changing this forces a new resource to be created.
* `secondary` - (Optional) Applicable only while binding a classic authentication policy; advance authentication policy uses nFactor. Changing this forces a new resource to be created.
* `groupextraction` - (Optional) Applicable only while binding a classic authentication policy; advance authentication policy uses nFactor. Changing this forces a new resource to be created.
* `nextfactor` - (Optional) Applicable only while binding an advance authentication policy; classic authentication policy does not support nFactor. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_authenticationsmartaccesspolicy_binding. It is a composite of the `name` and `policy` attributes in the format `name:<name>,policy:<policy>` (the values are URL-encoded).


## Import

An authenticationvserver_authenticationsmartaccesspolicy_binding can be imported using its id, in the format `name:<name>,policy:<policy>`, e.g.

```shell
terraform import citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding name:authvserver1,policy:smartaccesspolicy1
```
