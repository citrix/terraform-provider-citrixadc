---
subcategory: "Cloud"
---

# Resource: cloudtunnelparameter

This resource is used to manage the global cloud tunnel parameters on the Citrix ADC.

~> **Prerequisite:** The cloud tunnel feature is license-/feature-gated; on unsupported platforms the read is tolerated but configured values may not be echoed back by the appliance.


## Example usage

```hcl
resource "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
  controllerfqdn                 = "controller.citrixcloud.example.com"
  fqdn                           = "tunnel.citrixcloud.example.com"
  resourcelocation               = "00000000-0000-0000-0000-000000000000"
  subnetresourcelocationmappings = "10.0.0.0/24:00000000-0000-0000-0000-000000000000"
}
```


## Argument Reference

* `controllerfqdn` - (Optional) FQDN of the cloud controller that the appliance connects to when establishing the tunnel.
* `fqdn` - (Optional) FQDN advertised for the cloud tunnel endpoint on the appliance.
* `resourcelocation` - (Optional) Identifier of the resource location associated with the cloud tunnel.
* `subnetresourcelocationmappings` - (Optional) Mapping of on-prem subnets to their corresponding resource locations.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudtunnelparameter. It is set to `cloudtunnelparameter-config`.


## Import

A cloudtunnelparameter can be imported using its id, e.g.

```shell
terraform import citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter cloudtunnelparameter-config
```
