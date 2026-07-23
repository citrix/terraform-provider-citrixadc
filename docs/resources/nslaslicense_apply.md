---
subcategory: "NS"
---

# Resource: nslaslicense_apply

This resource is used to apply a staged LAS / fixed-bandwidth license file on the Citrix ADC.

!> **DISRUPTIVE / NON-IDEMPOTENT:** Applying this resource alters the licensed capacity and re-applies the license on every create/replace. Treat it as a one-shot operational action.


## Example usage

```hcl
resource "citrixadc_nslaslicense_apply" "tf_nslaslicense_apply" {
  filename       = "CNS_V3000_SERVER_PLT_Retail.lic"
  filelocation   = "/nsconfig/license/"
  fixedbandwidth = true
}
```


## Argument Reference

* `filename` - (Required) Name of the file. It should not include filepath. Changing this value forces the resource to be replaced.
* `filelocation` - (Required) Location of the file on Citrix ADC. Changing this value forces the resource to be replaced.
* `fixedbandwidth` - (Optional) Apply fixed bandwidth license on ADC. Changing this value forces the resource to be replaced.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslaslicense_apply resource. It is set to `nslaslicense_apply`.
