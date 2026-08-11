package nslimitidentifier

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

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

// NslimitidentifierResourceModel describes the resource data model.
type NslimitidentifierResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Limitidentifier  types.String `tfsdk:"limitidentifier"`
	Limittype        types.String `tfsdk:"limittype"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Mode             types.String `tfsdk:"mode"`
	Selectorname     types.String `tfsdk:"selectorname"`
	Threshold        types.Int64  `tfsdk:"threshold"`
	Timeslice        types.Int64  `tfsdk:"timeslice"`
	Trapsintimeslice types.Int64  `tfsdk:"trapsintimeslice"`
}

func (r *NslimitidentifierResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslimitidentifier resource.",
			},
			"limitidentifier": schema.StringAttribute{
				Required: true,
				// SDK v2 ForceNew: true -> RequiresReplace
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for a rate limit identifier. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Reserved words must not be used.",
			},
			"limittype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("BURSTY"),
				Description: "Smooth or bursty request type.\n* SMOOTH - When you want the permitted number of requests in a given interval of time to be spread evenly across the timeslice\n* BURSTY - When you want the permitted number of requests to exhaust the quota anytime within the timeslice.\nThis argument is needed only when the mode is set to REQUEST_RATE.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
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
				Default:     int64default.StaticInt64(1),
				Description: "Maximum number of requests that are allowed in the given timeslice when requests (mode is set as REQUEST_RATE) are tracked per timeslice.\nWhen connections (mode is set as CONNECTION) are tracked, it is the total number of connections that would be let through.",
			},
			"timeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Time interval, in milliseconds, specified in multiples of 10, during which requests are tracked to check if they cross the threshold. This argument is needed only when the mode is set to REQUEST_RATE.",
			},
			"trapsintimeslice": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Number of traps to be sent in the timeslice configured. A value of 0 indicates that traps are disabled.",
			},
		},
	}
}

func nslimitidentifierGetThePayloadFromthePlan(ctx context.Context, data *NslimitidentifierResourceModel) ns.Nslimitidentifier {
	tflog.Debug(ctx, "In nslimitidentifierGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nslimitidentifier := ns.Nslimitidentifier{}
	if !data.Limitidentifier.IsNull() && !data.Limitidentifier.IsUnknown() {
		nslimitidentifier.Limitidentifier = data.Limitidentifier.ValueString()
	}
	if !data.Limittype.IsNull() && !data.Limittype.IsUnknown() {
		nslimitidentifier.Limittype = data.Limittype.ValueString()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		nslimitidentifier.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		nslimitidentifier.Mode = data.Mode.ValueString()
	}
	if !data.Selectorname.IsNull() && !data.Selectorname.IsUnknown() {
		nslimitidentifier.Selectorname = data.Selectorname.ValueString()
	}
	if !data.Threshold.IsNull() && !data.Threshold.IsUnknown() {
		nslimitidentifier.Threshold = utils.IntPtr(int(data.Threshold.ValueInt64()))
	}
	if !data.Timeslice.IsNull() && !data.Timeslice.IsUnknown() {
		nslimitidentifier.Timeslice = utils.IntPtr(int(data.Timeslice.ValueInt64()))
	}
	if !data.Trapsintimeslice.IsNull() && !data.Trapsintimeslice.IsUnknown() {
		nslimitidentifier.Trapsintimeslice = utils.IntPtr(int(data.Trapsintimeslice.ValueInt64()))
	}

	return nslimitidentifier
}

// nslimitidentifierSetAttrFromGet maps the NITRO GET response onto the resource
// state. Omit-on-default guard: when NITRO omits an attribute from GET (e.g. a
// configured 0 / empty value that the appliance does not echo back), do NOT
// clobber a known configured/state value — only null it when it is still Unknown.
func nslimitidentifierSetAttrFromGet(ctx context.Context, data *NslimitidentifierResourceModel, getResponseData map[string]interface{}) *NslimitidentifierResourceModel {
	tflog.Debug(ctx, "In nslimitidentifierSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["limitidentifier"]; ok && val != nil {
		data.Limitidentifier = types.StringValue(val.(string))
	} else if data.Limitidentifier.IsUnknown() {
		data.Limitidentifier = types.StringNull()
	}
	if val, ok := getResponseData["limittype"]; ok && val != nil {
		data.Limittype = types.StringValue(val.(string))
	} else if data.Limittype.IsUnknown() {
		data.Limittype = types.StringNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else if data.Maxbandwidth.IsUnknown() {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else if data.Mode.IsUnknown() {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else if data.Selectorname.IsUnknown() {
		data.Selectorname = types.StringNull()
	}
	if val, ok := getResponseData["threshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Threshold = types.Int64Value(intVal)
		}
	} else if data.Threshold.IsUnknown() {
		data.Threshold = types.Int64Null()
	}
	if val, ok := getResponseData["timeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeslice = types.Int64Value(intVal)
		}
	} else if data.Timeslice.IsUnknown() {
		data.Timeslice = types.Int64Null()
	}
	if val, ok := getResponseData["trapsintimeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trapsintimeslice = types.Int64Value(intVal)
		}
	} else if data.Trapsintimeslice.IsUnknown() {
		data.Trapsintimeslice = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID (matches SDK v2 d.SetId(limitidentifier))
	data.Id = types.StringValue(data.Limitidentifier.ValueString())

	return data
}

// nslimitidentifierSetAttrFromGetForDatasource maps the NITRO GET response onto
// the shared model for the datasource. Unlike the resource setter, it always
// mirrors the GET response (copy-all, null-on-absent) and sets the ID.
func nslimitidentifierSetAttrFromGetForDatasource(ctx context.Context, data *NslimitidentifierResourceModel, getResponseData map[string]interface{}) *NslimitidentifierResourceModel {
	tflog.Debug(ctx, "In nslimitidentifierSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["limitidentifier"]; ok && val != nil {
		data.Limitidentifier = types.StringValue(val.(string))
	} else {
		data.Limitidentifier = types.StringNull()
	}
	if val, ok := getResponseData["limittype"]; ok && val != nil {
		data.Limittype = types.StringValue(val.(string))
	} else {
		data.Limittype = types.StringNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		} else {
			data.Maxbandwidth = types.Int64Null()
		}
	} else {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else {
		data.Selectorname = types.StringNull()
	}
	if val, ok := getResponseData["threshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Threshold = types.Int64Value(intVal)
		} else {
			data.Threshold = types.Int64Null()
		}
	} else {
		data.Threshold = types.Int64Null()
	}
	if val, ok := getResponseData["timeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeslice = types.Int64Value(intVal)
		} else {
			data.Timeslice = types.Int64Null()
		}
	} else {
		data.Timeslice = types.Int64Null()
	}
	if val, ok := getResponseData["trapsintimeslice"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trapsintimeslice = types.Int64Value(intVal)
		} else {
			data.Trapsintimeslice = types.Int64Null()
		}
	} else {
		data.Trapsintimeslice = types.Int64Null()
	}

	data.Id = types.StringValue(data.Limitidentifier.ValueString())

	return data
}
