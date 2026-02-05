---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwwsdl

The `citrixadc_appfwwsdl` data source is used to retrieve information about an existing Application Firewall WSDL file configured on the Citrix ADC.

## Example usage

```hcl

# Use the data source to retrieve the WSDL
data "citrixadc_appfwwsdl" "example_appfwwsdl" {
  name = citrixadc_appfwwsdl.example_appfwwsdl.name
}

# Use the data source outputs
output "appfwwsdl_id" {
  value = data.citrixadc_appfwwsdl.example_appfwwsdl.id
}

output "appfwwsdl_name" {
  value = data.citrixadc_appfwwsdl.example_appfwwsdl.name
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the WSDL file to retrieve.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The id of the application firewall WSDL. It has the same value as the `name` attribute.
* `src` - URL (protocol, host, path, and name) of the WSDL file that was imported.
* `comment` - Any comments to preserve information about the WSDL.
* `overwrite` - Whether to overwrite any existing WSDL of the same name.
