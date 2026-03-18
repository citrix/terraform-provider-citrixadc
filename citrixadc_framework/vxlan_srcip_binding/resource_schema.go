package vxlan_srcip_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VxlanSrcipBindingResourceModel describes the resource data model.
type VxlanSrcipBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Vxlanid types.Int64  `tfsdk:"vxlanid"`
	Srcip   types.String `tfsdk:"srcip"`
}

func (r *VxlanSrcipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlan_srcip_binding resource.",
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The source IP address to use in outgoing vxlan packets.",
			},
		},
	}
}

func vxlan_srcip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VxlanSrcipBindingResourceModel) network.Vxlansrcipbinding {
	tflog.Debug(ctx, "In vxlan_srcip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vxlan_srcip_binding := network.Vxlansrcipbinding{}
	if !data.Vxlanid.IsNull() {
		vxlan_srcip_binding.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Srcip.IsNull() {
		vxlan_srcip_binding.Srcip = data.Srcip.ValueString()
	}

	return vxlan_srcip_binding
}

func vxlan_srcip_bindingSetAttrFromGet(ctx context.Context, data *VxlanSrcipBindingResourceModel, getResponseData map[string]interface{}) *VxlanSrcipBindingResourceModel {
	tflog.Debug(ctx, "In vxlan_srcip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vxlanid = types.Int64Null()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vxlanid:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vxlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("srcip:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Srcip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
