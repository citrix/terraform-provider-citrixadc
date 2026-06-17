---
subcategory: "NS"
---

# Resource: nslimitsessions

The nslimitsessions resource flushes (clears) the active rate-limit sessions tracked on the Citrix ADC for a given rate-limit identifier. Use it when you want to reset the accumulated hit/drop counters and per-selectlet session state for a limit identifier so that throttling decisions start from a clean slate.

~> **One-shot action.** This resource maps to the NITRO `clear` action (`POST ?action=clear`); it does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the clear once. There is no readable server-side object: Read is a no-op, Delete only removes the resource from Terraform state, and changing `limitidentifier` forces a new clear (replacement).


## Example usage

```hcl
resource "citrixadc_nslimitsessions" "tf_nslimitsessions" {
  limitidentifier = "myratelimit"
}
```


## Argument Reference

* `limitidentifier` - (Required) Name of the rate limit identifier for which to display the sessions. Changing this value forces the resource to be replaced (re-running the clear action against the new identifier).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslimitsessions resource. It is a synthetic identifier that has the same value as the `limitidentifier` attribute.


## Import

An nslimitsessions resource can be imported using the limit identifier, e.g.

```shell
terraform import citrixadc_nslimitsessions.tf_nslimitsessions myratelimit
```
