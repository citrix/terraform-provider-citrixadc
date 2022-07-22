---
subcategory: "DNS"
---

# Resource: dnspolicy

The dnspolicy resource is used to create DNS policy.


## Example usage

```hcl
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
  }
```


## Argument Reference

* `name` - (Required) Name for the DNS policy.
* `rule` - (Required) Expression against which DNS traffic is evaluated. Note: * On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks. * If the expression itself includes double quotation marks, you must escape the quotations by using the  character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.  Example: CLIENT.UDP.DNS.DOMAIN.EQ("domainname")
* `actionname` - (Optional) Name of the DNS action to perform when the rule evaluates to TRUE. The built in actions function as follows: * dns_default_act_Drop. Drop the DNS request. * dns_default_act_Cachebypass. Bypass the DNS cache and forward the request to the name server. You can create custom actions by using the add dns action command in the CLI or the DNS > Actions > Create DNS Action dialog box in the Citrix ADC configuration utility.
* `cachebypass` - (Optional) By pass dns cache for this.
* `drop` - (Optional) The dns packet must be dropped.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `preferredlocation` - (Optional) The location used for the given policy. This is deprecated attribute. Please use -prefLocList
* `preferredloclist` - (Optional) The location list in priority order used for the given policy.
* `viewname` - (Optional) The view name that must be used for the given policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnspolicy. It has the same value as the `name` attribute.


## Import

A dnspolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_dnspolicy.dnspolicy policy_A
```
