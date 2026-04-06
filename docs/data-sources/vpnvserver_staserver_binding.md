---
subcategory: "VPN"
---

# Data Source: vpnvserver_staserver_binding

The vpnvserver_staserver_binding data source allows you to retrieve information about a specific vpnvserver to staserver binding.


## Example Usage

```terraform
data "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
  name      = "tf_vserver"
  staserver = "http://www.example.com/"
}

output "staaddresstype" {
  value = data.citrixadc_vpnvserver_staserver_binding.tf_binding.staaddresstype
}

output "binding_id" {
  value = data.citrixadc_vpnvserver_staserver_binding.tf_binding.id
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `staserver` - (Required) Configured Secure Ticketing Authority (STA) server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_staserver_binding. It is the concatenation of `name` and `staserver` attributes separated by comma.
* `staaddresstype` - Type of the STA server address(ipv4/v6).
