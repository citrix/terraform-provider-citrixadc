---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_jsonxssurl_binding

The appfwprofile_jsonxssurl_binding resource is used to bind jsonxssurl to appfwprofile resource.


## Example usage

```hclresource "citrixadc_appfwprofile" "tf_appfwprofile" {
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
resource "citrixadc_appfwprofile_jsonxssurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.tf_appfwprofile.name
  jsonxssurl     = "www.example.com"
  alertonly      = "OFF"
  state          = "ENABLED"
  isautodeployed = "NOTAUTODEPLOYED"
  comment        = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsonxssurl` - (Required) A regular expression that designates a URL on the Json XSS URL list for which XSS violations are relaxed. Enclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsonxssurl_binding. It is the concatenation of `name` and `jsonxssurl` attributes separated by comma.


## Import

A appfwprofile_jsonxssurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_jsonxssurl_binding.tf_binding tf_appfwprofile,www.example.com
```
