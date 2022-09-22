---
subcategory: "Basic"
---

# Resource: locationfile6

The locationfile6 resource is used to create locationfile6.


## Example usage

```hcl
resource "citrixadc_locationfile6" "tf_locationfile6" {
  locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6"
  format       = "netscaler6"
}
```


## Argument Reference

* `locationfile` - (Required) Name of the IPv6 location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both Citrix ADCs.
* `format` - (Required) Format of the IPv6 location file. Required for the Citrix ADC to identify how to read the location file.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationfile6. It has the same value as the `locationfile` attribute.


## Import

A locationfile6 can be imported using its name, e.g.

```shell
terraform import citrixadc_locationfile6.tf_locationfile6 /var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6
```
