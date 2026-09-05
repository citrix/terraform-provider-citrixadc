---
subcategory: "NS"
---

# Data Source: nscapacity

The nscapacity data source allows you to retrieve information about NetScaler capacity license configuration.


## Example usage

```terraform
data "citrixadc_nscapacity" "tf_nscapacity" {
}

output "bandwidth" {
  value = data.citrixadc_nscapacity.tf_nscapacity.bandwidth
}

output "edition" {
  value = data.citrixadc_nscapacity.tf_nscapacity.edition
}

output "unit" {
  value = data.citrixadc_nscapacity.tf_nscapacity.unit
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `bandwidth` - System bandwidth limit.
* `edition` - Product edition. Possible values: [ Standard, Enterprise, Platinum ]
* `nodeid` - Unique number that identifies the cluster node.
* `password` - Password to use when authenticating with ADM Agent for LAS licensing.
* `platform` - Appliance platform type. Possible values: [ VS10, VE10, VP10, VS25, VE25, VP25, VS200, VE200, VP200, VS1000, VE1000, VP1000, VS3000, VE3000, VP3000, VS5000, VE5000, VP5000, VS8000, VE8000, VP8000, VS10000, VE10000, VP10000, VS15000, VE15000, VP15000, VS25000, VE25000, VP25000, VS40000, VE40000, VP40000, CP1000 ]
* `unit` - Bandwidth unit. Possible values: [ Gbps, Mbps ]
* `username` - Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `vcpu` - Licensed using vcpu pool.
* `id` - The id of the nscapacity. It is a system-generated identifier.

### Read-only nscapacity metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nscapacity` resource). They are Computed / GET-only, and any attribute the appliance does not return (for example, those tied to a different license type) is `null`.

* `actualbandwidth` - Bandwith in MBPS.
* `vcpucount` - Number of vCPUs licensed.
* `maxvcpucount` - Number of max vCPUs.
* `maxbandwidth` - Maximum Bandwidth.
* `minbandwidth` - Minimum Bandwidth.
* `instancecount` - VPX will consume one instance and MPX will consume zero instance.
* `daystoexpiration` - Days to expire.
