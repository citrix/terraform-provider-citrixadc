---
subcategory: "GSLB"
---

# Data Source: gslbvserver_gslbservice_binding

The gslbvserver_gslbservice_binding data source allows you to retrieve information about a GSLB virtual server to GSLB service binding.

## Example usage

```terraform
data "citrixadc_gslbvserver_gslbservice_binding" "tf_gslbvserver_gslbservice_binding" {
  name        = "gslb_vserver"
  servicename = "gslb1vservice"
}

output "name" {
  value = data.citrixadc_gslbvserver_gslbservice_binding.tf_gslbvserver_gslbservice_binding.name
}

output "servicename" {
  value = data.citrixadc_gslbvserver_gslbservice_binding.tf_gslbvserver_gslbservice_binding.servicename
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `servicename` - (Required) Name of the GSLB service for which to change the weight.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_gslbservice_binding. It is a composite identifier in the format "name,servicename".
* `domainname` - Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `weight` - Weight for the service.

### Read-only gslbvserver_gslbservice_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_gslbvserver_gslbservice_binding` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `dynamicconfwt` - Weight obtained by virtue of bound service count or weight.
* `cumulativeweight` - Cumulative weight of the GSLB service considering both its configured weight and dynamic weight.
* `sitepersistcookie` - Cookie displayed for site persistence in a cluster setup.
* `orderstr` - Order number in string form assigned to the service when it is bound to the lb vserver.
* `gslbthreshold` - Indicates whether the GSLB service has reached its threshold.
* `preferredlocation` - The target site to be returned in the DNS response when a policy is successfully evaluated against the incoming DNS request.
* `svcsitepersistence` - Type of Site Persistence set on the bound service (`ConnectionProxy`, `HTTPRedirect`, `NONE`).
* `gslbboundsvctype` - Protocol used by services bound to the GSLB virtual server.
* `ipaddress` - IP address of the bound GSLB service.
* `iscname` - Whether the cname feature is set on the vserver (`ENABLED`, `DISABLED`).
* `thresholdvalue` - Indicates whether the threshold has been exceeded for this service participating in CUSTOMLB.
* `port` - Port number of the bound GSLB service.
* `curstate` - State of the GSLB vserver.
* `svreffgslbstate` - Effective state of the GSLB service.
* `cnameentry` - The cname of the GSLB service.
