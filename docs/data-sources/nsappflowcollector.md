---
subcategory: "NS"
---

# Data Source `nsappflowcollector`

The nsappflowcollector data source allows you to retrieve information about AppFlow collectors configured on the NetScaler ADC.


## Example usage

```terraform
data "citrixadc_nsappflowcollector" "tf_appflowcollector" {
  name = "my_appflowcollector"
}

output "ipaddress" {
  value = data.citrixadc_nsappflowcollector.tf_appflowcollector.ipaddress
}

output "port" {
  value = data.citrixadc_nsappflowcollector.tf_appflowcollector.port
}
```


## Argument Reference

* `name` - (Required) Name of the AppFlow collector.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ipaddress` - The IPv4 address of the AppFlow collector.
* `port` - The UDP port on which the AppFlow collector is listening.
* `id` - The id of the nsappflowcollector. It has the same value as the `name` attribute.


## Import

An nsappflowcollector can be imported using its name, e.g.

```shell
terraform import citrixadc_nsappflowcollector.tf_appflowcollector my_appflowcollector
```
