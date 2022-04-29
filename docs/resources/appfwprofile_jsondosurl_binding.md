---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_jsondosurl_binding

The appfwprofile_jsondosurl_binding resource is used to bind jsondosurl to appfw profile resource.


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
resource "citrixadc_appfwprofile_jsondosurl_binding" "tf_binding" {
  name                        = citrixadc_appfwprofile.tf_appfwprofile.name
  jsondosurl                  = ".*"
  state                       = "ENABLED"
  alertonly                   = "ON"
  isautodeployed              = "AUTODEPLOYED"
  jsonmaxarraylengthcheck     = "ON"
  jsonmaxdocumentlengthcheck  = "ON"
  jsonmaxcontainerdepth       = 5
  jsonmaxobjectkeylengthcheck = "OFF"
  jsonmaxarraylength          = 100000
  jsonmaxdocumentlength       = 200000
  jsonmaxobjectkeycountcheck  = "ON"
  jsonmaxobjectkeylength      = 128
  jsonmaxobjectkeycount       = 1000
  jsonmaxstringlengthcheck    = "ON"
  jsonmaxcontainerdepthcheck  = "ON"
  jsonmaxstringlength         = 1000
  comment                     = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsondosurl` - (Required) The URL on which we need to enforce the specified JSON denial-of-service (JSONDoS) attack protections. An JSON DoS configuration consists of the following items: * URL. PCRE-format regular expression for the URL. * Maximum-document-length-check toggle.  ON to enable this check, OFF to disable it. * Maximum document length. Positive integer representing the maximum length of the JSON document. * Maximum-container-depth-check toggle. ON to enable, OFF to disable.  * Maximum container depth. Positive integer representing the maximum container depth of the JSON document. * Maximum-object-key-count-check toggle. ON to enable, OFF to disable. * Maximum object key count. Positive integer representing the maximum allowed number of keys in any of the  JSON object. * Maximum-object-key-length-check toggle. ON to enable, OFF to disable. * Maximum object key length. Positive integer representing the maximum allowed length of key in any of the  JSON object. * Maximum-array-value-count-check toggle. ON to enable, OFF to disable. * Maximum array value count. Positive integer representing the maximum allowed number of values in any of the JSON array. * Maximum-string-length-check toggle. ON to enable, OFF to disable. * Maximum string length. Positive integer representing the maximum length of string in JSON.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `jsonmaxarraylength` - (Optional) Maximum array length in the any of JSON object. This check protects against arrays having large lengths.
* `jsonmaxarraylengthcheck` - (Optional) State if JSON Max array value count check is ON or OFF.
* `jsonmaxcontainerdepth` - (Optional) Maximum allowed nesting depth  of JSON document. JSON allows one to nest the containers (object and array) in any order to any depth. This check protects against documents that have excessive depth of hierarchy.
* `jsonmaxcontainerdepthcheck` - (Optional) State if JSON Max depth check is ON or OFF.
* `jsonmaxdocumentlength` - (Optional) Maximum document length of JSON document, in bytes.
* `jsonmaxdocumentlengthcheck` - (Optional) State if JSON Max document length check is ON or OFF.
* `jsonmaxobjectkeycount` - (Optional) Maximum key count in the any of JSON object. This check protects against objects that have large number of keys.
* `jsonmaxobjectkeycountcheck` - (Optional) State if JSON Max object key count check is ON or OFF.
* `jsonmaxobjectkeylength` - (Optional) Maximum key length in the any of JSON object. This check protects against objects that have large keys.
* `jsonmaxobjectkeylengthcheck` - (Optional) State if JSON Max object key length check is ON or OFF.
* `jsonmaxstringlength` - (Optional) Maximum string length in the JSON. This check protects against strings that have large length.
* `jsonmaxstringlengthcheck` - (Optional) State if JSON Max string value count check is ON or OFF.
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsondosurl_binding. It is the concatenation of  `name` and `jsondosurl` attributes separated by comma.


## Import

A appfwprofile_jsondosurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_jsondosurl_binding.tf_binding tf_appfwprofile,.*
```
