---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_xmlvalidationurl_binding

The appfwprofile_xmlvalidationurl_binding resource is used to bind xmlvalidationurl to appfwprofile resource.


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
resource "citrixadc_appfwprofile_xmlvalidationurl_binding" "tf_binding" {
  name             = citrixadc_appfwprofile.tf_appfwprofile.name
  xmlvalidationurl = ".*"
  isautodeployed   = "AUTODEPLOYED"
  comment          = "Testing"
  alertonly        = "ON"
  state            = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlvalidationurl` - (Required) XML Validation URL regular expression.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `xmladditionalsoapheaders` - (Optional) Allow addtional soap headers.
* `xmlendpointcheck` - (Optional) Modifies the behaviour of the Request URL validation w.r.t. the Service URL. 	If set to ABSOLUTE, the entire request URL is validated with the entire URL mentioned in Service of the associated WSDL. 		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would FAIL the validation. 	If set to RELAIVE, only the non-hostname part of the request URL is validated against the non-hostname part of the Service URL. 		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would PASS the validation.
* `xmlrequestschema` - (Optional) XML Schema object for request validation .
* `xmlresponseschema` - (Optional) XML Schema object for response validation.
* `xmlvalidateresponse` - (Optional) Validate response message.
* `xmlvalidatesoapenvelope` - (Optional) Validate SOAP Evelope only.
* `xmlwsdl` - (Optional) WSDL object for soap request validation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlvalidationurl_binding. It is the concatenation of `name` and `xmlvalidationurl` attributes separated by comma.


## Import

A appfwprofile_xmlvalidationurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_xmlvalidationurl_binding.tf_binding tf_appfwprofile,.*
```
