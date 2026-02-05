package sslcrl

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcrlDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basedn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN.",
			},
			"binary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the LDAP-based CRL retrieval mode to binary.",
			},
			"binddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL.",
			},
			"cacertfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CA certificate file.\n/nsconfig/ssl/ is the default path.",
			},
			"cakeyfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path",
			},
			"crlname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my crl\" or 'my crl').",
			},
			"crlpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the CRL file. /var/netscaler/ssl/ is the default path.",
			},
			"day": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY.",
			},
			"gencrl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path.",
			},
			"indexfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path",
			},
			"inform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the CRL file. The two formats supported on the appliance are:\nPEM - Privacy Enhanced Mail.\nDER - Distinguished Encoding Rule.",
			},
			"interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CRL refresh interval. Use the NONE setting to unset this parameter.",
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port for the LDAP server.",
			},
			"refresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set CRL auto refresh.",
			},
			"revoke": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Extent of the search operation on the LDAP server. Available settings function as follows:\nOne - One level below Base DN.\nBase - Exactly the same level as Base DN.",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the LDAP server from which to fetch the CRLs.",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the CRL distribution point.",
			},
		},
	}
}
