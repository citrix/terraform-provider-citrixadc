---
subcategory: "Cloud"
---

# Resource: cloudtunnelparameter

Configures the global cloud-tunnel parameters on the Citrix ADC. These settings identify the controller and resource-location endpoints used when the appliance establishes a management/data tunnel to a cloud service, and map on-prem subnets to their corresponding resource locations. Use this resource to point the appliance at the correct cloud controller and to advertise how local subnets are distributed across resource locations.

This is a singleton resource: a single `cloudtunnelparameter` configuration always exists on the appliance. Applying this resource updates the existing global configuration rather than creating a new object, so there is no delete operation and no name key.

~> **Prerequisite:** This feature is license/feature-gated. On platforms or releases where the cloud-tunnel feature is not enabled, the NITRO GET operation returns `Feature not supported in this release` (or a similar platform message). The provider tolerates this gracefully — the read is treated as non-fatal so apply is not broken — but the settable values may not be echoed back by the appliance. In that case the provider preserves the values you configured in Terraform state.


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
