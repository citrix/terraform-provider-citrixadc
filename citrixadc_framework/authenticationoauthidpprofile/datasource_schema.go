package authenticationoauthidpprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationoauthidpprofileDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationoauthidpprofileResourceModel. A data source is a
// pure read surface (Read only; no plan/apply lifecycle), so it can expose the
// FULL GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits. Every non-key attribute
// is Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares.
type AuthenticationoauthidpprofileDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Attributes                 types.String `tfsdk:"attributes"`
	Audience                   types.String `tfsdk:"audience"`
	Clientid                   types.String `tfsdk:"clientid"`
	Clientsecret               types.String `tfsdk:"clientsecret"`
	Configservice              types.String `tfsdk:"configservice"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Encrypttoken               types.String `tfsdk:"encrypttoken"`
	Issuer                     types.String `tfsdk:"issuer"`
	Name                       types.String `tfsdk:"name"`
	Redirecturl                types.String `tfsdk:"redirecturl"`
	Refreshinterval            types.Int64  `tfsdk:"refreshinterval"`
	Relyingpartymetadataurl    types.String `tfsdk:"relyingpartymetadataurl"`
	Sendpassword               types.String `tfsdk:"sendpassword"`
	Signaturealg               types.String `tfsdk:"signaturealg"`
	Signatureservice           types.String `tfsdk:"signatureservice"`
	Skewtime                   types.Int64  `tfsdk:"skewtime"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationoauthidpprofile.json). Never settable;
	// populated from GET.
	Oauthstatus types.String `tfsdk:"oauthstatus"`
}

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
				Sensitive:   true,
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

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"oauthstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Describes status information of oauth idp metadata fetch process. Possible values = INIT, CERTFETCH, AADFORGRAPH, GRAPH, AADFORMDM, MDMINFO, COMPLETE",
			},
		},
	}
}

// authenticationoauthidpprofileDataSourceSetAttrFromGet projects a NITRO
// authenticationoauthidpprofile GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationoauthidpprofileDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationoauthidpprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationoauthidpprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Attributes = utils.MapGetString(g, "attributes")
	data.Audience = utils.MapGetString(g, "audience")
	data.Clientid = utils.MapGetString(g, "clientid")
	data.Configservice = utils.MapGetString(g, "configservice")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Encrypttoken = utils.MapGetString(g, "encrypttoken")
	data.Issuer = utils.MapGetString(g, "issuer")
	data.Redirecturl = utils.MapGetString(g, "redirecturl")
	data.Refreshinterval = utils.MapGetInt64(g, "refreshinterval")
	data.Relyingpartymetadataurl = utils.MapGetString(g, "relyingpartymetadataurl")
	data.Sendpassword = utils.MapGetString(g, "sendpassword")
	data.Signaturealg = utils.MapGetString(g, "signaturealg")
	data.Signatureservice = utils.MapGetString(g, "signatureservice")
	data.Skewtime = utils.MapGetInt64(g, "skewtime")

	// clientsecret is a secret input the GET never returns -> Null.
	data.Clientsecret = types.StringNull()

	// Read-only attributes.
	data.Oauthstatus = utils.MapGetString(g, "oauthstatus")
}
