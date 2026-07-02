package sslservicegroup_sslcacertbundle_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslservicegroupSslcacertbundleBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertbundlename": schema.StringAttribute{
				Required:    true,
				Description: "CA certbundle name bound to the servicegroup.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}
