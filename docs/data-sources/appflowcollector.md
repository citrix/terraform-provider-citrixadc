---
subcategory: "AppFlow"
---

# Data Source `appflowcollector`

The appflowcollector data source allows you to retrieve information about an existing appflowcollector.


## Example usage

```terraform
data "citrixadc_appflowcollector" "tf_appflowcollector" {
  name = "tf_collector"
}

output "name" {
  value = data.citrixadc_appflowcollector.tf_appflowcollector.name
}

output "ipaddress" {
  value = data.citrixadc_appflowcollector.tf_appflowcollector.ipaddress
}

output "transport" {
  value = data.citrixadc_appflowcollector.tf_appflowcollector.transport
}

output "port" {
  value = data.citrixadc_appflowcollector.tf_appflowcollector.port
}
```


## Argument Reference

* `name` - (Required) Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  Only four collectors can be configured.   The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow collector" or 'my appflow collector').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowcollector. It has the same value as the `name` attribute.
* `ipaddress` - IPv4 address of the collector.
* `netprofile` - Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.
* `newname` - New name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at(@), equals (=), and hyphen (-) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow coll" or 'my appflow coll').
* `port` - Port on which the collector listens.
* `transport` - Type of collector: either logstream or ipfix or rest.
