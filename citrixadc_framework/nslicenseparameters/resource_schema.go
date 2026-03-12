package nslicenseparameters

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslicenseparametersResourceModel describes the resource data model.
type NslicenseparametersResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Alert1gracetimeout       types.Int64  `tfsdk:"alert1gracetimeout"`
	Alert2gracetimeout       types.Int64  `tfsdk:"alert2gracetimeout"`
	Heartbeatinterval        types.Int64  `tfsdk:"heartbeatinterval"`
	Inventoryrefreshinterval types.Int64  `tfsdk:"inventoryrefreshinterval"`
	Licenseexpiryalerttime   types.Int64  `tfsdk:"licenseexpiryalerttime"`
}

func (r *NslicenseparametersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicenseparameters resource.",
			},
			"alert1gracetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(6),
				Description: "If ADC remains in grace for the configured hours then first grace alert will be raised",
			},
			"alert2gracetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(240),
				Description: "If ADC remains in grace for the configured hours then major grace alert will be raised",
			},
			"heartbeatinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(280),
				Description: "Heartbeat between ADC and Licenseserver is configurable and applicable in case of pooled licensing",
			},
			"inventoryrefreshinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(360),
				Description: "Inventory refresh interval between ADC and Licenseserver is configurable and applicable in case of pooled licensing",
			},
			"licenseexpiryalerttime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "If ADC license contract expiry date is nearer then GUI/SNMP license expiry alert will be raised",
			},
		},
	}
}

func nslicenseparametersGetThePayloadFromtheConfig(ctx context.Context, data *NslicenseparametersResourceModel) ns.Nslicenseparameters {
	tflog.Debug(ctx, "In nslicenseparametersGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nslicenseparameters := ns.Nslicenseparameters{}
	if !data.Alert1gracetimeout.IsNull() {
		nslicenseparameters.Alert1gracetimeout = utils.IntPtr(int(data.Alert1gracetimeout.ValueInt64()))
	}
	if !data.Alert2gracetimeout.IsNull() {
		nslicenseparameters.Alert2gracetimeout = utils.IntPtr(int(data.Alert2gracetimeout.ValueInt64()))
	}
	if !data.Heartbeatinterval.IsNull() {
		nslicenseparameters.Heartbeatinterval = utils.IntPtr(int(data.Heartbeatinterval.ValueInt64()))
	}
	if !data.Inventoryrefreshinterval.IsNull() {
		nslicenseparameters.Inventoryrefreshinterval = utils.IntPtr(int(data.Inventoryrefreshinterval.ValueInt64()))
	}
	if !data.Licenseexpiryalerttime.IsNull() {
		nslicenseparameters.Licenseexpiryalerttime = utils.IntPtr(int(data.Licenseexpiryalerttime.ValueInt64()))
	}

	return nslicenseparameters
}

func nslicenseparametersSetAttrFromGet(ctx context.Context, data *NslicenseparametersResourceModel, getResponseData map[string]interface{}) *NslicenseparametersResourceModel {
	tflog.Debug(ctx, "In nslicenseparametersSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alert1gracetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Alert1gracetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Alert1gracetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["alert2gracetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Alert2gracetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Alert2gracetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["heartbeatinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Heartbeatinterval = types.Int64Value(intVal)
		}
	} else {
		data.Heartbeatinterval = types.Int64Null()
	}
	if val, ok := getResponseData["inventoryrefreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Inventoryrefreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Inventoryrefreshinterval = types.Int64Null()
	}
	if val, ok := getResponseData["licenseexpiryalerttime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Licenseexpiryalerttime = types.Int64Value(intVal)
		}
	} else {
		data.Licenseexpiryalerttime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nslicenseparameters-config")

	return data
}
