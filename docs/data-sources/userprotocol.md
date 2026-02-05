---
subcategory: "User"
---

# Data Source: citrixadc_userprotocol

The citrixadc_userprotocol data source is used to retrieve information about an existing userprotocol resource.

## Example Usage

```terraform
data "citrixadc_userprotocol" "example" {
  name = "my_userprotocol"
}

output "userprotocol_id" {
  value = data.citrixadc_userprotocol.example.id
}

output "userprotocol_transport" {
  value = data.citrixadc_userprotocol.example.transport
}

output "userprotocol_extension" {
  value = data.citrixadc_userprotocol.example.extension
}
```

## Argument Reference

* `name` - (Required) Unique name for the user protocol. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the userprotocol resource.
* `comment` - Any comments associated with the protocol.
* `extension` - Name of the extension to add parsing and runtime handling of the protocol packets.
* `transport` - Transport layer's protocol.
