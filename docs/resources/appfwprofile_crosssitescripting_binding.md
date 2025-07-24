---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_crosssitescripting_binding

The `appfwprofile_crosssitescripting_binding` resource is used to add binding between Applicatin Firewall Profile and CrossSiteScripting relaxation rule.

## Example usage

``` hcl
resource citrixadc_appfwprofile demo_appfw {
  name                     = "demo_appfwprofile"
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
resource citrixadc_appfwprofile_crosssitescripting_binding demo_binding {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  crosssitescripting   = "demoxss"
  formactionurl_xss    = "http://www.example.com"
  as_scan_location_xss = "HEADER"
  isregex_xss          = "NOTREGEX"
  isvalueregex_xss     = "NOTREGEX"
  as_value_type_xss    = "Attribute"
  as_value_expr_xss    = "nArtificialDelayMS"
  comment              = "democomment"
  state                = "ENABLED"
}
```

## Argument Reference

* `name` - Name of the profile to which to bind an exemption or rule.
* `crosssitescripting` - The web form field name.
* `isregex_xss` - Is the web form field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]
* `formactionurl_xss` - The web form action URL.
* `as_scan_location_xss` - Location of cross-site scripting exception - form field, header, cookie or URL. Possible values: [ FORMFIELD, HEADER, COOKIE, URL ]
* `as_value_type_xss` - (Optional) The web form value type. Possible values: [ Tag, Attribute, Pattern ]
* `as_value_expr_xss` - (Optional) The web form value expression.
* `isvalueregex_xss` - (Optional) Is the web form field value a regular expression?. Possible values: [ REGEX, NOTREGEX ]
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ on, off ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_crosssitescripting_binding`.It is the concatenation of the `name`, `crosssitescripting`, `formactionurl_xss`, `as_scan_location_xss`, `as_value_type_xss` and `as_value_expr_xss` attributes separated by a comma.


## Import

A appfwprofile_crosssitescripting_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_crosssitescripting_binding.demo_binding demo_appfw,demoxss,http://www.example.com,HEADER,Attribute,nArtificialDelayMS
```

A appfwprofile_crosssitescripting_binding which does not have values set for as_value_type_xss and as_value_expr_xss can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_crosssitescripting_binding.demo_binding demo_appfw,demoxss,http://www.example.com,HEADER
```