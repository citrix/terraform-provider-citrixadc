---
subcategory: "Basic"
---

# Resource: server

The server resource is used to create servers.


## Example usage

```hcl
resource "citrixadc_server" "tf_server" {
	name = "tf_server"
	ipaddress = "192.168.11.13"
}
```


## Argument Reference

* `name` - (Optional) Name for the server.
* `ipaddress` - (Optional) IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.
* `domain` - (Optional) Domain name of the server. For a domain based configuration, you must create the server first.
* `translationip` - (Optional) IP address used to transform the server's DNS-resolved IP address.
* `translationmask` - (Optional) The netmask of the translation ip.
* `domainresolveretry` - (Optional) Time, in seconds, for which the Citrix ADC must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.
* `state` - (Optional) Initial state of the server. Possible values: [ ENABLED, DISABLED ]
* `ipv6address` - (Optional) Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode. Possible values: [ YES, NO ]
* `comment` - (Optional) Any information about the server.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `querytype` - (Optional) Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument. Possible values: [ A, AAAA, SRV ]
* `domainresolvenow` - (Optional) Immediately send a DNS query to resolve the server's domain name.
* `delay` - (Optional) Time, in seconds, after which all the services configured on the server are disabled.
* `graceful` - (Optional) Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed. Possible values: [ YES, NO ]
* `internal` - (Optional) Display names of the servers that have been created for internal use.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the server. It has the same value as the `name` attribute.


## Import

A server can be imported using its name, e.g.

```shell
terraform import citrixadc_server.tf_server tf_server
```
