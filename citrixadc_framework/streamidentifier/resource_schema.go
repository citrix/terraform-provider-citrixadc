package streamidentifier

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// StreamidentifierResourceModel describes the resource data model.
type StreamidentifierResourceModel struct {
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
}

func (r *StreamidentifierResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the streamidentifier resource.",
			},
			"acceptancethreshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Non-Breaching transactions to Total transactions threshold expressed in percent.\nMaximum of 6 decimal places is supported.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     int64default.StaticInt64(1),
				Description: "Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Location where objects collected on the identifier will be logged.",
			},
			"loginterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time interval in minutes for logging the collected objects.\nLog interval should be greater than or equal to the inteval \nof the stream identifier.",
			},
			"loglimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of stream identifier.",
			},
			"samplecount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/disable SNMP trap for stream identifier",
			},
			"sort": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("REQUESTS"),
				Description: "Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).",
			},
			"trackackonlypackets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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

func streamidentifierGetThePayloadFromthePlan(ctx context.Context, data *StreamidentifierResourceModel) stream.Streamidentifier {
	tflog.Debug(ctx, "In streamidentifierGetThePayloadFromthePlan Function")

	// Create API request body from the model
	streamidentifier := stream.Streamidentifier{}
	if !data.Acceptancethreshold.IsNull() && !data.Acceptancethreshold.IsUnknown() {
		streamidentifier.Acceptancethreshold = data.Acceptancethreshold.ValueString()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		streamidentifier.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Breachthreshold.IsNull() && !data.Breachthreshold.IsUnknown() {
		streamidentifier.Breachthreshold = utils.IntPtr(int(data.Breachthreshold.ValueInt64()))
	}
	if !data.Interval.IsNull() && !data.Interval.IsUnknown() {
		streamidentifier.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Log.IsNull() && !data.Log.IsUnknown() {
		streamidentifier.Log = data.Log.ValueString()
	}
	if !data.Loginterval.IsNull() && !data.Loginterval.IsUnknown() {
		streamidentifier.Loginterval = utils.IntPtr(int(data.Loginterval.ValueInt64()))
	}
	if !data.Loglimit.IsNull() && !data.Loglimit.IsUnknown() {
		streamidentifier.Loglimit = utils.IntPtr(int(data.Loglimit.ValueInt64()))
	}
	if !data.Maxtransactionthreshold.IsNull() && !data.Maxtransactionthreshold.IsUnknown() {
		streamidentifier.Maxtransactionthreshold = utils.IntPtr(int(data.Maxtransactionthreshold.ValueInt64()))
	}
	if !data.Mintransactionthreshold.IsNull() && !data.Mintransactionthreshold.IsUnknown() {
		streamidentifier.Mintransactionthreshold = utils.IntPtr(int(data.Mintransactionthreshold.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		streamidentifier.Name = data.Name.ValueString()
	}
	if !data.Samplecount.IsNull() && !data.Samplecount.IsUnknown() {
		streamidentifier.Samplecount = utils.IntPtr(int(data.Samplecount.ValueInt64()))
	}
	if !data.Selectorname.IsNull() && !data.Selectorname.IsUnknown() {
		streamidentifier.Selectorname = data.Selectorname.ValueString()
	}
	if !data.Snmptrap.IsNull() && !data.Snmptrap.IsUnknown() {
		streamidentifier.Snmptrap = data.Snmptrap.ValueString()
	}
	if !data.Sort.IsNull() && !data.Sort.IsUnknown() {
		streamidentifier.Sort = data.Sort.ValueString()
	}
	if !data.Trackackonlypackets.IsNull() && !data.Trackackonlypackets.IsUnknown() {
		streamidentifier.Trackackonlypackets = data.Trackackonlypackets.ValueString()
	}
	if !data.Tracktransactions.IsNull() && !data.Tracktransactions.IsUnknown() {
		streamidentifier.Tracktransactions = data.Tracktransactions.ValueString()
	}

	return streamidentifier
}

func streamidentifierSetAttrFromGet(ctx context.Context, data *StreamidentifierResourceModel, getResponseData map[string]interface{}) *StreamidentifierResourceModel {
	tflog.Debug(ctx, "In streamidentifierSetAttrFromGet Function")

	// Convert API response to model.
	// The else-branches only null a value when it is Unknown (Computed, unset in
	// config/state). A known configured value that NITRO omits from GET
	// (omit-on-default) is preserved to avoid "inconsistent result after apply".
	if val, ok := getResponseData["acceptancethreshold"]; ok && val != nil {
		data.Acceptancethreshold = types.StringValue(val.(string))
	} else if data.Acceptancethreshold.IsUnknown() {
		data.Acceptancethreshold = types.StringNull()
	}
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else if data.Appflowlog.IsUnknown() {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["breachthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Breachthreshold = types.Int64Value(intVal)
		}
	} else if data.Breachthreshold.IsUnknown() {
		data.Breachthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	} else if data.Interval.IsUnknown() {
		data.Interval = types.Int64Null()
	}
	if val, ok := getResponseData["log"]; ok && val != nil {
		data.Log = types.StringValue(val.(string))
	} else if data.Log.IsUnknown() {
		data.Log = types.StringNull()
	}
	if val, ok := getResponseData["loginterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Loginterval = types.Int64Value(intVal)
		}
	} else if data.Loginterval.IsUnknown() {
		data.Loginterval = types.Int64Null()
	}
	if val, ok := getResponseData["loglimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Loglimit = types.Int64Value(intVal)
		}
	} else if data.Loglimit.IsUnknown() {
		data.Loglimit = types.Int64Null()
	}
	if val, ok := getResponseData["maxtransactionthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxtransactionthreshold = types.Int64Value(intVal)
		}
	} else if data.Maxtransactionthreshold.IsUnknown() {
		data.Maxtransactionthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["mintransactionthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mintransactionthreshold = types.Int64Value(intVal)
		}
	} else if data.Mintransactionthreshold.IsUnknown() {
		data.Mintransactionthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["samplecount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Samplecount = types.Int64Value(intVal)
		}
	} else if data.Samplecount.IsUnknown() {
		data.Samplecount = types.Int64Null()
	}
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else if data.Selectorname.IsUnknown() {
		data.Selectorname = types.StringNull()
	}
	if val, ok := getResponseData["snmptrap"]; ok && val != nil {
		data.Snmptrap = types.StringValue(val.(string))
	} else if data.Snmptrap.IsUnknown() {
		data.Snmptrap = types.StringNull()
	}
	if val, ok := getResponseData["sort"]; ok && val != nil {
		data.Sort = types.StringValue(val.(string))
	} else if data.Sort.IsUnknown() {
		data.Sort = types.StringNull()
	}
	if val, ok := getResponseData["trackackonlypackets"]; ok && val != nil {
		data.Trackackonlypackets = types.StringValue(val.(string))
	} else if data.Trackackonlypackets.IsUnknown() {
		data.Trackackonlypackets = types.StringNull()
	}
	if val, ok := getResponseData["tracktransactions"]; ok && val != nil {
		data.Tracktransactions = types.StringValue(val.(string))
	} else if data.Tracktransactions.IsUnknown() {
		data.Tracktransactions = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
