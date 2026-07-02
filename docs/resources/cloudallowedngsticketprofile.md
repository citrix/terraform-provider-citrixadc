---
subcategory: "Cloud"
---

# Resource: cloudallowedngsticketprofile

Configures a profile that identifies the set of next-generation service (NGS) tickets the Citrix ADC is allowed to accept in a cloud deployment. Create one when you need to authorize a named group of allowed tickets and, optionally, record who created the profile.


## Example usage

```hcl
resource "citrixadc_cloudallowedngsticketprofile" "tf_cloudallowedngsticketprofile" {
  name    = "allowed-tickets-prod"
  creator = "cloud-onboarding-team"
}
```


## Argument Reference

* `name` - (Required) Profile name for allowed tickets. Cannot be changed after the profile is created; changing this value forces a new resource. Maximum length = 127
* `creator` - (Optional) Created name for allowed tickets. This value can be updated in place. Maximum length = 255


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudallowedngsticketprofile. It has the same value as the `name` attribute.


## Import

A cloudallowedngsticketprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile allowed-tickets-prod
```
