---
subcategory: "DNS"
---

# Data Source: citrixadc_dnspolicylabel

This data source is used to retrieve information about an existing DNS policy label.

## Example Usage

```hcl
data "citrixadc_dnspolicylabel" "example" {
  labelname = "blue_label"
}
```

## Argument Reference

* `labelname` - (Required) Name of the DNS policy label.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the DNS policy label (same as `labelname`).
* `newname` - The new name of the DNS policy label.
* `transform` - The type of transformations allowed by the policies bound to the label.
