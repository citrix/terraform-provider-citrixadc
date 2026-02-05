---
subcategory: "NS"
---

# Data Source `nsdhcpparams`

The nsdhcpparams data source allows you to retrieve information about the DHCP parameters configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
}

output "dhcpclient" {
  value = data.citrixadc_nsdhcpparams.tf_nsdhcpparams.dhcpclient
}

output "saveroute" {
  value = data.citrixadc_nsdhcpparams.tf_nsdhcpparams.saveroute
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `dhcpclient` - Enables DHCP client to acquire IP address from the DHCP server in the next boot. When set to OFF, disables the DHCP client in the next boot. Possible values: [ ON, OFF ]
* `saveroute` - DHCP acquired routes are saved on the Citrix ADC. Possible values: [ ON, OFF ]

## Attribute Reference

* `id` - The id of the nsdhcpparams. It is a system-generated identifier.
