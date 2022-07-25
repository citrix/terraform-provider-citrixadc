---
subcategory: "DNS"
---

# Resource: dnspolicylabel_dnspolicy_binding

The dnspolicylabel_dnspolicy_binding resource is used to create DNS policylabel policy binding


## Example usage

```hcl
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
}
resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
  labelname = "blue_label"
  transform = "dns_req"

}
resource "citrixadc_dnspolicylabel_dnspolicy_binding" "dnspolicylabel_dnspolicy_binding" {
  labelname = citrixadc_dnspolicylabel.dnspolicylabel.labelname
  policyname = citrixadc_dnspolicy.dnspolicy.name
  priority = 2

}

```


## Argument Reference

* `policyname` - (Required) The dns policy name.
* `labelname` - (Required) Name of the dns policy label.
* `priority` - (Requireed) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label invocation.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnspolicylabel_dnspolicy_binding. It has the same value as the `policyname,labelname` attributes.


## Import

A dnspolicylabel_dnspolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_dnspolicylabel_dnspolicy_binding.dnspolicylabel_dnspolicy_binding policy_A,blue_label
```
