---
subcategory: "System"
---

# Resource: systemglobal_auditnslogpolicy_binding

The systemglobal_auditnslogpolicy_binding resource is used to create systemglobal_auditnslogpolicy_binding.


## Example usage

```hcl
resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
  policyname = "tf_auditnslogpolicy"
  priority   = 50
}
```


## Argument Reference

* `policyname` - (Required) The name of the  command policy.
* `priority` - (Required) The priority of the command policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]
* `feature` - (Optional) The feature to be checked while applying this config. Possible values: [ WL, WebLogging, SP, SurgeProtection, LB, LoadBalancing, CS, ContentSwitching, CR, CacheRedirection, SC, SureConnect, CMP, CMPcntl, CompressionControl, PQ, PriorityQueuing, HDOSP, HttpDoSProtection, SSLVPN, AAA, GSLB, GlobalServerLoadBalancing, SSL, SSLOffload, SSLOffloading, CF, ContentFiltering, IC, IntegratedCaching, OSPF, OSPFRouting, RIP, RIPRouting, BGP, BGPRouting, REWRITE, IPv6PT, IPv6protocoltranslation, AppFw, ApplicationFirewall, RESPONDER, HTMLInjection, push, NSPush, NetScalerPush, AppFlow, CloudBridge, ISIS, ISISRouting, CH, CallHome, AppQoE, ContentAccelerator, SYSTEM, RISE, FEO, LSN, LargeScaleNAT, RDPProxy, Rep, Reputation, URLFiltering, VideoOptimization, ForwardProxy, SSLInterception, AdaptiveTCP, CQA, CI, ContentInspection, Bot, APIGateway ]
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_auditnslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding tf_auditnslogpolicy
```
