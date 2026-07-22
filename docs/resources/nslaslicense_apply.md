---
subcategory: "NS"
---

# Resource: nslaslicense_apply

The nslaslicense_apply resource applies a local-area-services (LAS) / fixed-bandwidth license file that has already been staged on the Citrix ADC, activating the licensed capacity it grants. Use it to bring a previously uploaded license file into effect on the appliance.

!> **DISRUPTIVE / NON-IDEMPOTENT.** Applying this resource alters the licensed capacity of the appliance and is **not idempotent** — every create or replace re-applies the license. Treat this resource as a one-shot operational action rather than ordinary declarative configuration.

~> **Note.** This resource is immutable, so changing any of its attributes forces the apply action to run again as a replacement. Import is not meaningful for this resource.


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
