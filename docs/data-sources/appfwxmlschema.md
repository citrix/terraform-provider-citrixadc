---
subcategory: "Application Firewall"
---

# Data Source `appfwxmlschema`

The appfwxmlschema data source allows you to retrieve information about Application Firewall XML Schema objects.


## Example usage

```terraform
data "citrixadc_appfwxmlschema" "tf_appfwxmlschema" {
  name = "tf_appfwxmlschema"
}

output "src" {
  value = data.citrixadc_appfwxmlschema.tf_appfwxmlschema.src
}

output "comment" {
  value = data.citrixadc_appfwxmlschema.tf_appfwxmlschema.comment
}
```


## Argument Reference

* `name` - (Required) Name of the XML Schema object to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about the XML Schema object.
* `overwrite` - Overwrite any existing XML Schema object of the same name.
* `src` - URL (protocol, host, path, and file name) for the location at which to store the imported XML Schema. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.

## Attribute Reference

* `id` - The id of the appfwxmlschema. It has the same value as the `name` attribute.


## Import

A appfwxmlschema can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwxmlschema.tf_appfwxmlschema tf_appfwxmlschema
```
