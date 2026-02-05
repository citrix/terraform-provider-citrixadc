package netprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"mbf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.",
			},
			"overridelsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "USNIP/USIP settings override LSN settings for configured\n              service/virtual server traffic..",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Action (Enabled/Disabled)",
			},
			"proxyprotocolaftertlshandshake": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake",
			},
			"proxyprotocoltxversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Protocol Version (V1/V2)",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or the name of an IP set.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}
