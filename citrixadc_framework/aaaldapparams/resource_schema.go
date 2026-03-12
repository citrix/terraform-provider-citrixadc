package aaaldapparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaaldapparamsResourceModel describes the resource data model.
type AaaldapparamsResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Groupnameidentifier        types.String `tfsdk:"groupnameidentifier"`
	Groupsearchattribute       types.String `tfsdk:"groupsearchattribute"`
	Groupsearchfilter          types.String `tfsdk:"groupsearchfilter"`
	Groupsearchsubattribute    types.String `tfsdk:"groupsearchsubattribute"`
	Ldapbase                   types.String `tfsdk:"ldapbase"`
	Ldapbinddn                 types.String `tfsdk:"ldapbinddn"`
	Ldapbinddnpassword         types.String `tfsdk:"ldapbinddnpassword"`
	Ldaploginname              types.String `tfsdk:"ldaploginname"`
	Maxnestinglevel            types.Int64  `tfsdk:"maxnestinglevel"`
	Nestedgroupextraction      types.String `tfsdk:"nestedgroupextraction"`
	Passwdchange               types.String `tfsdk:"passwdchange"`
	Searchfilter               types.String `tfsdk:"searchfilter"`
	Sectype                    types.String `tfsdk:"sectype"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Ssonameattribute           types.String `tfsdk:"ssonameattribute"`
	Subattributename           types.String `tfsdk:"subattributename"`
	Svrtype                    types.String `tfsdk:"svrtype"`
}

func (r *AaaldapparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaldapparams resource.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
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
				Default:     int64default.StaticInt64(2),
				Description: "Number of levels up to which the system can query nested LDAP groups.",
			},
			"nestedgroupextraction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Queries the external LDAP server to determine whether the specified group belongs to another group.",
			},
			"passwdchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Accept password change requests.",
			},
			"searchfilter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to be combined with the default LDAP user search string to form the value to use when executing an LDAP search.\nFor example, the following values:\nvpnallowed=true,\nldaploginame=\"\"samaccount\"\"\nwhen combined with the user-supplied username \"\"bob\"\", yield the following LDAP search string:\n\"\"(&(vpnallowed=true)(samaccount=bob)\"\"",
			},
			"sectype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("TLS"),
				Description: "Type of security used for communications between the Citrix ADC and the LDAP server. For the PLAINTEXT setting, no encryption is required.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your LDAP server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(389),
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
				Default:     stringdefault.StaticString("AAA_LDAP_SERVER_TYPE_DEFAULT"),
				Description: "The type of LDAP server.",
			},
		},
	}
}

func aaaldapparamsGetThePayloadFromtheConfig(ctx context.Context, data *AaaldapparamsResourceModel) aaa.Aaaldapparams {
	tflog.Debug(ctx, "In aaaldapparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaaldapparams := aaa.Aaaldapparams{}
	if !data.Authtimeout.IsNull() {
		aaaldapparams.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		aaaldapparams.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupattrname.IsNull() {
		aaaldapparams.Groupattrname = data.Groupattrname.ValueString()
	}
	if !data.Groupnameidentifier.IsNull() {
		aaaldapparams.Groupnameidentifier = data.Groupnameidentifier.ValueString()
	}
	if !data.Groupsearchattribute.IsNull() {
		aaaldapparams.Groupsearchattribute = data.Groupsearchattribute.ValueString()
	}
	if !data.Groupsearchfilter.IsNull() {
		aaaldapparams.Groupsearchfilter = data.Groupsearchfilter.ValueString()
	}
	if !data.Groupsearchsubattribute.IsNull() {
		aaaldapparams.Groupsearchsubattribute = data.Groupsearchsubattribute.ValueString()
	}
	if !data.Ldapbase.IsNull() {
		aaaldapparams.Ldapbase = data.Ldapbase.ValueString()
	}
	if !data.Ldapbinddn.IsNull() {
		aaaldapparams.Ldapbinddn = data.Ldapbinddn.ValueString()
	}
	if !data.Ldapbinddnpassword.IsNull() {
		aaaldapparams.Ldapbinddnpassword = data.Ldapbinddnpassword.ValueString()
	}
	if !data.Ldaploginname.IsNull() {
		aaaldapparams.Ldaploginname = data.Ldaploginname.ValueString()
	}
	if !data.Maxnestinglevel.IsNull() {
		aaaldapparams.Maxnestinglevel = utils.IntPtr(int(data.Maxnestinglevel.ValueInt64()))
	}
	if !data.Nestedgroupextraction.IsNull() {
		aaaldapparams.Nestedgroupextraction = data.Nestedgroupextraction.ValueString()
	}
	if !data.Passwdchange.IsNull() {
		aaaldapparams.Passwdchange = data.Passwdchange.ValueString()
	}
	if !data.Searchfilter.IsNull() {
		aaaldapparams.Searchfilter = data.Searchfilter.ValueString()
	}
	if !data.Sectype.IsNull() {
		aaaldapparams.Sectype = data.Sectype.ValueString()
	}
	if !data.Serverip.IsNull() {
		aaaldapparams.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		aaaldapparams.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Ssonameattribute.IsNull() {
		aaaldapparams.Ssonameattribute = data.Ssonameattribute.ValueString()
	}
	if !data.Subattributename.IsNull() {
		aaaldapparams.Subattributename = data.Subattributename.ValueString()
	}
	if !data.Svrtype.IsNull() {
		aaaldapparams.Svrtype = data.Svrtype.ValueString()
	}

	return aaaldapparams
}

func aaaldapparamsSetAttrFromGet(ctx context.Context, data *AaaldapparamsResourceModel, getResponseData map[string]interface{}) *AaaldapparamsResourceModel {
	tflog.Debug(ctx, "In aaaldapparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Authtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
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
	if val, ok := getResponseData["ldaploginname"]; ok && val != nil {
		data.Ldaploginname = types.StringValue(val.(string))
	} else {
		data.Ldaploginname = types.StringNull()
	}
	if val, ok := getResponseData["maxnestinglevel"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxnestinglevel = types.Int64Value(intVal)
		}
	} else {
		data.Maxnestinglevel = types.Int64Null()
	}
	if val, ok := getResponseData["nestedgroupextraction"]; ok && val != nil {
		data.Nestedgroupextraction = types.StringValue(val.(string))
	} else {
		data.Nestedgroupextraction = types.StringNull()
	}
	if val, ok := getResponseData["passwdchange"]; ok && val != nil {
		data.Passwdchange = types.StringValue(val.(string))
	} else {
		data.Passwdchange = types.StringNull()
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
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
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

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaaldapparams-config")

	return data
}
