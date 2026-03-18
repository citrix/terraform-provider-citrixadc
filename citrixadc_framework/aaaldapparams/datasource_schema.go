package aaaldapparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaaldapparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the LDAP server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute name used for group extraction from the LDAP server.",
			},
			"groupnameidentifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP-group attribute that uniquely identifies the group. No two groups on one LDAP server can have the same group name identifier.",
			},
			"groupsearchattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP-group attribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.",
			},
			"groupsearchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Search-expression that can be specified for sending group-search requests to the LDAP server.",
			},
			"groupsearchsubattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LDAP-group subattribute that designates the parent group of the specified group. Use this attribute to search for a group's parent group.",
			},
			"ldapbase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base (the server and location) from which LDAP search commands should start.\nIf the LDAP server is running locally, the default value of base is dc=netscaler, dc=com.",
			},
			"ldapbinddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Complete distinguished name (DN) string used for binding to the LDAP server.",
			},
			"ldapbinddnpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password for binding to the LDAP server.",
			},
			"ldaploginname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name attribute that the Citrix ADC uses to query the external LDAP server or an Active Directory.",
			},
			"maxnestinglevel": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of levels up to which the system can query nested LDAP groups.",
			},
			"nestedgroupextraction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Queries the external LDAP server to determine whether the specified group belongs to another group.",
			},
			"passwdchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept password change requests.",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search.\nFor example, the following values:\nvpnallowed=true,\nldaploginame=\"\"samaccount\"\"\nwhen combined with the user-supplied username \"\"bob\"\", yield the following LDAP search string:\n\"\"(&(vpnallowed=true)(samaccount=bob)\"\"",
			},
			"sectype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your LDAP server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the LDAP server listens for connections.",
			},
			"ssonameattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute used by the Citrix ADC to query an external LDAP server or Active Directory for an alternative username.\nThis alternative username is then used for single sign-on (SSO).",
			},
			"subattributename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subattribute name used for group extraction from the LDAP server.",
			},
			"svrtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of LDAP server.",
			},
		},
	}
}
