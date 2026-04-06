---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmlsqlinjection_binding

The appfwprofile_xmlsqlinjection_binding data source allows you to retrieve information about an XML SQL injection binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding" {
  name                     = "tf_appfwprofile"
  xmlsqlinjection          = "hello"
  as_scan_location_xmlsql  = "ELEMENT"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding.comment
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlsqlinjection` - (Required) Exempt the specified URL from the XML SQL injection check. An XML SQL injection exemption (relaxation) consists of the following items: Name - Name to exempt, as a string or a PCRE-format regular expression. ISREGEX flag - REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string. Location - ELEMENT if the injection is located in an XML element, ATTRIBUTE if located in an XML attribute.
* `as_scan_location_xmlsql` - (Required) Location of SQL injection exception - XML Element or Attribute.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_xmlsqlinjection_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregex_xmlsql` - Is the XML SQL Injection exempted field name a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
