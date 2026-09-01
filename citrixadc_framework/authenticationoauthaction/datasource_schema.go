package authenticationoauthaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationoauthactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationoauthactionResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits. Every non-key attribute
// is Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares.
type AuthenticationoauthactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Allowedalgorithms          types.List   `tfsdk:"allowedalgorithms"`
	Attribute1                 types.String `tfsdk:"attribute1"`
	Attribute10                types.String `tfsdk:"attribute10"`
	Attribute11                types.String `tfsdk:"attribute11"`
	Attribute12                types.String `tfsdk:"attribute12"`
	Attribute13                types.String `tfsdk:"attribute13"`
	Attribute14                types.String `tfsdk:"attribute14"`
	Attribute15                types.String `tfsdk:"attribute15"`
	Attribute16                types.String `tfsdk:"attribute16"`
	Attribute2                 types.String `tfsdk:"attribute2"`
	Attribute3                 types.String `tfsdk:"attribute3"`
	Attribute4                 types.String `tfsdk:"attribute4"`
	Attribute5                 types.String `tfsdk:"attribute5"`
	Attribute6                 types.String `tfsdk:"attribute6"`
	Attribute7                 types.String `tfsdk:"attribute7"`
	Attribute8                 types.String `tfsdk:"attribute8"`
	Attribute9                 types.String `tfsdk:"attribute9"`
	Attributes                 types.String `tfsdk:"attributes"`
	Audience                   types.String `tfsdk:"audience"`
	Authentication             types.String `tfsdk:"authentication"`
	Authorizationendpoint      types.String `tfsdk:"authorizationendpoint"`
	Certendpoint               types.String `tfsdk:"certendpoint"`
	Certfilepath               types.String `tfsdk:"certfilepath"`
	Clientid                   types.String `tfsdk:"clientid"`
	Clientsecret               types.String `tfsdk:"clientsecret"`
	ClientsecretWo             types.String `tfsdk:"clientsecret_wo"`
	ClientsecretWoVersion      types.Int64  `tfsdk:"clientsecret_wo_version"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Granttype                  types.String `tfsdk:"granttype"`
	Graphendpoint              types.String `tfsdk:"graphendpoint"`
	Idtokendecryptendpoint     types.String `tfsdk:"idtokendecryptendpoint"`
	Introspecturl              types.String `tfsdk:"introspecturl"`
	Intunedeviceidexpression   types.String `tfsdk:"intunedeviceidexpression"`
	Issuer                     types.String `tfsdk:"issuer"`
	Metadataurl                types.String `tfsdk:"metadataurl"`
	Name                       types.String `tfsdk:"name"`
	Oauthmiscflags             types.List   `tfsdk:"oauthmiscflags"`
	Oauthtype                  types.String `tfsdk:"oauthtype"`
	Pkce                       types.String `tfsdk:"pkce"`
	Refreshinterval            types.Int64  `tfsdk:"refreshinterval"`
	Requestattribute           types.String `tfsdk:"requestattribute"`
	Resourceuri                types.String `tfsdk:"resourceuri"`
	Scopes                     types.String `tfsdk:"scopes"`
	Skewtime                   types.Int64  `tfsdk:"skewtime"`
	Tenantid                   types.String `tfsdk:"tenantid"`
	Tokenendpoint              types.String `tfsdk:"tokenendpoint"`
	Tokenendpointauthmethod    types.String `tfsdk:"tokenendpointauthmethod"`
	Userinfourl                types.String `tfsdk:"userinfourl"`
	Usernamefield              types.String `tfsdk:"usernamefield"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationoauthaction.json). Never settable;
	// populated from GET.
	Oauthstatus types.String `tfsdk:"oauthstatus"`
}

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
				Sensitive:   true,
				Description: "Secret string established by user and authorization server",
			},
			"clientsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Secret string established by user and authorization server",
			},
			"clientsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a clientsecret_wo update.",
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
			"scopes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OAuth Scopes expected. Please specify scopes in space separated format as per RFC 6749 (OAuth 2.0). Each scope value can contain any printable ASCII character except double-quote (\") and backslash (\\). Maximum length is 1024.",
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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"oauthstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes status information of oauth server. Possible values = INIT, CERTFETCH, AADFORGRAPH, GRAPH, AADFORMDM, MDMINFO, COMPLETE",
			},
		},
	}
}

// authenticationoauthactionDataSourceSetAttrFromGet projects a NITRO
// authenticationoauthaction GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled from
// the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationoauthactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationoauthactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationoauthactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Allowedalgorithms = utils.MapGetStringList(g, "allowedalgorithms")
	data.Attribute1 = utils.MapGetString(g, "attribute1")
	data.Attribute10 = utils.MapGetString(g, "attribute10")
	data.Attribute11 = utils.MapGetString(g, "attribute11")
	data.Attribute12 = utils.MapGetString(g, "attribute12")
	data.Attribute13 = utils.MapGetString(g, "attribute13")
	data.Attribute14 = utils.MapGetString(g, "attribute14")
	data.Attribute15 = utils.MapGetString(g, "attribute15")
	data.Attribute16 = utils.MapGetString(g, "attribute16")
	data.Attribute2 = utils.MapGetString(g, "attribute2")
	data.Attribute3 = utils.MapGetString(g, "attribute3")
	data.Attribute4 = utils.MapGetString(g, "attribute4")
	data.Attribute5 = utils.MapGetString(g, "attribute5")
	data.Attribute6 = utils.MapGetString(g, "attribute6")
	data.Attribute7 = utils.MapGetString(g, "attribute7")
	data.Attribute8 = utils.MapGetString(g, "attribute8")
	data.Attribute9 = utils.MapGetString(g, "attribute9")
	data.Attributes = utils.MapGetString(g, "attributes")
	data.Audience = utils.MapGetString(g, "audience")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authorizationendpoint = utils.MapGetString(g, "authorizationendpoint")
	data.Certendpoint = utils.MapGetString(g, "certendpoint")
	data.Certfilepath = utils.MapGetString(g, "certfilepath")
	data.Clientid = utils.MapGetString(g, "clientid")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Granttype = utils.MapGetString(g, "granttype")
	data.Graphendpoint = utils.MapGetString(g, "graphendpoint")
	data.Idtokendecryptendpoint = utils.MapGetString(g, "idtokendecryptendpoint")
	data.Introspecturl = utils.MapGetString(g, "introspecturl")
	data.Intunedeviceidexpression = utils.MapGetString(g, "intunedeviceidexpression")
	data.Issuer = utils.MapGetString(g, "issuer")
	data.Metadataurl = utils.MapGetString(g, "metadataurl")
	data.Oauthmiscflags = utils.MapGetStringList(g, "oauthmiscflags")
	data.Oauthtype = utils.MapGetString(g, "oauthtype")
	data.Pkce = utils.MapGetString(g, "pkce")
	data.Refreshinterval = utils.MapGetInt64(g, "refreshinterval")
	data.Requestattribute = utils.MapGetString(g, "requestattribute")
	data.Resourceuri = utils.MapGetString(g, "resourceuri")
	data.Scopes = utils.MapGetString(g, "scopes")
	data.Skewtime = utils.MapGetInt64(g, "skewtime")
	data.Tenantid = utils.MapGetString(g, "tenantid")
	data.Tokenendpoint = utils.MapGetString(g, "tokenendpoint")
	data.Tokenendpointauthmethod = utils.MapGetString(g, "tokenendpointauthmethod")
	data.Userinfourl = utils.MapGetString(g, "userinfourl")
	data.Usernamefield = utils.MapGetString(g, "usernamefield")

	// clientsecret / clientsecret_wo(+version) are write-only or action-only
	// inputs the GET never returns -> Null.
	data.Clientsecret = types.StringNull()
	data.ClientsecretWo = types.StringNull()
	data.ClientsecretWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Oauthstatus = utils.MapGetString(g, "oauthstatus")
}
