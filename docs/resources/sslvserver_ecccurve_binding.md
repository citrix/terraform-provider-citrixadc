---
subcategory: "SSL"
---

# Resource: sslvserver_ecccurve_binding

The sslvserver_ecccurve_binding resource is used to add an ecc curve to ssl vserver.


## Example usage

```hcl
resource "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
	ecccurvename = "P_256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
	
}

resource "citrixadc_lbvserver" "tf_sslvserver" {
	name        = "tf_sslvserver"
	servicetype = "SSL"
}
```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to vserver/service. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `vservername` - (Required) Name of the SSL virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_ecccurve_binding. It is the concatenation of the `vservername` and `ecccurvename` attributes separated by a comma.


## Import

A sslvserver_ecccurve_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslvserver_ecccurve_binding.tf_sslvserver_ecccurve_binding tf_sslvserver,P_256
```
