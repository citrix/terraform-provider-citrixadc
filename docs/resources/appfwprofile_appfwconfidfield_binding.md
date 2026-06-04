---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_appfwconfidfield_binding

Designates a form field as confidential within an application firewall profile on the Citrix ADC. Confidential fields (for example, password or credit-card input fields on a web form) are masked in the application firewall logs and traces, preventing sensitive user-submitted values from being recorded. Create this binding to attach a confidential-field rule to an existing `appfwprofile`.


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

resource "citrixadc_appfwprofile_appfwconfidfield_binding" "tf_binding" {
  name            = citrixadc_appfwprofile.tf_appfwprofile.name
  confidfield     = "password"
  cffield_url     = "https://www.example.com/login"
  isregex_cffield = "NOTREGEX"
  state           = "ENABLED"
  comment         = "Mask the login password field"
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the confidential field is bound. Changing this value forces a new resource to be created.
* `confidfield` - (Required) Name of the form field to designate as confidential. Changing this value forces a new resource to be created.
* `cffield_url` - (Optional) URL of the web page that contains the web form. Changing this value forces a new resource to be created.
* `isregex_cffield` - (Optional) Whether the confidential field name is a regular expression. Changing this value forces a new resource to be created. Possible values: [ REGEX, NOTREGEX ]
* `comment` - (Optional) Any comments about the purpose of the profile, or other useful information about the profile. Changing this value forces a new resource to be created.
* `state` - (Optional) Enable or disable the confidential-field rule. Changing this value forces a new resource to be created. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_appfwconfidfield_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `cffield_url:<cffield_url>,confidfield:<confidfield>,name:<name>`.
* `resourceid` - (Read-only) A system-generated identifier that identifies the rule.
* `isautodeployed` - (Read-only) Indicates whether the rule was auto-deployed by a dynamic profile.
* `alertonly` - (Read-only) Indicates whether an SNMP alert is sent.


## Import

A appfwprofile_appfwconfidfield_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_appfwprofile_appfwconfidfield_binding.tf_binding cffield_url:https://www.example.com/login,confidfield:password,name:tf_appfwprofile
```
