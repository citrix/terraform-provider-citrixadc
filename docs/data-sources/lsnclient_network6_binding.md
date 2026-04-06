---
subcategory: "LSN"
---

# Data Source: lsnclient_network6_binding

The lsnclient_network6_binding data source allows you to retrieve information about LSN client IPv6 network bindings.

## Example Usage

```terraform
data "citrixadc_lsnclient_network6_binding" "tf_lsnclient_network6_binding" {
  clientname = "my_lsn_client"
  network6   = "2001:db8:5001::/96"
}

output "clientname" {
  value = data.citrixadc_lsnclient_network6_binding.tf_lsnclient_network6_binding.clientname
}

output "network6" {
  value = data.citrixadc_lsnclient_network6_binding.tf_lsnclient_network6_binding.network6
}
```

## Argument Reference

* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').
* `network6` - (Required) IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `td` - ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.
* `id` - The id of the lsnclient_network6_binding. It is a system-generated identifier.
