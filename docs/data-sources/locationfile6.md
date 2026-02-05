---
subcategory: "Basic"
---

# Data Source: citrixadc_locationfile6

The citrixadc_locationfile6 data source allows you to retrieve information about the IPv6 location file configuration.

## Example usage

```terraform
data "citrixadc_locationfile6" "tf_locationfile6" {
}

output "locationfile" {
  value = data.citrixadc_locationfile6.tf_locationfile6.locationfile
}

output "format" {
  value = data.citrixadc_locationfile6.tf_locationfile6.format
}

output "src" {
  value = data.citrixadc_locationfile6.tf_locationfile6.src
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `format` - Format of the IPv6 location file. Required for the NetScaler to identify how to read the location file.
* `id` - The id of the locationfile6 datasource.
* `locationfile` - Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.
* `src` - URL (protocol, host, path, and file name) from where the location file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
