---
subcategory: "Application Firewall"
---

# Data Source: appfwarchive

The `citrixadc_appfwarchive` data source is used to retrieve information about an existing Application Firewall tar archive configured on the Citrix ADC.

Note: NITRO's `appfwarchive` GET response carries no per-archive identifying fields. The data source confirms the archive collection is reachable on the ADC and surfaces the lookup name and any optional values supplied in configuration; it does not read individual archive contents back from the ADC.


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
