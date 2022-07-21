---
subcategory: "GSLB"
---

# Resource: gslbservice_dnsview_binding

The gslbservice_dnsview_binding resource is used to create gslbservice_dnsview_binding.


## Example usage

```hcl
resource "citrixadc_gslbservice_dnsview_binding" "tf_gslbservice_dnsview_binding" {
  servicename = citrixadc_gslbservice.gslb_svc1.servicename
  viewname    = citrixadc_dnsview.tf_dnsview.viewname
  viewip      = "192.168.2.1"
}

resource "citrixadc_gslbsite" "site_remote" {
  sitename        = "Site-Remote"
  siteipaddress   = "172.31.48.18"
  sessionexchange = "ENABLED"
}

resource "citrixadc_gslbservice" "gslb_svc1" {
  ip          = "172.16.1.121"
  port        = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_remote.sitename
}

resource "citrixadc_dnsview" "tf_dnsview" {
  viewname = "view1"
}
```


## Argument Reference

* `servicename` - (Required) Name of the GSLB service.
* `viewname` - (Required) Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.
* `viewip` - (Optional) IP address to be used for the given view


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservice_dnsview_binding. It is the concatenation of the `servicename` and `viewname` attributes separated by a comma.


## Import

A gslbservice_dnsview_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservice_dnsview_binding.tf_gslbservice_dnsview_binding gslb1vservice,view1
```
