---
subcategory: "DNS"
---

# Resource: dnsaction

The dnsaction resource is used to create DNS action.


## Example usage

```hcl
resource "citrixadc_dnsaction" "dnsaction" {
	actionname       = "tf_action1"
	actiontype       = "ViewName"
	ipaddress        = ["192.0.2.20","192.0.2.56","198.51.100.10"]
	ttl              = 3600
	viewname         = "view1"
	preferredloclist = ["NA.tx.ns1.*.*.*","NA.tx.ns2.*.*.*","NA.tx.ns3.*.*.*"]
	dnsprofilename   = "tf_profile1"
  
  }
```


## Argument Reference

* `actionname` - (Required) Name of the dns action.
* `actiontype` - (Required) The type of DNS action that is being configured.
* `dnsprofilename` - (Optional) Name of the DNS profile to be associated with the transaction for which the action is chosen
* `ipaddress` - (Optional) List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type. 	    In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.
* `preferredloclist` - (Optional) The location list in priority order used for the given action.
* `ttl` - (Optional) Time to live, in seconds.
* `viewname` - (Optional) The view name that must be used for the given action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaction. It has the same value as the `actionname` attribute.


## Import

A dnsaction can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsaction.dnssaction tf_action1
```
