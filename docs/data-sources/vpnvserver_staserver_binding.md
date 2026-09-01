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

* `id` - The id of the vpnvserver_staserver_binding. It is the concatenation of the `name` and `staserver` attributes separated by a comma.
* `staaddresstype` - Type of the STA server address(ipv4/v6).

### Read-only vpnvserver_staserver_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnvserver_staserver_binding` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `stastate` - State of the STA Server. If Authority ID is set then STA Server is UP else DOWN. Possible values: [ UP, DOWN ]
* `staauthid` - Authority ID of the STA Server. Authority ID is used to match incoming STA tickets in the SOCKS/CGP protocol with the right STA server.
* `acttype` - Action type of the binding.
