---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_fakeaccount_binding

The appfwprofile_fakeaccount_binding data source allows you to retrieve information about a fake-account detection binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_fakeaccount_binding" "tf_binding" {
  name           = "tf_appfwprofile"
  fakeaccount    = "email"
  formexpression = "^[a-z0-9._%+-]+@example\\.com$"
  formurl_fad    = "/register/submit"
  tag            = "signup_form"
}

output "state" {
  value = data.citrixadc_appfwprofile_fakeaccount_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_fakeaccount_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the fake-account rule is bound.
* `fakeaccount` - (Required) Field name of the fake account rule.
* `formexpression` - (Required) A regular expression that defines the fake account.
* `formurl_fad` - (Required) The fake account detection URL.
* `tag` - (Required) A tag expression that defines the fake account.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fakeaccount_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `fakeaccount:<fakeaccount>,formexpression:<formexpression>,formurl_fad:<formurl_fad>,name:<name>,tag:<tag>`.
* `isfieldnameregex` - Whether the fake-account detection field name is a regular expression. Possible values: [ REGEX, NOTREGEX ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the fake-account rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
```
