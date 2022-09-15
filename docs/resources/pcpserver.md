---
subcategory: "Pcp"
---

# Resource: pcpserver

The pcpserverresource is used to create pcpserver.


## Example usage

```hcl
resource "citrixadc_pcpserver" "tf_pcpserver" {
  name      = "my_pcpserver"
  ipaddress = "10.222.74.185"
}
```


## Argument Reference

* `ipaddress` - (Required) The IP address of the PCP server.
* `name` - (Required) Name for the PCP server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpServer" or my pcpServer).
* `pcpprofile` - (Optional) pcp profile name
* `port` - (Optional) Port number for the PCP server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the pcpserver. It has the same value as the `name` attribute.


## Import

A pcpserver can be imported using its name, e.g.

```shell
terraform import citrixadc_pcpserver.tf_pcpserver my_pcpserver
```
