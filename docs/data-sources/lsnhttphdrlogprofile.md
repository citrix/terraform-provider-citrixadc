---
subcategory: "Lsn"
---

# Data Source: citrixadc_lsnhttphdrlogprofile

The lsnhttphdrlogprofile data source allows you to retrieve information about an LSN HTTP header logging profile.

## Example usage

```terraform
data "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
  httphdrlogprofilename = "my_lsn_httphdrlogprofile"
}

output "logurl" {
  value = data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.logurl
}

output "logversion" {
  value = data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.logversion
}

output "loghost" {
  value = data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile.loghost
}
```

## Argument Reference

* `httphdrlogprofilename` - (Required) The name of the HTTP header logging Profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnhttphdrlogprofile. It has the same value as the `httphdrlogprofilename` attribute.
* `loghost` - Host information is logged if option is enabled.
* `logmethod` - HTTP method information is logged if option is enabled.
* `logurl` - URL information is logged if option is enabled.
* `logversion` - Version information is logged if option is enabled.

## Import

A lsnhttphdrlogprofile can be imported using its httphdrlogprofilename, e.g.

```shell
terraform import citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile my_lsn_httphdrlogprofile
```
