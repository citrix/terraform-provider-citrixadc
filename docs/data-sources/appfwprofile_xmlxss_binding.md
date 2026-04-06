---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmlxss_binding

The appfwprofile_xmlxss_binding data source allows you to retrieve information about an existing appfwprofile xmlxss binding.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmlxss_binding" "tf_binding_data" {
  name                    = "tf_appfwprofile"
  xmlxss                  = "tf_xmlxss"
  as_scan_location_xmlxss = "ELEMENT"
}

output "state" {
  value = data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data.state
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data.alertonly
}

output "isregex_xmlxss" {
  value = data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data.isregex_xmlxss
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlxss` - (Required) Exempt the specified URL from the XML cross-site scripting (XSS) check. An XML cross-site scripting exemption (relaxation) consists of the following items: * URL. URL to exempt, as a string or a PCRE-format regular expression. * ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string. * Location. ELEMENT if the attachment is located in an XML element, ATTRIBUTE if located in an XML attribute.
* `as_scan_location_xmlxss` - (Required) Location of XSS injection exception - XML Element or Attribute.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_xmlxss_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregex_xmlxss` - Is the XML XSS exempted field name a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
