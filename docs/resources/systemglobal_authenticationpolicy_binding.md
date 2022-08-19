---
subcategory: "System"
---

# Resource: systemglobal_authenticationpolicy_binding

The systemglobal_authenticationpolicy_binding resource is used to create systemglobal_authenticationpolicy_binding.


## Example usage

```hcl
resource "citrixadc_systemglobal_authenticationpolicy_binding" "tf_systemglobal_authenticationpolicy_binding" {
  policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
  priority   = 50
}

resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
  name   = "tf_authenticationpolicy"
  rule   = "true"
  action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}>
```


## Argument Reference

* `policyname` - (Required) The name of the  command policy.
* `priority` - (Required) The priority of the command policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]
* `feature` - (Optional) The feature to be checked while applying this config. Possible values: [ WL, WebLogging, SP, SurgeProtection, LB, LoadBalancing, CS, ContentSwitching, CR, CacheRedirection, SC, SureConnect, CMP, CMPcntl, CompressionControl, PQ, PriorityQueuing, HDOSP, HttpDoSProtection, SSLVPN, AAA, GSLB, GlobalServerLoadBalancing, SSL, SSLOffload, SSLOffloading, CF, ContentFiltering, IC, IntegratedCaching, OSPF, OSPFRouting, RIP, RIPRouting, BGP, BGPRouting, REWRITE, IPv6PT, IPv6protocoltranslation, AppFw, ApplicationFirewall, RESPONDER, HTMLInjection, push, NSPush, NetScalerPush, AppFlow, CloudBridge, ISIS, ISISRouting, CH, CallHome, AppQoE, ContentAccelerator, SYSTEM, RISE, FEO, LSN, LargeScaleNAT, RDPProxy, Rep, Reputation, URLFiltering, VideoOptimization, ForwardProxy, SSLInterception, AdaptiveTCP, CQA, CI, ContentInspection, Bot, APIGateway ]
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Applicable only for advanced authentication policies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_authenticationpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_authenticationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_systemglobal_authenticationpolicy_binding.tf_systemglobal_authenticationpolicy_binding tf_authenticationpolicy
```
