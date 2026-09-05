package streamidentifier

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// StreamidentifierDataSourceModel is the data-source-specific model, decoupled
// from StreamidentifierResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type StreamidentifierDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Acceptancethreshold     types.String `tfsdk:"acceptancethreshold"`
	Appflowlog              types.String `tfsdk:"appflowlog"`
	Breachthreshold         types.Int64  `tfsdk:"breachthreshold"`
	Interval                types.Int64  `tfsdk:"interval"`
	Log                     types.String `tfsdk:"log"`
	Loginterval             types.Int64  `tfsdk:"loginterval"`
	Loglimit                types.Int64  `tfsdk:"loglimit"`
	Maxtransactionthreshold types.Int64  `tfsdk:"maxtransactionthreshold"`
	Mintransactionthreshold types.Int64  `tfsdk:"mintransactionthreshold"`
	Name                    types.String `tfsdk:"name"`
	Samplecount             types.Int64  `tfsdk:"samplecount"`
	Selectorname            types.String `tfsdk:"selectorname"`
	Snmptrap                types.String `tfsdk:"snmptrap"`
	Sort                    types.String `tfsdk:"sort"`
	Trackackonlypackets     types.String `tfsdk:"trackackonlypackets"`
	Tracktransactions       types.String `tfsdk:"tracktransactions"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/streamidentifier.json). Never settable; populated from GET.
	Rule types.List `tfsdk:"rule"`
}

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

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"rule": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Rule.",
			},
		},
	}
}

// streamidentifierDataSourceSetAttrFromGet projects a NITRO streamidentifier GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func streamidentifierDataSourceSetAttrFromGet(ctx context.Context, data *StreamidentifierDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In streamidentifierDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Acceptancethreshold = utils.MapGetString(g, "acceptancethreshold")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Breachthreshold = utils.MapGetInt64(g, "breachthreshold")
	data.Interval = utils.MapGetInt64(g, "interval")
	data.Log = utils.MapGetString(g, "log")
	data.Loginterval = utils.MapGetInt64(g, "loginterval")
	data.Loglimit = utils.MapGetInt64(g, "loglimit")
	data.Maxtransactionthreshold = utils.MapGetInt64(g, "maxtransactionthreshold")
	data.Mintransactionthreshold = utils.MapGetInt64(g, "mintransactionthreshold")
	data.Samplecount = utils.MapGetInt64(g, "samplecount")
	data.Selectorname = utils.MapGetString(g, "selectorname")
	data.Snmptrap = utils.MapGetString(g, "snmptrap")
	data.Sort = utils.MapGetString(g, "sort")
	data.Trackackonlypackets = utils.MapGetString(g, "trackackonlypackets")
	data.Tracktransactions = utils.MapGetString(g, "tracktransactions")

	// Read-only attributes.
	data.Rule = utils.MapGetStringList(g, "rule")
}
