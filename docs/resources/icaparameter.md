---
subcategory: "ICA"
---

# Resource: icaparameter

The icaparameter resource is used to update icaparameter.


## Example usage

```hcl
resource "citrixadc_icaparameter" "tf_icaparameter" {
  edtpmtuddf           = "ENABLED"
  edtpmtuddftimeout    = 200
  l7latencyfrequency   = 30
  enablesronhafailover = "YES"
}
```


## Argument Reference

* `enablesronhafailover` - (Optional) Enable/Disable Session Reliability on HA failover. The default value is No. Possible values: [ YES, NO ]
* `edtpmtuddf` -  (Optional) Enable/Disable DF enforcement for EDT PMTUD Control Blocks.Default value: DISABLED Possible values = [ENABLED, DISABLED]
* `edtpmtuddftimeout` (Optional) DF enforcement timeout for EDTPMTUDDF. Default value: 100 Minimum value = 10 Maximum value = 65535
* `hdxinsightnonnsap` - (Optional) Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes. Possible values: [ YES, NO ]
* `l7latencyfrequency` - (Optional) Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0. Minimum value =  0 Maximum value =  60
* `dfpersistence` - (Optional) Enable/Disable DF Persistence
* `edtlosstolerant` - (Optional) Enable/Disable EDT Loss Tolerant feature
* `edtpmtuddftimeout` - (Optional) DF enforcement timeout for EDTPMTUDDF
* `edtpmtudrediscovery` - (Optional) Enable/Disable EDT PMTUD Rediscovery


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of theicaparameter. It is a unique string prefixed with `tf-icaparameter-`.