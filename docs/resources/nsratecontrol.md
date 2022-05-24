---
subcategory: "NS"
---

# Resource: nsratecontrol

The nsratecontrol resource is used to create rate control resource.


## Example usage

```hcl
resource "citrixadc_nsratecontrol" "tf_nsratecontrol" {
  tcpthreshold    = 10
  udpthreshold    = 10
  icmpthreshold   = 100
  tcprstthreshold = 100
}
```


## Argument Reference

* `tcpthreshold` - (Optional) Number of SYNs permitted per 10 milliseconds.
* `udpthreshold` - (Optional) Number of UDP packets permitted per 10 milliseconds.
* `icmpthreshold` - (Optional) Number of ICMP packets permitted per 10 milliseconds.
* `tcprstthreshold` - (Optional) The number of TCP RST packets permitted per 10 milli second. zero means rate control is disabled and 0xffffffff means every thing is rate controlled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsratecontrol. It is a unique string prefixed with "tf-nsratecontrol-"

