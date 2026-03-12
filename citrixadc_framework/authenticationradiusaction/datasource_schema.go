package authenticationradiusaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
