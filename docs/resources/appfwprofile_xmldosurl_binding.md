---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_xmldosurl_binding

The appfwprofile_xmldosurl_binding resource is used to bind xmldosurl to appfwprofile resource.


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
resource "citrixadc_appfwprofile_xmldosurl_binding" "tf_binding" {
  name                           = citrixadc_appfwprofile.tf_appfwprofile.name
  xmldosurl                      = ".*"
  state                          = "ENABLED"
  xmlsoaparraycheck              = "ON"
  xmlmaxelementdepthcheck        = "ON"
  xmlmaxfilesize                 = 100000
  xmlmaxfilesizecheck            = "OFF"
  xmlmaxnamespaceurilength       = 200
  xmlmaxnamespaceurilengthcheck  = "ON"
  xmlmaxelementnamelength        = 300
  xmlmaxelementnamelengthcheck   = "ON"
  xmlmaxelements                 = 30
  xmlmaxelementscheck            = "ON"
  xmlmaxattributes               = 20
  xmlmaxattributescheck          = "ON"
  xmlmaxchardatalength           = 1000
  xmlmaxchardatalengthcheck      = "ON"
  xmlmaxnamespaces               = 30
  xmlmaxnamespacescheck          = "ON"
  xmlmaxattributenamelength      = 200
  xmlmaxattributenamelengthcheck = "ON"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmldosurl` - (Required) XML DoS URL regular expression length.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `xmlblockdtd` - (Optional) State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages.
* `xmlblockexternalentities` - (Optional) State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.
* `xmlblockpi` - (Optional) State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.
* `xmlmaxattributenamelength` - (Optional) Specify the longest name of any XML attribute. Protects against overflow attacks.
* `xmlmaxattributenamelengthcheck` - (Optional) State if XML Max attribute name length check is ON or OFF.
* `xmlmaxattributes` - (Optional) Specify maximum number of attributes per XML element. Protects against overflow attacks.
* `xmlmaxattributescheck` - (Optional) State if XML Max attributes check is ON or OFF.
* `xmlmaxattributevaluelength` - (Optional) Specify the longest value of any XML attribute. Protects against overflow attacks.
* `xmlmaxattributevaluelengthcheck` - (Optional) State if XML Max atribute value length is ON or OFF.
* `xmlmaxchardatalength` - (Optional) Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.
* `xmlmaxchardatalengthcheck` - (Optional) State if XML Max CDATA length check is ON or OFF.
* `xmlmaxelementchildren` - (Optional) Specify the maximum number of children allowed per XML element. Protects against overflow attacks.
* `xmlmaxelementchildrencheck` - (Optional) State if XML Max element children check is ON or OFF.
* `xmlmaxelementdepth` - (Optional) Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.
* `xmlmaxelementdepthcheck` - (Optional) State if XML Max element depth check is ON or OFF.
* `xmlmaxelementnamelength` - (Optional) Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.
* `xmlmaxelementnamelengthcheck` - (Optional) State if XML Max element name length check is ON or OFF.
* `xmlmaxelements` - (Optional) Specify the maximum number of XML elements allowed. Protects against overflow attacks.
* `xmlmaxelementscheck` - (Optional) State if XML Max elements check is ON or OFF.
* `xmlmaxentityexpansiondepth` - (Optional) Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.
* `xmlmaxentityexpansiondepthcheck` - (Optional) State if XML Max Entity Expansions Depth Check is ON or OFF.
* `xmlmaxentityexpansions` - (Optional) Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.
* `xmlmaxentityexpansionscheck` - (Optional) State if XML Max Entity Expansions Check is ON or OFF.
* `xmlmaxfilesize` - (Optional) Specify the maximum size of XML messages. Protects against overflow attacks.
* `xmlmaxfilesizecheck` - (Optional) State if XML Max file size check is ON or OFF.
* `xmlmaxnamespaces` - (Optional) Specify maximum number of active namespaces. Protects against overflow attacks.
* `xmlmaxnamespacescheck` - (Optional) State if XML Max namespaces check is ON or OFF.
* `xmlmaxnamespaceurilength` - (Optional) Specify the longest URI of any XML namespace. Protects against overflow attacks.
* `xmlmaxnamespaceurilengthcheck` - (Optional) State if XML Max namespace URI length check is ON or OFF.
* `xmlmaxnodes` - (Optional) Specify the maximum number of XML nodes. Protects against overflow attacks.
* `xmlmaxnodescheck` - (Optional) State if XML Max nodes check is ON or OFF.
* `xmlmaxsoaparrayrank` - (Optional) XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.
* `xmlmaxsoaparraysize` - (Optional) XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.
* `xmlminfilesize` - (Optional) Enforces minimum message size.
* `xmlminfilesizecheck` - (Optional) State if XML Min file size check is ON or OFF.
* `xmlsoaparraycheck` - (Optional) State if XML SOAP Array check is ON or OFF.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmldosurl_binding. It is the concatenation of `name` and `xmldosurl` attributes separated by comma.


## Import

A appfwprofile_xmldosurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_xmldosurl_binding.tf_binding tf_appfwprofile,.*
```
