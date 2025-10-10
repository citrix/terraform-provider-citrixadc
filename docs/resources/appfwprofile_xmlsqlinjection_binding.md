---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_xmlsqlinjection_binding

The appfwprofile_xmlsqlinjection_binding resource is used to bind xmlsqlinjection to appfwprofile.


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
resource "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding" {
  name                    = citrixadc_appfwprofile.tf_appfwprofile.name
  xmlsqlinjection         = "hello"
  as_scan_location_xmlsql = "ELEMENT"
  alertonly               = "ON"
  isautodeployed          = "AUTODEPLOYED"
  state                   = "ENABLED"
  comment                 = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlsqlinjection` - (Required) Exempt the specified URL from the XML SQL injection check.  An XML SQL injection exemption (relaxation) consists of the following items: * Name. Name to exempt, as a string or a PCRE-format regular expression. * ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string. * Location. ELEMENT if the injection is located in an XML element, ATTRIBUTE if located in an XML attribute.
* `alertonly` - (Optional) Send SNMP alert?
* `as_scan_location_xmlsql` - (Optional) Location of SQL injection exception - XML Element or Attribute.
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_xmlsql` - (Optional) Is the XML SQL Injection exempted field name a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlsqlinjection_binding. It is the concatenation of `name`, `xmlsqlinjection` and `as_scan_location_xmlsql` attributes separated by comma.


## Import

A appfwprofile_xmlsqlinjection_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding tf_appfwprofile,hello,ELEMENT
```
