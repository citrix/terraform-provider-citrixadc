package hanode_routemonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// HanodeRoutemonitorBindingResourceModel describes the resource data model.
type HanodeRoutemonitorBindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Hanodeid     types.Int64  `tfsdk:"hanode_id"`
	Netmask      types.String `tfsdk:"netmask"`
	Routemonitor types.String `tfsdk:"routemonitor"`
}

func (r *HanodeRoutemonitorBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hanode_routemonitor_binding resource.",
			},
			"hanode_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number that uniquely identifies the local node. The ID of the local node is always 0.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
		},
	}
}

func hanode_routemonitor_bindingGetThePayloadFromtheConfig(ctx context.Context, data *HanodeRoutemonitorBindingResourceModel) ha.Hanoderoutemonitorbinding {
	tflog.Debug(ctx, "In hanode_routemonitor_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	hanode_routemonitor_binding := ha.Hanoderoutemonitorbinding{}
	if !data.Hanodeid.IsNull() {
		hanode_routemonitor_binding.Id = utils.IntPtr(int(data.Hanodeid.ValueInt64()))
	}
	if !data.Netmask.IsNull() {
		hanode_routemonitor_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Routemonitor.IsNull() {
		hanode_routemonitor_binding.Routemonitor = data.Routemonitor.ValueString()
	}

	return hanode_routemonitor_binding
}

func hanode_routemonitor_bindingSetAttrFromGet(ctx context.Context, data *HanodeRoutemonitorBindingResourceModel, getResponseData map[string]interface{}) *HanodeRoutemonitorBindingResourceModel {
	tflog.Debug(ctx, "In hanode_routemonitor_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hanodeid = types.Int64Value(intVal)
		}
	} else {
		data.Hanodeid = types.Int64Null()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["routemonitor"]; ok && val != nil {
		data.Routemonitor = types.StringValue(val.(string))
	} else {
		data.Routemonitor = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("hanode_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Hanodeid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
