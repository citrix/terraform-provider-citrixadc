---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationauthnprofile

The `citrixadc_authenticationauthnprofile` data source is used to retrieve information about an existing Authentication Profile configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve an authentication profile by name
data "citrixadc_authenticationauthnprofile" "example" {
  name = "demo_authnprofile"
}

# Use the retrieved data in other resources
output "profile_authnvsname" {
  value = data.citrixadc_authenticationauthnprofile.example.authnvsname
}

output "profile_authenticationhost" {
  value = data.citrixadc_authenticationauthnprofile.example.authenticationhost
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the authentication profile to retrieve. This is the unique identifier for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the authentication profile. It has the same value as the `name` attribute.
* `authenticationdomain` - Domain for which TM cookie must to be set. If unspecified, cookie will be set for FQDN.
* `authenticationhost` - Hostname of the authentication vserver to which user must be redirected for authentication.
* `authenticationlevel` - Authentication weight or level of the vserver to which this will bound. This is used to order TM vservers based on the protection required. A session that is created by authenticating against TM vserver at given level cannot be used to access TM vserver at a higher level.
* `authnvsname` - Name of the authentication vserver at which authentication should be done.
