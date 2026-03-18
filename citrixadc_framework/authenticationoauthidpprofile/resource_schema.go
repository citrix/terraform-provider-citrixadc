package authenticationoauthidpprofile

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

// AuthenticationoauthidpprofileResourceModel describes the resource data model.
type AuthenticationoauthidpprofileResourceModel struct {
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
}

func (r *AuthenticationoauthidpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationoauthidpprofile resource.",
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
				Default:     int64default.StaticInt64(50),
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
				Default:     stringdefault.StaticString("RS256"),
				Description: "Algorithm to be used to sign OpenID tokens.",
			},
			"signatureservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service in cloud used to sign the data. This is applicable only if signature if offloaded to cloud.",
			},
			"skewtime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "This option specifies the duration for which the token sent by Citrix ADC IdP is valid. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.",
			},
		},
	}
}

func authenticationoauthidpprofileGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationoauthidpprofileResourceModel) authentication.Authenticationoauthidpprofile {
	tflog.Debug(ctx, "In authenticationoauthidpprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationoauthidpprofile := authentication.Authenticationoauthidpprofile{}
	if !data.Attributes.IsNull() {
		authenticationoauthidpprofile.Attributes = data.Attributes.ValueString()
	}
	if !data.Audience.IsNull() {
		authenticationoauthidpprofile.Audience = data.Audience.ValueString()
	}
	if !data.Clientid.IsNull() {
		authenticationoauthidpprofile.Clientid = data.Clientid.ValueString()
	}
	if !data.Clientsecret.IsNull() {
		authenticationoauthidpprofile.Clientsecret = data.Clientsecret.ValueString()
	}
	if !data.Configservice.IsNull() {
		authenticationoauthidpprofile.Configservice = data.Configservice.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationoauthidpprofile.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Encrypttoken.IsNull() {
		authenticationoauthidpprofile.Encrypttoken = data.Encrypttoken.ValueString()
	}
	if !data.Issuer.IsNull() {
		authenticationoauthidpprofile.Issuer = data.Issuer.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationoauthidpprofile.Name = data.Name.ValueString()
	}
	if !data.Redirecturl.IsNull() {
		authenticationoauthidpprofile.Redirecturl = data.Redirecturl.ValueString()
	}
	if !data.Refreshinterval.IsNull() {
		authenticationoauthidpprofile.Refreshinterval = utils.IntPtr(int(data.Refreshinterval.ValueInt64()))
	}
	if !data.Relyingpartymetadataurl.IsNull() {
		authenticationoauthidpprofile.Relyingpartymetadataurl = data.Relyingpartymetadataurl.ValueString()
	}
	if !data.Sendpassword.IsNull() {
		authenticationoauthidpprofile.Sendpassword = data.Sendpassword.ValueString()
	}
	if !data.Signaturealg.IsNull() {
		authenticationoauthidpprofile.Signaturealg = data.Signaturealg.ValueString()
	}
	if !data.Signatureservice.IsNull() {
		authenticationoauthidpprofile.Signatureservice = data.Signatureservice.ValueString()
	}
	if !data.Skewtime.IsNull() {
		authenticationoauthidpprofile.Skewtime = utils.IntPtr(int(data.Skewtime.ValueInt64()))
	}

	return authenticationoauthidpprofile
}

func authenticationoauthidpprofileSetAttrFromGet(ctx context.Context, data *AuthenticationoauthidpprofileResourceModel, getResponseData map[string]interface{}) *AuthenticationoauthidpprofileResourceModel {
	tflog.Debug(ctx, "In authenticationoauthidpprofileSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["configservice"]; ok && val != nil {
		data.Configservice = types.StringValue(val.(string))
	} else {
		data.Configservice = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["encrypttoken"]; ok && val != nil {
		data.Encrypttoken = types.StringValue(val.(string))
	} else {
		data.Encrypttoken = types.StringNull()
	}
	if val, ok := getResponseData["issuer"]; ok && val != nil {
		data.Issuer = types.StringValue(val.(string))
	} else {
		data.Issuer = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["redirecturl"]; ok && val != nil {
		data.Redirecturl = types.StringValue(val.(string))
	} else {
		data.Redirecturl = types.StringNull()
	}
	if val, ok := getResponseData["refreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Refreshinterval = types.Int64Null()
	}
	if val, ok := getResponseData["relyingpartymetadataurl"]; ok && val != nil {
		data.Relyingpartymetadataurl = types.StringValue(val.(string))
	} else {
		data.Relyingpartymetadataurl = types.StringNull()
	}
	if val, ok := getResponseData["sendpassword"]; ok && val != nil {
		data.Sendpassword = types.StringValue(val.(string))
	} else {
		data.Sendpassword = types.StringNull()
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
