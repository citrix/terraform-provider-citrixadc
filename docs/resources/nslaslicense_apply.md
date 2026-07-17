---
subcategory: "NS"
---

# Resource: nslaslicense_apply

The nslaslicense_apply resource applies a local-area-services (LAS) / fixed-bandwidth license file that has already been staged on the Citrix ADC, activating the licensed capacity it grants. Use it to bring a previously uploaded license file into effect on the appliance.

!> **DISRUPTIVE / NON-IDEMPOTENT.** This resource maps to the NITRO `apply` action (`POST ?action=apply`). Applying it alters the licensed capacity of the appliance and is **not idempotent** — every create or replace re-applies the license. NITRO exposes no get/add/update/delete endpoint for this object, so the applied license is **not readable back**: Read is a no-op, drift cannot be detected, and Delete only removes the resource from Terraform state without affecting the appliance. Treat this resource as a one-shot operational action rather than ordinary declarative configuration.

~> **Note.** All attributes are marked `RequiresReplace`, so changing any of them forces the apply action to run again as a replacement. Import is not meaningful for this resource because there is no server-side object to read back.


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nslaslicense_apply`. It does not correspond to any object on the Citrix ADC.
