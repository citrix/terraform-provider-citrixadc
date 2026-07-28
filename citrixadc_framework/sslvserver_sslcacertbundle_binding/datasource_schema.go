package sslvserver_sslcacertbundle_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverSslcacertbundleBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertbundlename": schema.StringAttribute{
				Required:    true,
				Description: "CA certbundle name bound to the vserver.",
			},
			"skipcacertbundle": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}
