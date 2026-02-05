package lbsipparameters

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbsipparametersResourceModel describes the resource data model.
type LbsipparametersResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Addrportvip         types.String `tfsdk:"addrportvip"`
	Retrydur            types.Int64  `tfsdk:"retrydur"`
	Rnatdstport         types.Int64  `tfsdk:"rnatdstport"`
	Rnatsecuredstport   types.Int64  `tfsdk:"rnatsecuredstport"`
	Rnatsecuresrcport   types.Int64  `tfsdk:"rnatsecuresrcport"`
	Rnatsrcport         types.Int64  `tfsdk:"rnatsrcport"`
	Sip503ratethreshold types.Int64  `tfsdk:"sip503ratethreshold"`
}

func (r *LbsipparametersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbsipparameters resource.",
			},
			"addrportvip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers.",
			},
			"retrydur": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the \"Retry-After\" header in the 503 response.",
			},
			"rnatdstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuredstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuresrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"sip503ratethreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of 503 Service Unavailable responses to generate, once every 10 milliseconds, when a SIP virtual server becomes unavailable.",
			},
		},
	}
}

func lbsipparametersGetThePayloadFromtheConfig(ctx context.Context, data *LbsipparametersResourceModel) lb.Lbsipparameters {
	tflog.Debug(ctx, "In lbsipparametersGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbsipparameters := lb.Lbsipparameters{}
	if !data.Addrportvip.IsNull() {
		lbsipparameters.Addrportvip = data.Addrportvip.ValueString()
	}
	if !data.Retrydur.IsNull() {
		lbsipparameters.Retrydur = utils.IntPtr(int(data.Retrydur.ValueInt64()))
	}
	if !data.Rnatdstport.IsNull() {
		lbsipparameters.Rnatdstport = utils.IntPtr(int(data.Rnatdstport.ValueInt64()))
	}
	if !data.Rnatsecuredstport.IsNull() {
		lbsipparameters.Rnatsecuredstport = utils.IntPtr(int(data.Rnatsecuredstport.ValueInt64()))
	}
	if !data.Rnatsecuresrcport.IsNull() {
		lbsipparameters.Rnatsecuresrcport = utils.IntPtr(int(data.Rnatsecuresrcport.ValueInt64()))
	}
	if !data.Rnatsrcport.IsNull() {
		lbsipparameters.Rnatsrcport = utils.IntPtr(int(data.Rnatsrcport.ValueInt64()))
	}
	if !data.Sip503ratethreshold.IsNull() {
		lbsipparameters.Sip503ratethreshold = utils.IntPtr(int(data.Sip503ratethreshold.ValueInt64()))
	}

	return lbsipparameters
}

func lbsipparametersSetAttrFromGet(ctx context.Context, data *LbsipparametersResourceModel, getResponseData map[string]interface{}) *LbsipparametersResourceModel {
	tflog.Debug(ctx, "In lbsipparametersSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addrportvip"]; ok && val != nil {
		data.Addrportvip = types.StringValue(val.(string))
	} else {
		data.Addrportvip = types.StringNull()
	}
	if val, ok := getResponseData["retrydur"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retrydur = types.Int64Value(intVal)
		}
	} else {
		data.Retrydur = types.Int64Null()
	}
	if val, ok := getResponseData["rnatdstport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rnatdstport = types.Int64Value(intVal)
		}
	} else {
		data.Rnatdstport = types.Int64Null()
	}
	if val, ok := getResponseData["rnatsecuredstport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rnatsecuredstport = types.Int64Value(intVal)
		}
	} else {
		data.Rnatsecuredstport = types.Int64Null()
	}
	if val, ok := getResponseData["rnatsecuresrcport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rnatsecuresrcport = types.Int64Value(intVal)
		}
	} else {
		data.Rnatsecuresrcport = types.Int64Null()
	}
	if val, ok := getResponseData["rnatsrcport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rnatsrcport = types.Int64Value(intVal)
		}
	} else {
		data.Rnatsrcport = types.Int64Null()
	}
	if val, ok := getResponseData["sip503ratethreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sip503ratethreshold = types.Int64Value(intVal)
		}
	} else {
		data.Sip503ratethreshold = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("lbsipparameters-config")

	return data
}
