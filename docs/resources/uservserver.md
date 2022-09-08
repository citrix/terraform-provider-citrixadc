---
subcategory: "User"
---

# Resource: uservserver

The uservserver resource is used to create uservserver.


## Example usage

```hcl
resource "citrixadc_uservserver" "tf_uservserver" {
  name         = "my_user_vserver"
  userprotocol = "MQTT"
  ipaddress    = "10.222.74.180"
  port         = 3200
  defaultlb    = "mysv"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). . Minimum length =  1
* `userprotocol` - (Required) User protocol uesd by the service.
* `ipaddress` - (Required) IPv4 or IPv6 address to assign to the virtual server.
* `port` - (Required) Port number for the virtual server. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `defaultlb` - (Optional) Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.
* `Params` - (Optional) Any comments associated with the protocol.
* `comment` - (Optional) Any comments that you might want to associate with the virtual server.
* `state` - (Optional) Initial state of the user vserver. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the uservserver. It has the same value as the `name` attribute.


## Import

A uservserver can be imported using its name, e.g.

```shell
terraform import citrixadc_uservserver.tf_uservserver my_user_vserver
```
