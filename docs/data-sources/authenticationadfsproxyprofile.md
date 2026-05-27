---
subcategory: "Authentication"
---

# Data Source: authenticationadfsproxyprofile

The authenticationadfsproxyprofile data source allows you to retrieve information about an existing ADFS proxy profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_authenticationadfsproxyprofile" "tf_adfsproxyprofile" {
  name = "example_adfsproxyprofile"
}

output "serverurl" {
  value = data.citrixadc_authenticationadfsproxyprofile.tf_adfsproxyprofile.serverurl
}

output "username" {
  value = data.citrixadc_authenticationadfsproxyprofile.tf_adfsproxyprofile.username
}
```


## Argument Reference

* `name` - (Required) Name of the ADFS proxy profile to look up. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `certkeyname` - SSL certificate of the proxy that is registered at the ADFS server for trust.
* `serverurl` - Fully qualified URL of the ADFS server.
* `username` - Name of an account in the directory that is used to authenticate the trust request from the Citrix ADC acting as a proxy.
* `id` - The id of the authenticationadfsproxyprofile. It has the same value as the `name` attribute.
