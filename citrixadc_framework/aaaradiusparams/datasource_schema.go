package aaaradiusparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaaradiusparamsDataSourceModel is the data-source-specific model, decoupled
// from AaaradiusparamsResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (groupauthname, ipaddress, builtin, feature). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type AaaradiusparamsDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Authentication             types.String `tfsdk:"authentication"`
	Authservretry              types.Int64  `tfsdk:"authservretry"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Callingstationid           types.String `tfsdk:"callingstationid"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Ipattributetype            types.Int64  `tfsdk:"ipattributetype"`
	Ipvendorid                 types.Int64  `tfsdk:"ipvendorid"`
	Messageauthenticator       types.String `tfsdk:"messageauthenticator"`
	Passencoding               types.String `tfsdk:"passencoding"`
	Pwdattributetype           types.Int64  `tfsdk:"pwdattributetype"`
	Pwdvendorid                types.Int64  `tfsdk:"pwdvendorid"`
	Radattributetype           types.Int64  `tfsdk:"radattributetype"`
	Radgroupseparator          types.String `tfsdk:"radgroupseparator"`
	Radgroupsprefix            types.String `tfsdk:"radgroupsprefix"`
	Radkey                     types.String `tfsdk:"radkey"`
	RadkeyWo                   types.String `tfsdk:"radkey_wo"`
	RadkeyWoVersion            types.Int64  `tfsdk:"radkey_wo_version"`
	Radnasid                   types.String `tfsdk:"radnasid"`
	Radnasip                   types.String `tfsdk:"radnasip"`
	Radvendorid                types.Int64  `tfsdk:"radvendorid"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tunnelendpointclientip     types.String `tfsdk:"tunnelendpointclientip"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/aaaradiusparams.json). Never settable; populated from GET.
	Groupauthname types.String `tfsdk:"groupauthname"`
	Ipaddress     types.String `tfsdk:"ipaddress"`
	Builtin       types.List   `tfsdk:"builtin"`
	Feature       types.String `tfsdk:"feature"`
}

func AaaradiusparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure the RADIUS server state to accept or refuse accounting messages.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure the RADIUS server state to accept or refuse authentication messages.",
			},
			"authservretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of retry by the Citrix ADC before getting response from the RADIUS server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server.",
			},
			"callingstationid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"ipattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IP attribute type in the RADIUS response.",
			},
			"ipvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID attribute in the RADIUS response.\nIf the attribute is not vendor-encoded, it is set to 0.",
			},
			"messageauthenticator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.",
			},
			"passencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server.",
			},
			"pwdattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute type of the Vendor ID in the RADIUS response.",
			},
			"pwdvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the password in the RADIUS response. Used to extract the user password.",
			},
			"radattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute type for RADIUS group extraction.",
			},
			"radgroupseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radgroupsprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "The key shared between the RADIUS server and clients.\nRequired for allowing the Citrix ADC to communicate with the RADIUS server.",
			},
			"radkey_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The key shared between the RADIUS server and clients.\nRequired for allowing the Citrix ADC to communicate with the RADIUS server.",
			},
			"radkey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a radkey_wo update.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server as the nasid part of the Radius protocol.",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the Citrix ADC IP (NSIP) address to the RADIUS server as the Network Access Server IP (NASIP) part of the Radius protocol.",
			},
			"radvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID for RADIUS group extraction.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your RADIUS server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the RADIUS server listens for connections.",
			},
			"tunnelendpointclientip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send Tunnel Endpoint Client IP address to the RADIUS server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"groupauthname": schema.StringAttribute{
				Computed:    true,
				Description: "Attribute name for group extraction from the RADIUS server.",
			},
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "IP Address.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// aaaradiusparamsDataSourceSetAttrFromGet projects a NITRO aaaradiusparams GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func aaaradiusparamsDataSourceSetAttrFromGet(ctx context.Context, data *AaaradiusparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaaradiusparamsDataSourceSetAttrFromGet Function")

	// aaaradiusparams is a singleton; use the same static ID as the resource.
	data.Id = types.StringValue("aaaradiusparams-config")

	// Read/write attributes as read-back outputs.
	data.Accounting = utils.MapGetString(g, "accounting")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authservretry = utils.MapGetInt64(g, "authservretry")
	data.Authtimeout = utils.MapGetInt64(g, "authtimeout")
	data.Callingstationid = utils.MapGetString(g, "callingstationid")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Ipattributetype = utils.MapGetInt64(g, "ipattributetype")
	data.Ipvendorid = utils.MapGetInt64(g, "ipvendorid")
	data.Messageauthenticator = utils.MapGetString(g, "messageauthenticator")
	data.Passencoding = utils.MapGetString(g, "passencoding")
	data.Pwdattributetype = utils.MapGetInt64(g, "pwdattributetype")
	data.Pwdvendorid = utils.MapGetInt64(g, "pwdvendorid")
	data.Radattributetype = utils.MapGetInt64(g, "radattributetype")
	data.Radgroupseparator = utils.MapGetString(g, "radgroupseparator")
	data.Radgroupsprefix = utils.MapGetString(g, "radgroupsprefix")
	data.Radnasid = utils.MapGetString(g, "radnasid")
	data.Radnasip = utils.MapGetString(g, "radnasip")
	data.Radvendorid = utils.MapGetInt64(g, "radvendorid")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Tunnelendpointclientip = utils.MapGetString(g, "tunnelendpointclientip")

	// radkey / radkey_wo(+version) are write-only or secret inputs the GET never
	// returns -> Null.
	data.Radkey = types.StringNull()
	data.RadkeyWo = types.StringNull()
	data.RadkeyWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Groupauthname = utils.MapGetString(g, "groupauthname")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
