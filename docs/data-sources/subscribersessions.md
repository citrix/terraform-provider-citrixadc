---
subcategory: "Subscriber"
---

# Data Source: subscribersessions

The subscribersessions data source retrieves subscriber session telemetry from the Citrix ADC. It is backed by the ADC "show subscriber sessions" (get-all) call and lets you inspect the Subscriber/Gx (PCRF) session state, optionally filtered by subscriber IP, VLAN, or cluster node.

-> **Note:** The Subscriber/Gx/PCRF (Telco) feature must be licensed and enabled on the Citrix ADC for subscriber sessions to exist. When no sessions are present (or none match the filters), the query returns successfully with an empty result; an empty list is a valid outcome and not an error.


## Example usage

```hcl
data "citrixadc_subscribersessions" "example" {
  ip = "198.51.100.25"
}

output "subscriber_ttl" {
  value = data.citrixadc_subscribersessions.example.ttl
}

output "subscriber_rules" {
  value = data.citrixadc_subscribersessions.example.subscriberrules
}
```


## Argument Reference

All filters are optional. Omit them to retrieve all subscriber sessions.

* `ip` - (Optional) Subscriber IP address to filter on.
* `vlan` - (Optional) The VLAN number on which the subscriber is located.
* `nodeid` - (Optional) Unique number that identifies the cluster node.


## Attribute Reference

In addition to the arguments, the following read-only attributes are available:

* `id` - A synthetic identifier for the data source instance.
* `subscriptionidtype` - Subscription-Id type. Possible values: E164, IMSI, SIP_URI, NAI, PRIVATE.
* `subscriptionidvalue` - Subscription-Id value.
* `subscriberrules` - Rules stored in this session for this subscriber.
* `flags` - Subscriber session flags.
* `ttl` - Subscriber session revalidation timeout remaining.
* `idlettl` - Subscriber session activity timeout remaining.
* `avpdisplaybuffer` - Subscriber attributes display.
* `servicepath` - Name of the service path to be taken for this subscriber.
