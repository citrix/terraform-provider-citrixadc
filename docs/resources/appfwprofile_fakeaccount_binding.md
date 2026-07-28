---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_fakeaccount_binding

This resource is used to bind a fake-account detection rule to an application firewall profile.

~> **Note:** `formexpression` and `formurl_fad` are mutually exclusive — set at most one.


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

resource "citrixadc_appfwprofile_fakeaccount_binding" "tf_binding" {
  name             = citrixadc_appfwprofile.tf_appfwprofile.name
  fakeaccount      = "email"
  # formexpression and formurl_fad are mutually exclusive; set at most one of the two
  formexpression   = "^[a-z0-9._%+-]+@example\\.com$"
  tag              = "signup_form"
  isfieldnameregex = "NOTREGEX"
  state            = "ENABLED"
  comment          = "Detect fake account registrations on the signup form"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the fake-account rule is bound. Changing this value forces a new resource to be created.
* `fakeaccount` - (Required) Field name of the fake account rule. Changing this value forces a new resource to be created.
* `formexpression` - (Optional) A regular expression that defines the fake account. Mutually exclusive with `formurl_fad`; set at most one of the two. Changing this value forces a new resource to be created.
* `formurl_fad` - (Optional) The fake account detection URL. Mutually exclusive with `formexpression`; set at most one of the two. Changing this value forces a new resource to be created.
* `tag` - (Required) A tag expression that defines the fake account. Changing this value forces a new resource to be created.
* `isfieldnameregex` - (Optional) Whether the fake-account detection field name is a regular expression. Changing this value forces a new resource to be created. Possible values: [ REGEX, NOTREGEX ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the fake-account rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - (Optional) Indicates whether the rule was auto-deployed by a dynamic profile. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fakeaccount_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded). Only the populated arm of the mutually-exclusive `formexpression`/`formurl_fad` pair appears, so the format is `fakeaccount:<fakeaccount>,formexpression:<formexpression>,name:<name>,tag:<tag>` or `fakeaccount:<fakeaccount>,formurl_fad:<formurl_fad>,name:<name>,tag:<tag>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_fakeaccount_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_fakeaccount_binding.tf_binding "fakeaccount:email,formexpression:^[a-z0-9._%+-]+@example\\.com$,name:tf_appfwprofile,tag:signup_form"
```
