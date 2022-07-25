---
subcategory: "DNS"
---

# Resource: dnsnameserver

The dnsnameserver resource is used to create    DNS nameserver.


## Example usage

```hcl
resource "citrixadc_dnsnameserver" "dnsnameserver" {
	ip = "192.0.2.0"
    local = true
    state = "DISABLED"
    type = "UDP"
    dnsprofilename = "tf_profile1"
}
```


## Argument Reference

* `dnsvservername` - (Required either) Name of a DNS virtual server. Overrides any IP address-based name servers configured on the Citrix ADC.
* `ip` - (Required either) IP address of an external name server or, if the Local parameter is set, IP address of a local DNS server (LDNS).
* `dnsprofilename` - (Optional) Name of the DNS profile to be associated with the name server
* `local` - (Optional) Mark the IP address as one that belongs to a local recursive DNS server on the Citrix ADC. The appliance recursively resolves queries received on an IP address that is marked as being local. For recursive resolution to work, the global DNS parameter, Recursion, must also be set.   If no name server is marked as being local, the appliance functions as a stub resolver and load balances the name servers.
* `state` - (Optional) Administrative state of the name server.
* `type` - (Optional) Protocol used by the name server. UDP_TCP is not valid if the name server is a DNS virtual server configured on the appliance.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsnameserver. It has the same value as the `ip` or `dnsvserver` attribute.


## Import

A dnsnameserver can be imported using its ip or dnsvservername, e.g.

```shell
terraform import citrixadc_dnsnameserver.dnsnameserver 192.0.2.0
`or`
terraform import citrixadc_dnsnameserver.dnsnameserver dnsvservername1
```
