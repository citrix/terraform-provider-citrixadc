package service_lbmonitor_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServiceLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"monitor_name": schema.StringAttribute{
				Required:    true,
				Description: "The monitor Names.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The configured state (enable/disable) of the monitor on this server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service to which to bind a monitor.",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},
		},
	}
}
