---
subcategory: "Application Firewall"
---

# Data Source: appfwxmlerrorpage

The appfwxmlerrorpage data source allows you to retrieve information about an existing appfwxmlerrorpage.


## Example usage

```terraform

data "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage" {
  name = citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage.name
}

output "name" {
  value = data.citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage.name
}

output "src" {
  value = data.citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage.src
}
```


## Argument Reference

* `name` - (Required) Indicates name of the imported xml error page to be removed.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwxmlerrorpage. It has the same value as the `name` attribute.
* `comment` - Any comments to preserve information about the XML error object.
* `overwrite` - Overwrite any existing XML error object of the same name.
* `src` - URL (protocol, host, path, and name) for the location at which to store the imported XML error object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.

### Read-only appfwxmlerrorpage metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwxmlerrorpage` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `response` - Response returned by the appliance for the imported XML error object.
