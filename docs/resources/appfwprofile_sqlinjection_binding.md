---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_sqlinjection_binding

The `appfwprofile_sqlinjection_binding` resource is used to bind an HTML SQL injection relaxation rule to an Application Firewall profile.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "demo_appfw" {
  name                     = "demo_appfwprofile1"
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

resource "citrixadc_appfwprofile_sqlinjection_binding" "demo_binding" {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  sqlinjection         = "data"
  formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
  as_scan_location_sql = "FORMFIELD"
  as_value_type_sql    = "Keyword"
  as_value_expr_sql    = ".*"
  isautodeployed       = "NOTAUTODEPLOYED"
  isvalueregex_sql     = "REGEX"
  state                = "ENABLED"
  ruletype             = "DENY"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `sqlinjection` - (Required) The web form field name.
* `formactionurl_sql` - (Required) The web form action URL.
* `as_scan_location_sql` - (Required) Location of SQL injection exception - form field, header or cookie.
* `as_value_type_sql` - (Optional) The web form value type.
* `as_value_expr_sql` - (Optional) The web form value expression.
* `isregex_sql` - (Optional) Is the web form field name a regular expression?
* `isvalueregex_sql` - (Optional) Is the web form field value a regular expression?
* `state` - (Optional) Enabled.
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile?
* `alertonly` - (Optional) Send SNMP alert?
* `ruletype` - (Optional) Specifies rule type of binding.
* `resourceid` - (Optional) A "id" that identifies the rule.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_sqlinjection_binding`. It is the concatenation of the `name`, `sqlinjection`, `formactionurl_sql`, `as_scan_location_sql`, `as_value_type_sql`, `as_value_expr_sql` and `ruletype` attributes separated by a comma.


## Import

An `appfwprofile_sqlinjection_binding` can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_sqlinjection_binding.demo_binding demo_appfwprofile1,data,^https://citrix.csg.com/analytics/saw.dll$,FORMFIELD,Keyword,.*,DENY
```

An `appfwprofile_sqlinjection_binding` which does not have values set for `as_value_type_sql`, `as_value_expr_sql` and `ruletype` can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_sqlinjection_binding.demo_binding demo_appfwprofile1,data,^https://citrix.csg.com/analytics/saw.dll$,FORMFIELD
```
