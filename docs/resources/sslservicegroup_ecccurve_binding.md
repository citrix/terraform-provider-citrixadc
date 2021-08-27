---
subcategory: "SSL"
---

# Resource: sslservicegroup_ecccurve_binding

The sslservicegroup_ecccurve_binding resource is used to add an ecc curve to ssl service group.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_ecccurve_binding" "tf_sslservicegroup_ecccurve_binding" {
	ecccurvename = "P_256"
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}
```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to servicegroup. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `servicegroupname` - (Required) The name of the SSL service to which the SSL policy needs to be bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_ecccurve_binding.  It is the concatenation of the `servicegroupname` and `ecccurvename` attributes separated by a comma.



## Import

A sslservicegroup_ecccurve_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslservicegroup_ecccurve_binding.tf_sslservicegroup_ecccurve_binding tf_servicegroup,P_256
```
