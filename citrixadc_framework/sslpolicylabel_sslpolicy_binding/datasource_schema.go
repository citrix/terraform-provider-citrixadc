package sslpolicylabel_sslpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslpolicylabelSslpolicyBindingDataSourceModel describes the datasource data model.
// It mirrors the resource model but exposes the invoke label name under the
// datasource-facing attribute name "invoke_labelname".
type SslpolicylabelSslpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func SslpolicylabelSslpolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL policy label to which to bind policies.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

// sslpolicylabel_sslpolicy_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response into the datasource model and sets the composite ID.
func sslpolicylabel_sslpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslpolicylabelSslpolicyBindingDataSourceModel, getResponseData map[string]interface{}) *SslpolicylabelSslpolicyBindingDataSourceModel {
	tflog.Debug(ctx, "In sslpolicylabel_sslpolicy_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
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

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
