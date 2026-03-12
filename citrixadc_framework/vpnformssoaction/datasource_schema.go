package vpnformssoaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnformssoactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Root-relative URL to which the completed form is submitted.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the form based single sign-on profile.",
			},
			"namevaluepair": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Other name-value pair attributes to send to the server, in addition to sending the user name and password. Value names are separated by an ampersand (&), such as in name1=value1&name2=value2.",
			},
			"nvtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How to process the name-value pair. Available settings function as follows:\n* STATIC - The administrator-configured values are used.\n* DYNAMIC - The response is parsed, the form is extracted, and then submitted.",
			},
			"passwdfield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form field in which the user types in the password.",
			},
			"responsesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of bytes to allow in the response size. Specifies the number of bytes in the response to be parsed for extracting the forms.",
			},
			"ssosuccessrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that defines the criteria for SSO success. Expression such as checking for cookie in the response is a common example.",
			},
			"submitmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method (GET or POST) used by the single sign-on form to send the logon credentials to the logon server.",
			},
			"userfield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form field in which the user types in the user ID.",
			},
		},
	}
}
