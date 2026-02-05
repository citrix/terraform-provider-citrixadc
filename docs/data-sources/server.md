---
subcategory: "Basic"
---

# Data Source: citrixadc_server

The server data source allows you to retrieve information about a server configuration.

## Example Usage

```terraform
data "citrixadc_server" "tf_server" {
  name = "test_server"
}

output "ipaddress" {
  value = data.citrixadc_server.tf_server.ipaddress
}

output "state" {
  value = data.citrixadc_server.tf_server.state
}
```

## Argument Reference

* `name` - (Required) Name of the server.

## Attribute Reference

The following attributes are available:

* `id` - The id of the server. It is a system-generated identifier.
* `name` - Name for the server.
* `ipaddress` - IPv4 address of the server.
* `ipv6address` - IPv6 address of the server.
* `domain` - Domain name of the server.
* `state` - Initial state of the server. Possible values: `ENABLED`, `DISABLED`.
* `comment` - Any information about the server.
* `td` - Traffic Domain ID.
* `translationip` - IP address used to transform the server's IP address.
* `translationmask` - The netmask of the translation IP.
