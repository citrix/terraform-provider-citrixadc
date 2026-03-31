package vxlanvlanmap_vxlan_binding

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

// VxlanvlanmapVxlanBindingResourceModel describes the resource data model.
type VxlanvlanmapVxlanBindingResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Vlan  types.List   `tfsdk:"vlan"`
	Vxlan types.Int64  `tfsdk:"vxlan"`
}

func (r *VxlanvlanmapVxlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlanvlanmap_vxlan_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the mapping table.",
			},
			"vlan": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The vlan id or the range of vlan ids in the on-premise network.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VXLAN assigned to the vlan inside the cloud.",
			},
		},
	}
}

func vxlanvlanmap_vxlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VxlanvlanmapVxlanBindingResourceModel) network.Vxlanvlanmapvxlanbinding {
	tflog.Debug(ctx, "In vxlanvlanmap_vxlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vxlanvlanmap_vxlan_binding := network.Vxlanvlanmapvxlanbinding{}
	if !data.Name.IsNull() {
		vxlanvlanmap_vxlan_binding.Name = data.Name.ValueString()
	}
	if !data.Vxlan.IsNull() {
		vxlanvlanmap_vxlan_binding.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return vxlanvlanmap_vxlan_binding
}

func vxlanvlanmap_vxlan_bindingSetAttrFromGet(ctx context.Context, data *VxlanvlanmapVxlanBindingResourceModel, getResponseData map[string]interface{}) *VxlanvlanmapVxlanBindingResourceModel {
	tflog.Debug(ctx, "In vxlanvlanmap_vxlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vxlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
