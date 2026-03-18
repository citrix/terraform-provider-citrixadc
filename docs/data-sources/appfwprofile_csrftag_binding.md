---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_csrftag_binding

The `citrixadc_appfwprofile_csrftag_binding` data source allows you to retrieve information about a specific CSRF tag binding for an Application Firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_csrftag_binding" "example" {
  name              = "tf_appfwprofile"
  csrftag           = "www.source.com"
  csrfformactionurl = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_csrftag_binding.example.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_csrftag_binding.example.state
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `csrftag` - (Required) The web form originating URL.
* `csrfformactionurl` - (Required) The web form action URL.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the binding.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
