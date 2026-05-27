---
subcategory: "Application Firewall"
---

# Resource: appfwprotofile

The appfwprotofile resource is used to import a gRPC schema (proto) file into the Citrix ADC Application Firewall.

Every attribute is marked `RequiresReplace`. Any change to `name`, `src`, `comment`, or `overwrite` forces Terraform to destroy and recreate the resource, since the NITRO API only exposes a POST `?action=Import` endpoint for this object and no compatible update path.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_protofile" {
  filename     = "sample.proto"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/sample.proto")
}

resource "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name       = "tf_appfwprotofile"
  src        = "local://sample.proto"
  comment    = "Imported via Terraform"
  overwrite  = true
  depends_on = [citrixadc_systemfile.tf_protofile]
}
```


## Argument Reference

* `name` - (Required) Name of the gRPC schema object.
* `src` - (Required) Indicates source path of the gRPC schema file. Typically a `local://` URL pointing to a file that has been uploaded to the ADC (for example via `citrixadc_systemfile`), or an HTTP(S) URL accessible from the ADC.
* `comment` - (Optional) Comments associated with this gRPC schema file. Write-only on the NITRO API: the value is sent to the ADC during import but is not echoed back on subsequent reads.
* `overwrite` - (Optional) Overwrite any existing gRPC schema object of the same name. Write-only on the NITRO API: the value is sent to the ADC during import but is not echoed back on subsequent reads.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprotofile. It has the same value as the `name` attribute.


## Import

An appfwprotofile can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwprotofile.tf_appfwprotofile tf_appfwprotofile
```
