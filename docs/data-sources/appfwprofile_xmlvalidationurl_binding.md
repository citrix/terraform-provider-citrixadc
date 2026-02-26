---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmlvalidationurl_binding

The appfwprofile_xmlvalidationurl_binding data source allows you to retrieve information about an XML validation URL binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmlvalidationurl_binding" "tf_binding" {
  name              = "tf_appfwprofile"
  xmlvalidationurl  = ".*"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_xmlvalidationurl_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_xmlvalidationurl_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_xmlvalidationurl_binding.tf_binding.comment
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlvalidationurl` - (Required) XML Validation URL regular expression.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_xmlvalidationurl_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
* `xmladditionalsoapheaders` - Allow addtional soap headers.
* `xmlendpointcheck` - Modifies the behaviour of the Request URL validation w.r.t. the Service URL. If set to ABSOLUTE, the entire request URL is validated with the entire URL mentioned in Service of the associated WSDL. eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would FAIL the validation. If set to RELAIVE, only the non-hostname part of the request URL is validated against the non-hostname part of the Service URL. eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would PASS the validation.
* `xmlrequestschema` - XML Schema object for request validation.
* `xmlresponseschema` - XML Schema object for response validation.
* `xmlvalidateresponse` - Validate response message.
* `xmlvalidatesoapenvelope` - Validate SOAP Evelope only.
* `xmlwsdl` - WSDL object for soap request validation.
