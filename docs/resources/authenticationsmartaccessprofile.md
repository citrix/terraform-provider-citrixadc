---
subcategory: "Authentication"
---

# Resource: authenticationsmartaccessprofile

The authenticationsmartaccessprofile resource is used to create and manage Citrix ADC authentication Smartaccess profiles.


## Example usage

```hcl
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "tag1"
  comment = "Smartaccess profile managed by Terraform"
}
```


## Argument Reference

* `name` - (Required) Name of the Smartaccess profile. Cannot be changed after the profile is created.
* `tags` - (Required) The tag that is associated with the Smartaccess profile.
* `comment` - (Optional) Optional comment for the profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsmartaccessprofile. It has the same value as the `name` attribute.


## Import

An authenticationsmartaccessprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile tf_authenticationsmartaccessprofile
```
