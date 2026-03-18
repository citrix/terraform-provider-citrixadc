package streamidentifier

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func StreamidentifierDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acceptancethreshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Non-Breaching transactions to Total transactions threshold expressed in percent.\nMaximum of 6 decimal places is supported.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable Appflow logging for stream identifier",
			},
			"breachthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Breaching transactions threshold calculated over interval.",
			},
			"interval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location where objects collected on the identifier will be logged.",
			},
			"loginterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval in minutes for logging the collected objects.\nLog interval should be greater than or equal to the inteval \nof the stream identifier.",
			},
			"loglimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of objects to be logged in the log interval.",
			},
			"maxtransactionthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.",
			},
			"mintransactionthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum per transcation value of metric. Metric to be tracked is specified by tracktransactions attribute.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of stream identifier.",
			},
			"samplecount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the sample from which to select a request for evaluation. The smaller the sample count, the more accurate is the statistical data. To evaluate all requests, set the sample count to 1. However, such a low setting can result in excessive consumption of memory and processing resources.",
			},
			"selectorname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the selector to use with the stream identifier.",
			},
			"snmptrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable SNMP trap for stream identifier",
			},
			"sort": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).",
			},
			"trackackonlypackets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Track ack only packets as well. This setting is applicable only when packet rate limiting is being used.",
			},
			"tracktransactions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Track transactions exceeding configured threshold. Transaction tracking can be enabled for following metric: ResponseTime.\nBy default transaction tracking is disabled",
			},
		},
	}
}
