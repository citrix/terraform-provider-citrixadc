package streamidentifier

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Default:     int64default.StaticInt64(1),
				Description: "Number of minutes of data to use when calculating session statistics (number of requests, bandwidth, and response times). The interval is a moving window that keeps the most recently collected data. Older data is discarded at regular intervals.",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Location where objects collected on the identifier will be logged.",
			},
			"loginterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time interval in minutes for logging the collected objects.\nLog interval should be greater than or equal to the inteval \nof the stream identifier.",
			},
			"loglimit": schema.Int64Attribute{
				Optional:    true,
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
				Required:    true,
				Description: "The name of stream identifier.",
			},
			"samplecount": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Size of the sample from which to select a request for evaluation. The smaller the sample count, the more accurate is the statistical data. To evaluate all requests, set the sample count to 1. However, such a low setting can result in excessive consumption of memory and processing resources.",
			},
			"selectorname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the selector to use with the stream identifier.",
			},
			"snmptrap": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable/disable SNMP trap for stream identifier",
			},
			"sort": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("REQUESTS"),
				Description: "Sort stored records by the specified statistics column, in descending order. Performed during data collection, the sorting enables real-time data evaluation through Citrix ADC policies (for example, compression and caching policies) that use functions such as IS_TOP(n).",
			},
			"trackackonlypackets": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Track ack only packets as well. This setting is applicable only when packet rate limiting is being used.",
			},
			"tracktransactions": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Track transactions exceeding configured threshold. Transaction tracking can be enabled for following metric: ResponseTime.\nBy default transaction tracking is disabled",
			},
		},
	}
}

func streamidentifierGetThePayloadFromtheConfig(ctx context.Context, data *StreamidentifierResourceModel) stream.Streamidentifier {
	tflog.Debug(ctx, "In streamidentifierGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	streamidentifier := stream.Streamidentifier{}
	if !data.Acceptancethreshold.IsNull() {
		streamidentifier.Acceptancethreshold = data.Acceptancethreshold.ValueString()
	}
	if !data.Appflowlog.IsNull() {
		streamidentifier.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Breachthreshold.IsNull() {
		streamidentifier.Breachthreshold = utils.IntPtr(int(data.Breachthreshold.ValueInt64()))
	}
	if !data.Interval.IsNull() {
		streamidentifier.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Log.IsNull() {
		streamidentifier.Log = data.Log.ValueString()
	}
	if !data.Loginterval.IsNull() {
		streamidentifier.Loginterval = utils.IntPtr(int(data.Loginterval.ValueInt64()))
	}
	if !data.Loglimit.IsNull() {
		streamidentifier.Loglimit = utils.IntPtr(int(data.Loglimit.ValueInt64()))
	}
	if !data.Maxtransactionthreshold.IsNull() {
		streamidentifier.Maxtransactionthreshold = utils.IntPtr(int(data.Maxtransactionthreshold.ValueInt64()))
	}
	if !data.Mintransactionthreshold.IsNull() {
		streamidentifier.Mintransactionthreshold = utils.IntPtr(int(data.Mintransactionthreshold.ValueInt64()))
	}
	if !data.Name.IsNull() {
		streamidentifier.Name = data.Name.ValueString()
	}
	if !data.Samplecount.IsNull() {
		streamidentifier.Samplecount = utils.IntPtr(int(data.Samplecount.ValueInt64()))
	}
	if !data.Selectorname.IsNull() {
		streamidentifier.Selectorname = data.Selectorname.ValueString()
	}
	if !data.Snmptrap.IsNull() {
		streamidentifier.Snmptrap = data.Snmptrap.ValueString()
	}
	if !data.Sort.IsNull() {
		streamidentifier.Sort = data.Sort.ValueString()
	}
	if !data.Trackackonlypackets.IsNull() {
		streamidentifier.Trackackonlypackets = data.Trackackonlypackets.ValueString()
	}
	if !data.Tracktransactions.IsNull() {
		streamidentifier.Tracktransactions = data.Tracktransactions.ValueString()
	}

	return streamidentifier
}

func streamidentifierSetAttrFromGet(ctx context.Context, data *StreamidentifierResourceModel, getResponseData map[string]interface{}) *StreamidentifierResourceModel {
	tflog.Debug(ctx, "In streamidentifierSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["acceptancethreshold"]; ok && val != nil {
		data.Acceptancethreshold = types.StringValue(val.(string))
	} else {
		data.Acceptancethreshold = types.StringNull()
	}
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["breachthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Breachthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Breachthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	} else {
		data.Interval = types.Int64Null()
	}
	if val, ok := getResponseData["log"]; ok && val != nil {
		data.Log = types.StringValue(val.(string))
	} else {
		data.Log = types.StringNull()
	}
	if val, ok := getResponseData["loginterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Loginterval = types.Int64Value(intVal)
		}
	} else {
		data.Loginterval = types.Int64Null()
	}
	if val, ok := getResponseData["loglimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Loglimit = types.Int64Value(intVal)
		}
	} else {
		data.Loglimit = types.Int64Null()
	}
	if val, ok := getResponseData["maxtransactionthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxtransactionthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Maxtransactionthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["mintransactionthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mintransactionthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Mintransactionthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["samplecount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Samplecount = types.Int64Value(intVal)
		}
	} else {
		data.Samplecount = types.Int64Null()
	}
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else {
		data.Selectorname = types.StringNull()
	}
	if val, ok := getResponseData["snmptrap"]; ok && val != nil {
		data.Snmptrap = types.StringValue(val.(string))
	} else {
		data.Snmptrap = types.StringNull()
	}
	if val, ok := getResponseData["sort"]; ok && val != nil {
		data.Sort = types.StringValue(val.(string))
	} else {
		data.Sort = types.StringNull()
	}
	if val, ok := getResponseData["trackackonlypackets"]; ok && val != nil {
		data.Trackackonlypackets = types.StringValue(val.(string))
	} else {
		data.Trackackonlypackets = types.StringNull()
	}
	if val, ok := getResponseData["tracktransactions"]; ok && val != nil {
		data.Tracktransactions = types.StringValue(val.(string))
	} else {
		data.Tracktransactions = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
