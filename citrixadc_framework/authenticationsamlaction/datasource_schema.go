package authenticationsamlaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsamlactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationsamlactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (metadataimportstatus). Every non-key attribute is Computed.
type AuthenticationsamlactionDataSourceModel struct {
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

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationsamlaction.json). Never settable;
	// populated from GET.
	Metadataimportstatus types.String `tfsdk:"metadataimportstatus"`
}

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
				Description: "This element specifies the authentication class types that are requested from IdP (IdentityProvider).",
			},
			"customauthnctxclassref": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the custom authentication class reference to be sent as a part of the Authentication Request that is sent by the SP to SAML IDP.",
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
				Description: "Boolean expression that will be evaluated to validate the SAML Response.",
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
				Description: "Boolean expression that will be evaluated to validate HTTP requests on SAML endpoints.",
			},
			"storesamlresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to store entire SAML Response through the life of user session.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"metadataimportstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes metadata import status. Possible values: INIT, SUCCESS, FAIL.",
			},
		},
	}
}

// authenticationsamlactionDataSourceSetAttrFromGet projects a NITRO
// authenticationsamlaction GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationsamlactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationsamlactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationsamlactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Artifactresolutionserviceurl = utils.MapGetString(g, "artifactresolutionserviceurl")
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
	data.Attributeconsumingserviceindex = utils.MapGetInt64(g, "attributeconsumingserviceindex")
	data.Attributes = utils.MapGetString(g, "attributes")
	data.Audience = utils.MapGetString(g, "audience")
	data.Authnctxclassref = utils.MapGetStringList(g, "authnctxclassref")
	data.Customauthnctxclassref = utils.MapGetString(g, "customauthnctxclassref")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Digestmethod = utils.MapGetString(g, "digestmethod")
	data.Enforceusername = utils.MapGetString(g, "enforceusername")
	data.Forceauthn = utils.MapGetString(g, "forceauthn")
	data.Groupnamefield = utils.MapGetString(g, "groupnamefield")
	data.Logoutbinding = utils.MapGetString(g, "logoutbinding")
	data.Logouturl = utils.MapGetString(g, "logouturl")
	data.Metadatarefreshinterval = utils.MapGetInt64(g, "metadatarefreshinterval")
	data.Metadataurl = utils.MapGetString(g, "metadataurl")
	data.Preferredbindtype = utils.MapGetStringList(g, "preferredbindtype")
	data.Relaystaterule = utils.MapGetString(g, "relaystaterule")
	data.Requestedauthncontext = utils.MapGetString(g, "requestedauthncontext")
	data.Samlacsindex = utils.MapGetInt64(g, "samlacsindex")
	data.Samlbinding = utils.MapGetString(g, "samlbinding")
	data.Samlidpcertname = utils.MapGetString(g, "samlidpcertname")
	data.Samlissuername = utils.MapGetString(g, "samlissuername")
	data.Samlredirecturl = utils.MapGetString(g, "samlredirecturl")
	data.Samlrejectunsignedassertion = utils.MapGetString(g, "samlrejectunsignedassertion")
	data.Samlsigningcertname = utils.MapGetString(g, "samlsigningcertname")
	data.Samltwofactor = utils.MapGetString(g, "samltwofactor")
	data.Samluserfield = utils.MapGetString(g, "samluserfield")
	data.Sendthumbprint = utils.MapGetString(g, "sendthumbprint")
	data.Signaturealg = utils.MapGetString(g, "signaturealg")
	data.Skewtime = utils.MapGetInt64(g, "skewtime")
	data.Statechecks = utils.MapGetString(g, "statechecks")
	data.Storesamlresponse = utils.MapGetString(g, "storesamlresponse")

	// Read-only metadata.
	data.Metadataimportstatus = utils.MapGetString(g, "metadataimportstatus")
}
