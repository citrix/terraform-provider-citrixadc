---
subcategory: "Application Firewall"
---

# Resource: appfwarchive_export

The `appfwarchive_export` resource exports an existing Application Firewall tar archive on the Citrix ADC to a target file path. It models the NITRO `appfwarchive ?action=export` action.

Note: This is a one-shot side-effect action. NITRO exposes no inverse API (no "un-export" and no delete-by-export-target), no update endpoint, and no GET endpoint that reports export state. As a result:

* `Read` is a no-op that preserves Terraform state.
* `Update` is a no-op; every attribute is `RequiresReplace`, so any plan change forces destroy + recreate, which triggers another export.
* `Delete` only removes the resource from Terraform state; the exported file on the ADC is not removed.

This resource is intentionally split from `citrixadc_appfwarchive` (which models `?action=Import`) because the two actions have incompatible payload requirements.


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


## Import

An `appfwarchive_export` resource can be imported using its id (the archive name), e.g.

```shell
terraform import citrixadc_appfwarchive_export.tf_appfwarchive_export tf_appfwarchive
```

Because NITRO has no GET endpoint for export state, an imported resource carries only the values present in Terraform configuration; the original `target` cannot be recovered from the ADC.
