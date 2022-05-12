---
subcategory: "Ns"
---

# Resource: nslimitidentifier

The nslimitidentifier resource is used to create limit Indetifier resource.


## Example usage

```hcl
resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
  limitidentifier  = "tf_nslimitidentifier"
  threshold        = 1
  timeslice        = 1000
  limittype        = "BURSTY"
  mode             = "REQUEST_RATE"
  maxbandwidth     = 0
  trapsintimeslice = 1
}
```


## Argument Reference

* `limitidentifier` - (Required) Name for a rate limit identifier. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Reserved words must not be used.
* `threshold` - (Optional) Maximum number of requests that are allowed in the given timeslice when requests (mode is set as REQUEST_RATE) are tracked per timeslice. When connections (mode is set as CONNECTION) are tracked, it is the total number of connections that would be let through. Minimum value =  1
* `timeslice` - (Optional) Time interval, in milliseconds, specified in multiples of 10, during which requests are tracked to check if they cross the threshold. This argument is needed only when the mode is set to REQUEST_RATE. Minimum value =  10
* `mode` - (Optional) Defines the type of traffic to be tracked. * REQUEST_RATE - Tracks requests/timeslice. * CONNECTION - Tracks active transactions. Examples 1. To permit 20 requests in 10 ms and 2 traps in 10 ms: add limitidentifier limit_req -mode request_rate -limitType smooth -timeslice 1000 -Threshold 2000 -trapsInTimeSlice 200 2. To permit 50 requests in 10 ms: set  limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5000 -limitType smooth 3. To permit 1 request in 40 ms: set limitidentifier limit_req -mode request_rate -timeslice 2000 -Threshold 50 -limitType smooth 4. To permit 1 request in 200 ms and 1 trap in 130 ms: set limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5 -limitType smooth -trapsInTimeSlice 8 5. To permit 5000 requests in 1000 ms and 200 traps in 1000 ms: set limitidentifier limit_req  -mode request_rate -timeslice 1000 -Threshold 5000 -limitType BURSTY. Possible values: [ CONNECTION, REQUEST_RATE, NONE ]
* `limittype` - (Optional) Smooth or bursty request type. * SMOOTH - When you want the permitted number of requests in a given interval of time to be spread evenly across the timeslice * BURSTY - When you want the permitted number of requests to exhaust the quota anytime within the timeslice. This argument is needed only when the mode is set to REQUEST_RATE. Possible values: [ BURSTY, SMOOTH ]
* `selectorname` - (Optional) Name of the rate limit selector. If this argument is NULL, rate limiting will be applied on all traffic received by the virtual server or the Citrix ADC (depending on whether the limit identifier is bound to a virtual server or globally) without any filtering. Minimum length =  1
* `maxbandwidth` - (Optional) Maximum bandwidth permitted, in kbps. Minimum value =  0 Maximum value =  4294967287
* `trapsintimeslice` - (Optional) Number of traps to be sent in the timeslice configured. A value of 0 indicates that traps are disabled. Minimum value =  0 Maximum value =  65535


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslimitidentifier. It has the same value as the `limitidentifier` attribute.


## Import

A nslimitidentifier can be imported using its limitidentifier, e.g.

```shell
terraform import citrixadc_nslimitidentifier.tf_nslimitidentifier tf_nslimitidentifier
```
