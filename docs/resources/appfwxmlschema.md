---
subcategory: "Application Firewall"
---

# Resource: appfwxmlschema

The appfwxmlschema resource is used to Configure XML schema resource.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_xmlschema" {
  filename     = "appfwxmlschema.xml"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/appfwxmlschema.xml")
}
resource "citrixadc_appfwxmlschema" "tf_appfwxmlschema" {
  name       = "tf_appfwxmlschema"
  src        = "local://appfwxmlschema.xml"
  depends_on = [citrixadc_systemfile.tf_xmlschema]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Name of the XML Schema object to remove. Minimum length =  1
* `src` - (Required) URL (protocol, host, path, and file name) for the location at which to store the imported XML Schema. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Minimum length =  1 Maximum length =  2047
* `comment` - (Optional) Any comments to preserve information about the XML Schema object. Maximum length =  128
* `overwrite` - (Optional) Overwrite any existing XML Schema object of the same name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwxmlschema. It has the same value as the `name` attribute.


## Import

A appfwxmlschema can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwxmlschema.tf_appfwxmlschema tf_appfwxmlschema
```
