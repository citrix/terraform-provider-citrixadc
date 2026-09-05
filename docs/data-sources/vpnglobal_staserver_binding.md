---
subcategory: "VPN"
---

# Data Source: vpnglobal_staserver_binding

The vpnglobal_staserver_binding data source allows you to retrieve information about a vpnglobal staserver binding configuration.


## Example usage

```terraform
data "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
  staserver = "http://www.example.com/"
}

output "staaddresstype" {
  value = data.citrixadc_vpnglobal_staserver_binding.tf_bind.staaddresstype
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_staserver_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `staserver` - (Required) Configured Secure Ticketing Authority (STA) server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_staserver_binding. It has the same value as the `staserver` attribute.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `staaddresstype` - Type of the STA server address(ipv4/v6).

### Read-only vpnglobal_staserver_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnglobal_staserver_binding` resource). They are Computed and are `null` when the appliance does not return them.

* `staauthid` - Authority ID of the STA Server. Authority ID is used to match incoming STA Tickets in the SOCKS/CGP protocol with the right STA Server.
* `stastate` - State of the STA Server. If Authority ID is set then STA Server is UP else DOWN. Possible values: `UP`, `DOWN`.
