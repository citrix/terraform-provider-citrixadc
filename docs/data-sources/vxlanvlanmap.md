---
subcategory: "Network"
---

# Data Source: vxlanvlanmap

The vxlanvlanmap data source allows you to retrieve information about a VXLAN VLAN mapping table.

## Example usage

```terraform
data "citrixadc_vxlanvlanmap" "example" {
  name = "tf_vxlanvlanmp"
}

output "name" {
  value = data.citrixadc_vxlanvlanmap.example.name
}
```

## Argument Reference

* `name` - (Required) Name of the mapping table.

## Attribute Reference

In addition to the arguments, the following attributes are available:

No additional attributes are available for this resource beyond the name identifier.
