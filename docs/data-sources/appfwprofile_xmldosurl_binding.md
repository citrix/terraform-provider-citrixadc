---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_xmldosurl_binding

The appfwprofile_xmldosurl_binding data source allows you to retrieve information about appfwprofile xmldosurl bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_xmldosurl_binding" "tf_binding" {
  name      = "tf_appfwprofile"
  xmldosurl = ".*"
}

output "state" {
  value = data.citrixadc_appfwprofile_xmldosurl_binding.tf_binding.state
}

output "xmlmaxfilesize" {
  value = data.citrixadc_appfwprofile_xmldosurl_binding.tf_binding.xmlmaxfilesize
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmldosurl` - (Required) XML DoS URL regular expression length.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmldosurl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
* `xmlblockdtd` - State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages.
* `xmlblockexternalentities` - State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.
* `xmlblockpi` - State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.
* `xmlmaxattributenamelength` - Specify the longest name of any XML attribute. Protects against overflow attacks.
* `xmlmaxattributenamelengthcheck` - State if XML Max attribute name length check is ON or OFF.
* `xmlmaxattributes` - Specify maximum number of attributes per XML element. Protects against overflow attacks.
* `xmlmaxattributescheck` - State if XML Max attributes check is ON or OFF.
* `xmlmaxattributevaluelength` - Specify the longest value of any XML attribute. Protects against overflow attacks.
* `xmlmaxattributevaluelengthcheck` - State if XML Max atribute value length is ON or OFF.
* `xmlmaxchardatalength` - Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.
* `xmlmaxchardatalengthcheck` - State if XML Max CDATA length check is ON or OFF.
* `xmlmaxelementchildren` - Specify the maximum number of children allowed per XML element. Protects against overflow attacks.
* `xmlmaxelementchildrencheck` - State if XML Max element children check is ON or OFF.
* `xmlmaxelementdepth` - Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.
* `xmlmaxelementdepthcheck` - State if XML Max element depth check is ON or OFF.
* `xmlmaxelementnamelength` - Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.
* `xmlmaxelementnamelengthcheck` - State if XML Max element name length check is ON or OFF.
* `xmlmaxelements` - Specify the maximum number of XML elements allowed. Protects against overflow attacks.
* `xmlmaxelementscheck` - State if XML Max elements check is ON or OFF.
* `xmlmaxentityexpansiondepth` - Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.
* `xmlmaxentityexpansiondepthcheck` - State if XML Max Entity Expansions Depth Check is ON or OFF.
* `xmlmaxentityexpansions` - Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.
* `xmlmaxentityexpansionscheck` - State if XML Max Entity Expansions Check is ON or OFF.
* `xmlmaxfilesize` - Specify the maximum size of XML messages. Protects against overflow attacks.
* `xmlmaxfilesizecheck` - State if XML Max file size check is ON or OFF.
* `xmlmaxnamespaces` - Specify maximum number of active namespaces. Protects against overflow attacks.
* `xmlmaxnamespacescheck` - State if XML Max namespaces check is ON or OFF.
* `xmlmaxnamespaceurilength` - Specify the longest URI of any XML namespace. Protects against overflow attacks.
* `xmlmaxnamespaceurilengthcheck` - State if XML Max namespace URI length check is ON or OFF.
* `xmlmaxnodes` - Specify the maximum number of XML nodes. Protects against overflow attacks.
* `xmlmaxnodescheck` - State if XML Max nodes check is ON or OFF.
* `xmlmaxsoaparrayrank` - XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.
* `xmlmaxsoaparraysize` - XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.
* `xmlminfilesize` - Enforces minimum message size.
* `xmlminfilesizecheck` - State if XML Min file size check is ON or OFF.
* `xmlsoaparraycheck` - State if XML SOAP Array check is ON or OFF.
