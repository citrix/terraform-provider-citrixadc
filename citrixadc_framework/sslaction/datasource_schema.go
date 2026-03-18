package sslaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertgrpname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This action will allow to pick CA(s) from the specific CA group, to verify the client certificate.",
			},
			"certfingerprintdigest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Digest algorithm used to compute the fingerprint of the client certificate.",
			},
			"certfingerprintheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client certificate fingerprint.",
			},
			"certhashheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client certificate signature (hash).",
			},
			"certheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client certificate.",
			},
			"certissuerheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client certificate issuer details.",
			},
			"certnotafterheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the certificate's expiry date.",
			},
			"certnotbeforeheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the date and time from which the certificate is valid.",
			},
			"certserialheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client serial number.",
			},
			"certsubjectheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the client certificate subject.",
			},
			"cipher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the cipher suite that the client and the Citrix ADC negotiated for the SSL session into the HTTP header of the request being sent to the web server. The appliance inserts the cipher-suite name, SSL protocol, export or non-export string, and cipher strength bit, depending on the type of browser connecting to the SSL virtual server or service (for example, Cipher-Suite: RC4- MD5 SSLv3 Non-Export 128-bit).",
			},
			"cipherheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the name of the cipher suite.",
			},
			"clientauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform client certificate authentication.",
			},
			"clientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the entire client certificate into the HTTP header of the request being sent to the web server. The certificate is inserted in ASCII (PEM) format.",
			},
			"clientcertfingerprint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the certificate's fingerprint into the HTTP header of the request being sent to the web server. The fingerprint is derived by computing the specified hash value (SHA256, for example) of the DER-encoding of the client certificate.",
			},
			"clientcerthash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the certificate's signature into the HTTP header of the request being sent to the web server. The signature is the value extracted directly from the X.509 certificate signature field. All X.509 certificates contain a signature field.",
			},
			"clientcertissuer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the certificate issuer details into the HTTP header of the request being sent to the web server.",
			},
			"clientcertnotafter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the date of expiry of the certificate into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time at which the certificate expires.",
			},
			"clientcertnotbefore": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the date from which the certificate is valid into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time from which it is valid.",
			},
			"clientcertserialnumber": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the entire client serial number into the HTTP header of the request being sent to the web server.",
			},
			"clientcertsubject": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the client certificate subject, also known as the distinguished name (DN), into the HTTP header of the request being sent to the web server.",
			},
			"clientcertverification": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate verification is mandatory or optional.",
			},
			"forward": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This action takes an argument a vserver name, to this vserver one will be able to forward all the packets.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SSL action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"owasupport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the appliance is in front of an Outlook Web Access (OWA) server, insert a special header field, FRONT-END-HTTPS: ON, into the HTTP requests going to the OWA server. This header communicates to the server that the transaction is HTTPS and not HTTP.",
			},
			"sessionid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the SSL session ID into the HTTP header of the request being sent to the web server. Every SSL connection that the client and the Citrix ADC share has a unique ID that identifies the specific connection.",
			},
			"sessionidheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header into which to insert the Session ID.",
			},
			"ssllogprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the ssllogprofile.",
			},
		},
	}
}
