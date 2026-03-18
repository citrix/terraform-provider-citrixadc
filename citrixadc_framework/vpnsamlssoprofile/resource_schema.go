package vpnsamlssoprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnsamlssoprofileResourceModel describes the resource data model.
type VpnsamlssoprofileResourceModel struct {
	Id                          types.String `tfsdk:"id"`
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
	Digestmethod                types.String `tfsdk:"digestmethod"`
	Encryptassertion            types.String `tfsdk:"encryptassertion"`
	Encryptionalgorithm         types.String `tfsdk:"encryptionalgorithm"`
	Name                        types.String `tfsdk:"name"`
	Nameidexpr                  types.String `tfsdk:"nameidexpr"`
	Nameidformat                types.String `tfsdk:"nameidformat"`
	Relaystaterule              types.String `tfsdk:"relaystaterule"`
	Samlissuername              types.String `tfsdk:"samlissuername"`
	Samlsigningcertname         types.String `tfsdk:"samlsigningcertname"`
	Samlspcertname              types.String `tfsdk:"samlspcertname"`
	Sendpassword                types.String `tfsdk:"sendpassword"`
	Signassertion               types.String `tfsdk:"signassertion"`
	Signaturealg                types.String `tfsdk:"signaturealg"`
	Signatureservice            types.String `tfsdk:"signatureservice"`
	Skewtime                    types.Int64  `tfsdk:"skewtime"`
}

func (r *VpnsamlssoprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnsamlssoprofile resource.",
			},
			"assertionconsumerserviceurl": schema.StringAttribute{
				Required:    true,
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
				Default:     stringdefault.StaticString("SHA256"),
				Description: "Algorithm to be used to compute/verify digest for SAML transactions",
			},
			"encryptassertion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to encrypt assertion when Citrix ADC sends one.",
			},
			"encryptionalgorithm": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AES256"),
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
				Default:     stringdefault.StaticString("transient"),
				Description: "Format of Name Identifier sent in Assertion.",
			},
			"relaystaterule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a target url to which user is redirected after the recipient validates SAML token",
			},
			"samlissuername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.",
			},
			"samlsigningcertname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the signing authority as given in the SAML server's SSL certificate.",
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
				Default:     stringdefault.StaticString("ASSERTION"),
				Description: "Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RSA-SHA256"),
				Description: "Algorithm to be used to sign/verify SAML transactions",
			},
			"signatureservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service in cloud used to sign the data",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
		},
	}
}

func vpnsamlssoprofileGetThePayloadFromtheConfig(ctx context.Context, data *VpnsamlssoprofileResourceModel) vpn.Vpnsamlssoprofile {
	tflog.Debug(ctx, "In vpnsamlssoprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnsamlssoprofile := vpn.Vpnsamlssoprofile{}
	if !data.Assertionconsumerserviceurl.IsNull() {
		vpnsamlssoprofile.Assertionconsumerserviceurl = data.Assertionconsumerserviceurl.ValueString()
	}
	if !data.Attribute1.IsNull() {
		vpnsamlssoprofile.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		vpnsamlssoprofile.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute10expr.IsNull() {
		vpnsamlssoprofile.Attribute10expr = data.Attribute10expr.ValueString()
	}
	if !data.Attribute10format.IsNull() {
		vpnsamlssoprofile.Attribute10format = data.Attribute10format.ValueString()
	}
	if !data.Attribute10friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute10friendlyname = data.Attribute10friendlyname.ValueString()
	}
	if !data.Attribute11.IsNull() {
		vpnsamlssoprofile.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute11expr.IsNull() {
		vpnsamlssoprofile.Attribute11expr = data.Attribute11expr.ValueString()
	}
	if !data.Attribute11format.IsNull() {
		vpnsamlssoprofile.Attribute11format = data.Attribute11format.ValueString()
	}
	if !data.Attribute11friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute11friendlyname = data.Attribute11friendlyname.ValueString()
	}
	if !data.Attribute12.IsNull() {
		vpnsamlssoprofile.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute12expr.IsNull() {
		vpnsamlssoprofile.Attribute12expr = data.Attribute12expr.ValueString()
	}
	if !data.Attribute12format.IsNull() {
		vpnsamlssoprofile.Attribute12format = data.Attribute12format.ValueString()
	}
	if !data.Attribute12friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute12friendlyname = data.Attribute12friendlyname.ValueString()
	}
	if !data.Attribute13.IsNull() {
		vpnsamlssoprofile.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute13expr.IsNull() {
		vpnsamlssoprofile.Attribute13expr = data.Attribute13expr.ValueString()
	}
	if !data.Attribute13format.IsNull() {
		vpnsamlssoprofile.Attribute13format = data.Attribute13format.ValueString()
	}
	if !data.Attribute13friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute13friendlyname = data.Attribute13friendlyname.ValueString()
	}
	if !data.Attribute14.IsNull() {
		vpnsamlssoprofile.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute14expr.IsNull() {
		vpnsamlssoprofile.Attribute14expr = data.Attribute14expr.ValueString()
	}
	if !data.Attribute14format.IsNull() {
		vpnsamlssoprofile.Attribute14format = data.Attribute14format.ValueString()
	}
	if !data.Attribute14friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute14friendlyname = data.Attribute14friendlyname.ValueString()
	}
	if !data.Attribute15.IsNull() {
		vpnsamlssoprofile.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute15expr.IsNull() {
		vpnsamlssoprofile.Attribute15expr = data.Attribute15expr.ValueString()
	}
	if !data.Attribute15format.IsNull() {
		vpnsamlssoprofile.Attribute15format = data.Attribute15format.ValueString()
	}
	if !data.Attribute15friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute15friendlyname = data.Attribute15friendlyname.ValueString()
	}
	if !data.Attribute16.IsNull() {
		vpnsamlssoprofile.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute16expr.IsNull() {
		vpnsamlssoprofile.Attribute16expr = data.Attribute16expr.ValueString()
	}
	if !data.Attribute16format.IsNull() {
		vpnsamlssoprofile.Attribute16format = data.Attribute16format.ValueString()
	}
	if !data.Attribute16friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute16friendlyname = data.Attribute16friendlyname.ValueString()
	}
	if !data.Attribute1expr.IsNull() {
		vpnsamlssoprofile.Attribute1expr = data.Attribute1expr.ValueString()
	}
	if !data.Attribute1format.IsNull() {
		vpnsamlssoprofile.Attribute1format = data.Attribute1format.ValueString()
	}
	if !data.Attribute1friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute1friendlyname = data.Attribute1friendlyname.ValueString()
	}
	if !data.Attribute2.IsNull() {
		vpnsamlssoprofile.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute2expr.IsNull() {
		vpnsamlssoprofile.Attribute2expr = data.Attribute2expr.ValueString()
	}
	if !data.Attribute2format.IsNull() {
		vpnsamlssoprofile.Attribute2format = data.Attribute2format.ValueString()
	}
	if !data.Attribute2friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute2friendlyname = data.Attribute2friendlyname.ValueString()
	}
	if !data.Attribute3.IsNull() {
		vpnsamlssoprofile.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute3expr.IsNull() {
		vpnsamlssoprofile.Attribute3expr = data.Attribute3expr.ValueString()
	}
	if !data.Attribute3format.IsNull() {
		vpnsamlssoprofile.Attribute3format = data.Attribute3format.ValueString()
	}
	if !data.Attribute3friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute3friendlyname = data.Attribute3friendlyname.ValueString()
	}
	if !data.Attribute4.IsNull() {
		vpnsamlssoprofile.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute4expr.IsNull() {
		vpnsamlssoprofile.Attribute4expr = data.Attribute4expr.ValueString()
	}
	if !data.Attribute4format.IsNull() {
		vpnsamlssoprofile.Attribute4format = data.Attribute4format.ValueString()
	}
	if !data.Attribute4friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute4friendlyname = data.Attribute4friendlyname.ValueString()
	}
	if !data.Attribute5.IsNull() {
		vpnsamlssoprofile.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute5expr.IsNull() {
		vpnsamlssoprofile.Attribute5expr = data.Attribute5expr.ValueString()
	}
	if !data.Attribute5format.IsNull() {
		vpnsamlssoprofile.Attribute5format = data.Attribute5format.ValueString()
	}
	if !data.Attribute5friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute5friendlyname = data.Attribute5friendlyname.ValueString()
	}
	if !data.Attribute6.IsNull() {
		vpnsamlssoprofile.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute6expr.IsNull() {
		vpnsamlssoprofile.Attribute6expr = data.Attribute6expr.ValueString()
	}
	if !data.Attribute6format.IsNull() {
		vpnsamlssoprofile.Attribute6format = data.Attribute6format.ValueString()
	}
	if !data.Attribute6friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute6friendlyname = data.Attribute6friendlyname.ValueString()
	}
	if !data.Attribute7.IsNull() {
		vpnsamlssoprofile.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute7expr.IsNull() {
		vpnsamlssoprofile.Attribute7expr = data.Attribute7expr.ValueString()
	}
	if !data.Attribute7format.IsNull() {
		vpnsamlssoprofile.Attribute7format = data.Attribute7format.ValueString()
	}
	if !data.Attribute7friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute7friendlyname = data.Attribute7friendlyname.ValueString()
	}
	if !data.Attribute8.IsNull() {
		vpnsamlssoprofile.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute8expr.IsNull() {
		vpnsamlssoprofile.Attribute8expr = data.Attribute8expr.ValueString()
	}
	if !data.Attribute8format.IsNull() {
		vpnsamlssoprofile.Attribute8format = data.Attribute8format.ValueString()
	}
	if !data.Attribute8friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute8friendlyname = data.Attribute8friendlyname.ValueString()
	}
	if !data.Attribute9.IsNull() {
		vpnsamlssoprofile.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Attribute9expr.IsNull() {
		vpnsamlssoprofile.Attribute9expr = data.Attribute9expr.ValueString()
	}
	if !data.Attribute9format.IsNull() {
		vpnsamlssoprofile.Attribute9format = data.Attribute9format.ValueString()
	}
	if !data.Attribute9friendlyname.IsNull() {
		vpnsamlssoprofile.Attribute9friendlyname = data.Attribute9friendlyname.ValueString()
	}
	if !data.Audience.IsNull() {
		vpnsamlssoprofile.Audience = data.Audience.ValueString()
	}
	if !data.Digestmethod.IsNull() {
		vpnsamlssoprofile.Digestmethod = data.Digestmethod.ValueString()
	}
	if !data.Encryptassertion.IsNull() {
		vpnsamlssoprofile.Encryptassertion = data.Encryptassertion.ValueString()
	}
	if !data.Encryptionalgorithm.IsNull() {
		vpnsamlssoprofile.Encryptionalgorithm = data.Encryptionalgorithm.ValueString()
	}
	if !data.Name.IsNull() {
		vpnsamlssoprofile.Name = data.Name.ValueString()
	}
	if !data.Nameidexpr.IsNull() {
		vpnsamlssoprofile.Nameidexpr = data.Nameidexpr.ValueString()
	}
	if !data.Nameidformat.IsNull() {
		vpnsamlssoprofile.Nameidformat = data.Nameidformat.ValueString()
	}
	if !data.Relaystaterule.IsNull() {
		vpnsamlssoprofile.Relaystaterule = data.Relaystaterule.ValueString()
	}
	if !data.Samlissuername.IsNull() {
		vpnsamlssoprofile.Samlissuername = data.Samlissuername.ValueString()
	}
	if !data.Samlsigningcertname.IsNull() {
		vpnsamlssoprofile.Samlsigningcertname = data.Samlsigningcertname.ValueString()
	}
	if !data.Samlspcertname.IsNull() {
		vpnsamlssoprofile.Samlspcertname = data.Samlspcertname.ValueString()
	}
	if !data.Sendpassword.IsNull() {
		vpnsamlssoprofile.Sendpassword = data.Sendpassword.ValueString()
	}
	if !data.Signassertion.IsNull() {
		vpnsamlssoprofile.Signassertion = data.Signassertion.ValueString()
	}
	if !data.Signaturealg.IsNull() {
		vpnsamlssoprofile.Signaturealg = data.Signaturealg.ValueString()
	}
	if !data.Signatureservice.IsNull() {
		vpnsamlssoprofile.Signatureservice = data.Signatureservice.ValueString()
	}
	if !data.Skewtime.IsNull() {
		vpnsamlssoprofile.Skewtime = utils.IntPtr(int(data.Skewtime.ValueInt64()))
	}

	return vpnsamlssoprofile
}

func vpnsamlssoprofileSetAttrFromGet(ctx context.Context, data *VpnsamlssoprofileResourceModel, getResponseData map[string]interface{}) *VpnsamlssoprofileResourceModel {
	tflog.Debug(ctx, "In vpnsamlssoprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["assertionconsumerserviceurl"]; ok && val != nil {
		data.Assertionconsumerserviceurl = types.StringValue(val.(string))
	} else {
		data.Assertionconsumerserviceurl = types.StringNull()
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
	if val, ok := getResponseData["attribute10expr"]; ok && val != nil {
		data.Attribute10expr = types.StringValue(val.(string))
	} else {
		data.Attribute10expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute10format"]; ok && val != nil {
		data.Attribute10format = types.StringValue(val.(string))
	} else {
		data.Attribute10format = types.StringNull()
	}
	if val, ok := getResponseData["attribute10friendlyname"]; ok && val != nil {
		data.Attribute10friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute10friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute11"]; ok && val != nil {
		data.Attribute11 = types.StringValue(val.(string))
	} else {
		data.Attribute11 = types.StringNull()
	}
	if val, ok := getResponseData["attribute11expr"]; ok && val != nil {
		data.Attribute11expr = types.StringValue(val.(string))
	} else {
		data.Attribute11expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute11format"]; ok && val != nil {
		data.Attribute11format = types.StringValue(val.(string))
	} else {
		data.Attribute11format = types.StringNull()
	}
	if val, ok := getResponseData["attribute11friendlyname"]; ok && val != nil {
		data.Attribute11friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute11friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute12"]; ok && val != nil {
		data.Attribute12 = types.StringValue(val.(string))
	} else {
		data.Attribute12 = types.StringNull()
	}
	if val, ok := getResponseData["attribute12expr"]; ok && val != nil {
		data.Attribute12expr = types.StringValue(val.(string))
	} else {
		data.Attribute12expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute12format"]; ok && val != nil {
		data.Attribute12format = types.StringValue(val.(string))
	} else {
		data.Attribute12format = types.StringNull()
	}
	if val, ok := getResponseData["attribute12friendlyname"]; ok && val != nil {
		data.Attribute12friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute12friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute13"]; ok && val != nil {
		data.Attribute13 = types.StringValue(val.(string))
	} else {
		data.Attribute13 = types.StringNull()
	}
	if val, ok := getResponseData["attribute13expr"]; ok && val != nil {
		data.Attribute13expr = types.StringValue(val.(string))
	} else {
		data.Attribute13expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute13format"]; ok && val != nil {
		data.Attribute13format = types.StringValue(val.(string))
	} else {
		data.Attribute13format = types.StringNull()
	}
	if val, ok := getResponseData["attribute13friendlyname"]; ok && val != nil {
		data.Attribute13friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute13friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute14"]; ok && val != nil {
		data.Attribute14 = types.StringValue(val.(string))
	} else {
		data.Attribute14 = types.StringNull()
	}
	if val, ok := getResponseData["attribute14expr"]; ok && val != nil {
		data.Attribute14expr = types.StringValue(val.(string))
	} else {
		data.Attribute14expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute14format"]; ok && val != nil {
		data.Attribute14format = types.StringValue(val.(string))
	} else {
		data.Attribute14format = types.StringNull()
	}
	if val, ok := getResponseData["attribute14friendlyname"]; ok && val != nil {
		data.Attribute14friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute14friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute15"]; ok && val != nil {
		data.Attribute15 = types.StringValue(val.(string))
	} else {
		data.Attribute15 = types.StringNull()
	}
	if val, ok := getResponseData["attribute15expr"]; ok && val != nil {
		data.Attribute15expr = types.StringValue(val.(string))
	} else {
		data.Attribute15expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute15format"]; ok && val != nil {
		data.Attribute15format = types.StringValue(val.(string))
	} else {
		data.Attribute15format = types.StringNull()
	}
	if val, ok := getResponseData["attribute15friendlyname"]; ok && val != nil {
		data.Attribute15friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute15friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute16"]; ok && val != nil {
		data.Attribute16 = types.StringValue(val.(string))
	} else {
		data.Attribute16 = types.StringNull()
	}
	if val, ok := getResponseData["attribute16expr"]; ok && val != nil {
		data.Attribute16expr = types.StringValue(val.(string))
	} else {
		data.Attribute16expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute16format"]; ok && val != nil {
		data.Attribute16format = types.StringValue(val.(string))
	} else {
		data.Attribute16format = types.StringNull()
	}
	if val, ok := getResponseData["attribute16friendlyname"]; ok && val != nil {
		data.Attribute16friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute16friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute1expr"]; ok && val != nil {
		data.Attribute1expr = types.StringValue(val.(string))
	} else {
		data.Attribute1expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute1format"]; ok && val != nil {
		data.Attribute1format = types.StringValue(val.(string))
	} else {
		data.Attribute1format = types.StringNull()
	}
	if val, ok := getResponseData["attribute1friendlyname"]; ok && val != nil {
		data.Attribute1friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute1friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute2"]; ok && val != nil {
		data.Attribute2 = types.StringValue(val.(string))
	} else {
		data.Attribute2 = types.StringNull()
	}
	if val, ok := getResponseData["attribute2expr"]; ok && val != nil {
		data.Attribute2expr = types.StringValue(val.(string))
	} else {
		data.Attribute2expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute2format"]; ok && val != nil {
		data.Attribute2format = types.StringValue(val.(string))
	} else {
		data.Attribute2format = types.StringNull()
	}
	if val, ok := getResponseData["attribute2friendlyname"]; ok && val != nil {
		data.Attribute2friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute2friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute3"]; ok && val != nil {
		data.Attribute3 = types.StringValue(val.(string))
	} else {
		data.Attribute3 = types.StringNull()
	}
	if val, ok := getResponseData["attribute3expr"]; ok && val != nil {
		data.Attribute3expr = types.StringValue(val.(string))
	} else {
		data.Attribute3expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute3format"]; ok && val != nil {
		data.Attribute3format = types.StringValue(val.(string))
	} else {
		data.Attribute3format = types.StringNull()
	}
	if val, ok := getResponseData["attribute3friendlyname"]; ok && val != nil {
		data.Attribute3friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute3friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute4"]; ok && val != nil {
		data.Attribute4 = types.StringValue(val.(string))
	} else {
		data.Attribute4 = types.StringNull()
	}
	if val, ok := getResponseData["attribute4expr"]; ok && val != nil {
		data.Attribute4expr = types.StringValue(val.(string))
	} else {
		data.Attribute4expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute4format"]; ok && val != nil {
		data.Attribute4format = types.StringValue(val.(string))
	} else {
		data.Attribute4format = types.StringNull()
	}
	if val, ok := getResponseData["attribute4friendlyname"]; ok && val != nil {
		data.Attribute4friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute4friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute5"]; ok && val != nil {
		data.Attribute5 = types.StringValue(val.(string))
	} else {
		data.Attribute5 = types.StringNull()
	}
	if val, ok := getResponseData["attribute5expr"]; ok && val != nil {
		data.Attribute5expr = types.StringValue(val.(string))
	} else {
		data.Attribute5expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute5format"]; ok && val != nil {
		data.Attribute5format = types.StringValue(val.(string))
	} else {
		data.Attribute5format = types.StringNull()
	}
	if val, ok := getResponseData["attribute5friendlyname"]; ok && val != nil {
		data.Attribute5friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute5friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute6"]; ok && val != nil {
		data.Attribute6 = types.StringValue(val.(string))
	} else {
		data.Attribute6 = types.StringNull()
	}
	if val, ok := getResponseData["attribute6expr"]; ok && val != nil {
		data.Attribute6expr = types.StringValue(val.(string))
	} else {
		data.Attribute6expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute6format"]; ok && val != nil {
		data.Attribute6format = types.StringValue(val.(string))
	} else {
		data.Attribute6format = types.StringNull()
	}
	if val, ok := getResponseData["attribute6friendlyname"]; ok && val != nil {
		data.Attribute6friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute6friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute7"]; ok && val != nil {
		data.Attribute7 = types.StringValue(val.(string))
	} else {
		data.Attribute7 = types.StringNull()
	}
	if val, ok := getResponseData["attribute7expr"]; ok && val != nil {
		data.Attribute7expr = types.StringValue(val.(string))
	} else {
		data.Attribute7expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute7format"]; ok && val != nil {
		data.Attribute7format = types.StringValue(val.(string))
	} else {
		data.Attribute7format = types.StringNull()
	}
	if val, ok := getResponseData["attribute7friendlyname"]; ok && val != nil {
		data.Attribute7friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute7friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute8"]; ok && val != nil {
		data.Attribute8 = types.StringValue(val.(string))
	} else {
		data.Attribute8 = types.StringNull()
	}
	if val, ok := getResponseData["attribute8expr"]; ok && val != nil {
		data.Attribute8expr = types.StringValue(val.(string))
	} else {
		data.Attribute8expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute8format"]; ok && val != nil {
		data.Attribute8format = types.StringValue(val.(string))
	} else {
		data.Attribute8format = types.StringNull()
	}
	if val, ok := getResponseData["attribute8friendlyname"]; ok && val != nil {
		data.Attribute8friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute8friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["attribute9"]; ok && val != nil {
		data.Attribute9 = types.StringValue(val.(string))
	} else {
		data.Attribute9 = types.StringNull()
	}
	if val, ok := getResponseData["attribute9expr"]; ok && val != nil {
		data.Attribute9expr = types.StringValue(val.(string))
	} else {
		data.Attribute9expr = types.StringNull()
	}
	if val, ok := getResponseData["attribute9format"]; ok && val != nil {
		data.Attribute9format = types.StringValue(val.(string))
	} else {
		data.Attribute9format = types.StringNull()
	}
	if val, ok := getResponseData["attribute9friendlyname"]; ok && val != nil {
		data.Attribute9friendlyname = types.StringValue(val.(string))
	} else {
		data.Attribute9friendlyname = types.StringNull()
	}
	if val, ok := getResponseData["audience"]; ok && val != nil {
		data.Audience = types.StringValue(val.(string))
	} else {
		data.Audience = types.StringNull()
	}
	if val, ok := getResponseData["digestmethod"]; ok && val != nil {
		data.Digestmethod = types.StringValue(val.(string))
	} else {
		data.Digestmethod = types.StringNull()
	}
	if val, ok := getResponseData["encryptassertion"]; ok && val != nil {
		data.Encryptassertion = types.StringValue(val.(string))
	} else {
		data.Encryptassertion = types.StringNull()
	}
	if val, ok := getResponseData["encryptionalgorithm"]; ok && val != nil {
		data.Encryptionalgorithm = types.StringValue(val.(string))
	} else {
		data.Encryptionalgorithm = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nameidexpr"]; ok && val != nil {
		data.Nameidexpr = types.StringValue(val.(string))
	} else {
		data.Nameidexpr = types.StringNull()
	}
	if val, ok := getResponseData["nameidformat"]; ok && val != nil {
		data.Nameidformat = types.StringValue(val.(string))
	} else {
		data.Nameidformat = types.StringNull()
	}
	if val, ok := getResponseData["relaystaterule"]; ok && val != nil {
		data.Relaystaterule = types.StringValue(val.(string))
	} else {
		data.Relaystaterule = types.StringNull()
	}
	if val, ok := getResponseData["samlissuername"]; ok && val != nil {
		data.Samlissuername = types.StringValue(val.(string))
	} else {
		data.Samlissuername = types.StringNull()
	}
	if val, ok := getResponseData["samlsigningcertname"]; ok && val != nil {
		data.Samlsigningcertname = types.StringValue(val.(string))
	} else {
		data.Samlsigningcertname = types.StringNull()
	}
	if val, ok := getResponseData["samlspcertname"]; ok && val != nil {
		data.Samlspcertname = types.StringValue(val.(string))
	} else {
		data.Samlspcertname = types.StringNull()
	}
	if val, ok := getResponseData["sendpassword"]; ok && val != nil {
		data.Sendpassword = types.StringValue(val.(string))
	} else {
		data.Sendpassword = types.StringNull()
	}
	if val, ok := getResponseData["signassertion"]; ok && val != nil {
		data.Signassertion = types.StringValue(val.(string))
	} else {
		data.Signassertion = types.StringNull()
	}
	if val, ok := getResponseData["signaturealg"]; ok && val != nil {
		data.Signaturealg = types.StringValue(val.(string))
	} else {
		data.Signaturealg = types.StringNull()
	}
	if val, ok := getResponseData["signatureservice"]; ok && val != nil {
		data.Signatureservice = types.StringValue(val.(string))
	} else {
		data.Signatureservice = types.StringNull()
	}
	if val, ok := getResponseData["skewtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Skewtime = types.Int64Value(intVal)
		}
	} else {
		data.Skewtime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
