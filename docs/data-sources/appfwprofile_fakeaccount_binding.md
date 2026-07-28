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
* `formexpression` - (Optional) A regular expression that defines the fake account. Mutually exclusive with `formurl_fad`; set at most one of the two.
* `formurl_fad` - (Optional) The fake account detection URL. Mutually exclusive with `formexpression`; set at most one of the two.
* `tag` - (Required) A tag expression that defines the fake account.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fakeaccount_binding. It is a composite identifier composed of comma-separated `key:value` pairs. Only the populated arm of the mutually-exclusive `formexpression`/`formurl_fad` pair appears, so the format is `fakeaccount:<fakeaccount>,formexpression:<formexpression>,name:<name>,tag:<tag>` or `fakeaccount:<fakeaccount>,formurl_fad:<formurl_fad>,name:<name>,tag:<tag>`.
* `isfieldnameregex` - Whether the fake-account detection field name is a regular expression. Possible values: [ REGEX, NOTREGEX ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the fake-account rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
