---
subcategory: "NS"
---

# citrixadc_nslimitidentifier (Data Source)

Data source for querying Citrix ADC limit identifier configuration. This data source retrieves information about a configured limit identifier, which is used for rate limiting traffic based on various criteria.

## Example Usage

```hcl
data "citrixadc_nslimitidentifier" "example" {
  limitidentifier = "my_limit_identifier"
}

# Output limit identifier information
output "threshold" {
  value = data.citrixadc_nslimitidentifier.example.threshold
}

output "timeslice" {
  value = data.citrixadc_nslimitidentifier.example.timeslice
}

output "mode" {
  value = data.citrixadc_nslimitidentifier.example.mode
}
```

## Argument Reference

The following arguments are required:

* `limitidentifier` - (Required) Name of the rate limit identifier to retrieve. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nslimitidentifier datasource.
* `limittype` - Smooth or bursty request type. Possible values: `SMOOTH`, `BURSTY`. When `SMOOTH`, the permitted number of requests in a given interval of time are spread evenly across the timeslice. When `BURSTY`, the permitted number of requests can exhaust the quota anytime within the timeslice. This argument is needed only when the mode is set to `REQUEST_RATE`.
* `maxbandwidth` - Maximum bandwidth permitted, in kbps.
* `mode` - Defines the type of traffic to be tracked. Possible values: `REQUEST_RATE`, `CONNECTION`. When `REQUEST_RATE`, tracks requests per timeslice. When `CONNECTION`, tracks active transactions.
* `selectorname` - Name of the rate limit selector. If this argument is NULL, rate limiting will be applied on all traffic received by the virtual server or the Citrix ADC (depending on whether the limit identifier is bound to a virtual server or globally) without any filtering.
* `threshold` - Maximum number of requests that are allowed in the given timeslice when requests (mode is set as `REQUEST_RATE`) are tracked per timeslice. When connections (mode is set as `CONNECTION`) are tracked, it is the total number of connections that would be let through.
* `timeslice` - Time interval, in milliseconds, specified in multiples of 10, during which requests are tracked to check if they cross the threshold. This argument is needed only when the mode is set to `REQUEST_RATE`.
* `trapsintimeslice` - Number of traps to be sent in the timeslice configured. A value of 0 indicates that traps are disabled.

## Example Usage with Resource

```hcl
resource "citrixadc_nslimitidentifier" "rate_limiter" {
  limitidentifier  = "api_rate_limiter"
  threshold        = 100
  timeslice        = 1000
  limittype        = "BURSTY"
  mode             = "REQUEST_RATE"
  maxbandwidth     = 500
  trapsintimeslice = 5
}

data "citrixadc_nslimitidentifier" "rate_limiter_info" {
  limitidentifier = citrixadc_nslimitidentifier.rate_limiter.limitidentifier
}
```
