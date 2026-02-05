package nscqaparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NscqaparamResourceModel describes the resource data model.
type NscqaparamResourceModel struct {
	Id            types.String  `tfsdk:"id"`
	Harqretxdelay types.Int64   `tfsdk:"harqretxdelay"`
	Lr1coeflist   types.String  `tfsdk:"lr1coeflist"`
	Lr1probthresh types.Float64 `tfsdk:"lr1probthresh"`
	Lr2coeflist   types.String  `tfsdk:"lr2coeflist"`
	Lr2probthresh types.Float64 `tfsdk:"lr2probthresh"`
	Minrttnet1    types.Int64   `tfsdk:"minrttnet1"`
	Minrttnet2    types.Int64   `tfsdk:"minrttnet2"`
	Minrttnet3    types.Int64   `tfsdk:"minrttnet3"`
	Net1cclscale  types.String  `tfsdk:"net1cclscale"`
	Net1csqscale  types.String  `tfsdk:"net1csqscale"`
	Net1label     types.String  `tfsdk:"net1label"`
	Net1logcoef   types.String  `tfsdk:"net1logcoef"`
	Net2cclscale  types.String  `tfsdk:"net2cclscale"`
	Net2csqscale  types.String  `tfsdk:"net2csqscale"`
	Net2label     types.String  `tfsdk:"net2label"`
	Net2logcoef   types.String  `tfsdk:"net2logcoef"`
	Net3cclscale  types.String  `tfsdk:"net3cclscale"`
	Net3csqscale  types.String  `tfsdk:"net3csqscale"`
	Net3label     types.String  `tfsdk:"net3label"`
	Net3logcoef   types.String  `tfsdk:"net3logcoef"`
}

func (r *NscqaparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nscqaparam resource.",
			},
			"harqretxdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HARQ retransmission delay (in ms).",
			},
			"lr1coeflist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "coefficients values for Label1.",
			},
			"lr1probthresh": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Probability threshold values for LR model to differentiate between NET1 and reset(NET2 and NET3).",
			},
			"lr2coeflist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "coefficients values for Label 2.",
			},
			"lr2probthresh": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Probability threshold values for LR model to differentiate between NET2 and NET3.",
			},
			"minrttnet1": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the first network.",
			},
			"minrttnet2": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the second network.",
			},
			"minrttnet3": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MIN RTT (in ms) for the third network.",
			},
			"net1cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net1csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net1label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label.",
			},
			"net1logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection quality ranking Log coefficients of network 1.",
			},
			"net2cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net2csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net2label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label 2.",
			},
			"net2logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connnection quality ranking Log coefficients of network 2.",
			},
			"net3cclscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three congestion level scores limits corresponding to None, Low, Medium.",
			},
			"net3csqscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three signal quality level scores limits corresponding to Excellent, Good, Fair.",
			},
			"net3label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network label 3.",
			},
			"net3logcoef": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection quality ranking Log coefficients of network 3.",
			},
		},
	}
}

func nscqaparamGetThePayloadFromtheConfig(ctx context.Context, data *NscqaparamResourceModel) ns.Nscqaparam {
	tflog.Debug(ctx, "In nscqaparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nscqaparam := ns.Nscqaparam{}
	if !data.Harqretxdelay.IsNull() {
		nscqaparam.Harqretxdelay = utils.IntPtr(int(data.Harqretxdelay.ValueInt64()))
	}
	if !data.Lr1coeflist.IsNull() {
		nscqaparam.Lr1coeflist = data.Lr1coeflist.ValueString()
	}
	if !data.Lr1probthresh.IsNull() {
		nscqaparam.Lr1probthresh = data.Lr1probthresh.ValueFloat64()
	}
	if !data.Lr2coeflist.IsNull() {
		nscqaparam.Lr2coeflist = data.Lr2coeflist.ValueString()
	}
	if !data.Lr2probthresh.IsNull() {
		nscqaparam.Lr2probthresh = data.Lr2probthresh.ValueFloat64()
	}
	if !data.Minrttnet1.IsNull() {
		nscqaparam.Minrttnet1 = utils.IntPtr(int(data.Minrttnet1.ValueInt64()))
	}
	if !data.Minrttnet2.IsNull() {
		nscqaparam.Minrttnet2 = utils.IntPtr(int(data.Minrttnet2.ValueInt64()))
	}
	if !data.Minrttnet3.IsNull() {
		nscqaparam.Minrttnet3 = utils.IntPtr(int(data.Minrttnet3.ValueInt64()))
	}
	if !data.Net1cclscale.IsNull() {
		nscqaparam.Net1cclscale = data.Net1cclscale.ValueString()
	}
	if !data.Net1csqscale.IsNull() {
		nscqaparam.Net1csqscale = data.Net1csqscale.ValueString()
	}
	if !data.Net1label.IsNull() {
		nscqaparam.Net1label = data.Net1label.ValueString()
	}
	if !data.Net1logcoef.IsNull() {
		nscqaparam.Net1logcoef = data.Net1logcoef.ValueString()
	}
	if !data.Net2cclscale.IsNull() {
		nscqaparam.Net2cclscale = data.Net2cclscale.ValueString()
	}
	if !data.Net2csqscale.IsNull() {
		nscqaparam.Net2csqscale = data.Net2csqscale.ValueString()
	}
	if !data.Net2label.IsNull() {
		nscqaparam.Net2label = data.Net2label.ValueString()
	}
	if !data.Net2logcoef.IsNull() {
		nscqaparam.Net2logcoef = data.Net2logcoef.ValueString()
	}
	if !data.Net3cclscale.IsNull() {
		nscqaparam.Net3cclscale = data.Net3cclscale.ValueString()
	}
	if !data.Net3csqscale.IsNull() {
		nscqaparam.Net3csqscale = data.Net3csqscale.ValueString()
	}
	if !data.Net3label.IsNull() {
		nscqaparam.Net3label = data.Net3label.ValueString()
	}
	if !data.Net3logcoef.IsNull() {
		nscqaparam.Net3logcoef = data.Net3logcoef.ValueString()
	}

	return nscqaparam
}

func nscqaparamSetAttrFromGet(ctx context.Context, data *NscqaparamResourceModel, getResponseData map[string]interface{}) *NscqaparamResourceModel {
	tflog.Debug(ctx, "In nscqaparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["harqretxdelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Harqretxdelay = types.Int64Value(intVal)
		}
	} else {
		data.Harqretxdelay = types.Int64Null()
	}
	if val, ok := getResponseData["lr1coeflist"]; ok && val != nil {
		data.Lr1coeflist = types.StringValue(val.(string))
	} else {
		data.Lr1coeflist = types.StringNull()
	}
	if val, ok := getResponseData["lr1probthresh"]; ok && val != nil {
		data.Lr1probthresh = types.Float64Value(val.(float64))
	} else {
		data.Lr1probthresh = types.Float64Null()
	}
	if val, ok := getResponseData["lr2coeflist"]; ok && val != nil {
		data.Lr2coeflist = types.StringValue(val.(string))
	} else {
		data.Lr2coeflist = types.StringNull()
	}
	if val, ok := getResponseData["lr2probthresh"]; ok && val != nil {
		data.Lr2probthresh = types.Float64Value(val.(float64))
	} else {
		data.Lr2probthresh = types.Float64Null()
	}
	if val, ok := getResponseData["minrttnet1"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrttnet1 = types.Int64Value(intVal)
		}
	} else {
		data.Minrttnet1 = types.Int64Null()
	}
	if val, ok := getResponseData["minrttnet2"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrttnet2 = types.Int64Value(intVal)
		}
	} else {
		data.Minrttnet2 = types.Int64Null()
	}
	if val, ok := getResponseData["minrttnet3"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minrttnet3 = types.Int64Value(intVal)
		}
	} else {
		data.Minrttnet3 = types.Int64Null()
	}
	if val, ok := getResponseData["net1cclscale"]; ok && val != nil {
		data.Net1cclscale = types.StringValue(val.(string))
	} else {
		data.Net1cclscale = types.StringNull()
	}
	if val, ok := getResponseData["net1csqscale"]; ok && val != nil {
		data.Net1csqscale = types.StringValue(val.(string))
	} else {
		data.Net1csqscale = types.StringNull()
	}
	if val, ok := getResponseData["net1label"]; ok && val != nil {
		data.Net1label = types.StringValue(val.(string))
	} else {
		data.Net1label = types.StringNull()
	}
	if val, ok := getResponseData["net1logcoef"]; ok && val != nil {
		data.Net1logcoef = types.StringValue(val.(string))
	} else {
		data.Net1logcoef = types.StringNull()
	}
	if val, ok := getResponseData["net2cclscale"]; ok && val != nil {
		data.Net2cclscale = types.StringValue(val.(string))
	} else {
		data.Net2cclscale = types.StringNull()
	}
	if val, ok := getResponseData["net2csqscale"]; ok && val != nil {
		data.Net2csqscale = types.StringValue(val.(string))
	} else {
		data.Net2csqscale = types.StringNull()
	}
	if val, ok := getResponseData["net2label"]; ok && val != nil {
		data.Net2label = types.StringValue(val.(string))
	} else {
		data.Net2label = types.StringNull()
	}
	if val, ok := getResponseData["net2logcoef"]; ok && val != nil {
		data.Net2logcoef = types.StringValue(val.(string))
	} else {
		data.Net2logcoef = types.StringNull()
	}
	if val, ok := getResponseData["net3cclscale"]; ok && val != nil {
		data.Net3cclscale = types.StringValue(val.(string))
	} else {
		data.Net3cclscale = types.StringNull()
	}
	if val, ok := getResponseData["net3csqscale"]; ok && val != nil {
		data.Net3csqscale = types.StringValue(val.(string))
	} else {
		data.Net3csqscale = types.StringNull()
	}
	if val, ok := getResponseData["net3label"]; ok && val != nil {
		data.Net3label = types.StringValue(val.(string))
	} else {
		data.Net3label = types.StringNull()
	}
	if val, ok := getResponseData["net3logcoef"]; ok && val != nil {
		data.Net3logcoef = types.StringValue(val.(string))
	} else {
		data.Net3logcoef = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nscqaparam-config")

	return data
}
