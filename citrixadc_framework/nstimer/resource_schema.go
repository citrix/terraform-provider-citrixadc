package nstimer

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
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this timer.",
			},
			"interval": schema.Int64Attribute{
				Required:    true,
				Default:     int64default.StaticInt64(5),
				Description: "The frequency at which the policies bound to this timer are invoked. The minimum value is 20 msec. The maximum value is 20940 in seconds and 349 in minutes",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Timer name.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the timer.",
			},
			"unit": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SEC"),
				Description: "Timer interval unit",
			},
		},
	}
}

func nstimerGetThePayloadFromtheConfig(ctx context.Context, data *NstimerResourceModel) ns.Nstimer {
	tflog.Debug(ctx, "In nstimerGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstimer := ns.Nstimer{}
	if !data.Comment.IsNull() {
		nstimer.Comment = data.Comment.ValueString()
	}
	if !data.Interval.IsNull() {
		nstimer.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Name.IsNull() {
		nstimer.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		nstimer.Newname = data.Newname.ValueString()
	}
	if !data.Unit.IsNull() {
		nstimer.Unit = data.Unit.ValueString()
	}

	return nstimer
}

func nstimerSetAttrFromGet(ctx context.Context, data *NstimerResourceModel, getResponseData map[string]interface{}) *NstimerResourceModel {
	tflog.Debug(ctx, "In nstimerSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
