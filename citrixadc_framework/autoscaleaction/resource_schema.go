package autoscaleaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AutoscaleactionResourceModel describes the resource data model.
type AutoscaleactionResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Parameters           types.String `tfsdk:"parameters"`
	Profilename          types.String `tfsdk:"profilename"`
	Quiettime            types.Int64  `tfsdk:"quiettime"`
	Type                 types.String `tfsdk:"type"`
	Vmdestroygraceperiod types.Int64  `tfsdk:"vmdestroygraceperiod"`
	Vserver              types.String `tfsdk:"vserver"`
}

func (r *AutoscaleactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the autoscaleaction resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "ActionScale action name.",
			},
			"parameters": schema.StringAttribute{
				Required:    true,
				Description: "Parameters to use in the action",
			},
			"profilename": schema.StringAttribute{
				Required:    true,
				Description: "AutoScale profile name.",
			},
			"quiettime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time in seconds no other policy is evaluated or action is taken",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The type of action.",
			},
			"vmdestroygraceperiod": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Time in minutes a VM is kept in inactive state before destroying",
			},
			"vserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the vserver on which autoscale action has to be taken.",
			},
		},
	}
}

func autoscaleactionGetThePayloadFromtheConfig(ctx context.Context, data *AutoscaleactionResourceModel) autoscale.Autoscaleaction {
	tflog.Debug(ctx, "In autoscaleactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	autoscaleaction := autoscale.Autoscaleaction{}
	if !data.Name.IsNull() {
		autoscaleaction.Name = data.Name.ValueString()
	}
	if !data.Parameters.IsNull() {
		autoscaleaction.Parameters = data.Parameters.ValueString()
	}
	if !data.Profilename.IsNull() {
		autoscaleaction.Profilename = data.Profilename.ValueString()
	}
	if !data.Quiettime.IsNull() {
		autoscaleaction.Quiettime = utils.IntPtr(int(data.Quiettime.ValueInt64()))
	}
	if !data.Type.IsNull() {
		autoscaleaction.Type = data.Type.ValueString()
	}
	if !data.Vmdestroygraceperiod.IsNull() {
		autoscaleaction.Vmdestroygraceperiod = utils.IntPtr(int(data.Vmdestroygraceperiod.ValueInt64()))
	}
	if !data.Vserver.IsNull() {
		autoscaleaction.Vserver = data.Vserver.ValueString()
	}

	return autoscaleaction
}

func autoscaleactionSetAttrFromGet(ctx context.Context, data *AutoscaleactionResourceModel, getResponseData map[string]interface{}) *AutoscaleactionResourceModel {
	tflog.Debug(ctx, "In autoscaleactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["parameters"]; ok && val != nil {
		data.Parameters = types.StringValue(val.(string))
	} else {
		data.Parameters = types.StringNull()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["quiettime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quiettime = types.Int64Value(intVal)
		}
	} else {
		data.Quiettime = types.Int64Null()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vmdestroygraceperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vmdestroygraceperiod = types.Int64Value(intVal)
		}
	} else {
		data.Vmdestroygraceperiod = types.Int64Null()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
