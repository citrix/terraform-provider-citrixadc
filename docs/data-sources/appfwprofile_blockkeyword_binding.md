---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_blockkeyword_binding

The appfwprofile_blockkeyword_binding data source allows you to retrieve information about a block keyword binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_blockkeyword_binding" "tf_binding" {
  name                    = "tf_appfwprofile"
  blockkeyword            = "drop table"
  fieldname               = "comments"
  as_blockkeyword_formurl = "https://www.example.com/submit"
}

output "state" {
  value = data.citrixadc_appfwprofile_blockkeyword_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_blockkeyword_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the block keyword is bound.
* `blockkeyword` - (Required) The block keyword bound to the profile.
* `fieldname` - (Required) Name of the form field in which the block keyword is evaluated.
* `as_blockkeyword_formurl` - (Required) The form action URL on which the block keyword rule applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_blockkeyword_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `as_blockkeyword_formurl:<as_blockkeyword_formurl>,blockkeyword:<blockkeyword>,fieldname:<fieldname>,name:<name>`.
* `as_fieldname_isregex_blockkeyword` - Whether the block keyword field name is a regular expression. Possible values: [ REGEX, NOTREGEX ]
* `blockkeywordtype` - The type of the block keyword. Possible values: [ literal, PCRE ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the block keyword rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
