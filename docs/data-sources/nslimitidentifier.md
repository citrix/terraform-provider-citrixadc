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

### Read-only nslimitidentifier metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nslimitidentifier` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `ngname` - Nodegroup name to which this identifier belongs.
* `hits` - The number of times this identifier was evaluated.
* `drop` - The number of times action was taken.
* `rule` - Rule. A list of strings.
* `time` - Time interval considered for rate limiting.
* `total` - Maximum number of requests permitted in the computed timeslice.
* `trapscomputedintimeslice` - The number of traps that would be sent in the timeslice configured.
* `computedtraptimeslice` - The time interval computed for sending traps.
* `alertscomputedintimeslice` - The number of appflow alerts that would be sent in the timeslice configured.
* `computedalerttimeslice` - The time interval computed for sending appflow alerts.
* `referencecount` - Total number of transactions pointing to this entry.
