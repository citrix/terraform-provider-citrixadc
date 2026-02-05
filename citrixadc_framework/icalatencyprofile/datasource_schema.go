package icalatencyprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IcalatencyprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"l7latencymaxnotifycount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Max notify Count. This is the upper limit on the number of notifications sent to the Insight Center within an interval where the Latency is above the threshold.",
			},
			"l7latencymonitoring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable L7 Latency monitoring for L7 latency notifications",
			},
			"l7latencynotifyinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Notify Interval. This is the interval at which the Citrix ADC sends out notifications to the Insight Center after the wait time has passed.",
			},
			"l7latencythresholdfactor": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency threshold factor. This is the factor by which the active latency should be greater than the minimum observed value to determine that the latency is high and may need to be reported",
			},
			"l7latencywaittime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Wait time. This is the time for which the Citrix ADC waits after the threshold is exceeded before it sends out a Notification to the Insight Center.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica l7latencyprofile\" or 'my ica l7latencyprofile').",
			},
		},
	}
}
