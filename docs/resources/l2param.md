---
subcategory: "Network"
---

# Resource: l2param

The l2param resource is used to update l2param.


## Example usage

```hcl
resource "citrixadc_l2param" "tf_l2param" {
  mbfpeermacupdate   = 20
  maxbridgecollision = 30
  bdggrpproxyarp     = "DISABLED"
}
```


## Argument Reference

* `mbfpeermacupdate` - (Optional) When mbf_instant_learning is enabled, learn any changes in peer's MAC after this time interval, which is in 10ms ticks.
* `maxbridgecollision` - (Optional) Maximum bridge collision for loop detection .
* `bdggrpproxyarp` - (Optional) Set/reset proxy ARP in bridge group deployment. Possible values: [ ENABLED, DISABLED ]
* `bdgsetting` - (Optional) Bridging settings for C2C behavior. If enabled, each PE will learn MAC entries independently. Otherwise, when L2 mode is ON, learned MAC entries on a PE will be broadcasted to all other PEs. Possible values: [ ENABLED, DISABLED ]
* `garponvridintf` - (Optional) Send GARP messagess on VRID-configured interfaces upon failover . Possible values: [ ENABLED, DISABLED ]
* `macmodefwdmypkt` - (Optional) Allows MAC mode vserver to pick and forward the packets even if it is destined to Citrix ADC owned VIP. Possible values: [ ENABLED, DISABLED ]
* `usemymac` - (Optional) Use Citrix ADC MAC for all outgoing packets. Possible values: [ ENABLED, DISABLED ]
* `proxyarp` - (Optional) Proxies the ARP as Citrix ADC MAC for FreeBSD. Possible values: [ ENABLED, DISABLED ]
* `garpreply` - (Optional) Set/reset REPLY form of GARP . Possible values: [ ENABLED, DISABLED ]
* `mbfinstlearning` - (Optional) Enable instant learning of MAC changes in MBF mode. Possible values: [ ENABLED, DISABLED ]
* `rstintfonhafo` - (Optional) Enable the reset interface upon HA failover. Possible values: [ ENABLED, DISABLED ]
* `skipproxyingbsdtraffic` - (Optional) Control source parameters (IP and Port) for FreeBSD initiated traffic. If Enabled, source parameters are retained. Else proxy the source parameters based on next hop. Possible values: [ ENABLED, DISABLED ]
* `returntoethernetsender` - (Optional) Return to ethernet sender. Possible values: [ ENABLED, DISABLED ]
* `stopmacmoveupdate` - (Optional) Stop Update of server mac change to NAT sessions. Possible values: [ ENABLED, DISABLED ]
* `bridgeagetimeout` - (Optional) Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value. Minimum value =  60 Maximum value =  300
* `usenetprofilebsdtraffic` - (Optional) Control source parameters (IP and Port) for FreeBSD initiated traffic. If enabled proxy the source parameters based on netprofile source ip. If netprofile does not have ip configured, then it will continue to use NSIP as earlier. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the l2param. It is a unique string prefixed with "tf-l2param-".

