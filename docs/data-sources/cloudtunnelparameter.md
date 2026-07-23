---
subcategory: "Cloud"
---

# Data Source: cloudtunnelparameter

The cloudtunnelparameter data source allows you to retrieve information about the global cloud tunnel parameters.

~> **Note:** This feature is license/feature-gated; on platforms where it is not enabled the read is tolerated and the returned attribute values may be empty.


## Example usage

```hcl
data "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
}

output "cloudtunnel_controllerfqdn" {
  value = data.citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter.controllerfqdn
}

output "cloudtunnel_resourcelocation" {
  value = data.citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter.resourcelocation
}
```


## Argument Reference

This data source takes no arguments; it always reads the singleton `cloudtunnelparameter` configuration.


## Attribute Reference

The following attributes are available:

* `id` - The id of the cloudtunnelparameter. It is set to `cloudtunnelparameter-config`.
* `controllerfqdn` - FQDN of the cloud controller that the appliance connects to when establishing the tunnel.
* `fqdn` - FQDN advertised for the cloud tunnel endpoint on the appliance.
* `resourcelocation` - Identifier of the resource location associated with the cloud tunnel.
* `subnetresourcelocationmappings` - Mapping of on-prem subnets to their corresponding resource locations.
