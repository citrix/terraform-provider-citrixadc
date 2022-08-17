---
subcategory: "AAA"
---

# Resource: aaaglobal_aaapreauthenticationpolicy_binding

The aaaglobal_aaapreauthenticationpolicy_binding resource is used to create aaaglobal_aaapreauthenticationpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaaglobal_aaapreauthenticationpolicy_binding" "tf_aaaglobal_aaapreauthenticationpolicy_binding" {
  policy    = "my_preauthentication_policy"
  priority  = 50
  }
```


## Argument Reference

* `policy` - (Required) Name of the policy to be unbound. Minimum length =  1
* `priority` - (Required) Priority of the bound policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaglobal_aaapreauthenticationpolicy_binding. It has the same value as the `policy` attribute.


## Import

A aaaglobal_aaapreauthenticationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_aaaglobal_aaapreauthenticationpolicy_binding my_preauthentication_policy
```
