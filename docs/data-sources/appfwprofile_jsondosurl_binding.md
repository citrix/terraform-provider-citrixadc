---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_jsondosurl_binding

The appfwprofile_jsondosurl_binding data source allows you to retrieve information about appfwprofile jsondosurl bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_jsondosurl_binding" "tf_binding" {
  name       = "tf_appfwprofile"
  jsondosurl = ".*"
}

output "state" {
  value = data.citrixadc_appfwprofile_jsondosurl_binding.tf_binding.state
}

output "jsonmaxdocumentlength" {
  value = data.citrixadc_appfwprofile_jsondosurl_binding.tf_binding.jsonmaxdocumentlength
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsondosurl` - (Required) The URL on which we need to enforce the specified JSON denial-of-service (JSONDoS) attack protections. An JSON DoS configuration consists of the following items: URL (PCRE-format regular expression), Maximum-document-length-check toggle, Maximum document length, Maximum-container-depth-check toggle, Maximum container depth, Maximum-object-key-count-check toggle, Maximum object key count, Maximum-object-key-length-check toggle, Maximum object key length, Maximum-array-value-count-check toggle, Maximum array value count, Maximum-string-length-check toggle, Maximum string length.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsondosurl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `jsonmaxarraylength` - Maximum array length in the any of JSON object. This check protects against arrays having large lengths.
* `jsonmaxarraylengthcheck` - State if JSON Max array value count check is ON or OFF.
* `jsonmaxcontainerdepth` - Maximum allowed nesting depth of JSON document. JSON allows one to nest the containers (object and array) in any order to any depth. This check protects against documents that have excessive depth of hierarchy.
* `jsonmaxcontainerdepthcheck` - State if JSON Max depth check is ON or OFF.
* `jsonmaxdocumentlength` - Maximum document length of JSON document, in bytes.
* `jsonmaxdocumentlengthcheck` - State if JSON Max document length check is ON or OFF.
* `jsonmaxobjectkeycount` - Maximum key count in the any of JSON object. This check protects against objects that have large number of keys.
* `jsonmaxobjectkeycountcheck` - State if JSON Max object key count check is ON or OFF.
* `jsonmaxobjectkeylength` - Maximum key length in the any of JSON object. This check protects against objects that have large keys.
* `jsonmaxobjectkeylengthcheck` - State if JSON Max object key length check is ON or OFF.
* `jsonmaxstringlength` - Maximum string length in the JSON. This check protects against strings that have large length.
* `jsonmaxstringlengthcheck` - State if JSON Max string value count check is ON or OFF.
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
