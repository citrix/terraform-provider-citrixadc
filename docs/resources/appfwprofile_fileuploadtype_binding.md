---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_fileuploadtype_binding

The appfwprofile_fileuploadtype_binding resource is used to bind fileuploadtype to appfwprofile.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}
resource "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding" {
  name                   = citrixadc_appfwprofile.tf_appfwprofile.name
  fileuploadtype         = "tf_uploadtype"
  as_fileuploadtypes_url = "www.example.com"
  filetype               = ["pdf", "text"]
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `fileuploadtype` - (Required) FileUploadTypes to allow/deny.
* `as_fileuploadtypes_url` - (Required) FileUploadTypes action URL.
* `filetype` - (Required) FileUploadTypes file types.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_fileuploadtypes_url` - (Optional) Is a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `isnameregex` - (Optional) Is field name a regular expression?


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fileuploadtype_binding. It is the concatenation of `name`, `fileuploadtype`, `as_fileuploadtypes_url`, and `filetype` attributes separated by comma. The `filetype` is a space-separated string of all file types (e.g., `pdf text`).


## Import

A appfwprofile_fileuploadtype_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_fileuploadtype_binding.tf_binding tf_appfwprofile,tf_uploadtype,www.example.com,pdf%20text
```
