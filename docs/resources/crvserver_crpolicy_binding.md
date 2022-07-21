---
subcategory: "Cache Redirection"
---

# Resource: crvserver_crpolicy_binding

The crvserver_crpolicy_binding resource is used to create CRvserver CRpolicy Binding.


## Example usage

```hcl
resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_crvserver_crpolicy_binding" "crvserver_crpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_crpolicy.crpolicy.policyname
    priority = 10 
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_crpolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_crpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_crpolicy_binding.crvserver_crpolicy_binding my_vserver,crpolicy1
```
