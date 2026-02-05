package nslimitidentifier

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NslimitidentifierDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"limitidentifier": schema.StringAttribute{
				Required:    true,
				Description: "Name for a rate limit identifier. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Reserved words must not be used.",
			},
			"limittype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Smooth or bursty request type.\n* SMOOTH - When you want the permitted number of requests in a given interval of time to be spread evenly across the timeslice\n* BURSTY - When you want the permitted number of requests to exhaust the quota anytime within the timeslice.\nThis argument is needed only when the mode is set to REQUEST_RATE.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth permitted, in kbps.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines the type of traffic to be tracked.\n* REQUEST_RATE - Tracks requests/timeslice.\n* CONNECTION - Tracks active transactions.\n\nExamples\n\n1. To permit 20 requests in 10 ms and 2 traps in 10 ms:\nadd limitidentifier limit_req -mode request_rate -limitType smooth -timeslice 1000 -Threshold 2000 -trapsInTimeSlice 200\n\n2. To permit 50 requests in 10 ms:\nset  limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5000 -limitType smooth\n\n3. To permit 1 request in 40 ms:\nset limitidentifier limit_req -mode request_rate -timeslice 2000 -Threshold 50 -limitType smooth\n\n4. To permit 1 request in 200 ms and 1 trap in 130 ms:\nset limitidentifier limit_req -mode request_rate -timeslice 1000 -Threshold 5 -limitType smooth -trapsInTimeSlice 8\n\n5. To permit 5000 requests in 1000 ms and 200 traps in 1000 ms:\nset limitidentifier limit_req  -mode request_rate -timeslice 1000 -Threshold 5000 -limitType BURSTY",
			},
			"selectorname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the rate limit selector. If this argument is NULL, rate limiting will be applied on all traffic received by the virtual server or the Citrix ADC (depending on whether the limit identifier is bound to a virtual server or globally) without any filtering.",
			},
			"threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that are allowed in the given timeslice when requests (mode is set as REQUEST_RATE) are tracked per timeslice.\nWhen connections (mode is set as CONNECTION) are tracked, it is the total number of connections that would be let through.",
			},
			"timeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in milliseconds, specified in multiples of 10, during which requests are tracked to check if they cross the threshold. This argument is needed only when the mode is set to REQUEST_RATE.",
			},
			"trapsintimeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of traps to be sent in the timeslice configured. A value of 0 indicates that traps are disabled.",
			},
		},
	}
}
