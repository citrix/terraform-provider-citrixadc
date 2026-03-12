---
subcategory: "Subscriber"
---

# Data Source: subscriberparam

The subscriberparam data source allows you to retrieve information about the subscriber parameters configuration.

## Example usage

```terraform
data "citrixadc_subscriberparam" "tf_subscriberparam" {
}

output "keytype" {
  value = data.citrixadc_subscriberparam.tf_subscriberparam.keytype
}

output "interfacetype" {
  value = data.citrixadc_subscriberparam.tf_subscriberparam.interfacetype
}

output "idlettl" {
  value = data.citrixadc_subscriberparam.tf_subscriberparam.idlettl
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the subscriberparam. It is a system-generated identifier.
* `idleaction` - Once idleTTL exprires on a subscriber session, Citrix ADC will take an idle action on that session. idleAction could be chosen from one of these ==> 1. ccrTerminate: (default) send CCR-T to inform PCRF about session termination and delete the session. 2. delete: Just delete the subscriber session without informing PCRF. 3. ccrUpdate: Do not delete the session and instead send a CCR-U to PCRF requesting for an updated session.
* `idlettl` - Idle Timeout, in seconds, after which Citrix ADC will take an idleAction on a subscriber session (refer to 'idleAction' arguement in 'set subscriber param' for more details on idleAction). Any data-plane or control plane activity updates the idleTimeout on subscriber session. idleAction could be to 'just delete the session' or 'delete and CCR-T' (if PCRF is configured) or 'do not delete but send a CCR-U'. Zero value disables the idle timeout.
* `interfacetype` - Subscriber Interface refers to Citrix ADC interaction with control plane protocols, RADIUS and GX. Types of subscriber interface: NONE, RadiusOnly, RadiusAndGx, GxOnly. NONE: Only static subscribers can be configured. RadiusOnly: GX interface is absent. Subscriber information is obtained through RADIUS Accounting messages. RadiusAndGx: Subscriber ID obtained through RADIUS Accounting is used to query PCRF. Subscriber information is obtained from both RADIUS and PCRF. GxOnly: RADIUS interface is absent. Subscriber information is queried using Subscriber IP or IP+VLAN.
* `ipv6prefixlookuplist` - The ipv6PrefixLookupList should consist of all the ipv6 prefix lengths assigned to the UE's.
* `keytype` - Type of subscriber key type IP or IPANDVLAN. IPANDVLAN option can be used only when the interfaceType is set to gxOnly. Changing the lookup method should result to the subscriber session database being flushed.
