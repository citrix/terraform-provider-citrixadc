---
subcategory: "Cloud"
---

# Resource: cloudallowedngsticketprofile

This resource is used to manage cloud allowed NGS ticket profiles.


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
