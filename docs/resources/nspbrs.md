---
subcategory: "NS"
---

# Resource: nspbrs

The nspbrs resource is used to configure Policy based routing resource.


## Example usage

```hcl
resource "citrixadc_nspbrs" "name" {
  action = "apply"
}
```

## Argument Reference

* `action` - (Required) actions to perform to configure Policy based routing resource. Possible values `apply`, `clear` and `renumber`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `nspbrs`. It is a unique string prefixed with `tf-nspbrs-`.

