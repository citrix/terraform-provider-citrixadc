---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_blockkeyword_binding

Binds a block keyword to an application firewall profile on the Citrix ADC. A block keyword identifies a forbidden word or pattern (literal or PCRE) that, when found in a specified form field, causes the application firewall to block or flag the request. Create this binding to attach a deny-keyword rule to an existing `appfwprofile` for a given form field and form action URL.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
}

resource "citrixadc_appfwprofile_blockkeyword_binding" "tf_binding" {
  name                              = citrixadc_appfwprofile.tf_appfwprofile.name
  blockkeyword                      = "drop table"
  fieldname                         = "comments"
  as_blockkeyword_formurl           = "https://www.example.com/submit"
  as_fieldname_isregex_blockkeyword = "NOTREGEX"
  blockkeywordtype                  = "literal"
  state                             = "ENABLED"
  comment                           = "Block SQL drop attempts in the comments field"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the block keyword is bound. Changing this value forces a new resource to be created.
* `blockkeyword` - (Required) The block keyword (word or pattern) to deny within the specified form field. Changing this value forces a new resource to be created.
* `fieldname` - (Required) Name of the form field in which the block keyword is evaluated. Changing this value forces a new resource to be created.
* `as_blockkeyword_formurl` - (Required) The form action URL on which the block keyword rule applies. Changing this value forces a new resource to be created.
* `as_fieldname_isregex_blockkeyword` - (Optional) Whether the block keyword field name is a regular expression. Changing this value forces a new resource to be created. Defaults to `"NOTREGEX"`. Possible values: [ REGEX, NOTREGEX ]
* `blockkeywordtype` - (Optional) The type of the block keyword. Changing this value forces a new resource to be created. Possible values: [ literal, PCRE ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the block keyword rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - (Optional) Indicates whether the rule was auto-deployed by a dynamic profile. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_blockkeyword_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `as_blockkeyword_formurl:<as_blockkeyword_formurl>,blockkeyword:<blockkeyword>,fieldname:<fieldname>,name:<name>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_blockkeyword_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_blockkeyword_binding.tf_binding as_blockkeyword_formurl:https://www.example.com/submit,blockkeyword:drop table,fieldname:comments,name:tf_appfwprofile
```
