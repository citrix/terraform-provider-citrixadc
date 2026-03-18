package nsratecontrol

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

// NsratecontrolResourceModel describes the resource data model.
type NsratecontrolResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Icmpthreshold   types.Int64  `tfsdk:"icmpthreshold"`
	Tcprstthreshold types.Int64  `tfsdk:"tcprstthreshold"`
	Tcpthreshold    types.Int64  `tfsdk:"tcpthreshold"`
	Udpthreshold    types.Int64  `tfsdk:"udpthreshold"`
}

func (r *NsratecontrolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsratecontrol resource.",
			},
			"icmpthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Number of ICMP packets permitted per 10 milliseconds.",
			},
			"tcprstthreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "The number of TCP RST packets permitted per 10 milli second. zero means rate control is disabled and 0xffffffff means every thing is rate controlled",
			},
			"tcpthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of SYNs permitted per 10 milliseconds.",
			},
			"udpthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of UDP packets permitted per 10 milliseconds.",
			},
		},
	}
}

func nsratecontrolGetThePayloadFromtheConfig(ctx context.Context, data *NsratecontrolResourceModel) ns.Nsratecontrol {
	tflog.Debug(ctx, "In nsratecontrolGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsratecontrol := ns.Nsratecontrol{}
	if !data.Icmpthreshold.IsNull() {
		nsratecontrol.Icmpthreshold = utils.IntPtr(int(data.Icmpthreshold.ValueInt64()))
	}
	if !data.Tcprstthreshold.IsNull() {
		nsratecontrol.Tcprstthreshold = utils.IntPtr(int(data.Tcprstthreshold.ValueInt64()))
	}
	if !data.Tcpthreshold.IsNull() {
		nsratecontrol.Tcpthreshold = utils.IntPtr(int(data.Tcpthreshold.ValueInt64()))
	}
	if !data.Udpthreshold.IsNull() {
		nsratecontrol.Udpthreshold = utils.IntPtr(int(data.Udpthreshold.ValueInt64()))
	}

	return nsratecontrol
}

func nsratecontrolSetAttrFromGet(ctx context.Context, data *NsratecontrolResourceModel, getResponseData map[string]interface{}) *NsratecontrolResourceModel {
	tflog.Debug(ctx, "In nsratecontrolSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["icmpthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Icmpthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Icmpthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["tcprstthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcprstthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Tcprstthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["tcpthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Tcpthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["udpthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Udpthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Udpthreshold = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsratecontrol-config")

	return data
}
