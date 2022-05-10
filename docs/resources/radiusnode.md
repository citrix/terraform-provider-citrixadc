---
subcategory: "Basic"
---

# Resource: radiusnode

The radiusnode resource is used for Configuration of RADIUS Node resource


## Example usage

```hcl
resource "citrixadc_radiusnode" "tf_radiusnode" {
  nodeprefix = "10.10.10.10/32"
  radkey     = "secret"
}
```


## Argument Reference

* `nodeprefix` - (Required) IP address/IP prefix of radius node in CIDR format
* `radkey` - (Optional) The key shared between the RADIUS server and clients.       Required for Citrix ADC to communicate with the RADIUS nodes.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the radiusnode. It has the same value as the `nodeprefix` attribute.


## Import

A radiusnode can be imported using its nodeprefix, e.g.

```shell
terraform import citrixadc_radiusnode.tf_radiusnode 10.10.10.10/32
```
