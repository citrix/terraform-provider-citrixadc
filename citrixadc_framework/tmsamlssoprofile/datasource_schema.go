package tmsamlssoprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TmsamlssoprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
			"digestmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to compute/verify digest for SAML transactions",
			},
			"encryptassertion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to encrypt assertion when Citrix ADC sends one.",
			},
			"encryptionalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to be used to encrypt SAML assertion",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new saml single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
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
			"relaystaterule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a targ\net url to which user is redirected after the recipient validates SAML token",
			},
			"samlissuername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.",
			},
			"samlsigningcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate that is used to Sign Assertion.",
			},
			"samlspcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate of peer/receving party using which Assertion is encrypted.",
			},
			"sendpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to send password in assertion.",
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
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
		},
	}
}
