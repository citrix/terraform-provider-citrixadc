---
subcategory: "Cloud"
---

# Data Source: cloudallowedngsticketprofile

The cloudallowedngsticketprofile data source allows you to retrieve information about an existing allowed NGS ticket profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_cloudallowedngsticketprofile" "example" {
  name = "allowed-tickets-prod"
}

output "cloudallowedngsticketprofile_creator" {
  value = data.citrixadc_cloudallowedngsticketprofile.example.creator
}
```


## Argument Reference

* `name` - (Required) Profile name for allowed tickets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudallowedngsticketprofile. It has the same value as the `name` attribute.
* `creator` - Created name for allowed tickets.
