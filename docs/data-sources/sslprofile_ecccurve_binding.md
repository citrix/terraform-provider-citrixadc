---
subcategory: "SSL"
---

# Data Source: sslprofile_ecccurve_binding

The sslprofile_ecccurve_binding data source allows you to retrieve information about a specific binding between an SSL profile and an ECC curve.


## Example Usage

```terraform
data "citrixadc_sslprofile_ecccurve_binding" "tf_bind" {
  name         = citrixadc_sslprofile.tf_sslprofile.name
  ecccurvename = "P_256"
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_ecccurve_binding.tf_bind.cipherpriority
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `ecccurvename` - (Required) Named ECC curve bound to vserver/service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipherpriority` - Priority of the cipher binding.
* `id` - The id of the sslprofile_ecccurve_binding. It is the concatenation of both `name` and `ecccurvename` attributes.

