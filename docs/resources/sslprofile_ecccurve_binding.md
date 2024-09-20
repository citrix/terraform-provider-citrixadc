---
subcategory: "SSL"
---

# Resource: sslprofile_ecccurve_binding

The sslprofile_ecccurve_binding resource is used to create bindings between sslprofiles and ecccurves.

~>  * If you are using this resource to bind ecccurves to a sslprofile, do not define the `ecccurvebindings` attribute in the sslprofile resource.

~>  * The order of ecccurve is importance, to maintain the order you need to create dependencies between multiple `citrixadc_sslprofile_ecccurve_binding` resources, as shown below.

~>  * The attribute `remove_existing_ecccurve_binding` should only be used for first resource, as it unbinds existing bindings to create only the single mentioned binding



## Example usage

```hcl

resource "citrixadc_sslprofile" "tf_sslprofile" {
    name = "tf_sslprofile"
}

resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
    name                                = citrixadc_sslprofile.tf_sslprofile.name
    ecccurvename                        = "P_256"
    remove_existing_ecccurve_binding    = true
}

resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding_1" {
    name         = citrixadc_sslprofile.tf_sslprofile.name
    ecccurvename = "P_384"
    depends_on   = [citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding]
}

resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding_2" {
    name         = citrixadc_sslprofile.tf_sslprofile.name
    ecccurvename = "P_521"
    depends_on   = [citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding_1]
}
```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to vserver/service. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `name` - (Required) Name of the SSL profile. Minimum length =  1 Maximum length =  127
* `remove_existing_ecccurve_binding` - (Optional) If you want to unbind all the existing ecccurve bindings, then set this as true else false. Note: If you want to bind multiple ecccurve to sslprofile as shown in above example, then this attribute shoild be true to only first resource.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_ecccurve_binding. It has is the conatenation of the `name` and `ecccurvename` attributes.


## Import

A sslprofile_ecccurve_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding tf_sslprofile,P_256
terraform import citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding_1 tf_sslprofile,P_384
terraform import citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding_2 tf_sslprofile,P_521
```
