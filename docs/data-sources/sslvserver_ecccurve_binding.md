---
subcategory: "SSL"
---

# Data Source: sslvserver_ecccurve_binding

The sslvserver_ecccurve_binding data source allows you to retrieve information about the ECC curve binding to an SSL virtual server.


## Example Usage

```terraform
data "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
  vservername  = "tf_sslvserver"
  ecccurvename = "P_256"
}

output "ecccurvename" {
  value = data.citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding.ecccurvename
}

output "vservername" {
  value = data.citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding.vservername
}
```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to vserver/service.
* `vservername` - (Required) Name of the SSL virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_ecccurve_binding. It is the concatenation of the `vservername` and `ecccurvename` attributes separated by a comma.
