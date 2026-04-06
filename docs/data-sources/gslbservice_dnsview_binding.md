---
subcategory: "GSLB"
---

# Data Source: gslbservice_dnsview_binding

The gslbservice_dnsview_binding data source allows you to retrieve information about a GSLB service DNS view binding.


## Example Usage

```terraform
data "citrixadc_gslbservice_dnsview_binding" "tf_gslbservice_dnsview_binding" {
  servicename = "gslb1vservice"
  viewname    = "view4"
}

output "viewip" {
  value = data.citrixadc_gslbservice_dnsview_binding.tf_gslbservice_dnsview_binding.viewip
}

output "servicename" {
  value = data.citrixadc_gslbservice_dnsview_binding.tf_gslbservice_dnsview_binding.servicename
}
```


## Argument Reference

* `servicename` - (Required) Name of the GSLB service.
* `viewname` - (Required) Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservice_dnsview_binding. It is the concatenation of the `servicename` and `viewname` attributes separated by a comma.
* `viewip` - IP address to be used for the given view
