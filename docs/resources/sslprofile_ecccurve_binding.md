---
subcategory: "SSL"
---

# Resource: sslprofile_ecccurve_binding

This resource is used to bind ECC curves to an SSL profile.


## Example usage

```hcl

resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"
}

resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
  name                             = citrixadc_sslprofile.tf_sslprofile.name
  ecccurvename                     = ["X_25519", "P_521", "P_384"]
  remove_existing_ecccurve_binding = true
}

```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to vserver/service. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `name` - (Required) Name of the SSL profile. Minimum length =  1 Maximum length =  127
* `remove_existing_ecccurve_binding` - (Optional) If you want to unbind all the existing ecccurve bindings, then set this as true else false.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_ecccurve_binding. It has is the conatenation of the `name` and `ecccurvename` attributes.
