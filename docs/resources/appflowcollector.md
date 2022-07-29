---
subcategory: "AppFlow"
---

# Resource: appflowcollector

The appflowcollector resource is used to create appflowcollector.


## Example usage

```hcl
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
  name      = "tf_collector"
  ipaddress = "192.168.2.2"
  transport = "logstream"
  port      =  80
}
```


## Argument Reference

* `name` - (Required) Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  Only four collectors can be configured.   The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow collector" or 'my appflow collector').
* `ipaddress` - (Optional) IPv4 address of the collector.
* `netprofile` - (Optional) Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.
* `port` - (Optional) Port on which the collector listens.
* `transport` - (Optional) Type of collector: either logstream or ipfix or rest.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowcollector. It has the same value as the `name` attribute.


## Import

A appflowcollector can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowcollector.tf_appflowcollector tf_collector
```
