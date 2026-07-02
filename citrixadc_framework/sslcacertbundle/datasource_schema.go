package sslcacertbundle

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslcacertbundleDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bundlefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the X509 CA certificate bundle file that is used to form cacertbundle entity. The CA certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The CA certificate bundle file consists of list of certificates.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the CA certbundle. The name will be used for bind/unbind/update operations. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
		},
	}
}
