package icalatencyprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IcalatencyprofileResourceModel describes the resource data model.
type IcalatencyprofileResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	L7latencymaxnotifycount  types.Int64  `tfsdk:"l7latencymaxnotifycount"`
	L7latencymonitoring      types.String `tfsdk:"l7latencymonitoring"`
	L7latencynotifyinterval  types.Int64  `tfsdk:"l7latencynotifyinterval"`
	L7latencythresholdfactor types.Int64  `tfsdk:"l7latencythresholdfactor"`
	L7latencywaittime        types.Int64  `tfsdk:"l7latencywaittime"`
	Name                     types.String `tfsdk:"name"`
}

func (r *IcalatencyprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icalatencyprofile resource.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica l7latencyprofile\" or 'my ica l7latencyprofile').",
			},
		},
	}
}

func icalatencyprofileGetThePayloadFromthePlan(ctx context.Context, data *IcalatencyprofileResourceModel) ica.Icalatencyprofile {
	tflog.Debug(ctx, "In icalatencyprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	icalatencyprofile := ica.Icalatencyprofile{}
	if !data.L7latencymaxnotifycount.IsNull() && !data.L7latencymaxnotifycount.IsUnknown() {
		icalatencyprofile.L7latencymaxnotifycount = utils.IntPtr(int(data.L7latencymaxnotifycount.ValueInt64()))
	}
	if !data.L7latencymonitoring.IsNull() && !data.L7latencymonitoring.IsUnknown() {
		icalatencyprofile.L7latencymonitoring = data.L7latencymonitoring.ValueString()
	}
	if !data.L7latencynotifyinterval.IsNull() && !data.L7latencynotifyinterval.IsUnknown() {
		icalatencyprofile.L7latencynotifyinterval = utils.IntPtr(int(data.L7latencynotifyinterval.ValueInt64()))
	}
	if !data.L7latencythresholdfactor.IsNull() && !data.L7latencythresholdfactor.IsUnknown() {
		icalatencyprofile.L7latencythresholdfactor = utils.IntPtr(int(data.L7latencythresholdfactor.ValueInt64()))
	}
	if !data.L7latencywaittime.IsNull() && !data.L7latencywaittime.IsUnknown() {
		icalatencyprofile.L7latencywaittime = utils.IntPtr(int(data.L7latencywaittime.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		icalatencyprofile.Name = data.Name.ValueString()
	}

	return icalatencyprofile
}

func icalatencyprofileSetAttrFromGet(ctx context.Context, data *IcalatencyprofileResourceModel, getResponseData map[string]interface{}) *IcalatencyprofileResourceModel {
	tflog.Debug(ctx, "In icalatencyprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["l7latencymaxnotifycount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.L7latencymaxnotifycount = types.Int64Value(intVal)
		}
	} else if data.L7latencymaxnotifycount.IsUnknown() {
		data.L7latencymaxnotifycount = types.Int64Null()
	}
	if val, ok := getResponseData["l7latencymonitoring"]; ok && val != nil {
		data.L7latencymonitoring = types.StringValue(val.(string))
	} else if data.L7latencymonitoring.IsUnknown() {
		data.L7latencymonitoring = types.StringNull()
	}
	if val, ok := getResponseData["l7latencynotifyinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.L7latencynotifyinterval = types.Int64Value(intVal)
		}
	} else if data.L7latencynotifyinterval.IsUnknown() {
		data.L7latencynotifyinterval = types.Int64Null()
	}
	if val, ok := getResponseData["l7latencythresholdfactor"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.L7latencythresholdfactor = types.Int64Value(intVal)
		}
	} else if data.L7latencythresholdfactor.IsUnknown() {
		data.L7latencythresholdfactor = types.Int64Null()
	}
	if val, ok := getResponseData["l7latencywaittime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.L7latencywaittime = types.Int64Value(intVal)
		}
	} else if data.L7latencywaittime.IsUnknown() {
		data.L7latencywaittime = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value (name) as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
