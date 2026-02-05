---
subcategory: "Load Balancing"
---

# Data Source `lbaction`

The lbaction data source allows you to retrieve information about an existing lbaction.


## Example usage

```terraform
data "citrixadc_lbaction" "tf_lbaction" {
  name = "my_lbaction"
}

output "name" {
  value = data.citrixadc_lbaction.tf_lbaction.name
}

output "type" {
  value = data.citrixadc_lbaction.tf_lbaction.type
}

output "value" {
  value = data.citrixadc_lbaction.tf_lbaction.value
}
```


## Argument Reference

* `name` - (Required) Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbaction. It has the same value as the `name` attribute.
* `comment` - Comment. Any type of information about this LB action.
* `newname` - New name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `type` - Type of an LB action. Available settings function as follows:
  * `NOLBACTION` - Does not consider LB action in making LB decision.
  * `SELECTIONORDER` - Services bound to vserver with order specified in value parameter is considered for lb/gslb decision.
* `value` - The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on. For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.
