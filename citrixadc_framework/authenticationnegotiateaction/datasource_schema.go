package authenticationnegotiateaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationnegotiateactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the service principal that represnts Citrix ADC.",
			},
			"domainuser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.",
			},
			"domainuserpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password of the account that is mapped to the Citrix ADC principal.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the AD KDC server profile (negotiate action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KDC server profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"ntlmpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.",
			},
			"ou": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Active Directory organizational units (OU) attribute.",
			},
		},
	}
}
