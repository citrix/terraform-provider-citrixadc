package appfwprofile_xmldosurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmldosurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
			"xmlblockdtd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML DTD is ON or OFF. Protects against recursive Document Type Declaration (DTD) entity expansion attacks. Also, SOAP messages cannot have DTDs in messages.",
			},
			"xmlblockexternalentities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Block External Entities Check is ON or OFF. Protects against XML External Entity (XXE) attacks that force applications to parse untrusted external entities (sources) in XML documents.",
			},
			"xmlblockpi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Block PI is ON or OFF. Protects resources from denial of service attacks as SOAP messages cannot have processing instructions (PI) in messages.",
			},
			"xmldosurl": schema.StringAttribute{
				Required:    true,
				Description: "XML DoS URL regular expression length.",
			},
			"xmlmaxattributenamelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest name of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributenamelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attribute name length check is ON or OFF.",
			},
			"xmlmaxattributes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum number of attributes per XML element. Protects against overflow attacks.",
			},
			"xmlmaxattributescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max attributes check is ON or OFF.",
			},
			"xmlmaxattributevaluelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest value of any XML attribute. Protects against overflow attacks.",
			},
			"xmlmaxattributevaluelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max atribute value length is ON or OFF.",
			},
			"xmlmaxchardatalength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum size of CDATA. Protects against overflow attacks and large quantities of unparsed data within XML messages.",
			},
			"xmlmaxchardatalengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max CDATA length check is ON or OFF.",
			},
			"xmlmaxelementchildren": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of children allowed per XML element. Protects against overflow attacks.",
			},
			"xmlmaxelementchildrencheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element children check is ON or OFF.",
			},
			"xmlmaxelementdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum nesting (depth) of XML elements. This check protects against documents that have excessive hierarchy depths.",
			},
			"xmlmaxelementdepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element depth check is ON or OFF.",
			},
			"xmlmaxelementnamelength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest name of any element (including the expanded namespace) to protect against overflow attacks.",
			},
			"xmlmaxelementnamelengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max element name length check is ON or OFF.",
			},
			"xmlmaxelements": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of XML elements allowed. Protects against overflow attacks.",
			},
			"xmlmaxelementscheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max elements check is ON or OFF.",
			},
			"xmlmaxentityexpansiondepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum entity expansion depth. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansiondepthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max Entity Expansions Depth Check is ON or OFF.",
			},
			"xmlmaxentityexpansions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum allowed number of entity expansions. Protects aganist Entity Expansion Attack.",
			},
			"xmlmaxentityexpansionscheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max Entity Expansions Check is ON or OFF.",
			},
			"xmlmaxfilesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum size of XML messages. Protects against overflow attacks.",
			},
			"xmlmaxfilesizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max file size check is ON or OFF.",
			},
			"xmlmaxnamespaces": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify maximum number of active namespaces. Protects against overflow attacks.",
			},
			"xmlmaxnamespacescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max namespaces check is ON or OFF.",
			},
			"xmlmaxnamespaceurilength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the longest URI of any XML namespace. Protects against overflow attacks.",
			},
			"xmlmaxnamespaceurilengthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max namespace URI length check is ON or OFF.",
			},
			"xmlmaxnodes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum number of XML nodes. Protects against overflow attacks.",
			},
			"xmlmaxnodescheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Max nodes check is ON or OFF.",
			},
			"xmlmaxsoaparrayrank": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Max Individual SOAP Array Rank. This is the dimension of the SOAP array.",
			},
			"xmlmaxsoaparraysize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "XML Max Total SOAP Array Size. Protects against SOAP Array Abuse attack.",
			},
			"xmlminfilesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Enforces minimum message size.",
			},
			"xmlminfilesizecheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML Min file size check is ON or OFF.",
			},
			"xmlsoaparraycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State if XML SOAP Array check is ON or OFF.",
			},
		},
	}
}
