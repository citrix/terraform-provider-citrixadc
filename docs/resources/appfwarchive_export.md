---
subcategory: "Application Firewall"
---

# Resource: appfwarchive_export

This resource is used to export an Application Firewall archive from the Citrix ADC to a target file path.


## Example usage

```hcl
# Import the archive first
resource "citrixadc_appfwarchive" "tf_appfwarchive" {
  name = "tf_appfwarchive"
  src  = "http://archive.example.com/appfw/tf_appfwarchive.tar"
}

# Then export it to a local path on the ADC
resource "citrixadc_appfwarchive_export" "tf_appfwarchive_export" {
  name   = citrixadc_appfwarchive.tf_appfwarchive.name
  target = "/var/tmp/tf_appfwarchive_exported.tar"
}
```


## Argument Reference

* `name` - (Required) Name of the tar archive to export. Forces replacement on change.
* `target` - (Required) Path to the file to which the archive is exported. Forces replacement on change.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwarchive_export`. It has the same value as the `name` attribute.
