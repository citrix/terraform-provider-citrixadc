---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmlwsiurl_binding

The appfwprofile_xmlwsiurl_binding data source allows you to retrieve information about an xmlwsiurl binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmlwsiurl_binding" "tf_binding" {
  name       = "tf_appfwprofile"
  xmlwsiurl  = ".*"
}

output "state" {
  value = data.citrixadc_appfwprofile_xmlwsiurl_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_xmlwsiurl_binding.tf_binding.comment
}

output "xmlwsichecks" {
  value = data.citrixadc_appfwprofile_xmlwsiurl_binding.tf_binding.xmlwsichecks
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlwsiurl` - (Required) XML WS-I URL regular expression length.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlwsiurl_binding. It is the composite identifier in the format `<name>,<xmlwsiurl>`.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
* `xmlwsichecks` - Specify a comma separated list of relevant WS-I rule IDs. (R1140, R1141)
