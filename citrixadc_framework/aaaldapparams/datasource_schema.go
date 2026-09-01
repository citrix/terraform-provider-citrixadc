package aaaldapparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaaldapparamsDataSourceModel is the data-source-specific model, decoupled from
// AaaldapparamsResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AaaldapparamsDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Authtimeout                 types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup  types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname               types.String `tfsdk:"groupattrname"`
	Groupnameidentifier         types.String `tfsdk:"groupnameidentifier"`
	Groupsearchattribute        types.String `tfsdk:"groupsearchattribute"`
	Groupsearchfilter           types.String `tfsdk:"groupsearchfilter"`
	Groupsearchsubattribute     types.String `tfsdk:"groupsearchsubattribute"`
	Ldapbase                    types.String `tfsdk:"ldapbase"`
	Ldapbinddn                  types.String `tfsdk:"ldapbinddn"`
	Ldapbinddnpassword          types.String `tfsdk:"ldapbinddnpassword"`
	LdapbinddnpasswordWo        types.String `tfsdk:"ldapbinddnpassword_wo"`
	LdapbinddnpasswordWoVersion types.Int64  `tfsdk:"ldapbinddnpassword_wo_version"`
	Ldaploginname               types.String `tfsdk:"ldaploginname"`
	Maxnestinglevel             types.Int64  `tfsdk:"maxnestinglevel"`
	Nestedgroupextraction       types.String `tfsdk:"nestedgroupextraction"`
	Passwdchange                types.String `tfsdk:"passwdchange"`
	Searchfilter                types.String `tfsdk:"searchfilter"`
	Sectype                     types.String `tfsdk:"sectype"`
	Serverip                    types.String `tfsdk:"serverip"`
	Serverport                  types.Int64  `tfsdk:"serverport"`
	Ssonameattribute            types.String `tfsdk:"ssonameattribute"`
	Subattributename            types.String `tfsdk:"subattributename"`
	Svrtype                     types.String `tfsdk:"svrtype"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/aaaldapparams.json). Never settable; populated from GET.
	Ldapcontimeout types.Int64  `tfsdk:"ldapcontimeout"`
	Groupauthname  types.String `tfsdk:"groupauthname"`
	Builtin        types.List   `tfsdk:"builtin"`
	Feature        types.String `tfsdk:"feature"`
}

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
				Sensitive:   true,
				Description: "Password for binding to the LDAP server.",
			},
			"ldapbinddnpassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for binding to the LDAP server.",
			},
			"ldapbinddnpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a ldapbinddnpassword_wo update.",
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

			// Read-only (GET-only) attributes surfaced by the data source. All Computed.
			"ldapcontimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of seconds the Citrix ADC waits for the LDAP connection (TCP connection and SSL/TLS handshake) to be established with the LDAP server.",
			},
			"groupauthname": schema.StringAttribute{
				Computed:    true,
				Description: "AAA group used to associate AAA users with an AAA group.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// aaaldapparamsDataSourceSetAttrFromGet projects a NITRO aaaldapparams GET
// response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaaldapparamsDataSourceSetAttrFromGet(ctx context.Context, data *AaaldapparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaaldapparamsDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Authtimeout = utils.MapGetInt64(g, "authtimeout")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Groupattrname = utils.MapGetString(g, "groupattrname")
	data.Groupnameidentifier = utils.MapGetString(g, "groupnameidentifier")
	data.Groupsearchattribute = utils.MapGetString(g, "groupsearchattribute")
	data.Groupsearchfilter = utils.MapGetString(g, "groupsearchfilter")
	data.Groupsearchsubattribute = utils.MapGetString(g, "groupsearchsubattribute")
	data.Ldapbase = utils.MapGetString(g, "ldapbase")
	data.Ldapbinddn = utils.MapGetString(g, "ldapbinddn")
	data.Ldaploginname = utils.MapGetString(g, "ldaploginname")
	data.Maxnestinglevel = utils.MapGetInt64(g, "maxnestinglevel")
	data.Nestedgroupextraction = utils.MapGetString(g, "nestedgroupextraction")
	data.Passwdchange = utils.MapGetString(g, "passwdchange")
	data.Searchfilter = utils.MapGetString(g, "searchfilter")
	data.Sectype = utils.MapGetString(g, "sectype")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Ssonameattribute = utils.MapGetString(g, "ssonameattribute")
	data.Subattributename = utils.MapGetString(g, "subattributename")
	data.Svrtype = utils.MapGetString(g, "svrtype")

	// ldapbinddnpassword / _wo(+version) are write-only secrets the GET never
	// returns -> Null.
	data.Ldapbinddnpassword = types.StringNull()
	data.LdapbinddnpasswordWo = types.StringNull()
	data.LdapbinddnpasswordWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Ldapcontimeout = utils.MapGetInt64(g, "ldapcontimeout")
	data.Groupauthname = utils.MapGetString(g, "groupauthname")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// aaaldapparams is a singleton; the ID is a fixed system-generated identifier.
	data.Id = types.StringValue("aaaldapparams-config")
}
