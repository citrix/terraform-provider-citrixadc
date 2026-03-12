package authenticationsamlaction

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

// AuthenticationsamlactionResourceModel describes the resource data model.
type AuthenticationsamlactionResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Artifactresolutionserviceurl   types.String `tfsdk:"artifactresolutionserviceurl"`
	Attribute1                     types.String `tfsdk:"attribute1"`
	Attribute10                    types.String `tfsdk:"attribute10"`
	Attribute11                    types.String `tfsdk:"attribute11"`
	Attribute12                    types.String `tfsdk:"attribute12"`
	Attribute13                    types.String `tfsdk:"attribute13"`
	Attribute14                    types.String `tfsdk:"attribute14"`
	Attribute15                    types.String `tfsdk:"attribute15"`
	Attribute16                    types.String `tfsdk:"attribute16"`
	Attribute2                     types.String `tfsdk:"attribute2"`
	Attribute3                     types.String `tfsdk:"attribute3"`
	Attribute4                     types.String `tfsdk:"attribute4"`
	Attribute5                     types.String `tfsdk:"attribute5"`
	Attribute6                     types.String `tfsdk:"attribute6"`
	Attribute7                     types.String `tfsdk:"attribute7"`
	Attribute8                     types.String `tfsdk:"attribute8"`
	Attribute9                     types.String `tfsdk:"attribute9"`
	Attributeconsumingserviceindex types.Int64  `tfsdk:"attributeconsumingserviceindex"`
	Attributes                     types.String `tfsdk:"attributes"`
	Audience                       types.String `tfsdk:"audience"`
	Authnctxclassref               types.List   `tfsdk:"authnctxclassref"`
	Customauthnctxclassref         types.String `tfsdk:"customauthnctxclassref"`
	Defaultauthenticationgroup     types.String `tfsdk:"defaultauthenticationgroup"`
	Digestmethod                   types.String `tfsdk:"digestmethod"`
	Enforceusername                types.String `tfsdk:"enforceusername"`
	Forceauthn                     types.String `tfsdk:"forceauthn"`
	Groupnamefield                 types.String `tfsdk:"groupnamefield"`
	Logoutbinding                  types.String `tfsdk:"logoutbinding"`
	Logouturl                      types.String `tfsdk:"logouturl"`
	Metadatarefreshinterval        types.Int64  `tfsdk:"metadatarefreshinterval"`
	Metadataurl                    types.String `tfsdk:"metadataurl"`
	Name                           types.String `tfsdk:"name"`
	Preferredbindtype              types.List   `tfsdk:"preferredbindtype"`
	Relaystaterule                 types.String `tfsdk:"relaystaterule"`
	Requestedauthncontext          types.String `tfsdk:"requestedauthncontext"`
	Samlacsindex                   types.Int64  `tfsdk:"samlacsindex"`
	Samlbinding                    types.String `tfsdk:"samlbinding"`
	Samlidpcertname                types.String `tfsdk:"samlidpcertname"`
	Samlissuername                 types.String `tfsdk:"samlissuername"`
	Samlredirecturl                types.String `tfsdk:"samlredirecturl"`
	Samlrejectunsignedassertion    types.String `tfsdk:"samlrejectunsignedassertion"`
	Samlsigningcertname            types.String `tfsdk:"samlsigningcertname"`
	Samltwofactor                  types.String `tfsdk:"samltwofactor"`
	Samluserfield                  types.String `tfsdk:"samluserfield"`
	Sendthumbprint                 types.String `tfsdk:"sendthumbprint"`
	Signaturealg                   types.String `tfsdk:"signaturealg"`
	Skewtime                       types.Int64  `tfsdk:"skewtime"`
	Statechecks                    types.String `tfsdk:"statechecks"`
	Storesamlresponse              types.String `tfsdk:"storesamlresponse"`
}

func (r *AuthenticationsamlactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationsamlaction resource.",
			},
			"artifactresolutionserviceurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the Artifact Resolution Service on IdP to which Citrix ADC will post artifact to get actual SAML token.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute1. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute10. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute11. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute12. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute13. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute14. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute15. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute16. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute2. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute3. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute4. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute5. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute6. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute7. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute8. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute9. Maximum length of the extracted attribute is 239 bytes.",
			},
			"attributeconsumingserviceindex": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "Index/ID of the attribute specification at Identity Provider (IdP). IdP will locate attributes requested by SP using this index and send those attributes in Assertion",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of attribute names separated by ',' which needs to be extracted.\nNote that preceeding and trailing spaces will be removed.\nAttribute name can be 127 bytes and total length of this string should not cross 2047 bytes.\nThese attributes have multi-value support separated by ',' and stored as key-value pair in AAA session",
			},
			"audience": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider",
			},
			"authnctxclassref": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the authentication class types that are requested from IdP (IdentityProvider).\nInternetProtocol: This is applicable when a principal is authenticated through the use of a provided IP address.\nInternetProtocolPassword: This is applicable when a principal is authenticated through the use of a provided IP address, in addition to a username/password.\nKerberos: This is applicable when the principal has authenticated using a password to a local authentication authority, in order to acquire a Kerberos ticket.\nMobileOneFactorUnregistered: This indicates authentication of the mobile device without requiring explicit end-user interaction.\nMobileTwoFactorUnregistered: This indicates two-factor based authentication during mobile customer registration process, such as secure device and user PIN.\nMobileOneFactorContract: Reflects mobile contract customer registration procedures and a single factor authentication.\nMobileTwoFactorContract: Reflects mobile contract customer registration procedures and a two-factor based authentication.\nPassword: This class is applicable when a principal authenticates using password over unprotected http session.\nPasswordProtectedTransport: This class is applicable when a principal authenticates to an authentication authority through the presentation of a password over a protected session.\nPreviousSession: This class is applicable when a principal had authenticated to an authentication authority at some point in the past using any authentication context.\nX509: This indicates that the principal authenticated by means of a digital signature where the key was validated as part of an X.509 Public Key Infrastructure.\nPGP: This indicates that the principal authenticated by means of a digital signature where the key was validated as part of a PGP Public Key Infrastructure.\nSPKI: This indicates that the principal authenticated by means of a digital signature where the key was validated via an SPKI Infrastructure.\nXMLDSig: This indicates that the principal authenticated by means of a digital signature according to the processing rules specified in the XML Digital Signature specification.\nSmartcard: This indicates that the principal has authenticated using smartcard.\nSmartcardPKI: This class is applicable when a principal authenticates to an authentication authority through a two-factor authentication mechanism using a smartcard with enclosed private key and a PIN.\nSoftwarePKI: This class is applicable when a principal uses an X.509 certificate stored in software to authenticate to the authentication authority.\nTelephony: This class is used to indicate that the principal authenticated via the provision of a fixed-line telephone number, transported via a telephony protocol such as ADSL.\nNomadTelephony: Indicates that the principal is \"roaming\" and authenticates via the means of the line number, a user suffix, and a password element.\nPersonalTelephony: This class is used to indicate that the principal authenticated via the provision of a fixed-line telephone.\nAuthenticatedTelephony: Indicates that the principal authenticated via the means of the line number, a user suffix, and a password element.\nSecureRemotePassword: This class is applicable when the authentication was performed by means of Secure Remote Password.\nTLSClient: This class indicates that the principal authenticated by means of a client certificate, secured with the SSL/TLS transport.\nTimeSyncToken: This is applicable when a principal authenticates through a time synchronization token.\nUnspecified: This indicates that the authentication was performed by unspecified means.\nWindows: This indicates that Windows integrated authentication is utilized for authentication.",
			},
			"customauthnctxclassref": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the custom authentication class reference to be sent as a part of the Authentication Request that is sent by the SP to SAML IDP. The input string must be the body of the authentication class being requested.\nInput format: Alphanumeric string or URL specifying the body of the Request.If more than one string has to be provided, then the same can be done by specifying the classes as a string of comma separated values.\nExample input: set authentication samlaction samlact1 -customAuthnCtxClassRef http://www.class1.com/LoA1,http://www.class2.com/LoA2",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"digestmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SHA256"),
				Description: "Algorithm to be used to compute/verify digest for SAML transactions",
			},
			"enforceusername": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Option to choose whether the username that is extracted from SAML assertion can be edited in login page while doing second factor",
			},
			"forceauthn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option that forces authentication at the Identity Provider (IdP) that receives Citrix ADC's request",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the tag in assertion that contains user groups.",
			},
			"logoutbinding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("POST"),
				Description: "This element specifies the transport mechanism of saml logout messages.",
			},
			"logouturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SingleLogout URL on IdP to which logoutRequest will be sent on Citrix ADC session cleanup.",
			},
			"metadatarefreshinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Interval in minutes for fetching metadata from specified metadata URL",
			},
			"metadataurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This URL is used for obtaining saml metadata. Note that it fills samlIdPCertName and samlredirectUrl fields so those fields should not be updated when metadataUrl present",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SAML server profile (action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after SAML profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"preferredbindtype": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the preferred binding types for sso and logout for metadata configuration.",
			},
			"relaystaterule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean expression that will be evaluated to validate the SAML Response.\nExamples:\nset authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.EQ(\"https://fqdn.com/\")'\nset authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.CONTAINS(\"https://fqdn.com/\")'\nset authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.CONTAINS_ANY(\"patset_name\")'\nset authentication samlAction samlsp -relaystateRule 'AAA.LOGIN.RELAYSTATE.REGEX_MATCH(re#http://<regex>.com/#)'.",
			},
			"requestedauthncontext": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("exact"),
				Description: "This element specifies the authentication context requirements of authentication statements returned in the response.",
			},
			"samlacsindex": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "Index/ID of the metadata entry corresponding to this configuration.",
			},
			"samlbinding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("POST"),
				Description: "This element specifies the transport mechanism of saml messages.",
			},
			"samlidpcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate used to verify responses from SAML Identity Provider (IdP). Note that if metadateURL is present then this filed should be empty.",
			},
			"samlissuername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.",
			},
			"samlredirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which users are redirected for authentication. Note that if metadateURL is present then this filed should be empty",
			},
			"samlrejectunsignedassertion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Reject unsigned SAML assertions. ON option results in rejection of Assertion that is received without signature. STRICT option ensures that both Response and Assertion are signed.",
			},
			"samlsigningcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate to sign requests from ServiceProvider (SP) to Identity Provider (IdP).",
			},
			"samltwofactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable second factor after SAML",
			},
			"samluserfield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SAML user ID, as given in the SAML assertion.",
			},
			"sendthumbprint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to send thumbprint instead of x509 certificate in SAML request",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RSA-SHA256"),
				Description: "Algorithm to be used to sign/verify SAML transactions",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "This option specifies the allowed clock skew in number of minutes that Citrix ADC ServiceProvider allows on an incoming assertion. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
			"statechecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean expression that will be evaluated to validate HTTP requests on SAML endpoints.\nExamples:\nset authentication samlaction <actionname> -stateChecks 'HTTP.REQ.HOSTNAME.EQ(\"https://fqdn.com/\")'",
			},
			"storesamlresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to store entire SAML Response through the life of user session.",
			},
		},
	}
}

func authenticationsamlactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationsamlactionResourceModel) authentication.Authenticationsamlaction {
	tflog.Debug(ctx, "In authenticationsamlactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationsamlaction := authentication.Authenticationsamlaction{}
	if !data.Artifactresolutionserviceurl.IsNull() {
		authenticationsamlaction.Artifactresolutionserviceurl = data.Artifactresolutionserviceurl.ValueString()
	}
	if !data.Attribute1.IsNull() {
		authenticationsamlaction.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		authenticationsamlaction.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute11.IsNull() {
		authenticationsamlaction.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute12.IsNull() {
		authenticationsamlaction.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute13.IsNull() {
		authenticationsamlaction.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute14.IsNull() {
		authenticationsamlaction.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute15.IsNull() {
		authenticationsamlaction.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute16.IsNull() {
		authenticationsamlaction.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute2.IsNull() {
		authenticationsamlaction.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute3.IsNull() {
		authenticationsamlaction.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute4.IsNull() {
		authenticationsamlaction.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute5.IsNull() {
		authenticationsamlaction.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute6.IsNull() {
		authenticationsamlaction.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute7.IsNull() {
		authenticationsamlaction.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute8.IsNull() {
		authenticationsamlaction.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute9.IsNull() {
		authenticationsamlaction.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Attributeconsumingserviceindex.IsNull() {
		authenticationsamlaction.Attributeconsumingserviceindex = utils.IntPtr(int(data.Attributeconsumingserviceindex.ValueInt64()))
	}
	if !data.Attributes.IsNull() {
		authenticationsamlaction.Attributes = data.Attributes.ValueString()
	}
	if !data.Audience.IsNull() {
		authenticationsamlaction.Audience = data.Audience.ValueString()
	}
	if !data.Customauthnctxclassref.IsNull() {
		authenticationsamlaction.Customauthnctxclassref = data.Customauthnctxclassref.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationsamlaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Digestmethod.IsNull() {
		authenticationsamlaction.Digestmethod = data.Digestmethod.ValueString()
	}
	if !data.Enforceusername.IsNull() {
		authenticationsamlaction.Enforceusername = data.Enforceusername.ValueString()
	}
	if !data.Forceauthn.IsNull() {
		authenticationsamlaction.Forceauthn = data.Forceauthn.ValueString()
	}
	if !data.Groupnamefield.IsNull() {
		authenticationsamlaction.Groupnamefield = data.Groupnamefield.ValueString()
	}
	if !data.Logoutbinding.IsNull() {
		authenticationsamlaction.Logoutbinding = data.Logoutbinding.ValueString()
	}
	if !data.Logouturl.IsNull() {
		authenticationsamlaction.Logouturl = data.Logouturl.ValueString()
	}
	if !data.Metadatarefreshinterval.IsNull() {
		authenticationsamlaction.Metadatarefreshinterval = utils.IntPtr(int(data.Metadatarefreshinterval.ValueInt64()))
	}
	if !data.Metadataurl.IsNull() {
		authenticationsamlaction.Metadataurl = data.Metadataurl.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationsamlaction.Name = data.Name.ValueString()
	}
	if !data.Relaystaterule.IsNull() {
		authenticationsamlaction.Relaystaterule = data.Relaystaterule.ValueString()
	}
	if !data.Requestedauthncontext.IsNull() {
		authenticationsamlaction.Requestedauthncontext = data.Requestedauthncontext.ValueString()
	}
	if !data.Samlacsindex.IsNull() {
		authenticationsamlaction.Samlacsindex = utils.IntPtr(int(data.Samlacsindex.ValueInt64()))
	}
	if !data.Samlbinding.IsNull() {
		authenticationsamlaction.Samlbinding = data.Samlbinding.ValueString()
	}
	if !data.Samlidpcertname.IsNull() {
		authenticationsamlaction.Samlidpcertname = data.Samlidpcertname.ValueString()
	}
	if !data.Samlissuername.IsNull() {
		authenticationsamlaction.Samlissuername = data.Samlissuername.ValueString()
	}
	if !data.Samlredirecturl.IsNull() {
		authenticationsamlaction.Samlredirecturl = data.Samlredirecturl.ValueString()
	}
	if !data.Samlrejectunsignedassertion.IsNull() {
		authenticationsamlaction.Samlrejectunsignedassertion = data.Samlrejectunsignedassertion.ValueString()
	}
	if !data.Samlsigningcertname.IsNull() {
		authenticationsamlaction.Samlsigningcertname = data.Samlsigningcertname.ValueString()
	}
	if !data.Samltwofactor.IsNull() {
		authenticationsamlaction.Samltwofactor = data.Samltwofactor.ValueString()
	}
	if !data.Samluserfield.IsNull() {
		authenticationsamlaction.Samluserfield = data.Samluserfield.ValueString()
	}
	if !data.Sendthumbprint.IsNull() {
		authenticationsamlaction.Sendthumbprint = data.Sendthumbprint.ValueString()
	}
	if !data.Signaturealg.IsNull() {
		authenticationsamlaction.Signaturealg = data.Signaturealg.ValueString()
	}
	if !data.Skewtime.IsNull() {
		authenticationsamlaction.Skewtime = utils.IntPtr(int(data.Skewtime.ValueInt64()))
	}
	if !data.Statechecks.IsNull() {
		authenticationsamlaction.Statechecks = data.Statechecks.ValueString()
	}
	if !data.Storesamlresponse.IsNull() {
		authenticationsamlaction.Storesamlresponse = data.Storesamlresponse.ValueString()
	}

	return authenticationsamlaction
}

func authenticationsamlactionSetAttrFromGet(ctx context.Context, data *AuthenticationsamlactionResourceModel, getResponseData map[string]interface{}) *AuthenticationsamlactionResourceModel {
	tflog.Debug(ctx, "In authenticationsamlactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["artifactresolutionserviceurl"]; ok && val != nil {
		data.Artifactresolutionserviceurl = types.StringValue(val.(string))
	} else {
		data.Artifactresolutionserviceurl = types.StringNull()
	}
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
	if val, ok := getResponseData["attributeconsumingserviceindex"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Attributeconsumingserviceindex = types.Int64Value(intVal)
		}
	} else {
		data.Attributeconsumingserviceindex = types.Int64Null()
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
	if val, ok := getResponseData["customauthnctxclassref"]; ok && val != nil {
		data.Customauthnctxclassref = types.StringValue(val.(string))
	} else {
		data.Customauthnctxclassref = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["digestmethod"]; ok && val != nil {
		data.Digestmethod = types.StringValue(val.(string))
	} else {
		data.Digestmethod = types.StringNull()
	}
	if val, ok := getResponseData["enforceusername"]; ok && val != nil {
		data.Enforceusername = types.StringValue(val.(string))
	} else {
		data.Enforceusername = types.StringNull()
	}
	if val, ok := getResponseData["forceauthn"]; ok && val != nil {
		data.Forceauthn = types.StringValue(val.(string))
	} else {
		data.Forceauthn = types.StringNull()
	}
	if val, ok := getResponseData["groupnamefield"]; ok && val != nil {
		data.Groupnamefield = types.StringValue(val.(string))
	} else {
		data.Groupnamefield = types.StringNull()
	}
	if val, ok := getResponseData["logoutbinding"]; ok && val != nil {
		data.Logoutbinding = types.StringValue(val.(string))
	} else {
		data.Logoutbinding = types.StringNull()
	}
	if val, ok := getResponseData["logouturl"]; ok && val != nil {
		data.Logouturl = types.StringValue(val.(string))
	} else {
		data.Logouturl = types.StringNull()
	}
	if val, ok := getResponseData["metadatarefreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metadatarefreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Metadatarefreshinterval = types.Int64Null()
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
	if val, ok := getResponseData["relaystaterule"]; ok && val != nil {
		data.Relaystaterule = types.StringValue(val.(string))
	} else {
		data.Relaystaterule = types.StringNull()
	}
	if val, ok := getResponseData["requestedauthncontext"]; ok && val != nil {
		data.Requestedauthncontext = types.StringValue(val.(string))
	} else {
		data.Requestedauthncontext = types.StringNull()
	}
	if val, ok := getResponseData["samlacsindex"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Samlacsindex = types.Int64Value(intVal)
		}
	} else {
		data.Samlacsindex = types.Int64Null()
	}
	if val, ok := getResponseData["samlbinding"]; ok && val != nil {
		data.Samlbinding = types.StringValue(val.(string))
	} else {
		data.Samlbinding = types.StringNull()
	}
	if val, ok := getResponseData["samlidpcertname"]; ok && val != nil {
		data.Samlidpcertname = types.StringValue(val.(string))
	} else {
		data.Samlidpcertname = types.StringNull()
	}
	if val, ok := getResponseData["samlissuername"]; ok && val != nil {
		data.Samlissuername = types.StringValue(val.(string))
	} else {
		data.Samlissuername = types.StringNull()
	}
	if val, ok := getResponseData["samlredirecturl"]; ok && val != nil {
		data.Samlredirecturl = types.StringValue(val.(string))
	} else {
		data.Samlredirecturl = types.StringNull()
	}
	if val, ok := getResponseData["samlrejectunsignedassertion"]; ok && val != nil {
		data.Samlrejectunsignedassertion = types.StringValue(val.(string))
	} else {
		data.Samlrejectunsignedassertion = types.StringNull()
	}
	if val, ok := getResponseData["samlsigningcertname"]; ok && val != nil {
		data.Samlsigningcertname = types.StringValue(val.(string))
	} else {
		data.Samlsigningcertname = types.StringNull()
	}
	if val, ok := getResponseData["samltwofactor"]; ok && val != nil {
		data.Samltwofactor = types.StringValue(val.(string))
	} else {
		data.Samltwofactor = types.StringNull()
	}
	if val, ok := getResponseData["samluserfield"]; ok && val != nil {
		data.Samluserfield = types.StringValue(val.(string))
	} else {
		data.Samluserfield = types.StringNull()
	}
	if val, ok := getResponseData["sendthumbprint"]; ok && val != nil {
		data.Sendthumbprint = types.StringValue(val.(string))
	} else {
		data.Sendthumbprint = types.StringNull()
	}
	if val, ok := getResponseData["signaturealg"]; ok && val != nil {
		data.Signaturealg = types.StringValue(val.(string))
	} else {
		data.Signaturealg = types.StringNull()
	}
	if val, ok := getResponseData["skewtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Skewtime = types.Int64Value(intVal)
		}
	} else {
		data.Skewtime = types.Int64Null()
	}
	if val, ok := getResponseData["statechecks"]; ok && val != nil {
		data.Statechecks = types.StringValue(val.(string))
	} else {
		data.Statechecks = types.StringNull()
	}
	if val, ok := getResponseData["storesamlresponse"]; ok && val != nil {
		data.Storesamlresponse = types.StringValue(val.(string))
	} else {
		data.Storesamlresponse = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
