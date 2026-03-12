package appfwprofile_xmlsqlinjection_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmlsqlinjectionBindingDataSourceSchema() schema.Schema {
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
			"as_scan_location_xmlsql": schema.StringAttribute{
				Required:    true,
				Description: "Location of SQL injection exception - XML Element or Attribute.",
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
			"isregex_xmlsql": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the XML SQL Injection exempted field name a regular expression?",
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
			"xmlsqlinjection": schema.StringAttribute{
				Required:    true,
				Description: "Exempt the specified URL from the XML SQL injection check.\nAn XML SQL injection exemption (relaxation) consists of the following items:\n* Name. Name to exempt, as a string or a PCRE-format regular expression.\n* ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string.\n* Location. ELEMENT if the injection is located in an XML element, ATTRIBUTE if located in an XML attribute.",
			},
		},
	}
}
