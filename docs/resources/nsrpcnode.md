---
subcategory: "NS"
---

# Resource: nsrpcnode

The nsrpcnode resource is used to manage rpc nodes.


## Example usage

```hcl
resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.78.60.201"
    password = "verysecret"
    secure = "ON"
    srcip = "10.78.60.201"
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the node. This has to be in the same subnet as the NSIP address.
* `password` - (Optional) Password to be used in authentication with the peer system node.
* `srcip` - (Optional) Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.
* `secure` - (Optional) State of the channel when talking to the node. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsrpcnode. It has the same value as the `ipaddress` attribute.


## Import

A nsrpcnode can be imported using its ipaddress, e.g.

```shell
terraform import citrixadc_nsrpcnode.tf_nsrpcnode 10.78.60.201
```
