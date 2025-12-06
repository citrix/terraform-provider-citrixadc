package sslcertkey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslCertKeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read SSL certificate key pair configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the SSL certificate key pair.",
			},
			"certkey": schema.StringAttribute{
				Required:    true,
				Description: "Name of the certificate and private-key pair to read.",
			},
			"cert": schema.StringAttribute{
				Computed:    true,
				Description: "Name of and path to the X509 certificate file.",
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "Name of and path to the private-key file.",
			},
			"password": schema.BoolAttribute{
				Computed:    true,
				Description: "Passphrase that was used to encrypt the private-key.",
			},
			"fipskey": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the FIPS key in the Hardware Security Module (HSM).",
			},
			"hsmkey": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the HSM key in the External Hardware Security Module (HSM).",
			},
			"inform": schema.StringAttribute{
				Computed:    true,
				Description: "Input format of the certificate and the private-key files (PEM, DER, or PFX).",
			},
			"expirymonitor": schema.StringAttribute{
				Computed:    true,
				Description: "Issue an alert when the certificate is about to expire.",
			},
			"notificationperiod": schema.Int64Attribute{
				Computed:    true,
				Description: "Time, in days, before certificate expiration at which to generate an alert.",
			},
			"bundle": schema.StringAttribute{
				Computed:    true,
				Description: "Parse the certificate chain as a single file.",
			},
			"linkcertkeyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the Certificate Authority certificate-key pair linked to this certificate.",
			},
			"nodomaincheck": schema.BoolAttribute{
				Computed:    true,
				Description: "Override the check for matching domain names during certificate update.",
			},
			"ocspstaplingcache": schema.BoolAttribute{
				Computed:    true,
				Description: "Clear cached ocspStapling response.",
			},
			"deletefromdevice": schema.BoolAttribute{
				Computed:    true,
				Description: "Delete cert/key file from file system.",
			},
			"deletecertkeyfilesonremoval": schema.StringAttribute{
				Computed:    true,
				Description: "Delete certificate and key files when the certificate is removed.",
			},
			"passplain": schema.StringAttribute{
				Computed:    true,
				Description: "Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.",
			},
			"passplain_wo": schema.StringAttribute{
				Computed:    true,
				Description: "Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.",
			},
			"passplain_wo_version": schema.Int64Attribute{
				Description: "Increment this version to signal a passplain_wo update.",
				Computed:    true,
			},
		},
	}
}
