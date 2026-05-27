---
subcategory: "API Definition"
---

# Resource: apiprofile

The apiprofile resource is used to create and manage Citrix ADC API profiles. An API profile groups configuration that controls how requests matching bound API specifications are processed.


## Example usage

```hcl
resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "my_apiprofile"
  apivisibility = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name of the API profile to add. This value cannot be changed after the resource is created; updating it forces resource replacement.
* `apivisibility` - (Optional) Enable/Disable the schema lookup for the requests/apispecs that are bounded to the API profile. The default value of this parameter is `DISABLED`. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the apiprofile. It has the same value as the `name` attribute.


## Import

An apiprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_apiprofile.tf_apiprofile my_apiprofile
```
