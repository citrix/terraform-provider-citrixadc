package sslvserver_sslpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslvserverSslpolicyBindingResourceModel describes the resource data model.
type SslvserverSslpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
	Vservername            types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke flag. This attribute is relevant only for ADVANCED policies",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the SSL policy binding.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policies bound to this SSL service",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("REQUEST"),
				Description: "Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean:\n1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert.\n2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET.\n3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslvserverSslpolicyBindingResourceModel) ssl.Sslvserversslpolicybinding {
	tflog.Debug(ctx, "In sslvserver_sslpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslvserver_sslpolicy_binding := ssl.Sslvserversslpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		sslvserver_sslpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		sslvserver_sslpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() {
		sslvserver_sslpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		sslvserver_sslpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		sslvserver_sslpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		sslvserver_sslpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		sslvserver_sslpolicy_binding.Type = data.Type.ValueString()
	}
	if !data.Vservername.IsNull() {
		sslvserver_sslpolicy_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslpolicy_binding
}

func sslvserver_sslpolicy_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslpolicyBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslpolicyBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
