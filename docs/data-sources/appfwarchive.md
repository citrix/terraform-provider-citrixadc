---
subcategory: "Application Firewall"
---

# Data Source: appfwarchive

The appfwarchive data source allows you to retrieve information about an Application Firewall tar archive configured on the Citrix ADC.

~> **Note:** The NITRO GET response carries no per-archive identifying fields; the data source confirms reachability and surfaces the lookup name rather than reading archive contents back.


## Example usage

```hcl
# Look up an existing application firewall archive by name
data "citrixadc_appfwarchive" "example_appfwarchive" {
  name = citrixadc_appfwarchive.tf_appfwarchive.name
}

# Use the data source outputs
output "appfwarchive_id" {
  value = data.citrixadc_appfwarchive.example_appfwarchive.id
}

output "appfwarchive_name" {
  value = data.citrixadc_appfwarchive.example_appfwarchive.name
}
```


## Argument Reference

The following argument is required:

* `name` - (Required) Name of the tar archive to retrieve.


## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The id of the `appfwarchive`. It has the same value as the `name` attribute.
* `comment` - Comments associated with this archive.
* `src` - URL of the form `<protocol>://<host>[:<port>][/<path>]` indicating the source of the tar archive file that was imported.
* `target` - Path to the file to be exported (export-action attribute; populated only if it was set in configuration).
