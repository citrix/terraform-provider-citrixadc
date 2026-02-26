---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_contenttype_binding

The appfwprofile_contenttype_binding data source allows you to retrieve information about a contenttype binding to an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_contenttype_binding" "tf_binding" {
  name        = "tf_appfwprofile"
  contenttype = "hello"
}

output "state" {
  value = data.citrixadc_appfwprofile_contenttype_binding.tf_binding.state
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_contenttype_binding.tf_binding.alertonly
}

output "comment" {
  value = data.citrixadc_appfwprofile_contenttype_binding.tf_binding.comment
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `contenttype` - (Required) A regular expression that designates a content-type on the content-types list.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_contenttype_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
