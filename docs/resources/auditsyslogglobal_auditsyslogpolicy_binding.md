---
subcategory: "Audit"
---

# Resource: auditsyslogglobal_auditsyslogpolicy_binding

The auditsyslogglobal_auditsyslogpolicy_binding resource is used to create auditsyslogglobal_auditsyslogpolicy_binding.


## Example usage

```hcl
resource "citrixadc_auditsyslogglobal_auditsyslogpolicy_binding" "tf_auditsyslogglobal_auditsyslogpolicy_binding" {
  policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
  priority   = 100
  globalbindtype = "SYSTEM_GLOBAL"
}

resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
    name = "tf_auditsyslogpolicy"
    rule = "true"
    action = citrixadc_auditsyslogaction.tf_syslogaction.name
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
    name = "tf_syslogaction"
    serverip = "10.78.60.33"
    serverport = 514
    loglevel = [
        "ERROR",
        "NOTICE",
    ]
}
```


## Argument Reference

* `policyname` - (Required) Name of the audit syslog policy.
* `priority` - (Required) Specifies the priority of the policy. Minimum value =  1 Maximum value =  2147483647
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]
* `feature` - (Optional) The feature to be checked while applying this config. Possible values: [ WL, WebLogging, SP, SurgeProtection, LB, LoadBalancing, CS, ContentSwitching, CR, CacheRedirection, SC, SureConnect, CMP, CMPcntl, CompressionControl, PQ, PriorityQueuing, HDOSP, HttpDoSProtection, SSLVPN, AAA, GSLB, GlobalServerLoadBalancing, SSL, SSLOffload, SSLOffloading, CF, ContentFiltering, IC, IntegratedCaching, OSPF, OSPFRouting, RIP, RIPRouting, BGP, BGPRouting, REWRITE, IPv6PT, IPv6protocoltranslation, AppFw, ApplicationFirewall, RESPONDER, HTMLInjection, push, NSPush, NetScalerPush, AppFlow, CloudBridge, ISIS, ISISRouting, CH, CallHome, AppQoE, ContentAccelerator, SYSTEM, RISE, FEO, LSN, LargeScaleNAT, RDPProxy, Rep, Reputation, URLFiltering, VideoOptimization, ForwardProxy, SSLInterception, AdaptiveTCP, CQA, CI, ContentInspection, Bot, APIGateway ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditsyslogglobal_auditsyslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A auditsyslogglobal_auditsyslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy tf_auditsyslogpolicy
```
