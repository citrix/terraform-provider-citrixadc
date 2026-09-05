package nslimitidentifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslimitidentifierDataSourceModel describes the DATASOURCE data model. It
// mirrors the configurable attributes surfaced by the datasource PLUS the
// read-only rate-limit metadata the NITRO `nslimitidentifier` GET returns. It is
// decoupled from the resource model so the data source can expose the full GET
// projection (including GET-only fields the resource intentionally omits).
type NslimitidentifierDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Limitidentifier   types.String `tfsdk:"limitidentifier"`
	Alertsintimeslice types.Int64  `tfsdk:"alertsintimeslice"`
	Limittype         types.String `tfsdk:"limittype"`
	Maxbandwidth      types.Int64  `tfsdk:"maxbandwidth"`
	Mode              types.String `tfsdk:"mode"`
	Selectorname      types.String `tfsdk:"selectorname"`
	Threshold         types.Int64  `tfsdk:"threshold"`
	Timealign         types.String `tfsdk:"timealign"`
	Timeslice         types.Int64  `tfsdk:"timeslice"`
	Trapsintimeslice  types.Int64  `tfsdk:"trapsintimeslice"`

	// Read-only (GET-only) rate-limit metadata from the NITRO doc read-only set
	// (zion73x_readonly/nslimitidentifier.json). Never settable; from GET.
	Ngname                    types.String `tfsdk:"ngname"`
	Hits                      types.Int64  `tfsdk:"hits"`
	Drop                      types.Int64  `tfsdk:"drop"`
	Rule                      types.List   `tfsdk:"rule"`
	Time                      types.Int64  `tfsdk:"time"`
	Total                     types.Int64  `tfsdk:"total"`
	Trapscomputedintimeslice  types.Int64  `tfsdk:"trapscomputedintimeslice"`
	Computedtraptimeslice     types.Int64  `tfsdk:"computedtraptimeslice"`
	Alertscomputedintimeslice types.Int64  `tfsdk:"alertscomputedintimeslice"`
	Computedalerttimeslice    types.Int64  `tfsdk:"computedalerttimeslice"`
	Referencecount            types.Int64  `tfsdk:"referencecount"`
}

// nslimitidentifierDataSourceSetAttrFromGet projects a NITRO nslimitidentifier
// GET response onto the data-source model using the shared utils.MapGet*
// helpers. Attributes the GET omits are left Null.
func nslimitidentifierDataSourceSetAttrFromGet(ctx context.Context, data *NslimitidentifierDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nslimitidentifierDataSourceSetAttrFromGet Function")

	data.Limitidentifier = utils.MapGetString(g, "limitidentifier")
	data.Alertsintimeslice = utils.MapGetInt64(g, "alertsintimeslice")
	data.Limittype = utils.MapGetString(g, "limittype")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Mode = utils.MapGetString(g, "mode")
	data.Selectorname = utils.MapGetString(g, "selectorname")
	data.Threshold = utils.MapGetInt64(g, "threshold")
	data.Timealign = utils.MapGetString(g, "timealign")
	data.Timeslice = utils.MapGetInt64(g, "timeslice")
	data.Trapsintimeslice = utils.MapGetInt64(g, "trapsintimeslice")

	// Read-only rate-limit metadata.
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Drop = utils.MapGetInt64(g, "drop")
	data.Rule = utils.MapGetStringList(g, "rule")
	data.Time = utils.MapGetInt64(g, "time")
	data.Total = utils.MapGetInt64(g, "total")
	data.Trapscomputedintimeslice = utils.MapGetInt64(g, "trapscomputedintimeslice")
	data.Computedtraptimeslice = utils.MapGetInt64(g, "computedtraptimeslice")
	data.Alertscomputedintimeslice = utils.MapGetInt64(g, "alertscomputedintimeslice")
	data.Computedalerttimeslice = utils.MapGetInt64(g, "computedalerttimeslice")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")

	// ID matches the resource: limitidentifier.
	data.Id = types.StringValue(data.Limitidentifier.ValueString())
}

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
			"alertsintimeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of appflow alerts to be sent in the timeslice configured. A value of 0 indicates that alerts are disabled. A value of 65535 indicates no limit on number of appflow alerts.",
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
			"timealign": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value MINUTE will align the time windows for a configured timeslice to Minute boundary. TimeSlice values should be integrals of 60000ms when value MINUTE is choosen. Default : NONE, timeslice alignments will happen with next 10ms.",
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

			// Read-only (GET-only) rate-limit metadata surfaced by the data source.
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Nodegroup name to which this identifier belongs. Null when the appliance omits it.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times this identifier was evaluated. Null when the appliance omits it.",
			},
			"drop": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times action was taken. Null when the appliance omits it.",
			},
			"rule": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Rule. Null when the appliance omits it.",
			},
			"time": schema.Int64Attribute{
				Computed:    true,
				Description: "Time interval considered for rate limiting. Null when the appliance omits it.",
			},
			"total": schema.Int64Attribute{
				Computed:    true,
				Description: "Maximum number of requests permitted in the computed timeslice. Null when the appliance omits it.",
			},
			"trapscomputedintimeslice": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of traps that would be sent in the timeslice configured. Null when the appliance omits it.",
			},
			"computedtraptimeslice": schema.Int64Attribute{
				Computed:    true,
				Description: "The time interval computed for sending traps. Null when the appliance omits it.",
			},
			"alertscomputedintimeslice": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of appflow alerts that would be sent in the timeslice configured. Null when the appliance omits it.",
			},
			"computedalerttimeslice": schema.Int64Attribute{
				Computed:    true,
				Description: "The time interval computed for sending appflow alerts. Null when the appliance omits it.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of transactions pointing to this entry. Null when the appliance omits it.",
			},
		},
	}
}
