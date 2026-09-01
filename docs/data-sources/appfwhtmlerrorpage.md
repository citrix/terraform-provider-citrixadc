---
subcategory: "Application Firewall"
---

# Data Source: appfwhtmlerrorpage

The appfwhtmlerrorpage data source allows you to retrieve information about an Application Firewall HTML error page.

## Example usage

```terraform
data "citrixadc_appfwhtmlerrorpage" "tf_appfwhtmlerrorpage" {
  name = "tf_appfwhtmlerrorpage"
}

output "appfwhtmlerrorpage_name" {
  value = data.citrixadc_appfwhtmlerrorpage.tf_appfwhtmlerrorpage.name
}

output "appfwhtmlerrorpage_src" {
  value = data.citrixadc_appfwhtmlerrorpage.tf_appfwhtmlerrorpage.src
}
```

## Argument Reference

* `name` - (Required) Name of the HTML error object to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwhtmlerrorpage. It has the same value as the `name` attribute.
* `comment` - Any comments to preserve information about the HTML error object.
* `overwrite` - Overwrite any existing HTML error object of the same name.
* `src` - URL (protocol, host, path, and name) for the location at which to store the imported HTML error object.

### Read-only appfwhtmlerrorpage metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwhtmlerrorpage` resource). They are Computed/GET-only and are `null` when the appliance omits them.

* `response` - Response returned by the appliance for the appfwhtmlerrorpage operation.

## Import

A appfwhtmlerrorpage can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwhtmlerrorpage.tf_appfwhtmlerrorpage tf_appfwhtmlerrorpage
```
