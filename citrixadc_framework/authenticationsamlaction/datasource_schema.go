package authenticationsamlaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AuthenticationsamlactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Computed:    true,
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
				Computed:    true,
				Description: "Algorithm to be used to compute/verify digest for SAML transactions",
			},
			"enforceusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "This element specifies the transport mechanism of saml logout messages.",
			},
			"logouturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SingleLogout URL on IdP to which logoutRequest will be sent on Citrix ADC session cleanup.",
			},
			"metadatarefreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "This element specifies the authentication context requirements of authentication statements returned in the response.",
			},
			"samlacsindex": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Index/ID of the metadata entry corresponding to this configuration.",
			},
			"samlbinding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
				Computed:    true,
				Description: "Algorithm to be used to sign/verify SAML transactions",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
