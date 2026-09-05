package authenticationradiusaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationradiusactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationradiusactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (ipaddress, success, failure). Every non-key attribute is Computed.
type AuthenticationradiusactionDataSourceModel struct {
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
	Name                       types.String `tfsdk:"name"`
	Passencoding               types.String `tfsdk:"passencoding"`
	Pwdattributetype           types.Int64  `tfsdk:"pwdattributetype"`
	Pwdvendorid                types.Int64  `tfsdk:"pwdvendorid"`
	Radattributetype           types.Int64  `tfsdk:"radattributetype"`
	Radgroupseparator          types.String `tfsdk:"radgroupseparator"`
	Radgroupsprefix            types.String `tfsdk:"radgroupsprefix"`
	Radkey                     types.String `tfsdk:"radkey"`
	Radnasid                   types.String `tfsdk:"radnasid"`
	Radnasip                   types.String `tfsdk:"radnasip"`
	Radvendorid                types.Int64  `tfsdk:"radvendorid"`
	Serverip                   types.String `tfsdk:"serverip"`
	Servername                 types.String `tfsdk:"servername"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Targetlbvserver            types.String `tfsdk:"targetlbvserver"`
	Transport                  types.String `tfsdk:"transport"`
	Tunnelendpointclientip     types.String `tfsdk:"tunnelendpointclientip"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationradiusaction.json). Never settable;
	// populated from GET.
	Ipaddress types.String `tfsdk:"ipaddress"`
	Success   types.Int64  `tfsdk:"success"`
	Failure   types.Int64  `tfsdk:"failure"`
}

func AuthenticationradiusactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the RADIUS server is currently accepting accounting messages.",
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
				Description: "Number of seconds the Citrix ADC waits for a response from the RADIUS server.",
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
				Description: "Remote IP address attribute type in a RADIUS response.",
			},
			"ipvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the intranet IP attribute in the RADIUS response.\nNOTE: A value of 0 indicates that the attribute is not vendor encoded.",
			},
			"messageauthenticator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RADIUS action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.",
			},
			"passencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encoding type for passwords in RADIUS packets that the Citrix ADC sends to the RADIUS server.",
			},
			"pwdattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-specific password attribute type in a RADIUS response.",
			},
			"pwdvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor ID of the attribute, in the RADIUS response, used to extract the user password.",
			},
			"radattributetype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS attribute type, used for RADIUS group extraction.",
			},
			"radgroupseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS group separator string\nThe group separator delimits group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radgroupsprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS groups prefix string.\nThis groups prefix precedes the group names within a RADIUS attribute for RADIUS group extraction.",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Key shared between the RADIUS server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the RADIUS server.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If configured, this string is sent to the RADIUS server as the Network Access Server ID (NASID).",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, the Citrix ADC IP address (NSIP) is sent to the RADIUS server as the  Network Access Server IP (NASIP) address.\nThe RADIUS protocol defines the meaning and use of the NASIP address.",
			},
			"radvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS vendor ID attribute, used for RADIUS group extraction.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address assigned to the RADIUS server.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS server name as a FQDN.  Mutually exclusive with RADIUS IP address.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the RADIUS server listens for connections.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If transport mode is TLS, specify the name of LB vserver to associate. The LB vserver needs to be of type TCP and service associated needs to be SSL_TCP",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transport mode to RADIUS server.",
			},
			"tunnelendpointclientip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send Tunnel Endpoint Client IP address to the RADIUS server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "IP address.",
			},
			"success": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of successful authentication requests.",
			},
			"failure": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of failed authentication requests.",
			},
		},
	}
}

// authenticationradiusactionDataSourceSetAttrFromGet projects a NITRO
// authenticationradiusaction GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationradiusactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationradiusactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationradiusactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

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
	data.Servername = utils.MapGetString(g, "servername")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Targetlbvserver = utils.MapGetString(g, "targetlbvserver")
	data.Transport = utils.MapGetString(g, "transport")
	data.Tunnelendpointclientip = utils.MapGetString(g, "tunnelendpointclientip")

	// radkey is a secret input the GET never returns -> Null.
	data.Radkey = types.StringNull()

	// Read-only metadata.
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Success = utils.MapGetInt64(g, "success")
	data.Failure = utils.MapGetInt64(g, "failure")
}
