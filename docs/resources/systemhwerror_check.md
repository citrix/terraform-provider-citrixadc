---
subcategory: "System"
---

# Resource: systemhwerror_check

Runs a hardware error check on the Citrix ADC as a one-shot diagnostic action. Use this resource when you want to trigger an on-demand disk error scan (for example, as part of a maintenance or troubleshooting workflow) and have the result surface in the ADC's diagnostic logs.

~>
* This is an **action-only** resource. NITRO exposes only the `check` action (`POST ?action=check`); there is no GET, update, or delete endpoint.
* **Apply runs the hardware/disk error check.** Creating the resource invokes the `check` action once. It is a side effect, not a managed object.
* Because `diskcheck` uses `RequiresReplace`, changing it forces the resource to be destroyed and recreated, which re-runs the check.
* **Read and Update are no-ops.** The check result cannot be read back through this resource; there is no NITRO GET endpoint, so the provider preserves prior state unchanged.
* **Destroy is state-only.** There is no inverse NITRO endpoint, so destroying the resource only removes it from Terraform state — it does not undo the check. Re-running `terraform apply` after a destroy (or after tainting the resource) runs the check again.
* **Import is not meaningful** for this resource because there is no underlying queryable object.


## Example usage

```hcl
resource "citrixadc_systemhwerror_check" "tf_systemhwerror_check" {
  diskcheck = true
}
```


## Argument Reference

* `diskcheck` - (Required) Perform only disk error checking. Changing this value forces a new resource to be created (re-runs the check).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemhwerror_check resource. It is a synthetic constant string `"systemhwerror_check"`. The ID is purely a Terraform state handle, not a NITRO lookup key (there is no GET endpoint for this resource).
