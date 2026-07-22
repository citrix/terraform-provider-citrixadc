---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_appfwconfidfield_binding

The appfwprofile_appfwconfidfield_binding data source allows you to retrieve information about a confidential-field binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_appfwconfidfield_binding" "tf_binding" {
  name        = "tf_appfwprofile"
  confidfield = "password"
  cffield_url = "https://www.example.com/login"
}

output "state" {
  value = data.citrixadc_appfwprofile_appfwconfidfield_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_appfwconfidfield_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the confidential field is bound.
* `confidfield` - (Required) Name of the form field designated as confidential.
* `cffield_url` - (Required) URL of the web page that contains the web form.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_appfwconfidfield_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `cffield_url:<cffield_url>,confidfield:<confidfield>,name:<name>`.
* `isregex_cffield` - Whether the confidential field name is a regular expression. Possible values: [ REGEX, NOTREGEX ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the confidential-field rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `resourceid` - A system-generated identifier that identifies the rule.
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `alertonly` - Indicates whether an SNMP alert is sent.
