---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmlattachmenturl_binding

The appfwprofile_xmlattachmenturl_binding data source allows you to retrieve information about appfwprofile XML attachment URL bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
  name             = "tf_appfwprofile"
  xmlattachmenturl = ".*"
}

output "state" {
  value = data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding.state
}

output "xmlattachmentcontenttype" {
  value = data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding.xmlattachmentcontenttype
}

output "xmlmaxattachmentsize" {
  value = data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding.xmlmaxattachmentsize
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlattachmenturl` - (Required) XML attachment URL regular expression length.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlattachmenturl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
* `xmlattachmentcontenttype` - Specify content-type regular expression.
* `xmlattachmentcontenttypecheck` - State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.
* `xmlmaxattachmentsize` - Specify maximum attachment size.
* `xmlmaxattachmentsizecheck` - State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.
