---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_jsonblockkeyword_binding

The appfwprofile_jsonblockkeyword_binding data source allows you to retrieve information about an existing JSON block-keyword binding on an application firewall profile.

## Example usage

```terraform
data "citrixadc_appfwprofile_jsonblockkeyword_binding" "tf_binding" {
  name                      = "tf_appfwprofile"
  jsonblockkeyword          = "passwd"
  keyname_json_blockkeyword = "user.credentials"
  jsonblockkeywordurl       = "/api/v1/login"
}

output "state" {
  value = data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_binding.resourceid
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which the exemption or rule is bound.
* `jsonblockkeyword` - (Required) Field name of the JSON block keyword binding.
* `keyname_json_blockkeyword` - (Required) JSON block keyword keyname (the JSON key under which the keyword is matched).
* `jsonblockkeywordurl` - (Required) The JSON block keyword URL on which the keyword is inspected.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the appfwprofile_jsonblockkeyword_binding. It is a composite of comma-separated `key:value` pairs in the order `jsonblockkeyword:<value>,jsonblockkeywordurl:<value>,keyname_json_blockkeyword:<value>,name:<value>`.
* `iskeyregex_json_blockkeyword` - Is the JSON block keyword key a regular expression?
* `jsonblockkeywordtype` - JSON block keyword type.
* `state` - Whether the binding is enabled.
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by a dynamic profile?
* `alertonly` - Send SNMP alert?
* `resourceid` - A server-assigned identifier that identifies the rule.
