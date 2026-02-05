package authenticationldapaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
