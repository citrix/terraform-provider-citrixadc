---
subcategory: "Application Firewall"
---

# Resource: appfwjsonerrorpage

The appfwjsonerrorpage resource is used to Configuration for JSON error page resource.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_jsonerrorpage" {
  filename     = "appfwjsonerrorpage.json"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwjsonerrorpage.json")
}
resource "citrixadc_appfwjsonerrorpage" "tf_appfwjsonerrorpage" {
  name       = "tf_appfwjsonerrorpage"
  src        = "local://appfwjsonerrorpage.json"
  depends_on = [citrixadc_systemfile.tf_jsonerrorpage]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Indicates name of the imported json error page to be removed.
* `comment` - (Optional) Any comments to preserve information about the JSON error object.
* `overwrite` - (Optional) Overwrite any existing JSON error object of the same name.
* `src` - (Optional) URL (protocol, host, path, and name) for the location at which to store the imported JSON error object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwjsonerrorpage. It has the same value as the `name` attribute.


## Import

A appfwjsonerrorpage can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwjsonerrorpage.tf_appfwjsonerrorpage tf_appfwjsonerrorpage
```
