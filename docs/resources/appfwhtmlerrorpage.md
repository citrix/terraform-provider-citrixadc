---
subcategory: "Application Firewall"
---

# Resource: appfwhtmlerrorpage

The appfwhtmlerrorpage resource is used to import appfw htmlerror page.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_errorpage" {
  filename     = "appfwhtmlerrorpage.html"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwhtmlerrorpage.html")
}
resource "citrixadc_appfwhtmlerrorpage" "tf_appfwhtmlerrorpage" {
  name       = "tf_appfwhtmlerrorpage"
  src        = "local://appfwhtmlerrorpage.html"
  depends_on = [citrixadc_systemfile.tf_errorpage]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Name of the XML error object to remove.
* `comment` - (Optional) Any comments to preserve information about the HTML error object.
* `overwrite` - (Optional) Overwrite any existing HTML error object of the same name.
* `src` - (Required) URL (protocol, host, path, and name) for the location at which to store the imported HTML error object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwhtmlerrorpage. It has the same value as the `name` attribute.


## Import

A appfwhtmlerrorpage can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwhtmlerrorpage.tf_appfwhtmlerrorpage tf_appfwhtmlerrorpage
```
