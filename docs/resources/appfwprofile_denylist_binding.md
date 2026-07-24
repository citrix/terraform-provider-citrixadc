---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_denylist_binding

This resource is used to bind a deny-list rule to an application firewall profile.


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

resource "citrixadc_appfwprofile_denylist_binding" "tf_binding" {
  name                    = citrixadc_appfwprofile.tf_appfwprofile.name
  as_deny_list            = "X-Forwarded-For"
  as_deny_list_location   = "HEADER"
  as_deny_list_value_type = "Keyword"
  as_deny_list_action     = ["RESET", "log"]
  state                   = "ENABLED"
  comment                 = "Block requests matching the X-Forwarded-For header"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the deny-list rule is bound. Changing this value forces a new resource to be created.
* `as_deny_list` - (Required) The deny-list value (the target pattern to be blocked). Changing this value forces a new resource to be created.
* `as_deny_list_location` - (Required) The scan location to which the deny-list rule applies. Changing this value forces a new resource to be created.
* `as_deny_list_value_type` - (Required) The value type of the deny-list entry. Changing this value forces a new resource to be created.
* `as_deny_list_action` - (Optional) The action(s) to take when the deny-list rule matches, expressed as a list of strings. Changing this value forces a new resource to be created. Defaults to `["REDIRECT"]`. Possible values: [ none, log, RESET, REDIRECT ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the deny-list rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - (Optional) Indicates whether the rule was auto-deployed by a dynamic profile. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_denylist_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `as_deny_list:<as_deny_list>,as_deny_list_location:<as_deny_list_location>,as_deny_list_value_type:<as_deny_list_value_type>,name:<name>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_denylist_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_denylist_binding.tf_binding as_deny_list:X-Forwarded-For,as_deny_list_location:HEADER,as_deny_list_value_type:Keyword,name:tf_appfwprofile
```
