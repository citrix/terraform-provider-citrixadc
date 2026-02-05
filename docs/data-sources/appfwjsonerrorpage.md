---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwjsonerrorpage

The appfwjsonerrorpage data source allows you to retrieve information about a JSON error page configuration.


## Example usage

```terraform
data "citrixadc_appfwjsonerrorpage" "tf_appfwjsonerrorpage" {
  name = "tf_appfwjsonerrorpage"
}

output "appfwjsonerrorpage_name" {
  value = data.citrixadc_appfwjsonerrorpage.tf_appfwjsonerrorpage.name
}

output "appfwjsonerrorpage_comment" {
  value = data.citrixadc_appfwjsonerrorpage.tf_appfwjsonerrorpage.comment
}
```


## Argument Reference

* `name` - (Required) Indicates name of the imported json error page to be removed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about the JSON error object.
* `overwrite` - Overwrite any existing JSON error object of the same name.
* `src` - URL (protocol, host, path, and name) for the location at which to store the imported JSON error object. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.

## Attribute Reference

* `id` - The id of the appfwjsonerrorpage. It has the same value as the `name` attribute.


## Import

A appfwjsonerrorpage can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwjsonerrorpage.tf_appfwjsonerrorpage tf_appfwjsonerrorpage
```
