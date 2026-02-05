---
subcategory: "DNS"
---

# Data Source `dnsview`

The dnsview data source allows you to retrieve information about DNS views.


## Example usage

```terraform
data "citrixadc_dnsview" "tf_dnsview" {
  viewname = "view3"
}

output "viewname" {
  value = data.citrixadc_dnsview.tf_dnsview.viewname
}
```


## Argument Reference

* `viewname` - (Required) Name for the DNS view.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsview. It has the same value as the `viewname` attribute.
