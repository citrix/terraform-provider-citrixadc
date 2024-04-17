---
subcategory: "Load Balancing"
---

# Resource: lbaction

The lbaction resource is used to configure lbaction resource.


## Example usage

```hcl
resource "citrixadc_lbaction" "tf_lbact" {
  name  = "tf_lbact"
  type  = "SELECTIONORDER"
  value = [1]
}

```


## Argument Reference

* `name` - (Required) Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb action" or 'my lb action').
* `type` - (Required) Type of an LB action. Available settings function as follows: * NOLBACTION - Does not consider LB action in making LB decision. * SELECTIONORDER - services bound to vserver with order specified in value parameter is considerd for lb/gslb decision.
* `comment` - (Optional) Comment. Any type of information about this LB action.
* `newname` - (Optional) New name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb action" or my lb action').
* `value` - (Optional) The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on. For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and  service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbaction resource. It has the same value as the `name` attribute.


## Import

A lbaction resource can be imported using its name, e.g.

```shell
terraform import citrixadc_lbaction.tf_lbact tf_lbact
```
