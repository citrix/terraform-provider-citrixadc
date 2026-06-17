---
subcategory: "NS"
---

# Resource: nslimitselector

Configures a rate limit selector on the Citrix ADC. A limit selector defines a set of request characteristics (expressed as default-syntax expressions) used to group traffic for rate limiting. The selector is referenced by a rate limit identifier (`nslimitidentifier`) so that requests sharing the same selected field values are counted together against a threshold.

~> **Note** The classic `nslimitselector` (the `add ns limitSelector` CLI command) is deprecated in favor of the stream selector (`add stream selector`), but it remains functional and is still supported by this resource.


## Example usage

```hcl
resource "citrixadc_nslimitselector" "tf_nslimitselector" {
  selectorname = "tf_limitselector"
  rule = [
    "HTTP.REQ.URL",
    "CLIENT.IP.SRC",
  ]
}
```


## Argument Reference

* `selectorname` - (Required) Name for the rate limit selector. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Changing this value forces a new resource to be created.
* `rule` - (Required) List of default-syntax expressions that identify the request fields tracked by the selector (for example `HTTP.REQ.URL` or `CLIENT.IP.SRC`). Requests are grouped by the combined values of these expressions. This attribute is updateable in place.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslimitselector. It has the same value as the `selectorname` attribute.


## Import

A nslimitselector can be imported using its selectorname, e.g.

```shell
terraform import citrixadc_nslimitselector.tf_nslimitselector tf_limitselector
```
