---
subcategory: "ICA"
---

# Data Source `icaparameter`

The icaparameter data source allows you to retrieve information about ICA (Independent Computing Architecture) parameters configuration.


## Example usage

```terraform
data "citrixadc_icaparameter" "tf_icaparameter" {
}

output "edtpmtuddf" {
  value = data.citrixadc_icaparameter.tf_icaparameter.edtpmtuddf
}

output "l7latencyfrequency" {
  value = data.citrixadc_icaparameter.tf_icaparameter.l7latencyfrequency
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `dfpersistence` - Enable/Disable DF Persistence.
* `edtlosstolerant` - Enable/Disable EDT Loss Tolerant feature.
* `edtpmtuddf` - Enable/Disable DF enforcement for EDT PMTUD Control Blocks.
* `edtpmtuddftimeout` - DF enforcement timeout for EDTPMTUDDF.
* `edtpmtudrediscovery` - Enable/Disable EDT PMTUD Rediscovery.
* `enablesronhafailover` - Enable/Disable Session Reliability on HA failover. The default value is No.
* `hdxinsightnonnsap` - Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes.
* `l7latencyfrequency` - Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0.
* `id` - The id of the icaparameter. It is a system-generated identifier.
