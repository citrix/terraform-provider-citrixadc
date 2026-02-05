---
subcategory: "Autoscale"
---

# Data Source `autoscaleprofile`

The autoscaleprofile data source allows you to retrieve information about autoscale profiles.


## Example usage

```terraform
data "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name = "my_autoscaleprofile"
}

output "type" {
  value = data.citrixadc_autoscaleprofile.tf_autoscaleprofile.type
}

output "url" {
  value = data.citrixadc_autoscaleprofile.tf_autoscaleprofile.url
}
```


## Argument Reference

* `name` - (Required) AutoScale profile name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `apikey` - API key for authentication with service.
* `sharedsecret` - Shared secret for authentication with service.
* `type` - The type of profile.
* `url` - URL providing the service.
* `id` - The id of the autoscaleprofile. It has the same value as the `name` attribute.
