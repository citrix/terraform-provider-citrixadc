---
subcategory: "NS"
---

# citrixadc_nsxmlnamespace (Data Source)

Data source for querying Citrix ADC XML namespaces. This data source retrieves information about XML namespace prefixes configured on the ADC appliance.

## Example Usage

```hcl
data "citrixadc_nsxmlnamespace" "example" {
  prefix = "my_namespace"
}

# Output XML namespace attributes
output "namespace_uri" {
  value = data.citrixadc_nsxmlnamespace.example.namespace
}

output "description" {
  value = data.citrixadc_nsxmlnamespace.example.description
}
```

## Argument Reference

The following arguments are supported:

* `prefix` - (Required) XML namespace prefix. This is used to identify the namespace in XML documents and policy expressions.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nsxmlnamespace datasource.
* `namespace` - The XML namespace URI associated with the prefix.
* `description` - Description of the XML namespace.

## Notes

XML namespaces are used in policy expressions to work with XML content. They allow you to define prefix-to-URI mappings that can be referenced when parsing and manipulating XML data in policies.
