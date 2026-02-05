---
subcategory: "AAA"
---

# Data Source: citrixadc_radiusnode

The `citrixadc_radiusnode` data source is used to retrieve information about a specific RADIUS node configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_radiusnode" "example" {
  nodeprefix = "192.168.1.0/24"
}
```

## Argument Reference

* `nodeprefix` - (Required) IP address/IP prefix of RADIUS node in CIDR format.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the RADIUS node.
* `radkey` - The key shared between the RADIUS server and clients. Required for NetScaler to communicate with the RADIUS nodes.
