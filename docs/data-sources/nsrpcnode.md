---
subcategory: "NS"
---

# Data Source `nsrpcnode`

The nsrpcnode data source allows you to retrieve information about RPC nodes configured on the NetScaler appliance. This is typically used in cluster configurations.


## Example usage

```terraform
resource "citrixadc_nsrpcnode" "my_nsrpcnode" {
  ipaddress = "10.222.74.146"
  password  = "mypassword"
  secure    = "ON"
  srcip     = "10.222.74.146"
}

data "citrixadc_nsrpcnode" "my_nsrpcnode_data" {
  ipaddress = citrixadc_nsrpcnode.my_nsrpcnode.ipaddress
}

output "rpcnode_secure" {
  value = data.citrixadc_nsrpcnode.my_nsrpcnode_data.secure
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the node. This has to be in the same subnet as the NSIP address.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `password` - Password to be used in authentication with the peer system node.
* `secure` - State of the channel when talking to the node.
* `srcip` - Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.
* `validatecert` - Validate the server certificate for secure SSL connections.
* `id` - The id of the nsrpcnode. It has the same value as the `ipaddress` attribute.
