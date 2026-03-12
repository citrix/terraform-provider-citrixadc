package dnskey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnskeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to generate the key.",
			},
			"autorollover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable/disable key rollover automatically.\nNote:\n* Key name will be appended with _AR1 for successor key. For e.g. current key=k1, successor key=k1_AR1.\n* Key name can be truncated if current name length is more than 58 bytes to accomodate the suffix.",
			},
			"expires": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which to consider the key valid, after the key is used to sign a zone.",
			},
			"filenameprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.",
			},
			"keyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the public-private key pair to publish in the zone.",
			},
			"keysize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the key, in bits.",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of key to create.",
			},
			"notificationperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm. \nIn case autorollover option is enabled, rollover for successor key will be intiated at this time. No notification trap will be sent.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Passphrase for reading the encrypted public/private DNS keys",
			},
			"privatekey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the private key.",
			},
			"publickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the public key.",
			},
			"revoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Revoke the key. Note: This operation is non-reversible.",
			},
			"rollovermethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method used for automatic rollover.\n* Key type: ZSK, Method: PrePublication or DoubleSignature.\n* Key type: KSK, Method: DoubleRRSet.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. This is a mandatory argument",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used.",
			},
			"units1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the expiry period.",
			},
			"units2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the notification period.",
			},
			"zonename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the zone for which to create a key.",
			},
		},
	}
}
