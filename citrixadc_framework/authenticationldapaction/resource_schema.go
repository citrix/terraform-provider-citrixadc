package authenticationldapaction

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

// AuthenticationldapactionResourceModel describes the resource data model.
type AuthenticationldapactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Alternateemailattr         types.String `tfsdk:"alternateemailattr"`
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
	Authentication             types.String `tfsdk:"authentication"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Cloudattributes            types.String `tfsdk:"cloudattributes"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Email                      types.String `tfsdk:"email"`
	Followreferrals            types.String `tfsdk:"followreferrals"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Groupnameidentifier        types.String `tfsdk:"groupnameidentifier"`
	Groupsearchattribute       types.String `tfsdk:"groupsearchattribute"`
	Groupsearchfilter          types.String `tfsdk:"groupsearchfilter"`
	Groupsearchsubattribute    types.String `tfsdk:"groupsearchsubattribute"`
	Kbattribute                types.String `tfsdk:"kbattribute"`
	Ldapbase                   types.String `tfsdk:"ldapbase"`
	Ldapbinddn                 types.String `tfsdk:"ldapbinddn"`
	Ldapbinddnpassword         types.String `tfsdk:"ldapbinddnpassword"`
	Ldaphostname               types.String `tfsdk:"ldaphostname"`
	Ldaploginname              types.String `tfsdk:"ldaploginname"`
	Maxldapreferrals           types.Int64  `tfsdk:"maxldapreferrals"`
	Maxnestinglevel            types.Int64  `tfsdk:"maxnestinglevel"`
	Mssrvrecordlocation        types.String `tfsdk:"mssrvrecordlocation"`
	Name                       types.String `tfsdk:"name"`
	Nestedgroupextraction      types.String `tfsdk:"nestedgroupextraction"`
	Otpsecret                  types.String `tfsdk:"otpsecret"`
	Passwdchange               types.String `tfsdk:"passwdchange"`
	Pushservice                types.String `tfsdk:"pushservice"`
	Referraldnslookup          types.String `tfsdk:"referraldnslookup"`
	Requireuser                types.String `tfsdk:"requireuser"`
	Searchfilter               types.String `tfsdk:"searchfilter"`
	Sectype                    types.String `tfsdk:"sectype"`
	Serverip                   types.String `tfsdk:"serverip"`
	Servername                 types.String `tfsdk:"servername"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Sshpublickey               types.String `tfsdk:"sshpublickey"`
	Ssonameattribute           types.String `tfsdk:"ssonameattribute"`
	Subattributename           types.String `tfsdk:"subattributename"`
	Svrtype                    types.String `tfsdk:"svrtype"`
	Validateservercert         types.String `tfsdk:"validateservercert"`
}

func (r *AuthenticationldapactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationldapaction resource.",
			},
			"alternateemailattr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The NetScaler appliance uses the alternateive email attribute to query the Active Directory for the alternative email id of a user",
			},
			"attribute1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute1 from the ldap response",
			},
			"attribute10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute10 from the ldap response",
			},
			"attribute11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute11 from the ldap response",
			},
			"attribute12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute12 from the ldap response",
			},
			"attribute13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute13 from the ldap response",
			},
			"attribute14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute14 from the ldap response",
			},
			"attribute15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute15 from the ldap response",
			},
			"attribute16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute16 from the ldap response",
			},
			"attribute2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute2 from the ldap response",
			},
			"attribute3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute3 from the ldap response",
			},
			"attribute4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute4 from the ldap response",
			},
			"attribute5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute5 from the ldap response",
			},
			"attribute6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute6 from the ldap response",
			},
			"attribute7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute7 from the ldap response",
			},
			"attribute8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute8 from the ldap response",
			},
			"attribute9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that would be evaluated to extract attribute9 from the ldap response",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of attribute names separated by ',' which needs to be fetched from ldap server.\nNote that preceeding and trailing spaces will be removed.\nAttribute name can be 127 bytes and total length of this string should not cross 2047 bytes.\nThese attributes have multi-value support separated by ',' and stored as key-value pair in AAA session",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Perform LDAP authentication.\nIf authentication is disabled, any LDAP authentication attempt returns authentication success if the user is found.\nCAUTION! Authentication should be disabled only for authorization group extraction or where other (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Number of seconds the Citrix ADC waits for a response from the RADIUS server.",
			},
			"cloudattributes": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "The Citrix ADC uses the cloud attributes to extract additional attributes from LDAP servers required for Citrix Cloud operations",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"email": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("mail"),
				Description: "The Citrix ADC uses the email attribute to query the Active Directory for the email id of a user",
			},
			"followreferrals": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this option to ON enables following LDAP referrals received from the LDAP server.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP group attribute name.\nUsed for group extraction on the LDAP server.",
			},
			"groupnameidentifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name that uniquely identifies a group in LDAP or Active Directory.",
			},
			"groupsearchattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP group search attribute.\nUsed to determine to which groups a group belongs.",
			},
			"groupsearchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to be combined with the default LDAP group search string to form the search value.  For example, the group search filter \"\"vpnallowed=true\"\" when combined with the group identifier \"\"samaccount\"\" and the group name \"\"g1\"\" yields the LDAP search string \"\"(&(vpnallowed=true)(samaccount=g1)\"\". If nestedGroupExtraction is ENABLED, the filter is applied on the first level group search as well, otherwise first level groups (of which user is a direct member of) will be fetched without applying this filter. (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.)",
			},
			"groupsearchsubattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP group search subattribute.\nUsed to determine to which groups a group belongs.",
			},
			"kbattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "KnowledgeBasedAuthentication(KBA) attribute on AD. This attribute is used to store and retrieve preconfigured Question and Answer knowledge base used for KBA authentication.",
			},
			"ldapbase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base (node) from which to start LDAP searches.\nIf the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.",
			},
			"ldapbinddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Full distinguished name (DN) that is used to bind to the LDAP server.\nDefault: cn=Manager,dc=netscaler,dc=com",
			},
			"ldapbinddnpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password used to bind to the LDAP server.",
			},
			"ldaphostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname for the LDAP server.  If -validateServerCert is ON then this must be the host name on the certificate from the LDAP server.\nA hostname mismatch will cause a connection failure.",
			},
			"ldaploginname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP login name attribute.\nThe Citrix ADC uses the LDAP login name to query external LDAP servers or Active Directories.",
			},
			"maxldapreferrals": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Specifies the maximum number of nested referrals to follow.",
			},
			"maxnestinglevel": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "If nested group extraction is ON, specifies the number of levels up to which group extraction is performed.",
			},
			"mssrvrecordlocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MSSRV Specific parameter. Used to locate the DNS node to which the SRV record pertains in the domainname. The domainname is appended to it to form the srv record.\nExample : For \"dc._msdcs\", the srv record formed is _ldap._tcp.dc._msdcs.<domainname>.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new LDAP action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the LDAP action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"nestedgroupextraction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow nested group extraction, in which the Citrix ADC queries external LDAP servers to determine whether a group is part of another group.",
			},
			"otpsecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OneTimePassword(OTP) Secret key attribute on AD. This attribute is used to store and retrieve secret key used for OTP check",
			},
			"passwdchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow password change requests.",
			},
			"pushservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service used to send push notifications",
			},
			"referraldnslookup": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("A-REC"),
				Description: "Specifies the DNS Record lookup Type for the referrals",
			},
			"requireuser": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Require a successful user search for authentication.\nCAUTION!  This field should be set to NO only if usersearch not required [Both username validation as well as password validation skipped] and (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to be combined with the default LDAP user search string to form the search value. For example, if the search filter \"vpnallowed=true\" is combined with the LDAP login name \"samaccount\" and the user-supplied username is \"bob\", the result is the LDAP search string \"\"&(vpnallowed=true)(samaccount=bob)\"\" (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.).",
			},
			"sectype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PLAINTEXT"),
				Description: "Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address assigned to the LDAP server.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP server name as a FQDN.  Mutually exclusive with LDAP IP address.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(389),
				Description: "Port on which the LDAP server accepts connections.",
			},
			"sshpublickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSH PublicKey is attribute on AD. This attribute is used to retrieve ssh PublicKey for RBA authentication",
			},
			"ssonameattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP single signon (SSO) attribute.\nThe Citrix ADC uses the SSO name attribute to query external LDAP servers or Active Directories for an alternate username.",
			},
			"subattributename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP group sub-attribute name.\nUsed for group extraction from the LDAP server.",
			},
			"svrtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AAA_LDAP_SERVER_TYPE_DEFAULT"),
				Description: "The type of LDAP server.",
			},
			"validateservercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When to validate LDAP server certs",
			},
		},
	}
}

func authenticationldapactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationldapactionResourceModel) authentication.Authenticationldapaction {
	tflog.Debug(ctx, "In authenticationldapactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationldapaction := authentication.Authenticationldapaction{}
	if !data.Alternateemailattr.IsNull() {
		authenticationldapaction.Alternateemailattr = data.Alternateemailattr.ValueString()
	}
	if !data.Attribute1.IsNull() {
		authenticationldapaction.Attribute1 = data.Attribute1.ValueString()
	}
	if !data.Attribute10.IsNull() {
		authenticationldapaction.Attribute10 = data.Attribute10.ValueString()
	}
	if !data.Attribute11.IsNull() {
		authenticationldapaction.Attribute11 = data.Attribute11.ValueString()
	}
	if !data.Attribute12.IsNull() {
		authenticationldapaction.Attribute12 = data.Attribute12.ValueString()
	}
	if !data.Attribute13.IsNull() {
		authenticationldapaction.Attribute13 = data.Attribute13.ValueString()
	}
	if !data.Attribute14.IsNull() {
		authenticationldapaction.Attribute14 = data.Attribute14.ValueString()
	}
	if !data.Attribute15.IsNull() {
		authenticationldapaction.Attribute15 = data.Attribute15.ValueString()
	}
	if !data.Attribute16.IsNull() {
		authenticationldapaction.Attribute16 = data.Attribute16.ValueString()
	}
	if !data.Attribute2.IsNull() {
		authenticationldapaction.Attribute2 = data.Attribute2.ValueString()
	}
	if !data.Attribute3.IsNull() {
		authenticationldapaction.Attribute3 = data.Attribute3.ValueString()
	}
	if !data.Attribute4.IsNull() {
		authenticationldapaction.Attribute4 = data.Attribute4.ValueString()
	}
	if !data.Attribute5.IsNull() {
		authenticationldapaction.Attribute5 = data.Attribute5.ValueString()
	}
	if !data.Attribute6.IsNull() {
		authenticationldapaction.Attribute6 = data.Attribute6.ValueString()
	}
	if !data.Attribute7.IsNull() {
		authenticationldapaction.Attribute7 = data.Attribute7.ValueString()
	}
	if !data.Attribute8.IsNull() {
		authenticationldapaction.Attribute8 = data.Attribute8.ValueString()
	}
	if !data.Attribute9.IsNull() {
		authenticationldapaction.Attribute9 = data.Attribute9.ValueString()
	}
	if !data.Attributes.IsNull() {
		authenticationldapaction.Attributes = data.Attributes.ValueString()
	}
	if !data.Authentication.IsNull() {
		authenticationldapaction.Authentication = data.Authentication.ValueString()
	}
	if !data.Authtimeout.IsNull() {
		authenticationldapaction.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Cloudattributes.IsNull() {
		authenticationldapaction.Cloudattributes = data.Cloudattributes.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationldapaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Email.IsNull() {
		authenticationldapaction.Email = data.Email.ValueString()
	}
	if !data.Followreferrals.IsNull() {
		authenticationldapaction.Followreferrals = data.Followreferrals.ValueString()
	}
	if !data.Groupattrname.IsNull() {
		authenticationldapaction.Groupattrname = data.Groupattrname.ValueString()
	}
	if !data.Groupnameidentifier.IsNull() {
		authenticationldapaction.Groupnameidentifier = data.Groupnameidentifier.ValueString()
	}
	if !data.Groupsearchattribute.IsNull() {
		authenticationldapaction.Groupsearchattribute = data.Groupsearchattribute.ValueString()
	}
	if !data.Groupsearchfilter.IsNull() {
		authenticationldapaction.Groupsearchfilter = data.Groupsearchfilter.ValueString()
	}
	if !data.Groupsearchsubattribute.IsNull() {
		authenticationldapaction.Groupsearchsubattribute = data.Groupsearchsubattribute.ValueString()
	}
	if !data.Kbattribute.IsNull() {
		authenticationldapaction.Kbattribute = data.Kbattribute.ValueString()
	}
	if !data.Ldapbase.IsNull() {
		authenticationldapaction.Ldapbase = data.Ldapbase.ValueString()
	}
	if !data.Ldapbinddn.IsNull() {
		authenticationldapaction.Ldapbinddn = data.Ldapbinddn.ValueString()
	}
	if !data.Ldapbinddnpassword.IsNull() {
		authenticationldapaction.Ldapbinddnpassword = data.Ldapbinddnpassword.ValueString()
	}
	if !data.Ldaphostname.IsNull() {
		authenticationldapaction.Ldaphostname = data.Ldaphostname.ValueString()
	}
	if !data.Ldaploginname.IsNull() {
		authenticationldapaction.Ldaploginname = data.Ldaploginname.ValueString()
	}
	if !data.Maxldapreferrals.IsNull() {
		authenticationldapaction.Maxldapreferrals = utils.IntPtr(int(data.Maxldapreferrals.ValueInt64()))
	}
	if !data.Maxnestinglevel.IsNull() {
		authenticationldapaction.Maxnestinglevel = utils.IntPtr(int(data.Maxnestinglevel.ValueInt64()))
	}
	if !data.Mssrvrecordlocation.IsNull() {
		authenticationldapaction.Mssrvrecordlocation = data.Mssrvrecordlocation.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationldapaction.Name = data.Name.ValueString()
	}
	if !data.Nestedgroupextraction.IsNull() {
		authenticationldapaction.Nestedgroupextraction = data.Nestedgroupextraction.ValueString()
	}
	if !data.Otpsecret.IsNull() {
		authenticationldapaction.Otpsecret = data.Otpsecret.ValueString()
	}
	if !data.Passwdchange.IsNull() {
		authenticationldapaction.Passwdchange = data.Passwdchange.ValueString()
	}
	if !data.Pushservice.IsNull() {
		authenticationldapaction.Pushservice = data.Pushservice.ValueString()
	}
	if !data.Referraldnslookup.IsNull() {
		authenticationldapaction.Referraldnslookup = data.Referraldnslookup.ValueString()
	}
	if !data.Requireuser.IsNull() {
		authenticationldapaction.Requireuser = data.Requireuser.ValueString()
	}
	if !data.Searchfilter.IsNull() {
		authenticationldapaction.Searchfilter = data.Searchfilter.ValueString()
	}
	if !data.Sectype.IsNull() {
		authenticationldapaction.Sectype = data.Sectype.ValueString()
	}
	if !data.Serverip.IsNull() {
		authenticationldapaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() {
		authenticationldapaction.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() {
		authenticationldapaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Sshpublickey.IsNull() {
		authenticationldapaction.Sshpublickey = data.Sshpublickey.ValueString()
	}
	if !data.Ssonameattribute.IsNull() {
		authenticationldapaction.Ssonameattribute = data.Ssonameattribute.ValueString()
	}
	if !data.Subattributename.IsNull() {
		authenticationldapaction.Subattributename = data.Subattributename.ValueString()
	}
	if !data.Svrtype.IsNull() {
		authenticationldapaction.Svrtype = data.Svrtype.ValueString()
	}
	if !data.Validateservercert.IsNull() {
		authenticationldapaction.Validateservercert = data.Validateservercert.ValueString()
	}

	return authenticationldapaction
}

func authenticationldapactionSetAttrFromGet(ctx context.Context, data *AuthenticationldapactionResourceModel, getResponseData map[string]interface{}) *AuthenticationldapactionResourceModel {
	tflog.Debug(ctx, "In authenticationldapactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alternateemailattr"]; ok && val != nil {
		data.Alternateemailattr = types.StringValue(val.(string))
	} else {
		data.Alternateemailattr = types.StringNull()
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
	if val, ok := getResponseData["attributes"]; ok && val != nil {
		data.Attributes = types.StringValue(val.(string))
	} else {
		data.Attributes = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Authtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["cloudattributes"]; ok && val != nil {
		data.Cloudattributes = types.StringValue(val.(string))
	} else {
		data.Cloudattributes = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["email"]; ok && val != nil {
		data.Email = types.StringValue(val.(string))
	} else {
		data.Email = types.StringNull()
	}
	if val, ok := getResponseData["followreferrals"]; ok && val != nil {
		data.Followreferrals = types.StringValue(val.(string))
	} else {
		data.Followreferrals = types.StringNull()
	}
	if val, ok := getResponseData["groupattrname"]; ok && val != nil {
		data.Groupattrname = types.StringValue(val.(string))
	} else {
		data.Groupattrname = types.StringNull()
	}
	if val, ok := getResponseData["groupnameidentifier"]; ok && val != nil {
		data.Groupnameidentifier = types.StringValue(val.(string))
	} else {
		data.Groupnameidentifier = types.StringNull()
	}
	if val, ok := getResponseData["groupsearchattribute"]; ok && val != nil {
		data.Groupsearchattribute = types.StringValue(val.(string))
	} else {
		data.Groupsearchattribute = types.StringNull()
	}
	if val, ok := getResponseData["groupsearchfilter"]; ok && val != nil {
		data.Groupsearchfilter = types.StringValue(val.(string))
	} else {
		data.Groupsearchfilter = types.StringNull()
	}
	if val, ok := getResponseData["groupsearchsubattribute"]; ok && val != nil {
		data.Groupsearchsubattribute = types.StringValue(val.(string))
	} else {
		data.Groupsearchsubattribute = types.StringNull()
	}
	if val, ok := getResponseData["kbattribute"]; ok && val != nil {
		data.Kbattribute = types.StringValue(val.(string))
	} else {
		data.Kbattribute = types.StringNull()
	}
	if val, ok := getResponseData["ldapbase"]; ok && val != nil {
		data.Ldapbase = types.StringValue(val.(string))
	} else {
		data.Ldapbase = types.StringNull()
	}
	if val, ok := getResponseData["ldapbinddn"]; ok && val != nil {
		data.Ldapbinddn = types.StringValue(val.(string))
	} else {
		data.Ldapbinddn = types.StringNull()
	}
	if val, ok := getResponseData["ldapbinddnpassword"]; ok && val != nil {
		data.Ldapbinddnpassword = types.StringValue(val.(string))
	} else {
		data.Ldapbinddnpassword = types.StringNull()
	}
	if val, ok := getResponseData["ldaphostname"]; ok && val != nil {
		data.Ldaphostname = types.StringValue(val.(string))
	} else {
		data.Ldaphostname = types.StringNull()
	}
	if val, ok := getResponseData["ldaploginname"]; ok && val != nil {
		data.Ldaploginname = types.StringValue(val.(string))
	} else {
		data.Ldaploginname = types.StringNull()
	}
	if val, ok := getResponseData["maxldapreferrals"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxldapreferrals = types.Int64Value(intVal)
		}
	} else {
		data.Maxldapreferrals = types.Int64Null()
	}
	if val, ok := getResponseData["maxnestinglevel"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxnestinglevel = types.Int64Value(intVal)
		}
	} else {
		data.Maxnestinglevel = types.Int64Null()
	}
	if val, ok := getResponseData["mssrvrecordlocation"]; ok && val != nil {
		data.Mssrvrecordlocation = types.StringValue(val.(string))
	} else {
		data.Mssrvrecordlocation = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nestedgroupextraction"]; ok && val != nil {
		data.Nestedgroupextraction = types.StringValue(val.(string))
	} else {
		data.Nestedgroupextraction = types.StringNull()
	}
	if val, ok := getResponseData["otpsecret"]; ok && val != nil {
		data.Otpsecret = types.StringValue(val.(string))
	} else {
		data.Otpsecret = types.StringNull()
	}
	if val, ok := getResponseData["passwdchange"]; ok && val != nil {
		data.Passwdchange = types.StringValue(val.(string))
	} else {
		data.Passwdchange = types.StringNull()
	}
	if val, ok := getResponseData["pushservice"]; ok && val != nil {
		data.Pushservice = types.StringValue(val.(string))
	} else {
		data.Pushservice = types.StringNull()
	}
	if val, ok := getResponseData["referraldnslookup"]; ok && val != nil {
		data.Referraldnslookup = types.StringValue(val.(string))
	} else {
		data.Referraldnslookup = types.StringNull()
	}
	if val, ok := getResponseData["requireuser"]; ok && val != nil {
		data.Requireuser = types.StringValue(val.(string))
	} else {
		data.Requireuser = types.StringNull()
	}
	if val, ok := getResponseData["searchfilter"]; ok && val != nil {
		data.Searchfilter = types.StringValue(val.(string))
	} else {
		data.Searchfilter = types.StringNull()
	}
	if val, ok := getResponseData["sectype"]; ok && val != nil {
		data.Sectype = types.StringValue(val.(string))
	} else {
		data.Sectype = types.StringNull()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["sshpublickey"]; ok && val != nil {
		data.Sshpublickey = types.StringValue(val.(string))
	} else {
		data.Sshpublickey = types.StringNull()
	}
	if val, ok := getResponseData["ssonameattribute"]; ok && val != nil {
		data.Ssonameattribute = types.StringValue(val.(string))
	} else {
		data.Ssonameattribute = types.StringNull()
	}
	if val, ok := getResponseData["subattributename"]; ok && val != nil {
		data.Subattributename = types.StringValue(val.(string))
	} else {
		data.Subattributename = types.StringNull()
	}
	if val, ok := getResponseData["svrtype"]; ok && val != nil {
		data.Svrtype = types.StringValue(val.(string))
	} else {
		data.Svrtype = types.StringNull()
	}
	if val, ok := getResponseData["validateservercert"]; ok && val != nil {
		data.Validateservercert = types.StringValue(val.(string))
	} else {
		data.Validateservercert = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
