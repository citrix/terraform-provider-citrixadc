---
subcategory: "User"
---

# Resource: userprotocol

The userprotocol resource is used to create userprotocol.


## Example usage

```hcl
resource "citrixadc_userprotocol" "tf_userprotocol" {
  name      = "my_userprotocol"
  transport = "TCP"
  extension = "my_extension"
  comment   = "my_comment"
}
```


## Argument Reference

* `name` - (Required) Unique name for the user protocol. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Minimum length =  1
* `transport` - (Required) Transport layer's protocol. Possible values: [ TCP, SSL ]
* `extension` - (Required) Name of the extension to add parsing and runtime handling of the protocol packets.
* `comment` - (Optional) Any comments associated with the protocol.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the userprotocol. It has the same value as the `name` attribute.


## Import

A userprotocol can be imported using its name, e.g.

```shell
terraform import citrixadc_userprotocol.tf_userprotocol my_userprotocol
```
