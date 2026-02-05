package tmformssoaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TmformssoactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which the completed form is submitted.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new form-based single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"namevaluepair": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name-value pair attributes to send to the server in addition to sending the username and password. Value names are separated by an ampersand (&) (for example, name1=value1&name2=value2).",
			},
			"nvtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of processing of the name-value pair. If you specify STATIC, the values configured by the administrator are used. For DYNAMIC, the response is parsed, and the form is extracted and then submitted.",
			},
			"passwdfield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form field in which the user types in the password.",
			},
			"responsesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bytes, in the response, to parse for extracting the forms.",
			},
			"ssosuccessrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, that checks to see if single sign-on is successful.",
			},
			"submitmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method used by the single sign-on form to send the logon credentials to the logon server. Applies only to STATIC name-value type.",
			},
			"userfield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the form field in which the user types in the user ID.",
			},
		},
	}
}
