---
subcategory: "Content Switching"
---

# Resource: cspolicylabel\_cspolicy\_binding

This resource is used to bind a content switching policy to a content switching policy label.


## Example usage

```hcl
resource "citrixadc_cspolicylabel" "tf_cspolicylabel" {
  labelname    = "tf_cspolicylabel"
  cspolicylabeltype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "tf_cspolicy"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicylabel_cspolicy_binding" "tf_binding" {
  labelname     = citrixadc_cspolicylabel.tf_cspolicylabel.labelname
  policyname    = citrixadc_cspolicy.tf_cspolicy.policyname
  priority      = 100
  targetvserver = citrixadc_lbvserver.tf_lbvserver.name
}
```


## Argument Reference

* `labelname` - (Required) Name of the policy label to which to bind a content switching policy. Changing this forces a new resource to be created.
* `policyname` - (Required) Name of the content switching policy. Changing this forces a new resource to be created.
* `priority` - (Required) Specifies the priority of the policy. Changing this forces a new resource to be created.
* `targetvserver` - (Optional) Name of the virtual server to which to forward requests that match the policy. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.
* `invoke` - (Optional) Invoke a policy label if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.
* `labeltype` - (Optional) Type of policy label invocation. Changing this forces a new resource to be created.
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.

Every attribute is replace-only, so changing any of them recreates the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the cspolicylabel\_cspolicy\_binding resource. It is a comma-separated list of `key:value` pairs built from the unique attributes, in the form `labelname:<labelname>,policyname:<policyname>` (values are URL-encoded).


## Import

A cspolicylabel\_cspolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_cspolicylabel_cspolicy_binding.tf_binding labelname:tf_cspolicylabel,policyname:tf_cspolicy
```
