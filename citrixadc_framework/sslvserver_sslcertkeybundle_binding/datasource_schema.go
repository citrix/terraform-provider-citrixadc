package sslvserver_sslcertkeybundle_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverSslcertkeybundleBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"certkeybundlename": schema.StringAttribute{
				Required:    true,
				Description: "Certkeybundle name bound to the vserver.",
			},
			"snicertkeybundle": schema.BoolAttribute{
				Required:    true,
				Description: "Use this option to bind certkeybundle which will be used in SNI processing.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}
