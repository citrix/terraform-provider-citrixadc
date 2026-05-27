---
subcategory: "API Definition"
---

# Data Source: apiprofile

The apiprofile data source allows you to retrieve information about an existing Citrix ADC API profile.


## Example usage

```terraform
data "citrixadc_apiprofile" "tf_apiprofile" {
  name = "my_apiprofile"
}

output "apivisibility" {
  value = data.citrixadc_apiprofile.tf_apiprofile.apivisibility
}
```


## Argument Reference

* `name` - (Required) Name of the API profile to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `apivisibility` - Enable/Disable the schema lookup for the requests/apispecs that are bounded to the API profile. The default value of this parameter is `DISABLED`. Possible values: [ ENABLED, DISABLED ].
* `id` - The id of the apiprofile. It has the same value as the `name` attribute.
