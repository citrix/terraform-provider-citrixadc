package gslbservice_lbmonitor_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbserviceLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"monitor_name": schema.StringAttribute{
				Required:    true,
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the monitor bound to gslb service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},
		},
	}
}
