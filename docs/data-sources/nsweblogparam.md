---
subcategory: "NS"
---

# citrixadc_nsweblogparam (Data Source)

Data source for querying Citrix ADC Web Logging parameters. This data source retrieves information about the web logging configuration on the ADC appliance.

## Example Usage

```hcl
data "citrixadc_nsweblogparam" "example" {
}

# Output web logging parameters
output "buffer_size" {
  value = data.citrixadc_nsweblogparam.example.buffersizemb
}

output "custom_request_headers" {
  value = data.citrixadc_nsweblogparam.example.customreqhdrs
}

output "custom_response_headers" {
  value = data.citrixadc_nsweblogparam.example.customrsphdrs
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the nsweblogparam datasource.
* `buffersizemb` - Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system.
* `customreqhdrs` - List of HTTP request header names whose values should be exported by the Web Logging feature.
* `customrsphdrs` - List of HTTP response header names whose values should be exported by the Web Logging feature.

## Notes

The nsweblogparam resource is a singleton resource on the Citrix ADC appliance that contains web logging parameters. These parameters control how HTTP headers are logged and the buffer size for log data.
