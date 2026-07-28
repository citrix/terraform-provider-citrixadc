---
subcategory: "NS"
---

# Data Source: nscentralmanagementserver

The nscentralmanagementserver data source allows you to retrieve information about the central management server (Citrix ADM) registration configured on the Citrix ADC, looking it up by its type.


## Example usage

```terraform
data "citrixadc_nscentralmanagementserver" "example" {
  type = "ONPREM"
}

output "centralmgmtserver_ipaddress" {
  value = data.citrixadc_nscentralmanagementserver.example.ipaddress
}
```


## Argument Reference

* `type` - (Required) Type of the central management server. Must be either `CLOUD` or `ONPREM` depending on whether the server is on the cloud or on premise. Possible values: [ CLOUD, ONPREM ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nscentralmanagementserver. It has the same value as the `type` attribute.
* `ipaddress` - IP address of the central management server.
* `servername` - Fully qualified domain name of the central management server, or the service-url used to locate the ADM service.
* `username` - Username for access to the central management server.
* `adcusername` - ADC username used to create the device profile on ADM.
* `deviceprofilename` - Device profile created on ADM that contains the user name and password of the instance(s).
* `activationcode` - Activation code used to register to the ADM service.
* `validatecert` - Validate the server certificate for secure SSL connections.
