---
subcategory: "SSL"
---

# Resource: sslprofile_ecccurve_binding

The sslprofile_ecccurve_binding resource is used to create bindings between sslprofiles and ecccurves.

~>  If you are using this resource to bind ecccurves to a sslprofile, do not define the `ecccurvebindings` attribute in the sslprofile resource.

~>  Set the attribute `remove_existing_ecccurve_binding` to `true` if you want to delete all the existing ecccurve bindings before binding the configured curves to the sslprofile.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"
}

resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
  name                             = citrixadc_sslprofile.tf_sslprofile.name
  ecccurvename                     = ["P_256", "P_384", "P_521"]
  remove_existing_ecccurve_binding = true
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile. Minimum length = 1 Maximum length = 127
* `ecccurvename` - (Required) List of named ECC curves to bind to the SSL profile. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `remove_existing_ecccurve_binding` - (Required) Remove all existing ecccurve bindings from the SSL profile before binding the configured curves. Possible values: [ true, false ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_ecccurve_binding. It is the concatenation of the `name` and `ecccurvename` attributes separated by a comma.


## Import

A sslprofile_ecccurve_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding tf_sslprofile,P_256
```
