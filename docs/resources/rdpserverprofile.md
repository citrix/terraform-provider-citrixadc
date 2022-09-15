---
subcategory: "Rdp"
---

# Resource: rdpserverprofile

The rdpserverprofile resource is used to create rdpserverprofile.


## Example usage

```hcl
resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
  name           = "my_rdpserverprofile"
  psk            = "key"
  rdpredirection = "ENABLE"
  rdpport        = 4000
}
```


## Argument Reference

* `name` - (Required) The name of the rdp server profile
* `psk` - (Required) Pre shared key value
* `rdpip` - (Optional) IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.
* `rdpport` - (Optional) TCP port on which the RDP connection is established.
* `rdpredirection` - (Optional) Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rdpserverprofile. It has the same value as the `name` attribute.


## Import

A rdpserverprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_rdpserverprofile.tf_rdpserverprofile my_rdpserverprofile
```
