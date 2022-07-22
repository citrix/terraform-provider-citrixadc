---
subcategory: "Cache Redirection"
---

# Resource: crvserver_icapolicy_binding

The crvserver_icapolicy_binding resource is used to create CR vserver ICA policy Binding.


## Example usage

```hcl
# Since the icapolicy resource is not yet available on Terraform,
# the tf_icapolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add ica action tf_icaaction -accessProfileName default_ica_accessprofile
# add ica policy tf_icapolicy -rule true -action tf_icaaction
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_icapolicy_binding" "crvserver_icapolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_icapolicy"
  priority   = 1
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE (valid only for default-syntax policies such as application firewall, transform, integrated cache, rewrite, responder, and content switching).
* `labelname` - (Optional) Name of the label to be invoked.
* `labeltype` - (Optional) Type of label to be invoked.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_icapolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_icapolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_icapolicy_binding.crvserver_icapolicy_binding my_v,tf_icapolicyserver
```
