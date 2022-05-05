---
subcategory: "Application Firewall"
---

# Resource: appfwxmlerrorpage

The appfwxmlerrorpage resource is used to Configuration for xml error page resource.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_xmlerrorpage" {
  filename     = "appfwxmlerrorpage.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwxmlerrorpage.xml")
}
resource "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage" {
  name       = "tf_appfwxmlerrorpage"
  src        = "local://appfwxmlerrorpage.xml"
  depends_on = [citrixadc_systemfile.tf_xmlerrorpage]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Indicates name of the imported xml error page to be removed.
* `comment` - (Optional) Any comments to preserve information about the XML error object.
* `overwrite` - (Optional) Overwrite any existing XML error object of the same name.
* `src` - (Optional) URL (protocol, host, path, and name) for the location at which to store the imported XML error object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwxmlerrorpage. It has the same value as the `name` attribute.


## Import

A appfwxmlerrorpage can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage tf_appfwxmlerrorpage
```
