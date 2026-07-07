---
subcategory: "SSL VPN"
---

# Resource: vpnepaprofile

The vpnepaprofile resource (`citrixadc_vpnepaprofile`) configures an EPA (Endpoint Analysis) device profile on the Citrix ADC. The profile is defined from a device-profile data XML, supplied either inline via `data` or by referencing an uploaded file via `filename`.

-> **NOTE** vpnepaprofile has no in-place update: `name`, `filename`, and `data` are all RequiresReplace, so any change recreates the resource.


## Example usage

```hcl
resource "citrixadc_vpnepaprofile" "tf_vpnepaprofile" {
  name     = "tf_vpnepaprofile"
  filename = "tf_vpnepaprofile.xml"
}
```


## Argument Reference

* `name` - (Required) Name of the device profile. Minimum length = 1. Changing this attribute forces a new resource to be created.
* `filename` - (Optional) Filename of the device-profile data XML. Minimum length = 1. Changing this attribute forces a new resource to be created.
* `data` - (Optional) Device-profile data XML. Minimum length = 1. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnepaprofile. It has the same value as the `name` attribute.


## Import

A vpnepaprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnepaprofile.tf_vpnepaprofile tf_vpnepaprofile
```
