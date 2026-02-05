package aaatacacsparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaatacacsparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send accounting messages to the TACACS+ server.",
			},
			"auditfailedcmds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The option for sending accounting messages to the TACACS+ server.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use streaming authorization on the TACACS+ server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TACACS+ group attribute name.Used for group extraction on the TACACS+ server.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your TACACS+ server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the TACACS+ server listens for connections.",
			},
			"tacacssecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
		},
	}
}
