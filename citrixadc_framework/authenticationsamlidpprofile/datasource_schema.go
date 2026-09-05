package authenticationsamlidpprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsamlidpprofileDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationsamlidpprofileResourceModel. Every non-key
// attribute is Computed, and it additionally exposes read-only (GET-only)
// attributes the resource deliberately omits (metadataimportstatus).
type AuthenticationsamlidpprofileDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Acsurlrule                  types.String `tfsdk:"acsurlrule"`
	Assertionconsumerserviceurl types.String `tfsdk:"assertionconsumerserviceurl"`
	Attribute1                  types.String `tfsdk:"attribute1"`
	Attribute10                 types.String `tfsdk:"attribute10"`
	Attribute10expr             types.String `tfsdk:"attribute10expr"`
	Attribute10format           types.String `tfsdk:"attribute10format"`
	Attribute10friendlyname     types.String `tfsdk:"attribute10friendlyname"`
	Attribute11                 types.String `tfsdk:"attribute11"`
	Attribute11expr             types.String `tfsdk:"attribute11expr"`
	Attribute11format           types.String `tfsdk:"attribute11format"`
	Attribute11friendlyname     types.String `tfsdk:"attribute11friendlyname"`
	Attribute12                 types.String `tfsdk:"attribute12"`
	Attribute12expr             types.String `tfsdk:"attribute12expr"`
	Attribute12format           types.String `tfsdk:"attribute12format"`
	Attribute12friendlyname     types.String `tfsdk:"attribute12friendlyname"`
	Attribute13                 types.String `tfsdk:"attribute13"`
	Attribute13expr             types.String `tfsdk:"attribute13expr"`
	Attribute13format           types.String `tfsdk:"attribute13format"`
	Attribute13friendlyname     types.String `tfsdk:"attribute13friendlyname"`
	Attribute14                 types.String `tfsdk:"attribute14"`
	Attribute14expr             types.String `tfsdk:"attribute14expr"`
	Attribute14format           types.String `tfsdk:"attribute14format"`
	Attribute14friendlyname     types.String `tfsdk:"attribute14friendlyname"`
	Attribute15                 types.String `tfsdk:"attribute15"`
	Attribute15expr             types.String `tfsdk:"attribute15expr"`
	Attribute15format           types.String `tfsdk:"attribute15format"`
	Attribute15friendlyname     types.String `tfsdk:"attribute15friendlyname"`
	Attribute16                 types.String `tfsdk:"attribute16"`
	Attribute16expr             types.String `tfsdk:"attribute16expr"`
	Attribute16format           types.String `tfsdk:"attribute16format"`
	Attribute16friendlyname     types.String `tfsdk:"attribute16friendlyname"`
	Attribute1expr              types.String `tfsdk:"attribute1expr"`
	Attribute1format            types.String `tfsdk:"attribute1format"`
	Attribute1friendlyname      types.String `tfsdk:"attribute1friendlyname"`
	Attribute2                  types.String `tfsdk:"attribute2"`
	Attribute2expr              types.String `tfsdk:"attribute2expr"`
	Attribute2format            types.String `tfsdk:"attribute2format"`
	Attribute2friendlyname      types.String `tfsdk:"attribute2friendlyname"`
	Attribute3                  types.String `tfsdk:"attribute3"`
	Attribute3expr              types.String `tfsdk:"attribute3expr"`
	Attribute3format            types.String `tfsdk:"attribute3format"`
	Attribute3friendlyname      types.String `tfsdk:"attribute3friendlyname"`
	Attribute4                  types.String `tfsdk:"attribute4"`
	Attribute4expr              types.String `tfsdk:"attribute4expr"`
	Attribute4format            types.String `tfsdk:"attribute4format"`
	Attribute4friendlyname      types.String `tfsdk:"attribute4friendlyname"`
	Attribute5                  types.String `tfsdk:"attribute5"`
	Attribute5expr              types.String `tfsdk:"attribute5expr"`
	Attribute5format            types.String `tfsdk:"attribute5format"`
	Attribute5friendlyname      types.String `tfsdk:"attribute5friendlyname"`
	Attribute6                  types.String `tfsdk:"attribute6"`
	Attribute6expr              types.String `tfsdk:"attribute6expr"`
	Attribute6format            types.String `tfsdk:"attribute6format"`
	Attribute6friendlyname      types.String `tfsdk:"attribute6friendlyname"`
	Attribute7                  types.String `tfsdk:"attribute7"`
	Attribute7expr              types.String `tfsdk:"attribute7expr"`
	Attribute7format            types.String `tfsdk:"attribute7format"`
	Attribute7friendlyname      types.String `tfsdk:"attribute7friendlyname"`
	Attribute8                  types.String `tfsdk:"attribute8"`
	Attribute8expr              types.String `tfsdk:"attribute8expr"`
	Attribute8format            types.String `tfsdk:"attribute8format"`
	Attribute8friendlyname      types.String `tfsdk:"attribute8friendlyname"`
	Attribute9                  types.String `tfsdk:"attribute9"`
	Attribute9expr              types.String `tfsdk:"attribute9expr"`
	Attribute9format            types.String `tfsdk:"attribute9format"`
	Attribute9friendlyname      types.String `tfsdk:"attribute9friendlyname"`
	Audience                    types.String `tfsdk:"audience"`
	Defaultauthenticationgroup  types.String `tfsdk:"defaultauthenticationgroup"`
	Digestmethod                types.String `tfsdk:"digestmethod"`
	Encryptassertion            types.String `tfsdk:"encryptassertion"`
	Encryptionalgorithm         types.String `tfsdk:"encryptionalgorithm"`
	Keytransportalg             types.String `tfsdk:"keytransportalg"`
	Logoutbinding               types.String `tfsdk:"logoutbinding"`
	Metadatarefreshinterval     types.Int64  `tfsdk:"metadatarefreshinterval"`
	Metadataurl                 types.String `tfsdk:"metadataurl"`
	Name                        types.String `tfsdk:"name"`
	Nameidexpr                  types.String `tfsdk:"nameidexpr"`
	Nameidformat                types.String `tfsdk:"nameidformat"`
	Rejectunsignedrequests      types.String `tfsdk:"rejectunsignedrequests"`
	Samlbinding                 types.String `tfsdk:"samlbinding"`
	Samlidpcertname             types.String `tfsdk:"samlidpcertname"`
	Samlissuername              types.String `tfsdk:"samlissuername"`
	Samlsigningcertversion      types.String `tfsdk:"samlsigningcertversion"`
	Samlspcertname              types.String `tfsdk:"samlspcertname"`
	Samlspcertversion           types.String `tfsdk:"samlspcertversion"`
	Sendpassword                types.String `tfsdk:"sendpassword"`
	Serviceproviderid           types.String `tfsdk:"serviceproviderid"`
	Signassertion               types.String `tfsdk:"signassertion"`
	Signaturealg                types.String `tfsdk:"signaturealg"`
	Signatureservice            types.String `tfsdk:"signatureservice"`
	Skewtime                    types.Int64  `tfsdk:"skewtime"`
	Splogouturl                 types.String `tfsdk:"splogouturl"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/authenticationsamlidpprofile.json). Never settable;
	// populated from GET.
	Metadataimportstatus types.String `tfsdk:"metadataimportstatus"`
}

func AuthenticationsamlidpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acsurlrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to allow Assertion Consumer Service URI coming in the SAML Request",
			},
			"assertionconsumerserviceurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which the assertion is to be sent.",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute1 that needs to be sent in SAML Assertion",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute10 that needs to be sent in SAML Assertion",
			},
			"attribute10expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute10's value to be sent in Assertion",
			},
			"attribute10format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute10 to be sent in Assertion.",
			},
			"attribute10friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute10 that needs to be sent in SAML Assertion",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute11 that needs to be sent in SAML Assertion",
			},
			"attribute11expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute11's value to be sent in Assertion",
			},
			"attribute11format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute11 to be sent in Assertion.",
			},
			"attribute11friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute11 that needs to be sent in SAML Assertion",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute12 that needs to be sent in SAML Assertion",
			},
			"attribute12expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute12's value to be sent in Assertion",
			},
			"attribute12format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute12 to be sent in Assertion.",
			},
			"attribute12friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute12 that needs to be sent in SAML Assertion",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute13 that needs to be sent in SAML Assertion",
			},
			"attribute13expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute13's value to be sent in Assertion",
			},
			"attribute13format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute13 to be sent in Assertion.",
			},
			"attribute13friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute13 that needs to be sent in SAML Assertion",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute14 that needs to be sent in SAML Assertion",
			},
			"attribute14expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute14's value to be sent in Assertion",
			},
			"attribute14format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute14 to be sent in Assertion.",
			},
			"attribute14friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute14 that needs to be sent in SAML Assertion",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute15 that needs to be sent in SAML Assertion",
			},
			"attribute15expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute15's value to be sent in Assertion",
			},
			"attribute15format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute15 to be sent in Assertion.",
			},
			"attribute15friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute15 that needs to be sent in SAML Assertion",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute16 that needs to be sent in SAML Assertion",
			},
			"attribute16expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute16's value to be sent in Assertion",
			},
			"attribute16format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute16 to be sent in Assertion.",
			},
			"attribute16friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute16 that needs to be sent in SAML Assertion",
			},
			"attribute1expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute1's value to be sent in Assertion",
			},
			"attribute1format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute1 to be sent in Assertion.",
			},
			"attribute1friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute1 that needs to be sent in SAML Assertion",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute2 that needs to be sent in SAML Assertion",
			},
			"attribute2expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute2's value to be sent in Assertion",
			},
			"attribute2format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute2 to be sent in Assertion.",
			},
			"attribute2friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute2 that needs to be sent in SAML Assertion",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute3 that needs to be sent in SAML Assertion",
			},
			"attribute3expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute3's value to be sent in Assertion",
			},
			"attribute3format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute3 to be sent in Assertion.",
			},
			"attribute3friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute3 that needs to be sent in SAML Assertion",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute4 that needs to be sent in SAML Assertion",
			},
			"attribute4expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute4's value to be sent in Assertion",
			},
			"attribute4format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute4 to be sent in Assertion.",
			},
			"attribute4friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute4 that needs to be sent in SAML Assertion",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute5 that needs to be sent in SAML Assertion",
			},
			"attribute5expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute5's value to be sent in Assertion",
			},
			"attribute5format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute5 to be sent in Assertion.",
			},
			"attribute5friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute5 that needs to be sent in SAML Assertion",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute6 that needs to be sent in SAML Assertion",
			},
			"attribute6expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute6's value to be sent in Assertion",
			},
			"attribute6format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute6 to be sent in Assertion.",
			},
			"attribute6friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute6 that needs to be sent in SAML Assertion",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute7 that needs to be sent in SAML Assertion",
			},
			"attribute7expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute7's value to be sent in Assertion",
			},
			"attribute7format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute7 to be sent in Assertion.",
			},
			"attribute7friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute7 that needs to be sent in SAML Assertion",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute8 that needs to be sent in SAML Assertion",
			},
			"attribute8expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute8's value to be sent in Assertion",
			},
			"attribute8format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute8 to be sent in Assertion.",
			},
			"attribute8friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute8 that needs to be sent in SAML Assertion",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of attribute9 that needs to be sent in SAML Assertion",
			},
			"attribute9expr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain attribute9's value to be sent in Assertion",
			},
			"attribute9format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Attribute9 to be sent in Assertion.",
			},
			"attribute9friendlyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User-Friendly Name of attribute9 that needs to be sent in SAML Assertion",
			},
			"audience": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This group will be part of AAA session's internal group list. This will be helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party. In authentication policy AAA.USER.IS_MEMBER_OF(\"<default_auth_group>\")  is way to use this feature.",
			},
			"digestmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to compute/verify digest for SAML transactions",
			},
			"encryptassertion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to encrypt assertion when Citrix ADC IDP sends one.",
			},
			"encryptionalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to encrypt SAML assertion",
			},
			"keytransportalg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key transport algorithm to be used in encryption of SAML assertion",
			},
			"logoutbinding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the transport mechanism of saml logout messages.",
			},
			"metadatarefreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval in minute for fetching metadata from specified metadata URL",
			},
			"metadataurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This URL is used for obtaining samlidp metadata",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new saml single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"nameidexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that will be evaluated to obtain NameIdentifier to be sent in assertion",
			},
			"nameidformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of Name Identifier sent in Assertion.",
			},
			"rejectunsignedrequests": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to Reject unsigned SAML Requests. ON option denies any authentication requests that arrive without signature.",
			},
			"samlbinding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This element specifies the transport mechanism of saml messages.",
			},
			"samlidpcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the certificate used to sign the SAMLResposne that is sent to Relying Party or Service Provider after successful authentication",
			},
			"samlissuername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.",
			},
			"samlsigningcertversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "version of the certificate in signature service used to sign the SAMLResposne that is sent to Relying Party or Service Provider after successful authentication",
			},
			"samlspcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate of SAML Relying Party. This certificate is used to verify signature of the incoming AuthnRequest from a Relying Party or Service Provider",
			},
			"samlspcertversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "version of the certificate in signature service used to verify the signature of the incoming AuthnRequest from a Relying Party or Service Provider",
			},
			"sendpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to send password in assertion.",
			},
			"serviceproviderid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identifier of the Service Provider that sends SAML Request. Citrix ADC will ensure that the Issuer of the SAML Request matches this URI. In case of SP initiated sign-in scenarios, this value must be same as samlIssuerName configured in samlAction.",
			},
			"signassertion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to sign/verify SAML transactions",
			},
			"signatureservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service in cloud used to sign the data",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
			"splogouturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Endpoint on the ServiceProvider (SP) to which logout messages are to be sent",
			},
			"metadataimportstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes metadata import status. Possible values = INIT, SUCCESS, FAIL.",
			},
		},
	}
}

// authenticationsamlidpprofileDataSourceSetAttrFromGet projects a NITRO
// authenticationsamlidpprofile GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationsamlidpprofileDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationsamlidpprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationsamlidpprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Acsurlrule = utils.MapGetString(g, "acsurlrule")
	data.Assertionconsumerserviceurl = utils.MapGetString(g, "assertionconsumerserviceurl")
	data.Attribute1 = utils.MapGetString(g, "attribute1")
	data.Attribute10 = utils.MapGetString(g, "attribute10")
	data.Attribute10expr = utils.MapGetString(g, "attribute10expr")
	data.Attribute10format = utils.MapGetString(g, "attribute10format")
	data.Attribute10friendlyname = utils.MapGetString(g, "attribute10friendlyname")
	data.Attribute11 = utils.MapGetString(g, "attribute11")
	data.Attribute11expr = utils.MapGetString(g, "attribute11expr")
	data.Attribute11format = utils.MapGetString(g, "attribute11format")
	data.Attribute11friendlyname = utils.MapGetString(g, "attribute11friendlyname")
	data.Attribute12 = utils.MapGetString(g, "attribute12")
	data.Attribute12expr = utils.MapGetString(g, "attribute12expr")
	data.Attribute12format = utils.MapGetString(g, "attribute12format")
	data.Attribute12friendlyname = utils.MapGetString(g, "attribute12friendlyname")
	data.Attribute13 = utils.MapGetString(g, "attribute13")
	data.Attribute13expr = utils.MapGetString(g, "attribute13expr")
	data.Attribute13format = utils.MapGetString(g, "attribute13format")
	data.Attribute13friendlyname = utils.MapGetString(g, "attribute13friendlyname")
	data.Attribute14 = utils.MapGetString(g, "attribute14")
	data.Attribute14expr = utils.MapGetString(g, "attribute14expr")
	data.Attribute14format = utils.MapGetString(g, "attribute14format")
	data.Attribute14friendlyname = utils.MapGetString(g, "attribute14friendlyname")
	data.Attribute15 = utils.MapGetString(g, "attribute15")
	data.Attribute15expr = utils.MapGetString(g, "attribute15expr")
	data.Attribute15format = utils.MapGetString(g, "attribute15format")
	data.Attribute15friendlyname = utils.MapGetString(g, "attribute15friendlyname")
	data.Attribute16 = utils.MapGetString(g, "attribute16")
	data.Attribute16expr = utils.MapGetString(g, "attribute16expr")
	data.Attribute16format = utils.MapGetString(g, "attribute16format")
	data.Attribute16friendlyname = utils.MapGetString(g, "attribute16friendlyname")
	data.Attribute1expr = utils.MapGetString(g, "attribute1expr")
	data.Attribute1format = utils.MapGetString(g, "attribute1format")
	data.Attribute1friendlyname = utils.MapGetString(g, "attribute1friendlyname")
	data.Attribute2 = utils.MapGetString(g, "attribute2")
	data.Attribute2expr = utils.MapGetString(g, "attribute2expr")
	data.Attribute2format = utils.MapGetString(g, "attribute2format")
	data.Attribute2friendlyname = utils.MapGetString(g, "attribute2friendlyname")
	data.Attribute3 = utils.MapGetString(g, "attribute3")
	data.Attribute3expr = utils.MapGetString(g, "attribute3expr")
	data.Attribute3format = utils.MapGetString(g, "attribute3format")
	data.Attribute3friendlyname = utils.MapGetString(g, "attribute3friendlyname")
	data.Attribute4 = utils.MapGetString(g, "attribute4")
	data.Attribute4expr = utils.MapGetString(g, "attribute4expr")
	data.Attribute4format = utils.MapGetString(g, "attribute4format")
	data.Attribute4friendlyname = utils.MapGetString(g, "attribute4friendlyname")
	data.Attribute5 = utils.MapGetString(g, "attribute5")
	data.Attribute5expr = utils.MapGetString(g, "attribute5expr")
	data.Attribute5format = utils.MapGetString(g, "attribute5format")
	data.Attribute5friendlyname = utils.MapGetString(g, "attribute5friendlyname")
	data.Attribute6 = utils.MapGetString(g, "attribute6")
	data.Attribute6expr = utils.MapGetString(g, "attribute6expr")
	data.Attribute6format = utils.MapGetString(g, "attribute6format")
	data.Attribute6friendlyname = utils.MapGetString(g, "attribute6friendlyname")
	data.Attribute7 = utils.MapGetString(g, "attribute7")
	data.Attribute7expr = utils.MapGetString(g, "attribute7expr")
	data.Attribute7format = utils.MapGetString(g, "attribute7format")
	data.Attribute7friendlyname = utils.MapGetString(g, "attribute7friendlyname")
	data.Attribute8 = utils.MapGetString(g, "attribute8")
	data.Attribute8expr = utils.MapGetString(g, "attribute8expr")
	data.Attribute8format = utils.MapGetString(g, "attribute8format")
	data.Attribute8friendlyname = utils.MapGetString(g, "attribute8friendlyname")
	data.Attribute9 = utils.MapGetString(g, "attribute9")
	data.Attribute9expr = utils.MapGetString(g, "attribute9expr")
	data.Attribute9format = utils.MapGetString(g, "attribute9format")
	data.Attribute9friendlyname = utils.MapGetString(g, "attribute9friendlyname")
	data.Audience = utils.MapGetString(g, "audience")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Digestmethod = utils.MapGetString(g, "digestmethod")
	data.Encryptassertion = utils.MapGetString(g, "encryptassertion")
	data.Encryptionalgorithm = utils.MapGetString(g, "encryptionalgorithm")
	data.Keytransportalg = utils.MapGetString(g, "keytransportalg")
	data.Logoutbinding = utils.MapGetString(g, "logoutbinding")
	data.Metadatarefreshinterval = utils.MapGetInt64(g, "metadatarefreshinterval")
	data.Metadataurl = utils.MapGetString(g, "metadataurl")
	data.Nameidexpr = utils.MapGetString(g, "nameidexpr")
	data.Nameidformat = utils.MapGetString(g, "nameidformat")
	data.Rejectunsignedrequests = utils.MapGetString(g, "rejectunsignedrequests")
	data.Samlbinding = utils.MapGetString(g, "samlbinding")
	data.Samlidpcertname = utils.MapGetString(g, "samlidpcertname")
	data.Samlissuername = utils.MapGetString(g, "samlissuername")
	data.Samlsigningcertversion = utils.MapGetString(g, "samlsigningcertversion")
	data.Samlspcertname = utils.MapGetString(g, "samlspcertname")
	data.Samlspcertversion = utils.MapGetString(g, "samlspcertversion")
	data.Sendpassword = utils.MapGetString(g, "sendpassword")
	data.Serviceproviderid = utils.MapGetString(g, "serviceproviderid")
	data.Signassertion = utils.MapGetString(g, "signassertion")
	data.Signaturealg = utils.MapGetString(g, "signaturealg")
	data.Signatureservice = utils.MapGetString(g, "signatureservice")
	data.Skewtime = utils.MapGetInt64(g, "skewtime")
	data.Splogouturl = utils.MapGetString(g, "splogouturl")

	// Read-only attribute.
	data.Metadataimportstatus = utils.MapGetString(g, "metadataimportstatus")
}
