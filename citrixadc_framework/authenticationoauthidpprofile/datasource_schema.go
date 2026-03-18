package authenticationoauthidpprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationoauthidpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name-Value pairs of attributes to be inserted in idtoken. Configuration format is name=value_expr@@@name2=value2_expr@@@.\n'@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character.\nValue is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.",
			},
			"audience": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audience for which token is being sent by Citrix ADC IdP. This is typically entity name or url that represents the recipient",
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity of the relying party requesting for authentication.",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique secret string to authorize relying party at authorization server.",
			},
			"configservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the entity that is used to obtain configuration for the current authentication request. It is used only in Citrix Cloud.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This group will be part of AAA session's internal group list. This will be helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party. In authentication policy AAA.USER.IS_MEMBER_OF(\"<default_auth_group>\")  is way to use this feature.",
			},
			"encrypttoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to encrypt token when Citrix ADC IDP sends one.",
			},
			"issuer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new OAuth Identity Provider (IdP) single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL endpoint on relying party to which the OAuth token is to be sent.",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which Relying Party metadata is refreshed.",
			},
			"relyingpartymetadataurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the endpoint at which Citrix ADC IdP can get details about Relying Party (RP) being configured. Metadata response should include endpoints for jwks_uri for RP public key(s).",
			},
			"sendpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to send encrypted password in idtoken.",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to sign OpenID tokens.",
			},
			"signatureservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service in cloud used to sign the data. This is applicable only if signature if offloaded to cloud.",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option specifies the duration for which the token sent by Citrix ADC IdP is valid. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
		},
	}
}
