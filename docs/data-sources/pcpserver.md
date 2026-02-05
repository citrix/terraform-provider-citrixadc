---
subcategory: "PCP"
---

# Data Source: citrixadc_pcpserver

The pcpserver data source allows you to retrieve information about a PCP server configuration.

## Example usage

```terraform
data "citrixadc_pcpserver" "tf_pcpserver" {
  name = "my_pcpserver"
}

output "ipaddress" {
  value = data.citrixadc_pcpserver.tf_pcpserver.ipaddress
}

output "port" {
  value = data.citrixadc_pcpserver.tf_pcpserver.port
}
```

## Argument Reference

* `name` - (Required) Name for the PCP server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpServer" or my pcpServer).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the pcpserver. It has the same value as the `name` attribute.
* `ipaddress` - The IP address of the PCP server.
* `pcpprofile` - pcp profile name
* `port` - Port number for the PCP server.

## Import

A pcpserver can be imported using its name, e.g.

```shell
terraform import citrixadc_pcpserver.tf_pcpserver my_pcpserver
```
