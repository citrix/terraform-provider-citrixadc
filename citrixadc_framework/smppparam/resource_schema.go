package smppparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/smpp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SmppparamResourceModel describes the resource data model.
type SmppparamResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Addrnpi      types.Int64  `tfsdk:"addrnpi"`
	Addrrange    types.String `tfsdk:"addrrange"`
	Addrton      types.Int64  `tfsdk:"addrton"`
	Clientmode   types.String `tfsdk:"clientmode"`
	Msgqueue     types.String `tfsdk:"msgqueue"`
	Msgqueuesize types.Int64  `tfsdk:"msgqueuesize"`
}

func (r *SmppparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the smppparam resource.",
			},
			"addrnpi": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numbering Plan Indicator, such as landline, data, or WAP client, used in the ESME address sent in the bind request.",
			},
			"addrrange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("\\d*"),
				Description: "Set of SME addresses, sent in the bind request, serviced by the ESME.",
			},
			"addrton": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of Number, such as an international number or a national number, used in the ESME address sent in the bind request.",
			},
			"clientmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("TRANSCEIVER"),
				Description: "Mode in which the client binds to the ADC. Applicable settings function as follows:\n* TRANSCEIVER - Client can send and receive messages to and from the message center.\n* TRANSMITTERONLY - Client can only send messages.\n* RECEIVERONLY - Client can only receive messages.",
			},
			"msgqueue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Queue SMPP messages if a client that is capable of receiving the destination address messages is not available.",
			},
			"msgqueuesize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10000),
				Description: "Maximum number of SMPP messages that can be queued. After the limit is reached, the Citrix ADC sends a deliver_sm_resp PDU, with an appropriate error message, to the message center.",
			},
		},
	}
}

func smppparamGetThePayloadFromtheConfig(ctx context.Context, data *SmppparamResourceModel) smpp.Smppparam {
	tflog.Debug(ctx, "In smppparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	smppparam := smpp.Smppparam{}
	if !data.Addrnpi.IsNull() {
		smppparam.Addrnpi = utils.IntPtr(int(data.Addrnpi.ValueInt64()))
	}
	if !data.Addrrange.IsNull() {
		smppparam.Addrrange = data.Addrrange.ValueString()
	}
	if !data.Addrton.IsNull() {
		smppparam.Addrton = utils.IntPtr(int(data.Addrton.ValueInt64()))
	}
	if !data.Clientmode.IsNull() {
		smppparam.Clientmode = data.Clientmode.ValueString()
	}
	if !data.Msgqueue.IsNull() {
		smppparam.Msgqueue = data.Msgqueue.ValueString()
	}
	if !data.Msgqueuesize.IsNull() {
		smppparam.Msgqueuesize = utils.IntPtr(int(data.Msgqueuesize.ValueInt64()))
	}

	return smppparam
}

func smppparamSetAttrFromGet(ctx context.Context, data *SmppparamResourceModel, getResponseData map[string]interface{}) *SmppparamResourceModel {
	tflog.Debug(ctx, "In smppparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addrnpi"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Addrnpi = types.Int64Value(intVal)
		}
	} else {
		data.Addrnpi = types.Int64Null()
	}
	if val, ok := getResponseData["addrrange"]; ok && val != nil {
		data.Addrrange = types.StringValue(val.(string))
	} else {
		data.Addrrange = types.StringNull()
	}
	if val, ok := getResponseData["addrton"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Addrton = types.Int64Value(intVal)
		}
	} else {
		data.Addrton = types.Int64Null()
	}
	if val, ok := getResponseData["clientmode"]; ok && val != nil {
		data.Clientmode = types.StringValue(val.(string))
	} else {
		data.Clientmode = types.StringNull()
	}
	if val, ok := getResponseData["msgqueue"]; ok && val != nil {
		data.Msgqueue = types.StringValue(val.(string))
	} else {
		data.Msgqueue = types.StringNull()
	}
	if val, ok := getResponseData["msgqueuesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Msgqueuesize = types.Int64Value(intVal)
		}
	} else {
		data.Msgqueuesize = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("smppparam-config")

	return data
}
