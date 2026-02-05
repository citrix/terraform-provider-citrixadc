---
subcategory: "NS"
---

# citrixadc_ntpparam (Data Source)

Data source for querying Citrix ADC NTP parameters. This data source retrieves information about the NTP (Network Time Protocol) configuration parameters on the ADC appliance.

## Example Usage

```hcl
data "citrixadc_ntpparam" "example" {
}

# Output NTP parameters
output "authentication" {
  value = data.citrixadc_ntpparam.example.authentication
}

output "trusted_keys" {
  value = data.citrixadc_ntpparam.example.trustedkey
}

output "autokey_log_sec" {
  value = data.citrixadc_ntpparam.example.autokeylogsec
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the ntpparam datasource.
* `authentication` - Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted. Options: YES or NO.
* `autokeylogsec` - Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys, in seconds, expressed as a power of 2.
* `revokelogsec` - Interval between re-randomizations of the autokey seeds to prevent brute-force attacks on the autokey algorithms.
* `trustedkey` - List of key identifiers that are trusted for server authentication with symmetric key cryptography in the keys file.

## Notes

The ntpparam resource is a singleton resource on the Citrix ADC appliance that contains NTP parameters. These parameters control how the ADC synchronizes its time with NTP servers and manages authentication.
