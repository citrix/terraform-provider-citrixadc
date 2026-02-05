package aaaradiusparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
				Description: "The key shared between the RADIUS server and clients.\nRequired for allowing the Citrix ADC to communicate with the RADIUS server.",
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
		},
	}
}
