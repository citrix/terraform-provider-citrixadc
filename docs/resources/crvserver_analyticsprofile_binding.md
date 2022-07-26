---
subcategory: "Cache Redirection"
---

# Resource: crvserver_analyticsprofile_binding

The crvserver_analyticsprofile_binding resource is used to create CRvserver Analyticsprofile Binding.


## Example usage

```hcl
# Since the analyticsprofile resource is not yet available on Terraform,
# the new_profile profile must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add analyticsprofile new_profile -type tcpinsight
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_analyticsprofile_binding" "crvserver_analyticsprofile_binding" {
  name             = citrixadc_crvserver.crvserver.name
  analyticsprofile = "new_profile"
}
```


## Argument Reference

* `analyticsprofile` - (Optional) Name of the analytics profile bound to the CR vserver.
* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_analyticsprofile_binding. It has the same value as the `name` attribute.


## Import

A crvserver_analyticsprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_analyticsprofile_binding.crvserver_analyticsprofile_binding my_vserver,new_profile
```
