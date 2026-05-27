---
subcategory: "Authentication"
---

# Data Source: authenticationsmartaccessprofile

The authenticationsmartaccessprofile data source allows you to retrieve information about a Citrix ADC authentication Smartaccess profile.


## Example usage

```terraform
data "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name = "my_smartaccess_profile"
}

output "tags" {
  value = data.citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.tags
}
```


## Argument Reference

* `name` - (Required) Name of the Smartaccess profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `tags` - The tag that is associated with the Smartaccess profile.
* `comment` - Optional comment for the profile.
* `id` - The id of the authenticationsmartaccessprofile. It has the same value as the `name` attribute.
