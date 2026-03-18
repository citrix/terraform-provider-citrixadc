package gslbservice_dnsview_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbserviceDnsviewBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view",
			},
			"viewname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.",
			},
		},
	}
}
