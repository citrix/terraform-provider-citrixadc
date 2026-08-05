package nstimer

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstimerResourceModel describes the resource data model.
type NstimerResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Comment  types.String `tfsdk:"comment"`
	Interval types.Int64  `tfsdk:"interval"`
	Name     types.String `tfsdk:"name"`
	Newname  types.String `tfsdk:"newname"`
	Unit     types.String `tfsdk:"unit"`
}

func (r *NstimerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstimer resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed, no Default (value read from ADC).
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this timer.",
			},
			"interval": schema.Int64Attribute{
				// SDK v2 parity: Required with NO Default (a Default is invalid without
				// Computed, and Required makes a default meaningless anyway).
				Required:    true,
				Description: "The frequency at which the policies bound to this timer are invoked. The minimum value is 20 msec. The maximum value is 20940 in seconds and 349 in minutes",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timer name.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It must be
				// Optional-only: NOT Computed (pure user input, never echoed by GET) and
				// NOT RequiresReplace (a change drives an in-place rename via Update).
				Optional:    true,
				Description: "The new name of the timer.",
			},
			"unit": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed, no Default (value read from ADC).
				Optional:    true,
				Computed:    true,
				Description: "Timer interval unit",
			},
		},
	}
}

func nstimerGetThePayloadFromthePlan(ctx context.Context, data *NstimerResourceModel) ns.Nstimer {
	tflog.Debug(ctx, "In nstimerGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nstimer := ns.Nstimer{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nstimer.Comment = data.Comment.ValueString()
	}
	if !data.Interval.IsNull() && !data.Interval.IsUnknown() {
		nstimer.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nstimer.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add/update payload, so it is deliberately excluded from the create/update body.
	if !data.Unit.IsNull() && !data.Unit.IsUnknown() {
		nstimer.Unit = data.Unit.ValueString()
	}

	return nstimer
}

// nstimerSetAttrFromGet is the RESOURCE-side setter. It overlays the GET response
// onto the existing plan/state model, preserving user-configured values that NITRO
// omits from GET (omit-on-default trap) and never touching the ID (Create/Update own
// the ID, which is critical for the rename lifecycle).
func nstimerSetAttrFromGet(ctx context.Context, data *NstimerResourceModel, getResponseData map[string]interface{}) *NstimerResourceModel {
	tflog.Debug(ctx, "In nstimerSetAttrFromGet Function")

	// Convert API response to model. else-branch nulls ONLY when the value is unknown
	// (a Computed attribute the user omitted) so we never clobber a known configured value.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	} else if data.Interval.IsUnknown() {
		data.Interval = types.Int64Null()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name and GET returns the live
	// name; overwriting name from GET would clobber the user's value and cause a
	// spurious RequiresReplace diff. Only adopt the GET value when we don't already
	// have one (e.g. import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["unit"]; ok && val != nil {
		data.Unit = types.StringValue(val.(string))
	} else if data.Unit.IsUnknown() {
		data.Unit = types.StringNull()
	}

	return data
}

// nstimerSetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it must populate
// the model directly from the API response and set the ID itself.
func nstimerSetAttrFromGetForDatasource(ctx context.Context, data *NstimerResourceModel, getResponseData map[string]interface{}) *NstimerResourceModel {
	tflog.Debug(ctx, "In nstimerSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		} else {
			data.Interval = types.Int64Null()
		}
	} else {
		data.Interval = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["unit"]; ok && val != nil {
		data.Unit = types.StringValue(val.(string))
	} else {
		data.Unit = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
