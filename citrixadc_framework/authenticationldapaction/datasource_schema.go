package authenticationldapaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationldapactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationldapactionResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes the resource
// deliberately omits (ldapcontimeout, success, failure). The Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AuthenticationldapactionDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Alternateemailattr          types.String `tfsdk:"alternateemailattr"`
	Attribute1                  types.String `tfsdk:"attribute1"`
	Attribute10                 types.String `tfsdk:"attribute10"`
	Attribute11                 types.String `tfsdk:"attribute11"`
	Attribute12                 types.String `tfsdk:"attribute12"`
	Attribute13                 types.String `tfsdk:"attribute13"`
	Attribute14                 types.String `tfsdk:"attribute14"`
	Attribute15                 types.String `tfsdk:"attribute15"`
	Attribute16                 types.String `tfsdk:"attribute16"`
	Attribute2                  types.String `tfsdk:"attribute2"`
	Attribute3                  types.String `tfsdk:"attribute3"`
	Attribute4                  types.String `tfsdk:"attribute4"`
	Attribute5                  types.String `tfsdk:"attribute5"`
	Attribute6                  types.String `tfsdk:"attribute6"`
	Attribute7                  types.String `tfsdk:"attribute7"`
	Attribute8                  types.String `tfsdk:"attribute8"`
	Attribute9                  types.String `tfsdk:"attribute9"`
	Attributes                  types.String `tfsdk:"attributes"`
	Authentication              types.String `tfsdk:"authentication"`
	Authtimeout                 types.Int64  `tfsdk:"authtimeout"`
	Cloudattributes             types.String `tfsdk:"cloudattributes"`
	Defaultauthenticationgroup  types.String `tfsdk:"defaultauthenticationgroup"`
	Email                       types.String `tfsdk:"email"`
	Followreferrals             types.String `tfsdk:"followreferrals"`
	Groupattrname               types.String `tfsdk:"groupattrname"`
	Groupnameidentifier         types.String `tfsdk:"groupnameidentifier"`
	Groupsearchattribute        types.String `tfsdk:"groupsearchattribute"`
	Groupsearchfilter           types.String `tfsdk:"groupsearchfilter"`
	Groupsearchsubattribute     types.String `tfsdk:"groupsearchsubattribute"`
	Kbattribute                 types.String `tfsdk:"kbattribute"`
	Ldapbase                    types.String `tfsdk:"ldapbase"`
	Ldapbinddn                  types.String `tfsdk:"ldapbinddn"`
	Ldapbinddnpassword          types.String `tfsdk:"ldapbinddnpassword"`
	LdapbinddnpasswordWo        types.String `tfsdk:"ldapbinddnpassword_wo"`
	LdapbinddnpasswordWoVersion types.Int64  `tfsdk:"ldapbinddnpassword_wo_version"`
	Ldaphostname                types.String `tfsdk:"ldaphostname"`
	Ldaploginname               types.String `tfsdk:"ldaploginname"`
	Maxldapreferrals            types.Int64  `tfsdk:"maxldapreferrals"`
	Maxnestinglevel             types.Int64  `tfsdk:"maxnestinglevel"`
	Mssrvrecordlocation         types.String `tfsdk:"mssrvrecordlocation"`
	Name                        types.String `tfsdk:"name"`
	Nestedgroupextraction       types.String `tfsdk:"nestedgroupextraction"`
	Otpsecret                   types.String `tfsdk:"otpsecret"`
	Passwdchange                types.String `tfsdk:"passwdchange"`
	Passwordlessmgmtaccess      types.String `tfsdk:"passwordlessmgmtaccess"`
	Pushservice                 types.String `tfsdk:"pushservice"`
	Referraldnslookup           types.String `tfsdk:"referraldnslookup"`
	Requireuser                 types.String `tfsdk:"requireuser"`
	Searchfilter                types.String `tfsdk:"searchfilter"`
	Sectype                     types.String `tfsdk:"sectype"`
	Serverip                    types.String `tfsdk:"serverip"`
	Servername                  types.String `tfsdk:"servername"`
	Serverport                  types.Int64  `tfsdk:"serverport"`
	Sshpublickey                types.String `tfsdk:"sshpublickey"`
	Ssonameattribute            types.String `tfsdk:"ssonameattribute"`
	Subattributename            types.String `tfsdk:"subattributename"`
	Svrtype                     types.String `tfsdk:"svrtype"`
	Validateservercert          types.String `tfsdk:"validateservercert"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationldapaction.json). Never settable.
	Ldapcontimeout types.Int64 `tfsdk:"ldapcontimeout"`
	Success        types.Int64 `tfsdk:"success"`
	Failure        types.Int64 `tfsdk:"failure"`
}

func AuthenticationldapactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Computed:    true,
				Description: "Perform LDAP authentication.\nIf authentication is disabled, any LDAP authentication attempt returns authentication success if the user is found.\nCAUTION! Authentication should be disabled only for authorization group extraction or where other (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds the Citrix ADC waits for a response from the RADIUS server.",
			},
			"cloudattributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Citrix ADC uses the cloud attributes to extract additional attributes from LDAP servers required for Citrix Cloud operations",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"email": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
			"ldapbinddnpassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password used to bind to the LDAP server.",
			},
			"ldapbinddnpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a ldapbinddnpassword_wo update.",
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
				Computed:    true,
				Description: "Specifies the maximum number of nested referrals to follow.",
			},
			"maxnestinglevel": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Allow password change requests.",
			},
			"passwordlessmgmtaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This feature configures NetScaler management access to use LDAP exclusively for retrieving user group information. It ensures that LDAP is not used for authenticating user logins (i.e., verifying passwords) for NetScaler management access.",
			},
			"pushservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service used to send push notifications",
			},
			"referraldnslookup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the DNS Record lookup Type for the referrals",
			},
			"requireuser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Require a successful user search for authentication.\nCAUTION!  This field should be set to NO only if usersearch not required [Both username validation as well as password validation skipped] and (non-LDAP) authentication methods are in use and either bound to a primary list or flagged as secondary.",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to be combined with the default LDAP user search string to form the search value. For example, if the search filter \"vpnallowed=true\" is combined with the LDAP login name \"samaccount\" and the user-supplied username is \"bob\", the result is the LDAP search string \"\"&(vpnallowed=true)(samaccount=bob)\"\" (Be sure to enclose the search string in two sets of double quotation marks; both sets are needed.).",
			},
			"sectype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
				Computed:    true,
				Description: "The type of LDAP server.",
			},
			"validateservercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When to validate LDAP server certs",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"ldapcontimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of seconds the Citrix ADC waits for the LDAP connection (TCP connection and SSL/TLS handshake) to be established with the LDAP server.",
			},
			"success": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successful authentications through this LDAP action.",
			},
			"failure": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of failed authentications through this LDAP action.",
			},
		},
	}
}

// authenticationldapactionDataSourceSetAttrFromGet projects a NITRO
// authenticationldapaction GET response onto the data-source model. Attributes
// are simply filled from the GET (or left Null when the GET omits them) via the
// shared utils.MapGet* helpers.
func authenticationldapactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationldapactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationldapactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Alternateemailattr = utils.MapGetString(g, "alternateemailattr")
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
	data.Attributes = utils.MapGetString(g, "attributes")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authtimeout = utils.MapGetInt64(g, "authtimeout")
	data.Cloudattributes = utils.MapGetString(g, "cloudattributes")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Email = utils.MapGetString(g, "email")
	data.Followreferrals = utils.MapGetString(g, "followreferrals")
	data.Groupattrname = utils.MapGetString(g, "groupattrname")
	data.Groupnameidentifier = utils.MapGetString(g, "groupnameidentifier")
	data.Groupsearchattribute = utils.MapGetString(g, "groupsearchattribute")
	data.Groupsearchfilter = utils.MapGetString(g, "groupsearchfilter")
	data.Groupsearchsubattribute = utils.MapGetString(g, "groupsearchsubattribute")
	data.Kbattribute = utils.MapGetString(g, "kbattribute")
	data.Ldapbase = utils.MapGetString(g, "ldapbase")
	data.Ldapbinddn = utils.MapGetString(g, "ldapbinddn")
	data.Ldaphostname = utils.MapGetString(g, "ldaphostname")
	data.Ldaploginname = utils.MapGetString(g, "ldaploginname")
	data.Maxldapreferrals = utils.MapGetInt64(g, "maxldapreferrals")
	data.Maxnestinglevel = utils.MapGetInt64(g, "maxnestinglevel")
	data.Mssrvrecordlocation = utils.MapGetString(g, "mssrvrecordlocation")
	data.Nestedgroupextraction = utils.MapGetString(g, "nestedgroupextraction")
	data.Otpsecret = utils.MapGetString(g, "otpsecret")
	data.Passwdchange = utils.MapGetString(g, "passwdchange")
	data.Passwordlessmgmtaccess = utils.MapGetString(g, "passwordlessmgmtaccess")
	data.Pushservice = utils.MapGetString(g, "pushservice")
	data.Referraldnslookup = utils.MapGetString(g, "referraldnslookup")
	data.Requireuser = utils.MapGetString(g, "requireuser")
	data.Searchfilter = utils.MapGetString(g, "searchfilter")
	data.Sectype = utils.MapGetString(g, "sectype")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Servername = utils.MapGetString(g, "servername")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Sshpublickey = utils.MapGetString(g, "sshpublickey")
	data.Ssonameattribute = utils.MapGetString(g, "ssonameattribute")
	data.Subattributename = utils.MapGetString(g, "subattributename")
	data.Svrtype = utils.MapGetString(g, "svrtype")
	data.Validateservercert = utils.MapGetString(g, "validateservercert")

	// ldapbinddnpassword / ldapbinddnpassword_wo_version are secret or version
	// tracker inputs the GET never returns -> Null. ldapbinddnpassword_wo is an
	// Optional config-side helper -> left as configured.
	data.Ldapbinddnpassword = types.StringNull()
	data.LdapbinddnpasswordWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Ldapcontimeout = utils.MapGetInt64(g, "ldapcontimeout")
	data.Success = utils.MapGetInt64(g, "success")
	data.Failure = utils.MapGetInt64(g, "failure")
}
