---
subcategory: "Basic"
---

# Data Source `locationfile`

The locationfile data source allows you to retrieve information about the location file configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_locationfile" "tf_locationfile" {
}

output "locationfile" {
  value = data.citrixadc_locationfile.tf_locationfile.locationfile
}

output "format" {
  value = data.citrixadc_locationfile.tf_locationfile.format
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `format` - Format of the location file. Required for the NetScaler to identify how to read the location file.
* `locationfile` - Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.
* `src` - URL (protocol, host, path, and file name) from where the location file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.

## Attribute Reference

* `id` - The id of the locationfile. It is a system-generated identifier.
