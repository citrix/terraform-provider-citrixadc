---
subcategory: "NS"
---

# Data Source `nsratecontrol`

The nsratecontrol data source allows you to retrieve information about rate control configuration on the NetScaler appliance.


## Example usage

```terraform
resource "citrixadc_nsratecontrol" "my_nsratecontrol" {
  tcpthreshold    = 100
  udpthreshold    = 100
  icmpthreshold   = 100
  tcprstthreshold = 100
}

data "citrixadc_nsratecontrol" "my_nsratecontrol_data" {
  depends_on = [citrixadc_nsratecontrol.my_nsratecontrol]
}

output "tcp_threshold" {
  value = data.citrixadc_nsratecontrol.my_nsratecontrol_data.tcpthreshold
}

output "udp_threshold" {
  value = data.citrixadc_nsratecontrol.my_nsratecontrol_data.udpthreshold
}
```


## Argument Reference

This is a singleton resource, so no arguments are required.

## Attribute Reference

The following attributes are available:

* `icmpthreshold` - Number of ICMP packets permitted per 10 milliseconds.
* `tcprstthreshold` - The number of TCP RST packets permitted per 10 milliseconds. Zero means rate control is disabled and 0xffffffff means everything is rate controlled.
* `tcpthreshold` - Number of SYNs permitted per 10 milliseconds.
* `udpthreshold` - Number of UDP packets permitted per 10 milliseconds.
* `id` - The id of the nsratecontrol resource (singleton).
