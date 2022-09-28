---
subcategory: "Lsn"
---

# Resource: lsnhttphdrlogprofile

The lsnhttphdrlogprofile resource is used to create lsnhttphdrlogprofile.


## Example usage

```hcl
resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
  httphdrlogprofilename = "my_lsn_httphdrlogprofile"
  logurl                = "DISABLED"
  logversion            = "DISABLED"
  loghost               = "DISABLED"
}
```


## Argument Reference

* `httphdrlogprofilename` - (Required) The name of the HTTP header logging Profile.
* `loghost` - (Optional) Host information is logged if option is enabled.
* `logmethod` - (Optional) HTTP method information is logged if option is enabled.
* `logurl` - (Optional) URL information is logged if option is enabled.
* `logversion` - (Optional) Version information is logged if option is enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnhttphdrlogprofile. It has the same value as the `name` attribute.


## Import

A lsnhttphdrlogprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile my_lsn_httphdrlogprofile
```
