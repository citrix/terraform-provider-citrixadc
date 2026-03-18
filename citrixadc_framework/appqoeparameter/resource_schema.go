package appqoeparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppqoeparameterResourceModel describes the resource data model.
type AppqoeparameterResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Avgwaitingclient    types.Int64  `tfsdk:"avgwaitingclient"`
	Dosattackthresh     types.Int64  `tfsdk:"dosattackthresh"`
	Maxaltrespbandwidth types.Int64  `tfsdk:"maxaltrespbandwidth"`
	Sessionlife         types.Int64  `tfsdk:"sessionlife"`
}

func (r *AppqoeparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appqoeparameter resource.",
			},
			"avgwaitingclient": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000000),
				Description: "average number of client connections, that can sit in service waiting queue",
			},
			"dosattackthresh": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2000),
				Description: "average number of client connection that can queue up on vserver level without triggering DoS mitigation module",
			},
			"maxaltrespbandwidth": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "maximum bandwidth which will determine whether to send alternate content response",
			},
			"sessionlife": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time, in seconds, between the first time and the next time the AppQoE alternative content window is displayed. The alternative content window is displayed only once during a session for the same browser accessing a configured URL, so this parameter determines the length of a session.",
			},
		},
	}
}

func appqoeparameterGetThePayloadFromtheConfig(ctx context.Context, data *AppqoeparameterResourceModel) appqoe.Appqoeparameter {
	tflog.Debug(ctx, "In appqoeparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appqoeparameter := appqoe.Appqoeparameter{}
	if !data.Avgwaitingclient.IsNull() {
		appqoeparameter.Avgwaitingclient = utils.IntPtr(int(data.Avgwaitingclient.ValueInt64()))
	}
	if !data.Dosattackthresh.IsNull() {
		appqoeparameter.Dosattackthresh = utils.IntPtr(int(data.Dosattackthresh.ValueInt64()))
	}
	if !data.Maxaltrespbandwidth.IsNull() {
		appqoeparameter.Maxaltrespbandwidth = utils.IntPtr(int(data.Maxaltrespbandwidth.ValueInt64()))
	}
	if !data.Sessionlife.IsNull() {
		appqoeparameter.Sessionlife = utils.IntPtr(int(data.Sessionlife.ValueInt64()))
	}

	return appqoeparameter
}

func appqoeparameterSetAttrFromGet(ctx context.Context, data *AppqoeparameterResourceModel, getResponseData map[string]interface{}) *AppqoeparameterResourceModel {
	tflog.Debug(ctx, "In appqoeparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["avgwaitingclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Avgwaitingclient = types.Int64Value(intVal)
		}
	} else {
		data.Avgwaitingclient = types.Int64Null()
	}
	if val, ok := getResponseData["dosattackthresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dosattackthresh = types.Int64Value(intVal)
		}
	} else {
		data.Dosattackthresh = types.Int64Null()
	}
	if val, ok := getResponseData["maxaltrespbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaltrespbandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Maxaltrespbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["sessionlife"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionlife = types.Int64Value(intVal)
		}
	} else {
		data.Sessionlife = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("appqoeparameter-config")

	return data
}
