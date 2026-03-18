package authenticationoauthaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationoauthactionResourceModel describes the resource data model.
type AuthenticationoauthactionResourceModel struct {
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
	Skewtime                   types.Int64  `tfsdk:"skewtime"`
	Tenantid                   types.String `tfsdk:"tenantid"`
	Tokenendpoint              types.String `tfsdk:"tokenendpoint"`
	Tokenendpointauthmethod    types.String `tfsdk:"tokenendpointauthmethod"`
	Userinfourl                types.String `tfsdk:"userinfourl"`
	Usernamefield              types.String `tfsdk:"usernamefield"`
}

func (r *AuthenticationoauthactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationoauthaction resource.",
			},
			"allowedalgorithms": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
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
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("CODE"),
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
				Default:     stringdefault.StaticString("GENERIC"),
				Description: "Type of the OAuth implementation. Default value is generic implementation that is applicable for most deployments.",
			},
			"pkce": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Option to enable/disable PKCE flow during authentication.",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1440),
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
				Default:     int64default.StaticInt64(5),
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
				Default:     stringdefault.StaticString("client_secret_post"),
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

func authenticationoauthactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationoauthactionResourceModel) authentication.Authenticationoauthaction {
	tflog.Debug(ctx, "In authenticationoauthactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationoauthaction := authentication.Authenticationoauthaction{}
	if !data.Attribute1.IsNull() {
		authenticationoauthaction.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		authenticationoauthaction.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute11.IsNull() {
		authenticationoauthaction.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute12.IsNull() {
		authenticationoauthaction.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute13.IsNull() {
		authenticationoauthaction.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute14.IsNull() {
		authenticationoauthaction.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute15.IsNull() {
		authenticationoauthaction.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute16.IsNull() {
		authenticationoauthaction.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute2.IsNull() {
		authenticationoauthaction.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute3.IsNull() {
		authenticationoauthaction.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute4.IsNull() {
		authenticationoauthaction.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute5.IsNull() {
		authenticationoauthaction.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute6.IsNull() {
		authenticationoauthaction.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute7.IsNull() {
		authenticationoauthaction.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute8.IsNull() {
		authenticationoauthaction.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute9.IsNull() {
		authenticationoauthaction.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Attributes.IsNull() {
		authenticationoauthaction.Attributes = data.Attributes.ValueString()
	}
	if !data.Audience.IsNull() {
		authenticationoauthaction.Audience = data.Audience.ValueString()
	}
	if !data.Authentication.IsNull() {
		authenticationoauthaction.Authentication = data.Authentication.ValueString()
	}
	if !data.Authorizationendpoint.IsNull() {
		authenticationoauthaction.Authorizationendpoint = data.Authorizationendpoint.ValueString()
	}
	if !data.Certendpoint.IsNull() {
		authenticationoauthaction.Certendpoint = data.Certendpoint.ValueString()
	}
	if !data.Certfilepath.IsNull() {
		authenticationoauthaction.Certfilepath = data.Certfilepath.ValueString()
	}
	if !data.Clientid.IsNull() {
		authenticationoauthaction.Clientid = data.Clientid.ValueString()
	}
	if !data.Clientsecret.IsNull() {
		authenticationoauthaction.Clientsecret = data.Clientsecret.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationoauthaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Granttype.IsNull() {
		authenticationoauthaction.Granttype = data.Granttype.ValueString()
	}
	if !data.Graphendpoint.IsNull() {
		authenticationoauthaction.Graphendpoint = data.Graphendpoint.ValueString()
	}
	if !data.Idtokendecryptendpoint.IsNull() {
		authenticationoauthaction.Idtokendecryptendpoint = data.Idtokendecryptendpoint.ValueString()
	}
	if !data.Introspecturl.IsNull() {
		authenticationoauthaction.Introspecturl = data.Introspecturl.ValueString()
	}
	if !data.Intunedeviceidexpression.IsNull() {
		authenticationoauthaction.Intunedeviceidexpression = data.Intunedeviceidexpression.ValueString()
	}
	if !data.Issuer.IsNull() {
		authenticationoauthaction.Issuer = data.Issuer.ValueString()
	}
	if !data.Metadataurl.IsNull() {
		authenticationoauthaction.Metadataurl = data.Metadataurl.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationoauthaction.Name = data.Name.ValueString()
	}
	if !data.Oauthtype.IsNull() {
		authenticationoauthaction.Oauthtype = data.Oauthtype.ValueString()
	}
	if !data.Pkce.IsNull() {
		authenticationoauthaction.Pkce = data.Pkce.ValueString()
	}
	if !data.Refreshinterval.IsNull() {
		authenticationoauthaction.Refreshinterval = utils.IntPtr(int(data.Refreshinterval.ValueInt64()))
	}
	if !data.Requestattribute.IsNull() {
		authenticationoauthaction.Requestattribute = data.Requestattribute.ValueString()
	}
	if !data.Resourceuri.IsNull() {
		authenticationoauthaction.Resourceuri = data.Resourceuri.ValueString()
	}
	if !data.Skewtime.IsNull() {
		authenticationoauthaction.Skewtime = utils.IntPtr(int(data.Skewtime.ValueInt64()))
	}
	if !data.Tenantid.IsNull() {
		authenticationoauthaction.Tenantid = data.Tenantid.ValueString()
	}
	if !data.Tokenendpoint.IsNull() {
		authenticationoauthaction.Tokenendpoint = data.Tokenendpoint.ValueString()
	}
	if !data.Tokenendpointauthmethod.IsNull() {
		authenticationoauthaction.Tokenendpointauthmethod = data.Tokenendpointauthmethod.ValueString()
	}
	if !data.Userinfourl.IsNull() {
		authenticationoauthaction.Userinfourl = data.Userinfourl.ValueString()
	}
	if !data.Usernamefield.IsNull() {
		authenticationoauthaction.Usernamefield = data.Usernamefield.ValueString()
	}

	return authenticationoauthaction
}

func authenticationoauthactionSetAttrFromGet(ctx context.Context, data *AuthenticationoauthactionResourceModel, getResponseData map[string]interface{}) *AuthenticationoauthactionResourceModel {
	tflog.Debug(ctx, "In authenticationoauthactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["attribute1"]; ok && val != nil {
		data.Attribute1 = types.StringValue(val.(string))
	} else {
		data.Attribute1 = types.StringNull()
	}
	if val, ok := getResponseData["attribute10"]; ok && val != nil {
		data.Attribute10 = types.StringValue(val.(string))
	} else {
		data.Attribute10 = types.StringNull()
	}
	if val, ok := getResponseData["attribute11"]; ok && val != nil {
		data.Attribute11 = types.StringValue(val.(string))
	} else {
		data.Attribute11 = types.StringNull()
	}
	if val, ok := getResponseData["attribute12"]; ok && val != nil {
		data.Attribute12 = types.StringValue(val.(string))
	} else {
		data.Attribute12 = types.StringNull()
	}
	if val, ok := getResponseData["attribute13"]; ok && val != nil {
		data.Attribute13 = types.StringValue(val.(string))
	} else {
		data.Attribute13 = types.StringNull()
	}
	if val, ok := getResponseData["attribute14"]; ok && val != nil {
		data.Attribute14 = types.StringValue(val.(string))
	} else {
		data.Attribute14 = types.StringNull()
	}
	if val, ok := getResponseData["attribute15"]; ok && val != nil {
		data.Attribute15 = types.StringValue(val.(string))
	} else {
		data.Attribute15 = types.StringNull()
	}
	if val, ok := getResponseData["attribute16"]; ok && val != nil {
		data.Attribute16 = types.StringValue(val.(string))
	} else {
		data.Attribute16 = types.StringNull()
	}
	if val, ok := getResponseData["attribute2"]; ok && val != nil {
		data.Attribute2 = types.StringValue(val.(string))
	} else {
		data.Attribute2 = types.StringNull()
	}
	if val, ok := getResponseData["attribute3"]; ok && val != nil {
		data.Attribute3 = types.StringValue(val.(string))
	} else {
		data.Attribute3 = types.StringNull()
	}
	if val, ok := getResponseData["attribute4"]; ok && val != nil {
		data.Attribute4 = types.StringValue(val.(string))
	} else {
		data.Attribute4 = types.StringNull()
	}
	if val, ok := getResponseData["attribute5"]; ok && val != nil {
		data.Attribute5 = types.StringValue(val.(string))
	} else {
		data.Attribute5 = types.StringNull()
	}
	if val, ok := getResponseData["attribute6"]; ok && val != nil {
		data.Attribute6 = types.StringValue(val.(string))
	} else {
		data.Attribute6 = types.StringNull()
	}
	if val, ok := getResponseData["attribute7"]; ok && val != nil {
		data.Attribute7 = types.StringValue(val.(string))
	} else {
		data.Attribute7 = types.StringNull()
	}
	if val, ok := getResponseData["attribute8"]; ok && val != nil {
		data.Attribute8 = types.StringValue(val.(string))
	} else {
		data.Attribute8 = types.StringNull()
	}
	if val, ok := getResponseData["attribute9"]; ok && val != nil {
		data.Attribute9 = types.StringValue(val.(string))
	} else {
		data.Attribute9 = types.StringNull()
	}
	if val, ok := getResponseData["attributes"]; ok && val != nil {
		data.Attributes = types.StringValue(val.(string))
	} else {
		data.Attributes = types.StringNull()
	}
	if val, ok := getResponseData["audience"]; ok && val != nil {
		data.Audience = types.StringValue(val.(string))
	} else {
		data.Audience = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authorizationendpoint"]; ok && val != nil {
		data.Authorizationendpoint = types.StringValue(val.(string))
	} else {
		data.Authorizationendpoint = types.StringNull()
	}
	if val, ok := getResponseData["certendpoint"]; ok && val != nil {
		data.Certendpoint = types.StringValue(val.(string))
	} else {
		data.Certendpoint = types.StringNull()
	}
	if val, ok := getResponseData["certfilepath"]; ok && val != nil {
		data.Certfilepath = types.StringValue(val.(string))
	} else {
		data.Certfilepath = types.StringNull()
	}
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	if val, ok := getResponseData["clientsecret"]; ok && val != nil {
		data.Clientsecret = types.StringValue(val.(string))
	} else {
		data.Clientsecret = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["granttype"]; ok && val != nil {
		data.Granttype = types.StringValue(val.(string))
	} else {
		data.Granttype = types.StringNull()
	}
	if val, ok := getResponseData["graphendpoint"]; ok && val != nil {
		data.Graphendpoint = types.StringValue(val.(string))
	} else {
		data.Graphendpoint = types.StringNull()
	}
	if val, ok := getResponseData["idtokendecryptendpoint"]; ok && val != nil {
		data.Idtokendecryptendpoint = types.StringValue(val.(string))
	} else {
		data.Idtokendecryptendpoint = types.StringNull()
	}
	if val, ok := getResponseData["introspecturl"]; ok && val != nil {
		data.Introspecturl = types.StringValue(val.(string))
	} else {
		data.Introspecturl = types.StringNull()
	}
	if val, ok := getResponseData["intunedeviceidexpression"]; ok && val != nil {
		data.Intunedeviceidexpression = types.StringValue(val.(string))
	} else {
		data.Intunedeviceidexpression = types.StringNull()
	}
	if val, ok := getResponseData["issuer"]; ok && val != nil {
		data.Issuer = types.StringValue(val.(string))
	} else {
		data.Issuer = types.StringNull()
	}
	if val, ok := getResponseData["metadataurl"]; ok && val != nil {
		data.Metadataurl = types.StringValue(val.(string))
	} else {
		data.Metadataurl = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["oauthtype"]; ok && val != nil {
		data.Oauthtype = types.StringValue(val.(string))
	} else {
		data.Oauthtype = types.StringNull()
	}
	if val, ok := getResponseData["pkce"]; ok && val != nil {
		data.Pkce = types.StringValue(val.(string))
	} else {
		data.Pkce = types.StringNull()
	}
	if val, ok := getResponseData["refreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Refreshinterval = types.Int64Null()
	}
	if val, ok := getResponseData["requestattribute"]; ok && val != nil {
		data.Requestattribute = types.StringValue(val.(string))
	} else {
		data.Requestattribute = types.StringNull()
	}
	if val, ok := getResponseData["resourceuri"]; ok && val != nil {
		data.Resourceuri = types.StringValue(val.(string))
	} else {
		data.Resourceuri = types.StringNull()
	}
	if val, ok := getResponseData["skewtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Skewtime = types.Int64Value(intVal)
		}
	} else {
		data.Skewtime = types.Int64Null()
	}
	if val, ok := getResponseData["tenantid"]; ok && val != nil {
		data.Tenantid = types.StringValue(val.(string))
	} else {
		data.Tenantid = types.StringNull()
	}
	if val, ok := getResponseData["tokenendpoint"]; ok && val != nil {
		data.Tokenendpoint = types.StringValue(val.(string))
	} else {
		data.Tokenendpoint = types.StringNull()
	}
	if val, ok := getResponseData["tokenendpointauthmethod"]; ok && val != nil {
		data.Tokenendpointauthmethod = types.StringValue(val.(string))
	} else {
		data.Tokenendpointauthmethod = types.StringNull()
	}
	if val, ok := getResponseData["userinfourl"]; ok && val != nil {
		data.Userinfourl = types.StringValue(val.(string))
	} else {
		data.Userinfourl = types.StringNull()
	}
	if val, ok := getResponseData["usernamefield"]; ok && val != nil {
		data.Usernamefield = types.StringValue(val.(string))
	} else {
		data.Usernamefield = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
