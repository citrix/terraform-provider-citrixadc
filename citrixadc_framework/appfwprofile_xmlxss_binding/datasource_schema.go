package appfwprofile_xmlxss_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmlxssBindingDataSourceSchema() schema.Schema {
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
			"as_scan_location_xmlxss": schema.StringAttribute{
				Required:    true,
				Description: "Location of XSS injection exception - XML Element or Attribute.",
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
			"isregex_xmlxss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the XML XSS exempted field name a regular expression?",
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
			"xmlxss": schema.StringAttribute{
				Required:    true,
				Description: "Exempt the specified URL from the XML cross-site scripting (XSS) check.\nAn XML cross-site scripting exemption (relaxation) consists of the following items:\n* URL. URL to exempt, as a string or a PCRE-format regular expression.\n* ISREGEX flag. REGEX if URL is a regular expression, NOTREGEX if URL is a fixed string.\n* Location. ELEMENT if the attachment is located in an XML element, ATTRIBUTE if located in an XML attribute.",
			},
		},
	}
}
