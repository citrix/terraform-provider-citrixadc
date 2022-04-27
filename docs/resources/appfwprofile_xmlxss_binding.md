---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_xmlxss_binding

The appfwprofile_xmlxss_binding resource is used to bind xmlxss with appfw profile.


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
resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding" {
  name                    = citrixadc_appfwprofile.tf_appfwprofile.name
  xmlxss                  = "tf_xmlxss"
  as_scan_location_xmlxss = "ELEMENT"
  state                   = "ENABLED"
  alertonly               = "ON"
  isregex_xmlxss          = "NOTREGEX"
  isautodeployed          = "AUTODEPLOYED"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlxss` - (Required) Exempt the specified URL from the XML cross-site scripting (XSS) check. An XML cross-site scripting exemption (relaxation) consists of the following items: * URL. URL to exempt, as a string or a PCRE-format regular expression. * ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string. * Location. ELEMENT if the attachment is located in an XML element, ATTRIBUTE if located in an XML attribute.
* `as_scan_location_xmlxss` - (Required) Location of XSS injection exception - XML Element or Attribute.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_xmlxss` - (Optional) Is the XML XSS exempted field name a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlxss_binding. It is the concatenation `name` ,`xmlxss` and `as_scan_location_xmlxss` attributes separated by comma.


## Import

A appfwprofile_xmlxss_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_xmlxss_binding.tf_binding tf_appfwprofile,tf_xmlxss,ELEMENT
```
