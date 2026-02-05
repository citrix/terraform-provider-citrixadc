---
subcategory: "NS"
---

# Data Source: citrixadc_nsdiameter

The `citrixadc_nsdiameter` data source is used to retrieve information about the Diameter configuration on a Citrix ADC appliance. Diameter is a protocol used for Authentication, Authorization, and Accounting (AAA) in modern networks.

## Example usage

```hcl
# Retrieve the diameter configuration for the default node
data "citrixadc_nsdiameter" "default" {
  ownernode = -1
}

# Use the retrieved data in outputs
output "diameter_identity" {
  value = data.citrixadc_nsdiameter.default.identity
}

output "diameter_realm" {
  value = data.citrixadc_nsdiameter.default.realm
}

```

## Argument Reference

The following arguments are required:

* `ownernode` - (Required) ID of the cluster node for which the diameter id is set. Use `-1` for standalone or default node configuration. This can only be configured through CLIP (Command Line Interface Protocol).

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the nsdiameter configuration. It has the same value as the `ownernode` attribute.
* `identity` - DiameterIdentity to be used by the Citrix ADC. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity. This identity will be used as Origin-Host AVP as defined in RFC3588.
  
  Example:
  ```
  set ns diameter -identity netscaler.com
  ```
  Now whenever Citrix ADC needs to use identity in diameter messages, it will use 'netscaler.com' as Origin-Host AVP.

* `realm` - Diameter Realm to be used by the Citrix ADC. This realm will be used as Origin-Realm AVP as defined in RFC3588.
  
  Example:
  ```
  set ns diameter -realm com
  ```
  Now whenever Citrix ADC system needs to use realm in diameter messages, it will use 'com' as Origin-Realm AVP.

* `serverclosepropagation` - When a server connection goes down, whether to close the corresponding client connection if there were requests pending on the server. Possible values: `YES`, `NO`.

## Notes

* The nsdiameter configuration is typically a singleton resource, meaning there is only one diameter configuration per node.
* In standalone deployments, use `ownernode = -1` to reference the default node.
* In cluster deployments, use the appropriate cluster node ID.
