---
subcategory: "Application Firewall"
---

# Resource: appfwsignatures

The appfwsignatures resource is used to import appfw signature.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_signature" {
  filename     = "appfw_signatures.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfw_signatures.xml")
}
resource "citrixadc_appfwsignatures" "tf_appfwsignatures" {
  name       = "tf_appfwsignatures"
  src        = "local://appfw_signatures.xml"
  depends_on = [citrixadc_systemfile.tf_signature]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Name of the signature object.
* `comment` - (Optional) Any comments to preserve information about the signatures object.
* `merge` - (Optional) Merges the existing Signature with new signature rules
* `mergedefault` - (Optional) Merges signature file with default signature file.
* `overwrite` - (Optional) Overwrite any existing signatures object of the same name.
* `preservedefactions` - (Optional) preserves def actions of signature rules
* `sha1` - (Optional) File path for sha1 file to validate signature file
* `src` - (Optional) URL (protocol, host, path, and file name) for the location at which to store the imported signatures object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
* `vendortype` - (Optional) Third party vendor type for which WAF signatures has to be generated.
* `xslt` - (Optional) XSLT file source.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwsignatures. It has the same value as the `name` attribute.


## Import

A appfwsignatures can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwsignatures.tf_appfwsignatures tf_appfwsignatures
```
