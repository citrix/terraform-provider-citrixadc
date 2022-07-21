---
subcategory: "GSLB"
---

# Resource: gslbservice_lbmonitor_binding

The gslbservice_lbmonitor_binding resource is used to create gslbservice_lbmonitor_binding.


## Example usage

```hcl
resource "citrixadc_gslbservice_lbmonitor_binding" "tf_gslbservice_lbmonitor_binding" {
  monitor_name = citrixadc_lbmonitor.tfmonitor1.monitorname
  monstate    = "DISABLED"
  servicename = citrixadc_gslbservice.tf_gslbservice.servicename
  weight      = "20"

}

resource "citrixadc_gslbservice" "tf_gslbservice" {
  ip          = "172.16.1.128"
  port        = "80"
  servicename = "test_gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.tf_gslbsite.sitename
}

resource "citrixadc_gslbsite" "tf_gslbsite" {
  sitename      = "test_sitename"
  siteipaddress = "10.222.70.79"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "test_monitor"
  type        = "HTTP"
}
```


## Argument Reference

* `servicename` - (Required) Name of the GSLB service.
* `monitor_name` - (Required) Monitor name.
* `monstate` - (Optional) State of the monitor bound to gslb service.
* `weight` - (Optional) Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservice_lbmonitor_binding is the concatenation of the `servicename` and `monitor_name` attributes separated by a comma.


## Import

A gslbservice_lbmonitor_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding test_gslb1vservice,test_monitor
```
