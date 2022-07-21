---
subcategory: "Cache Reduction"
---

# Resource: crpolicy

The crpolicy resource is used to create CR policy.


## Example usage

```hcl
resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}

```


## Argument Reference

* `policyname` - (Required) Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI:  If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `rule` - (Required) Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: *  If the expression includes one or more spaces, enclose the entire expression in double quotation marks. *  If the expression itself includes double quotation marks, escape the quotations by using the \ character.  *  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Optional) Name of the built-in cache redirection action: CACHE/ORIGIN.
* `logaction` - (Optional) The log action associated with the cache redirection policy

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crpolicy. It has the same value as the `policyname` attribute.


## Import

A crpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_crpolicy.crpolicy crpolicy1
```
