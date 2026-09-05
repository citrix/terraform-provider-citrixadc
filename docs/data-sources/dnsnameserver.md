---
subcategory: "DNS"
---

# Data Source: dnsnameserver

The dnsnameserver data source allows you to retrieve information about a DNS name server configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsnameserver" "tf_dnsnameserver" {
  ip   = "192.0.2.1"
  type = "UDP"
}

output "dns_state" {
  value = data.citrixadc_dnsnameserver.tf_dnsnameserver.state
}

output "dns_local" {
  value = data.citrixadc_dnsnameserver.tf_dnsnameserver.local
}

output "dns_profile" {
  value = data.citrixadc_dnsnameserver.tf_dnsnameserver.dnsprofilename
}
```


## Argument Reference

* `ip` - (Optional) IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS). Either `dnsvservername` or `ip` must be specified.
* `dnsvservername` - (Optional) Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC. Either `dnsvservername` or `ip` must be specified.
* `type` - (Optional) Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsnameserver. It is a composite key of the `ip` (or `dnsvservername`) and `type` attributes.
* `dnsprofilename` - Name of the DNS profile to be associated with the name server.
* `local` - Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set. If no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.
* `state` - Administrative state of the name server.

### Read-only dnsnameserver metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_dnsnameserver` resource) and are `Computed`. Any attribute the appliance does not return is `null`.

* `servicename` - The name of the DNS vserver.
* `port` - Port of the service. Range 1 - 65535 (`*` in the CLI is represented as 65535 in the NITRO API).
* `nameserverstate` - State of the server. Possible values = `UP`, `DOWN`, `UNKNOWN`, `BUSY`, `OUT OF SERVICE`, `GOING OUT OF SERVICE`, `DOWN WHEN GOING OUT OF SERVICE`, `NS_EMPTY_STR`, `Unknown`, `DISABLED`.
* `clmonowner` - Tells the mon owner of the service.
* `clmonview` - Tells the view id by which state of the service is updated.
