---
subcategory: "SSL"
---

# Resource: sslprofile_ecccurve_binding

The sslprofile_ecccurve_binding resource is used to create bindings between sslprofiles and ecccurves.

~>  If you are using this resource to bind ecccurves to a sslprofile, do not define the `ecccurvebindings` attribute in the sslprofile resource.

~>  The attribute `remove_existing_ecccurve_binding` should be `true` if you want to delete all the exiting bindings, and bind the new ecccurve to sslprofile.



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
