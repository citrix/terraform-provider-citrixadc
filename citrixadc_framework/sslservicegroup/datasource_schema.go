package sslservicegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslservicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server",
			},
			"ocspstapling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:\nENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.\nDISABLED: The appliance does not check the status of the server certificate.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable sending SSL Close-Notify at the end of a transaction",
			},
			"serverauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of server authentication support for the SSL service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service group for which to set advanced configuration.",
			},
			"sessreuse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the Server Name Indication (SNI) feature on the service. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.",
			},
			"ssl3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of SSLv3 protocol support for the SSL service group.\nNote: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.",
			},
			"sslclientlogs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile that contains SSL settings for the Service Group.",
			},
			"strictsigdigestcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter indicating to check whether peer's certificate is signed with one of signature-hash combination supported by Citrix ADC",
			},
			"tls1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.0 protocol support for the SSL service group.",
			},
			"tls11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.1 protocol support for the SSL service group.",
			},
			"tls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.2 protocol support for the SSL service group.",
			},
			"tls13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of TLSv1.3 protocol support for the SSL service group.",
			},
		},
	}
}
