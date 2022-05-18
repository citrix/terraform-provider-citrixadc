---
subcategory: "Application Firewall"
---

# Resource: appfwwsdl

The appfwwsdl resource is used to create WSDL file resource.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_wsdl" {
  filename     = "sample.wsdl"
  filelocation = "/var/tmp"
  filecontent  = file("${path.module}/sample.wsdl")
}
resource "citrixadc_appfwwsdl" "tf_appfwwsdl" {
  name       = "tf_appfwwsdl"
  src        = "local://sample.wsdl"
  depends_on = [citrixadc_systemfile.tf_wsdl]
  comment    = "TestingExample"
}
```


## Argument Reference

* `name` - (Required) Name of the WSDL file to remove. Minimum length =  1
* `src` - (Optional) URL (protocol, host, path, and name) of the WSDL file to be imported is stored. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Minimum length =  1 Maximum length =  2047
* `comment` - (Optional) Any comments to preserve information about the WSDL. Maximum length =  128
* `overwrite` - (Optional) Overwrite any existing WSDL of the same name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwwsdl. It has the same value as the `name` attribute.


## Import

A appfwwsdl can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwwsdl.tf_appfwwsdl tf_appfwwsdl
```
