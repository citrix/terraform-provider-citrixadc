---
subcategory: "Basic"
---

# Resource: locationfile

The locationfile resource is used to create locationfile.


## Example usage

```hcl
resource "citrixadc_locationfile" "tf_locationfile" {
  locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"
  format       = "netscaler"
}
```


## Argument Reference

* `locationfile` - (Required) Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both Citrix ADCs.
* `format` - (Required) Format of the location file. Required for the Citrix ADC to identify how to read the location file.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationfile. It has the same value as the `locationfile` attribute.


## Import

A locationfile can be imported using its name, e.g.

```shell
terraform import citrixadc_locationfile.tf_locationfile /var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4
```
