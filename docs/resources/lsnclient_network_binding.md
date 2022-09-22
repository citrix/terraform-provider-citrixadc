---
subcategory: "LSN"
---

# Resource: lsnclient_network_binding

The lsnclient_network_binding resource is used to create lsnclient_network_binding.


## Example usage

```hcl
resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
  clientname = "my_lsnclient"
  network    = "10.222.74.160"
}
```


## Argument Reference

* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').
* `network` - (Required) IPv4 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.
* `netmask` - (Optional) Subnet mask for the IPv4 address specified in the Network parameter.
* `td` - (Optional) ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs.  If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient_network_binding. It is the concatenation of `clientname` and `network` attributes separated by a comma.


## Import

A lsnclient_network_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding my_lsnclient,10.222.74.160
```
