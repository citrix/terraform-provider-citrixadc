---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_bypasslist_binding

This resource is used to bind a bypass-list (relaxation) rule to an application firewall profile.


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

resource "citrixadc_appfwprofile_bypasslist_binding" "tf_binding" {
  name                      = citrixadc_appfwprofile.tf_appfwprofile.name
  as_bypass_list            = "X-Forwarded-For"
  as_bypass_list_location   = "HEADER"
  as_bypass_list_value_type = "Keyword"
  as_bypass_list_action     = "log"
  state                     = "ENABLED"
  comment                   = "Bypass checks for the X-Forwarded-For header"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the bypass-list rule is bound. Changing this value forces a new resource to be created.
* `as_bypass_list` - (Required) The bypass-list value (the target pattern to be exempted from security checks). Changing this value forces a new resource to be created.
* `as_bypass_list_location` - (Required) The scan location to which the bypass-list rule applies. Changing this value forces a new resource to be created.
* `as_bypass_list_value_type` - (Required) The value type of the bypass-list entry. Changing this value forces a new resource to be created.
* `as_bypass_list_action` - (Optional) The action to take when the bypass-list rule matches. Changing this value forces a new resource to be created. Possible values: [ none, log ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the bypass-list rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - (Optional) Indicates whether the rule was auto-deployed by a dynamic profile. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_bypasslist_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `as_bypass_list:<as_bypass_list>,as_bypass_list_location:<as_bypass_list_location>,as_bypass_list_value_type:<as_bypass_list_value_type>,name:<name>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_bypasslist_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_bypasslist_binding.tf_binding as_bypass_list:X-Forwarded-For,as_bypass_list_location:HEADER,as_bypass_list_value_type:Keyword,name:tf_appfwprofile
```
