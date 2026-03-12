package authenticationoauthaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AuthenticationoauthactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowedalgorithms": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Multivalued option to specify allowed token verification algorithms.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute1",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute10",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute11",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute12",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute13",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute14",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute15",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute16",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute2",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute3",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute4",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute5",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute6",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute7",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute8",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute to be extracted from OAuth Token and to be stored in the attribute9",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of attribute names separated by ',' which needs to be extracted.\nNote that preceding and trailing spaces will be removed.\nAttribute name can be 127 bytes and total length of this string should not cross 1023 bytes.\nThese attributes have multi-value support separated by ',' and stored as key-value pair in AAA session",
			},
			"audience": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audience for which token sent by Authorization server is applicable. This is typically entity name or url that represents the recipient",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If authentication is disabled, password is not sent in the request.",
			},
			"authorizationendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authorization endpoint/url to which unauthenticated user will be redirected. Citrix ADC redirects user to this endpoint by adding query parameters including clientid. If this parameter not specified then as default value we take Token Endpoint/URL value. Please note that Authorization Endpoint or Token Endpoint is mandatory for oauthAction",
			},
			"certendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the endpoint that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.",
			},
			"certfilepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the file that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.",
			},
			"clientid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identity of the client/user who is getting authenticated. Authorization server infers client configuration using this ID",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secret string established by user and authorization server",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"granttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Grant type support. value can be code or password",
			},
			"graphendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the Graph API service to learn Enterprise Mobility Services (EMS) endpoints.",
			},
			"idtokendecryptendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which obtained idtoken will be posted to get a decrypted user identity. Encrypted idtoken will be obtained by posting OAuth token to token endpoint. In order to decrypt idtoken, Citrix ADC posts request to the URL configured",
			},
			"introspecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which access token would be posted for validation",
			},
			"intunedeviceidexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The expression that will be evaluated to obtain IntuneDeviceId for compliance check against IntuneNAC device compliance endpoint. The expression is applicable when the OAuthType is INTUNE. The maximum length allowed to be used as IntuneDeviceId for the device compliance check from the computed response after the expression evaluation is 41.\nExamples:\nadd authentication oauthAction <actionName> -intuneDeviceIdExpression 'AAA.LOGIN.INTUNEURI.AFTER_STR(\"IntuneDeviceId://\")'",
			},
			"issuer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Identity of the server whose tokens are to be accepted.",
			},
			"metadataurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Well-known configuration endpoint of the Authorization Server. Citrix ADC fetches server details from this endpoint.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the OAuth Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"oauthmiscflags": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Option to set/unset miscellaneous feature flags.\nAvailable values function as follows:\n* Base64Encode_Authorization_With_Padding - On setting this value, for endpoints (token and introspect), basic authorization header will be base64 encoded with padding.\n* EnableJWTRequest - By enabling this field, Authorisation request to IDP will have jwt signed 'request' parameter",
			},
			"oauthtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the OAuth implementation. Default value is generic implementation that is applicable for most deployments.",
			},
			"pkce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable/disable PKCE flow during authentication.",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which services are monitored for necessary configuration.",
			},
			"requestattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name-Value pairs of attributes to be inserted in request parameter. Configuration format is name=value_expr@@@name2=value2_expr@@@.\n'@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character.\nValue is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.",
			},
			"resourceuri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Resource URL for Oauth configuration.",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option specifies the allowed clock skew in number of minutes that Citrix ADC allows on an incoming token. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
			"tenantid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.",
			},
			"tokenendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which OAuth token will be posted to verify its authenticity. User obtains this token from Authorization server upon successful authentication. Citrix ADC will validate presented token by posting it to the URL configured",
			},
			"tokenendpointauthmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to select the variant of token authentication method. This method is used while exchanging code with IdP.",
			},
			"userinfourl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which OAuth access token will be posted to obtain user information.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute in the token from which username should be extracted.",
			},
		},
	}
}
