package sslcertkeybundle

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcertkeybundleDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bundlefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. The certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The certificate bundle file consists of list of certificates and one key in PEM format.",
			},
			"certkeybundlename": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the cerKeyBundle. The name will be used to bind/unbind certkey bundle to vip. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"passplain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a passplain_wo update.",
			},
		},
	}
}
