---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_fileuploadtype_binding

The appfwprofile_fileuploadtype_binding data source allows you to retrieve information about a specific fileuploadtype binding to an appfwprofile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding" {
  name                   = "tf_appfwprofile"
  fileuploadtype         = "tf_uploadtype"
  as_fileuploadtypes_url = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
  filetype               = ["pdf", "text"]
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_fileuploadtype_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_fileuploadtype_binding.tf_binding.state
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which the fileuploadtype is bound.
* `fileuploadtype` - (Required) FileUploadTypes to allow/deny.
* `as_fileuploadtypes_url` - (Required) FileUploadTypes action URL.
* `filetype` - (Required) FileUploadTypes file types.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_fileuploadtype_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isnameregex` - Is field name a regular expression?
* `isregex_fileuploadtypes_url` - Is a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
