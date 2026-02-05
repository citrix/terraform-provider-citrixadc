package ipsecparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsec"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IpsecparameterResourceModel describes the resource data model.
type IpsecparameterResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Encalgo               types.List   `tfsdk:"encalgo"`
	Hashalgo              types.List   `tfsdk:"hashalgo"`
	Ikeretryinterval      types.Int64  `tfsdk:"ikeretryinterval"`
	Ikeversion            types.String `tfsdk:"ikeversion"`
	Lifetime              types.Int64  `tfsdk:"lifetime"`
	Livenesscheckinterval types.Int64  `tfsdk:"livenesscheckinterval"`
	Perfectforwardsecrecy types.String `tfsdk:"perfectforwardsecrecy"`
	Replaywindowsize      types.Int64  `tfsdk:"replaywindowsize"`
	Retransmissiontime    types.Int64  `tfsdk:"retransmissiontime"`
}

func (r *IpsecparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipsecparameter resource.",
			},
			"encalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Type of encryption algorithm (Note: Selection of AES enables AES128)",
			},
			"hashalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Type of hashing algorithm",
			},
			"ikeretryinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE retry interval for bringing up the connection",
			},
			"ikeversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("V2"),
				Description: "IKE Protocol Version",
			},
			"lifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8)",
			},
			"livenesscheckinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.",
			},
			"perfectforwardsecrecy": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "Enable/Disable PFS.",
			},
			"replaywindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPSec Replay window size for the data traffic",
			},
			"retransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure,\nincreases for every retransmit till 6 retransmits.",
			},
		},
	}
}

func ipsecparameterGetThePayloadFromtheConfig(ctx context.Context, data *IpsecparameterResourceModel) ipsec.Ipsecparameter {
	tflog.Debug(ctx, "In ipsecparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ipsecparameter := ipsec.Ipsecparameter{}
	if !data.Ikeretryinterval.IsNull() {
		ipsecparameter.Ikeretryinterval = utils.IntPtr(int(data.Ikeretryinterval.ValueInt64()))
	}
	if !data.Ikeversion.IsNull() {
		ipsecparameter.Ikeversion = data.Ikeversion.ValueString()
	}
	if !data.Lifetime.IsNull() {
		ipsecparameter.Lifetime = utils.IntPtr(int(data.Lifetime.ValueInt64()))
	}
	if !data.Livenesscheckinterval.IsNull() {
		ipsecparameter.Livenesscheckinterval = utils.IntPtr(int(data.Livenesscheckinterval.ValueInt64()))
	}
	if !data.Perfectforwardsecrecy.IsNull() {
		ipsecparameter.Perfectforwardsecrecy = data.Perfectforwardsecrecy.ValueString()
	}
	if !data.Replaywindowsize.IsNull() {
		ipsecparameter.Replaywindowsize = utils.IntPtr(int(data.Replaywindowsize.ValueInt64()))
	}
	if !data.Retransmissiontime.IsNull() {
		ipsecparameter.Retransmissiontime = utils.IntPtr(int(data.Retransmissiontime.ValueInt64()))
	}

	return ipsecparameter
}

func ipsecparameterSetAttrFromGet(ctx context.Context, data *IpsecparameterResourceModel, getResponseData map[string]interface{}) *IpsecparameterResourceModel {
	tflog.Debug(ctx, "In ipsecparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ikeretryinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ikeretryinterval = types.Int64Value(intVal)
		}
	} else {
		data.Ikeretryinterval = types.Int64Null()
	}
	if val, ok := getResponseData["ikeversion"]; ok && val != nil {
		data.Ikeversion = types.StringValue(val.(string))
	} else {
		data.Ikeversion = types.StringNull()
	}
	if val, ok := getResponseData["lifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Lifetime = types.Int64Value(intVal)
		}
	} else {
		data.Lifetime = types.Int64Null()
	}
	if val, ok := getResponseData["livenesscheckinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Livenesscheckinterval = types.Int64Value(intVal)
		}
	} else {
		data.Livenesscheckinterval = types.Int64Null()
	}
	if val, ok := getResponseData["perfectforwardsecrecy"]; ok && val != nil {
		data.Perfectforwardsecrecy = types.StringValue(val.(string))
	} else {
		data.Perfectforwardsecrecy = types.StringNull()
	}
	if val, ok := getResponseData["replaywindowsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Replaywindowsize = types.Int64Value(intVal)
		}
	} else {
		data.Replaywindowsize = types.Int64Null()
	}
	if val, ok := getResponseData["retransmissiontime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retransmissiontime = types.Int64Value(intVal)
		}
	} else {
		data.Retransmissiontime = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("ipsecparameter-config")

	return data
}
