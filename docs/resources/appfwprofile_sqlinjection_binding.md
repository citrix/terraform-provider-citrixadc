---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_sqlinjection_binding

The `appfwprofile_sqlinjection_binding` resource is used to add bindings between Applicatin Firewall Profile and HTML SQLInjection relaxation rule.

## Example usage

``` hcl
resource citrixadc_appfwprofile_sqlinjection_binding demo_binding {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  sqlinjection         = "data"
  isautodeployed       = "NOTAUTODEPLOYED"
  as_scan_location_sql = "FORMFIELD"
  formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
  as_value_type_sql    = "Keyword"
  isvalueregex_sql     = "REGEX"
  as_value_expr_sql    = ".*"
  state                = "ENABLED"
  depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

resource citrixadc_appfwprofile demo_appfw {
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
```

## Argument Reference

* `name` - Name of the profile to which to bind an exemption or rule.
* `sqlinjection` - The web form field name.
* `formactionurl_sql` - The web form action URL.
* `as_scan_location_sql` - Location of SQL injection exception - form field, header or cookie. Possible values: [ FORMFIELD, HEADER, COOKIE ]
* `isregex_sql` - (Optional) Is the web form field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]
* `as_value_type_sql` - (Optional) The web form value type. Possible values: [ Keyword, SpecialString, Wildchar ]
* `as_value_expr_sql` - (Optional) The web form value expression.
* `isvalueregex_sql` - (Optional) Is the web form field value a regular expression?. Possible values: [ REGEX, NOTREGEX ]
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ on, off ]
* `ruletype` - (Optional) Specifies rule type of binding.
* `resourceid` - (Optional) A "id" that identifies the rule.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_sqlinjection_binding`. It is the concatenation of the `name`, `sqlinjection`, `formactionurl_sql`, `as_scan_location_sql`, `as_value_type_sql` and `as_value_expr_sql` attributes separated by a comma.


## Import

An `appfwprofile_sqlinjection_binding` can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_sqlinjection_binding.demo_binding demo_appfw,data,^https://citrix.csg.com/analytics/saw.dll$,FORMFIELD,Keyword,.*
```

An `appfwprofile_sqlinjection_binding` which does not have values set for `as_value_type_sql` and `as_value_expr_sql` can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_sqlinjection_binding.demo_binding demo_appfw,data,^https://citrix.csg.com/analytics/saw.dll$,FORMFIELD
```
