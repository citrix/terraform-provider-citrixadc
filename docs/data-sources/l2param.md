---
subcategory: "Network"
---

# Data Source: citrixadc_l2param

This data source is used to retrieve information about the Layer 2 parameters configuration.

## Example Usage

```hcl
data "citrixadc_l2param" "example" {
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the l2param resource.
* `bdggrpproxyarp` - Set/reset proxy ARP in bridge group deployment. Possible values: `ENABLED`, `DISABLED`. Default: `ENABLED`.
* `bdgsetting` - Bridging settings for C2C behavior. If enabled, each PE will learn MAC entries independently. Otherwise, when L2 mode is ON, learned MAC entries on a PE will be broadcasted to all other PEs. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `bridgeagetimeout` - Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value. Default: `300`.
* `garponvridintf` - Send GARP messages on VRID-configured interfaces upon failover. Possible values: `ENABLED`, `DISABLED`. Default: `ENABLED`.
* `garpreply` - Set/reset REPLY form of GARP. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `macmodefwdmypkt` - Allows MAC mode vserver to pick and forward the packets even if it is destined to Citrix ADC owned VIP. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `maxbridgecollision` - Maximum bridge collision for loop detection. Default: `20`.
* `mbfinstlearning` - Enable instant learning of MAC changes in MBF mode. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `mbfpeermacupdate` - When mbf_instant_learning is enabled, learn any changes in peer's MAC after this time interval, which is in 10ms ticks. Default: `10`.
* `proxyarp` - Proxies the ARP as Citrix ADC MAC for FreeBSD. Possible values: `ENABLED`, `DISABLED`. Default: `ENABLED`.
* `returntoethernetsender` - Return to ethernet sender. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `rstintfonhafo` - Enable the reset interface upon HA failover. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `skipproxyingbsdtraffic` - Control source parameters (IP and Port) for FreeBSD initiated traffic. If Enabled, source parameters are retained. Else proxy the source parameters based on next hop. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `stopmacmoveupdate` - Stop Update of server mac change to NAT sessions. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `usemymac` - Use Citrix ADC MAC for all outgoing packets. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `usenetprofilebsdtraffic` - Control source parameters (IP and Port) for FreeBSD initiated traffic. If enabled proxy the source parameters based on netprofile source ip. If netprofile does not have ip configured, then it will continue to use NSIP as earlier. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
